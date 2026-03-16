# Security Requirements Template

---

## Template Fields

| Field | Description |
|---|---|
| **Requirement ID** | SR-[Category]-[Seq] (e.g., SR-AUTH-001) |
| **Title** | Short descriptive title |
| **Category** | Authentication / Authorization / Encryption / Network / Application / Data / Monitoring / Incident Response |
| **Priority** | Critical / High / Medium / Low |
| **Related Functional Req(s)** | Linked FR IDs |
| **Threat/Risk Addressed** | What threat or vulnerability this mitigates (reference OWASP, STRIDE, or risk register) |
| **Description** | Detailed security requirement |
| **Implementation Standard** | Reference standard (OWASP ASVS, NIST 800-53, CIS Benchmarks, ISO 27001) |
| **Technical Specification** | Specific technical implementation details |
| **Data Classification** | Data sensitivity level this protects: Public / Internal / Confidential / Restricted (ref: PM-BP-008) |
| **Attack Vectors Mitigated** | Specific attack types prevented |
| **Testing Requirements** | How this will be validated (SAST, DAST, pen test, manual review) |
| **Monitoring Requirements** | What should be monitored/alerted |
| **Audit Trail Requirements** | Security events to log: event type, fields captured, retention period, tamper-proofing method (ref: PM-BP-009) |
| **Incident Response** | Expected response if this control fails |
| **DR/BCP Considerations** | How this security control operates during failover/recovery; backup of security infrastructure (keys, certs, policies) (ref: PM-BP-010) |
| **Acceptance Criteria** | Pass/fail conditions |
| **Dependencies** | Other security or infrastructure requirements |
| **Exception Process** | How to request an exception if this cannot be met |

---

## Security Categories & Standard Requirements for Banking

### Authentication
| Requirement | Standard | Details |
|---|---|---|
| Multi-factor authentication | PSD2 SCA, NIST 800-63B | Required for all customer-facing and privileged operations |
| Session management | OWASP ASVS 3.x | Secure tokens, 15-min timeout for sensitive ops, absolute timeout 8 hours |
| Password policy | NIST 800-63B | Min 12 chars, breach database check, no periodic rotation forced |
| Account lockout | OWASP | Lock after 5 failed attempts, progressive delay, alert on anomaly |
| Token management | OAuth 2.0 / OIDC | Short-lived access tokens (15 min), secure refresh token rotation |

### Authorization
| Requirement | Standard | Details |
|---|---|---|
| Role-based access control (RBAC) | ISO 27001 A.9 | Least privilege, segregation of duties for financial operations |
| Resource-level authorization | OWASP ASVS 4.x | Every API endpoint validates user authorization for the specific resource |
| Privileged access management | CIS Controls | Just-in-time access, approval workflow, session recording |
| API authorization | OAuth 2.0 scopes | Fine-grained scopes per API operation |

### Encryption
| Requirement | Standard | Details |
|---|---|---|
| Data in transit | TLS 1.3 (min TLS 1.2) | All communications encrypted, no fallback to insecure protocols |
| Data at rest | AES-256 | All PII, financial data, credentials encrypted at rest |
| Key management | NIST 800-57 | HSM for critical keys, automated rotation, separation of duties |
| Database encryption | TDE + column-level | TDE for all databases, column-level for PII/sensitive fields |
| Tokenization | PCI-DSS | Card numbers, account numbers tokenized in non-production and logs |

### Application Security
| Requirement | Standard | Details |
|---|---|---|
| Input validation | OWASP ASVS 5.x | Whitelist validation on all inputs, server-side enforcement |
| Output encoding | OWASP | Context-aware encoding to prevent XSS |
| SQL injection prevention | OWASP | Parameterized queries only, no dynamic SQL |
| CSRF protection | OWASP | Anti-CSRF tokens on all state-changing operations |
| Security headers | OWASP | CSP, X-Frame-Options, HSTS, X-Content-Type-Options |
| Dependency management | OWASP | Automated scanning, patch within 30 days (critical: 24 hours) |

### Monitoring & Logging
| Requirement | Standard | Details |
|---|---|---|
| Security event logging | ISO 27001 A.12.4 | Log auth events, access to sensitive data, admin actions, failures |
| Tamper-proof logs | SOX, PCI-DSS | Immutable centralized logging, integrity verification |
| Real-time alerting | NIST CSF | Alert on brute force, privilege escalation, anomalous access patterns |
| Log retention | Regulatory | 7 years for financial, 3 years for access logs |

---

## Example — Security Requirement

| Field | Value |
|---|---|
| **Requirement ID** | SR-AUTH-003 |
| **Title** | Multi-Factor Authentication for Fund Transfers |
| **Category** | Authentication |
| **Priority** | Critical |
| **Related Functional Req(s)** | FR-PAY-012 (Fund Transfer), FR-AUTH-003 (MFA) |
| **Threat/Risk Addressed** | Account takeover, unauthorized transactions (STRIDE: Spoofing) |
| **Description** | All fund transfer operations must require step-up multi-factor authentication using a second factor independent of the primary channel |
| **Implementation Standard** | PSD2 SCA (Article 97), NIST 800-63B AAL2 |
| **Technical Specification** | TOTP (RFC 6238) or push notification via registered device; SMS OTP as fallback only where regulatory permitted; 90-second OTP validity; single-use enforcement |
| **Data Classification** | Restricted |
| **Attack Vectors Mitigated** | Credential stuffing, session hijacking, phishing, man-in-the-middle |
| **Testing Requirements** | DAST: test bypass attempts; Pen test: MFA bypass scenarios; Unit test: OTP validation logic, expiry, replay prevention |
| **Monitoring Requirements** | Alert on: 3+ failed MFA attempts, MFA bypass attempts, MFA from new device/location, MFA method change |
| **Audit Trail Requirements** | Log all MFA events: user ID, timestamp, MFA method, device fingerprint, IP address, outcome (success/failure/lockout), correlation ID. Retain for 7 years. Logs must be immutable and centrally stored. |
| **Incident Response** | Failed MFA triggers: temporary account lock after 5 attempts, notify customer via registered email/SMS, escalate to fraud team if pattern detected |
| **DR/BCP Considerations** | MFA infrastructure (TOTP seed storage, push notification service) must be replicated to DR site. HSM containing TOTP seeds must have DR backup. During failover, MFA must remain enforced — no bypass permitted. Fallback to SMS OTP if push notification service is unavailable in DR. |
| **Acceptance Criteria** | Given an authenticated user initiating a fund transfer, When step-up MFA is triggered, Then the user must provide a valid second factor AND expired/replayed OTPs are rejected AND 5 consecutive failures lock the operation for 30 minutes |
| **Dependencies** | SR-AUTH-001 (Primary Authentication), SR-ENC-002 (TLS for OTP delivery) |
| **Exception Process** | No exceptions permitted for fund transfers; exceptions for other operations require CISO + Compliance Officer approval with compensating controls documented |

---

## Usage Guidelines

1. **Every requirement must have a unique ID** following the SR-[Category]-[Seq] naming convention
2. **Threat/risk addressed is mandatory** — reference OWASP, STRIDE, or internal risk register
3. **Implementation standard must be cited** — ensures alignment with industry frameworks
4. **Testing requirements must specify method** — SAST, DAST, pen test, or manual review
5. **Cross-reference** to related Functional (FR) and Compliance (CR) requirements
6. **Review cadence**: Quarterly + on regulatory change
7. **Approval**: Product Owner + Security Engineer sign-off required
