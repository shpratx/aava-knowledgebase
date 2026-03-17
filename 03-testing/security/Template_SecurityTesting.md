# Security Testing — Template, Best Practices & Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Part A: Test Scenario Template

### Metadata
| Field | Description |
|---|---|
| **Scenario ID** | TS-SEC-[Category]-[Seq] (e.g., TS-SEC-AUTH-001) |
| **Category** | Authentication / Authorization / Injection / Session / Encryption / Data Exposure / API Security / Compliance |
| **OWASP Reference** | OWASP Top 10 category (e.g., A01:2021 Broken Access Control) |
| **Risk Level** | Critical / High / Medium / Low |
| **Target** | UI / API / Database / Infrastructure / Mobile |
| **Tool** | OWASP ZAP / Burp Suite / Custom / Manual |

### Security Test Categories

#### 1. Authentication Testing (OWASP A07:2021)
| Scenario | Test | Expected |
|---|---|---|
| No credentials | Access protected resource without token | 401; redirect to login (UI) |
| Invalid credentials | Login with wrong password | "Invalid credentials" (no differentiation) |
| Expired token | Use expired JWT | 401; no data returned |
| Tampered token | Modify JWT payload (change userId, role) | 401; signature validation fails |
| Algorithm confusion | JWT with alg:"none" or alg:"HS256" (when RS256 expected) | 401; rejected |
| Brute force | 10+ failed login attempts | Account locked after 5; progressive delay |
| Credential stuffing | Automated login with breached credentials | Rate limited; CAPTCHA triggered |
| MFA bypass | Access MFA-protected endpoint without MFA token | 403; operation blocked |
| MFA replay | Reuse previously valid OTP | Rejected (single-use) |
| MFA expiry | Use OTP after 90 seconds | Rejected (expired) |
| Session fixation | Set session ID before login; verify new ID after login | New session ID issued on authentication |
| Password reset | Request reset for non-existent email | Same response as valid email (no enumeration) |

#### 2. Authorization Testing (OWASP A01:2021)
| Scenario | Test | Expected |
|---|---|---|
| Horizontal privilege (IDOR) | User A accesses User B's account via ID manipulation | 403; no data from User B |
| Vertical privilege | Regular user accesses admin endpoint | 403 |
| Scope enforcement | Call API with wrong OAuth scope | 403 |
| Missing authorization | Endpoint without auth check (misconfiguration) | Should return 401/403; not 200 |
| Forced browsing | Access /admin, /debug, /actuator/env directly | 403 or 404 |
| Parameter tampering | Modify accountId in request body to another user's | 403; ownership check fails |
| Role escalation | Modify role claim in token | 401 (signature invalid) |
| Maker-checker bypass | Same user initiates and approves | Rejected; different user required |

#### 3. Injection Testing (OWASP A03:2021)
| Scenario | Test | Expected |
|---|---|---|
| SQL injection | `' OR 1=1 --` in all string fields | 400 (validation) or sanitized; no SQL execution |
| SQL injection (blind) | `' AND SLEEP(5) --` | No delay; no SQL execution |
| NoSQL injection | `{"$gt": ""}` in JSON fields | 400; no query manipulation |
| XSS (reflected) | `<script>alert(1)</script>` in input fields | Sanitized; no script execution |
| XSS (stored) | Script in beneficiary name/reference | Sanitized on storage and display |
| XSS (DOM) | Script in URL fragment/params | CSP blocks; no execution |
| Command injection | `; ls -la` in any field | 400; no command execution |
| Path traversal | `../../etc/passwd` in file paths | 400; no file access |
| LDAP injection | `*)(uid=*))(|(uid=*` in search | 400; no LDAP query manipulation |
| XML/XXE | External entity in XML payload | Disabled; no external entity processing |
| Header injection | Newline characters in header values | Sanitized; no header injection |
| SSRF | Internal URL in user-supplied URL field | Blocked; no internal network access |

#### 4. Session Management Testing
| Scenario | Test | Expected |
|---|---|---|
| Session timeout | Idle for 15 minutes | Session expired; re-auth required |
| Absolute timeout | Active for 8 hours | Session expired regardless of activity |
| Session invalidation | Logout | Server-side session destroyed; token revoked |
| Concurrent sessions | Login from two devices | Alert on second device; option to terminate other |
| Session token in URL | Check if session/token appears in URL | Never in URL |
| Cookie flags | Inspect session cookie | HttpOnly, Secure, SameSite=Strict |
| Token rotation | Refresh token used | Old refresh token invalidated; new one issued |

#### 5. Data Protection Testing
| Scenario | Test | Expected |
|---|---|---|
| PII in response | Check API responses for unmasked PII | Account numbers masked (****1234); email/phone partially masked |
| PII in logs | Check application logs | No passwords, tokens, PAN, CVV, full account numbers |
| PII in URL | Check URLs and query parameters | No sensitive data in URLs |
| PII in storage | Check localStorage, sessionStorage, cookies | No sensitive data in client storage |
| PII in errors | Trigger errors; check messages | No PII, no SQL, no stack traces |
| Encryption at rest | Inspect database storage | PII encrypted (AES-256) |
| Encryption in transit | Inspect network traffic | TLS 1.2+; no plaintext |
| Source maps | Check production deployment | No .map files accessible |
| Debug endpoints | Check /debug, /actuator/env, /swagger-ui | Disabled or secured in production |
| Server headers | Check response headers | No Server, X-Powered-By, X-AspNet-Version |

#### 6. API-Specific Security Testing
| Scenario | Test | Expected |
|---|---|---|
| Rate limiting | Exceed rate limit | 429 + Retry-After |
| CORS | Request from unauthorized origin | Blocked; no wildcard |
| CSRF | State-changing request without CSRF token | Rejected |
| Content-Type validation | Send wrong Content-Type | 415 Unsupported Media Type |
| Request size | Send oversized payload (> 1MB) | 413 Payload Too Large |
| HTTP methods | Send unsupported method (e.g., TRACE) | 405 Method Not Allowed |
| TLS version | Connect with TLS 1.0/1.1 | Connection refused |
| Security headers | Check all responses | HSTS, CSP, X-Content-Type-Options, X-Frame-Options present |
| Idempotency key collision | Same key, different body | 422 (not silently accepted) |

#### 7. Compliance Security Testing
| Scenario | Test | Expected |
|---|---|---|
| PSD2 SCA | Payment without two-factor auth | Blocked; MFA required |
| PCI-DSS PAN | Full PAN in any response/log | Never present; masked or tokenized |
| PCI-DSS CVV | CVV stored post-authorization | Never stored |
| GDPR consent | Process data without consent (optional processing) | Blocked until consent granted |
| GDPR erasure | Request erasure; verify data removed | PII anonymized within 30 days |
| AML screening | Transaction without sanctions check | Blocked; screening mandatory |

---