# GDPR Articles 13-14 — Information to Data Subjects (Privacy Notices)

## Regulation Reference
- Regulation: GDPR (2016/679) — Articles 13 (direct collection) and 14 (indirect collection)
- Enforcing Body: National DPAs

## Obligation (Plain Language)
When collecting personal data, controllers must provide specific information to data subjects. Art. 13 applies when data is collected directly from the individual; Art. 14 when obtained from other sources.

## Required Information (Art. 13)
| Information | Example |
|---|---|
| Controller identity and contact | "Bank Name, 123 Street, dpo@bank.com" |
| DPO contact | "dpo@bank.com" |
| Purposes and legal basis | "Account management — contractual necessity (Art. 6(1)(b))" |
| Legitimate interests (if Art. 6(1)(f)) | "Fraud prevention" |
| Recipients/categories | "Payment processors, fraud detection services, regulators" |
| Cross-border transfers | "Data transferred to US under SCCs" |
| Retention period | "Duration of account + 7 years" |
| Data subject rights | Access, rectification, erasure, restriction, portability, objection |
| Right to withdraw consent | If consent-based processing |
| Right to complain to DPA | "You can complain to the ICO at ico.org.uk" |
| Whether provision is statutory/contractual | "KYC data is required by law; marketing consent is optional" |
| Automated decision-making | If applicable: logic, significance, consequences |

## Technical Controls Required
1. Layered privacy notice: short notice at point of collection + link to full notice
2. Privacy notice accessible from every page (footer link)
3. Purpose-specific notice per data collection form
4. Just-in-time notices for sensitive data fields ("Why do we need this?")
5. Privacy notice versioning with change tracking
6. Notification of material changes to privacy notice
7. Privacy notice in plain language; WCAG 2.1 AA accessible
8. Consent mechanism where consent is the legal basis (Art. 7 compliant)

## Evidence Required
- Privacy notice (current version) with all required information
- Version history of privacy notices
- UI screenshots showing notice placement and accessibility
- Change notification records for material updates
- Layered notice implementation (short + full)

## Acceptance Criteria
Given a data collection point (form, onboarding, feature),
When personal data is collected from the customer,
Then a privacy notice containing all Art. 13 required information is presented before or at collection
  AND the notice is in plain language and accessible (WCAG 2.1 AA)
  AND a layered approach is used (short notice + link to full)
  AND the notice is versioned and changes are notified to customers.

## Penalties
- Up to €20M or 4% of annual global turnover
