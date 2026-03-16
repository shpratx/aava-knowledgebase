# Best Practices & Standards for Database Design

---

## 1. Normalization Standards (ref: DA-BP-013)

### Normalization Requirements

| Normal Form | Requirement | Banking Rationale |
|---|---|---|
| 1NF | Atomic values; no repeating groups | Each transaction field is a single value |
| 2NF | No partial dependencies on composite keys | Transaction attributes depend on full transaction ID |
| 3NF (minimum) | No transitive dependencies | Customer name in customer table, not duplicated in transactions |
| BCNF | Every determinant is a candidate key | Recommended for reference data (currency, country) |

### When to Denormalize (with justification)

| Scenario | Approach | Justification Required |
|---|---|---|
| Read-heavy dashboards | Materialized views or read replicas | Query pattern, refresh frequency, staleness tolerance |
| Transaction history | Pre-computed aggregates | Aggregation logic, reconciliation process |
| Reporting / analytics | Star schema in data warehouse | ETL process, data freshness SLA |
| Search / full-text | Elasticsearch alongside relational | Sync mechanism, consistency model |
| Caching | Redis/Memcached | TTL, invalidation strategy |

**Denormalization Decision Record:**

| Field | Description |
|---|---|
| Table/View | What is being denormalized |
| Source tables | Normalized source(s) |
| Justification | Performance requirement |
| Consistency model | How denormalized data stays in sync |
| Staleness tolerance | Maximum acceptable lag |
| Reconciliation | How to detect and fix inconsistencies |
| Approved by | Architect + DBA |

### Banking Data Model

```sql
-- Core tables (3NF)
customers (customer_id PK, first_name, last_name, date_of_birth, kyc_status, created_at)
customer_addresses (address_id PK, customer_id FK, address_type, line1, line2, city, postal_code, country, is_primary)
customer_contacts (contact_id PK, customer_id FK, contact_type, value, is_verified, is_primary)

accounts (account_id PK, customer_id FK, account_type, currency, status, opened_at, closed_at)
account_balances (account_id PK, available_balance, ledger_balance, hold_amount, last_updated)

transactions (transaction_id PK, account_id FK, transaction_type, amount, currency,
             direction, status, value_date, booking_date, reference, correlation_id, created_at)
transaction_details (transaction_id PK FK, counterparty_account, counterparty_name,
                     remittance_info, exchange_rate, original_amount, original_currency)
```

### Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Tables | Plural, snake_case | customer_addresses, account_balances |
| Columns | Singular, snake_case | first_name, account_id, created_at |
| Primary keys | {singular_table}_id | customer_id, transaction_id |
| Foreign keys | Referenced table's PK name | customer_id (in accounts table) |
| Indexes | idx_{table}_{columns} | idx_transactions_account_id_created_at |
| Constraints | {type}_{table}_{columns} | uk_customers_email, chk_accounts_status |
| Timestamps | _at suffix (UTC) | created_at, updated_at, deleted_at |
| Booleans | is_ or has_ prefix | is_active, is_verified, has_mfa |
| Amounts | Explicit name + separate currency | amount + currency, not just "value" |

### Data Type Standards

| Data | Type | Rationale |
|---|---|---|
| Monetary amounts | DECIMAL(19,4) or NUMERIC(19,4) | Never FLOAT/DOUBLE — precision loss causes financial discrepancies |
| Currency codes | CHAR(3) | ISO 4217 |
| Dates | DATE or TIMESTAMPTZ | Always UTC; convert for display |
| UUIDs | UUID native type or CHAR(36) | External-facing IDs; never expose sequential integers |
| Account numbers | VARCHAR with encryption | Column-level encryption for PII |
| Status fields | VARCHAR with CHECK or lookup table | Explicit allowed values |
| Country codes | CHAR(2) | ISO 3166-1 alpha-2 |
| Phone numbers | VARCHAR(20) | E.164 format |
| Email | VARCHAR(254) | RFC 5321 max |
| IP addresses | INET or VARCHAR(45) | IPv4 and IPv6 |

---

## 2. Row-Level Security (ref: DA-BP-014)

