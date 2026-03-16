# AML Directive Article 40 — Record Retention

## Regulation Reference
- Directive: 4AMLD (2015/849) — Article 40

## Obligation
Firms must retain CDD data and transaction records for 5 years after the end of the business relationship or the date of the occasional transaction.

## Retention Requirements
| Data Type | Retention Period | Start Date |
|---|---|---|
| CDD documents (ID, verification) | 5 years | End of business relationship |
| Transaction records | 5 years | Date of transaction |
| SAR records | 5 years | Date of SAR filing |
| Risk assessments | 5 years | Date of assessment |
| Correspondence with FIU | 5 years | Date of correspondence |

## Technical Controls
1. Automated retention policy: 5-year retention from relationship end or transaction date
2. Retention hold: prevent deletion of records under active investigation
3. Secure storage: encrypted; access-controlled; tamper-proof
4. Retrieval capability: records retrievable within reasonable time for regulatory/law enforcement requests
5. Destruction: secure deletion after retention period (unless other regulation requires longer)
6. Conflict with GDPR: AML retention prevails; anonymize customer reference in retained records; document legal basis

## Acceptance Criteria
Given customer and transaction records, Then CDD data is retained for 5 years post-relationship AND transaction records retained for 5 years from transaction date AND retention is automated AND records are retrievable for regulatory requests AND GDPR conflict is resolved by anonymizing customer reference while retaining transaction data.
