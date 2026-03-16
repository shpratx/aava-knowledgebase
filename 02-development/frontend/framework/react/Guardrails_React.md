# Guardrails for React Application Development
---

## 1. Client-Side Data Protection Guardrails (ref: DV-GR-001)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-DG-001 | No sensitive data (PII, account numbers, balances, tokens, passwords) in client-side code or localStorage/sessionStorage (ref: DV-GR-001) | Reject if sensitive data found in client storage |
| RG-DG-002 | No sensitive data in React component state that persists beyond the component visible lifecycle | Reject if sensitive data remains in state after unmount |
| RG-DG-003 | All sensitive data must be cleared from memory on logout, session timeout, and component unmount | Reject if cleanup is missing |
| RG-DG-004 | No sensitive data in URL parameters, query strings, or browser history | Reject if sensitive data appears in URLs |
| RG-DG-005 | No sensitive data in Redux/global store — server state via React Query only (not persisted) | Reject if sensitive data is in persisted global state |
| RG-DG-006 | Tokens must be stored in HTTP-only secure cookies or in-memory only — never localStorage | Reject if tokens are in localStorage/sessionStorage |
| RG-DG-007 | No sensitive data in service worker cache or IndexedDB | Reject if sensitive data found in browser storage APIs |
| RG-DG-008 | Account numbers must be masked in UI by default (****1234) — full display only on explicit user action | Reject if full account numbers displayed by default |
| RG-DG-009 | Clipboard operations on sensitive data must auto-clear after 60 seconds | Flag if no clipboard auto-clear |
| RG-DG-010 | Browser auto-fill must be disabled for OTP, CVV, and PIN fields | Reject if sensitive fields allow auto-fill |

---

## 2. Dependency Security Scanning Guardrails (ref: DV-GR-002)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-DS-001 | All external libraries must be security-scanned before use and on every build (ref: DV-GR-002) | Reject if any dependency is unscanned |
| RG-DS-002 | npm audit must run in CI/CD on every build — critical vulnerabilities block merge | CI/CD gate |
| RG-DS-003 | Snyk or Dependabot must be configured for continuous monitoring | Reject if no continuous monitoring |
| RG-DS-004 | New dependencies must pass security review checklist before addition | Reject if added without review |
| RG-DS-005 | Only approved licenses permitted: MIT, Apache 2.0, BSD, ISC — GPL/AGPL prohibited | Reject if prohibited license |
| RG-DS-006 | Transitive dependency tree must be reviewed for new packages | Flag if transitive deps not reviewed |
| RG-DS-007 | package-lock.json must be committed; npm ci used in CI (not npm install) | Reject if lock file missing |
| RG-DS-008 | No dependencies from untrusted registries — only npmjs.com or approved private registry | Reject if untrusted registry |

---

## 3. Dependency Patching Guardrails (ref: DV-GR-003)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-DP-001 | Dependencies must be updated within 30 days of security patch release (ref: DV-GR-003) | Flag if patch > 30 days old |
| RG-DP-002 | Critical CVE patches must be applied within 24 hours | Reject deployment if unpatched > 24h |
| RG-DP-003 | High CVE patches must be applied within 7 days | Reject deployment if unpatched > 7d |
| RG-DP-004 | Medium CVE patches must be applied within 30 days | Flag if unpatched > 30d |
| RG-DP-005 | Dependency update PRs must be reviewed and merged within SLA per severity | Process guardrail |
| RG-DP-006 | Major version upgrades of core dependencies require architecture review | Reject if no arch review |
| RG-DP-007 | Vulnerable dependencies with no patch must have documented mitigation or be replaced | Reject if no mitigation plan |

---

## 4. Production Logging Guardrails (ref: DV-GR-004)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-LG-001 | No console.log/warn/error/debug/info in production code (ref: DV-GR-004) | CI/CD gate — ESLint no-console; build strips console |
| RG-LG-002 | Production builds must strip all console statements via build config | Reject if production build contains console |
| RG-LG-003 | Structured logging must use dedicated monitoring service (Datadog, Sentry) | Reject if error reporting uses console |
| RG-LG-004 | Error logging to monitoring must not include sensitive data | Reject if monitoring payloads contain sensitive data |
| RG-LG-005 | Debug/dev-only code must be excluded from production builds | Reject if dev code in production bundle |

