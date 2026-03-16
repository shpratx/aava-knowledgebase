# Guardrails for Security Requirements

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-SG-001 | Every SR must have a unique ID following SR-[Category]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| SR-SG-002 | Every SR must specify a Category (Authentication/Authorization/Encryption/Network/Application/Data/Monitoring/Incident Response) | Reject if Category is empty or non-standard |
| SR-SG-003 | Every SR must reference the Threat/Risk Addressed using STRIDE, OWASP, or risk register | Reject if Threat/Risk Addressed is empty |
| SR-SG-004 | Every SR must cite an Implementation Standard (OWASP ASVS, NIST, CIS, ISO 27001) | Reject if Implementation Standard is empty |
| SR-SG-005 | Every SR must specify Testing Requirements (SAST, DAST, pen test, manual review) | Reject if Testing Requirements is empty |
| SR-SG-006 | Every SR must specify Monitoring Requirements | Reject if Monitoring Requirements is empty |
| SR-SG-007 | Every SR must define Incident Response actions if the control fails | Reject if Incident Response is empty |
| SR-SG-008 | Every SR must have testable Acceptance Criteria including both positive and bypass-attempt tests | Reject if Acceptance Criteria is empty or tests only positive path |

---

## 2. Authentication Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-AuthG-001 | All customer-facing applications must require multi-factor authentication for financial transactions | Reject if financial transaction SR does not mandate MFA |
| SR-AuthG-002 | Session timeout must not exceed 15 minutes of inactivity for sensitive operations | Reject if session timeout exceeds 15 minutes for sensitive ops |
| SR-AuthG-003 | Absolute session timeout must not exceed 8 hours | Reject if absolute timeout exceeds 8 hours |
| SR-AuthG-004 | Password requirements must align with NIST 800-63B (min 12 chars, breach database check) | Flag if password policy contradicts NIST 800-63B |
| SR-AuthG-005 | Account lockout must activate after maximum 5 consecutive failed attempts | Reject if lockout threshold exceeds 5 attempts |
| SR-AuthG-006 | OTP validity must not exceed 90 seconds and must be single-use | Reject if OTP validity exceeds 90 seconds or allows reuse |
| SR-AuthG-007 | SMS OTP must only be used as fallback where regulatory permitted — TOTP or FIDO2 preferred | Flag if SMS OTP is the primary MFA method |
| SR-AuthG-008 | Authentication tokens must use cryptographically random generation with minimum 128-bit entropy | Reject if token generation spec is weaker |

---

## 3. Authorization Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-AuthzG-001 | Every API endpoint must enforce authorization at the resource level (not just role level) | Reject if authorization is role-only without resource-level checks |
| SR-AuthzG-002 | Authorization must follow principle of least privilege — users get minimum permissions needed | Reject if authorization grants broader access than required |
| SR-AuthzG-003 | Segregation of duties must be enforced for financial operations (e.g., maker-checker) | Reject if financial operation SR has no segregation of duties |
| SR-AuthzG-004 | Privileged access must be just-in-time with approval workflow and session recording | Flag if privileged access is permanent or unmonitored |
| SR-AuthzG-005 | Authorization decisions must be logged (who requested, what resource, decision, timestamp) | Reject if authorization logging is not specified |
| SR-AuthzG-006 | RBAC role definitions must be reviewed quarterly and certified by business owners | Process guardrail — enforce via scheduled review |

---

## 4. Encryption Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-EncG-001 | All data in transit must use TLS 1.2 minimum, TLS 1.3 preferred — SSL and TLS 1.0/1.1 are prohibited | Reject if any protocol below TLS 1.2 is permitted |
| SR-EncG-002 | All PII and financial data at rest must be encrypted with AES-256 or equivalent | Reject if at-rest encryption is weaker than AES-256 |
| SR-EncG-003 | Encryption keys must be managed via HSM (FIPS 140-2 Level 3) in production | Reject if production keys are not HSM-managed |
| SR-EncG-004 | Encryption keys must be rotated at minimum every 90 days | Flag if key rotation period exceeds 90 days |
| SR-EncG-005 | Card data (PAN) must be tokenized in all non-production environments and logs | Reject if PAN appears untokenized outside production processing |
| SR-EncG-006 | Password storage must use bcrypt (cost 12+) or Argon2id — MD5, SHA-1, plain SHA-256 are prohibited | Reject if prohibited hashing algorithm is specified |
| SR-EncG-007 | Encryption implementation must not use custom/proprietary algorithms — only industry-standard algorithms | Reject if custom cryptography is proposed |

---

