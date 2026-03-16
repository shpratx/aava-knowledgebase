# Banking Design System — Templates, Pages & Governance
---

## 5. Templates

Templates are page-level layouts that define content structure using placeholders. They show where organisms and molecules are placed without real content.

### 5.1 Dashboard Template

```
┌──────────────────────────────────────────────────────────────┐
│ [Header Organism]                                            │
├────────────┬─────────────────────────────────────────────────┤
│            │                                                 │
│ [Sidebar   │  [Account Summary Organism]                     │
│  Navigation│                                                 │
│  Organism] │  ┌─────────────────┐ ┌────────────────────────┐ │
│            │  │ [Quick Actions  │ │ [Transaction List      │ │
│            │  │  Organism]      │ │  Organism]             │ │
│            │  │                 │ │                        │ │
│            │  │ • Transfer      │ │  {transaction rows}    │ │
│            │  │ • Pay Bill      │ │  {pagination}          │ │
│            │  │ • Add Benefic.  │ │                        │ │
│            │  └─────────────────┘ └────────────────────────┘ │
│            │                                                 │
│            │  [Notification Panel Organism] (collapsed)       │
│            │                                                 │
├────────────┴─────────────────────────────────────────────────┤
│ [Footer Organism]                                            │
└──────────────────────────────────────────────────────────────┘
```

**Layout:** Sidebar (240px fixed) + Main content (fluid). Sidebar collapses to bottom nav on mobile.

**Content Structure:**
- Account summary: 1 account selector + 3 stat cards
- Quick actions: 3-5 action buttons
- Transaction list: 10 rows default, paginated
- Notifications: collapsed by default, expandable

### 5.2 Transfer Template

```
┌──────────────────────────────────────────────────────────────┐
│ [Header Organism]                                            │
├────────────┬─────────────────────────────────────────────────┤
│            │                                                 │
│ [Sidebar]  │  [Breadcrumb: Dashboard > Transfers > New]      │
│            │                                                 │
│            │  [Page Title: "New Transfer"]                    │
│            │                                                 │
│            │  [Progress Steps: Details → Confirm → Verify]   │
│            │                                                 │
│            │  ┌────────────────────────────────────────────┐ │
│            │  │ [Transfer Form Organism]                   │ │
│            │  │                                            │ │
│            │  │ {current step content}                     │ │
│            │  │                                            │ │
│            │  │ [Cancel]              [Continue / Confirm] │ │
│            │  └────────────────────────────────────────────┘ │
│            │                                                 │
│            │  [Alert Molecule] (conditional: success/error)   │
│            │                                                 │
├────────────┴─────────────────────────────────────────────────┤
│ [Footer Organism]                                            │
└──────────────────────────────────────────────────────────────┘
```

**Content Structure:**
- 3-step wizard (Details → Confirm → Verify)
- Form fields: account selector, beneficiary selector, amount, reference
- Confirmation: read-only summary of all fields
- Verification: OTP input with timer

### 5.3 Account Detail Template

```
┌──────────────────────────────────────────────────────────────┐
│ [Header Organism]                                            │
├────────────┬─────────────────────────────────────────────────┤
│            │                                                 │
│ [Sidebar]  │  [Breadcrumb: Dashboard > Accounts > ****1234]  │
│            │                                                 │
│            │  [Account Summary Organism]                      │
│            │  {account type, number (masked), status badge}   │
│            │  {available balance, ledger balance, hold}       │
│            │                                                 │
│            │  [Tab Navigation: Transactions | Statements |    │
│            │                   Details | Settings]            │
│            │                                                 │
│            │  ┌────────────────────────────────────────────┐ │
│            │  │ {Active tab content}                       │ │
│            │  │                                            │ │
│            │  │ Transactions: [Transaction List Organism]  │ │
│            │  │ Statements: [Statement List + Download]    │ │
│            │  │ Details: [Account info read-only]          │ │
│            │  │ Settings: [Alerts, Limits, Preferences]    │ │
│            │  └────────────────────────────────────────────┘ │
│            │                                                 │
├────────────┴─────────────────────────────────────────────────┤
│ [Footer Organism]                                            │
└──────────────────────────────────────────────────────────────┘
```

### 5.4 Settings Template

```
┌──────────────────────────────────────────────────────────────┐
│ [Header Organism]                                            │
├────────────┬─────────────────────────────────────────────────┤
│            │                                                 │
│ [Sidebar]  │  [Page Title: "Settings"]                       │
│            │                                                 │
│            │  [Settings Nav: Profile | Security | Privacy |   │
│            │                 Notifications | Preferences]     │
│            │                                                 │
│            │  ┌────────────────────────────────────────────┐ │
│            │  │ Profile: Name, email, phone, address       │ │
│            │  │ Security: Password, MFA, sessions, devices │ │
│            │  │ Privacy: Consents (toggles), data export,  │ │
│            │  │          erasure request                    │ │
│            │  │ Notifications: Email/SMS/Push toggles       │ │
│            │  │ Preferences: Language, theme, currency      │ │
│            │  └────────────────────────────────────────────┘ │
│            │                                                 │
├────────────┴─────────────────────────────────────────────────┤
│ [Footer Organism]                                            │
└──────────────────────────────────────────────────────────────┘
```