---

## 5. Source Map Guardrails (ref: DV-GR-005)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-SM-001 | Source maps must not be deployed to production or be publicly accessible (ref: DV-GR-005) | Reject if source maps accessible in production |
| RG-SM-002 | Production build must set GENERATE_SOURCEMAP=false or sourcemap: false | CI/CD gate |
| RG-SM-003 | If needed for monitoring, upload privately to Sentry — do not serve publicly | Flag if served publicly |
| RG-SM-004 | CI/CD must verify no .map files in deployment artifact | CI/CD gate — scan for .map files |
| RG-SM-005 | CDN/server must block .map file requests as defense in depth | Flag if not blocked |

---

## 6. Third-Party Script Guardrails (ref: DV-GR-006)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-TP-001 | Third-party scripts require security review before integration (ref: DV-GR-006) | Reject if added without review |
| RG-TP-002 | Third-party scripts must load from approved CDN domains only — listed in CSP | Reject if unapproved domain |
| RG-TP-003 | Third-party scripts must not access sensitive banking data | Reject if script can access banking data |
| RG-TP-004 | Third-party scripts must not block page rendering — use async/defer | Flag if scripts block rendering |
| RG-TP-005 | Third-party scripts must comply with privacy policy and GDPR | Reject if violates privacy policy |
| RG-TP-006 | Third-party script updates require re-review | Process guardrail |
| RG-TP-007 | Minimize third-party scripts — each increases attack surface | Flag if > 5 third-party scripts |

---

## 7. XSS Prevention Guardrails (ref: DV-GR-007)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-XG-001 | dangerouslySetInnerHTML must only be used with DOMPurify sanitization (ref: DV-GR-007) | Reject if used without DOMPurify |
| RG-XG-002 | Every dangerouslySetInnerHTML use must be approved by security review | Reject if no security approval |
| RG-XG-003 | eval(), Function(), setTimeout(string), setInterval(string) are prohibited | Reject immediately — no exceptions |
| RG-XG-004 | innerHTML and outerHTML DOM manipulation are prohibited | Reject if direct DOM HTML manipulation |
| RG-XG-005 | href/src attributes must validate URLs — reject javascript: protocol | Reject if URL validation missing |
| RG-XG-006 | Markdown rendering must use sanitizing renderer (react-markdown + rehype-sanitize) | Reject if no sanitization |
| RG-XG-007 | Template literals must not construct HTML strings for DOM insertion | Reject if template literals create HTML |
| RG-XG-008 | All user input must go through React JSX rendering or explicit sanitization | Reject if input bypasses escaping |
| RG-XG-009 | CSP must block inline scripts (no unsafe-inline in script-src) | Reject if CSP allows unsafe-inline |
| RG-XG-010 | ESLint react/no-danger rule must be enabled and set to error | CI/CD gate |

---

## 8. Subresource Integrity Guardrails (ref: DV-GR-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-SI-001 | All CDN resources must include SRI hashes (ref: DV-GR-008) | Reject if CDN resources lack SRI |
| RG-SI-002 | SRI hashes must use SHA-384 or SHA-512 | Reject if weaker hash |
| RG-SI-003 | SRI hashes must be regenerated on resource version change | Process guardrail |
| RG-SI-004 | crossorigin="anonymous" must be set on elements with SRI | Reject if missing |
| RG-SI-005 | SRI verification failures must be monitored | Flag if no monitoring |

---

## 9. Component and Code Quality Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-CQ-001 | All components must be functional — class components prohibited (except ErrorBoundary) | Reject if class component introduced |
| RG-CQ-002 | All components must be typed with TypeScript — no any type | CI/CD gate — tsc strict |
| RG-CQ-003 | Error boundaries must wrap every feature/route | Reject if feature lacks error boundary |
| RG-CQ-004 | Error boundary fallback must not display technical details | Reject if fallback exposes internals |
| RG-CQ-005 | Max file: 300 lines; max function: 50 lines; max complexity: 10 | CI/CD gate — ESLint |
| RG-CQ-006 | Test coverage >= 80% overall; >= 90% for shared utils/hooks | CI/CD gate |
| RG-CQ-007 | Accessibility scan must pass with zero critical violations | CI/CD gate |
| RG-CQ-008 | All forms must have server-side validation — client-side is UX only | Reject if client-side only |

