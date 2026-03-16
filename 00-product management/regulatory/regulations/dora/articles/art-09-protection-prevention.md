# DORA Article 9 — Protection and Prevention

## Regulation Reference
- Regulation: DORA (2022/2554) — Article 9

## Obligation
Financial entities must implement ICT security measures to ensure the protection and prevention of ICT risks, including policies on access control, encryption, network security, and vulnerability management.

## Technical Controls
1. Access control: RBAC, least privilege, MFA, privileged access management
2. Encryption: data at rest and in transit per data classification
3. Network security: segmentation, firewalls, IDS/IPS, mTLS for internal
4. Vulnerability management: scanning, patching (critical within 24h), pen testing
5. Secure development: SAST/DAST, code review, secure coding standards
6. Physical security: data center access controls (where applicable)
7. Endpoint protection: anti-malware, EDR, device management
8. Data loss prevention: DLP controls for sensitive data

## Acceptance Criteria
Given ICT systems processing financial data, Then access control follows least privilege with MFA AND data is encrypted at rest and in transit AND networks are segmented with IDS/IPS AND vulnerabilities are scanned continuously and patched per SLA AND development follows secure coding practices with SAST/DAST.
