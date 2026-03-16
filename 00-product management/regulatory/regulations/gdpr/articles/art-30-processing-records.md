# GDPR Article 30 — Records of Processing Activities

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 30
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Controllers and processors must maintain a register of all processing activities. This register must be available to supervisory authorities on request.

## Required Fields (Controller Register)
| Field | Description | Example |
|---|---|---|
| Processing activity name | Descriptive name | "Customer Account Management" |
| Controller identity | Organization name and contact | "Bank Name, DPO: dpo@bank.com" |
| Purpose of processing | Why data is processed | "Manage customer accounts per banking contract" |
| Legal basis | Art. 6 ground | "Contractual necessity — Art. 6(1)(b)" |
| Data subject categories | Who the data belongs to | "Retail banking customers" |
| Personal data categories | Types of data processed | "Name, address, DOB, email, phone, account data, transaction history" |
| Recipients | Who receives the data | "Core banking system, fraud detection, payment processors" |
| Cross-border transfers | Countries and mechanisms | "EU only" or "US — Standard Contractual Clauses" |
| Retention period | How long data is kept | "Duration of account + 7 years (regulatory)" |
| Security measures | Technical and organizational | "AES-256 encryption, RBAC, audit logging, TDE" |

## Technical Controls Required
1. Maintain a centralized, version-controlled processing register
2. Update register when new processing activities are introduced or changed
3. Link register entries to technical systems (databases, APIs, services)
4. Make register available to DPA within 24 hours of request
5. Review register quarterly and after significant system changes

## Evidence Required
- Complete processing register with all required fields
- Version history showing updates
- Quarterly review records
- Link to technical implementation per activity

## Acceptance Criteria
Given a new data processing activity is introduced,
When the activity goes live,
Then the processing register is updated with all required fields
  AND the register links to the technical systems involved
  AND the register is version-controlled with change history
  AND the register is available to supervisory authority within 24 hours.

## Penalties
- Up to €10M or 2% of annual global turnover
