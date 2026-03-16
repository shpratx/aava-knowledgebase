# Guardrails for Database Design
### Banking Domain — Agentic Knowledge Base
### References: DA-GR-013, DA-GR-014, DA-GR-015, DA-GR-016, DA-GR-017, DA-GR-018, DA-GR-019

---

## 1. Column Encryption Guardrails (ref: DA-GR-013)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-EG-001 | All PII and sensitive columns must be encrypted at the column level using AES-256-GCM (ref: DA-GR-013) | Reject if PII columns are stored unencrypted |
| DB-EG-002 | TDE (Transparent Data Encryption) must be enabled on all databases as baseline — column-level encryption is additional for PII | Reject if TDE is not enabled |
| DB-EG-003 | Columns requiring encryption: account numbers, card numbers (PAN), national ID/SSN, customer name, email, phone, date of birth, address, biometric data | Reject if any of these columns are unencrypted |
| DB-EG-004 | Card data (PAN) must be tokenized in addition to encryption — full PAN stored only in PCI-DSS scoped systems | Reject if PAN is stored without tokenization outside CDE |
| DB-EG-005 | CVV, PIN, and track data must never be stored post-authorization — no column, no table, no log | Reject immediately — no exceptions |
| DB-EG-006 | Encryption keys must be managed via HSM (FIPS 140-2 Level 3) — not stored in database, config files, or application code | Reject if keys are stored alongside encrypted data |
| DB-EG-007 | Data Encryption Keys (DEK) must be rotated every 90 days; Master Keys annually | Flag if rotation schedule is not configured |
| DB-EG-008 | Separate encryption keys per environment — production keys must never be used in non-production | Reject if production keys are used in non-production |
| DB-EG-009 | Encryption must be verified — not just configured. Test by inspecting raw storage to confirm data is not readable | Reject if encryption verification test is not performed |
| DB-EG-010 | Encrypted columns must have documented searchability strategy (tokenization, hash index, encrypted index, or search via separate index) | Flag if no search strategy is documented for encrypted columns |

---

## 2. Database Credential Management Guardrails (ref: DA-GR-014)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-CG-001 | Database credentials must be stored in secure vaults (HashiCorp Vault, AWS Secrets Manager, Azure Key Vault) — not in config files, environment variables, or container images (ref: DA-GR-014) | Reject if credentials are stored outside secure vault |
| DB-CG-002 | Database credentials must not be hardcoded in source code | Reject — CI/CD secret scanning must block |
| DB-CG-003 | Database credentials must not be committed to version control | Reject if credentials detected in repository |
| DB-CG-004 | Database credentials must not appear in application logs | Reject if connection strings with passwords are logged |
| DB-CG-005 | Database credentials must be rotated at minimum every 90 days | Flag if rotation exceeds 90 days |
| DB-CG-006 | Each application/service must have its own database credentials — no shared credentials across services | Reject if multiple services share database credentials |
| DB-CG-007 | Each environment must have separate database credentials — no credential sharing across environments | Reject if same credentials used across environments |
| DB-CG-008 | Database credentials must follow least privilege — application accounts get only required permissions (SELECT, INSERT, UPDATE on specific tables) | Reject if application account has DBA or superuser privileges |
| DB-CG-009 | Credential retrieval from vault must be audited — log which service retrieved which credential and when | Flag if vault access is not audited |
| DB-CG-010 | Compromised credentials must be rotated immediately — incident response process must be defined | Reject if no credential compromise process exists |

---

## 3. Production Data Protection Guardrails (ref: DA-GR-015)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-PG-001 | Production data must not be used in non-production environments (dev, test, staging, QA) (ref: DA-GR-015) | Reject if production data is found in non-production environments |
| DB-PG-002 | Production database backups must not be restored to non-production environments without anonymization | Reject if production backup is restored to non-production without anonymization |
| DB-PG-003 | Production database access from non-production networks must be blocked at the network level | Reject if non-production networks can reach production databases |
| DB-PG-004 | Database connection strings for production must not be accessible from non-production environments | Reject if production connection strings are in non-production configs |
| DB-PG-005 | Data generation tools must be used for test data — synthetic data that mimics production patterns without using real customer data | Flag if no synthetic data generation process exists |
| DB-PG-006 | Any exception to this rule (e.g., production data needed for debugging) requires CISO + DPO approval, time-limited access, and full audit trail | Reject exception without required approvals |
| DB-PG-007 | Periodic scans must verify no production data exists in non-production environments | Process guardrail — quarterly scan |

