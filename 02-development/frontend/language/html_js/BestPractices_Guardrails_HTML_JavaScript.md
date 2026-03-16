# HTML & JavaScript Development Standards, Best Practices & Guardrails
---

## Part A: HTML Standards & Best Practices

### 1. Semantic HTML

| Standard | Requirement | Banking Example |
|---|---|---|
| Use semantic elements | header, nav, main, section, article, aside, footer — not div for everything | Account dashboard: main > section (balance) + section (transactions) |
| Heading hierarchy | h1 → h2 → h3 sequential; never skip levels | h1: Account Overview, h2: Balance, h2: Recent Transactions |
| Form semantics | form, fieldset, legend, label — every input has a label | Transfer form: fieldset (Transfer Details) > label + input pairs |
| Table semantics | table, thead, tbody, th (with scope), caption | Transaction history: caption + thead (Date, Description, Amount) |
| Lists | ul/ol for lists; never div sequences | Beneficiary list: ul > li per beneficiary |
| Landmarks | role attributes only when no semantic equivalent | nav for navigation; main for content; no role="navigation" on nav |
| Language | lang attribute on html element | <html lang="en"> |

### 2. Accessibility (WCAG 2.1 AA)

| Standard | Implementation |
|---|---|
| All images | alt text (descriptive) or alt="" (decorative) |
| All form inputs | Associated label via for/id or wrapping; aria-describedby for help text |
| All interactive elements | Keyboard accessible; visible focus indicator (>= 2px) |
| Color contrast | >= 4.5:1 normal text; >= 3:1 large text |
| Status indicators | Color + icon + text — never color alone |
| Error messages | aria-describedby linking error to field; aria-live for dynamic errors |
| Modals | aria-modal, focus trap, ESC to close, return focus on close |
| Loading states | aria-busy, aria-live="polite" for completion announcement |
| Skip navigation | Skip to main content link as first focusable element |
| Touch targets | Minimum 44x44px on mobile |

### 3. HTML Security

| Standard | Implementation |
|---|---|
| No inline scripts | All JavaScript in external files; CSP blocks inline |
| No inline event handlers | No onclick, onsubmit in HTML — use addEventListener |
| No sensitive data in HTML | No tokens, account numbers, PII in hidden fields, data attributes, or comments |
| No developer comments in production | Strip HTML comments from production builds |
| Form security | autocomplete="off" for OTP/CVV; method="POST" for sensitive forms |
| iframe restrictions | sandbox attribute if iframes needed; X-Frame-Options: DENY |
| Meta tags | charset UTF-8; viewport for responsive; no referrer for sensitive pages |

### 4. HTML Performance

| Standard | Implementation |
|---|---|
| Minimal DOM | Avoid deeply nested divs; flatten structure where possible |
| Async/defer scripts | script defer for non-critical; script async for independent |
| Preload critical resources | <link rel="preload"> for fonts, critical CSS |
| Lazy load images | loading="lazy" on below-fold images |
| Responsive images | srcset + sizes for resolution-appropriate images |
| Preconnect | <link rel="preconnect"> for API domains, CDN |

---

## Part B: JavaScript Standards & Best Practices

### 5. Language Standards

| Standard | Requirement |
|---|---|
| ECMAScript version | ES2020+ (transpiled for target browsers via Babel/SWC) |
| Strict mode | "use strict" enforced (automatic in ES modules) |
| TypeScript preferred | TypeScript for all new code; strict mode enabled |
| Variable declarations | const by default; let when reassignment needed; never var |
| Equality | Always === and !==; never == or != |
| Null handling | Optional chaining (?.) and nullish coalescing (??); no non-null assertions |
| Error handling | try/catch for async; .catch() for promises; never swallow errors silently |
| Async patterns | async/await preferred over .then() chains; Promise.all for parallel |

### 6. Security Standards

#### 6.1 XSS Prevention

