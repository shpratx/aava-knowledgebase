# FCA SYSC — Operational Resilience

## Regulation Reference
- Regulator: FCA (jointly with PRA)
- Rule: SYSC 15A (Operational Resilience)
- Policy Statement: PS21/3
- Effective: 31 March 2022; full compliance by 31 March 2025

## Obligation (Plain Language)
Firms must identify their important business services, set impact tolerances for disruption, and ensure they can remain within those tolerances through severe but plausible scenarios.

## Key Requirements

### Important Business Services (IBS)
| Requirement | Implementation |
|---|---|
| Identify IBS | Map services that, if disrupted, would cause intolerable harm to consumers, market integrity, or firm safety |
| Set impact tolerances | Maximum tolerable disruption for each IBS (time-based) |
| Map resources | Identify people, processes, technology, facilities, and third parties supporting each IBS |
| Scenario testing | Test ability to remain within impact tolerances under severe but plausible scenarios |

### Impact Tolerance Examples for Banking
| Important Business Service | Impact Tolerance |
|---|---|
| Payment processing | 4 hours maximum disruption |
| Customer account access (online/mobile) | 4 hours maximum disruption |
| Mortgage application processing | 24 hours maximum disruption |
| Complaint handling | 48 hours maximum disruption |
| Regulatory reporting | Per regulatory deadline |

## Technical Controls Required
1. Service dependency mapping (technology, third parties, data)
2. Impact tolerance monitoring (real-time availability tracking per IBS)
3. Scenario testing framework (annual minimum; after significant changes)
4. Incident management aligned to IBS impact tolerances
5. Third-party resilience assessment (critical third parties mapped to IBS)
6. Communication procedures for disruption events
7. Lessons learned process after incidents and tests

## Evidence Required
- IBS register with impact tolerances
- Resource mapping per IBS (technology, people, third parties)
- Scenario test plans and results
- Incident records showing response within impact tolerances
- Third-party resilience assessments
- Board-level reporting on operational resilience

## Acceptance Criteria
Given an important business service,
When a severe but plausible disruption scenario is tested,
Then the service recovers within the defined impact tolerance
  AND all supporting resources (technology, third parties) are mapped
  AND communication procedures are executed
  AND lessons learned are documented and acted upon.