---

## 4. Data Masking & Anonymization Guardrails (ref: DA-GR-016)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-MG-001 | Data masking or anonymization is required for all non-production environments (ref: DA-GR-016) | Reject if non-production environments contain unmasked sensitive data |
| DB-MG-002 | Masking must be irreversible for anonymization use cases — masked data must not be reversible to original | Reject if anonymization is reversible |
| DB-MG-003 | Masking must preserve data format and referential integrity — masked data must be usable for testing | Flag if masking breaks application functionality |
| DB-MG-004 | Masking rules must be defined per column based on data classification | Reject if no masking rules are defined |

**Masking Rules by Data Type:**

| Data Type | Masking Method | Example |
|---|---|---|
| Customer name | Fake name generation | "John Smith" → "Alice Johnson" |
| Email | Domain-preserving fake | "john@bank.com" → "user1234@bank.com" |
| Phone | Format-preserving random | "+1-555-123-4567" → "+1-555-987-6543" |
| National ID / SSN | Format-preserving random | "123-45-6789" → "987-65-4321" |
| Account number | Format-preserving random | "1234567890" → "9876543210" |
| Card number (PAN) | Format-preserving with valid Luhn | "4111111111111111" → "4532015112830366" |
| Date of birth | Random within realistic range | "1985-03-15" → "1990-07-22" |
| Address | Fake address generation | Real address → synthetic address |
| Transaction amount | Proportional randomization | Preserve distribution, randomize values |
| Balance | Randomized | Preserve realistic ranges |

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-MG-005 | Masking must be automated — manual masking is error-prone and unscalable | Flag if masking is manual |
| DB-MG-006 | Masking process must be tested — verify no sensitive data leaks through unmapped columns, foreign keys, or free-text fields | Reject if masking verification is not performed |
| DB-MG-007 | Free-text fields (notes, comments, remittance info) must be scrubbed or replaced — they often contain PII | Reject if free-text fields are not addressed in masking |
| DB-MG-008 | Masking must cover all data stores: relational databases, caches, search indexes, message queues, file storage | Flag if masking only covers primary database |

---

## 5. Database Access Control Guardrails (ref: DA-GR-017)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-AG-001 | Direct database access (SQL clients, CLI) is restricted to DBAs only (ref: DA-GR-017) | Reject if non-DBA personnel have direct database access |
| DB-AG-002 | Application access must be through application service accounts with least-privilege permissions | Reject if application uses DBA or superuser account |
| DB-AG-003 | Developer access to production databases is prohibited — use application logs and monitoring for debugging | Reject if developers have production database access |
| DB-AG-004 | DBA production access must use just-in-time (JIT) provisioning — no standing privileged access | Flag if DBA access is permanent rather than JIT |
| DB-AG-005 | All direct database access must be logged: who, when, what queries, from where | Reject if direct access is not logged |
| DB-AG-006 | DBA sessions on production must be recorded (session recording) | Flag if session recording is not enabled |
| DB-AG-007 | Emergency database access requires approval from on-call lead + security team; time-limited; fully audited | Reject if emergency access has no approval process |
| DB-AG-008 | Database accounts must be individual — no shared accounts (e.g., no shared "admin" account) | Reject if shared database accounts exist |
| DB-AG-009 | Database account permissions must be reviewed quarterly — remove unused accounts and excessive privileges | Process guardrail — quarterly access review |
| DB-AG-010 | Network access to databases must be restricted — databases must not be accessible from the internet; only from authorized application subnets | Reject if database is publicly accessible |

---

