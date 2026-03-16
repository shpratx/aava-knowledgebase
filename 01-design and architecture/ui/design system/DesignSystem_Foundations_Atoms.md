# Banking Design System — Foundations & Atoms
---

## Overview

This design system follows Brad Frost's Atomic Design methodology, organizing UI components into five hierarchical levels: Atoms → Molecules → Organisms → Templates → Pages. Each level builds upon the previous, creating a consistent, accessible, and secure banking interface.

```
Pages          → Specific instances with real content (Account Dashboard with John's data)
Templates      → Page layouts with content structure (Dashboard layout with placeholders)
Organisms      → Complex UI sections (Header, Transaction List, Transfer Form)
Molecules      → Simple component groups (Search Bar, Account Card, Form Field)
Atoms          → Foundational elements (Button, Input, Label, Icon, Typography)
```

---

## 1. Design Tokens (Foundation Layer)

Design tokens are the sub-atomic particles — the raw values that feed into every atom.

### 1.1 Color Tokens

| Token | Light Theme | Dark Theme | Usage |
|---|---|---|---|
| `--color-primary-500` | #1A56DB | #3B82F6 | Primary actions, links, focus |
| `--color-primary-600` | #1E40AF | #2563EB | Primary hover |
| `--color-primary-700` | #1E3A8A | #1D4ED8 | Primary active/pressed |
| `--color-secondary-500` | #6B7280 | #9CA3AF | Secondary actions |
| `--color-success-500` | #059669 | #10B981 | Success states, completed transactions |
| `--color-warning-500` | #D97706 | #F59E0B | Warning states, pending items |
| `--color-error-500` | #DC2626 | #EF4444 | Error states, failed transactions |
| `--color-info-500` | #2563EB | #3B82F6 | Informational messages |
| `--color-neutral-50` | #F9FAFB | #111827 | Page background |
| `--color-neutral-100` | #F3F4F6 | #1F2937 | Card background |
| `--color-neutral-200` | #E5E7EB | #374151 | Borders, dividers |
| `--color-neutral-700` | #374151 | #D1D5DB | Secondary text |
| `--color-neutral-900` | #111827 | #F9FAFB | Primary text |
| `--color-surface` | #FFFFFF | #1F2937 | Card/panel surface |
| `--color-overlay` | rgba(0,0,0,0.5) | rgba(0,0,0,0.7) | Modal overlay |

**Accessibility:** All text/background combinations must meet WCAG 2.1 AA contrast ratios (≥ 4.5:1 normal text, ≥ 3:1 large text).

### 1.2 Typography Tokens

| Token | Value | Usage |
|---|---|---|
| `--font-family-primary` | 'Inter', system-ui, sans-serif | Body text, UI elements |
| `--font-family-mono` | 'JetBrains Mono', monospace | Account numbers, amounts, code |
| `--font-size-xs` | 0.75rem (12px) | Captions, helper text |
| `--font-size-sm` | 0.875rem (14px) | Secondary text, table cells |
| `--font-size-base` | 1rem (16px) | Body text, inputs |
| `--font-size-lg` | 1.125rem (18px) | Subheadings |
| `--font-size-xl` | 1.25rem (20px) | Section headings |
| `--font-size-2xl` | 1.5rem (24px) | Page headings |
| `--font-size-3xl` | 1.875rem (30px) | Hero headings |
| `--font-weight-regular` | 400 | Body text |
| `--font-weight-medium` | 500 | Labels, emphasis |
| `--font-weight-semibold` | 600 | Headings, buttons |
| `--font-weight-bold` | 700 | Amounts, key figures |
| `--line-height-tight` | 1.25 | Headings |
| `--line-height-normal` | 1.5 | Body text |
| `--line-height-relaxed` | 1.75 | Long-form content |

### 1.3 Spacing Tokens

