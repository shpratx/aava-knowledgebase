# NoSQL Database Design — Standards, Best Practices & Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Part A: Standards & Best Practices

### 1. NoSQL Database Selection

#### When to Use NoSQL in Banking

| Use Case | NoSQL Type | Database | Why Not Relational |
|---|---|---|---|
| Session/cache store | Key-Value | Redis, DynamoDB | Sub-millisecond reads; TTL-based expiry; ephemeral data |
| Customer 360 view | Document | MongoDB, Cosmos DB | Flexible schema; aggregated view from multiple sources |
| Transaction event store | Wide-Column | Cassandra, ScyllaDB | High write throughput; time-series partitioning; linear scalability |
| Fraud graph analysis | Graph | Neo4j, Neptune | Relationship traversal; pattern detection across entities |
| Full-text search | Search | Elasticsearch, OpenSearch | Transaction search; log analysis; fuzzy matching |
| Real-time analytics | Time-Series | TimescaleDB, InfluxDB | Metrics, monitoring, transaction volume analytics |
| Configuration/feature flags | Key-Value / Document | Redis, Consul, DynamoDB | Fast reads; simple structure; distributed |

#### When NOT to Use NoSQL in Banking

| Scenario | Use Relational Instead | Reason |
|---|---|---|
| Core account balances | PostgreSQL / Oracle | ACID transactions; strong consistency; referential integrity |
| Financial ledger | PostgreSQL / Oracle | Double-entry bookkeeping requires transactions |
| Regulatory reporting source of truth | PostgreSQL | Joins, aggregations, complex queries |
| Audit trail (primary) | PostgreSQL + immutable store | Referential integrity; regulatory retention |
| Any data requiring multi-record ACID | PostgreSQL | NoSQL trades consistency for scalability |

### 2. Data Modeling Standards

#### Document Database (MongoDB / Cosmos DB)

**Embedding vs. Referencing:**

| Pattern | When to Use | Banking Example |
|---|---|---|
| Embed | Data always accessed together; 1:1 or 1:few; child doesn't exist independently | Customer + addresses; Transfer + transfer lines |
| Reference | Data accessed independently; 1:many (unbounded); many:many; large sub-documents | Customer → Accounts (reference); Account → Transactions (reference) |
| Hybrid | Embed summary, reference detail | Account embeds last 5 transactions; references full history |

**Document Design Rules:**

| Rule | Standard |
|---|---|
| Document size | Keep under 1MB (MongoDB 16MB limit; aim for < 1MB for performance) |
| Nesting depth | Maximum 3 levels deep |
| Array growth | Avoid unbounded arrays — use reference pattern for growing collections |
| Schema version | Include `schemaVersion` field in every document for evolution |
| Timestamps | ISO 8601 UTC; include `createdAt`, `updatedAt` |
| IDs | Use UUID or domain-specific ID (not auto-increment ObjectId for external use) |
| Null handling | Omit field rather than store null (saves storage; simplifies queries) |

**Banking Document Example:**
```json
{
  "_id": "CUS-001234",
  "schemaVersion": 2,
  "firstName": "John",
  "lastName": "Doe",
  "email": "j***@example.com",
  "kycStatus": "VERIFIED",
  "riskScore": "LOW",
  "addresses": [
    {
      "type": "HOME",
      "line1": "encrypted:AES256:...",
      "city": "London",
      "country": "GB",
      "isPrimary": true
    }
  ],
  "accountSummaries": [
    {
      "accountId": "ACC-001234",
      "type": "CHECKING",
      "currency": "GBP",
      "status": "ACTIVE"
    }
  ],
  "createdAt": "2024-01-15T10:00:00Z",
  "updatedAt": "2026-03-15T15:30:00Z"
}
```

#### Key-Value Store (Redis / DynamoDB)

**Key Design Standards:**

| Standard | Implementation |
|---|---|
| Key naming | `{entity}:{id}:{attribute}` — e.g., `session:USR-001:token`, `cache:ACC-001:balance` |
| Key length | Keep short (< 100 bytes) for memory efficiency |
| TTL | Always set TTL for cache/session data; never store without expiry |
| Namespacing | Prefix with service name to avoid collisions in shared clusters |
| No sensitive data in keys | Account numbers, PII must not appear in key names |

