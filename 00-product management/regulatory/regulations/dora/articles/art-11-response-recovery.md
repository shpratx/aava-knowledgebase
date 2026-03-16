# DORA Article 11 — Response and Recovery

## Regulation Reference
- Regulation: DORA (2022/2554)
- Article: 11
- Enforcing Body: ESAs / national authorities

## Obligation (Plain Language)
Financial entities must establish ICT business continuity policies and disaster recovery plans. These must be tested regularly and ensure continuity of critical functions.

## Technical Controls Required
1. ICT business continuity policy covering all critical functions
2. Disaster recovery plans for all critical ICT systems
3. RTO and RPO defined per system criticality tier
4. Automated failover for Tier 1 critical systems
5. Regular testing: at least annually; after significant changes
6. Communication procedures during ICT disruptions
7. Lessons learned process after incidents and tests
8. Third-party continuity: ensure ICT providers have adequate continuity arrangements

## Testing Requirements
| Test Type | Frequency | Scope |
|---|---|---|
| Tabletop exercise | Semi-annual | Scenario walkthrough with all stakeholders |
| Failover test | Annual (minimum) | Actual failover to DR site |
| Full DR drill | Annual | Complete recovery of critical systems |
| Third-party continuity test | Annual | Verify provider's continuity capabilities |

## Evidence Required
- ICT business continuity policy (approved by management body)
- DR plans per critical system
- RTO/RPO targets and test results showing targets met
- Test reports (tabletop, failover, full DR)
- Communication procedure documentation
- Lessons learned reports
- Third-party continuity assessment records

## Acceptance Criteria
Given a critical ICT system,
When a DR test is conducted,
Then the system recovers within the defined RTO
  AND data loss is within the defined RPO
  AND all critical functions are available after recovery
  AND communication procedures are executed successfully
  AND lessons learned are documented and acted upon.
