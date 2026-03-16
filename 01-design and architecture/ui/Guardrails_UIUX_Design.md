# Guardrails for UI/UX Design and Architecture

---

## 1. Data Exposure Prevention Guardrails (ref: DA-GR-001)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-DEG-001 | No sensitive data (PII, account numbers, balances, transaction details) in URLs or query parameters (ref: DA-GR-001) | Reject if any sensitive data appears in URLs |
| UI-DEG-002 | No sensitive data in browser history — use replaceState for sensitive page transitions; POST for sensitive operations | Reject if sensitive data is recorded in browser history |
| UI-DEG-003 | No sensitive data in Referrer headers — set Referrer-Policy: strict-origin-when-cross-origin | Reject if referrer policy allows sensitive data leakage |
| UI-DEG-004 | No sensitive data in client-side storage (localStorage, sessionStorage) — use in-memory state only | Reject if sensitive data is stored in localStorage/sessionStorage |
| UI-DEG-005 | No sensitive data in cookies unless HTTP-only, Secure, and SameSite=Strict | Reject if sensitive cookies lack any of these flags |
| UI-DEG-006 | No sensitive data in browser console logs — remove all console.log statements containing data in production builds | Reject if production build contains console.log with data |
| UI-DEG-007 | No sensitive data in HTML source/DOM that is not currently displayed — clear data from DOM when no longer visible | Flag if hidden DOM elements contain sensitive data |
| UI-DEG-008 | No sensitive data in auto-complete suggestions — use autocomplete="off" for OTP, CVV; autocomplete="new-password" for password fields | Reject if sensitive fields allow auto-complete |
| UI-DEG-009 | Account numbers must be masked in UI display (show last 4 digits only) unless user explicitly requests full view | Reject if full account numbers are displayed by default |
| UI-DEG-010 | Sensitive data must be masked in mobile app switcher/recent apps view — implement background blur on app backgrounding | Reject if app content is visible in app switcher |
| UI-DEG-011 | Screenshot/screen recording must be prevented on screens displaying sensitive data (OTP, card details, full account numbers) | Reject for mobile apps if FLAG_SECURE (Android) / screen capture prevention (iOS) is not implemented on sensitive screens |
| UI-DEG-012 | PDF/document downloads containing sensitive data must use Content-Disposition: attachment and Cache-Control: no-store | Reject if sensitive downloads allow inline display or caching |

---

## 2. Multi-Factor Authentication Guardrails (ref: DA-GR-002)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-MFA-001 | Multi-factor authentication is required for all high-risk operations (ref: DA-GR-002) | Reject if high-risk operation has no MFA |
| UI-MFA-002 | Fund transfers must require step-up MFA before submission — regardless of session auth level | Reject if transfer flow has no MFA step |
| UI-MFA-003 | Beneficiary add/modify/delete must require MFA | Reject if beneficiary management has no MFA |
| UI-MFA-004 | Profile changes (email, phone, address) must require re-authentication or MFA | Reject if profile modification has no re-authentication |
| UI-MFA-005 | Password/PIN change must require current password/PIN verification plus MFA | Reject if credential change has no current credential verification |
| UI-MFA-006 | Large transfers (above configurable threshold) must require additional verification beyond standard MFA | Flag if no enhanced verification for large transfers |
| UI-MFA-007 | MFA UI must be accessible — OTP input fields must have labels, purpose, error states announced by screen readers | Reject if MFA UI fails WCAG 2.1 AA |
| UI-MFA-008 | MFA must offer accessible alternatives — if primary method is biometric, provide PIN/password fallback | Reject if MFA has no accessible fallback |
| UI-MFA-009 | MFA failure UI must not reveal which factor failed — display generic "Authentication failed" message | Reject if MFA error differentiates between factors |
| UI-MFA-010 | MFA timeout must be displayed to user (countdown timer) with clear indication when OTP expires | Flag if no expiry indication is shown |
| UI-MFA-011 | MFA resend/retry must be rate-limited in UI — disable resend button for 30 seconds after each request | Reject if MFA resend has no rate limiting |

---

