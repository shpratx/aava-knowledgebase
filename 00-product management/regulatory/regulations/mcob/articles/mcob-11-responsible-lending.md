# MCOB 11 — Responsible Lending and Affordability

## Regulation Reference
- Sourcebook: MCOB (FCA Handbook)
- Chapter: 11
- Enforcing Body: FCA

## Obligation (Plain Language)
Lenders must not enter into a regulated mortgage contract unless they can demonstrate the customer can afford the repayments. Affordability must be assessed based on income, expenditure, and stress-tested against interest rate increases.

## Technical Controls Required
1. **Income verification:** Automated income verification (payslips, bank statements, credit reference data)
2. **Expenditure assessment:** ONS data or customer-declared expenditure; committed expenditure from credit file
3. **Affordability calculation:** Net disposable income after all commitments; must be positive after mortgage payment
4. **Stress testing:** Affordability tested at stressed interest rate (typically SVR + buffer, or minimum stress rate per FCA)
5. **Hard/soft limits:** System-enforced maximum LTV, LTI, and DTI ratios
6. **Audit trail:** Full record of affordability assessment inputs, calculations, and decision
7. **Decline reasons:** Clear, specific reasons provided to customer on decline

## Affordability Assessment Model
| Input | Source | Validation |
|---|---|---|
| Gross income | Payslips, P60, tax returns, bank statements | Verified against multiple sources |
| Other income | Benefits, rental, investments | Documented and sustainable |
| Committed expenditure | Credit file (loans, cards, other mortgages) | Real-time credit check |
| Essential expenditure | ONS data or customer-declared | Minimum floor per household type |
| Mortgage payment | Calculated at product rate AND stressed rate | Both must be affordable |
| Stress rate | SVR + buffer (or FCA minimum) | Applied to full term |

## Evidence Required
- Affordability calculation records (inputs, formula, result)
- Income verification documents
- Credit check results
- Stress test results at stressed rate
- Decision record (approve/decline with rationale)
- Customer communication of decision

## Acceptance Criteria
Given a mortgage application,
When affordability is assessed,
Then income is verified against reliable sources
  AND all committed expenditure is captured from credit file
  AND affordability is positive at both product rate and stressed rate
  AND LTV, LTI, DTI ratios are within policy limits
  AND full assessment is audit-logged with all inputs and calculations
  AND decline reasons are specific and communicated to customer.
