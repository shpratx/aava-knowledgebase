# Best Practices & Standards for Microfrontend Architecture

---

## 1. Microfrontend Principles

### When to Use Microfrontends in Banking

| Use When | Avoid When |
|---|---|
| Multiple teams own different banking domains (payments, lending, cards) | Single team owns the entire frontend |
| Independent deployment of UI features is required | Application is small with few features |
| Different domains have different release cadences | Tight coupling between all UI features |
| Legacy modernization — incrementally replacing monolith | Greenfield with single domain |
| Regulatory requirements differ by domain (PCI-DSS for cards vs. general for accounts) | Uniform regulatory requirements across all features |

### Composition Patterns

| Pattern | How It Works | Pros | Cons | Banking Use Case |
|---|---|---|---|---|
| **Build-time** | NPM packages composed at build | Type safety, optimized bundle | Coupled releases | Shared component library (design system) |
| **Server-side** | Server assembles fragments (SSI, ESI, Tailor) | Fast initial load, SEO | Server complexity | Account dashboard with widgets from different teams |
| **Client-side (Module Federation)** | Webpack Module Federation loads remotes at runtime | Independent deployment, shared dependencies | Runtime errors, version conflicts | Payments module loaded independently from accounts |
| **Client-side (iframes)** | Each microfrontend in an iframe | Complete isolation | Performance, UX limitations, accessibility challenges | Legacy system integration (temporary) |
| **Edge-side** | CDN/edge assembles fragments | Performance, caching | Limited interactivity | Static content composition |

**Recommended for Banking: Module Federation** (primary) + **Build-time** (shared design system)

### Domain Decomposition for Banking

```
Shell (Host Application)
├── Authentication Microfrontend (login, MFA, session management)
├── Accounts Microfrontend (balances, transactions, statements)
├── Payments Microfrontend (transfers, bill pay, standing orders)
├── Cards Microfrontend (card management, limits, PIN) [PCI-DSS scope]
├── Lending Microfrontend (loan applications, repayments)
├── Onboarding Microfrontend (KYC, account opening)
├── Settings Microfrontend (profile, preferences, consent management)
└── Shared Design System (component library, tokens, utilities)
```

---

## 2. Architecture Standards

### Shell (Host) Responsibilities

| Responsibility | Implementation |
|---|---|
| Routing | Top-level route mapping to microfrontends; lazy-load on navigation |
| Authentication | Centralized auth; pass tokens to microfrontends via shared context |
| Session management | Centralized timeout, warning, logout — microfrontends consume session state |
| Global navigation | Header, footer, sidebar owned by shell |
| Error boundaries | Catch microfrontend load/render failures; show fallback UI |
| Shared state | Minimal — only auth context, user preferences, feature flags |
| Security headers | CSP, HSTS, security headers set by shell/gateway |
| Analytics | Centralized analytics collection; microfrontends emit events |

### Microfrontend Responsibilities

| Responsibility | Standard |
|---|---|
| Own domain UI | Complete ownership of domain-specific screens and components |
| Own domain state | Local state management; no dependency on other microfrontends' state |
| Own domain API calls | Direct API calls to domain backend; no routing through other microfrontends |
| Own domain tests | Unit, integration, accessibility tests within the microfrontend |
| Expose contract | Defined entry point, props interface, events emitted |
| Use shared design system | Consume shared components for consistency |
| Handle own errors | Graceful error handling within domain; propagate critical errors to shell |

### Communication Between Microfrontends

| Method | When to Use | Banking Example |
|---|---|---|
| Custom Events (pub/sub) | Loose coupling; one-to-many notifications | Transfer completed → Accounts refreshes balance |
| Shared state (minimal) | Auth context, user preferences | Session token, language preference |
| URL/route params | Navigation with context | Navigate to transfer with pre-selected account |
| Props (parent → child) | Shell passes config to microfrontend | Feature flags, user role |
| **Avoid:** Direct imports | Never — creates coupling | — |
| **Avoid:** Shared mutable state | Never — race conditions, coupling | — |