**Banking Key Patterns:**
```
session:{sessionId}              → session data (TTL: 15 min)
cache:balance:{accountId}        → cached balance (TTL: 30 sec)
rate-limit:{userId}:{endpoint}   → rate limit counter (TTL: 60 sec)
idempotency:{idempotencyKey}     → cached response (TTL: 24 hours)
feature-flag:{flagName}          → flag value (TTL: 5 min)
lock:transfer:{transferId}       → distributed lock (TTL: 30 sec)
```

#### Wide-Column Store (Cassandra / ScyllaDB)

**Partition Key Design:**

| Rule | Standard |
|---|---|
| Partition key | Must distribute data evenly; avoid hot partitions |
| Partition size | Target < 100MB per partition; < 100K rows |
| Clustering key | Determines sort order within partition; choose for query pattern |
| Query-first design | Design tables around query patterns, not entity relationships |
| Denormalization | Expected and necessary — duplicate data across tables for different query patterns |

**Banking Table Example:**
```cql
-- Transactions by account (query: get recent transactions for account)
CREATE TABLE transactions_by_account (
    account_id    TEXT,
    transaction_date DATE,
    transaction_id TIMEUUID,
    type          TEXT,
    amount        DECIMAL,
    currency      TEXT,
    status        TEXT,
    description   TEXT,
    PRIMARY KEY ((account_id), transaction_date, transaction_id)
) WITH CLUSTERING ORDER BY (transaction_date DESC, transaction_id DESC)
  AND default_time_to_live = 220752000;  -- 7 years in seconds

-- Transactions by correlation ID (query: trace a transfer across services)
CREATE TABLE transactions_by_correlation (
    correlation_id TEXT,
    service        TEXT,
    event_time     TIMESTAMP,
    event_type     TEXT,
    payload        TEXT,
    PRIMARY KEY ((correlation_id), event_time, service)
) WITH CLUSTERING ORDER BY (event_time ASC);
```

#### Graph Database (Neo4j / Neptune)

**Banking Graph Patterns:**

| Pattern | Nodes | Relationships | Use Case |
|---|---|---|---|
| Fraud ring detection | Customer, Account, Transaction, Device, IP | OWNS, TRANSFERS_TO, USES_DEVICE, FROM_IP | Detect connected fraud networks |
| Beneficial ownership | Person, Company, Trust | OWNS, CONTROLS, DIRECTS | KYC/AML beneficial ownership chains |
| Payment network | Account, Bank | PAYS, RECEIVES | Payment flow analysis; correspondent banking |

### 3. Consistency Models

| Model | Guarantee | Use Case | Banking Example |
|---|---|---|---|
| Strong consistency | Read always returns latest write | Financial balances, ledger | Account balance (use relational) |
| Eventual consistency | Read may return stale data temporarily | Caches, search indexes, analytics | Transaction search; customer 360 |
| Read-your-writes | User sees their own writes immediately | Session data, user preferences | Customer sees their transfer immediately after submission |
| Bounded staleness | Read is at most N seconds behind | Dashboards, reporting | Account dashboard (max 5-second lag) |
| Causal consistency | Causally related operations seen in order | Event processing | Transfer events processed in causal order |

**Banking Consistency Rules:**

| Data | Required Consistency | Implementation |
|---|---|---|
| Account balance | Strong | Relational DB (not NoSQL) |
| Transaction history (query) | Eventual (< 5 sec) | NoSQL read model; CDC sync |
| Session data | Read-your-writes | Redis with read-after-write |
| Search index | Eventual (< 10 sec) | Elasticsearch; async indexing |
| Cache | Eventual (TTL-based) | Redis with appropriate TTL |
| Audit events | Eventual (< 5 sec) | Kafka → NoSQL/Elasticsearch |

### 4. Security Standards

#### Encryption

| Layer | Standard | Implementation |
|---|---|---|
| At rest | AES-256 | Database-native encryption (MongoDB encryption at rest, DynamoDB SSE, Cassandra TDE) |
| In transit | TLS 1.2+ | All client-to-database and node-to-node communication |
| Field-level | AES-256-GCM | PII fields encrypted at application layer before storage |
| Key management | HSM / KMS | AWS KMS, Azure Key Vault, HashiCorp Vault — never store keys with data |