| Rule | Implementation |
|---|---|
| No eval() | Prohibited — no dynamic code execution |
| No Function() constructor | Prohibited |
| No setTimeout/setInterval with strings | Use function references only |
| No innerHTML/outerHTML | Use textContent for text; DOM API for elements; framework rendering |
| No document.write() | Prohibited |
| Template literal safety | Never interpolate user input into HTML template literals |
| URL validation | Validate all dynamic URLs; reject javascript: protocol |
| JSON.parse safety | Wrap in try/catch; validate structure after parsing |

#### 6.2 Input Handling

| Rule | Implementation |
|---|---|
| Client-side validation is UX only | Never trust client validation for security — server validates authoritatively |
| Sanitize before display | Use framework escaping (React JSX) or DOMPurify for raw HTML |
| Encode output | Context-aware encoding: HTML entities for HTML, URL encoding for URLs |
| Limit input length | maxlength on inputs; validate in JS before submission |
| Whitelist characters | Regex validation for allowed patterns (account refs: alphanumeric only) |
| File upload validation | Check type, size, extension client-side (server validates authoritatively) |

#### 6.3 Sensitive Data Handling

| Rule | Implementation |
|---|---|
| No sensitive data in global scope | No window.accountNumber, no global variables with PII |
| No sensitive data in localStorage/sessionStorage | Tokens in HTTP-only cookies or in-memory only |
| No sensitive data in URLs | POST for sensitive operations; no PII in query params |
| No sensitive data in console | No console.log with account data; strip console in production |
| No sensitive data in error messages | Catch errors; display generic messages; log details server-side |
| Clear sensitive data | Null out variables holding sensitive data after use; clear on logout |
| Clipboard protection | Auto-clear clipboard after 60s for copied sensitive data |

#### 6.4 Authentication & Session

| Rule | Implementation |
|---|---|
| Token storage | HTTP-only Secure SameSite=Strict cookies (server-managed) or in-memory |
| Token refresh | Automatic refresh before expiry; redirect to login on failure |
| Session timeout | Client-side timer for UX warning (2 min before); server enforces authoritatively |
| Logout cleanup | Clear all state, tokens, caches, timers; redirect to login |
| CSRF protection | Include CSRF token in state-changing requests; SameSite cookies |
| Concurrent session | Detect and alert user of concurrent login from different location |

### 7. API Communication

| Standard | Implementation |
|---|---|
| HTTPS only | All API calls over HTTPS; reject HTTP |
| Correlation ID | Generate UUID per request; send as X-Correlation-ID header |
| Error handling | Never display raw API errors; map to user-friendly messages |
| Timeout | Configure request timeout (30s default); handle timeout gracefully |
| Retry | Exponential backoff for idempotent requests; max 3 retries |
| Rate limiting | Respect 429 responses; implement client-side throttle for UX |
| Request cancellation | Cancel in-flight requests on component unmount (AbortController) |
| Response validation | Validate response structure before use; don't trust API blindly |

### 8. Error Handling

| Standard | Implementation |
|---|---|
| Global error handler | window.onerror + window.onunhandledrejection → monitoring service |
| Never swallow errors | Every catch must log or re-throw; no empty catch blocks |
| User-friendly messages | "Something went wrong. Please try again." — no stack traces |
| Error boundaries (React) | Per-feature error boundaries with fallback UI |
| Network errors | Detect offline; show appropriate message; retry when online |
| Validation errors | Inline, field-level, accessible (aria-describedby) |
| Monitoring integration | Send errors to Sentry/Datadog with correlation ID, user context (no PII) |

### 9. Performance

| Standard | Implementation |
|---|---|
| Bundle splitting | Route-level code splitting; lazy load non-critical features |
| Tree shaking | Named imports only; avoid import * |
| Debounce/throttle | Debounce search inputs (300ms); throttle scroll handlers (100ms) |
| Web Workers | Offload heavy computation (encryption, data transformation) |
| Virtualization | Virtual scrolling for lists > 50 items |
| Memory management | Clean up event listeners, timers, subscriptions on unmount |
| Avoid memory leaks | No closures holding large data; WeakMap/WeakRef for caches |
| requestAnimationFrame | Use for animations instead of setTimeout/setInterval |

