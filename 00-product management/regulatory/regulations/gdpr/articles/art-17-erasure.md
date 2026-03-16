# GDPR Article 17 — Right to Erasure (Right to be Forgotten)

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 17
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Data subjects have the right to request erasure of their personal data without undue delay. The controller must erase the data and notify all downstream processors. Exceptions exist for legal obligations (e.g., AML retention).

## Technical Controls Required
1. Erasure request intake: customer portal, branch, email — all channels
2. Identity verification before processing erasure
3. Data inventory: identify all systems storing the data subject's personal data
4. Exemption assessment: check each data field against retention obligations
5. Erasure execution: anonymize PII fields (soft-delete with field-level anonymization)
6. Downstream notification: notify all third-party processors of erasure
7. Confirmation: send erasure confirmation to data subject
8. Audit trail: log all erasure actions (without re-creating erased data)

## Timeline Requirements
- Complete within 30 days of verified request
- Extension: additional 2 months for complex requests (must notify data subject within 30 days)

## Exemptions (Banking-Specific)
| Exemption | Legal Basis | Data Retained | Duration |
|---|---|---|---|
| AML/KYC retention | AML Directive Art. 40 | Transaction records, KYC documents | 5 years post-relationship |
| Tax reporting | Local tax law | Tax-relevant transaction data | Per local law (typically 7 years) |
| Legal claims | Art. 17(3)(e) | Data relevant to pending/potential claims | Duration of limitation period |
| Regulatory obligation | Art. 17(3)(b) | Data required by banking regulator | Per regulatory requirement |

## Conflicts with Other Regulations
- **AML Directive:** 5-year retention for transaction records → anonymize customer reference, retain transaction with anonymized ID
- **Tax law:** 7-year retention for tax-relevant data → retain with anonymized customer reference
- **Resolution pattern:** Anonymize the customer identifier; retain the transaction/record with a non-reversible reference; document the legal basis per retained field

## Applicability
- **Applies when:** Data subject requests erasure; processing no longer necessary; consent withdrawn; unlawful processing
- **Data types:** All personal data across all systems (databases, caches, backups, search indexes, third-party systems)

## Evidence Required
- Erasure request log (request date, verification date, completion date)
- Data inventory showing all systems checked
- Exemption documentation with legal basis per retained field
- Downstream processor notification confirmations
- Data subject confirmation of erasure
- Audit trail of erasure process

## Acceptance Criteria
Given a verified customer erasure request,
When the erasure workflow completes,
Then all PII is anonymized within 30 days
  AND downstream processors notified within 48 hours
  AND audit log captures: request date, verification date, erasure date, fields anonymized
  AND AML-exempt records retained with documented legal basis per field
  AND customer receives confirmation of erasure
  AND backup data is handled per backup erasure policy.

## Penalties
- Up to €20M or 4% of annual global turnover
