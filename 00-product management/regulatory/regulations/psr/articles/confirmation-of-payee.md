# PSR — Confirmation of Payee (CoP)

## Regulation Reference
- Regulator: PSR
- Direction: General Direction 4 (CoP)
- Effective: Phase 1: June 2020; Phase 2: October 2024 (expanded scope)

## Obligation (Plain Language)
Directed PSPs must implement Confirmation of Payee — a name-checking service that verifies the name provided by the payer matches the name on the payee's account before a payment is made.

## CoP Response Types
| Response | Meaning | UI Action |
|---|---|---|
| Full match | Name matches account holder | Proceed with payment; green indicator |
| Partial match | Close but not exact match | Show actual name; ask customer to verify; amber warning |
| No match | Name does not match | Strong warning; show actual name; require explicit confirmation; red warning |
| Not available | Payee's PSP doesn't support CoP | Inform customer that name could not be checked; proceed with caution warning |
| Account not found | Account does not exist | Block payment; error message |

## Technical Controls Required
1. CoP API integration with Pay.UK CoP service (or equivalent)
2. Real-time name check before payment submission
3. UI handling for all 5 response types with appropriate warnings
4. Customer response logging (acknowledged warning, proceeded anyway)
5. Secondary Reference Data (SRD) support for business accounts
6. Bulk payment CoP checking capability

## Evidence Required
- CoP integration test results
- Match rate statistics (full/partial/no match/unavailable)
- Warning message UI screenshots
- Customer acknowledgment logs
- SRD implementation for business accounts

## Acceptance Criteria
Given a customer entering a payment with payee name "John Smith" and sort code/account number,
When CoP check returns "No Match" with actual name "Jane Doe",
Then a red warning is displayed: "The name you entered doesn't match the account holder. Account holder name: Jane Doe"
  AND the customer must explicitly confirm to proceed
  AND the warning and customer decision are logged
  AND the payment is flagged for enhanced fraud monitoring.
