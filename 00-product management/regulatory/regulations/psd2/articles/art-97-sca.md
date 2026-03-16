# PSD2 Article 97 — Strong Customer Authentication (SCA)

## Regulation Reference
- Directive: PSD2 (2015/2366)
- Article: 97
- RTS: Commission Delegated Regulation (EU) 2018/389
- Enforcing Body: EBA / national authorities

## Obligation (Plain Language)
Payment service providers must apply Strong Customer Authentication when the payer initiates an electronic payment, accesses their account online, or carries out any action through a remote channel which may imply a risk of payment fraud.

## SCA Requirements
SCA requires TWO of the following THREE factors:
| Factor | Category | Banking Implementation |
|---|---|---|
| Something the user knows | Knowledge | Password, PIN, security question |
| Something the user has | Possession | Mobile device (push notification, TOTP), hardware token, SMS OTP (fallback only) |
| Something the user is | Inherence | Fingerprint, face recognition, voice recognition |

The two factors must be from DIFFERENT categories and must be independent (compromise of one does not compromise the other).

## When SCA is Required
| Action | SCA Required | Exemption Possible |
|---|---|---|
| Online payment initiation | Yes | See exemptions below |
| Account access (login to online banking) | Yes | 90-day re-authentication exemption |
| Adding/modifying beneficiary | Yes | No exemption |
| Setting up standing order | Yes (first time) | No |
| Contactless payment < €50 | No (exemption) | Cumulative limit: €150 or 5 transactions |

## SCA Exemptions (RTS Art. 10-18)
| Exemption | Condition |
|---|---|
| Low-value transactions | < €30 (cumulative < €100 or 5 transactions) |
| Trusted beneficiaries | Customer-maintained whitelist |
| Recurring transactions | Same amount, same payee (SCA on first) |
| Transaction Risk Analysis (TRA) | Based on fraud rate thresholds per PSP |
| Secure corporate payments | Dedicated corporate payment processes |

## Technical Controls Required
1. MFA implementation: TOTP (RFC 6238) or push notification as primary; SMS OTP as fallback only
2. Dynamic linking: authentication code linked to specific amount and payee (Art. 97(2))
3. Independence: compromise of one factor must not compromise the other
4. OTP validity: maximum 90 seconds; single-use
5. Session binding: SCA result bound to the specific transaction
6. Exemption engine: evaluate exemptions per RTS; apply SCA when no exemption applies
7. Fallback: if primary MFA unavailable, provide alternative (not weaker) method

## Evidence Required
- MFA implementation showing two independent factors from different categories
- Dynamic linking verification (auth code tied to amount + payee)
- OTP configuration (90-second expiry, single-use)
- Exemption engine logic and configuration
- Transaction logs showing SCA applied or exemption reason

## Acceptance Criteria
Given a customer initiating an online payment,
When the payment is submitted,
Then SCA is applied with two factors from different categories
  AND the authentication code is dynamically linked to the amount and payee
  AND OTP expires after 90 seconds and is single-use
  AND if an exemption applies, it is documented in the transaction log
  AND 5 consecutive SCA failures lock the payment function for 30 minutes.
