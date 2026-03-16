# PCI-DSS Requirement 6 — Develop and Maintain Secure Systems and Software

## Regulation Reference
- Standard: PCI-DSS v4.0
- Requirement: 6

## Obligation (Plain Language)
All system components must be protected from known vulnerabilities by installing applicable security patches. Custom software must be developed securely.

## Technical Controls Required
1. **Patch management:** Critical patches within 30 days of release; high within 90 days
2. **Secure development:** OWASP Top 10 addressed; secure coding guidelines followed
3. **Code review:** All custom code reviewed for vulnerabilities before production (manual or automated)
4. **SAST/DAST:** Static and dynamic analysis on every build
5. **Dependency scanning:** Third-party libraries scanned for known CVEs
6. **Change control:** All changes to system components follow change management process
7. **Web application protection:** WAF or equivalent for public-facing web applications
8. **Training:** Developers trained in secure coding annually

## Evidence Required
- Patch management records showing compliance with timelines
- Secure coding guidelines document
- Code review records (manual review or SAST/DAST reports)
- Dependency scan reports
- Change management records
- WAF configuration and logs
- Developer training records

## Acceptance Criteria
Given a system component in the CDE,
When assessed for Req 6 compliance,
Then all critical patches are applied within 30 days
  AND custom code is reviewed for OWASP Top 10 vulnerabilities
  AND SAST/DAST runs on every build with zero critical/high findings
  AND third-party dependencies are scanned for CVEs
  AND public-facing web applications are protected by WAF.
