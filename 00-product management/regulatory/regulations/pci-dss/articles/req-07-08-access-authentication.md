# PCI-DSS Requirements 7-8 — Access Control and Authentication

## Regulation Reference
- Standard: PCI-DSS v4.0 — Requirements 7 and 8

## Obligation
Req 7: Restrict access to cardholder data by business need to know. Req 8: Identify users and authenticate access to system components.

## Technical Controls
1. Role-based access control; least privilege; need-to-know basis
2. Unique user ID for every person with computer access — no shared/generic accounts
3. MFA for all access to CDE (administrative and remote)
4. MFA for all remote network access
5. Password policy: minimum 12 characters (or 8 + complexity); changed every 90 days; not reused for 4 cycles
6. Account lockout after maximum 10 failed attempts; 30-minute lockout or manual unlock
7. Session timeout: 15 minutes of inactivity
8. Access reviews: at least every 6 months; revoke terminated users immediately
9. Service accounts: limited privileges; managed passwords; no interactive login

## Acceptance Criteria
Given access to the CDE, Then every user has a unique ID AND MFA is required for all CDE and remote access AND passwords meet policy AND accounts lock after 10 failures AND sessions timeout at 15 minutes AND access is reviewed every 6 months.