### 5.5 Login Template

```
┌──────────────────────────────────────────────────────────────┐
│                     [Logo]                                    │
│                                                              │
│              ┌──────────────────────┐                        │
│              │  Welcome Back        │                        │
│              │                      │                        │
│              │  Email / User ID     │                        │
│              │  [________________]  │                        │
│              │                      │                        │
│              │  Password            │                        │
│              │  [________________]  │                        │
│              │                      │                        │
│              │  [Forgot Password?]  │                        │
│              │                      │                        │
│              │  [    Sign In     ]  │                        │
│              │                      │                        │
│              │  {MFA step - conditional}                     │
│              └──────────────────────┘                        │
│                                                              │
│  [Footer: Privacy | Terms | Security | Contact]              │
└──────────────────────────────────────────────────────────────┘
```

### 5.6 Mobile Templates

All templates must have mobile variants:

| Desktop Element | Mobile Adaptation |
|---|---|
| Sidebar navigation | Bottom tab bar (5 items max) + hamburger for overflow |
| Multi-column layout | Single column, stacked |
| Data tables | Card-based list view |
| Horizontal tabs | Scrollable tabs or dropdown |
| Side-by-side stat cards | Stacked or horizontal scroll |
| Modal dialogs | Full-screen sheets (bottom sheet on mobile) |

---

## 6. Pages

Pages are specific instances of templates with real representative content. They validate that the design system works with actual data.

### 6.1 Page Variations to Test

| Page | Variations |
|---|---|
| Dashboard | New user (no transactions, empty state); Active user (many transactions); User with pending transfers; User with alerts/notifications |
| Transfer | Small amount ($10); Large amount ($49,999); Limit exceeded; Insufficient funds; Fraud hold; Success; Failure |
| Account Detail | Active account; Frozen account; Dormant account; Account with 0 transactions; Account with 1000+ transactions |
| Beneficiary List | No beneficiaries (empty state); 1 beneficiary; 50+ beneficiaries (pagination); Search with no results |
| Settings - Privacy | All consents granted; No consents; Mixed consents |
| Login | First login; Returning user; MFA required; Account locked; Password expired |
| Error | 404 page; 500 page; Maintenance page; Session expired |

### 6.2 Content Guidelines for Pages

| Content Type | Guideline |
|---|---|
| Account numbers | Always masked: ****1234 (last 4 digits) |
| Amounts | Locale-formatted with currency symbol; monospace font; debit red, credit green |
| Dates | User locale format; relative for recent ("2 hours ago"); absolute for older |
| Names | Full name for beneficiaries; first name + last initial for privacy in lists |
| Status | Badge with color + icon + text (never color alone) |
| Empty states | Illustration + heading + description + primary action |
| Loading states | Skeleton screens matching content layout; never blank page |
| Error states | User-friendly message + action; never technical details |

---

## 7. Responsive Design Standards

### Breakpoint Behavior

| Breakpoint | Layout | Navigation | Content |
|---|---|---|---|
| Mobile (< 768px) | Single column | Bottom tab bar | Cards stacked; tables → card list |
| Tablet (768-1024px) | Collapsed sidebar + main | Collapsible sidebar | 2-column where appropriate |
| Desktop (> 1024px) | Sidebar + main | Full sidebar | Multi-column; data tables |

### Touch Target Standards

| Element | Minimum Size | Spacing |
|---|---|---|
| Buttons | 44 × 44px | 8px between targets |
| Links (inline) | 44px height | 8px vertical spacing |
| Checkboxes/Radio | 44 × 44px (including label) | 12px between options |
| List items (tappable) | 48px height | No gap (full-width tap) |
| Icons (interactive) | 44 × 44px tap area | 8px between icons |

---

## 8. Accessibility Standards

### WCAG 2.1 AA Compliance Checklist

| Category | Requirement | Implementation |
|---|---|---|
| Perceivable | Color contrast ≥ 4.5:1 (text), ≥ 3:1 (large text, UI) | Design tokens enforce; automated testing |
| Perceivable | No color-only information | Status: color + icon + text |
| Perceivable | All images have alt text | Decorative: alt=""; meaningful: descriptive alt |
| Operable | All interactive elements keyboard accessible | Tab order; Enter/Space activation; arrow keys for groups |
| Operable | Visible focus indicators (≥ 2px) | --border-focus token; never outline: none |
| Operable | No keyboard traps | Modals: focus trapped but ESC exits |
| Operable | Session timeout warning with extend option | 2-minute warning; accessible modal |
| Understandable | Form labels on all inputs | label + for/id; aria-describedby for help |
| Understandable | Error identification and suggestion | Inline errors linked to fields; aria-invalid |
| Understandable | Consistent navigation | Same nav position on all pages |
| Robust | Valid semantic HTML | Automated HTML validation |
| Robust | ARIA used correctly | Tested with screen readers |

### Screen Reader Announcements