## 3. Session Timeout Guardrails (ref: DA-GR-003)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-STG-001 | Session inactivity timeout must not exceed 15 minutes for sensitive operations (ref: DA-GR-003) | Reject if inactivity timeout exceeds 15 minutes for sensitive operations |
| UI-STG-002 | Absolute session timeout must not exceed 8 hours — force re-authentication regardless of activity | Reject if absolute timeout exceeds 8 hours |
| UI-STG-003 | Timeout warning must be displayed 2 minutes before session expiry with option to extend | Reject if no timeout warning is implemented |
| UI-STG-004 | Timeout warning must be accessible — announced by screen readers, keyboard-operable extend button | Reject if timeout warning is not accessible |
| UI-STG-005 | On session timeout, all sensitive data must be cleared from UI state, DOM, and memory | Reject if sensitive data persists after timeout |
| UI-STG-006 | On session timeout, user must be redirected to login page with clear message ("Session expired. Please log in again.") | Reject if timeout results in error page or blank screen |
| UI-STG-007 | Session timeout must be enforced server-side — client-side timeout is for UX only, not security | Reject if timeout is client-side only |
| UI-STG-008 | In-progress form data must be handled gracefully on timeout — warn user before clearing, or save non-sensitive draft | Flag if timeout causes silent data loss on forms |
| UI-STG-009 | Session extension must require user action (click "Continue") — no automatic silent extension | Reject if session auto-extends without user interaction |
| UI-STG-010 | Concurrent session detection must alert user — "You are logged in from another device. Continue here?" | Flag if no concurrent session detection exists |
| UI-STG-011 | Logout must invalidate session server-side immediately and clear all client-side state | Reject if logout does not invalidate server-side session |

---

## 4. Error Message Guardrails (ref: DA-GR-004)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-EMG-001 | Error messages must not leak sensitive information — no stack traces, internal paths, database details, SQL errors, or server versions (ref: DA-GR-004) | Reject if error messages expose technical details |
| UI-EMG-002 | Authentication error messages must not differentiate between invalid username and invalid password — use "Invalid credentials" | Reject if login error reveals which credential is wrong |
| UI-EMG-003 | Account enumeration must be prevented — "If an account exists, a reset link has been sent" not "No account found for this email" | Reject if password reset/registration reveals account existence |
| UI-EMG-004 | API error responses rendered in UI must be sanitized — never display raw API error bodies to users | Reject if raw API errors are displayed |
| UI-EMG-005 | Error messages must be user-friendly and actionable — tell the user what happened and what to do next | Flag if error messages are generic without guidance |
| UI-EMG-006 | Error messages must be accessible — linked to the relevant field via aria-describedby; announced by screen readers; not color-only | Reject if error messages are not accessible |
| UI-EMG-007 | HTTP error pages (404, 500, 503) must be custom-branded — no default server error pages exposing server technology | Reject if default server error pages are served |
| UI-EMG-008 | Error messages must not include correlation IDs or request IDs visible to the user — log them server-side for support reference | Flag if internal IDs are displayed to users (acceptable if shown as "Reference: XXX" for support purposes) |
| UI-EMG-009 | Rate limiting error messages must not reveal the exact limit or reset time — "Too many attempts. Please try again later." | Flag if rate limit details are exposed |
| UI-EMG-010 | Validation error messages must not reveal business rules that could be exploited — "Invalid amount" not "Amount must be between $0.01 and $999,999.99" for security-sensitive fields | Flag for security-sensitive fields; acceptable for user-friendly fields |

---

## 5. Internal System Detail Guardrails (ref: DA-GR-005)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-ISD-001 | UI must not expose internal system details — no server names, IP addresses, technology stack, framework versions, or internal API paths (ref: DA-GR-005) | Reject if internal system details are exposed |
| UI-ISD-002 | HTML source must not contain developer comments revealing architecture, TODOs, or internal references | Reject if production HTML contains developer comments |
| UI-ISD-003 | JavaScript source maps must not be deployed to production | Reject if source maps are accessible in production |
| UI-ISD-004 | API endpoint paths must not reveal internal architecture — use abstracted paths (/v1/transfers not /internal/core-banking/t24/transfer-service) | Reject if API paths expose internal system names |
| UI-ISD-005 | HTTP response headers must not reveal server technology — remove Server, X-Powered-By, X-AspNet-Version headers | Reject if technology-revealing headers are present |
| UI-ISD-006 | Debug endpoints, admin panels, and development tools must not be accessible in production | Reject if debug/admin endpoints are accessible in production |
| UI-ISD-007 | Feature flags must not be visible in client-side code or network requests in a way that reveals unreleased features | Flag if feature flag names/states are exposed to client |
| UI-ISD-008 | Third-party service names must not be exposed in error messages — "Payment processing unavailable" not "Stripe API returned 503" | Reject if third-party service names appear in user-facing errors |

---