### Event Bus Standards

```typescript
// Event contract — shared type definitions
interface BankingEvent {
  type: string;
  source: string;        // microfrontend name
  timestamp: string;     // ISO 8601
  correlationId: string; // trace across microfrontends
  payload: unknown;
}

// Domain events
type TransferCompletedEvent = BankingEvent & {
  type: 'transfer.completed';
  payload: { transferId: string; amount: number; currency: string };
};

type SessionExpiringEvent = BankingEvent & {
  type: 'session.expiring';
  payload: { expiresIn: number }; // seconds
};
```

---

## 3. Security Standards

| Concern | Standard |
|---|---|
| Authentication | Centralized in shell; tokens passed via secure context (not URL, not localStorage) |
| CSP | Single CSP policy managed by shell; microfrontends must not require unsafe-eval |
| Isolation | Microfrontends must not access other microfrontends' DOM or state directly |
| Dependencies | Shared dependencies (React, design system) version-locked; security patches coordinated |
| PCI-DSS scope | Cards microfrontend isolated; minimal shared surface with non-PCI components |
| Sensitive data | Each microfrontend clears its own sensitive data on unmount/logout |
| Third-party scripts | Only shell loads third-party scripts; microfrontends must not load external scripts independently |

---

## 4. Performance Standards

| Metric | Target | Implementation |
|---|---|---|
| Initial shell load | < 2 seconds | Minimal shell; lazy-load microfrontends |
| Microfrontend load | < 1 second | Code splitting; shared vendor chunk |
| Shared dependency size | < 100KB gzipped | React + design system shared; no duplication |
| Total bundle (per microfrontend) | < 150KB gzipped | Tree shaking; no unnecessary dependencies |
| Navigation between microfrontends | < 500ms | Prefetch on hover/visibility; cached modules |

### Shared Dependency Management

| Dependency | Sharing Strategy | Version Policy |
|---|---|---|
| React / React DOM | Singleton (Module Federation shared) | All microfrontends on same major version |
| Design system | Singleton | All on same minor version; patch updates independent |
| State management (if any) | Not shared — each microfrontend owns its state | Independent |
| HTTP client (axios/fetch) | Not shared — each microfrontend uses its own | Independent |
| Utility libraries (lodash, date-fns) | Not shared unless > 50KB savings | Independent |

---

## 5. Testing Standards

| Test Type | Scope | Tool |
|---|---|---|
| Unit tests | Individual microfrontend components | Jest, React Testing Library |
| Integration tests | Microfrontend with its own API | Cypress component tests, MSW for API mocking |
| Contract tests | Event bus contracts between microfrontends | Pact or custom schema validation |
| E2E tests | Full application with all microfrontends composed | Cypress, Playwright |
| Accessibility tests | Each microfrontend independently + composed | axe-core, Lighthouse |
| Visual regression | Each microfrontend independently | Chromatic, Percy |
| Performance tests | Each microfrontend load time + composed | Lighthouse CI |

---

## 6. Deployment Standards

| Standard | Implementation |
|---|---|
| Independent deployment | Each microfrontend deployable without redeploying others |
| Versioned artifacts | Each microfrontend versioned independently (semver) |
| Feature flags | Gradual rollout per microfrontend; instant rollback via flag |
| Canary deployment | Route percentage of traffic to new version |
| Rollback | Revert to previous version within 5 minutes |
| Environment parity | Same composition in staging and production |
| Health checks | Each microfrontend exposes health endpoint; shell monitors |

---

## 7. Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Shared mutable global state | Race conditions, tight coupling | Event bus for communication; minimal shared context |
| Microfrontend loading other microfrontends | Dependency chain, cascading failures | Only shell orchestrates composition |
| Inconsistent design | Fragmented UX | Shared design system enforced |
| Each microfrontend bundles React | Massive bundle size, version conflicts | Module Federation shared singleton |
| Direct DOM manipulation across boundaries | Fragile, breaks on updates | Custom events only |
| No error boundaries | One microfrontend crash takes down entire app | Error boundary per microfrontend |
