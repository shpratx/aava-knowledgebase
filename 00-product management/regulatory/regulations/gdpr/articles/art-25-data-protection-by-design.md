# GDPR Article 25 — Data Protection by Design and by Default

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 25
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Controllers must implement appropriate technical and organizational measures to implement data protection principles (like data minimization) effectively, both at the time of design and during processing. By default, only personal data necessary for each specific purpose should be processed.

## Technical Controls Required
1. Privacy Impact Assessment (PIA/DPIA) at design phase for new features processing PII
2. Data minimization review: justify every PII field collected against stated purpose
3. Purpose limitation: data collected for one purpose not used for another without new legal basis
4. Storage limitation: automated retention policies with defined periods per data type
5. Encryption by default: AES-256 at rest, TLS 1.2+ in transit for all personal data
6. Access control by default: least privilege; role-based + resource-level authorization
7. Pseudonymization where possible: use tokens/references instead of direct identifiers
8. Audit logging by default: log all access to personal data

## SDLC Integration
| Phase | Data Protection Activity |
|---|---|
| Requirements | Identify PII; determine legal basis; classify data; specify retention |
| Design | DPIA if high-risk; design encryption, access control, audit; plan erasure capability |
| Development | Implement controls; parameterized queries; input validation; output masking |
| Testing | Verify encryption, access control, audit logging, erasure, consent flows |
| Deployment | Compliance sign-off; monitoring active; evidence collection |
| Operations | Ongoing monitoring; data subject request handling; periodic review |

## Evidence Required
- DPIA documents for high-risk processing
- Data minimization assessments per feature
- Technical architecture showing encryption, access control, audit
- Privacy notice showing purpose limitation
- Retention policy with automated enforcement

## Acceptance Criteria
Given a new feature processing personal data,
When the feature is designed,
Then a data minimization assessment is completed (every PII field justified)
  AND a DPIA is completed if high-risk processing
  AND encryption at rest and in transit is designed
  AND access control follows least privilege
  AND retention policy is defined with automated enforcement
  AND erasure capability is designed from the start.

## Penalties
- Up to €10M or 2% of annual global turnover