## 6. Accessibility Guardrails (ref: DA-BP-001)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-ACG-001 | All customer-facing UI must meet WCAG 2.1 AA compliance (ref: DA-BP-001) | Reject if WCAG 2.1 AA audit fails |
| UI-ACG-002 | Automated accessibility scan (axe-core/Lighthouse) must run on every build — critical violations block merge | CI/CD gate — block on critical accessibility violations |
| UI-ACG-003 | Color contrast must meet minimum ratios: 4.5:1 for normal text, 3:1 for large text | Reject if contrast ratios are below minimum |
| UI-ACG-004 | All interactive elements must be keyboard-accessible with visible focus indicators (≥ 2px) | Reject if keyboard navigation is broken or focus is invisible |
| UI-ACG-005 | All form inputs must have associated labels — no placeholder-only labels | Reject if form inputs lack labels |
| UI-ACG-006 | All images and icons must have appropriate alt text or aria-label | Reject if non-decorative images lack text alternatives |
| UI-ACG-007 | Status indicators must not rely on color alone — use color + icon + text | Reject if status is communicated by color only |
| UI-ACG-008 | Touch targets must be minimum 44×44px on mobile | Reject if touch targets are below minimum |
| UI-ACG-009 | Reduced motion must be respected — honor prefers-reduced-motion media query | Flag if animations ignore reduced motion preference |
| UI-ACG-010 | Screen reader testing must be performed for all critical banking flows (login, transfer, payment, account management) before release | Reject if critical flows are not screen reader tested |

---

## 7. Transport Security Guardrails (ref: DA-BP-003)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-TSG-001 | All pages and resources must be served over HTTPS — zero HTTP resources (ref: DA-BP-003) | Reject if any HTTP resource is loaded |
| UI-TSG-002 | HSTS header must be present with max-age ≥ 31536000, includeSubDomains, and preload | Reject if HSTS is missing or misconfigured |
| UI-TSG-003 | Content-Security-Policy header must be configured — no unsafe-eval; minimize unsafe-inline | Reject if CSP is missing; flag if unsafe-eval is present |
| UI-TSG-004 | All mandatory security headers must be present (HSTS, CSP, X-Content-Type-Options, X-Frame-Options, Referrer-Policy, Permissions-Policy) | Reject if any mandatory header is missing |
| UI-TSG-005 | Third-party scripts must use Subresource Integrity (SRI) hashes | Reject if third-party scripts lack SRI |
| UI-TSG-006 | WebSocket connections must use WSS (secure WebSocket) — no WS | Reject if unsecured WebSocket is used |
| UI-TSG-007 | Sensitive pages must include Cache-Control: no-store and Pragma: no-cache | Reject if sensitive pages allow caching |

---

## 8. CSRF Protection Guardrails (ref: DA-BP-004)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-CSG-001 | CSRF protection must be implemented on all state-changing requests (POST, PUT, DELETE, PATCH) (ref: DA-BP-004) | Reject if any state-changing endpoint lacks CSRF protection |
| UI-CSG-002 | GET requests must not perform state-changing operations | Reject if GET requests modify data |
| UI-CSG-003 | CSRF tokens must be cryptographically random and validated server-side | Reject if CSRF tokens are predictable or client-validated only |
| UI-CSG-004 | Session cookies must include SameSite=Strict (or Lax as minimum) | Reject if SameSite attribute is missing or set to None without justification |
| UI-CSG-005 | CSRF validation failures must be logged as potential attack indicators | Flag if CSRF failures are not logged |
| UI-CSG-006 | Logout must be POST-based with CSRF token — not a GET link | Reject if logout is a GET request |

---

## 9. Privacy & Consent UI Guardrails (ref: DA-BP-005)

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-PCG-001 | Privacy notice must be accessible from every page (ref: DA-BP-005) | Reject if any page lacks privacy notice link |
| UI-PCG-002 | Privacy notice must be displayed before or at the point of data collection — not only in footer | Reject if data collection occurs without prior privacy notice |
| UI-PCG-003 | Consent checkboxes must be unchecked by default — no pre-ticked consent | Reject if consent is pre-selected |
| UI-PCG-004 | Consent must be granular — separate checkbox per purpose (marketing email, SMS, analytics, third-party sharing) | Reject if consent is bundled into a single checkbox |
| UI-PCG-005 | Consent withdrawal must be accessible from account settings — as easy as giving consent | Reject if withdrawal requires contacting support or is harder than consent |
| UI-PCG-006 | Cookie consent banner must be displayed on first visit with accept/reject/customize options | Reject if cookie banner lacks reject option |
| UI-PCG-007 | "Why do we need this?" help text must be available for sensitive data fields (national ID, income, employment) | Flag if sensitive fields lack purpose explanation |

---

