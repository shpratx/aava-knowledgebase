
## 7. Content Security Policy (ref: DV-BP-005)

### CSP Configuration

```html
<!-- Strict CSP for banking React app -->
<meta http-equiv="Content-Security-Policy" content="
  default-src 'self';
  script-src 'self';
  style-src 'self' 'unsafe-inline';
  img-src 'self' data: https://cdn.bank.com;
  font-src 'self' https://fonts.bank.com;
  connect-src 'self' https://api.bank.com wss://ws.bank.com;
  frame-src 'none';
  frame-ancestors 'none';
  object-src 'none';
  base-uri 'self';
  form-action 'self';
  upgrade-insecure-requests;
">
```

**Preferred: Set via HTTP response header (not meta tag) at server/CDN level.**

### CSP Standards

| Directive | Standard | Rationale |
|---|---|---|
| default-src | 'self' | Deny everything not explicitly allowed |
| script-src | 'self' (no 'unsafe-inline', no 'unsafe-eval') | Prevent XSS; use nonce-based if inline needed |
| style-src | 'self' 'unsafe-inline' (or nonce-based) | CSS-in-JS may require unsafe-inline; prefer nonce |
| img-src | 'self' data: + specific CDN domains | Allow images from own domain and CDN |
| connect-src | 'self' + specific API domains | Restrict API calls to known backends |
| frame-src | 'none' | No iframes (prevent clickjacking) |
| frame-ancestors | 'none' | Prevent embedding in other sites |
| object-src | 'none' | Block Flash/plugins |
| base-uri | 'self' | Prevent base tag injection |
| form-action | 'self' | Prevent form submission to external sites |

### CSP Implementation for React

| Build Tool | Implementation |
|---|---|
| Vite | vite-plugin-csp; nonce injection via server middleware |
| Create React App | Set via server (nginx/CDN); not in build |
| Next.js | next.config.js headers; middleware for nonce |

### CSP Violation Monitoring

```tsx
// Report CSP violations to monitoring
document.addEventListener('securitypolicyviolation', (event) => {
  auditService.logSecurityEvent({
    type: 'CSP_VIOLATION',
    blockedUri: event.blockedURI,
    violatedDirective: event.violatedDirective,
    originalPolicy: event.originalPolicy,
    timestamp: new Date().toISOString(),
  });
});
```

---

## 8. Dependency Security (ref: DV-BP-006)

### Dependency Management Standards

| Standard | Requirement | Enforcement |
|---|---|---|
| Lock file | package-lock.json or yarn.lock committed | CI fails without lock file |
| Audit on install | npm audit / yarn audit on every install | CI/CD gate |
| Automated scanning | Snyk or Dependabot continuous monitoring | Automated PRs for vulnerabilities |
| Critical CVE | Patch within 24 hours | Blocking — no deployment |
| High CVE | Patch within 7 days | Blocking before next release |
| Medium CVE | Patch within 30 days | Tracked in backlog |
| New dependency | Security review before adding | PR review checklist |
| License compliance | Only approved licenses (MIT, Apache 2.0, BSD) | License checker in CI |

### Dependency Review Checklist (for new packages)

| Check | Criteria |
|---|---|
| Maintenance | Last published < 6 months; active maintainers |
| Popularity | > 1000 weekly downloads; established community |
| Security | No known CVEs; security policy published |
| License | MIT, Apache 2.0, BSD, ISC — no GPL for proprietary banking apps |
| Size | Justified bundle size impact; tree-shakeable |
| Alternatives | Evaluated against alternatives; justified choice |
| Transitive deps | Reviewed transitive dependency tree for risks |

### Approved Core Dependencies (Banking React)

| Category | Package | Purpose |
|---|---|---|
| Framework | react, react-dom | UI framework |
| Routing | react-router-dom | Client-side routing |
| State | @reduxjs/toolkit, react-redux | Global state (if needed) |
| Data fetching | @tanstack/react-query | Server state, caching, sync |
| Forms | react-hook-form + zod | Form management + validation |
| HTTP | axios | API client with interceptors |
| Sanitization | dompurify | XSS prevention |
| Date | date-fns | Date formatting (tree-shakeable) |
| i18n | react-i18next | Internationalization |
| Testing | jest, @testing-library/react, msw | Unit/integration testing |
| Accessibility | @axe-core/react | Dev-time accessibility checking |
| Linting | eslint, prettier | Code quality |

---

## 9. State Management (ref: DV-BP-007)

### State Categories

| Category | Where to Store | Tool | Banking Example |
|---|---|---|---|
| Server state | React Query cache | @tanstack/react-query | Account balances, transactions, transfer status |
| Global UI state | Redux or Context | @reduxjs/toolkit | Theme, language, feature flags |
| Auth state | Secure context | Custom AuthProvider | Session status, user role, MFA status |
| Form state | Local (component) | react-hook-form | Transfer form inputs, validation |
| Component state | Local (useState) | React hooks | Modal open/close, dropdown selection |
| URL state | Router | react-router-dom | Current page, filters, pagination |

