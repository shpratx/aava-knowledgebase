# DORA Articles 17-19 — ICT-Related Incident Management and Reporting

## Regulation Reference
- Regulation: DORA (2022/2554) — Articles 17-19

## Obligation
Financial entities must classify ICT-related incidents, report major incidents to competent authorities, and voluntarily report significant cyber threats.

## Incident Classification Criteria
| Criterion | Threshold for "Major" |
|---|---|
| Clients affected | > 10% of clients or > 100,000 clients |
| Duration | > 2 hours for critical services |
| Geographic spread | > 2 EU member states |
| Data loss | Any personal data breach or confidentiality breach |
| Economic impact | Direct/indirect costs above materiality threshold |
| Critical services affected | Any critical or important function impacted |

## Reporting Timeline
| Report | Deadline | Content |
|---|---|---|
| Initial notification | Within 4 hours of classification as major (24 hours from detection) | Basic facts: what happened, when, impact |
| Intermediate report | Within 72 hours | Updated assessment, root cause (if known), mitigation actions |
| Final report | Within 1 month | Full RCA, lessons learned, remediation actions |

## Technical Controls
1. Incident detection: automated monitoring and alerting (Art. 10)
2. Incident classification engine: automated assessment against major incident criteria
3. Incident management workflow: detection → classification → containment → reporting → recovery → lessons learned
4. Regulatory reporting: pre-built templates; automated submission to competent authority
5. Communication procedures: internal escalation; customer notification; regulatory notification
6. Evidence preservation: forensic data captured before remediation
7. Lessons learned: documented and fed back into risk management framework

## Acceptance Criteria
Given a major ICT incident, Then the incident is classified within 4 hours of detection AND initial notification is submitted to competent authority within 4 hours of classification AND intermediate report within 72 hours AND final report within 1 month AND evidence is preserved AND lessons learned are documented.