## 10. Form & Input Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-FIG-001 | All user input must be validated server-side — client-side validation is for UX only | Reject if validation is client-side only |
| UI-FIG-002 | All user-generated content displayed in UI must be output-encoded to prevent XSS | Reject if output encoding is missing |
| UI-FIG-003 | File upload fields must validate file type, size, and content — executable files must be rejected | Reject if file upload has no validation |
| UI-FIG-004 | Form submission buttons must be disabled during processing to prevent double-submission | Flag if double-submit prevention is not implemented |
| UI-FIG-005 | Sensitive form fields (password, PIN, CVV) must use type="password" and prevent value exposure in DOM | Reject if sensitive fields use type="text" |
| UI-FIG-006 | Paste must be allowed on password and OTP fields — disabling paste harms accessibility and password manager usage | Reject if paste is disabled on password/OTP fields |

---

## 11. Mobile-Specific Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| UI-MOB-001 | Mobile apps must implement certificate pinning for API endpoints | Reject if certificate pinning is not implemented |
| UI-MOB-002 | Sensitive data must be stored in platform secure storage (iOS Keychain, Android Keystore) — not in shared preferences or plain files | Reject if sensitive data is in insecure storage |
| UI-MOB-003 | Jailbreak/root detection must be implemented with appropriate warning or restriction | Flag if no jailbreak/root detection exists |
| UI-MOB-004 | Deep links / universal links must validate parameters before processing — no blind navigation to sensitive screens | Reject if deep links are not validated |
| UI-MOB-005 | Push notifications must not contain sensitive data (account numbers, balances, transaction details) | Reject if push notifications contain sensitive data |
| UI-MOB-006 | Third-party keyboard input must be disabled for sensitive fields (password, PIN, OTP) on mobile | Flag if third-party keyboards are allowed on sensitive fields |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot proceed until fixed | Sensitive data in URLs (DA-GR-001), no MFA for transfers (DA-GR-002), timeout > 15 min (DA-GR-003), stack trace in error (DA-GR-004), server version exposed (DA-GR-005) |
| **Flag** | Can proceed but must be addressed before release | No concurrent session detection, no reduced motion support, no purpose help text on sensitive fields |
| **CI/CD Gate** | Automated enforcement in pipeline | Accessibility scan (axe-core), security header check, source map detection, console.log detection |

---

## Quick Reference: Guardrail Triggers by UI Component

| Component / Screen | Triggered Guardrails |
|---|---|
| Login page | UI-EMG-002 (no credential differentiation), UI-CSG-001 (CSRF), UI-MFA-001 (MFA if high-risk), UI-DEG-008 (autocomplete control) |
| Transfer form | UI-MFA-002 (MFA mandatory), UI-CSG-001 (CSRF), UI-DEG-001 (no data in URL), UI-STG-001 (15-min timeout), UI-FIG-001 (server-side validation) |
| Account dashboard | UI-DEG-009 (masked account numbers), UI-STG-001 (session timeout), UI-ACG-001 (WCAG 2.1 AA) |
| OTP/MFA screen | UI-MFA-007 (accessible), UI-MFA-008 (fallback), UI-DEG-011 (screenshot prevention), UI-DEG-008 (no autocomplete) |
| Profile settings | UI-MFA-004 (re-auth for changes), UI-PCG-005 (consent withdrawal), UI-CSG-001 (CSRF) |
| Error pages | UI-EMG-001 (no technical details), UI-EMG-007 (custom branded), UI-ISD-001 (no system details) |
| Mobile app | UI-MOB-001→006 (cert pinning, secure storage, jailbreak detection, deep link validation, push notification safety) |
| All pages | UI-TSG-001→007 (HTTPS, headers), UI-ACG-001→010 (accessibility), UI-PCG-001 (privacy notice) |

---

## Pre-Release UI/UX Checklist

| # | Check | Guardrail Ref |
|---|---|---|
| 1 | No sensitive data in URLs, localStorage, or browser history | DA-GR-001 |
| 2 | MFA implemented for all high-risk operations | DA-GR-002 |
| 3 | Session timeout ≤ 15 min inactivity with warning | DA-GR-003 |
| 4 | Error messages reveal no technical/internal details | DA-GR-004 |
| 5 | No internal system details exposed in UI, headers, or source | DA-GR-005 |
| 6 | WCAG 2.1 AA accessibility audit passed | DA-BP-001 |
| 7 | All security headers present and correctly configured | DA-BP-003 |
| 8 | CSRF protection on all state-changing requests | DA-BP-004 |
| 9 | Privacy notices and consent mechanisms in place | DA-BP-005 |
| 10 | Source maps removed from production build | UI-ISD-003 |
| 11 | Console.log statements removed from production | UI-DEG-006 |
| 12 | Mobile: cert pinning, secure storage, screenshot prevention | UI-MOB-001→006 |
