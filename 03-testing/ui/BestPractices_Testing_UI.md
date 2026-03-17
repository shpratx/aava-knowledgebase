# Best Practices for Test Scenario & Test Case Generation — UI Features
### Banking Domain — Agentic Knowledge Base

---

## 1. Test Scenario Design Principles

**Coverage Strategy:**
- Every user story acceptance criterion must have at least one test case
- Every security acceptance criterion must have a dedicated security test
- Every compliance acceptance criterion must have a compliance test
- Test the happy path FIRST, then systematically cover error paths

**Scenario Derivation:**
| Source | Derive |
|---|---|
| Acceptance criteria (Given/When/Then) | Direct test case per criterion |
| Business rules | Boundary value tests; rule violation tests |
| Data sensitivity tag | Security tests proportional to classification |
| Error/exception flows from user story | Error handling tests per flow |
| Regulatory linkage | Compliance verification tests |
| NFR targets | Performance/accessibility tests |

## 2. UI-Specific Best Practices

### Functional Testing
| Practice | Standard |
|---|---|
| Test user journeys, not components | End-to-end flow from entry to completion |
| Test with realistic data | Use data that mimics production patterns (masked) |
| Test all form states | Empty, partially filled, fully filled, invalid, submitted |
| Test all UI states | Default, loading, success, error, empty, disabled |
| Test navigation | Forward, back, refresh, deep link, breadcrumb |
| Test multi-step flows | Complete, abandon, go back and edit, timeout mid-flow |
| Test data persistence | Data survives page refresh (where expected); cleared on logout |
| Test concurrent actions | Double-click submit; rapid navigation; multiple tabs |

### Security Testing
| Practice | Standard |
|---|---|
| Test auth at every entry point | Direct URL access without login must redirect |
| Test IDOR via URL manipulation | Change account/transfer IDs in URL; verify 403 |
| Test session boundaries | Timeout, extend, concurrent, logout clears all |
| Test input injection | XSS in every text field; script in URL params |
| Test sensitive data exposure | Inspect localStorage, sessionStorage, DOM, network, console |
| Test error messages | No technical details; no account enumeration |
| Test MFA enforcement | Sensitive operations require MFA; expired OTP rejected |

### Accessibility Testing
| Practice | Standard |
|---|---|
| Keyboard-only complete flow | Every critical journey completable without mouse |
| Screen reader full flow | Test with VoiceOver (Mac/iOS), TalkBack (Android), NVDA (Windows) |
| Automated scan every build | axe-core integrated in CI; zero critical violations |
| Manual audit per release | Contrast, focus indicators, form labels, error announcements |
| Test at 200% zoom | No content loss; no horizontal scroll |
| Test reduced motion | Animations disabled when prefers-reduced-motion is set |

### Responsive Testing
| Practice | Standard |
|---|---|
| Test at breakpoints | 375px (mobile), 768px (tablet), 1024px (desktop), 1280px (large) |
| Test touch targets | All interactive elements ≥ 44×44px on mobile |
| Test orientation | Portrait and landscape on mobile/tablet |
| Test real devices | Not just browser emulation; test on actual iOS and Android devices |

## 3. Test Data Management

| Standard | Implementation |
|---|---|
| No production data | Synthetic data only; generated to mimic production patterns |
| Masked PII | Account numbers: ****1234; names: synthetic; emails: test@test.bank.com |
| Boundary values | Min amount ($0.01), max amount ($999,999,999.99), limit boundary ($50,000 ± $1) |
| Edge cases | Empty strings, very long strings (140 chars), special characters, Unicode, RTL text |
| State-specific data | Accounts: active, frozen, dormant, closed; Transfers: all statuses |
| Reusable fixtures | Shared test data fixtures; reset between test runs |

## 4. Test Automation Standards

| Standard | Implementation |
|---|---|
| Framework | Cypress or Playwright (not Selenium for new projects) |
| Page Object Model | Abstract UI interactions into page objects; tests read like user stories |
| Selectors | data-testid attributes (not CSS classes or XPath) |
| API mocking | MSW (Mock Service Worker) for unit/component tests; real API for E2E |
| Parallel execution | Tests run in parallel; no shared state between tests |
| CI/CD integration | Run on every PR; block merge on failure |
| Visual regression | Chromatic or Percy for design system components |
| Flaky test policy | Flaky tests quarantined within 24 hours; fixed within 1 week; never disabled permanently |

## 5. Common Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Only happy path tested | Errors, edge cases, security issues missed | Systematic coverage: happy + error + security + accessibility |
| Testing implementation, not behavior | Brittle tests that break on refactor | Test what user sees/does, not internal state |
| Hardcoded waits (sleep) | Slow, flaky tests | Use explicit waits (waitFor, assertions with retry) |
| Shared state between tests | Order-dependent; flaky | Each test sets up its own state; cleanup after |
| No accessibility tests | Excludes users; legal risk | axe-core in CI + manual screen reader testing |
| CSS selectors for test hooks | Break on styling changes | data-testid attributes |
| Testing only desktop | Mobile users can't use the app | Test at all breakpoints |
| No security tests in UI suite | Vulnerabilities shipped | Include auth, IDOR, injection, session tests |