### Implementation Approaches

| Approach | Database | Implementation |
|---|---|---|
| Native RLS | PostgreSQL | CREATE POLICY; ALTER TABLE ENABLE ROW LEVEL SECURITY |
| Virtual Private Database | Oracle | DBMS_RLS.ADD_POLICY |
| Application-enforced | All | WHERE clauses in repository layer |
| View-based | All | Filtered views per role/tenant |

### Banking RLS Patterns (PostgreSQL)

```sql
ALTER TABLE accounts ENABLE ROW LEVEL SECURITY;
ALTER TABLE transactions ENABLE ROW LEVEL SECURITY;

-- Customers see only their own accounts
CREATE POLICY customer_accounts ON accounts FOR ALL TO banking_app_role
  USING (customer_id = current_setting('app.current_customer_id')::uuid);

-- Customers see only their own transactions
CREATE POLICY customer_transactions ON transactions FOR ALL TO banking_app_role
  USING (account_id IN (
    SELECT account_id FROM accounts
    WHERE customer_id = current_setting('app.current_customer_id')::uuid));

-- Branch staff see branch customers
CREATE POLICY branch_staff ON accounts FOR SELECT TO branch_staff_role
  USING (branch_id = current_setting('app.current_branch_id')::uuid);

-- Compliance: full access (audit-logged)
CREATE POLICY compliance_access ON accounts FOR SELECT TO compliance_role
  USING (true);
```

### RLS by Role

| Role | Scope | Policy |
|---|---|---|
| Customer (API) | Own accounts only | Filter by customer_id |
| Branch teller | Branch customers | Filter by branch_id |
| Relationship manager | Portfolio customers | Filter by rm_id |
| Operations | All (read-only mostly) | Read access; restricted writes |
| Compliance/Audit | All (read-only) | Full read; all access logged |
| DBA | Schema only; no business data | No RLS bypass |

---

## 3. Database Encryption (ref: DA-BP-015)

### Encryption Layers

| Layer | Protects | Standard |
|---|---|---|
| TDE | Data files at rest | AES-256; mandatory for all databases |
| Column-level | Specific sensitive columns | AES-256-GCM; for PII and financial data |
| Connection | Data in transit | TLS 1.2+; reject unencrypted connections |
| Backup | Database backups | AES-256; key stored separately |
| Log | Logs with sensitive data | Encrypted storage |

### Column-Level Encryption Matrix

| Column | Encrypt? | Searchable? |
|---|---|---|
| Account number | Yes (AES-256-GCM) | Via token or encrypted index |
| Card number (PAN) | Yes (AES-256-GCM + tokenization) | Via token lookup |
| National ID / SSN | Yes (AES-256-GCM) | Via encrypted index |
| Customer name | Yes (AES-256-GCM) | Via search index |
| Email | Yes (AES-256-GCM) | Via hash for exact match |
| Phone | Yes (AES-256-GCM) | Via hash for exact match |
| Date of birth | Yes (AES-256-GCM) | Via encrypted index |
| Address | Yes (AES-256-GCM) | No (use search index) |
| Transaction amount | No (TDE covers) | Standard indexing |
| Account balance | No (TDE covers) | Standard indexing |

### Key Management

| Standard | Requirement |
|---|---|
| Storage | HSM (FIPS 140-2 Level 3) for master keys |
| Hierarchy | Master Key (HSM) → DEK (per table/column) → data |
| Rotation | DEK: 90 days; Master: annually |
| Separation | Per environment; per classification |
| Backup | HSM backup to separate HSM; documented recovery |
| Access | 2-person rule for master key operations |
| Audit | All key operations logged; quarterly inventory |

---

## 4. CQRS and Read/Write Separation (ref: DA-BP-016)

### Patterns

| Pattern | Implementation | Use Case |
|---|---|---|
| Read replicas | Primary writes; replicas read | Balance queries from replica |
| CQRS separate stores | Write normalized; read denormalized | Transactions → PostgreSQL; search → Elasticsearch |
| CQRS + event sourcing | Events as source of truth; projections for queries | Transfer lifecycle events → balance view |
| Materialized views | DB-maintained denormalized views | Daily balance summary |

