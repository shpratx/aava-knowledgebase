# React Application Standards & Best Practices — Part 1

---

## 1. Project Structure

### Standard Banking React Project Layout

```
src/
├── app/                          # Application shell
│   ├── App.tsx                   # Root component, providers, router
│   ├── routes.tsx                # Route definitions (lazy-loaded)
│   └── providers/                # Global providers (auth, theme, i18n)
│       ├── AuthProvider.tsx
│       ├── ThemeProvider.tsx
│       └── ErrorBoundaryProvider.tsx
│
├── features/                     # Feature-based modules (domain-driven)
│   ├── accounts/
│   │   ├── components/           # Feature-specific components
│   │   │   ├── AccountList.tsx
│   │   │   ├── AccountBalance.tsx
│   │   │   └── TransactionHistory.tsx
│   │   ├── hooks/                # Feature-specific hooks
│   │   │   ├── useAccounts.ts
│   │   │   └── useTransactions.ts
│   │   ├── services/             # API calls for this feature
│   │   │   └── accountApi.ts
│   │   ├── types/                # TypeScript types for this feature
│   │   │   └── account.types.ts
│   │   ├── utils/                # Feature-specific utilities
│   │   └── index.ts              # Public API (barrel export)
│   │
│   ├── transfers/
│   │   ├── components/
│   │   │   ├── TransferForm.tsx
│   │   │   ├── TransferConfirmation.tsx
│   │   │   ├── MfaStep.tsx
│   │   │   └── TransferStatus.tsx
│   │   ├── hooks/
│   │   │   ├── useTransfer.ts
│   │   │   └── useTransferValidation.ts
│   │   ├── services/
│   │   │   └── transferApi.ts
│   │   ├── types/
│   │   └── index.ts
│   │
│   ├── beneficiaries/
│   ├── cards/                    # PCI-DSS scoped — isolated module
│   ├── settings/                 # Profile, preferences, consent management
│   └── auth/                     # Login, MFA, session management
│
├── shared/                       # Shared across features
│   ├── components/               # Reusable UI components
│   │   ├── ui/                   # Design system primitives
│   │   │   ├── Button.tsx
│   │   │   ├── Input.tsx
│   │   │   ├── Modal.tsx
│   │   │   ├── Table.tsx
│   │   │   └── Alert.tsx
│   │   ├── layout/               # Layout components
│   │   │   ├── Header.tsx
│   │   │   ├── Footer.tsx
│   │   │   ├── Sidebar.tsx
│   │   │   └── PageLayout.tsx
│   │   ├── feedback/             # Loading, error, empty states
│   │   │   ├── LoadingSkeleton.tsx
│   │   │   ├── ErrorFallback.tsx
│   │   │   └── EmptyState.tsx
│   │   └── forms/                # Form primitives
│   │       ├── FormField.tsx
│   │       ├── AmountInput.tsx
│   │       ├── AccountSelector.tsx
│   │       └── OtpInput.tsx
│   │
│   ├── hooks/                    # Shared hooks
│   │   ├── useAuth.ts
│   │   ├── useSession.ts
│   │   ├── useApi.ts
│   │   └── useAccessibility.ts
│   │
│   ├── services/                 # Shared API infrastructure
│   │   ├── apiClient.ts          # Axios/fetch wrapper with interceptors
│   │   ├── authService.ts        # Token management
│   │   └── auditService.ts       # Client-side audit event emission
│   │
│   ├── utils/                    # Shared utilities
│   │   ├── formatters.ts         # Currency, date, account number formatting
│   │   ├── validators.ts         # Client-side validation helpers
│   │   ├── sanitizers.ts         # Input sanitization (DOMPurify)
│   │   ├── maskUtils.ts          # Account/card number masking
│   │   └── constants.ts          # App-wide constants
│   │
│   ├── types/                    # Shared TypeScript types
│   │   ├── api.types.ts          # API response/error types
│   │   ├── auth.types.ts         # Auth/session types
│   │   └── common.types.ts       # Shared domain types (Money, Currency)
│   │
│   └── config/                   # Configuration
│       ├── env.ts                # Environment variable access (validated)
│       ├── routes.ts             # Route path constants
│       └── featureFlags.ts       # Feature flag access
│
├── styles/                       # Global styles and design tokens
│   ├── tokens/                   # Design tokens (colors, spacing, typography)
│   ├── global.css                # Global styles, CSS reset
│   └── themes/                   # Theme definitions (light, dark, high-contrast)
│
├── test/                         # Test utilities and setup
│   ├── setup.ts                  # Test environment setup
│   ├── mocks/                    # API mocks (MSW handlers)
│   ├── fixtures/                 # Test data fixtures
│   └── utils/                    # Test helpers (render with providers)
│
└── index.tsx                     # Entry point
```

### Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Components | PascalCase | TransferForm.tsx, AccountBalance.tsx |
| Hooks | camelCase with use prefix | useTransfer.ts, useAuth.ts |
| Services | camelCase with descriptive suffix | transferApi.ts, authService.ts |
| Types | PascalCase with .types.ts suffix | account.types.ts |
| Utils | camelCase | formatters.ts, validators.ts |
| Constants | UPPER_SNAKE_CASE | MAX_TRANSFER_AMOUNT, SESSION_TIMEOUT_MS |
| CSS modules | camelCase | transferForm.module.css |
| Test files | Same name + .test.tsx | TransferForm.test.tsx |
| Barrel exports | index.ts per feature | features/transfers/index.ts |

---

## 2. Functional Components & Hooks (ref: DV-BP-001)

### Component Standards

```tsx
// Standard functional component pattern
import { memo } from 'react';
import type { FC } from 'react';

interface TransferFormProps {
  sourceAccountId: string;
  onSubmit: (transfer: TransferRequest) => void;
  onCancel: () => void;
}

export const TransferForm: FC<TransferFormProps> = memo(({ 
  sourceAccountId, 
  onSubmit, 
  onCancel 
}) => {
  // hooks at top level only
  const { accounts } = useAccounts();
  const { validate, errors } = useTransferValidation();
  
  // event handlers
  const handleSubmit = useCallback((e: FormEvent) => {
    e.preventDefault();
    // ...
  }, [/* deps */]);

  return (
    <form onSubmit={handleSubmit} noValidate>
      {/* ... */}
    </form>
  );
});

TransferForm.displayName = 'TransferForm';
```

### Component Rules

| Rule | Standard |
|---|---|
| Functional only | No class components — use functional components with hooks (ref: DV-BP-001) |
| TypeScript | All components must be typed — props interface, return type inferred |
| Props interface | Explicit interface for every component's props — no `any` |
| Default exports | Avoid — use named exports for better refactoring and tree-shaking |
| Memoization | Use `memo()` for components receiving complex props; `useMemo`/`useCallback` for expensive computations |
| Hook rules | Hooks at top level only; never inside conditions, loops, or nested functions |
| Single responsibility | One component = one responsibility; extract sub-components when > 150 lines |
| Display name | Set `displayName` on memo'd components for DevTools debugging |

### Custom Hook Standards

```tsx
// Custom hook pattern — encapsulate business logic
export function useTransfer(transferId: string) {
  const [transfer, setTransfer] = useState<Transfer | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<ApiError | null>(null);

  useEffect(() => {
    let cancelled = false;
    
    async function fetchTransfer() {
      try {
        setLoading(true);
        const data = await transferApi.getTransfer(transferId);
        if (!cancelled) setTransfer(data);
      } catch (err) {
        if (!cancelled) setError(toApiError(err));
      } finally {
        if (!cancelled) setLoading(false);
      }
    }
    
    fetchTransfer();
    return () => { cancelled = true; };
  }, [transferId]);

  return { transfer, loading, error };
}
```

---

## 3. Error Boundaries (ref: DV-BP-002)

### Error Boundary Architecture

```
App
├── GlobalErrorBoundary (catches unhandled errors → generic error page)
│   ├── AuthErrorBoundary (catches auth errors → redirect to login)
│   │   ├── FeatureErrorBoundary (per feature → feature-level fallback)
│   │   │   ├── AccountsFeature
│   │   │   ├── TransfersFeature
│   │   │   └── CardsFeature
│   │   └── FeatureErrorBoundary
│   └── AuthErrorBoundary
└── GlobalErrorBoundary
```

### Implementation

```tsx
// Reusable error boundary with fallback
import { Component, ErrorInfo, ReactNode } from 'react';

interface ErrorBoundaryProps {
  fallback: ReactNode | ((error: Error, reset: () => void) => ReactNode);
  onError?: (error: Error, errorInfo: ErrorInfo) => void;
  children: ReactNode;
}

export class ErrorBoundary extends Component<ErrorBoundaryProps, { error: Error | null }> {
  state = { error: null as Error | null };

  static getDerivedStateFromError(error: Error) {
    return { error };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // Log to monitoring service — never expose to user
    this.props.onError?.(error, errorInfo);
    auditService.logError({
      error: error.message,
      component: errorInfo.componentStack,
      timestamp: new Date().toISOString(),
    });
  }

  reset = () => this.setState({ error: null });

  render() {
    if (this.state.error) {
      const { fallback } = this.props;
      return typeof fallback === 'function' 
        ? fallback(this.state.error, this.reset) 
        : fallback;
    }
    return this.props.children;
  }
}
```

