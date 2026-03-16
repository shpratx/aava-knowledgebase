# Best Practices for Composing Security Requirements

---

## 1. Writing the Requirement ID & Title

**Best Practices:**
- Use consistent naming: \`SR-[Category]-[Seq]\` (e.g., SR-AUTH-001, SR-ENC-003, SR-APP-012)
- Standard categories: AUTH (Authentication), AUTHZ (Authorization), ENC (Encryption), NET (Network), APP (Application), DATA (Data Protection), MON (Monitoring), IR (Incident Response)
- Titles should state the security control and scope: "Multi-Factor Authentication for Fund Transfers" not "Security for payments"
- Include the protected asset or threat in the title when it adds clarity

**Common Mistakes:**
- Generic IDs like SR-001 without category context
- Titles that name a technology instead of a control ("Implement OAuth" vs. "API Authentication using OAuth 2.0")

---

## 2. Identifying Threats & Risks

**Best Practices:**
- Every security requirement must trace to a specific threat, vulnerability, or risk
- Use established threat modeling frameworks: STRIDE, OWASP Top 10, MITRE ATT&CK
- Reference the organization's risk register where applicable
- Quantify risk where possible (likelihood × impact)
- Consider both external threats (attackers) and internal threats (insider, misconfiguration)

**STRIDE Mapping for Banking:**

| Threat | Description | Banking Examples |
|---|---|---|
| **S**poofing | Impersonating a user or system | Account takeover, phishing, session hijacking |
| **T**ampering | Modifying data or code | Transaction manipulation, log tampering, API parameter tampering |
| **R**epudiation | Denying an action occurred | Disputed transactions, unauthorized transfers without audit trail |
| **I**nformation Disclosure | Exposing data to unauthorized parties | Data breach, API data leakage, error message information leak |
| **D**enial of Service | Making system unavailable | DDoS on banking portal, resource exhaustion, batch job blocking |
| **E**levation of Privilege | Gaining unauthorized access | Privilege escalation, IDOR, broken access control |

**OWASP Top 10 Banking Relevance:**

| OWASP Risk | Banking Impact | Typical Control |
|---|---|---|
| Broken Access Control | Unauthorized account access, data theft | RBAC, resource-level authorization, IDOR prevention |
| Cryptographic Failures | PII/financial data exposure | TLS 1.3, AES-256, proper key management |
| Injection | Data theft, system compromise | Parameterized queries, input validation |
| Insecure Design | Fundamental security flaws | Threat modeling, secure design patterns |
| Security Misconfiguration | Exposed admin panels, default credentials | Hardening standards, automated config scanning |
| Vulnerable Components | Supply chain attacks | Dependency scanning, SCA, SBOM |
| Auth Failures | Account takeover, session hijacking | MFA, secure session management, rate limiting |
| Data Integrity Failures | Tampered transactions, malicious updates | Code signing, integrity checks, secure CI/CD |
| Logging Failures | Undetected breaches, no forensic evidence | Centralized logging, tamper-proof audit trails |
| SSRF | Internal system access, data exfiltration | Input validation, network segmentation, allowlists |

---

## 3. Writing the Description

**Best Practices:**
- State what the security control does, what it protects, and against what threat
- Be specific about the scope: which systems, APIs, data, or user interactions
- Include the security principle being applied (defense in depth, least privilege, zero trust)
- Avoid implementation-specific language unless the standard mandates a specific approach

**Formula:**
> The system shall [implement/enforce/prevent/detect/respond to] [specific control] for [specific scope] to mitigate [specific threat] in accordance with [standard/framework].

**Good Example:**
> The system shall enforce multi-factor authentication for all fund transfer operations to mitigate account takeover and unauthorized transaction risks, in accordance with PSD2 SCA (Article 97) and NIST 800-63B AAL2.

**Bad Example:**
> Add MFA to the payment page.

---

## 4. Citing Implementation Standards

**Best Practices:**
- Reference specific standards and control numbers, not just framework names
- Use the most current version of the standard
- When multiple standards apply, cite all of them
- Include the specific assurance level where applicable (e.g., NIST AAL2, OWASP ASVS Level 2)

**Key Standards for Banking Security:**

| Domain | Standard | Use For |
|---|---|---|
| Authentication | NIST 800-63B, PSD2 SCA | Auth strength levels, MFA requirements |
| Application Security | OWASP ASVS v4.0 | Comprehensive application security controls |
| Encryption | NIST 800-57, FIPS 140-2/3 | Key management, cryptographic module validation |
| Network Security | CIS Benchmarks, NIST 800-53 | Network segmentation, firewall rules |
| Access Control | ISO 27001 A.9, NIST 800-53 AC | RBAC, privilege management |
| Logging | ISO 27001 A.12.4, PCI-DSS Req 10 | Audit logging, log protection |
| Incident Response | NIST 800-61, ISO 27035 | Incident handling procedures |
| Cloud Security | CIS Cloud Benchmarks, CSA CCM | Cloud-specific controls |

---

## 5. Writing Technical Specifications

**Best Practices:**
- Be specific enough for developers to implement correctly
- Specify algorithms, key lengths, protocol versions, and configuration parameters
- Include both what to do and what not to do (e.g., "TLS 1.3 preferred, TLS 1.2 minimum; SSL and TLS 1.0/1.1 prohibited")
- Specify secure defaults — the system should be secure out of the box
- Include configuration parameters with recommended values

**Banking-Specific Technical Specs:**

**Authentication:**
- MFA: TOTP (RFC 6238) or FIDO2/WebAuthn; SMS OTP only as fallback where regulatory permitted
- Session tokens: cryptographically random, minimum 128-bit entropy
- Session timeout: 15 minutes inactivity (sensitive ops), 8 hours absolute
- Password: minimum 12 characters, checked against breach databases (NIST 800-63B)
- Account lockout: 5 failed attempts → 30-minute progressive lockout

**Encryption:**
- Transit: TLS 1.3 (preferred), TLS 1.2 (minimum); disable SSL, TLS 1.0, TLS 1.1
- At rest: AES-256-GCM for data, AES-256-XTS for disk encryption
- Key management: HSM (FIPS 140-2 Level 3) for production keys; automated rotation every 90 days
- Hashing: bcrypt (cost factor 12+) or Argon2id for passwords; SHA-256 minimum for integrity

**API Security:**
- OAuth 2.0 with PKCE for public clients; client credentials for service-to-service
- Access tokens: JWT with RS256/ES256, 15-minute expiry
- Refresh tokens: opaque, single-use, secure storage, 24-hour expiry
- Rate limiting: 100 requests/minute per user (adjustable per endpoint criticality)
- Input validation: whitelist approach, server-side enforcement, reject on failure

---

## 6. Defining Testing Requirements

**Best Practices:**
- Specify which testing methods apply to this requirement (SAST, DAST, SCA, pen test, manual review)
- Define specific test scenarios, not just "test security"
- Include both positive tests (control works) and negative tests (control cannot be bypassed)
- Specify testing frequency and who performs the test
- Define pass/fail criteria for each test

**Testing Matrix by Security Category:**

| Category | SAST | DAST | SCA | Pen Test | Manual Review |
|---|---|---|---|---|---|
| Authentication | ✓ Logic flaws | ✓ Bypass attempts | — | ✓ MFA bypass | ✓ Flow review |
| Authorization | ✓ Missing checks | ✓ IDOR, privilege escalation | — | ✓ Access control bypass | ✓ RBAC review |
| Encryption | ✓ Weak algorithms | ✓ SSL/TLS config | — | ✓ Crypto attacks | ✓ Key management |
| Input Validation | ✓ Injection patterns | ✓ Fuzzing, injection | — | ✓ Advanced injection | — |
| Dependencies | — | — | ✓ CVE scanning | — | ✓ License review |
| Logging | ✓ Missing log calls | — | — | ✓ Log tampering | ✓ Log completeness |
| Configuration | ✓ Hardcoded secrets | ✓ Misconfig scanning | — | ✓ Config exploitation | ✓ Hardening review |

**Test Scenario Examples:**
```
# Authentication Bypass
Test: Attempt to access fund transfer API without valid session token
Expected: 401 Unauthorized, no data leakage in error response

