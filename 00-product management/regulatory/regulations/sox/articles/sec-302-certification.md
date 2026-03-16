# SOX Section 302 — Corporate Responsibility for Financial Reports

## Regulation Reference
- Act: Sarbanes-Oxley Act (2002) — Section 302

## Obligation
CEO and CFO must personally certify the accuracy of financial reports and the effectiveness of internal controls. IT systems supporting financial reporting must be reliable, accurate, and controlled.

## Technical Controls
1. Financial system access controls: segregation of duties; maker-checker
2. Data integrity: automated reconciliation; validation rules; checksums
3. Change management: all changes to financial systems approved, tested, documented
4. Audit trail: all modifications to financial data logged with before/after values
5. Report accuracy: automated validation of report outputs against source data
6. IT General Controls (ITGCs): access management, change management, operations, SDLC

## Acceptance Criteria
Given financial reporting systems, Then segregation of duties is enforced AND all changes follow change management AND financial data modifications are audit-logged AND report outputs are validated against source data AND ITGCs are documented and tested.