### Standards

| Standard | Requirement |
|---|---|
| Write path | Always primary; strong consistency; transactional |
| Read path | Replica or read store; eventual consistency acceptable for most |
| Replication lag | Monitor; alert if > 1 second for financial data |
| Consistency | Balance inquiry: primary (real-time); history: replica OK |
| Failover | Replica promoted automatically |
| Connection pooling | Separate pools for read/write |

### Banking CQRS Map

| Write Model | Read Model | Sync |
|---|---|---|
| accounts (PostgreSQL) | account_dashboard (Redis) | CDC → Kafka |
| transactions (PostgreSQL) | transaction_search (Elasticsearch) | CDC → Kafka |
| transfers (PostgreSQL) | transfer_status (Redis) | Domain events |
| customers (PostgreSQL) | customer_360 (MongoDB) | CDC → Kafka |

### Connection Pool Standards

| Parameter | Write | Read |
|---|---|---|
| Min connections | 5 | 10 |
| Max connections | 20 | 50 |
| Connection timeout | 5s | 5s |
| Idle timeout | 10 min | 10 min |
| Max lifetime | 30 min | 30 min |
| Leak detection | 60s | 60s |

## 5. Database Audit Triggers (ref: DA-BP-017)

### Audit Trigger Standards

Every table containing sensitive data (Confidential/Restricted) must have audit triggers that capture changes.

### Audit Log Table Schema

```sql
CREATE TABLE audit_log (
  audit_id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  table_name      VARCHAR(100) NOT NULL,
  record_id       VARCHAR(100) NOT NULL,
  operation       VARCHAR(10) NOT NULL CHECK (operation IN ('INSERT','UPDATE','DELETE')),
  old_values      JSONB,
  new_values      JSONB,
  changed_columns TEXT[],
  user_id         VARCHAR(100),
  session_id      VARCHAR(100),
  source_ip       INET,
  correlation_id  VARCHAR(100),
  application     VARCHAR(100),
  timestamp       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_audit_log_table_record ON audit_log (table_name, record_id);
CREATE INDEX idx_audit_log_timestamp ON audit_log (timestamp);
CREATE INDEX idx_audit_log_user ON audit_log (user_id);
CREATE INDEX idx_audit_log_correlation ON audit_log (correlation_id);
```

### Generic Audit Trigger (PostgreSQL)

```sql
CREATE OR REPLACE FUNCTION audit_trigger_func()
RETURNS TRIGGER AS $$
BEGIN
  IF TG_OP = 'INSERT' THEN
    INSERT INTO audit_log (table_name, record_id, operation, new_values, user_id, 
                           session_id, source_ip, correlation_id, application)
    VALUES (TG_TABLE_NAME, NEW.id::text, 'INSERT', to_jsonb(NEW),
            current_setting('app.current_user_id', true),
            current_setting('app.current_session_id', true),
            current_setting('app.current_source_ip', true)::inet,
            current_setting('app.current_correlation_id', true),
            current_setting('app.current_application', true));
    RETURN NEW;
  ELSIF TG_OP = 'UPDATE' THEN
    INSERT INTO audit_log (table_name, record_id, operation, old_values, new_values,
                           changed_columns, user_id, session_id, source_ip, 
                           correlation_id, application)
    VALUES (TG_TABLE_NAME, NEW.id::text, 'UPDATE', to_jsonb(OLD), to_jsonb(NEW),
            akeys(hstore(NEW) - hstore(OLD)),
            current_setting('app.current_user_id', true),
            current_setting('app.current_session_id', true),
            current_setting('app.current_source_ip', true)::inet,
            current_setting('app.current_correlation_id', true),
            current_setting('app.current_application', true));
    RETURN NEW;
  ELSIF TG_OP = 'DELETE' THEN
    INSERT INTO audit_log (table_name, record_id, operation, old_values, user_id,
                           session_id, source_ip, correlation_id, application)
    VALUES (TG_TABLE_NAME, OLD.id::text, 'DELETE', to_jsonb(OLD),
            current_setting('app.current_user_id', true),
            current_setting('app.current_session_id', true),
            current_setting('app.current_source_ip', true)::inet,
            current_setting('app.current_correlation_id', true),
            current_setting('app.current_application', true));
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
```