---

## 10. State and Configuration Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-SC-001 | No hardcoded API URLs or config values in source code | Reject if hardcoded config found |
| RG-SC-002 | No secrets in client-side code or environment variables | Reject immediately |
| RG-SC-003 | .env files with real values must be in .gitignore | Reject if .env committed |
| RG-SC-004 | Server state must use React Query — not Redux | Flag if API data in Redux |
| RG-SC-005 | Global state must be minimal: theme, language, feature flags only | Flag if business data in global state |
| RG-SC-006 | All state must be cleared on logout and session timeout | Reject if state persists after logout |

---

## 11. Performance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| RG-PG-001 | Every feature/route must be lazy-loaded via React.lazy() | Reject if eagerly loaded |
| RG-PG-002 | Initial bundle must not exceed 200KB gzipped | CI/CD gate |
| RG-PG-003 | Per-feature chunk must not exceed 100KB gzipped | CI/CD gate |
| RG-PG-004 | Bundle size increase > 10% on any PR must be justified | CI/CD gate — require approval |
| RG-PG-005 | Lighthouse performance score must be >= 90 | CI/CD gate |
| RG-PG-006 | Lists > 50 items must use virtualization | Flag if not virtualized |
| RG-PG-007 | Images must use lazy loading, WebP, responsive srcset | Flag if not optimized |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot merge/deploy | Sensitive data in localStorage (DV-GR-001), unscanned dependency (DV-GR-002), source maps in prod (DV-GR-005), dangerouslySetInnerHTML without DOMPurify (DV-GR-007), eval() |
| **Flag** | Address before release | Patch > 30 days (DV-GR-003), > 5 third-party scripts (DV-GR-006), large list not virtualized |
| **CI/CD Gate** | Automated | npm audit (DV-GR-002), console detection (DV-GR-004), .map scan (DV-GR-005), ESLint no-danger (DV-GR-007), bundle size, coverage, accessibility |
| **Process** | Workflow | Dependency review, third-party re-review, SRI regeneration |

---

## Quick Reference: Guardrail Triggers

| Activity | Triggered Guardrails |
|---|---|
| Adding npm package | RG-DS-001-008, RG-DP-006 |
| Rendering user input | RG-XG-001-010, RG-DG-008 |
| Storing data client-side | RG-DG-001-010 |
| Loading CDN resources | RG-SI-001-005, RG-TP-001-007 |
| Production build | RG-LG-001-005, RG-SM-001-005, RG-PG-001-007 |
| Every PR/merge | RG-DS-002, RG-CQ-002, RG-CQ-005-007 |

---

## Pre-Deployment Checklist

| # | Check | Ref |
|---|---|---|
| 1 | No sensitive data in localStorage, sessionStorage, or client state | DV-GR-001 |
| 2 | All dependencies security-scanned; zero critical/high CVEs | DV-GR-002 |
| 3 | All security patches applied within SLA | DV-GR-003 |
| 4 | No console.log in production build | DV-GR-004 |
| 5 | No source maps in deployment artifact | DV-GR-005 |
| 6 | All third-party scripts security-reviewed | DV-GR-006 |
| 7 | No unsanitized dangerouslySetInnerHTML; no eval() | DV-GR-007 |
| 8 | SRI hashes on all CDN resources | DV-GR-008 |
| 9 | Error boundaries on all features | DV-BP-002 |
| 10 | CSP configured with no unsafe-eval | DV-BP-005 |
| 11 | Test coverage >= 80%; accessibility scan passed | DV-BP-001 |
| 12 | Bundle size within limits; features lazy-loaded | DV-BP-008 |
| 13 | No secrets in client-side code or env vars | DV-BP-004 |
| 14 | State cleared on logout/timeout | DV-GR-001 |
