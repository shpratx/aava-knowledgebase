# Consumer Rights Act 2015 — Unfair Terms (Part 2, s62-69)

## Regulation Reference
- Act: Consumer Rights Act 2015
- Sections: 62-69 + Schedule 2 (Grey List)
- Enforcing Body: CMA / FCA

## Obligation (Plain Language)
Contract terms must be fair. A term is unfair if it causes a significant imbalance in the parties' rights and obligations to the consumer's detriment. Unfair terms are not binding on the consumer. Terms must be transparent (plain language) and prominent.

## Grey List — Potentially Unfair Terms (Schedule 2)
| Term Type | Example in Banking | Risk |
|---|---|---|
| Excluding/limiting liability | "We are not liable for any loss from system outage" | Unfair if blanket exclusion |
| Unilateral variation | "We may change any term at any time" | Unfair without reasonable notice and right to exit |
| Unilateral termination | "We may close your account at any time without reason" | Unfair without reasonable notice |
| Automatic renewal | "Contract auto-renews unless you cancel 60 days before" | Unfair if unreasonable notice period |
| Disproportionate penalties | "£50 fee for any returned direct debit" | Unfair if disproportionate to cost |
| Binding arbitration | "All disputes must go to arbitration" | Unfair if excludes court access |

## Technical Controls Required
1. **T&C plain language:** All terms in plain, intelligible language; readability scoring
2. **Key terms prominence:** Important terms (fees, liability, variation, termination) highlighted in UI
3. **Change notification:** Advance notice of term changes (minimum 30 days; 60 days for material changes)
4. **Right to exit:** Customer can exit without penalty if terms change materially
5. **Fee transparency:** All fees clearly disclosed before commitment; no hidden charges
6. **Cancellation ease:** Cancellation process no more difficult than sign-up (aligns with FCA Consumer Duty)
7. **Grey list review:** All T&C reviewed against Schedule 2 grey list before publication

## Evidence Required
- T&C readability assessment
- UI screenshots showing key term prominence
- Change notification records with advance notice
- Cancellation journey analysis (friction comparison)
- Grey list review records
- Customer complaint analysis related to terms

## Acceptance Criteria
Given banking terms and conditions,
When assessed against unfair terms provisions,
Then all terms are in plain, intelligible language
  AND key terms (fees, liability, variation) are prominently displayed
  AND term changes are notified with minimum 30 days advance notice
  AND customers can exit without penalty on material term changes
  AND cancellation is as easy as sign-up
  AND all terms are reviewed against the Schedule 2 grey list.
