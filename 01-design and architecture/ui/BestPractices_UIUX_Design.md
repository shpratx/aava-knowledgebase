# Best Practices & Standards for UI/UX Design and Architecture

---

## 1. Accessibility Standards (ref: DA-BP-001)

### WCAG 2.1 AA Compliance

**Principle 1: Perceivable**

| Guideline | Requirement | Banking Implementation |
|---|---|---|
| 1.1 Text Alternatives | All non-text content has text alternative | Alt text for icons (transfer, payment, account), chart descriptions, CAPTCHA alternatives |
| 1.2 Time-Based Media | Captions and audio descriptions for media | Video tutorials for banking features must have captions and transcripts |
| 1.3 Adaptable | Content can be presented in different ways | Semantic HTML for account tables, transaction lists; proper heading hierarchy; form labels |
| 1.4 Distinguishable | Content is easy to see and hear | Contrast ratio ≥ 4.5:1 for text, ≥ 3:1 for large text; no color-only indicators for transaction status (use icons + color) |

**Principle 2: Operable**

| Guideline | Requirement | Banking Implementation |
|---|---|---|
| 2.1 Keyboard Accessible | All functionality via keyboard | Tab navigation through transfer forms, account selectors, menus; no keyboard traps |
| 2.2 Enough Time | Users have enough time | Session timeout warning 2 minutes before expiry with option to extend; no auto-advancing carousels for product offers |
| 2.3 Seizures | No content that causes seizures | No flashing elements in dashboards or transaction confirmations |
| 2.4 Navigable | Users can navigate and find content | Skip navigation links; breadcrumbs for multi-step processes (loan application, onboarding); focus management on page transitions |
| 2.5 Input Modalities | Multiple input methods supported | Touch targets ≥ 44×44px for mobile banking; support for voice input on forms |

**Principle 3: Understandable**

| Guideline | Requirement | Banking Implementation |
|---|---|---|
| 3.1 Readable | Text is readable and understandable | Plain language for financial terms; reading level appropriate for general public; language attribute set |
| 3.2 Predictable | Pages behave predictably | Consistent navigation across banking sections; no unexpected context changes on input |
| 3.3 Input Assistance | Help users avoid and correct errors | Inline validation on transfer forms; clear error messages ("Amount exceeds daily limit of $50,000" not "Error 422"); error prevention for irreversible actions (transfer confirmation step) |

**Principle 4: Robust**

| Guideline | Requirement | Banking Implementation |
|---|---|---|
| 4.1 Compatible | Content works with assistive technologies | Valid HTML; ARIA roles for custom components (account selector, amount input); tested with screen readers (NVDA, VoiceOver, JAWS) |

**Testing Requirements:**
- Automated: axe-core, Lighthouse accessibility audit on every build
- Manual: Screen reader testing (VoiceOver for macOS/iOS, TalkBack for Android, NVDA for Windows)
- Keyboard-only navigation testing for all critical flows (login, transfer, payment, account management)
- Color contrast verification for all UI components
- User testing with assistive technology users (annually)

**Banking-Specific Accessibility Considerations:**
- Account balance and transaction amounts must be readable by screen readers with proper formatting
- OTP/MFA input fields must be accessible — label, purpose, and error states announced
- PDF statements must be tagged/accessible or have HTML alternative
- Biometric authentication must have accessible fallback (PIN, password)
- Financial charts/graphs must have data table alternatives

---

## 2. Secure Session Management (ref: DA-BP-002)

### Session Configuration Standards

| Parameter | Standard | Banking Requirement |
|---|---|---|
| Session ID generation | Cryptographically random, ≥ 128-bit entropy | Use framework-provided secure session (e.g., Spring Session, Express session with secure store) |
| Session ID transport | HTTP-only, Secure, SameSite=Strict cookies only | Never in URL, never in localStorage |
| Inactivity timeout | 15 minutes for sensitive operations | Configurable per operation risk level |
| Absolute timeout | 8 hours maximum | Force re-authentication after 8 hours regardless of activity |
| Timeout warning | 2 minutes before expiry | Modal dialog with "Extend Session" option; accessible to screen readers |
| Session on login | Issue new session ID on authentication | Prevent session fixation attacks |
| Session on logout | Invalidate server-side immediately | Clear all session data; revoke tokens |
| Session on privilege change | Issue new session ID | When user elevates privileges (e.g., step-up MFA) |
| Concurrent sessions | Configurable per policy | Alert user of concurrent login; option to terminate other sessions |

### Token Management (for SPA/Mobile)

