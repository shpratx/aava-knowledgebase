# PCI-DSS Requirement 3 — Protect Stored Account Data

## Regulation Reference
- Standard: PCI-DSS v4.0
- Requirement: 3
- Scope: All systems storing cardholder data within the CDE

## Obligation (Plain Language)
Stored account data must be protected. PAN must be rendered unreadable anywhere it is stored. Sensitive authentication data (CVV, PIN, track data) must NEVER be stored after authorization.

## Technical Controls Required
1. **PAN rendering:** Encrypt (AES-256), hash (SHA-256 with salt), truncate (first 6 + last 4), or tokenize
2. **CVV/PIN/Track data:** Never stored post-authorization — no database column, no log, no cache, no backup
3. **Encryption:** AES-256 for PAN at rest; key management per Req 3.6
4. **Tokenization:** Replace PAN with non-reversible token for all non-CDE systems, logs, and non-production environments
5. **Key management:** HSM (FIPS 140-2 Level 3); split knowledge; dual control; 90-day DEK rotation; annual master key rotation
6. **Data retention:** Define retention period; securely delete when no longer needed
7. **Display masking:** Show only first 6 + last 4 digits (BIN + last 4); full PAN only for authorized business need

## Evidence Required
- Encryption implementation verification (storage inspection)
- Key management procedures and HSM configuration
- Data flow diagram showing where PAN is stored/processed/transmitted
- Tokenization implementation for non-CDE systems
- Data retention policy with automated enforcement
- Verification that CVV/PIN/track data is not stored post-authorization

## Acceptance Criteria
Given a system storing cardholder data,
When assessed for PCI-DSS Req 3 compliance,
Then PAN is encrypted with AES-256 at rest
  AND CVV, PIN, and track data are never stored post-authorization
  AND PAN is tokenized in all non-CDE systems and logs
  AND encryption keys are managed via HSM with split knowledge
  AND PAN display is masked (first 6 + last 4) unless business need
  AND data retention policy is defined and automated.

## Penalties
- Fines up to $100,000/month; loss of card processing capability
