# Guardrails for UI Test Scenario & Test Case Generation
### Banking Domain — Agentic Knowledge Base

---

## 1. Coverage Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-UI-CG-001 | Every user story acceptance criterion must have at least one test case | Reject if any AC lacks a test |
| TG-UI-CG-002 | Every security acceptance criterion must have a dedicated security test case | Reject if security AC lacks test |
| TG-UI-CG-003 | Every compliance acceptance criterion must have a compliance test case | Reject if compliance AC lacks test |
| TG-UI-CG-004 | Every error/exception flow in the user story must have a test case | Reject if error flow is untested |
| TG-UI-CG-005 | Both positive (happy path) and negative (error/rejection) scenarios must be covered | Reject if only happy path tested |
| TG-UI-CG-006 | Boundary value tests must exist for all numeric inputs (min, max, limit, limit+1, limit-1) | Reject if no boundary tests for numeric fields |
| TG-UI-CG-007 | Empty state, loading state, and error state must be tested for every data-driven component | Flag if UI states are not tested |

---

## 2. Security Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-UI-SG-001 | Authentication test: every protected page must be tested without login → verify redirect | Reject if auth redirect not tested |
| TG-UI-SG-002 | IDOR test: URL manipulation with other user's resource IDs must be tested → verify 403/redirect | Reject if IDOR not tested for resource pages |
| TG-UI-SG-003 | Session timeout test: verify timeout at 15 min inactivity; warning at 13 min; data cleared | Reject if session timeout not tested |
| TG-UI-SG-004 | XSS test: script injection in every text input field must be tested | Reject if XSS not tested on input fields |
| TG-UI-SG-005 | Sensitive data exposure test: verify no PII in localStorage, sessionStorage, URL, DOM source | Reject if data exposure not tested |
| TG-UI-SG-006 | MFA test: sensitive operations must be tested with and without MFA | Reject if MFA enforcement not tested |
| TG-UI-SG-007 | Error message test: verify no technical details in any error state | Reject if error leakage not tested |

---

## 3. Accessibility Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-UI-AG-001 | Automated accessibility scan (axe-core) must run on every build — zero critical violations | CI/CD gate |
| TG-UI-AG-002 | Keyboard-only navigation must be tested for all critical flows (login, transfer, payment) | Reject if keyboard test missing for critical flows |
| TG-UI-AG-003 | Screen reader testing must be performed for all critical flows before release | Reject if screen reader test missing |
| TG-UI-AG-004 | Color contrast must be verified (≥ 4.5:1 text, ≥ 3:1 large text/UI) | CI/CD gate (axe-core) |
| TG-UI-AG-005 | Form error messages must be tested for screen reader announcement (aria-live, aria-describedby) | Reject if error announcement not tested |
| TG-UI-AG-006 | All images/icons must be tested for alt text or aria-label | CI/CD gate (axe-core) |

---

## 4. Responsive Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-UI-RG-001 | All critical flows must be tested at mobile (375px), tablet (768px), and desktop (1024px) | Reject if not tested at all breakpoints |
| TG-UI-RG-002 | Touch targets must be verified ≥ 44×44px on mobile | Reject if touch targets below minimum |
| TG-UI-RG-003 | No horizontal scroll at any breakpoint | Reject if horizontal overflow exists |
| TG-UI-RG-004 | Cross-browser testing on Chrome, Firefox, Safari, Edge (latest 2 versions) | Reject if not tested on all required browsers |

---

## 5. Test Quality Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-UI-QG-001 | Test cases must use data-testid selectors — not CSS classes or XPath | Reject if fragile selectors used |
| TG-UI-QG-002 | No hardcoded waits (sleep/setTimeout) — use explicit waits with assertions | Reject if hardcoded waits found |
| TG-UI-QG-003 | Each test must be independent — no shared state or execution order dependency | Reject if tests depend on each other |
| TG-UI-QG-004 | Test data must be synthetic — no production data in test fixtures | Reject if production data used |
| TG-UI-QG-005 | Flaky tests must be quarantined within 24 hours and fixed within 1 week | Process guardrail |
| TG-UI-QG-006 | Visual regression tests must exist for all design system components | Flag if no visual regression |
| TG-UI-QG-007 | E2E test suite must complete within 15 minutes | Flag if suite exceeds 15 minutes |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot release | Missing AC test, no IDOR test, no keyboard test, no session timeout test, XSS untested |
| **Flag** | Address before next release | No visual regression, UI states untested, suite > 15 min |
| **CI/CD Gate** | Automated | axe-core scan, test coverage threshold, flaky test detection |

---

## Pre-Release UI Test Checklist

| # | Check |
|---|---|
| 1 | All acceptance criteria have test cases (functional, security, compliance) |
| 2 | Happy path + all error paths tested |
| 3 | Boundary values tested for all numeric inputs |
| 4 | Auth redirect tested for all protected pages |
| 5 | IDOR tested via URL manipulation |
| 6 | Session timeout tested (warning + expiry + data cleared) |
| 7 | XSS tested in all text input fields |
| 8 | No sensitive data in localStorage/sessionStorage/URL/DOM |
| 9 | MFA enforcement tested for sensitive operations |
| 10 | Keyboard-only flow tested for critical journeys |
| 11 | Screen reader tested for critical journeys |
| 12 | axe-core scan passed (zero critical) |
| 13 | Tested at mobile, tablet, desktop breakpoints |
| 14 | Cross-browser tested (Chrome, Firefox, Safari, Edge) |
| 15 | All tests independent; no flaky tests; synthetic data only |