| Parameter | Standard | Implementation |
|---|---|---|
| Access token type | JWT (RS256 or ES256) | Short-lived, signed, not encrypted (no sensitive data in payload) |
| Access token lifetime | 15 minutes maximum | Refresh before expiry via refresh token |
| Refresh token type | Opaque (not JWT) | Stored server-side; single-use; rotated on each use |
| Refresh token lifetime | 24 hours maximum | Absolute expiry; revoked on logout |
| Token storage (web) | In-memory only (not localStorage/sessionStorage) | Use HTTP-only cookies for transport; BFF pattern preferred |
| Token storage (mobile) | Secure enclave / Keychain / Keystore | Platform-specific secure storage |
| Token revocation | Immediate on logout, password change, security event | Server-side revocation list or short token lifetime |

### Session Security Best Practices:
- Regenerate session ID after every authentication event (login, MFA, privilege escalation)
- Bind session to client fingerprint (IP range, user-agent) — alert on mismatch
- Implement session concurrency controls — detect and alert on simultaneous sessions from different locations
- Log all session events: creation, extension, timeout, logout, concurrent detection
- Display "last login" information to users — helps detect unauthorized access
- Implement "terminate all sessions" capability for users and administrators

---

## 3. HTTPS and Transport Security (ref: DA-BP-003)

### Transport Standards

| Requirement | Standard | Implementation |
|---|---|---|
| Protocol | HTTPS only — all pages, all resources | Redirect HTTP → HTTPS (301); no HTTP endpoints |
| TLS version | TLS 1.3 preferred, TLS 1.2 minimum | Disable SSL, TLS 1.0, TLS 1.1 |
| Certificate | Valid, trusted CA, ≥ 2048-bit RSA or ECC P-256 | Auto-renewal; monitor expiry; pin certificates for mobile apps |
| HSTS | Strict-Transport-Security: max-age=31536000; includeSubDomains; preload | Submit to HSTS preload list |
| Mixed content | Zero tolerance — no HTTP resources on HTTPS pages | CSP to block mixed content; automated scanning |
| Certificate transparency | Required | Monitor CT logs for unauthorized certificates |

### Security Headers

| Header | Value | Purpose |
|---|---|---|
| Strict-Transport-Security | max-age=31536000; includeSubDomains; preload | Force HTTPS |
| Content-Security-Policy | default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; connect-src 'self' api.bank.com; frame-ancestors 'none' | Prevent XSS, clickjacking, data injection |
| X-Content-Type-Options | nosniff | Prevent MIME-type sniffing |
| X-Frame-Options | DENY | Prevent clickjacking (backup for CSP frame-ancestors) |
| X-XSS-Protection | 0 | Disable browser XSS filter (rely on CSP instead) |
| Referrer-Policy | strict-origin-when-cross-origin | Control referrer information leakage |
| Permissions-Policy | camera=(), microphone=(), geolocation=(), payment=() | Restrict browser feature access |
| Cache-Control | no-store, no-cache, must-revalidate (for sensitive pages) | Prevent caching of sensitive data |
| Pragma | no-cache (for sensitive pages) | HTTP/1.0 backward compatibility |

### Banking-Specific Transport Considerations:
- API responses containing financial data must include Cache-Control: no-store
- PDF statement downloads must be served over HTTPS with Content-Disposition: attachment
- WebSocket connections (for real-time balance updates, notifications) must use WSS (WebSocket Secure)
- Third-party scripts (analytics, chat) must be loaded over HTTPS with Subresource Integrity (SRI)
- Mobile app certificate pinning for API endpoints — prevent MITM attacks

---

## 4. CSRF Protection (ref: DA-BP-004)

### CSRF Protection Standards

| Approach | When to Use | Implementation |
|---|---|---|
| Synchronizer Token Pattern | Server-rendered forms | Unique token per session/form; validate server-side on every state-changing request |
| Double Submit Cookie | SPA with API backend | CSRF token in cookie + request header; server validates match |
| SameSite Cookie Attribute | All applications | SameSite=Strict for session cookies; SameSite=Lax as minimum |
| Custom Request Header | API-only (no form submission) | Require custom header (e.g., X-Requested-With) that browsers don't send cross-origin |
| Origin/Referer Validation | Defense in depth (not sole protection) | Validate Origin header matches expected domain |

### CSRF Protection by Operation

