# PRA SS1/23 — Model Risk Management

## Regulation Reference
- Regulator: PRA — SS1/23
- Effective: 17 May 2024

## Obligation
Firms must have a model risk management (MRM) framework covering all models, including AI/ML models. Models must be validated, monitored, and governed throughout their lifecycle.

## Scope — What is a "Model"
Any quantitative method that processes input data into quantitative estimates used for decision-making: credit scoring, fraud detection, pricing, risk calculation, AML transaction monitoring, affordability assessment.

## Technical Controls
1. Model inventory: all models registered with owner, purpose, methodology, data inputs, materiality
2. Model development standards: documented methodology, assumptions, limitations
3. Independent validation: models validated by team independent from developers
4. Ongoing monitoring: model performance tracked; drift detection; outcome analysis
5. Model change management: changes follow governance process with re-validation
6. AI/ML specific: explainability requirements; bias testing; fairness assessment
7. Data quality: input data quality monitored; data lineage documented
8. Model risk tiering: Tier 1 (high impact) through Tier 3 (low impact) with proportionate governance

## AI/ML Specific Requirements
| Requirement | Implementation |
|---|---|
| Explainability | Model decisions must be explainable to customers and regulators |
| Bias testing | Test for discrimination across protected characteristics (Equality Act) |
| Fairness | Outcomes must not disproportionately disadvantage protected groups |
| Human oversight | Human-in-the-loop for high-impact decisions (credit, fraud) |
| Monitoring | Continuous monitoring for model drift, performance degradation |
| Audit trail | All model decisions logged with inputs, outputs, and version |

## Acceptance Criteria
Given a model used for decision-making, Then the model is in the model inventory AND independently validated before production AND performance is monitored continuously AND bias/fairness testing is conducted for models affecting customers AND model decisions are logged with inputs and outputs AND changes follow governance with re-validation.