### 10. Module & Import Standards

| Standard | Implementation |
|---|---|
| ES Modules | import/export exclusively; no CommonJS require() in frontend |
| Named exports | Prefer named over default exports for refactoring safety |
| Barrel exports | index.ts per feature for clean imports; avoid re-exporting large modules |
| Import order | External packages → shared modules → feature modules → relative imports |
| Dynamic imports | import() for code splitting; wrap in Suspense/error boundary |
| Circular dependencies | Prohibited — ESLint rule to detect |
| Side effects | No side effects in module scope; use initialization functions |

---

## Part C: Guardrails

### 11. HTML Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| HJ-HG-001 | All pages must have valid semantic HTML — validated by HTML linter | CI/CD gate |
| HJ-HG-002 | All form inputs must have associated labels — no unlabeled inputs | CI/CD gate — axe-core |
| HJ-HG-003 | All images must have alt attributes | CI/CD gate — axe-core |
| HJ-HG-004 | No inline scripts or inline event handlers in HTML | CI/CD gate — CSP + linter |
| HJ-HG-005 | No sensitive data in HTML source (hidden fields, data attributes, comments) | Reject if sensitive data found in HTML |
| HJ-HG-006 | No developer comments in production HTML | CI/CD gate — build strips comments |
| HJ-HG-007 | Color contrast must meet WCAG 2.1 AA minimums | CI/CD gate — axe-core |
| HJ-HG-008 | All interactive elements must be keyboard accessible | Reject if keyboard navigation broken |
| HJ-HG-009 | lang attribute must be set on html element | CI/CD gate — linter |
| HJ-HG-010 | All pages must include required meta tags (charset, viewport) | CI/CD gate — linter |

### 12. JavaScript Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| HJ-SG-001 | eval(), Function(), document.write() are prohibited | Reject immediately — ESLint no-eval |
| HJ-SG-002 | innerHTML/outerHTML are prohibited — use textContent or framework rendering | Reject — ESLint no-inner-html |
| HJ-SG-003 | setTimeout/setInterval with string arguments are prohibited | Reject — ESLint no-implied-eval |
| HJ-SG-004 | No sensitive data in global scope (window object) | Reject if sensitive data on window |
| HJ-SG-005 | No sensitive data in localStorage, sessionStorage, or IndexedDB | Reject if sensitive data in browser storage |
| HJ-SG-006 | No sensitive data in console output — console stripped from production | CI/CD gate — no-console rule + build strip |
| HJ-SG-007 | All user input must be sanitized before display (DOMPurify or framework escaping) | Reject if unsanitized input rendered |
| HJ-SG-008 | All dynamic URLs must be validated — reject javascript: protocol | Reject if URL validation missing |
| HJ-SG-009 | All API errors must be caught and translated — no raw error display | Reject if raw API errors shown to user |
| HJ-SG-010 | No hardcoded secrets, API keys, or credentials in JavaScript | Reject — secret scanning in CI/CD |

### 13. JavaScript Quality Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| HJ-QG-001 | TypeScript strict mode enabled (strict: true, noImplicitAny: true) | CI/CD gate — tsc |
| HJ-QG-002 | No any type — use unknown + type guards or specific types | CI/CD gate — ESLint no-explicit-any |
| HJ-QG-003 | ESLint must pass with zero errors on every build | CI/CD gate |
| HJ-QG-004 | Prettier formatting must be consistent — checked in CI | CI/CD gate |
| HJ-QG-005 | No var declarations — const by default, let when needed | CI/CD gate — ESLint no-var |
| HJ-QG-006 | Always use === and !== — never == or != | CI/CD gate — ESLint eqeqeq |
| HJ-QG-007 | No empty catch blocks — every catch must log or handle | CI/CD gate — ESLint no-empty |
| HJ-QG-008 | No unused variables or imports | CI/CD gate — ESLint no-unused-vars |
| HJ-QG-009 | Maximum cyclomatic complexity: 10 per function | CI/CD gate — ESLint complexity |
| HJ-QG-010 | No circular dependencies | CI/CD gate — ESLint import/no-cycle |
| HJ-QG-011 | Test coverage >= 80% overall | CI/CD gate |
| HJ-QG-012 | All async operations must have error handling (try/catch or .catch()) | CI/CD gate — ESLint no-floating-promises |