| Operation | CSRF Risk | Protection Required |
|---|---|---|
| Fund transfer | Critical | Synchronizer token + SameSite=Strict + re-authentication (MFA) |
| Beneficiary add/modify/delete | High | Synchronizer token + SameSite=Strict |
| Profile update (address, email, phone) | High | Synchronizer token + SameSite=Strict + re-authentication |
| Password change | Critical | Synchronizer token + SameSite=Strict + current password verification |
| Consent grant/withdrawal | High | Synchronizer token + SameSite=Strict |
| Login form | Medium | Synchronizer token (prevent login CSRF) |
| Search / read-only operations | Low | SameSite cookie sufficient |
| Logout | Low | POST-based logout with token (prevent logout CSRF) |

### CSRF Best Practices:
- Apply CSRF protection to ALL state-changing requests (POST, PUT, DELETE, PATCH)
- Never use GET requests for state-changing operations
- CSRF tokens must be cryptographically random, per-session minimum (per-request preferred for high-risk)
- Validate CSRF token server-side — never rely on client-side validation
- CSRF token must not be in the URL — use hidden form field or request header
- For SPAs: use the BFF (Backend for Frontend) pattern to handle CSRF at the server layer
- Log CSRF validation failures — they may indicate an attack

---

## 5. Privacy Notices and Consent Mechanisms (ref: DA-BP-005)

### Privacy Notice Standards

| Element | Requirement | Implementation |
|---|---|---|
| Visibility | Privacy notice accessible from every page | Footer link "Privacy Policy" on all pages; prominent during data collection |
| Timing | Presented before or at the point of data collection | Privacy notice displayed before form submission; not buried in T&C |
| Language | Plain, clear language — no legal jargon | Reading level appropriate for general public; avoid "hereinafter", "notwithstanding" |
| Content | Purpose, legal basis, recipients, retention, rights, contact | Structured with headings; layered approach (summary + full notice) |
| Layered approach | Short notice at point of collection + link to full notice | Tooltip/expandable section with key points; link to full privacy policy |
| Updates | Notify users of material changes | Banner notification; email for significant changes; version history |
| Accessibility | WCAG 2.1 AA compliant | Screen reader compatible; sufficient contrast; keyboard navigable |

### Consent Mechanism Standards

| Element | Requirement | Implementation |
|---|---|---|
| Granularity | Separate consent per purpose | Individual checkboxes for: marketing email, marketing SMS, analytics, third-party sharing |
| Freely given | Not bundled with service access | Service works without optional consent; no "consent wall" for optional processing |
| Informed | Clear description of what user consents to | "We will send you weekly product offers via email" not "We may contact you" |
| Unambiguous | Affirmative action required | Unchecked checkbox that user must check; no pre-ticked boxes |
| Withdrawable | Easy to withdraw — as easy as giving | Settings page with toggle switches; one-click unsubscribe in emails |
| Recorded | Consent captured with full context | Timestamp, purpose, mechanism, privacy notice version, user identifier |
| Age verification | Parental consent for minors (< 16 in most EU countries) | Age gate before data collection; parental consent workflow if minor |

### Banking-Specific Privacy UI Patterns:

**Data Collection Forms:**
- Display purpose statement above each form section that collects personal data
- Mark mandatory vs. optional fields clearly — optional fields must have consent if used for secondary purposes
- "Why do we need this?" expandable help text per field for sensitive data (national ID, income, employment)

**Account Dashboard:**
- Privacy settings section: view/manage consents, download personal data, request erasure
- "Who has accessed my data" transparency log (for customer-facing audit trail)
- Cookie preference center accessible from every page

**Transaction Screens:**
- Clear indication of data shared with third parties (payment processors, fraud detection)
- Beneficiary data handling notice when adding new beneficiaries
- Cross-border transfer data sharing notice (which jurisdictions, what data)

**Onboarding Flow:**
- Step-by-step consent collection — not a single "agree to everything" page
- KYC data collection with clear purpose statement per document type
- Biometric consent (if applicable) with explicit opt-in and alternative offered

---

## 6. UI Architecture Standards

### Component Architecture

| Standard | Requirement | Rationale |
|---|---|---|
| Component library | Use a shared, versioned component library | Consistency, accessibility built-in, security controls centralized |
| Design system | Maintain a banking design system with tokens, patterns, and guidelines | Brand consistency, accessibility compliance, faster development |
| Responsive design | Mobile-first responsive design | Banking customers use mobile primarily; WCAG 2.1 requires responsive |
| Progressive enhancement | Core functionality works without JavaScript | Graceful degradation for older browsers; accessibility |
| Error boundaries | Implement error boundaries around every major UI section | Prevent full-page crashes; show user-friendly error states |
| Loading states | Skeleton screens or spinners for all async operations | User feedback during API calls; prevent double-submission |
| Offline handling | Graceful offline detection and messaging | Mobile banking in low-connectivity areas |