| Token | Value | Usage |
|---|---|---|
| `--space-1` | 0.25rem (4px) | Inline element gaps |
| `--space-2` | 0.5rem (8px) | Tight spacing (icon + label) |
| `--space-3` | 0.75rem (12px) | Form field internal padding |
| `--space-4` | 1rem (16px) | Standard component padding |
| `--space-5` | 1.25rem (20px) | Card padding |
| `--space-6` | 1.5rem (24px) | Section spacing |
| `--space-8` | 2rem (32px) | Large section gaps |
| `--space-10` | 2.5rem (40px) | Page section separation |
| `--space-12` | 3rem (48px) | Major layout gaps |
| `--space-16` | 4rem (64px) | Page-level spacing |

### 1.4 Border & Shadow Tokens

| Token | Value | Usage |
|---|---|---|
| `--radius-sm` | 0.25rem (4px) | Buttons, badges |
| `--radius-md` | 0.5rem (8px) | Cards, inputs |
| `--radius-lg` | 0.75rem (12px) | Modals, panels |
| `--radius-full` | 9999px | Avatars, pills |
| `--border-width` | 1px | Default borders |
| `--border-focus` | 2px | Focus indicators (accessibility) |
| `--shadow-sm` | 0 1px 2px rgba(0,0,0,0.05) | Subtle elevation |
| `--shadow-md` | 0 4px 6px rgba(0,0,0,0.1) | Cards |
| `--shadow-lg` | 0 10px 15px rgba(0,0,0,0.1) | Dropdowns, popovers |
| `--shadow-xl` | 0 20px 25px rgba(0,0,0,0.15) | Modals |

### 1.5 Breakpoint Tokens

| Token | Value | Usage |
|---|---|---|
| `--breakpoint-sm` | 640px | Mobile landscape |
| `--breakpoint-md` | 768px | Tablet |
| `--breakpoint-lg` | 1024px | Desktop |
| `--breakpoint-xl` | 1280px | Large desktop |

### 1.6 Motion Tokens

| Token | Value | Usage |
|---|---|---|
| `--duration-fast` | 150ms | Hover, focus transitions |
| `--duration-normal` | 300ms | Panel open/close |
| `--duration-slow` | 500ms | Page transitions |
| `--easing-default` | ease-in-out | Standard transitions |

**Accessibility:** All motion must respect `prefers-reduced-motion: reduce` — disable animations when user prefers reduced motion.

---

## 2. Atoms

Atoms are the foundational building blocks — individual HTML elements that cannot be broken down further without losing meaning.

### 2.1 Button Atom

| Variant | Usage | Visual |
|---|---|---|
| Primary | Main actions (Submit Transfer, Confirm) | Filled, primary color, white text |
| Secondary | Alternative actions (Cancel, Back) | Outlined, primary color border |
| Tertiary | Low-emphasis actions (Learn More, Details) | Text only, primary color |
| Danger | Destructive actions (Delete Beneficiary, Block Card) | Filled, error color |
| Ghost | Minimal emphasis (Close, Dismiss) | Transparent, icon or text |

| Size | Height | Padding | Font Size | Touch Target |
|---|---|---|---|---|
| Small | 32px | 8px 12px | 14px | 44×44px (mobile) |
| Medium (default) | 40px | 10px 16px | 16px | 44×44px |
| Large | 48px | 12px 24px | 18px | 48×48px |

| State | Behavior |
|---|---|
| Default | Standard appearance |
| Hover | Darken background 10%; cursor: pointer |
| Focus | 2px focus ring (--color-primary-500); visible outline |
| Active/Pressed | Darken 20%; slight scale(0.98) |
| Disabled | 50% opacity; cursor: not-allowed; no interaction |
| Loading | Spinner replaces text; disabled; aria-busy="true" |

**Accessibility:** `role="button"` if not `<button>`; `aria-label` for icon-only buttons; `aria-disabled` for disabled state; keyboard: Enter/Space to activate.

### 2.2 Input Atom

| Variant | Usage |
|---|---|
| Text | Names, references, search |
| Number | Amounts (with currency formatting) |
| Password | Passwords, PINs (type="password") |
| Email | Email addresses |
| Tel | Phone numbers |
| Date | Date selection |
| Textarea | Multi-line (notes, remittance info) |