### State Management Rules

| Rule | Standard |
|---|---|
| Server state via React Query | Never store API data in Redux — use React Query for fetching, caching, sync |
| Minimal global state | Only auth, theme, language, feature flags in global state |
| No sensitive data in state | Never store passwords, tokens, full PAN in React state |
| Clear on logout | All state cleared on logout/session timeout |
| Immutable updates | Always create new objects/arrays — never mutate state directly |
| Derived state | Compute from source — don't store derived values in state |

### React Query Standards (Server State)

```tsx
// Standard query pattern for banking data
export function useAccountBalance(accountId: string) {
  return useQuery({
    queryKey: ['accounts', accountId, 'balance'],
    queryFn: () => accountApi.getBalance(accountId),
    staleTime: 0,           // Balance always fresh (no stale cache)
    gcTime: 5 * 60 * 1000,  // Keep in cache 5 min for back-navigation
    retry: 2,
    refetchOnWindowFocus: true,  // Refresh when user returns to tab
  });
}

// Mutation pattern for transfers
export function useInitiateTransfer() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (request: TransferRequest) => transferApi.initiate(request),
    onSuccess: (data) => {
      // Invalidate related queries
      queryClient.invalidateQueries({ queryKey: ['accounts'] });
      queryClient.invalidateQueries({ queryKey: ['transfers'] });
    },
  });
}
```

### Redux Standards (Global UI State — if needed)

```tsx
// Slice pattern — minimal global state
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface AppState {
  theme: 'light' | 'dark' | 'high-contrast';
  language: string;
  featureFlags: Record<string, boolean>;
}

const appSlice = createSlice({
  name: 'app',
  initialState: { theme: 'light', language: 'en', featureFlags: {} } as AppState,
  reducers: {
    setTheme: (state, action: PayloadAction<AppState['theme']>) => {
      state.theme = action.payload;
    },
    setLanguage: (state, action: PayloadAction<string>) => {
      state.language = action.payload;
    },
    setFeatureFlags: (state, action: PayloadAction<Record<string, boolean>>) => {
      state.featureFlags = action.payload;
    },
  },
});
```

---

## 10. Code Splitting & Performance (ref: DV-BP-008)

### Code Splitting Standards

```tsx
// Route-level code splitting — every feature lazy-loaded
import { lazy, Suspense } from 'react';
import { LoadingSkeleton } from '@/shared/components/feedback/LoadingSkeleton';

const AccountsFeature = lazy(() => import('@/features/accounts'));
const TransfersFeature = lazy(() => import('@/features/transfers'));
const CardsFeature = lazy(() => import('@/features/cards'));
const SettingsFeature = lazy(() => import('@/features/settings'));

export function AppRoutes() {
  return (
    <Suspense fallback={<LoadingSkeleton />}>
      <Routes>
        <Route path="/accounts/*" element={<AccountsFeature />} />
        <Route path="/transfers/*" element={<TransfersFeature />} />
        <Route path="/cards/*" element={<CardsFeature />} />
        <Route path="/settings/*" element={<SettingsFeature />} />
      </Routes>
    </Suspense>
  );
}
```

### Performance Standards

| Metric | Target | Measurement |
|---|---|---|
| Initial bundle (gzipped) | < 200KB | Build analysis (webpack-bundle-analyzer / vite-bundle-visualizer) |
| Feature chunk (gzipped) | < 100KB per feature | Build analysis |
| First Contentful Paint | < 1.5s | Lighthouse |
| Largest Contentful Paint | < 2.5s | Lighthouse |
| Time to Interactive | < 3s | Lighthouse |
| Cumulative Layout Shift | < 0.1 | Lighthouse |
| First Input Delay | < 100ms | RUM |

### Performance Best Practices

| Practice | Implementation |
|---|---|
| Route-level splitting | Every feature loaded via React.lazy() |
| Component-level splitting | Heavy components (charts, PDF viewer) lazy-loaded |
| Image optimization | WebP format; lazy loading; responsive srcset |
| Memoization | React.memo for list items; useMemo for expensive computations |
| Virtualization | react-window for long lists (transaction history > 50 items) |
| Prefetching | Prefetch next likely route on hover/visibility |
| Bundle analysis | Run on every PR; alert on > 10% size increase |
| Tree shaking | Named imports only (import { format } from 'date-fns', not import * as dateFns) |
| No barrel re-exports of large modules | Import directly from feature, not from shared/index.ts |
| Web Workers | Offload heavy computation (encryption, data transformation) |

### Rendering Optimization

```tsx
// Virtualized list for transaction history
import { FixedSizeList } from 'react-window';

export function TransactionList({ transactions }: { transactions: Transaction[] }) {
  const Row = memo(({ index, style }: { index: number; style: CSSProperties }) => (
    <div style={style}>
      <TransactionRow transaction={transactions[index]} />
    </div>
  ));

  return (
    <FixedSizeList
      height={600}
      itemCount={transactions.length}
      itemSize={72}
      width="100%"
    >
      {Row}
    </FixedSizeList>
  );
}
```