## 5. Application Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-AppG-001 | All user inputs must be validated server-side using whitelist approach | Reject if input validation is client-side only or uses blacklist |
| SR-AppG-002 | All database queries must use parameterized queries — dynamic SQL construction is prohibited | Reject if dynamic SQL is permitted |
| SR-AppG-003 | All state-changing operations must have CSRF protection | Reject if CSRF protection is not specified for state-changing ops |
| SR-AppG-004 | Security headers must be implemented: CSP, HSTS (max-age ≥ 31536000), X-Frame-Options, X-Content-Type-Options | Reject if mandatory security headers are missing |
| SR-AppG-005 | Error responses must not expose stack traces, internal paths, database details, or system information | Reject if error handling permits information leakage |
| SR-AppG-006 | Dependencies must be scanned for known vulnerabilities — critical CVEs must be patched within 24 hours, high within 7 days | Reject if no patching SLA is defined |
| SR-AppG-007 | Source maps, debug endpoints, and development tools must not be deployed to production | Reject if production deployment includes debug artifacts |
| SR-AppG-008 | File upload functionality must validate file type, size, and content — executable uploads are prohibited | Reject if file upload has no validation specified |

---

## 6. Data Protection Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-DataG-001 | Sensitive data must not appear in URLs, query parameters, or browser history | Reject if sensitive data is transmitted via URL |
| SR-DataG-002 | Sensitive data must not be stored in client-side storage (localStorage, sessionStorage, cookies without secure/httpOnly flags) | Reject if client-side sensitive data storage is proposed |
| SR-DataG-003 | Sensitive data must not appear in application logs in clear text | Reject if logging includes unmasked sensitive data |
| SR-DataG-004 | Production data must not be used in non-production environments without anonymization/masking | Reject if production data use in lower environments is proposed |
| SR-DataG-005 | Data masking must be applied to sensitive fields in all non-production displays and exports | Flag if no masking strategy is defined |
| SR-DataG-006 | Secrets (API keys, credentials, certificates) must be stored in secure vaults — never in code, config files, or environment variables in plain text | Reject if secrets are stored outside secure vaults |

---

## 7. Monitoring & Logging Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-MonG-001 | All authentication events (success, failure, lockout) must be logged | Reject if auth event logging is not specified |
| SR-MonG-002 | All access to sensitive data must be logged (who, what, when, where) | Reject if sensitive data access logging is missing |
| SR-MonG-003 | All privileged/admin actions must be logged and reviewed | Reject if privileged action logging is missing |
| SR-MonG-004 | Audit logs must be immutable and stored separately from application data | Reject if log immutability is not specified |
| SR-MonG-005 | Security alerts must have defined response times: Critical < 15 minutes, High < 1 hour, Medium < 4 hours | Flag if alert response times are not defined |
| SR-MonG-006 | Log retention must meet regulatory requirements (7 years financial, 3 years access) | Reject if retention is below regulatory minimum |
| SR-MonG-007 | Logs must not contain sensitive data (passwords, tokens, PII) in clear text | Reject if log content includes unmasked sensitive data |
| SR-MonG-008 | Real-time alerting must be configured for: brute force, privilege escalation, anomalous data access, configuration changes | Flag if real-time alerting is not specified for these events |

---

## 8. Network Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-NetG-001 | Production networks must be segmented from non-production | Reject if no network segmentation is specified |
| SR-NetG-002 | Database servers must not be directly accessible from the internet | Reject if database has public network exposure |
| SR-NetG-003 | Service-to-service communication must use mutual TLS (mTLS) where feasible | Flag if mTLS is not specified for internal service communication |
| SR-NetG-004 | Firewall rules must follow default-deny principle | Reject if default-allow is proposed |
| SR-NetG-005 | API gateways must be used for all external-facing APIs | Flag if external APIs bypass API gateway |

---

## 9. Incident Response Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-IRG-001 | Every SR must define what happens when the security control fails or is breached | Reject if no incident response is defined |
| SR-IRG-002 | Security incidents must be escalated within 1 hour of detection | Reject if escalation time exceeds 1 hour |
| SR-IRG-003 | Containment actions must include both automated (immediate) and manual (follow-up) responses | Flag if containment is manual-only |
| SR-IRG-004 | Post-incident review must be conducted within 48 hours | Flag if no post-incident review timeline is specified |
| SR-IRG-005 | Incident evidence must be preserved for forensic analysis before any remediation | Reject if evidence preservation is not addressed |

---

## 10. Exception & Waiver Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-EWG-001 | Every SR must define an Exception Process | Reject if Exception Process is empty |
| SR-EWG-002 | Security exceptions require CISO + Compliance Officer approval at minimum | Reject if exception approval is below CISO level |
| SR-EWG-003 | Exceptions must have compensating controls documented and tested | Reject if exception has no compensating controls |
| SR-EWG-004 | Exceptions must have an expiry date not exceeding 6 months | Reject if exception has no expiry or exceeds 6 months |
| SR-EWG-005 | No exceptions are permitted for: encryption of card data, MFA for fund transfers, audit logging for financial transactions | Reject any exception request for these controls |
| SR-EWG-006 | All exceptions must be tracked in a central exception register and reviewed monthly | Process guardrail — enforce via exception management |