### Tables Requiring Audit Triggers

| Table | Operations to Audit | Retention | Rationale |
|---|---|---|---|
| customers | INSERT, UPDATE, DELETE | 7 years | KYC/AML; GDPR processing records |
| customer_addresses | INSERT, UPDATE, DELETE | 7 years | Address change fraud detection |
| customer_contacts | INSERT, UPDATE, DELETE | 7 years | Contact change fraud detection |
| accounts | INSERT, UPDATE, DELETE | 7 years | Account lifecycle; regulatory |
| account_balances | UPDATE | 7 years | Balance reconciliation; dispute resolution |
| transactions | INSERT, UPDATE | 7 years | Financial audit trail |
| transfers | INSERT, UPDATE | 7 years | Transfer lifecycle; AML |
| beneficiaries | INSERT, UPDATE, DELETE | 7 years | Beneficiary fraud detection |
| cards | INSERT, UPDATE | 7 years | Card lifecycle; PCI-DSS |
| user_roles | INSERT, UPDATE, DELETE | 7 years | Access control audit; SOX |
| system_config | INSERT, UPDATE, DELETE | 7 years | Configuration change audit |
| consent_records | INSERT, UPDATE, DELETE | Duration + 3 years | GDPR consent evidence |

### Sensitive Data in Audit Logs

| Standard | Requirement |
|---|---|
| PII masking | Mask PII in audit log values (last 4 digits for account numbers, national IDs) |
| Password exclusion | Never store password hashes in audit log — exclude password columns |
| Card data exclusion | Never store full PAN, CVV, PIN in audit log — tokenize or mask |
| Encryption | Audit log table encrypted via TDE at minimum; column-level for PII fields |
| Immutability | Audit log table: no UPDATE or DELETE permissions for any application role |
| Separation | Audit log in separate schema/tablespace; separate access controls |

---

## 6. Parameterized Queries (ref: DA-BP-018)

### Standards

| Standard | Requirement |
|---|---|
| All queries parameterized | Every database query must use parameterized statements — no string concatenation |
| ORM usage | Use JPA/Hibernate named parameters or Spring Data query methods |
| Native queries | If native SQL is required, must use parameterized PreparedStatement |
| Dynamic queries | Use Criteria API or QueryDSL — never concatenate WHERE clauses |
| Stored procedures | Parameters must be typed and validated |
| Batch operations | Use batch parameterized statements |

### Implementation Examples

**Spring Data JPA (Correct):**
```java
// Repository method — automatically parameterized
List<Transaction> findByAccountIdAndStatusAndCreatedAtBetween(
    UUID accountId, TransactionStatus status, Instant from, Instant to);

// JPQL — parameterized
@Query("SELECT t FROM Transaction t WHERE t.accountId = :accountId AND t.amount > :minAmount")
List<Transaction> findLargeTransactions(@Param("accountId") UUID accountId, 
                                        @Param("minAmount") BigDecimal minAmount);

// Native query — parameterized
@Query(value = "SELECT * FROM transactions WHERE account_id = :accountId AND status = :status", 
       nativeQuery = true)
List<Transaction> findByAccountAndStatus(@Param("accountId") UUID accountId, 
                                          @Param("status") String status);
```

**Prohibited Patterns:**
```java
// NEVER — SQL injection vulnerability
String query = "SELECT * FROM accounts WHERE customer_id = '" + customerId + "'";

// NEVER — string concatenation in JPQL
String jpql = "SELECT t FROM Transaction t WHERE t.status = '" + status + "'";

// NEVER — dynamic WHERE clause via concatenation
String where = "WHERE 1=1";
if (status != null) where += " AND status = '" + status + "'";
```

