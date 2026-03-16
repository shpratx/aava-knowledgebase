# DORA Articles 23-25 — Digital Operational Resilience Testing

## Regulation Reference
- Regulation: DORA (2022/2554) — Articles 23-25

## Obligation
Financial entities must establish a testing program for ICT tools, systems, and processes. Significant entities must conduct Threat-Led Penetration Testing (TLPT) at least every 3 years.

## Testing Requirements
| Test Type | Frequency | Scope | Who |
|---|---|---|---|
| Vulnerability scanning | Continuous/quarterly | All ICT systems | Internal or external |
| Open-source analysis (SCA) | Every build | All software dependencies | Automated in CI/CD |
| Network security assessment | Annual | Network architecture, segmentation | Internal or external |
| Penetration testing | Annual | Critical systems and applications | Qualified external testers |
| Source code review | Per release | Custom-developed software | Internal + external |
| Scenario-based testing | Annual | Business continuity, DR, incident response | Internal with external validation |
| TLPT (Threat-Led Pen Test) | Every 3 years | Critical functions; red team exercise | Qualified external (TIBER-EU framework) |

## Technical Controls
1. Testing program: documented, approved by management, covering all test types
2. Vulnerability management: scan → prioritize → remediate → verify cycle
3. Pen test scope: includes critical functions, APIs, authentication, authorization
4. TLPT: follows TIBER-EU or equivalent framework; covers people, processes, technology
5. Remediation tracking: all findings tracked to closure with SLA per severity
6. Test results reporting: to management body; to competent authority (TLPT results)
7. Third-party testing: critical ICT third-party providers included in testing scope

## Acceptance Criteria
Given critical ICT systems, Then vulnerability scanning runs continuously AND penetration testing is conducted annually AND TLPT is conducted every 3 years for significant entities AND all findings are remediated per SLA AND test results are reported to management AND critical third-party providers are included in testing scope.
