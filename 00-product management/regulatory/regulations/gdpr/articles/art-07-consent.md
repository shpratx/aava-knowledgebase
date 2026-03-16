# GDPR Article 7 — Conditions for Consent

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 7
- Enforcing Body: National DPAs

## Obligation (Plain Language)
When consent is the legal basis for processing, it must be freely given, specific, informed, and unambiguous. The controller must be able to demonstrate consent was given. Withdrawal must be as easy as giving consent.

## Technical Controls Required
1. Consent checkboxes must be unchecked by default — no pre-ticked boxes
2. Consent must be granular — separate consent per purpose (marketing email, SMS, analytics, third-party sharing)
3. Consent must not be bundled with service access — service works without optional consent
4. Consent record must capture: data subject ID, timestamp, purpose, mechanism, privacy notice version
5. Withdrawal mechanism must be accessible from account settings — one-click toggle
6. Processing must stop immediately upon withdrawal — no grace period
7. Consent records retained for duration of processing + 3 years

## Applicability
- **Applies when:** Consent is used as legal basis (Art. 6(1)(a))
- **Banking use cases:** Marketing communications, analytics/profiling, third-party data sharing, cookie consent

## Evidence Required
- Consent collection UI screenshots showing unchecked defaults
- Consent withdrawal UI showing equal ease
- Consent database records with all required fields
- Processing stop verification after withdrawal
- Consent record retention policy

## Acceptance Criteria
Given a customer providing consent for marketing emails,
When they check the consent checkbox and submit,
Then consent is recorded with timestamp, purpose, mechanism, and privacy notice version
  AND the checkbox was unchecked by default
  AND the service functions without this consent
  AND the customer can withdraw via account settings with one click
  AND processing stops immediately upon withdrawal.

## Penalties
- Up to €20M or 4% of annual global turnover