**Dynamic Query (Correct — Criteria API):**
```java
CriteriaBuilder cb = entityManager.getCriteriaBuilder();
CriteriaQuery<Transaction> query = cb.createQuery(Transaction.class);
Root<Transaction> root = query.from(Transaction.class);

List<Predicate> predicates = new ArrayList<>();
predicates.add(cb.equal(root.get("accountId"), accountId));
if (status != null) predicates.add(cb.equal(root.get("status"), status));
if (fromDate != null) predicates.add(cb.greaterThanOrEqualTo(root.get("createdAt"), fromDate));

query.where(predicates.toArray(new Predicate[0]));
return entityManager.createQuery(query).getResultList();
```

---

## 7. Schema Migration Standards

### Migration Tool Standards

| Standard | Requirement |
|---|---|
| Tool | Flyway (recommended) or Liquibase |
| Version control | All migrations in version control alongside application code |
| Naming | V{version}__{description}.sql (e.g., V1.0.0__create_accounts_table.sql) |
| Idempotency | Migrations must be idempotent where possible (IF NOT EXISTS) |
| Rollback | Every migration must have a corresponding rollback script |
| Testing | Migrations tested in CI/CD before production |
| Review | Schema changes require DBA review |
| Backward compatibility | Migrations must be backward-compatible with previous application version |

### Migration Best Practices

| Practice | Standard |
|---|---|
| Additive preferred | Add columns/tables; avoid dropping or renaming |
| Column addition | Add as nullable first; backfill; then add NOT NULL if needed |
| Column removal | Deprecate in code first; remove column in later migration |
| Index creation | CREATE INDEX CONCURRENTLY (PostgreSQL) to avoid table locks |
| Large table changes | Use online DDL or pt-online-schema-change for zero-downtime |
| Data migration | Separate from schema migration; run as application-level task |
| Seed data | Reference data (currencies, countries) in migrations; business data via application |

---

## 8. Indexing Standards

### Index Strategy

| Index Type | When to Use | Banking Example |
|---|---|---|
| B-tree (default) | Equality and range queries | account_id, created_at ranges |
| Hash | Equality-only lookups | Exact match on correlation_id |
| GIN | Full-text search, JSONB queries | Search in transaction metadata |
| Partial | Subset of rows matching condition | Active accounts only: WHERE status = 'ACTIVE' |
| Composite | Multi-column queries | (account_id, created_at) for transaction history |
| Covering | Include all query columns | Avoid table lookup for frequent queries |
| Unique | Enforce uniqueness | (customer_id, account_type) for one account per type |

### Banking Index Recommendations

```sql
-- Transactions: most common query patterns
CREATE INDEX idx_transactions_account_created ON transactions (account_id, created_at DESC);
CREATE INDEX idx_transactions_correlation ON transactions (correlation_id);
CREATE INDEX idx_transactions_status ON transactions (status) WHERE status IN ('PENDING','PROCESSING');
CREATE INDEX idx_transactions_reference ON transactions (reference);

-- Accounts: lookup patterns
CREATE UNIQUE INDEX idx_accounts_customer_type ON accounts (customer_id, account_type) 
  WHERE status = 'ACTIVE';
CREATE INDEX idx_accounts_status ON accounts (status);

-- Audit log: investigation patterns
CREATE INDEX idx_audit_table_record_time ON audit_log (table_name, record_id, timestamp DESC);
CREATE INDEX idx_audit_user_time ON audit_log (user_id, timestamp DESC);
CREATE INDEX idx_audit_correlation ON audit_log (correlation_id);
```

### Index Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Index on every column | Write performance degradation; storage waste | Index only queried columns |
| No index on foreign keys | Slow JOINs and cascading deletes | Index all foreign keys |
| Missing composite index | Multiple single-column indexes less efficient | Composite index for multi-column WHERE |
| Wrong column order in composite | Index not used for leading column queries | Most selective column first; match query pattern |
| No partial indexes | Full index when only subset queried | Partial index for common filters |
| Never analyzing index usage | Unused indexes waste resources | Regular pg_stat_user_indexes review |

---

## 9. High Availability & Disaster Recovery

### Database HA Standards

