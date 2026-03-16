# GDPR Article 5 — Principles Relating to Processing of Personal Data

## Regulation Reference
- Regulation: GDPR (2016/679) — Article 5
- Enforcing Body: National DPAs

## Obligation (Plain Language)
All personal data processing must comply with seven fundamental principles. These are the foundation of GDPR — every other article builds on them.

## The Seven Principles
| Principle | Art. 5(1) | Obligation | Banking Implementation |
|---|---|---|---|
| Lawfulness, fairness, transparency | (a) | Process lawfully (Art. 6 basis); fairly (no unexpected use); transparently (privacy notices) | Legal basis per activity; privacy notices; no dark patterns |
| Purpose limitation | (b) | Collect for specified, explicit, legitimate purposes; no further processing incompatible with original purpose | Purpose documented per field; no repurposing without new basis |
| Data minimization | (c) | Adequate, relevant, limited to what is necessary | Justify every PII field; remove unnecessary collection |
| Accuracy | (d) | Accurate and kept up to date; inaccurate data erased/rectified without delay | Customer self-service profile updates; data quality checks |
| Storage limitation | (e) | Kept no longer than necessary for the purpose | Automated retention policies; defined periods per data type |
| Integrity and confidentiality | (f) | Appropriate security: protection against unauthorized/unlawful processing, accidental loss/destruction/damage | Encryption, access control, audit logging, DR/BCP |
| Accountability | 5(2) | Controller must demonstrate compliance with all principles | Processing records, DPIAs, audit trails, policies |

## Technical Controls Required
1. Processing register documenting legal basis and purpose per activity (Art. 30)
2. Data minimization review for every feature collecting PII
3. Automated retention policies with defined periods per data type
4. Encryption at rest (AES-256) and in transit (TLS 1.2+)
5. RBAC + resource-level access control
6. Audit logging of all access to personal data
7. Customer self-service for profile accuracy (view, update, correct)
8. Evidence collection for accountability (policies, DPIAs, training records)

## Evidence Required
- Processing register (Art. 30) with all seven principles addressed per activity
- Data minimization assessments per feature
- Retention policy with automated enforcement verification
- Encryption and access control configuration
- Audit log samples
- DPIA records for high-risk processing
- Staff training records

## Acceptance Criteria
Given any processing of personal data,
When assessed against Art. 5 principles,
Then a lawful basis is documented (lawfulness)
  AND purpose is specified and data is not repurposed (purpose limitation)
  AND every PII field is justified as necessary (minimization)
  AND customers can view and correct their data (accuracy)
  AND retention periods are defined and automated (storage limitation)
  AND data is encrypted and access-controlled (integrity/confidentiality)
  AND compliance is demonstrable with evidence (accountability).

## Penalties
- Up to €20M or 4% of annual global turnover
