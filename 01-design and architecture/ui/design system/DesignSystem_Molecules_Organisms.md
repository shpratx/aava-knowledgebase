# Banking Design System — Molecules & Organisms
---

## 3. Molecules

Molecules are simple groups of atoms functioning together as a unit. Each molecule has a single responsibility.

### 3.1 Form Field Molecule

**Composition:** Label atom + Input atom + Help text atom + Error text atom

```
┌─────────────────────────────────┐
│ Amount *                        │  ← Label (required indicator)
│ ┌─────────────────────────────┐ │
│ │ $ 5,000.00                  │ │  ← Input (number, currency formatted)
│ └─────────────────────────────┘ │
│ Daily limit: $50,000.00         │  ← Help text
│ ⚠ Exceeds daily limit          │  ← Error text (conditional)
└─────────────────────────────────┘
```

| Prop | Type | Description |
|---|---|---|
| label | string | Field label text |
| name | string | Input name/id |
| type | InputType | text, number, password, email, tel, date |
| required | boolean | Shows required indicator |
| helpText | string | Guidance below input |
| error | string | Error message (shows error state) |
| disabled | boolean | Disables interaction |
| value | string | Controlled value |
| onChange | function | Change handler |

### 3.2 Amount Input Molecule

**Composition:** Label atom + Input atom (number) + Currency selector atom

```
┌─────────────────────────────────┐
│ Transfer Amount *               │
│ ┌──────┐ ┌────────────────────┐ │
│ │ USD ▼│ │         5,000.00   │ │
│ └──────┘ └────────────────────┘ │
│ Available balance: $45,000.00   │
└─────────────────────────────────┘
```

**Banking-specific:** Locale-aware formatting; currency-appropriate decimal places (USD: 2, JPY: 0); min 0.01; max per daily limit; monospace font for amounts.

### 3.3 Account Selector Molecule

**Composition:** Label atom + Dropdown atom + Account info (number masked + balance)

```
┌─────────────────────────────────┐
│ From Account *                  │
│ ┌─────────────────────────────┐ │
│ │ Checking ****1234            │ │
│ │ Available: $45,000.00     ▼ │ │
│ └─────────────────────────────┘ │
│ ┌─────────────────────────────┐ │
│ │ ○ Checking ****1234         │ │  ← Dropdown options
│ │   Available: $45,000.00     │ │
│ │ ○ Savings ****5678          │ │
│ │   Available: $120,000.00    │ │
│ └─────────────────────────────┘ │
└─────────────────────────────────┘
```

**Security:** Account numbers masked (last 4 digits); only authenticated user's accounts shown; server-side ownership validation.

### 3.4 Search Bar Molecule

**Composition:** Input atom (search) + Button atom (submit) + Icon atom (search)

```
┌─────────────────────────────────────┐
│ 🔍 │ Search transactions...  │ Search │
└─────────────────────────────────────┘
```

**Accessibility:** `role="search"`; `aria-label="Search transactions"`; keyboard: Enter to submit.

### 3.5 Navigation Item Molecule

**Composition:** Icon atom + Label atom + Badge atom (optional, for notification count)

```
│ 🏠 Dashboard                    │
│ 💳 Accounts                     │
│ ↗️ Transfers              (2)   │  ← Badge: 2 pending
│ 📋 Beneficiaries                │
│ ⚙️ Settings                     │
```

### 3.6 Alert Molecule

**Composition:** Icon atom (status) + Text atom (message) + Button atom (dismiss/action)

| Variant | Icon | Color | Usage |
|---|---|---|---|
| Success | ✓ | Green | "Transfer completed successfully" |
| Warning | ⚠ | Yellow | "Your session expires in 2 minutes" |
| Error | ✕ | Red | "Transfer failed. Please try again." |
| Info | ℹ | Blue | "New feature: International transfers now available" |

**Accessibility:** `role="alert"` for errors/warnings; `role="status"` for success/info; `aria-live="assertive"` for errors, `"polite"` for info.

### 3.7 Transaction Row Molecule

**Composition:** Icon atom (type) + Text atoms (description, date) + Amount atom + Badge atom (status)

```
│ ↗️ │ Transfer to John Smith    │ 15 Mar 2026 │ -$5,000.00 │ ✓ Completed │
│ ↙️ │ Salary Deposit            │ 14 Mar 2026 │ +$8,500.00 │ ✓ Completed │
│ ↗️ │ Bill Payment - Electric   │ 13 Mar 2026 │   -$150.00 │ ⏳ Pending  │
```

**Banking-specific:** Debit amounts in red with minus; credit amounts in green with plus; monospace for amounts; date in user locale.

### 3.8 Stat Card Molecule

**Composition:** Label atom + Amount atom + Trend indicator (icon + percentage)