### Error Boundary Rules

| Rule | Standard |
|---|---|
| Global boundary | Wrap entire app — catch unhandled errors |
| Feature boundaries | Wrap each feature/route — isolate failures |
| Fallback UI | User-friendly message; retry button; no technical details |
| Error logging | Log to monitoring (Datadog/Sentry) — never expose stack traces to UI |
| Recovery | Provide "Try Again" or "Go to Dashboard" — don't leave user stuck |
| Sensitive data | Error fallback must not display any sensitive data from state |

---

## 4. Input Sanitization & XSS Prevention (ref: DV-BP-003)

### Sanitization Standards

| Layer | What | How |
|---|---|---|
| Input sanitization | Clean user input before processing | DOMPurify for rich text; escape for plain text |
| Output encoding | Encode before rendering | React's JSX auto-escapes by default; never use dangerouslySetInnerHTML without DOMPurify |
| URL sanitization | Validate URLs before navigation | Whitelist allowed protocols (https:); reject javascript: |
| API response | Treat all API data as untrusted | Sanitize before rendering; validate types |

### Implementation

```tsx
// Input sanitization utility
import DOMPurify from 'dompurify';

export const sanitize = {
  // For plain text inputs (transfer reference, search)
  text: (input: string): string => {
    return input.replace(/[<>&"']/g, (char) => ({
      '<': '&lt;', '>': '&gt;', '&': '&amp;', '"': '&quot;', "'": '&#x27;'
    }[char] || char));
  },
  
  // For rich text (if ever needed — avoid in banking)
  html: (input: string): string => {
    return DOMPurify.sanitize(input, { ALLOWED_TAGS: ['b', 'i', 'em', 'strong'] });
  },
  
  // For URLs
  url: (input: string): string | null => {
    try {
      const url = new URL(input);
      return ['https:', 'http:'].includes(url.protocol) ? url.toString() : null;
    } catch {
      return null;
    }
  }
};

// NEVER do this:
// <div dangerouslySetInnerHTML={{ __html: userInput }} />

// If absolutely necessary (rare):
// <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(content) }} />
```

### XSS Prevention Rules

| Rule | Standard |
|---|---|
| React JSX | Rely on React's auto-escaping — it handles most cases |
| dangerouslySetInnerHTML | Prohibited unless sanitized with DOMPurify AND approved by security review |
| href/src attributes | Validate URLs — reject javascript: protocol |
| eval() / Function() | Prohibited — never execute dynamic code |
| innerHTML | Prohibited — use React's rendering |
| Template literals in DOM | Never interpolate user input into HTML strings |
| Third-party HTML | Always sanitize with DOMPurify before rendering |
| Markdown rendering | Use a sanitizing markdown renderer (react-markdown with rehype-sanitize) |

---

## 5. Environment Variables & Configuration (ref: DV-BP-004)

### Configuration Standards

```tsx
// src/shared/config/env.ts — validated environment access
const requiredEnvVars = [
  'REACT_APP_API_BASE_URL',
  'REACT_APP_AUTH_URL',
  'REACT_APP_ENV',
] as const;

function getEnv(key: string): string {
  const value = process.env[key];
  if (!value) throw new Error(`Missing required env var: ${key}`);
  return value;
}

export const config = Object.freeze({
  apiBaseUrl: getEnv('REACT_APP_API_BASE_URL'),
  authUrl: getEnv('REACT_APP_AUTH_URL'),
  environment: getEnv('REACT_APP_ENV') as 'development' | 'staging' | 'production',
  
  // Feature flags (from remote config, not env vars)
  featureFlags: {
    enableInternationalTransfers: false, // overridden by remote config
    enableBiometricAuth: false,
  },
  
  // Session configuration
  session: {
    inactivityTimeoutMs: 15 * 60 * 1000,  // 15 minutes
    absoluteTimeoutMs: 8 * 60 * 60 * 1000, // 8 hours
    warningBeforeTimeoutMs: 2 * 60 * 1000,  // 2 minutes
  },
  
  // Rate limiting (client-side UX throttling)
  rateLimits: {
    transferSubmitCooldownMs: 5000,
    otpResendCooldownMs: 30000,
  },
});
```

### Configuration Rules

