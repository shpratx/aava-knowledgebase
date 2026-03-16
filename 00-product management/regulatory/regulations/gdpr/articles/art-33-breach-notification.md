# GDPR Article 33 — Notification of Breach to Supervisory Authority

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 33
- Enforcing Body: National DPAs

## Obligation (Plain Language)
In case of a personal data breach, the controller must notify the supervisory authority within 72 hours of becoming aware, unless the breach is unlikely to result in a risk to data subjects' rights and freedoms.

## Timeline
| Event | Deadline |
|---|---|
| Breach detected | T=0 (awareness) |
| Initial assessment | T+4 hours |
| Notification to DPA | T+72 hours (if risk to data subjects) |
| Notification to data subjects | Without undue delay (if high risk) — Art. 34 |
| Full incident report | Within 30 days |

## Notification Content (Art. 33(3))
| Field | Content |
|---|---|
| Nature of breach | Description of what happened |
| Data categories | Types of personal data affected |
| Data subject categories | Who is affected (customers, employees) |
| Approximate number affected | Number of data subjects and records |
| DPO contact | Name and contact details of DPO |
| Likely consequences | Assessment of impact on data subjects |
| Measures taken | Actions to address and mitigate the breach |

## Technical Controls Required
1. Breach detection: real-time monitoring for unauthorized access, data exfiltration, encryption failures
2. Breach classification: automated triage to determine if notification is required
3. Notification workflow: pre-built templates, approval chain, DPA submission mechanism
4. Evidence preservation: forensic data captured before remediation
5. Communication templates: pre-approved templates for DPA and data subject notification
6. Breach register: all breaches recorded regardless of notification requirement
7. Annual tabletop exercise: test the entire notification process

## Evidence Required
- Breach detection monitoring configuration
- Breach register (all incidents)
- Notification records (DPA submissions with timestamps)
- Tabletop exercise results (annual)
- Communication templates (pre-approved)

## Acceptance Criteria
Given a confirmed personal data breach,
When the breach is classified as reportable,
Then the supervisory authority is notified within 72 hours
  AND notification includes all required fields (nature, categories, numbers, consequences, measures)
  AND breach is recorded in the breach register
  AND forensic evidence is preserved
  AND data subjects are notified without undue delay if high risk (Art. 34).

## Penalties
- Up to €10M or 2% of annual global turnover
