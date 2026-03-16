# SOX Section 404 — Management Assessment of Internal Controls

## Regulation Reference
- Act: Sarbanes-Oxley Act (2002)
- Section: 404
- Enforcing Body: SEC / PCAOB

## Obligation (Plain Language)
Management must assess and report on the effectiveness of internal controls over financial reporting (ICFR). External auditors must attest to management's assessment.

## Technical Controls Required (IT General Controls)
1. **Access control:** Segregation of duties (maker-checker); least privilege; access reviews quarterly
2. **Change management:** All changes to financial systems require approval, testing, and documentation
3. **Audit trail:** All modifications to financial data logged with before/after values, user, timestamp
4. **Data integrity:** Reconciliation processes; validation rules; dual-control for financial calculations
5. **Backup and recovery:** Financial data backed up; recovery tested; RPO/RTO defined
6. **Logical security:** Authentication, authorization, encryption for financial systems
7. **Monitoring:** Continuous monitoring of control effectiveness; exception reporting

## Segregation of Duties Matrix
| Function | Initiator | Approver | Cannot Be Same Person |
|---|---|---|---|
| Payment initiation | Operations | Manager | Yes |
| Journal entry | Accountant | Controller | Yes |
| User access provisioning | Requester | Security admin | Yes |
| Code deployment | Developer | Release manager | Yes |
| Vendor setup | Procurement | Finance | Yes |

## Evidence Required
- Access control configuration and quarterly review records
- Change management records for all financial system changes
- Audit trail samples showing complete capture
- Reconciliation reports
- Backup and recovery test results
- Segregation of duties matrix and enforcement evidence
- Management's annual assessment report

## Acceptance Criteria
Given a financial system,
When assessed for SOX 404 compliance,
Then segregation of duties is enforced (maker-checker for all financial operations)
  AND all changes follow change management with approval and documentation
  AND all financial data modifications are audit-logged with before/after values
  AND access is reviewed quarterly with excessive privileges removed
  AND backup and recovery is tested with RPO/RTO targets met.