# MFA Replay
Test: Capture a valid OTP and attempt to reuse it
Expected: OTP rejected on second use, alert generated

# IDOR (Insecure Direct Object Reference)
Test: Authenticated as Customer A, attempt to access Customer B's account via API
Expected: 403 Forbidden, access attempt logged and alerted

# SQL Injection
Test: Submit malicious SQL in all input fields (account number, reference, amount)
Expected: Input rejected, parameterized query prevents execution, attempt logged

# Privilege Escalation
Test: Modify JWT claims to elevate role from "customer" to "admin"
Expected: Token signature validation fails, request rejected, alert generated
```

---

## 7. Defining Monitoring & Alerting Requirements

**Best Practices:**
- Specify what security events must be monitored in real-time
- Define alert thresholds and escalation paths
- Specify response time expectations for each alert severity
- Include both automated responses and human review requirements
- Define correlation rules for detecting complex attack patterns

**Banking Security Monitoring Requirements:**

| Event | Alert Threshold | Response Time | Escalation |
|---|---|---|---|
| Failed login attempts | 5 in 5 minutes per account | Immediate auto-lock | Security team if pattern across accounts |
| MFA failures | 3 consecutive per session | Immediate auto-lock | Fraud team |
| Privilege escalation attempt | Any occurrence | Immediate alert | Security team + CISO |
| Unusual transaction pattern | Risk score > threshold | < 5 minutes | Fraud team |
| Data exfiltration indicators | Bulk data access anomaly | < 15 minutes | Security team + DPO |
| API abuse | Rate limit exceeded | Immediate throttle | Security team if persistent |
| Configuration change | Any production change | < 30 minutes review | Change management |
| Vulnerability detected | Critical/High CVE | < 24 hours assessment | Security team + DevOps |

---

## 8. Defining Incident Response

**Best Practices:**
- Specify what happens when the security control fails or is breached
- Define containment actions (automated and manual)
- Define communication requirements (who is notified, within what timeframe)
- Include recovery steps and post-incident actions
- Reference the organization's incident response plan

**Incident Response Template per Security Requirement:**
> **Detection:** [How the breach/failure is detected]
> **Containment:** [Immediate automated actions + manual actions]
> **Notification:** [Who is notified, within what timeframe, via what channel]
> **Investigation:** [What evidence is preserved, who investigates]
> **Recovery:** [How normal operations are restored]
> **Post-Incident:** [RCA, lessons learned, control improvement]

---

## 9. Writing Acceptance Criteria

**Best Practices:**
- Security acceptance criteria must prove the control works AND cannot be bypassed
- Include both functional verification and attack simulation
- Specify the tools used for verification
- Include evidence requirements (scan reports, test results)

**Patterns:**
```
# Authentication
Given a user initiating a fund transfer,
When step-up MFA is triggered,
Then the user must provide a valid second factor
  AND expired OTPs (>90 seconds) are rejected
  AND replayed OTPs are rejected
  AND 5 consecutive failures lock the operation for 30 minutes
  AND all attempts are logged with timestamp, IP, and device fingerprint.

