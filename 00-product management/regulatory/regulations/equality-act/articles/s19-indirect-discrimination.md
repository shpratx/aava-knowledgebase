# Equality Act 2010 — Section 19: Indirect Discrimination

## Regulation Reference
- Act: Equality Act 2010 — Section 19
- Enforcing Body: EHRC / courts

## Obligation
A provision, criterion, or practice (PCP) that applies to everyone equally but puts people with a protected characteristic at a particular disadvantage is indirect discrimination — unless it can be objectively justified.

## Banking IT Relevance
| Risk Area | Example | Protected Characteristic |
|---|---|---|
| Automated credit scoring | Algorithm trained on biased historical data | Race, sex, age |
| Affordability assessment | Criteria that disadvantage part-time workers | Sex (disproportionately women) |
| Digital-only services | No branch/phone alternative | Age, disability |
| Identity verification | ID requirements that disadvantage certain nationalities | Race, nationality |
| Product eligibility | Age-based restrictions without justification | Age |
| Pricing algorithms | Differential pricing correlated with protected characteristics | Race, sex, age |

## Technical Controls
1. Algorithmic bias testing: test all automated decision models for disparate impact across protected characteristics
2. Fairness metrics: demographic parity, equalized odds, predictive parity — measured and reported
3. Alternative access: digital services must have non-digital alternatives (phone, branch, post)
4. Inclusive design: user research with diverse participants including protected groups
5. Regular equality impact assessment: for new products, features, and policy changes
6. Human review: automated decisions with significant impact must have human review option
7. Monitoring: ongoing monitoring of outcomes by protected characteristic (where data available)

## Acceptance Criteria
Given an automated decision system or policy, When assessed for indirect discrimination, Then bias testing is conducted across protected characteristics AND fairness metrics are measured and within acceptable thresholds AND alternative access channels exist for digital services AND human review is available for significant automated decisions AND equality impact assessment is completed for new products/features.
