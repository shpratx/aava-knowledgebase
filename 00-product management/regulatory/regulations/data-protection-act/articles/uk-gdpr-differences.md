# DPA 2018 — Key Differences from EU GDPR for Banking

## Regulation Reference
- Act: Data Protection Act 2018
- Supplementing: UK GDPR (retained EU law)
- Enforcing Body: ICO

## Key Differences

### 1. International Transfers
| Aspect | EU GDPR | UK GDPR + DPA 2018 |
|---|---|---|
| Adequacy decisions | EU Commission decides | UK Secretary of State decides |
| EU adequacy | N/A | UK has EU adequacy decision (until June 2025; review pending) |
| Transfer mechanisms | EU SCCs | UK International Data Transfer Agreement (IDTA) or UK Addendum to EU SCCs |
| Transitional provisions | N/A | UK-specific transitional bridge for existing transfers |

### 2. Age of Consent for Children
| Aspect | EU GDPR | UK DPA 2018 |
|---|---|---|
| Age threshold | 16 (member states can lower to 13) | 13 years (s9 DPA 2018) |

### 3. Crime Prevention Exemption
| Aspect | EU GDPR | UK DPA 2018 |
|---|---|---|
| Exemption | Limited | Broader exemption for crime prevention/detection (Sch 2 Para 2) |
| Banking use | Limited data sharing | Can share with law enforcement without consent for fraud/AML |

### 4. Regulatory Functions Exemption
| Aspect | EU GDPR | UK DPA 2018 |
|---|---|---|
| Exemption | Not explicit | Explicit exemption for regulatory functions (Sch 2 Para 7) |
| Banking use | N/A | Can share with FCA/PRA without data subject consent for regulatory purposes |

## Technical Controls Required
1. **UK IDTA/Addendum:** Use UK-specific transfer mechanisms for international transfers (not EU SCCs alone)
2. **Age verification:** Implement age gate at 13 (not 16) for consent-based processing
3. **Crime prevention processing:** Implement lawful basis for fraud data sharing with law enforcement (DPA 2018 Sch 2 Para 2)
4. **Regulatory reporting:** Implement lawful basis for data sharing with FCA/PRA (DPA 2018 Sch 2 Para 7)
5. **ICO registration:** Maintain ICO registration (fee paid annually)
6. **UK representative:** If processing UK data from outside UK, appoint UK representative

## Evidence Required
- UK IDTA or UK Addendum for international transfers
- Age verification mechanism for consent-based processing
- Lawful basis documentation for crime prevention data sharing
- ICO registration confirmation
- UK representative appointment (if applicable)

## Acceptance Criteria
Given a UK banking operation processing personal data,
When assessed for DPA 2018 compliance,
Then international transfers use UK IDTA or UK Addendum (not EU SCCs alone)
  AND age of consent is set at 13 for consent-based processing
  AND crime prevention data sharing has documented lawful basis (Sch 2 Para 2)
  AND regulatory data sharing has documented lawful basis (Sch 2 Para 7)
  AND ICO registration is current.