### State Management Security

| Concern | Standard | Implementation |
|---|---|---|
| Sensitive data in state | Minimize; clear on logout/timeout | Never store passwords, tokens, full PAN in component state |
| State persistence | No sensitive data in localStorage/sessionStorage | Use in-memory state only for sensitive data; clear on tab close |
| URL state | No sensitive data in URL/query params | Account numbers, balances, PII must never appear in URLs |
| Browser history | No sensitive data in history entries | Use replaceState for sensitive page transitions |
| Clipboard | Warn before copying sensitive data | Disable copy on CVV fields; auto-clear clipboard after paste for OTP |
| Auto-fill | Allow for non-sensitive fields; disable for OTP/CVV | autocomplete="off" for OTP; autocomplete="cc-csc" disabled |

### Performance Standards

| Metric | Target | Measurement |
|---|---|---|
| First Contentful Paint (FCP) | < 1.5 seconds | Lighthouse |
| Largest Contentful Paint (LCP) | < 2.5 seconds | Lighthouse / RUM |
| First Input Delay (FID) | < 100ms | RUM |
| Cumulative Layout Shift (CLS) | < 0.1 | Lighthouse / RUM |
| Time to Interactive (TTI) | < 3 seconds | Lighthouse |
| Bundle size (initial) | < 200KB gzipped | Build analysis |
| API response rendering | < 500ms from response to UI update | APM |

---

## 7. Form Design Standards

### Input Validation

| Layer | What to Validate | Implementation |
|---|---|---|
| Client-side | Format, length, required fields | Real-time inline validation; UX feedback only — not security |
| Server-side | All validation rules (authoritative) | Whitelist validation; reject on failure; never trust client |
| Display | Sanitize all output | Context-aware output encoding; prevent XSS |

### Banking Form Patterns

| Form Type | Validation Rules | Security Controls |
|---|---|---|
| Transfer form | Amount: decimal, 2dp, min 0.01, max per limit; Account: format validation; Reference: alphanumeric, max 140 chars | CSRF token; MFA step-up; server-side limit check |
| Login form | Username: email/ID format; Password: no max length restriction | CSRF token; rate limiting; no error message differentiation ("Invalid credentials" not "Wrong password") |
| OTP input | 6-digit numeric; auto-focus; auto-submit on complete | 90-second expiry timer displayed; single-use; no paste restriction (accessibility) |
| Address form | Country-specific format validation | Sanitize all fields; no script injection |
| Search | Sanitize input; limit result count | Server-side sanitization; parameterized queries; rate limiting |

### Error Message Standards

| Principle | Good Example | Bad Example |
|---|---|---|
| Specific and actionable | "Daily transfer limit exceeded. Remaining: $5,000" | "Error" |
| No technical details | "Something went wrong. Please try again." | "NullPointerException at TransferService.java:142" |
| No security information leakage | "Invalid credentials" | "Password incorrect for user john@bank.com" |
| Accessible | Error linked to field via aria-describedby; announced by screen reader | Red border only (not perceivable by screen readers or color-blind users) |
| Preventive | Inline validation before submission | Errors only shown after full form submission |

---

## 8. Mobile Banking UI Standards

### Platform-Specific Standards

| Standard | iOS | Android |
|---|---|---|
| Design guidelines | Human Interface Guidelines | Material Design |
| Biometric auth | Face ID / Touch ID via LocalAuthentication | Fingerprint / Face via BiometricPrompt |
| Secure storage | Keychain Services | Android Keystore |
| Certificate pinning | NSURLSession with pinned certificates | OkHttp CertificatePinner |
| Screenshot prevention | applicationWillResignActive → blur | FLAG_SECURE on sensitive screens |
| Jailbreak/root detection | Detect and warn/restrict | Detect and warn/restrict |
| Deep linking | Universal Links with validation | App Links with verification |
| Push notifications | No sensitive data in notification payload | No sensitive data in notification payload |

### Mobile-Specific Security UI

| Feature | Implementation |
|---|---|
| Balance masking | Toggle to show/hide balance on dashboard; default: masked |
| Transaction masking | Mask account numbers in transaction list (show last 4 digits) |
| Biometric prompt | Clear purpose statement: "Authenticate to view account balance" |
| Session indicator | Visual indicator of active session; countdown for timeout |
| Secure keyboard | Use secure text input for passwords/PINs; disable third-party keyboards for sensitive fields |
| Copy protection | Disable screenshot on sensitive screens (OTP, account details, card details) |
| Background blur | Blur app content when switching apps (prevent screenshot in app switcher) |

---

## 9. Internationalization & Localization