```
┌──────────────────┐
│ Available Balance │
│ $45,000.00       │
│ ↑ 12% this month │
└──────────────────┘
```

### 3.9 Empty State Molecule

**Composition:** Icon/illustration + Heading atom + Description atom + Action button atom

```
┌─────────────────────────────────┐
│         📋                      │
│   No transactions yet           │
│   Your recent transactions      │
│   will appear here.             │
│   [Make a Transfer]             │
└─────────────────────────────────┘
```

### 3.10 Confirmation Step Molecule

**Composition:** Label-value pairs + Divider + Total amount

```
┌─────────────────────────────────┐
│ From:        Checking ****1234  │
│ To:          John Smith ****5678│
│ Amount:      $5,000.00 USD      │
│ Reference:   Invoice 2026-001   │
│ ─────────────────────────────── │
│ Total:       $5,000.00 USD      │
└─────────────────────────────────┘
```

### 3.11 OTP Input Molecule

**Composition:** 6 × Input atoms (single digit) + Timer atom + Resend link atom

```
┌─────────────────────────────────┐
│ Enter verification code         │
│ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ ┌───┐ │
│ │ 4 │ │ 8 │ │ 2 │ │   │ │   │ │   │ │
│ └───┘ └───┘ └───┘ └───┘ └───┘ └───┘ │
│ Code expires in 1:23            │
│ Didn't receive? Resend          │
└─────────────────────────────────┘
```

**Security:** Auto-focus next digit; auto-submit on complete; 90-second expiry; single-use; paste allowed (accessibility); `autocomplete="one-time-code"`.

### 3.12 Pagination Molecule

**Composition:** Button atoms (prev/next) + Page number atoms + Text atom (showing X of Y)

```
│ ← Previous │ 1 │ 2 │ [3] │ 4 │ ... │ 8 │ Next → │
│ Showing 41-60 of 156 transactions                  │
```

---

## 4. Organisms

Organisms are complex UI components composed of molecules and atoms, forming distinct sections of the interface.

### 4.1 Header Organism

**Composition:** Logo atom + Navigation molecule + Account menu molecule + Notification bell molecule

```
┌──────────────────────────────────────────────────────────────┐
│ [Logo]  Dashboard  Accounts  Transfers  Cards    🔔(3) 👤 JD │
└──────────────────────────────────────────────────────────────┘
```

| Element | Behavior |
|---|---|
| Logo | Links to dashboard; alt text: "Bank Name" |
| Navigation | Active state on current page; keyboard navigable |
| Notification bell | Badge count for unread; dropdown on click |
| Account menu | Dropdown: Profile, Settings, Logout |

**Responsive:** Collapses to hamburger menu on mobile (< 768px).

### 4.2 Sidebar Navigation Organism

**Composition:** Navigation item molecules (stacked) + User info molecule + Collapse toggle

```
┌────────────────────┐
│ 👤 John Doe        │
│ Checking ****1234  │
│ ────────────────── │
│ 🏠 Dashboard       │
│ 💳 Accounts        │
│ ↗️ Transfers  (2)  │
│ 👥 Beneficiaries   │
│ 💳 Cards           │
│ 📊 Statements      │
│ ⚙️ Settings        │
│ ────────────────── │
│ 🔒 Logout          │
└────────────────────┘
```

### 4.3 Account Summary Organism

**Composition:** Stat card molecules (balance, income, expenses) + Account selector molecule

```
┌──────────────────────────────────────────────────────────────┐
│ Account: Checking ****1234 ▼                                 │
│ ┌──────────────┐ ┌──────────────┐ ┌──────────────┐          │
│ │ Available    │ │ This Month   │ │ This Month   │          │
│ │ $45,000.00   │ │ Income       │ │ Expenses     │          │
│ │              │ │ +$12,500.00  │ │ -$7,500.00   │          │
│ └──────────────┘ └──────────────┘ └──────────────┘          │
└──────────────────────────────────────────────────────────────┘
```

### 4.4 Transaction List Organism

**Composition:** Search bar molecule + Filter molecules + Transaction row molecules + Pagination molecule + Empty state molecule

```
┌──────────────────────────────────────────────────────────────┐
│ Recent Transactions                                          │
│ ┌──────────────────────────────────────────────────────────┐ │
│ │ 🔍 Search...  │ All Types ▼ │ Date Range ▼ │ Export ↓  │ │
│ └──────────────────────────────────────────────────────────┘ │
│ ┌──────────────────────────────────────────────────────────┐ │
│ │ ↗️ Transfer to John Smith   │ 15 Mar │ -$5,000.00 │ ✓  │ │
│ │ ↙️ Salary Deposit           │ 14 Mar │ +$8,500.00 │ ✓  │ │
│ │ ↗️ Bill Payment - Electric  │ 13 Mar │   -$150.00 │ ⏳ │ │
│ │ ↗️ Transfer to Jane Doe     │ 12 Mar │ -$2,000.00 │ ✓  │ │
│ └──────────────────────────────────────────────────────────┘ │
│ │ ← Previous │ 1 │ [2] │ 3 │ Next → │ Showing 21-40 of 156│
└──────────────────────────────────────────────────────────────┘
```

