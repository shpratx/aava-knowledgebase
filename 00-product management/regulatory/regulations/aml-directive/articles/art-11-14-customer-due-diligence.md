# AML Directive Articles 11-14 — Customer Due Diligence (CDD)

## Regulation Reference
- Directive: 4AMLD (2015/849) / 5AMLD (2018/843)
- Articles: 11-14
- Enforcing Body: National FIUs / banking regulators

## Obligation (Plain Language)
Banks must identify and verify the identity of customers before establishing a business relationship or carrying out occasional transactions above thresholds. Ongoing monitoring of the business relationship is required.

## CDD Measures Required
| Measure | Description | Implementation |
|---|---|---|
| Customer identification | Collect name, DOB, address, national ID | Digital onboarding form; document upload |
| Identity verification | Verify against reliable, independent source | ID document verification (OCR + liveness); database checks |
| Beneficial ownership | Identify beneficial owners (>25% ownership) | Corporate registry checks; UBO declaration |
| Purpose of relationship | Understand nature and purpose of business relationship | Account purpose questionnaire |
| Ongoing monitoring | Monitor transactions for consistency with customer profile | Automated transaction monitoring; periodic review |

## Enhanced Due Diligence (EDD) Triggers
| Trigger | Additional Measures |
|---|---|
| High-risk country (FATF list) | Enhanced verification; senior management approval; source of funds |
| Politically Exposed Person (PEP) | Senior management approval; source of wealth; enhanced monitoring |
| Complex/unusual transactions | Enhanced scrutiny; documented rationale |
| High-value transactions | Source of funds verification |

## Technical Controls Required
1. KYC data collection workflow with document upload and verification
2. Sanctions screening against OFAC, EU, UN sanctions lists (real-time)
3. PEP screening against PEP databases (at onboarding + periodic)
4. Transaction monitoring engine with configurable rules
5. Risk scoring per customer (low/medium/high)
6. Periodic review triggers (annual for high-risk, 3-year for medium, 5-year for low)
7. Record retention: 5 years post-relationship for all CDD data

## Evidence Required
- KYC records for all customers (ID documents, verification results)
- Sanctions screening results (at onboarding + ongoing)
- PEP screening results
- Risk assessment per customer
- Transaction monitoring rules and alert records
- Periodic review records
- SAR filings (if applicable)

## Acceptance Criteria
Given a new customer onboarding,
When the KYC process completes,
Then customer identity is verified against reliable source
  AND sanctions screening is performed (OFAC, EU, UN lists)
  AND PEP screening is performed
  AND risk score is assigned (low/medium/high)
  AND all CDD data is retained for 5 years post-relationship
  AND ongoing transaction monitoring is activated.