## 6. Schema Change Control Guardrails (ref: DA-GR-018)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-SG-001 | All schema changes require change control approval before execution (ref: DA-GR-018) | Reject if schema change has no approval |
| DB-SG-002 | Schema changes must be version-controlled using migration tools (Flyway/Liquibase) — no manual DDL in production | Reject if manual DDL is executed in production |
| DB-SG-003 | Schema changes must be reviewed by a DBA before approval | Reject if schema change has no DBA review |
| DB-SG-004 | Schema changes must be backward-compatible with the current application version — support rolling deployments | Reject if schema change breaks current application version |
| DB-SG-005 | Every schema migration must have a tested rollback script | Reject if no rollback script exists |
| DB-SG-006 | Destructive changes (DROP TABLE, DROP COLUMN, column type change) require additional approval from Architect + Business Owner | Reject if destructive change lacks additional approvals |
| DB-SG-007 | Schema changes must be tested in staging with production-equivalent data volume before production execution | Reject if schema change is untested in staging |
| DB-SG-008 | Schema changes must include impact assessment: affected queries, indexes, application code, downstream consumers | Reject if no impact assessment is documented |
| DB-SG-009 | Large table alterations must use online DDL (e.g., CREATE INDEX CONCURRENTLY) to avoid downtime | Reject if large table DDL uses locking operations |
| DB-SG-010 | Schema changes to audit tables are prohibited — audit schema is immutable after creation | Reject if audit table schema modification is proposed |

---

## 7. Backup & Recovery Testing Guardrails (ref: DA-GR-019)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-BG-001 | Database backup and recovery must be tested at minimum quarterly (ref: DA-GR-019) | Reject if backup recovery has not been tested in the last quarter |
| DB-BG-002 | Continuous WAL/transaction log archiving must be enabled for all production databases | Reject if WAL archiving is not enabled |
| DB-BG-003 | Daily full backups must be performed and verified (checksum validation) | Reject if daily backups are not configured or verified |
| DB-BG-004 | Backup retention: daily backups for 30 days; monthly backups for 7 years (regulatory) | Reject if retention is below these minimums |
| DB-BG-005 | Backups must be encrypted with AES-256 — backup encryption key stored separately from backup | Reject if backups are unencrypted |
| DB-BG-006 | Backups must be stored in a separate location from the primary database (different AZ minimum; different region for Tier 1) | Reject if backups are co-located with primary |
| DB-BG-007 | Point-in-time recovery (PITR) must be supported and tested — recover to any point within the retention window | Reject if PITR is not configured or tested |
| DB-BG-008 | Automated weekly restore test must verify: backup integrity, data completeness, application connectivity | Reject if no automated restore testing exists |
| DB-BG-009 | Recovery time must meet RTO targets: Tier 1 < 30 minutes; Tier 2 < 1 hour; Tier 3 < 4 hours | Reject if recovery time exceeds RTO in testing |
| DB-BG-010 | Backup monitoring must alert on: backup failure, backup size anomaly (> 20% change), backup age exceeding schedule | Reject if no backup monitoring/alerting exists |
| DB-BG-011 | Pre-migration/pre-deployment snapshots must be taken before every schema change or data migration | Reject if deployment proceeds without pre-migration snapshot |
| DB-BG-012 | Backup access must be restricted and audited — same access controls as production data | Reject if backup access is less restricted than production |

---

## 8. Query Security Guardrails (ref: DA-BP-018)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-QG-001 | All database queries must use parameterized statements — string concatenation for query building is prohibited | Reject if string concatenation is used in queries |
| DB-QG-002 | ORM queries must use named parameters or query methods — no raw string JPQL/HQL with concatenation | Reject if ORM queries use string concatenation |
| DB-QG-003 | Dynamic queries must use Criteria API, QueryDSL, or equivalent type-safe builder — never string assembly | Reject if dynamic queries use string assembly |
| DB-QG-004 | Native SQL queries require security review — verify parameterization and no injection vectors | Reject if native SQL has no security review |
| DB-QG-005 | Stored procedures must use typed parameters — no dynamic SQL execution within procedures | Reject if stored procedures use dynamic SQL |
| DB-QG-006 | Database error messages must not be exposed to API consumers — catch and translate to generic errors | Reject if SQL errors propagate to API responses |

