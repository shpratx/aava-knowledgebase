# Equality Act 2010 — Duty to Make Reasonable Adjustments (s20-21)

## Regulation Reference
- Act: Equality Act 2010
- Sections: 20-21
- Enforcing Body: EHRC / courts

## Obligation (Plain Language)
Service providers must make reasonable adjustments to avoid putting disabled people at a substantial disadvantage compared to non-disabled people. This is an anticipatory duty — firms must plan ahead, not wait for individual requests.

## Three Requirements
| Requirement | Description | Banking IT Implementation |
|---|---|---|
| Change a provision, criterion, or practice | Modify policies or procedures that disadvantage disabled people | Flexible authentication (not solely biometric); alternative to phone-only support |
| Remove or alter a physical feature | Modify physical barriers | ATM accessibility; branch accessibility (not primarily IT) |
| Provide an auxiliary aid or service | Provide additional support | Screen reader compatibility; large print; BSL video relay; text relay |

## Technical Controls Required
1. **WCAG 2.1 AA compliance:** All digital services meet accessibility standards (anticipatory duty)
2. **Screen reader support:** Full compatibility with NVDA, JAWS, VoiceOver, TalkBack
3. **Keyboard navigation:** All functionality accessible via keyboard alone
4. **Alternative formats:** Statements in large print, Braille, audio (on request)
5. **Flexible authentication:** If biometric fails/unavailable, provide accessible alternative (PIN, password)
6. **Color independence:** No information conveyed by color alone
7. **Cognitive accessibility:** Plain language; clear navigation; consistent layout; error prevention
8. **Communication preferences:** Record and respect customer communication preferences (large print, email, BSL)
9. **Assistive technology testing:** Test with screen readers, magnifiers, voice control, switch access
10. **Adjustment request workflow:** System to record, track, and fulfill reasonable adjustment requests

## Evidence Required
- WCAG 2.1 AA audit results (automated + manual)
- Screen reader testing results for critical journeys
- Keyboard navigation testing results
- Alternative format request and fulfillment records
- Adjustment request tracking records
- Assistive technology compatibility test results
- Staff training records on disability awareness

## Acceptance Criteria
Given a disabled customer accessing banking services,
When they use digital channels,
Then the service meets WCAG 2.1 AA standards
  AND all functionality is keyboard-accessible
  AND screen readers can navigate all critical journeys (login, balance, transfer, payment)
  AND alternative authentication is available if biometric is not suitable
  AND alternative format statements are available on request
  AND reasonable adjustment requests are tracked and fulfilled.
