# PSD2 Articles 73-74 — Liability for Unauthorized Transactions

## Regulation Reference
- Directive: PSD2 (2015/2366) — Articles 73-74

## Obligation
PSPs must refund unauthorized payment transactions immediately (by end of next business day at latest) unless they have reasonable grounds to suspect fraud by the payer. Payer's liability is limited to €50 for unauthorized transactions where SCA was not applied.

## Technical Controls
1. Unauthorized transaction detection: real-time monitoring for unusual patterns
2. Customer dispute workflow: intake, investigation, provisional credit, resolution
3. Refund automation: immediate refund (D+1) unless fraud suspected
4. SCA enforcement: if SCA was not applied, PSP bears full liability
5. Fraud investigation: documented process for suspected customer fraud
6. Notification: customer notified of unauthorized transaction detection
7. Liability calculation: €50 max payer liability if no gross negligence; zero if SCA not applied by PSP

## Acceptance Criteria
Given an unauthorized transaction reported by customer, When the dispute is filed, Then provisional refund is issued by end of next business day AND if SCA was not applied by PSP, full refund with zero payer liability AND investigation is completed within regulatory timeline AND customer is kept informed throughout.
