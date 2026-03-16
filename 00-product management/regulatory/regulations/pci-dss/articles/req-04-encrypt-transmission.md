# PCI-DSS Requirement 4 — Protect Cardholder Data During Transmission

## Regulation Reference
- Standard: PCI-DSS v4.0
- Requirement: 4

## Obligation (Plain Language)
Cardholder data must be protected with strong cryptography during transmission over open, public networks.

## Technical Controls Required
1. TLS 1.2 minimum; TLS 1.3 preferred — SSL, TLS 1.0, TLS 1.1 prohibited
2. Strong cipher suites only — no RC4, DES, 3DES, MD5-based, export-grade
3. Trusted certificates from recognized CAs
4. Certificate pinning for mobile applications communicating with payment APIs
5. No PAN transmission via end-user messaging (email, chat, SMS)
6. Wireless networks transmitting cardholder data must use strong encryption (WPA2/WPA3)

## Evidence Required
- TLS configuration scan results (SSL Labs A+ rating)
- Cipher suite configuration
- Certificate management procedures
- Network diagram showing encrypted transmission paths
- Mobile app certificate pinning verification

## Acceptance Criteria
Given cardholder data transmitted over any network,
When the transmission is inspected,
Then TLS 1.2+ is used with strong cipher suites
  AND certificates are valid and from trusted CAs
  AND no PAN is sent via email, chat, or SMS
  AND mobile apps use certificate pinning for payment APIs.
