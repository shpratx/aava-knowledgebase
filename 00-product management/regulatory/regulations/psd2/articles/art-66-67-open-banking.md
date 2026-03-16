# PSD2 Articles 66-67 — Access to Accounts (Open Banking)

## Regulation Reference
- Directive: PSD2 (2015/2366) — Articles 66-67

## Obligation
Account Servicing Payment Service Providers (ASPSPs — banks) must allow access to customer accounts by authorized third-party providers: Account Information Service Providers (AISPs) for read access and Payment Initiation Service Providers (PISPs) for payment initiation.

## Technical Controls
1. Dedicated API interface (Open Banking API) for AISP/PISP access
2. Customer consent (explicit) before granting third-party access
3. SCA applied when customer authorizes third-party access
4. API must provide same data availability as customer's own online access
5. No obstacles to third-party access (no screen scraping blocks without API alternative)
6. API performance: same availability and performance as customer-facing channels
7. Consent dashboard: customer can view and revoke third-party access
8. 90-day re-authentication for AISP access

## Acceptance Criteria
Given an authorized AISP/PISP requesting account access, When the customer provides consent with SCA, Then the API provides account data (AISP) or initiates payment (PISP) AND the customer can view and revoke access AND API availability matches customer-facing channels AND re-authentication occurs every 90 days for AISPs.