#### Access Control

| Standard | Implementation |
|---|---|
| Authentication | Database-native auth (SCRAM-SHA-256 for MongoDB; IAM for DynamoDB; internal auth for Cassandra) |
| Authorization | Role-based; least privilege; separate read/write roles |
| Network | Database not internet-accessible; VPC/private subnet only; security groups restrict to app subnets |
| Credentials | Vault-managed; rotated every 90 days; per-service credentials |
| Audit | All access logged; admin operations logged with user, action, timestamp |

#### Data Classification in NoSQL

| Classification | Encryption | Access | Audit | Masking (non-prod) |
|---|---|---|---|---|
| Restricted | Field-level AES-256 + at-rest | MFA + role + resource | All access logged | Full anonymization |
| Confidential | Field-level AES-256 + at-rest | Role + resource | Modifications logged | Tokenization |
| Internal | At-rest encryption | Role-based | Auth events | Recommended |
| Public | At-rest encryption | Standard | Minimal | Not required |

### 5. Performance Standards

| Metric | Target | Measurement |
|---|---|---|
| Read latency (p95) | < 10ms (cache/KV); < 50ms (document/wide-column) | Database metrics |
| Write latency (p95) | < 20ms (KV); < 100ms (document/wide-column) | Database metrics |
| Query latency (p95) | < 100ms (indexed); < 500ms (aggregation) | Application metrics |
| Throughput | Per SLA; typically 1K-100K ops/sec depending on type | Load testing |
| Connection pool | Sized per service; monitored for exhaustion | Connection metrics |

#### Performance Best Practices

| Practice | Standard |
|---|---|
| Index strategy | Index all query patterns; avoid full collection scans |
| Query optimization | Project only needed fields; avoid returning entire documents |
| Pagination | Cursor-based for large result sets; never skip-based for large offsets |
| Connection pooling | Use driver connection pool; size appropriately; monitor |
| Read preference | Read from secondaries/replicas for non-critical reads (eventual consistency acceptable) |
| Write concern | Majority write concern for important data; acknowledge for ephemeral |
| Batch operations | Use bulk writes for high-volume inserts; batch reads where possible |
| TTL | Set TTL on all ephemeral data (sessions, caches, rate limits) |
| Compaction | Monitor and tune compaction (Cassandra); compact collections (MongoDB) |

### 6. Schema Evolution

| Strategy | Implementation |
|---|---|
| Schema version field | Every document includes `schemaVersion` integer |
| Backward compatibility | New code reads old schema; old code ignores new fields |
| Lazy migration | Migrate documents on read (update to new schema when accessed) |
| Batch migration | Background job migrates documents in batches (for large changes) |
| Default values | New required fields have defaults for old documents |
| Never remove fields | Deprecate (stop writing); remove after all documents migrated |
| Migration testing | Test migration with production-volume data in staging |

### 7. High Availability & Disaster Recovery

| Standard | Implementation |
|---|---|
| Replication | Minimum 3 replicas (MongoDB replica set; Cassandra RF=3; Redis Sentinel/Cluster) |
| Multi-AZ | Replicas across availability zones |
| Multi-region | For Tier 1 data; active-passive or active-active depending on consistency needs |
| Backup | Daily snapshots + continuous oplog/WAL; encrypted; stored separately |
| Backup testing | Weekly automated restore test |
| RTO | Tier 1: < 30 min (automatic failover); Tier 2: < 1 hour |
| RPO | Tier 1: < 1 min (synchronous replication); Tier 2: < 15 min |
| Failover | Automatic for replica set/cluster; tested quarterly |

### 8. Monitoring & Observability

| Metric | Alert Threshold | Tool |
|---|---|---|
| Replication lag | > 5 seconds | Database metrics / Datadog |
| Connection pool utilization | > 80% | Application metrics |
| Query latency (p95) | > 2x baseline | APM |
| Disk usage | > 80% | Infrastructure monitoring |
| Oplog/WAL size | Approaching retention limit | Database metrics |
| Hot partition (Cassandra/DynamoDB) | > 10x average partition size | Database metrics |
| Cache hit rate (Redis) | < 90% | Redis metrics |
| Eviction rate (Redis) | > 0 (unexpected) | Redis metrics |
| Slow queries | > 100ms | Database profiler |

