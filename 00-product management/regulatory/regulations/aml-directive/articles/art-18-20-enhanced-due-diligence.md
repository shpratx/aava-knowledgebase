# AML Directive Articles 18-20 — Enhanced Due Diligence (EDD)

## Regulation Reference
- Directive: 4AMLD (2015/849) / 5AMLD (2018/843) — Articles 18-20

## Obligation
Enhanced due diligence measures must be applied in high-risk situations: high-risk third countries, PEPs, complex/unusual transactions, and other situations identified by the firm's risk assessment.

## EDD Measures
| Trigger | Additional Measures |
|---|---|
| High-risk third country (FATF) | Enhanced identity verification; source of funds; source of wealth; senior management approval; enhanced ongoing monitoring |
| Politically Exposed Person (PEP) | Senior management approval; source of wealth; source of funds; enhanced ongoing monitoring; applies to PEP + family + close associates |
| Complex/unusual transactions | Enhanced scrutiny; documented business rationale; senior management awareness |
| Correspondent banking | Nature of respondent's business; AML controls assessment; senior management approval |

## Technical Controls
1. PEP screening: automated screening against PEP databases at onboarding + periodic (annual minimum)
2. High-risk country screening: automated check against FATF/EU high-risk country lists
3. Risk scoring: automated risk score incorporating PEP status, country risk, transaction patterns
4. Senior management approval workflow: system-enforced for high-risk customers
5. Enhanced monitoring: tighter transaction monitoring rules for EDD customers
6. Source of funds/wealth documentation: upload and verification workflow
7. Periodic review: annual for high-risk; triggered by risk events

## Acceptance Criteria
Given a customer identified as high-risk (PEP, high-risk country, unusual activity), When EDD is triggered, Then enhanced identity verification is performed AND source of funds/wealth is documented AND senior management approval is obtained AND enhanced ongoing monitoring is activated AND periodic review is scheduled (annual minimum).
