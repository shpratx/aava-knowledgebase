# Cross-Reference: Regulatory Conflicts and Resolutions

## Conflict 1: GDPR Erasure vs AML Retention
| Regulation | Requirement |
|---|---|
| GDPR Art. 17 | Erase personal data within 30 days of request |
| AML Directive Art. 40 | Retain transaction records for 5 years post-relationship |

**Resolution:**
- AML retention prevails for transaction records (GDPR Art. 17(3)(b) — legal obligation exemption)
- Anonymize the customer identifier in retained records (replace with non-reversible hash)
- Retain the transaction with anonymized reference — satisfies AML without retaining PII
- Document the legal basis per retained field
- Inform the data subject that certain records are retained under legal obligation

## Conflict 2: GDPR Minimization vs KYC Data Collection
| Regulation | Requirement |
|---|---|
| GDPR Art. 5(1)(c) | Collect only data adequate, relevant, and limited to purpose |
| AML Directive Art. 11-14 | Collect comprehensive identity data for CDD |

**Resolution:**
- KYC data collection is justified under GDPR Art. 6(1)(c) — legal obligation
- Collect only what AML regulation requires — no additional "nice to have" fields
- Document the legal obligation basis per field in the processing register
- Apply data minimization to non-KYC data (marketing, preferences)

## Conflict 3: GDPR Consent vs Banking Contractual Necessity
| Regulation | Requirement |
|---|---|
| GDPR Art. 7 | Consent must be freely given; withdrawable |
| Banking contract | Core services require data processing |

**Resolution:**
- Use contractual necessity (Art. 6(1)(b)) for core banking services — NOT consent
- Consent only for optional processing (marketing, analytics, third-party sharing)
- This avoids the problem of customers withdrawing consent for essential services
- Clearly separate mandatory (contractual) from optional (consent) processing in privacy notice

## Conflict 4: PCI-DSS Data Destruction vs SOX Audit Trail
| Regulation | Requirement |
|---|---|
| PCI-DSS Req 3 | Securely delete card data when no longer needed |
| SOX Sec 404 | Maintain audit trail of financial transactions |

**Resolution:**
- Tokenize card data after authorization — token satisfies audit trail without storing PAN
- Retain transaction records with token reference (not PAN) for SOX compliance
- Securely delete actual PAN per PCI-DSS retention policy
- Audit trail references the token, which can be resolved only within the CDE if needed

## Conflict 5: GDPR Cross-Border Transfer Restrictions vs Global Banking Operations
| Regulation | Requirement |
|---|---|
| GDPR Art. 44-49 | Restrict transfers to non-adequate countries |
| Banking operations | Global payment processing, correspondent banking |

**Resolution:**
- Use SCCs for processor transfers to non-adequate countries
- Use contractual necessity derogation (Art. 49(1)(b)) for customer-initiated cross-border payments
- Conduct Transfer Impact Assessment for each destination country
- Implement supplementary measures (encryption, access controls)
- Document transfer mechanism per data flow in the transfer inventory


## UK-Specific Conflicts

### Conflict 6: DPA 2018 Crime Prevention Exemption vs Data Subject Rights
| Regulation | Requirement |
|---|---|
| UK GDPR Art. 15 | Data subject right of access |
| DPA 2018 Sch 2 Para 2 | Exemption from data subject rights for crime prevention |

**Resolution:**
- When sharing data with law enforcement for fraud/AML investigation, the crime prevention exemption applies
- Data subject access requests can be refused/restricted if disclosure would prejudice crime prevention
- Document the exemption assessment per request
- Exemption is not blanket — assess on case-by-case basis

### Conflict 7: FCA Consumer Duty (Easy Cancellation) vs AML Retention
| Regulation | Requirement |
|---|---|
| FCA Consumer Duty | Cancellation as easy as sign-up; no unreasonable barriers |
| AML Directive / MLR 2017 | Retain customer records for 5 years post-relationship |

**Resolution:**
- Account closure process must be straightforward (Consumer Duty)
- BUT customer data retained for 5 years post-closure (AML)
- Inform customer that account is closed but certain records retained under legal obligation
- Anonymize non-required data; retain only what AML mandates

### Conflict 8: Equality Act Accessibility vs Security Controls
| Regulation | Requirement |
|---|---|
| Equality Act s20-21 | Reasonable adjustments; accessible authentication |
| PSD2 Art. 97 / FCA | Strong Customer Authentication (two factors) |

**Resolution:**
- SCA must be accessible — provide multiple factor options
- If biometric is not suitable (disability), provide accessible alternative (PIN + device possession)
- SMS OTP as accessible fallback (with security trade-off documented)
- Never exempt from SCA — provide accessible SCA, not no SCA