---

## Part B: Guardrails

### 9. Data Model Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-DG-001 | NoSQL must not be used as the primary store for financial balances or ledger data — use relational for ACID | Reject if NoSQL proposed for balance/ledger |
| NQ-DG-002 | Every document must include a `schemaVersion` field for evolution | Reject if schema version is missing |
| NQ-DG-003 | Document size must be kept under 1MB; unbounded arrays are prohibited | Reject if unbounded array growth is possible |
| NQ-DG-004 | Partition keys (Cassandra/DynamoDB) must distribute data evenly — hot partition analysis required | Reject if partition key creates hot spots |
| NQ-DG-005 | Key-Value keys must not contain sensitive data (account numbers, PII) | Reject if PII in key names |
| NQ-DG-006 | All ephemeral data (sessions, caches, rate limits) must have TTL configured | Reject if TTL is missing on ephemeral data |
| NQ-DG-007 | Data model must be designed query-first — document the access patterns before designing schema | Reject if no access pattern documentation |
| NQ-DG-008 | Monetary amounts must use exact decimal types (Decimal128 in MongoDB; DECIMAL in Cassandra) — never floating point | Reject if float/double used for money |

### 10. Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-SG-001 | All NoSQL databases must have encryption at rest enabled | Reject if at-rest encryption is disabled |
| NQ-SG-002 | All client-to-database and node-to-node communication must use TLS 1.2+ | Reject if unencrypted communication |
| NQ-SG-003 | PII fields must be encrypted at the application layer (field-level) before storage | Reject if PII stored unencrypted in NoSQL |
| NQ-SG-004 | Database must not be accessible from the internet — VPC/private subnet only | Reject if database is publicly accessible |
| NQ-SG-005 | Authentication must be enabled — no anonymous access | Reject if auth is disabled |
| NQ-SG-006 | Each service must have its own database credentials with least-privilege permissions | Reject if shared credentials or excessive permissions |
| NQ-SG-007 | Credentials must be stored in vault — not in code, config files, or environment variables | Reject if credentials outside vault |
| NQ-SG-008 | Production data must not be used in non-production — anonymize/mask for lower environments | Reject if production data in non-prod |
| NQ-SG-009 | Admin/DBA access must be JIT with approval and audit logging | Flag if standing admin access exists |
| NQ-SG-010 | No default passwords or default admin accounts in production | Reject if defaults exist |

### 11. Consistency & Reliability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-CG-001 | Consistency model must be documented per collection/table — strong, eventual, or bounded staleness | Reject if consistency model is undocumented |
| NQ-CG-002 | Financial data requiring strong consistency must use relational database, not NoSQL | Reject if strong consistency needed but NoSQL chosen |
| NQ-CG-003 | Write concern must be "majority" for important data (MongoDB); RF=3 with CL=QUORUM for Cassandra | Reject if write concern is weaker than majority for important data |
| NQ-CG-004 | Replication factor must be minimum 3 across availability zones | Reject if RF < 3 or single-AZ |
| NQ-CG-005 | Eventual consistency lag must be monitored and alerted — max acceptable lag documented per use case | Reject if no lag monitoring |
| NQ-CG-006 | Idempotent writes must be implemented for all event consumers writing to NoSQL | Reject if consumer writes are not idempotent |

### 12. Performance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-PG-001 | All query patterns must be indexed — full collection/table scans are prohibited on collections > 10K documents | Reject if unindexed query on large collection |
| NQ-PG-002 | Queries must project only needed fields — no returning entire documents when subset is needed | Flag if full document returned unnecessarily |
| NQ-PG-003 | Pagination must be cursor-based for large result sets — skip-based pagination prohibited for offset > 1000 | Reject if skip-based with large offsets |
| NQ-PG-004 | Connection pooling must be configured with appropriate limits and monitoring | Reject if no connection pool configured |
| NQ-PG-005 | Slow query logging must be enabled — alert on queries > 100ms | Flag if slow query logging is disabled |
| NQ-PG-006 | Cache (Redis) must have eviction policy configured — no unbounded memory growth | Reject if no eviction policy |
| NQ-PG-007 | Read-heavy workloads must use read replicas/secondaries — not primary for all reads | Flag if all reads go to primary |