**Performance:** Virtualized list for > 50 items; server-side pagination (max 20 per page); debounced search (300ms).

### 4.5 Transfer Form Organism

**Composition:** Account selector molecule + Beneficiary selector molecule + Amount input molecule + Form field molecule (reference) + Confirmation step molecule + OTP input molecule

```
Step 1: Transfer Details          Step 2: Confirm          Step 3: Verify
┌────────────────────────┐  →  ┌──────────────────┐  →  ┌──────────────────┐
│ From Account *         │     │ From: ****1234   │     │ Enter OTP        │
│ [Checking ****1234  ▼] │     │ To:   ****5678   │     │ [_ _ _ _ _ _]    │
│                        │     │ Amount: $5,000   │     │ Expires: 1:23    │
│ To Beneficiary *       │     │ Ref: Inv-001     │     │                  │
│ [John Smith ****5678▼] │     │ ──────────────── │     │ [Verify & Send]  │
│                        │     │ Total: $5,000.00 │     └──────────────────┘
│ Amount *               │     │                  │
│ [USD ▼] [5,000.00   ] │     │ [Edit] [Confirm] │
│ Available: $45,000     │     └──────────────────┘
│                        │
│ Reference              │
│ [Invoice 2026-001   ]  │
│                        │
│ [Cancel] [Continue →]  │
└────────────────────────┘
```

**Security:** CSRF token; MFA step-up; server-side validation; idempotency key; double-submit prevention (disable button on click).

### 4.6 Beneficiary List Organism

**Composition:** Search bar molecule + Beneficiary card molecules + Add button atom + Empty state molecule

```
┌──────────────────────────────────────────────────────────────┐
│ Beneficiaries                                    [+ Add New] │
│ ┌──────────────────────────────────────────────────────────┐ │
│ │ 🔍 Search beneficiaries...                               │ │
│ └──────────────────────────────────────────────────────────┘ │
│ ┌──────────────────────┐ ┌──────────────────────┐           │
│ │ 👤 John Smith        │ │ 👤 Jane Doe          │           │
│ │ ****5678             │ │ ****9012             │           │
│ │ Bank of America      │ │ Chase Bank           │           │
│ │ [Transfer] [Edit] [⋮]│ │ [Transfer] [Edit] [⋮]│           │
│ └──────────────────────┘ └──────────────────────┘           │
└──────────────────────────────────────────────────────────────┘
```

### 4.7 Session Timeout Warning Organism

**Composition:** Modal overlay + Heading atom + Description atom + Timer atom + Button atoms (Extend/Logout)

```
┌──────────────────────────────────────┐
│         Session Expiring             │
│                                      │
│  Your session will expire in 1:45    │
│                                      │
│  Would you like to continue?         │
│                                      │
│  [Logout]          [Continue Session]│
└──────────────────────────────────────┘
```

**Accessibility:** `aria-modal="true"`; focus trapped; ESC to dismiss (extends session); `aria-live="assertive"` for countdown; return focus to trigger on close.

### 4.8 Notification Panel Organism

**Composition:** Heading atom + Notification item molecules (icon + text + time + action) + Mark all read link

```
┌──────────────────────────────────┐
│ Notifications          Mark all  │
│ ────────────────────────────────│
│ ● Transfer of $5,000 completed  │
│   15 Mar 2026, 3:30 PM          │
│ ● New beneficiary added         │
│   15 Mar 2026, 2:15 PM          │
│ ○ Statement available           │
│   14 Mar 2026, 9:00 AM          │
└──────────────────────────────────┘
```

### 4.9 Footer Organism

**Composition:** Link groups (About, Legal, Support) + Copyright atom + Regulatory info atom

```
┌──────────────────────────────────────────────────────────────┐
│ About Us │ Privacy Policy │ Terms │ Security │ Contact Us    │
│ © 2026 Bank Name. All rights reserved. Member FDIC.         │
│ Regulated by [Authority]. License No. XXXXX.                │
└──────────────────────────────────────────────────────────────┘
```

### 4.10 Error Page Organism

**Composition:** Icon/illustration + Heading atom + Description atom + Action buttons

```
┌──────────────────────────────────────┐
│              ⚠️                       │
│     Something went wrong             │
│                                      │
│  We're having trouble loading this   │
│  page. Please try again.             │
│                                      │
│  [Try Again]  [Go to Dashboard]      │
│                                      │
│  Reference: REQ-abc-123              │
└──────────────────────────────────────┘
```

**Security:** No stack traces; no internal details; reference ID for support correlation.