| Rule | Standard |
|---|---|
| No hardcoded values | API URLs, auth endpoints, feature flags — all via env vars or remote config |
| No secrets in env vars | API keys, tokens, secrets must NEVER be in client-side env vars |
| Validated on startup | Missing required env vars fail fast at app initialization |
| Typed config | Config object is typed and frozen — no runtime mutation |
| Environment-specific | Different values per environment (dev, staging, production) |
| .env files | .env.local in .gitignore; .env.example committed with placeholder values |
| Build-time injection | REACT_APP_ prefix for Create React App; VITE_ for Vite |

---

## 6. Coding Guidelines

### TypeScript Standards

| Rule | Standard |
|---|---|
| Strict mode | tsconfig: strict: true, noImplicitAny: true, strictNullChecks: true |
| No any | Prohibited — use unknown + type guards, or specific types |
| Interface over type | Prefer interface for object shapes; type for unions/intersections |
| Enums | Use const enums or string literal unions — avoid numeric enums |
| Null handling | Use optional chaining (?.) and nullish coalescing (??); no non-null assertions (!) |
| Return types | Explicit return types on exported functions; inferred for internal |
| Generics | Use for reusable utilities; constrain with extends |

### Code Quality Rules

| Rule | Standard | Enforcement |
|---|---|---|
| Linting | ESLint with banking-specific rules | CI/CD gate — block on errors |
| Formatting | Prettier with consistent config | Pre-commit hook |
| Max file length | 300 lines per file | ESLint warning at 200, error at 300 |
| Max function length | 50 lines | ESLint warning at 30, error at 50 |
| Cyclomatic complexity | Max 10 | ESLint error |
| Import order | External → shared → feature → relative | ESLint auto-fix |
| No console.log | Prohibited in production builds | ESLint error; build strips console |
| No TODO/FIXME | Must be tracked as tasks; not in production code | ESLint warning; CI report |
| Comments | JSDoc for public APIs; inline for complex logic only | Code review |

### API Client Standards

```tsx
// src/shared/services/apiClient.ts
import axios, { AxiosInstance, InternalAxiosRequestConfig } from 'axios';
import { config } from '../config/env';
import { authService } from './authService';

const apiClient: AxiosInstance = axios.create({
  baseURL: config.apiBaseUrl,
  timeout: 30000,
  headers: { 'Content-Type': 'application/json' },
  withCredentials: true, // for HTTP-only cookies
});

// Request interceptor: attach auth + correlation ID
apiClient.interceptors.request.use((req: InternalAxiosRequestConfig) => {
  const token = authService.getAccessToken();
  if (token) req.headers.Authorization = `Bearer ${token}`;
  req.headers['X-Correlation-ID'] = crypto.randomUUID();
  req.headers['X-Request-ID'] = crypto.randomUUID();
  return req;
});

// Response interceptor: handle auth errors, sanitize errors
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Attempt token refresh; if fails, redirect to login
      const refreshed = await authService.refreshToken();
      if (refreshed) return apiClient.request(error.config);
      authService.logout();
    }
    // Never expose raw API error to UI
    throw toApiError(error);
  }
);

// Sanitize API errors — never expose internal details
function toApiError(error: unknown): ApiError {
  if (axios.isAxiosError(error) && error.response) {
    return {
      code: error.response.data?.error?.code || 'UNKNOWN_ERROR',
      message: error.response.data?.error?.message || 'Something went wrong',
      status: error.response.status,
      details: error.response.data?.error?.details,
    };
  }
  return { code: 'NETWORK_ERROR', message: 'Unable to connect. Please try again.', status: 0 };
}

export { apiClient };
```

### Formatting Utilities

```tsx
// src/shared/utils/formatters.ts
export const format = {
  // Currency — locale-aware, never floating point display issues
  currency: (amount: number, currency: string, locale = 'en-US'): string => {
    return new Intl.NumberFormat(locale, {
      style: 'currency',
      currency,
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    }).format(amount);
  },

  // Account number masking — show last 4 only
  accountNumber: (account: string): string => {
    return `****${account.slice(-4)}`;
  },

  // Date — always from UTC, display in user locale
  date: (isoDate: string, locale = 'en-US'): string => {
    return new Intl.DateTimeFormat(locale, {
      year: 'numeric', month: 'short', day: 'numeric',
    }).format(new Date(isoDate));
  },

  // Date-time with timezone
  dateTime: (isoDate: string, locale = 'en-US'): string => {
    return new Intl.DateTimeFormat(locale, {
      year: 'numeric', month: 'short', day: 'numeric',
      hour: '2-digit', minute: '2-digit', timeZoneName: 'short',
    }).format(new Date(isoDate));
  },
};
```