### 13. Operational Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-OG-001 | Backups must be configured: daily snapshots + continuous replication; encrypted; stored separately | Reject if no backup configured |
| NQ-OG-002 | Backup restore must be tested quarterly | Reject if not tested in last quarter |
| NQ-OG-003 | Automatic failover must be configured and tested for Tier 1 data | Reject if no automatic failover |
| NQ-OG-004 | Schema migrations must be backward-compatible — new code reads old schema | Reject if migration breaks backward compatibility |
| NQ-OG-005 | Monitoring must cover: replication lag, connection pool, query latency, disk usage, hot partitions | Reject if monitoring is incomplete |
| NQ-OG-006 | Capacity planning must account for data growth over retention period | Flag if no capacity projection |
| NQ-OG-007 | Data retention/TTL must align with regulatory requirements (7 years financial, 5 years AML) | Reject if retention is below regulatory minimum |

### 14. Audit & Compliance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NQ-AG-001 | All access to Confidential/Restricted data in NoSQL must be audit-logged | Reject if no audit logging for sensitive data |
| NQ-AG-002 | Audit logs must be stored separately from the NoSQL database being audited — not in the same cluster | Reject if audit logs co-located with audited data |
| NQ-AG-003 | NoSQL databases storing PII must be included in the GDPR processing register (Art. 30) | Reject if not in processing register |
| NQ-AG-004 | NoSQL databases must support data erasure/anonymization for GDPR Art. 17 compliance | Reject if erasure capability is not implemented |
| NQ-AG-005 | Data in NoSQL must be included in the data classification inventory with classification per collection/table | Reject if not in data classification inventory |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot proceed | NoSQL for financial ledger, no encryption, PII in keys, no TTL on sessions, float for money, no auth, public access |
| **Flag** | Address before release | No capacity planning, full documents returned, standing admin access, all reads to primary |
| **CI/CD Gate** | Automated | Schema version check, TTL verification, credential scanning |

---

## NoSQL Selection Decision Tree

```
Need ACID transactions across multiple records?
  → YES → Use Relational (PostgreSQL)
  → NO ↓

Need sub-millisecond reads for ephemeral data (sessions, cache)?
  → YES → Key-Value (Redis)
  → NO ↓

Need flexible schema for aggregated views?
  → YES → Document (MongoDB / Cosmos DB)
  → NO ↓

Need high write throughput for time-series/event data?
  → YES → Wide-Column (Cassandra / ScyllaDB)
  → NO ↓

Need relationship traversal (fraud detection, ownership chains)?
  → YES → Graph (Neo4j / Neptune)
  → NO ↓

Need full-text search across large datasets?
  → YES → Search (Elasticsearch / OpenSearch)
  → NO ↓

Default → Relational (PostgreSQL)
```

---

## Pre-Deployment Checklist

| # | Check |
|---|---|
| 1 | NoSQL type justified for use case (not used where relational is needed) |
| 2 | Data model designed query-first with documented access patterns |
| 3 | Schema version field in all documents |
| 4 | No unbounded arrays; documents < 1MB |
| 5 | Partition keys distribute evenly (no hot partitions) |
| 6 | Encryption at rest and in transit enabled |
| 7 | PII fields encrypted at application layer |
| 8 | Database not internet-accessible; auth enabled; least-privilege credentials |
| 9 | Consistency model documented per collection/table |
| 10 | All query patterns indexed; no full scans on large collections |
| 11 | TTL configured for all ephemeral data |
| 12 | Replication factor ≥ 3 across AZs; automatic failover configured |
| 13 | Backups configured, encrypted, and tested |
| 14 | Monitoring covers: replication lag, latency, connections, disk, hot partitions |
| 15 | Audit logging for sensitive data access |
| 16 | Included in data classification inventory and GDPR processing register |
| 17 | Erasure/anonymization capability for GDPR Art. 17 |
| 18 | Decimal types for monetary amounts (not float/double) |