| Standard | Requirement |
|---|---|
| Replication | Synchronous replication for financial data (zero data loss) |
| Failover | Automatic failover within 30 seconds |
| Read replicas | Minimum 2 read replicas for customer-facing services |
| Backup frequency | Continuous WAL archiving + daily full backup |
| Backup testing | Weekly automated restore test; quarterly manual DR drill |
| RPO | < 0 (synchronous replication) for Tier 1; < 15 minutes for Tier 2 |
| RTO | < 30 seconds (automatic failover) for Tier 1; < 1 hour for Tier 2 |
| Geographic redundancy | Multi-AZ minimum; multi-region for Tier 1 critical |
| Connection failover | Application connection string supports automatic failover |

### Backup Standards

| Backup Type | Frequency | Retention | Testing |
|---|---|---|---|
| WAL / transaction log | Continuous | 7 days | Included in PITR test |
| Full backup | Daily | 30 days | Weekly restore test |
| Monthly archive | Monthly | 7 years | Quarterly restore test |
| Pre-migration snapshot | Before every schema change | 30 days | Verified before migration |

---

## 10. Performance Standards

### Query Performance

| Metric | Target | Measurement |
|---|---|---|
| Simple lookup (by PK) | < 5ms | Query profiling |
| Indexed query | < 50ms | Query profiling |
| Complex join (3+ tables) | < 100ms | Query profiling |
| Aggregation query | < 500ms | Query profiling |
| Full table scan | Prohibited on tables > 10K rows | Query plan analysis |
| N+1 queries | Prohibited | ORM query logging |

### Performance Best Practices

| Practice | Standard |
|---|---|
| Query plan analysis | EXPLAIN ANALYZE for all new queries; review in code review |
| N+1 prevention | Use JOIN FETCH, @EntityGraph, or batch fetching |
| Pagination | All collection queries must be paginated (max 100 per page) |
| Connection pooling | HikariCP (Spring Boot default); sized per service |
| Statement caching | Enable prepared statement caching |
| Slow query logging | Log queries > 1 second; alert on queries > 5 seconds |
| Table statistics | Auto-vacuum and auto-analyze enabled; manual ANALYZE after bulk operations |
| Partitioning | Partition large tables by date (transactions: monthly partitions) |

### Table Partitioning for Banking

```sql
-- Partition transactions by month
CREATE TABLE transactions (
  transaction_id UUID NOT NULL,
  account_id UUID NOT NULL,
  amount DECIMAL(19,4) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  ...
) PARTITION BY RANGE (created_at);

CREATE TABLE transactions_2026_01 PARTITION OF transactions
  FOR VALUES FROM ('2026-01-01') TO ('2026-02-01');
CREATE TABLE transactions_2026_02 PARTITION OF transactions
  FOR VALUES FROM ('2026-02-01') TO ('2026-03-01');
-- Auto-create future partitions via pg_partman or scheduled job
```

---

## 11. Common Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| FLOAT for money | Precision loss; financial discrepancies | DECIMAL(19,4) or NUMERIC(19,4) |
| Sequential integer IDs exposed | Enumeration attack; data scraping | UUID for external-facing IDs |
| No encryption for PII columns | Data breach exposure | Column-level AES-256-GCM |
| String concatenation in queries | SQL injection | Parameterized queries exclusively |
| No audit triggers on sensitive tables | Failed audits; no forensic trail | Audit triggers on all Confidential/Restricted tables |
| Shared database across services | Tight coupling; schema change breaks others | Database per service |
| No RLS for multi-tenant data | Cross-customer data leakage | Row-level security policies |
| No index on foreign keys | Slow joins and cascading operations | Index all foreign keys |
| Full table scans on large tables | Performance degradation | Proper indexing; partitioning |
| No backup testing | Backups fail when needed | Weekly automated restore tests |
| Storing passwords in plain text | Catastrophic breach | bcrypt/Argon2id; never reversible encryption |
| No connection pooling | Connection exhaustion | HikariCP with proper sizing |
| Denormalization without justification | Data inconsistency; maintenance burden | Document justification; reconciliation process |
| No schema migration tool | Manual DDL; inconsistent environments | Flyway/Liquibase in CI/CD |
