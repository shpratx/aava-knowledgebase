# DPA 2018 / UK GDPR — Breach Notification (UK-Specific)

## Regulation Reference
- Act: UK GDPR Art. 33-34 (as retained) + DPA 2018
- Enforcing Body: ICO

## Obligation (Plain Language)
Same as EU GDPR Art. 33-34 but notification is to the ICO (not EU DPAs). Must notify within 72 hours of becoming aware of a breach that poses a risk to individuals' rights and freedoms.

## UK-Specific Requirements
| Requirement | Detail |
|---|---|
| Notify to | ICO (Information Commissioner's Office) |
| Method | ICO online breach reporting tool (ico.org.uk) |
| Timeline | 72 hours from awareness |
| Content | Same as EU GDPR Art. 33(3) |
| Data subject notification | Without undue delay if high risk (UK GDPR Art. 34) |
| Record keeping | All breaches recorded regardless of notification (breach register) |

## Technical Controls Required
1. Same as GDPR Art. 33 controls (see GDPR knowledge base)
2. ICO-specific notification template pre-built
3. ICO online reporting tool integration/familiarity
4. UK DPO contact details in notification
5. Breach register maintained per UK GDPR Art. 33(5)

## Evidence Required
- ICO notification records with timestamps
- Breach register (all incidents)
- Data subject notification records (if high risk)
- ICO acknowledgment of notification

## Acceptance Criteria
Given a personal data breach affecting UK data subjects,
When the breach is classified as reportable,
Then the ICO is notified within 72 hours via the ICO online reporting tool
  AND notification includes all required fields
  AND the breach is recorded in the breach register
  AND data subjects are notified without undue delay if high risk.