# Encryption
Given data transmitted between client and server,
When inspected via network analysis,
Then all traffic uses TLS 1.2 or higher
  AND no SSL/TLS 1.0/1.1 connections are accepted
  AND certificate validation is enforced
  AND HSTS header is present with max-age ≥ 31536000.

# Access Control
Given an authenticated user with "customer" role,
When they attempt to access another customer's account data via API,
Then the request returns 403 Forbidden
  AND no data from the other account is included in the response
  AND the access attempt is logged and triggers a security alert.
```

---

## 10. Exception Process

**Best Practices:**
- Define a clear process for when a security requirement cannot be met
- Require documented justification, risk assessment, and compensating controls
- Specify approval authority (typically CISO + Compliance Officer for banking)
- Set an expiry date on exceptions — they must be reviewed and renewed
- Track all exceptions in a central register

**Exception Template:**
> **Requirement:** [SR ID and title]
> **Justification:** [Why the requirement cannot be met]
> **Risk Assessment:** [Residual risk with the exception]
> **Compensating Controls:** [Alternative controls in place]
> **Approved By:** [CISO, Compliance Officer, Business Owner]
> **Expiry Date:** [Maximum 6 months, then re-review]
> **Review Date:** [Next review date]

**Banking-Specific Exception Rules:**
- No exceptions for: PCI-DSS encryption of card data, MFA for fund transfers, audit logging for financial transactions
- Time-limited exceptions only for: legacy system migration, third-party integration constraints
- All exceptions must have compensating controls documented and tested

---

## 11. Data Classification in Security Requirements (ref: PM-BP-008)

**Best Practices:**
- Security controls must be proportional to data classification — Restricted data needs the strongest controls
- Specify the exact classification levels the SR protects in the Data Classification field
- Use classification to determine: encryption strength, access control granularity, audit depth, monitoring sensitivity
- Security requirements for Restricted data are non-negotiable — no exceptions permitted

**Classification-Driven Security Controls:**

| Control | Restricted | Confidential | Internal | Public |
|---|---|---|---|---|
| Encryption at rest | AES-256 + HSM keys | AES-256 | Recommended | Not required |
| Encryption in transit | TLS 1.3 mandatory | TLS 1.2+ | TLS 1.2+ | TLS 1.2+ recommended |
| Authentication | MFA mandatory | MFA for modifications | Password-based | Anonymous allowed |
| Authorization | Resource-level + MFA step-up | Resource-level | Role-based | Open |
| Audit logging | All access logged | All modifications logged | Auth events | Minimal |
| Monitoring | Real-time alerting on any access | Real-time on anomalies | Daily review | None |
| Data masking (non-prod) | Full anonymization | Tokenization | Recommended | Not required |
| Incident response | Immediate (< 15 min) | Urgent (< 1 hour) | Standard (< 4 hours) | Best effort |

---

## 12. Audit Trail in Security Requirements (ref: PM-BP-009)

**Best Practices:**
- Every SR must specify what security events are logged as part of the Audit Trail Requirements field
- Separate monitoring (real-time detection) from audit trail (forensic evidence) — both are needed
- Audit logs for security events must be immutable, centrally stored, and separately access-controlled
- Specify log fields: event type, timestamp (UTC), user ID, source IP, action, resource, outcome, correlation ID
- Specify tamper-proofing method (cryptographic hashing, write-once storage, blockchain-based)
- Security audit logs must be available for forensic analysis within minutes, not hours

**Security Audit Events That Must Always Be Logged:**

| Event Category | Specific Events | Retention |
|---|---|---|
| Authentication | Login success/failure, MFA success/failure, lockout, unlock, password change/reset | 3 years |
| Authorization | Access granted/denied, privilege escalation attempt, role change | 3 years |
| Data access | Read/write to Restricted/Confidential data, bulk data export, API data retrieval | 7 years |
| Configuration | Security config change, firewall rule change, certificate rotation, key rotation | 7 years |
| Incident | Alert triggered, incident created, containment action, recovery action | 7 years |
| Admin | User creation/deletion, role assignment, system override, exception approval | 7 years |

---

## 13. DR/BCP in Security Requirements (ref: PM-BP-010)

**Best Practices:**
- Security controls must not be weakened during DR/BCP events — this is when systems are most vulnerable
- Specify how security infrastructure (HSMs, certificate stores, key vaults, IAM) is replicated to DR
- Define security monitoring continuity during failover — SIEM, alerting, and log collection must continue
- Specify that MFA, encryption, and access controls remain enforced in DR mode
- Define the security validation checklist for DR activation (verify controls are active before accepting traffic)
- Specify security testing requirements for DR environments (same SAST/DAST/pen test standards)

**Security DR/BCP Checklist:**

| Security Component | DR Requirement |
|---|---|
| HSM / Key Management | Keys replicated to DR HSM; key access verified on failover |
| Certificates / TLS | DR certificates pre-provisioned and valid; no certificate errors on failover |
| IAM / Access Control | IAM replicated to DR; RBAC enforced; no emergency bypass accounts |
| MFA Infrastructure | TOTP seeds / push notification service available in DR |
| SIEM / Monitoring | Log collection continues in DR; alerting rules active; no blind spots |
| Firewall / Network Security | DR network segmentation mirrors production; no relaxed rules |
| Secrets Management | Vault replicated to DR; secrets accessible; rotation continues |
| Audit Logging | Logs from DR environment captured with same completeness and integrity |

---

## 14. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| "Make it secure" | Unmeasurable, untestable | Specify control, threat, and standard |
| Security as last step | Expensive retrofitting, vulnerabilities | Security-by-design from inception |
| No threat reference | Control may not address actual risks | Map every SR to a threat/risk |
| Technology-only focus | Misses process and people controls | Include technical + process + training |
| No testing specified | Can't verify the control works | Define test method, scenario, and frequency |
| No monitoring | Breaches go undetected | Define real-time monitoring and alerting |
| No incident response | No plan when control fails | Define detection, containment, recovery |
| Blanket exceptions | Erodes security posture | Time-limited, compensating controls, CISO approval |
| Copy-paste from standards | Requirements don't fit context | Tailor to your system's specific threats and architecture |
| No exception process | Either everything blocked or everything bypassed | Formal exception with risk acceptance |
| Same security controls for all data | Over/under-protection | Tier controls by data classification |
| Monitoring without audit trail | Can detect but can't prove in court | Separate monitoring (real-time) from audit (forensic) |
| Security relaxed during DR | Most vulnerable when most exposed | Maintain all controls during failover |
| Security infrastructure not in DR | Controls unavailable on failover | Replicate HSM, IAM, SIEM, vault to DR |