| State | Visual |
|---|---|
| Default | 1px neutral border |
| Focus | 2px primary border; subtle shadow |
| Filled | Neutral border; value displayed |
| Error | 2px error border; error icon; error message below |
| Disabled | Gray background; no interaction |
| Read-only | No border; text displayed as static |

**Accessibility:** Always paired with `<label>` (via `for`/`id`); `aria-describedby` for help text and errors; `aria-invalid="true"` on error; `aria-required="true"` for required fields.

**Security:** `autocomplete="off"` for OTP/CVV; `type="password"` for sensitive fields; no sensitive data stored in DOM after submission.

### 2.3 Label Atom

| Variant | Usage |
|---|---|
| Form label | Associated with input via for/id |
| Required indicator | Red asterisk (*) with sr-only "required" text |
| Optional indicator | "(optional)" text suffix |
| Help text | Below input; lighter color; smaller font |
| Error text | Below input; error color; error icon prefix |

### 2.4 Typography Atoms

| Element | Tag | Token | Usage |
|---|---|---|---|
| Page Title | h1 | 3xl / semibold | One per page: "Account Overview" |
| Section Title | h2 | 2xl / semibold | Section headings: "Recent Transactions" |
| Subsection Title | h3 | xl / semibold | Card headings: "Transfer Details" |
| Body | p | base / regular | Paragraphs, descriptions |
| Body Small | p | sm / regular | Secondary information |
| Caption | span | xs / regular | Timestamps, metadata |
| Amount | span | lg-3xl / bold / mono | Financial amounts: "$5,000.00" |
| Account Number | span | base / mono | Masked: "****1234" |
| Status Text | span | sm / medium | "Completed", "Pending", "Failed" |
| Link | a | base / medium / primary | Navigation, actions |

### 2.5 Icon Atom

| Category | Icons | Size |
|---|---|---|
| Navigation | Home, Accounts, Transfers, Cards, Settings, Menu, Back, Close | 24×24px |
| Action | Add, Edit, Delete, Copy, Download, Upload, Search, Filter | 20×20px |
| Status | Success (checkmark), Warning (triangle), Error (circle-x), Info (circle-i) | 20×20px |
| Finance | Transfer, Payment, Deposit, Withdrawal, Exchange, Statement | 24×24px |
| Security | Lock, Unlock, Shield, Key, Fingerprint, Eye, Eye-off | 20×20px |

**Accessibility:** Decorative icons: `aria-hidden="true"`. Meaningful icons: `aria-label="description"`. Icon buttons: `aria-label` required.

### 2.6 Badge Atom

| Variant | Color | Usage |
|---|---|---|
| Success | Green bg, green text | "Completed", "Active", "Verified" |
| Warning | Yellow bg, yellow text | "Pending", "Processing", "Under Review" |
| Error | Red bg, red text | "Failed", "Rejected", "Frozen" |
| Info | Blue bg, blue text | "New", "Updated" |
| Neutral | Gray bg, gray text | "Draft", "Inactive" |

### 2.7 Avatar Atom

| Size | Dimension | Usage |
|---|---|---|
| Small | 32×32px | Inline mentions, compact lists |
| Medium | 40×40px | Navigation, cards |
| Large | 64×64px | Profile page |

Fallback: Initials on colored background when no image available.

### 2.8 Divider Atom

| Variant | Usage |
|---|---|
| Horizontal | Section separation within cards |
| Vertical | Inline element separation |
| With label | "OR" divider between options |

### 2.9 Skeleton Atom

| Variant | Usage |
|---|---|
| Text line | Placeholder for text content during loading |
| Circle | Placeholder for avatar during loading |
| Rectangle | Placeholder for card/image during loading |

**Accessibility:** `aria-busy="true"` on loading container; `aria-live="polite"` to announce when content loads.

### 2.10 Spinner Atom

| Size | Usage |
|---|---|
| Small (16px) | Inline loading (button, input) |
| Medium (24px) | Component loading |
| Large (48px) | Page/section loading |

**Accessibility:** `role="status"`; `aria-label="Loading"`.
