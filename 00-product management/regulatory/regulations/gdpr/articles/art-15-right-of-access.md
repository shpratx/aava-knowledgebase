# GDPR Article 15 — Right of Access (Subject Access Requests)

## Regulation Reference
- Regulation: GDPR (2016/679) — Article 15
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Data subjects have the right to obtain confirmation of whether their data is being processed, and if so, access to the data and specific information about the processing. This is commonly known as a Subject Access Request (SAR).

## Information to Provide
| Information | Content |
|---|---|
| Confirmation | Whether personal data is being processed |
| Copy of data | All personal data being processed |
| Purposes | Why the data is processed |
| Categories | Types of personal data |
| Recipients | Who has received the data |
| Retention | How long data will be stored |
| Rights | Rectification, erasure, restriction, objection rights |
| Source | Where data was obtained (if not from data subject) |
| Automated decisions | Logic, significance, consequences of automated processing |

## Timeline
- Respond within 1 month of receipt
- Extension: additional 2 months for complex/numerous requests (notify within 1 month)
- Free of charge (can charge reasonable fee for manifestly unfounded/excessive requests)

## Technical Controls Required
1. SAR intake channel: online form, email, branch — all channels accepted
2. Identity verification before processing SAR
3. Data discovery: automated search across all systems storing the data subject's data
4. Data compilation: aggregate data from all sources into structured format
5. Third-party data redaction: redact other individuals' data from the response
6. Secure delivery: encrypted delivery of SAR response
7. Exemption assessment: check if any exemptions apply (crime prevention, legal privilege)
8. SLA tracking: 1-month deadline with automated reminders
9. Audit trail: log SAR receipt, processing steps, and response

## Banking-Specific Considerations
| Consideration | Handling |
|---|---|
| Transaction history | Include all transactions (may be voluminous) |
| Internal notes | Include unless exemption applies (e.g., legal privilege) |
| Credit scoring data | Include score, factors, and logic explanation |
| Fraud flags | May be exempt under crime prevention (DPA 2018 Sch 2 Para 2) |
| Third-party data in transactions | Redact other customers' personal data |
| AML/SAR data | Exempt under crime prevention — do not disclose SAR existence |

## Evidence Required
- SAR intake and tracking records
- Identity verification records
- Data search and compilation records
- Exemption assessment records
- Response delivery records with timestamps
- SLA compliance metrics

## Acceptance Criteria
Given a valid Subject Access Request,
When the SAR is processed,
Then the data subject's identity is verified
  AND all systems are searched for the data subject's personal data
  AND data is compiled and third-party data redacted
  AND exemptions are assessed and documented
  AND response is delivered securely within 1 month
  AND the SAR process is audit-logged.

## Penalties
- Up to €20M or 4% of annual global turnover