---

## 9. Audit & Monitoring Guardrails (ref: DA-BP-017)

| ID | Guardrail | Enforcement |
|---|---|---|
| DB-AMG-001 | Audit triggers must be implemented on all tables containing Confidential or Restricted data | Reject if sensitive tables lack audit triggers |
| DB-AMG-002 | Audit triggers must capture: operation, old values, new values, changed columns, user, session, IP, correlation ID, timestamp | Reject if audit trigger captures incomplete data |
| DB-AMG-003 | Audit log tables must be immutable — no UPDATE or DELETE permissions for any role | Reject if audit tables allow modification |
| DB-AMG-004 | Audit log must be in a separate schema with independent access controls | Flag if audit log shares schema with application tables |
| DB-AMG-005 | Sensitive data in audit logs must be masked (account numbers, national IDs — last 4 digits only) | Reject if audit logs contain unmasked PII |
| DB-AMG-006 | Passwords, card data (PAN, CVV, PIN), and authentication tokens must never appear in audit logs | Reject immediately — no exceptions |
| DB-AMG-007 | Slow query monitoring must be enabled — alert on queries exceeding 1 second | Flag if slow query monitoring is not configured |
| DB-AMG-008 | Database connection monitoring must alert on: pool exhaustion, connection leaks, unusual connection patterns | Flag if connection monitoring is not configured |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot proceed | Unencrypted PII columns (DA-GR-013), credentials in config files (DA-GR-014), production data in non-prod (DA-GR-015), unmasked non-prod data (DA-GR-016), developer production DB access (DA-GR-017), unapproved schema change (DA-GR-018), untested backups (DA-GR-019) |
| **Flag** | Address before release | Key rotation > 90 days, manual masking, no JIT for DBA access, no session recording |
| **CI/CD Gate** | Automated | Secret scanning for credentials, migration script validation, rollback script verification |
| **Process** | Workflow/calendar | Quarterly backup testing, quarterly access review, quarterly production data scan |

---

## Quick Reference: Guardrail Triggers

| Database Activity | Triggered Guardrails |
|---|---|
| New table with PII | DB-EG-001→010 (encryption), DB-AMG-001→006 (audit triggers), DB-MG-001→008 (masking for non-prod) |
| New service needing DB access | DB-CG-001→010 (credentials via vault), DB-AG-002 (least privilege), DB-AG-008 (individual account) |
| Schema change | DB-SG-001→010 (change control, DBA review, rollback script, backward compatibility) |
| Environment provisioning | DB-PG-001→007 (no production data), DB-MG-001→008 (masking required), DB-CG-007 (separate credentials) |
| Production incident requiring DB access | DB-AG-004 (JIT), DB-AG-007 (emergency approval), DB-AG-005→006 (logging + recording) |
| Backup/DR activity | DB-BG-001→012 (quarterly testing, encryption, separate location, PITR, monitoring) |

---

## Pre-Deployment Database Checklist

| # | Check | Ref |
|---|---|---|
| 1 | PII/sensitive columns encrypted (AES-256-GCM) | DA-GR-013 |
| 2 | Credentials in secure vault; not in code/config | DA-GR-014 |
| 3 | No production data in non-production environments | DA-GR-015 |
| 4 | Non-production data masked/anonymized | DA-GR-016 |
| 5 | Direct DB access restricted to DBAs; app uses least-privilege account | DA-GR-017 |
| 6 | Schema change approved, reviewed by DBA, rollback script tested | DA-GR-018 |
| 7 | Backups configured, encrypted, tested within last quarter | DA-GR-019 |
| 8 | Audit triggers on all sensitive tables | DA-BP-017 |
| 9 | All queries parameterized; no string concatenation | DA-BP-018 |
| 10 | TDE enabled on database | DA-BP-015 |
| 11 | Row-level security configured for multi-tenant data | DA-BP-014 |
| 12 | Connection pooling configured for read/write separation | DA-BP-016 |