| Event | Announcement | ARIA |
|---|---|---|
| Page load | Page title | document.title |
| Form error | "Error: [field] [message]" | aria-live="assertive" |
| Transfer success | "Transfer completed successfully" | aria-live="polite" |
| Loading complete | "Content loaded" | aria-busy="false" |
| Session warning | "Session expires in [time]" | aria-live="assertive" |
| Notification | "[count] new notifications" | aria-live="polite" |

---

## 9. Design System Governance

### 9.1 Contribution Process

| Step | Activity | Owner |
|---|---|---|
| 1 | Propose new component or change (issue/RFC) | Any developer/designer |
| 2 | Review against existing components (avoid duplication) | Design system team |
| 3 | Design with tokens, accessibility, responsive specs | Designer |
| 4 | Implement with tests (unit, accessibility, visual regression) | Developer |
| 5 | Review: code, design, accessibility, security | Design system team |
| 6 | Document: props, variants, states, usage guidelines, do/don't | Author |
| 7 | Publish: version bump, changelog, migration guide if breaking | Design system team |

### 9.2 Versioning

| Change Type | Version Bump | Example |
|---|---|---|
| New component | Minor (1.x.0) | Add OTP Input molecule |
| New variant/prop | Minor (1.x.0) | Add "ghost" button variant |
| Bug fix | Patch (1.0.x) | Fix focus indicator on input |
| Breaking change | Major (x.0.0) | Rename prop; remove variant |
| Token change | Minor or Major | Minor if additive; Major if existing token changes |

### 9.3 Component Documentation Template

Every component must be documented with:

| Section | Content |
|---|---|
| Name | Component name and atomic level (atom/molecule/organism) |
| Description | What it does and when to use it |
| Props/API | All props with types, defaults, and descriptions |
| Variants | All visual variants with examples |
| States | All interactive states (default, hover, focus, active, disabled, loading, error) |
| Accessibility | ARIA attributes, keyboard behavior, screen reader behavior |
| Responsive | Behavior at each breakpoint |
| Security | Any security considerations (masking, autocomplete, sensitive data) |
| Do / Don't | Usage guidelines with examples |
| Related | Related components and when to use which |
| Changelog | Version history for this component |

### 9.4 Quality Gates

| Gate | Requirement | Enforcement |
|---|---|---|
| Design review | Matches design tokens; consistent with system | Design team approval |
| Accessibility | axe-core zero violations; keyboard tested; screen reader tested | CI/CD gate |
| Visual regression | No unintended visual changes | Chromatic/Percy approval |
| Unit tests | All props, variants, states tested | CI/CD gate (≥ 90% coverage) |
| Documentation | All sections of component doc template filled | Review checklist |
| Security | No sensitive data exposure; proper autocomplete; masking | Security review |
| Performance | No layout shift; lazy loaded if heavy; bundle impact assessed | CI/CD gate |
| Browser testing | Latest 2 versions of Chrome, Firefox, Safari, Edge | Cross-browser test |
| Mobile testing | iOS Safari, Android Chrome | Device testing |

### 9.5 Design System Health Metrics

| Metric | Target | Measurement |
|---|---|---|
| Component adoption | > 90% of UI built with design system components | Code analysis |
| Accessibility score | 100% WCAG 2.1 AA | Automated + manual audit |
| Visual consistency | < 5% custom overrides | CSS analysis |
| Component coverage | All UI patterns have a design system component | Pattern audit |
| Documentation completeness | 100% of components fully documented | Documentation audit |
| Stale components | 0 deprecated components without replacement | Quarterly review |

---

## 10. Component Inventory Summary

| Level | Count | Components |
|---|---|---|
| **Atoms** | 10 | Button, Input, Label, Typography, Icon, Badge, Avatar, Divider, Skeleton, Spinner |
| **Molecules** | 12 | Form Field, Amount Input, Account Selector, Search Bar, Nav Item, Alert, Transaction Row, Stat Card, Empty State, Confirmation Step, OTP Input, Pagination |
| **Organisms** | 10 | Header, Sidebar Nav, Account Summary, Transaction List, Transfer Form, Beneficiary List, Session Timeout Warning, Notification Panel, Footer, Error Page |
| **Templates** | 6 | Dashboard, Transfer, Account Detail, Settings, Login, Mobile variants |
| **Pages** | 7+ | Dashboard (4 variants), Transfer (6 variants), Account (4 variants), Beneficiaries (4 variants), Settings, Login (4 variants), Error (3 variants) |
| **Total** | 45+ | Baseline components for banking UI |

---

## 11. Evolution Roadmap

| Phase | Components to Add | Priority |
|---|---|---|
| Phase 1 (Baseline) | All atoms + core molecules + core organisms + primary templates | Now |
| Phase 2 (Expansion) | Card management organisms, Loan application templates, Statement viewer | Next quarter |
| Phase 3 (Advanced) | Data visualization (charts, graphs), Onboarding wizard, Biometric auth UI | Following quarter |
| Phase 4 (Optimization) | Dark mode, High-contrast mode, Animation library, Micro-interactions | Ongoing |
| Continuous | Accessibility improvements, Performance optimization, New banking features | Always |
