# PSR — APP Fraud Reimbursement

## Regulation Reference
- Regulator: PSR
- Rule: PS23/3 (APP Fraud Reimbursement)
- Effective: 7 October 2024

## Obligation (Plain Language)
Payment service providers must reimburse victims of Authorised Push Payment (APP) fraud, with costs split 50:50 between sending and receiving PSPs. Maximum reimbursement level set by PSR.

## Technical Controls Required
1. **Confirmation of Payee (CoP):** Implement name-checking before payment execution; warn customer on mismatch
2. **Fraud detection:** Real-time transaction monitoring for APP fraud patterns
3. **Warning messages:** Effective scam warnings at point of payment (not generic; context-specific)
4. **Reimbursement workflow:** Automated claim intake, assessment, and reimbursement within 5 business days
5. **Reporting:** Report APP fraud data to PSR per reporting requirements
6. **Customer vulnerability:** Enhanced protections for vulnerable customers
7. **Sending PSP obligations:** Warn customer; implement CoP; detect fraud patterns
8. **Receiving PSP obligations:** Detect mule accounts; act on fraud intelligence

## Evidence Required
- CoP implementation and match rate statistics
- Fraud detection rules and alert volumes
- Warning message effectiveness data
- Reimbursement claim records (volumes, outcomes, timelines)
- PSR reporting submissions
- Vulnerability identification records

## Acceptance Criteria
Given a customer initiating a payment,
When Confirmation of Payee returns a mismatch,
Then a clear warning is displayed explaining the mismatch
  AND the customer must actively acknowledge the warning to proceed
  AND the warning and customer response are logged for evidence
  AND if the payment is later confirmed as APP fraud, reimbursement is processed within 5 business days.
