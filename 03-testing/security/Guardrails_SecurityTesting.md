# Security Testing — Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Guardrails

 Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-SEC-001 | SAST must run on every build — critical/high findings block merge | CI/CD gate |
| TG-SEC-002 | SCA (dependency scan) must run on every build — critical CVEs block merge | CI/CD gate |
| TG-SEC-003 | DAST must run on staging before every production release | Reject release without DAST |
| TG-SEC-004 | Every endpoint must have authentication test (no token → 401) | Reject if auth untested |
| TG-SEC-005 | Every resource endpoint must have IDOR test (other user's ID → 403) | Reject if IDOR untested |
| TG-SEC-006 | Every string input must have injection test (SQL, XSS, command) | Reject if injection untested |
| TG-SEC-007 | Session timeout must be tested (15 min inactivity; 8 hour absolute) | Reject if timeout untested |
| TG-SEC-008 | MFA enforcement must be tested for all financial operations | Reject if MFA untested |
| TG-SEC-009 | Sensitive data masking must be verified in responses, logs, URLs, and storage | Reject if masking unverified |
| TG-SEC-010 | Security headers must be verified on all responses (HSTS, CSP, X-Content-Type-Options) | Reject if headers unverified |
| TG-SEC-011 | TLS configuration must be verified (TLS 1.2+; no weak ciphers) | Reject if TLS unverified |
| TG-SEC-012 | Rate limiting must be tested for all endpoints | Reject if rate limiting untested |
| TG-SEC-013 | Penetration test must be conducted annually by qualified external testers | Reject release if pen test overdue > 12 months |
| TG-SEC-014 | All OWASP Top 10 categories must be tested per release | Reject if OWASP coverage incomplete |
| TG-SEC-015 | Security scan reports must be retained for 7 years | Reject if reports not retained |
| TG-SEC-016 | Critical security findings must be remediated before production deployment — no exceptions | Reject deployment with open critical findings |

### Pre-Release Security Checklist
| # | Check |
|---|---|
| 1 | SAST passed — zero critical/high |
| 2 | SCA passed — zero critical CVEs |
| 3 | DAST passed — zero critical/high |
| 4 | Auth tested: no token, expired, tampered, wrong algorithm → 401 |
| 5 | IDOR tested: other user's resources → 403 |
| 6 | Injection tested: SQL, XSS, command in all string inputs |
| 7 | Session tested: timeout, invalidation, cookie flags |
| 8 | MFA tested: financial ops without MFA → blocked |
| 9 | Data masking verified: responses, logs, URLs, storage |
| 10 | Security headers verified: HSTS, CSP, X-Content-Type-Options |
| 11 | TLS verified: 1.2+; no weak ciphers |
| 12 | Rate limiting verified: 429 + Retry-After |
| 13 | PCI-DSS: no PAN/CVV in responses or logs |
| 14 | Pen test current (within 12 months) |