---

## 11. Data Classification Guardrails (ref: PM-BP-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-DCG-001 | Every SR must specify the Data Classification levels it protects using the four-level scheme (Public/Internal/Confidential/Restricted) | Reject if Data Classification uses generic description instead of explicit levels |
| SR-DCG-002 | Security controls must be proportional to data classification — Restricted requires strongest controls | Reject if Restricted data has weaker controls than Confidential |
| SR-DCG-003 | SRs for Restricted data must mandate: AES-256 encryption, HSM key management, MFA, full access logging, real-time monitoring | Reject if any of these controls are missing for Restricted data |
| SR-DCG-004 | SRs for Confidential data must mandate: AES-256 encryption, MFA for modifications, modification logging | Reject if these controls are missing for Confidential data |
| SR-DCG-005 | Security exceptions are not permitted for Restricted data controls | Reject any exception request for Restricted data SRs |
| SR-DCG-006 | Data classification must drive the incident response urgency: Restricted < 15 min, Confidential < 1 hour | Flag if response times are not tiered by classification |

---

## 12. Audit Trail Guardrails (ref: PM-BP-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-ATG-001 | Every SR must have an Audit Trail Requirements field specifying security events to log, fields to capture, and retention | Reject if Audit Trail Requirements is empty |
| SR-ATG-002 | Audit trail must be separate from monitoring — monitoring is real-time detection, audit trail is forensic evidence | Flag if monitoring and audit are conflated without distinction |
| SR-ATG-003 | Security audit logs must specify tamper-proofing method (cryptographic hashing, WORM storage) | Reject if no tamper-proofing method is specified |
| SR-ATG-004 | Audit log fields must include at minimum: event type, timestamp (UTC), user ID, source IP, action, resource, outcome, correlation ID | Reject if mandatory log fields are missing |
| SR-ATG-005 | Security audit logs must be accessible for forensic analysis within 5 minutes of query | Flag if no log query performance target is specified |
| SR-ATG-006 | Audit logs must not contain sensitive data in clear text — passwords, tokens, PII must be masked | Reject if log specification includes unmasked sensitive data |

---

## 13. DR/BCP Guardrails (ref: PM-BP-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-DRG-001 | Every SR must specify DR/BCP Considerations — how the security control operates during failover | Reject if DR/BCP Considerations is empty |
| SR-DRG-002 | Security controls must not be weakened or bypassed during DR/BCP events | Reject if DR mode permits security control bypass |
| SR-DRG-003 | MFA must remain enforced during DR — no emergency bypass for financial transactions | Reject if MFA bypass is permitted during DR |
| SR-DRG-004 | Security infrastructure (HSM, key vault, IAM, certificates) must be replicated to DR site | Reject if security infrastructure DR replication is not specified |
| SR-DRG-005 | SIEM/monitoring/log collection must continue during DR — no security blind spots during failover | Reject if monitoring continuity during DR is not addressed |
| SR-DRG-006 | DR activation must include a security validation checklist — verify all controls are active before accepting traffic | Flag if no security validation checklist is defined for DR activation |
| SR-DRG-007 | Encryption keys and certificates in DR must be pre-provisioned and validated — no expired certs on failover | Reject if DR certificate/key management is not addressed |
| SR-DRG-008 | Network security in DR must mirror production — same segmentation, firewall rules, no relaxed policies | Reject if DR network security is weaker than production |

---

## 14. Testing & Validation Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| SR-TVG-001 | SAST must run on every build — critical/high findings block merge | Enforce as CI/CD gate |
| SR-TVG-002 | DAST must run on staging before every production release | Enforce as release gate |
| SR-TVG-003 | Dependency scanning (SCA) must run on every build — critical CVEs block merge | Enforce as CI/CD gate |
| SR-TVG-004 | Penetration testing must be conducted annually and after major architectural changes | Process guardrail — enforce via security calendar |
| SR-TVG-005 | Security test results must be reviewed by the security team before release sign-off | Enforce as release gate |
| SR-TVG-006 | All critical/high security findings must be remediated before production deployment — no exceptions | Enforce as release gate |
| SR-TVG-007 | Security acceptance criteria must include bypass-attempt tests, not just positive verification | Reject if acceptance criteria only test positive path |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | SR cannot proceed until fixed | Missing threat reference, no testing requirements, weak encryption, no incident response |
| **Flag** | SR can proceed but must be addressed before release | SMS as primary MFA, no mTLS for internal services, no degraded-mode response |
| **Review** | SR requires additional review by specified role | Security Engineer for all SRs, CISO for exceptions, Compliance for regulatory SRs |
| **CI/CD Gate** | Automated enforcement in pipeline | SAST/SCA on every build, DAST before release, critical findings block deployment |
