# GDPR Article 32 — Security of Processing

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 32
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Controllers and processors must implement appropriate technical and organizational security measures, considering the state of the art, cost, nature/scope/context/purposes of processing, and risks to data subjects.

## Required Measures (Art. 32(1))
| Measure | Implementation | Banking Standard |
|---|---|---|
| (a) Pseudonymization and encryption | Tokenization for card data; AES-256 for PII at rest; TLS 1.2+ in transit | Mandatory for Confidential/Restricted data |
| (b) Confidentiality, integrity, availability | RBAC, input validation, audit logging, HA architecture | Mandatory for all systems processing personal data |
| (c) Restore availability after incident | DR/BCP: RTO < 1 hour, RPO < 15 minutes for Tier 1 | Mandatory; tested quarterly |
| (d) Regular testing and evaluation | Pen testing, SAST/DAST, DR drills, access reviews | Annual pen test; continuous SAST/DAST; quarterly DR test |

## Technical Controls Required
1. Encryption: AES-256 at rest; TLS 1.2+ in transit; HSM for key management
2. Access control: RBAC + resource-level; least privilege; MFA for sensitive operations
3. Audit logging: all access to personal data logged (who, what, when, where, outcome)
4. Vulnerability management: SAST/DAST on every build; pen test annually; patch critical CVEs within 24h
5. Incident response: breach detection, 72-hour notification capability, tested annually
6. DR/BCP: tested quarterly; RTO/RPO per data criticality tier

## Evidence Required
- Encryption implementation verification (storage inspection, network analysis)
- Access control configuration and review records
- Audit log samples showing completeness
- Vulnerability scan reports (SAST, DAST, pen test)
- DR test results
- Incident response test results

## Acceptance Criteria
Given a system processing personal data,
When assessed for Art. 32 compliance,
Then all personal data is encrypted at rest (AES-256) and in transit (TLS 1.2+)
  AND access control follows least privilege with MFA for sensitive operations
  AND all access to personal data is audit-logged
  AND vulnerability scanning runs on every build (SAST/DAST)
  AND DR is tested quarterly with RTO/RPO targets met
  AND incident response is tested annually with 72-hour notification capability.

## Penalties
- Up to €10M or 2% of annual global turnover