---

## 11. Testing Standards

### Test Pyramid

| Level | Coverage Target | Tool | What to Test |
|---|---|---|---|
| Unit | ≥ 80% overall; ≥ 90% utils/hooks | Jest + RTL | Components, hooks, utils, formatters, validators |
| Integration | All API interactions | MSW + RTL | Feature flows with mocked API |
| E2E | Critical paths | Cypress / Playwright | Login → transfer → confirmation (happy + error) |
| Accessibility | All components | axe-core + manual | WCAG 2.1 AA compliance |
| Visual regression | Design system components | Chromatic / Percy | UI consistency |

### Testing Rules

| Rule | Standard |
|---|---|
| Test behavior, not implementation | Test what user sees/does, not internal state |
| No snapshot tests for logic | Snapshots only for design system components |
| Mock at network level | Use MSW (Mock Service Worker), not mock axios |
| Test accessibility | Include axe checks in component tests |
| Test error states | Every component test includes error/loading/empty states |
| No production data | Use fixtures with synthetic data |
| Test security controls | Test that sensitive data is masked, auth redirects work, XSS is prevented |

### Component Test Pattern

```tsx
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { axe } from 'jest-axe';

describe('TransferForm', () => {
  it('submits valid transfer', async () => {
    render(<TransferForm sourceAccountId="ACC-001" onSubmit={mockSubmit} onCancel={mockCancel} />);
    
    await userEvent.selectOptions(screen.getByLabelText('Beneficiary'), 'BEN-001');
    await userEvent.type(screen.getByLabelText('Amount'), '5000');
    await userEvent.click(screen.getByRole('button', { name: /transfer/i }));
    
    await waitFor(() => expect(mockSubmit).toHaveBeenCalledWith(
      expect.objectContaining({ amount: 5000, beneficiaryId: 'BEN-001' })
    ));
  });

  it('shows validation error for amount exceeding limit', async () => {
    render(<TransferForm sourceAccountId="ACC-001" onSubmit={mockSubmit} onCancel={mockCancel} />);
    
    await userEvent.type(screen.getByLabelText('Amount'), '999999999999');
    await userEvent.click(screen.getByRole('button', { name: /transfer/i }));
    
    expect(screen.getByText(/exceeds.*limit/i)).toBeInTheDocument();
    expect(mockSubmit).not.toHaveBeenCalled();
  });

  it('has no accessibility violations', async () => {
    const { container } = render(
      <TransferForm sourceAccountId="ACC-001" onSubmit={mockSubmit} onCancel={mockCancel} />
    );
    expect(await axe(container)).toHaveNoViolations();
  });
});
```

---

## 12. CI/CD Pipeline Standards

### Pipeline Stages

```
1. Install → npm ci (clean install from lock file)
2. Lint → eslint + prettier check
3. Type check → tsc --noEmit
4. Unit tests → jest --coverage (fail if < 80%)
5. Build → production build
6. Bundle analysis → size check (fail if > threshold)
7. Accessibility scan → axe-core on built components
8. Security scan → npm audit + Snyk
9. Integration tests → MSW-based tests
10. Deploy to staging → preview deployment
11. E2E tests → Cypress/Playwright on staging
12. Deploy to production → with feature flags
```

### CI/CD Gates

| Gate | Blocks | Threshold |
|---|---|---|
| Lint errors | Merge | Zero errors |
| Type errors | Merge | Zero errors |
| Test coverage | Merge | < 80% overall |
| Bundle size | Merge | > 200KB initial (gzipped) |
| npm audit critical | Merge | Any critical CVE |
| Accessibility violations | Merge | Any critical violation |
| E2E failures | Deploy | Any failure on critical paths |

---

## 13. Common Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Class components | Inconsistent patterns; no hooks | Functional components only (DV-BP-001) |
| No error boundaries | One crash takes down entire app | Error boundary per feature (DV-BP-002) |
| dangerouslySetInnerHTML without sanitization | XSS vulnerability | DOMPurify or avoid entirely (DV-BP-003) |
| Hardcoded API URLs | Breaks across environments | Environment variables (DV-BP-004) |
| No CSP | XSS, data injection attacks | Strict CSP header (DV-BP-005) |
| Outdated dependencies with CVEs | Security vulnerabilities | Automated scanning + patching (DV-BP-006) |
| Everything in Redux | Unnecessary complexity; stale data | React Query for server state (DV-BP-007) |
| No code splitting | Massive initial bundle; slow load | Lazy loading per route (DV-BP-008) |
| Sensitive data in localStorage | Data breach via XSS | In-memory only; HTTP-only cookies |
| console.log in production | Information leakage | Strip in build; ESLint rule |
| any type everywhere | No type safety; runtime errors | Strict TypeScript; no any |
| No accessibility testing | Excludes users; legal risk | axe-core in tests + CI |
