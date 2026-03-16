# DORA Articles 28-30 — ICT Third-Party Risk Management

## Regulation Reference
- Regulation: DORA (2022/2554) — Articles 28-30

## Obligation
Financial entities must manage ICT third-party risk throughout the lifecycle: pre-contractual assessment, contractual provisions, ongoing monitoring, and exit planning. Critical ICT third-party providers are subject to direct oversight by ESAs.

## Key Contractual Requirements (Art. 30)
| Requirement | Detail |
|---|---|
| Service description | Clear description of functions and services |
| Data processing locations | Where data is processed and stored |
| Data protection | Compliance with data protection requirements |
| Service availability | SLAs for availability, performance, support |
| Incident notification | Obligation to report incidents affecting the financial entity |
| Audit rights | Financial entity and regulator right to audit/inspect |
| Exit strategy | Transition plan; data return; adequate transition period |
| Sub-outsourcing | Prior approval required; same contractual requirements flow down |
| Cooperation with authorities | Provider must cooperate with competent authorities |

## Technical Controls
1. Third-party register: all ICT third-party arrangements with criticality classification
2. Pre-contractual due diligence: security assessment, financial viability, regulatory compliance
3. Contractual framework: all Art. 30 requirements in contracts
4. Ongoing monitoring: SLA tracking, security posture monitoring, incident tracking
5. Concentration risk assessment: dependency on single providers; substitutability analysis
6. Exit planning: documented exit strategy for every critical provider; tested
7. Sub-outsourcing control: approval process; visibility into sub-outsourcing chain

## Acceptance Criteria
Given a critical ICT third-party provider, Then the arrangement is in the third-party register AND pre-contractual due diligence is completed AND contract includes all Art. 30 requirements AND SLA monitoring is active AND exit strategy is documented and tested AND concentration risk is assessed AND sub-outsourcing is controlled.
