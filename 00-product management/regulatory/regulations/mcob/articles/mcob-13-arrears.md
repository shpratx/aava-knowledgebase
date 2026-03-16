# MCOB 13 — Arrears, Payment Shortfalls, and Repossessions

## Regulation Reference
- Sourcebook: MCOB (FCA Handbook)
- Chapter: 13
- Enforcing Body: FCA

## Obligation (Plain Language)
Firms must treat customers in arrears fairly. Repossession must be a last resort. Firms must consider reasonable forbearance options before taking enforcement action.

## Technical Controls Required
1. **Arrears detection:** Automated detection of missed/partial payments
2. **Customer contact:** Automated communication within 15 business days of arrears
3. **Forbearance options:** System supports: payment holiday, term extension, interest-only period, reduced payments, capitalization of arrears
4. **Forbearance assessment:** Affordability re-assessment for forbearance arrangements
5. **Arrears charges:** Transparent; reasonable; disclosed in advance
6. **Repossession workflow:** Multi-stage process with mandatory forbearance consideration before enforcement
7. **Vulnerability identification:** Flag vulnerable customers for enhanced support
8. **Regulatory reporting:** Arrears and forbearance data for FCA Product Sales Data (PSD)

## Evidence Required
- Arrears detection and communication records
- Forbearance options offered and customer response
- Affordability re-assessment for forbearance
- Repossession process records showing forbearance was considered
- Vulnerability identification records
- Regulatory reporting submissions

## Acceptance Criteria
Given a mortgage customer in arrears,
When the arrears management process is triggered,
Then the customer is contacted within 15 business days
  AND at least one forbearance option is offered and assessed for affordability
  AND arrears charges are transparent and reasonable
  AND repossession is only pursued after forbearance options are exhausted
  AND vulnerable customers receive enhanced support
  AND all interactions are audit-logged.