### Banking-Specific i18n Standards

| Element | Standard | Example |
|---|---|---|
| Currency formatting | Use locale-specific formatting; always show currency code | $1,234.56 (US), €1.234,56 (DE), ¥1,234 (JP) |
| Date formatting | Use locale-specific; store as ISO 8601 (UTC) | 03/15/2026 (US), 15/03/2026 (UK), 2026-03-15 (ISO) |
| Number formatting | Locale-specific decimal and thousands separators | 1,234.56 (US), 1.234,56 (DE) |
| Account number formatting | Country-specific (IBAN, BBAN, routing number) | DE89 3704 0044 0532 0130 00 (IBAN), 12-34-56 78901234 (UK sort code) |
| Right-to-left (RTL) | Full RTL support for Arabic, Hebrew | Mirrored layouts, RTL text alignment, bidirectional text handling |
| Language selection | User preference stored; browser language as default | Language selector accessible from every page |
| Legal content | Jurisdiction-specific privacy notices, T&C | Serve correct legal content based on user's jurisdiction |
| Error messages | Localized error messages | All user-facing text externalized for translation |

---

## 10. Design System Standards

### Component Library Requirements

| Component | Accessibility | Security | Banking-Specific |
|---|---|---|---|
| Amount input | aria-label, currency announcement | Server-side validation; no client-side limit bypass | Locale-aware formatting; currency selector |
| Account selector | aria-listbox, keyboard navigation | Own-account filtering server-side | Masked account numbers (last 4 digits); account type indicator |
| OTP input | aria-label per digit; auto-focus management | Auto-clear after timeout; no value in DOM after submission | Countdown timer; resend link with rate limiting |
| Transaction table | aria-table, sortable headers announced | Pagination server-side; no full dataset on client | Status indicators (color + icon + text); amount formatting |
| Confirmation dialog | aria-modal, focus trap, ESC to close | CSRF token in confirmation; prevent double-submit | Clear summary of action; "Are you sure?" for irreversible actions |
| Notification banner | aria-live="polite" for info, "assertive" for errors | No sensitive data in notifications | Dismissible; action links; categorized (info/warning/error/success) |
| Date picker | Keyboard navigable; aria-label for selected date | Server-side date validation | Business day awareness; cut-off time indication |
| File upload | aria-label; progress announcement | File type/size/content validation; virus scan | Document type guidance (ID, proof of address); max file size displayed |

### Design Tokens

| Token Category | Examples | Purpose |
|---|---|---|
| Colors | Primary, secondary, error, warning, success, neutral | Consistent branding; accessibility contrast ratios |
| Typography | Font family, sizes (h1-h6, body, caption), weights, line heights | Readability; hierarchy; accessibility |
| Spacing | 4px base unit; 8, 12, 16, 24, 32, 48, 64 | Consistent layout; touch target sizing |
| Borders | Radius (4px, 8px), width (1px, 2px), colors | Component styling; focus indicators (≥ 2px for accessibility) |
| Shadows | Elevation levels (card, modal, dropdown) | Depth hierarchy; focus indication |
| Breakpoints | Mobile (< 768px), Tablet (768-1024px), Desktop (> 1024px) | Responsive design |
| Motion | Duration (150ms, 300ms), easing (ease-in-out) | Reduced motion support (prefers-reduced-motion) |

---

## 11. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Accessibility as afterthought | Fails WCAG audit; excludes users; legal risk | Build accessibility into design system and test on every build |
| Session timeout without warning | User loses work; poor UX | 2-minute warning with extend option |
| Mixed HTTP/HTTPS content | Security vulnerability; browser warnings | HTTPS everywhere; CSP to block mixed content |
| CSRF protection on some forms only | Inconsistent protection; vulnerable forms exploited | CSRF on ALL state-changing requests |
| Privacy notice buried in footer | GDPR non-compliance; user unaware | Layered notices at point of data collection |
| Sensitive data in URLs | Leaked via referrer, logs, browser history | POST for sensitive operations; no PII in URLs |
| Color-only status indicators | Inaccessible to color-blind users | Color + icon + text for all status indicators |
| Infinite session | Security risk; account takeover | Inactivity + absolute timeout enforced |
| Client-side only validation | Bypassable; security vulnerability | Server-side validation authoritative; client-side for UX only |
| Generic error messages | Poor UX; user can't resolve issue | Specific, actionable, accessible error messages |
| No loading states | User uncertainty; double-submission | Skeleton screens; disable submit button during processing |
| Sensitive data in push notifications | Data exposure on lock screen | Generic notification; detail only after authentication |
