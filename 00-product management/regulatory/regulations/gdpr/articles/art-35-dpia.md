# GDPR Article 35 — Data Protection Impact Assessment (DPIA)

## Regulation Reference
- Regulation: GDPR (2016/679) — Article 35
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Where processing is likely to result in a high risk to individuals' rights and freedoms, the controller must carry out a DPIA before the processing begins. This assesses the necessity, proportionality, and risks of the processing, and identifies measures to mitigate those risks.

## When DPIA is Required
| Trigger | Banking Example |
|---|---|
| Systematic and extensive profiling with significant effects | Automated credit scoring/decisioning |
| Large-scale processing of special category data | Health data for insurance products |
| Systematic monitoring of public areas | CCTV in branches (not primarily IT) |
| New technology | AI/ML for fraud detection; biometric authentication |
| Large-scale processing of personal data | Customer database; transaction monitoring |
| Matching or combining datasets | Combining credit bureau data with internal data |
| Data concerning vulnerable individuals | Products for elderly, children, financially vulnerable |
| Cross-border transfer to non-adequate country | Offshoring data processing |

## DPIA Content (Art. 35(7))
| Section | Content |
|---|---|
| Description of processing | What data, what purpose, what technology |
| Necessity and proportionality | Why this processing is needed; why less intrusive alternatives won't work |
| Risk assessment | Risks to data subjects' rights and freedoms |
| Mitigation measures | Technical and organizational measures to address risks |
| Consultation | DPO consulted; supervisory authority consulted if high residual risk (Art. 36) |

## Technical Controls Required
1. DPIA trigger assessment for every new feature/change involving personal data
2. DPIA template with all required sections
3. Risk scoring methodology (likelihood × impact)
4. Mitigation measures linked to technical controls (encryption, access control, minimization)
5. DPO review and sign-off
6. DPIA register tracking all assessments
7. DPIA review when processing changes materially
8. Art. 36 prior consultation with DPA if residual risk remains high after mitigation

## Evidence Required
- DPIA trigger assessment records (why DPIA was/wasn't required)
- Completed DPIA documents with all required sections
- DPO consultation records
- Mitigation measures implementation evidence
- DPIA register
- Review records when processing changed

## Acceptance Criteria
Given a new feature involving high-risk personal data processing,
When the feature is designed,
Then a DPIA trigger assessment is completed
  AND if triggered, a full DPIA is conducted before processing begins
  AND the DPIA assesses necessity, proportionality, and risks
  AND mitigation measures are identified and implemented
  AND the DPO is consulted and signs off
  AND the DPIA is recorded in the DPIA register
  AND if residual risk is high, the supervisory authority is consulted (Art. 36).

## Penalties
- Up to €10M or 2% of annual global turnover