### 14. Performance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| HJ-PG-001 | Initial bundle must not exceed 200KB gzipped | CI/CD gate |
| HJ-PG-002 | No synchronous blocking operations on main thread > 50ms | Flag — performance profiling |
| HJ-PG-003 | Event listeners must be cleaned up on component unmount | Reject if cleanup missing |
| HJ-PG-004 | Timers (setTimeout/setInterval) must be cleared on unmount | Reject if cleanup missing |
| HJ-PG-005 | Subscriptions (WebSocket, EventSource) must be closed on unmount | Reject if cleanup missing |
| HJ-PG-006 | API requests must be cancelled on component unmount (AbortController) | Flag if cancellation missing |
| HJ-PG-007 | Lists > 50 items must use virtualization | Flag if not virtualized |
| HJ-PG-008 | Images must use lazy loading for below-fold content | Flag if not lazy loaded |

### 15. Dependency Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| HJ-DG-001 | All dependencies must be security-scanned on every build | CI/CD gate — npm audit |
| HJ-DG-002 | Critical CVEs must be patched within 24 hours | Reject deployment |
| HJ-DG-003 | Only approved licenses (MIT, Apache 2.0, BSD, ISC) | CI/CD gate — license checker |
| HJ-DG-004 | New dependencies require security review | Reject if added without review |
| HJ-DG-005 | Lock file must be committed; deterministic installs (npm ci) | CI/CD gate |
| HJ-DG-006 | CDN resources must include SRI hashes | Reject if SRI missing |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot merge/deploy | eval() usage, innerHTML, sensitive data in storage, hardcoded secrets, unlabeled inputs |
| **Flag** | Address before release | Missing lazy loading, large list not virtualized, missing API cancellation |
| **CI/CD Gate** | Automated | ESLint, TypeScript, axe-core, npm audit, bundle size, coverage, SRI check |

---

## ESLint Configuration Reference

```json
{
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/strict-type-checked",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:jsx-a11y/recommended",
    "plugin:import/recommended",
    "prettier"
  ],
  "rules": {
    "no-eval": "error",
    "no-implied-eval": "error",
    "no-console": "error",
    "no-var": "error",
    "eqeqeq": "error",
    "no-empty": "error",
    "no-unused-vars": "error",
    "complexity": ["error", 10],
    "@typescript-eslint/no-explicit-any": "error",
    "@typescript-eslint/no-floating-promises": "error",
    "react/no-danger": "error",
    "import/no-cycle": "error",
    "jsx-a11y/label-has-associated-control": "error"
  }
}
```

---

## Pre-Deployment Checklist

| # | Check |
|---|---|
| 1 | Valid semantic HTML; all inputs labeled; all images have alt |
| 2 | WCAG 2.1 AA: contrast, keyboard, screen reader tested |
| 3 | No eval/innerHTML/document.write in codebase |
| 4 | No sensitive data in HTML, localStorage, console, URLs, global scope |
| 5 | All user input sanitized; all dynamic URLs validated |
| 6 | TypeScript strict; ESLint zero errors; Prettier consistent |
| 7 | Test coverage >= 80%; accessibility scan passed |
| 8 | All dependencies scanned; zero critical/high CVEs |
| 9 | SRI on CDN resources; CSP configured |
| 10 | Bundle size within limits; lazy loading implemented |
| 11 | All event listeners, timers, subscriptions cleaned up on unmount |
| 12 | No console statements in production build |
| 13 | No source maps in production |
| 14 | Error monitoring configured (Sentry/Datadog) |
