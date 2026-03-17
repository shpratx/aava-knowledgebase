# Security Testing — Best Practices
### Banking Domain — Agentic Knowledge Base

---

## Best Practices

 Best Practices

| Practice | Standard |
|---|---|
| SAST on every build | SonarQube/Checkmarx/Semgrep in CI; block on critical/high |
| DAST on staging | OWASP ZAP/Burp Suite before every release |
| SCA on every build | Snyk/OWASP Dependency-Check; block on critical CVEs |
| Pen test annually | External qualified testers; scope includes all critical functions |
| TLPT every 3 years | Threat-Led Penetration Testing per DORA (for significant entities) |
| Test OWASP Top 10 | Every release must verify no OWASP Top 10 vulnerabilities |
| Shift left | Security tests in developer workflow, not just pre-release |
| Threat model first | Threat model (STRIDE) before writing security tests |
| Test all input vectors | Every user-supplied field, header, URL parameter, file upload |
| Automate regression | Security tests in CI; manual for new attack vectors |
| Separate security tests | Don't mix with functional tests; dedicated security test suite |
| Evidence retention | Security scan reports retained for 7 years (audit) |

### Anti-Patterns
| Anti-Pattern | Fix |
|---|---|
| Security testing only before release | SAST/SCA on every build; DAST on staging continuously |
| Only testing happy path security | Test every failure mode: expired, tampered, missing, wrong scope |
| No IDOR testing | Test resource access with every other user role |
| Trusting client-side validation | Test API directly, bypassing UI |
| No injection testing | Automated injection payloads in every string field |
| Pen test as only security testing | Pen test supplements, not replaces, automated security testing |

---