# Test Scenario Template — UI Features
### Banking Domain — Agentic Knowledge Base

---

## Test Scenario Metadata

| Field | Description |
|---|---|
| **Scenario ID** | TS-UI-[Module]-[Seq] (e.g., TS-UI-TRF-001) |
| **Feature** | Feature/User Story ID and title |
| **Type** | Functional / Security / Accessibility / Performance / Usability / Cross-Browser / Responsive |
| **Priority** | Critical / High / Medium / Low |
| **Preconditions** | System state required before testing |
| **Test Data** | Specific data needed (accounts, amounts, beneficiaries) |
| **Environment** | Staging / QA / UAT |
| **Automation Status** | Automated / Manual / Hybrid |

---

## Test Scenario Structure

### Scenario Description
| Field | Content |
|---|---|
| **Objective** | What this scenario validates |
| **User Persona** | Who is performing the action (retail customer, premium, admin) |
| **Entry Point** | Where the user starts (dashboard, direct URL, deep link) |
| **Expected Flow** | Step-by-step user journey |
| **Exit Criteria** | What constitutes scenario pass/fail |

---

## UI Test Scenario Categories

### 1. Functional Scenarios

| Category | Scenarios to Cover |
|---|---|
| Happy path | Complete the feature successfully end-to-end |
| Validation | Submit with missing required fields; invalid formats; boundary values |
| Error handling | Server error (500); timeout; network offline; API unavailable |
| State management | Data persists across navigation; cleared on logout; refreshed on return |
| Navigation | Forward/back browser buttons; breadcrumb navigation; deep linking |
| Multi-step flows | Complete all steps; go back and edit; abandon mid-flow; resume |
| Conditional display | Elements shown/hidden based on user role, account type, feature flags |
| Empty states | No data available; first-time user; search with no results |
| Loading states | Skeleton screens displayed; spinner on actions; no flash of unstyled content |
| Pagination | First page; last page; page navigation; items per page change |
| Sorting/filtering | Sort ascending/descending; filter by type/date/status; clear filters; combined filters |

### 2. Security Scenarios

| Category | Scenarios to Cover |
|---|---|
| Authentication | Access without login → redirect; expired session → re-auth; MFA step-up for sensitive ops |
| Authorization | Access other user's data via URL manipulation; role-based element visibility |
| Session | Timeout warning at 13 min; timeout at 15 min; extend session; concurrent session detection |
| Input injection | XSS in text fields; SQL injection in search; script in URL params |
| CSRF | Submit form without CSRF token; replay old CSRF token |
| Sensitive data | No PII in URL; no sensitive data in localStorage; account numbers masked; no data in page source |
| Error leakage | Error messages don't expose stack traces, internal paths, or technology details |

### 3. Accessibility Scenarios

| Category | Scenarios to Cover |
|---|---|
| Keyboard | Complete entire flow using keyboard only; tab order is logical; no keyboard traps |
| Screen reader | All elements announced correctly; form labels read; errors announced; status changes announced |
| Visual | Color contrast meets WCAG 2.1 AA; no color-only information; focus indicators visible |
| Zoom | Page usable at 200% zoom; no horizontal scroll at 320px width |
| Motion | Animations respect prefers-reduced-motion |
| Forms | Error messages linked to fields; required fields indicated; help text available |

### 4. Responsive / Cross-Browser Scenarios

| Category | Scenarios to Cover |
|---|---|
| Mobile (< 768px) | Layout stacks correctly; touch targets ≥ 44px; bottom nav visible; no horizontal overflow |
| Tablet (768-1024px) | Sidebar collapses; content adapts; touch and mouse both work |
| Desktop (> 1024px) | Full layout; sidebar visible; data tables render correctly |
| Browsers | Chrome, Firefox, Safari, Edge (latest 2 versions) |
| OS | Windows, macOS, iOS, Android |

---

## UI Test Case Template

| Field | Description |
|---|---|
| **Test Case ID** | TC-UI-[Module]-[Seq] (e.g., TC-UI-TRF-001) |
| **Scenario ID** | Parent test scenario |
| **Title** | Descriptive title |
| **Type** | Functional / Security / Accessibility / Performance / Responsive |
| **Priority** | Critical / High / Medium / Low |
| **Preconditions** | State before test execution |
| **Test Data** | Specific values used |
| **Steps** | Numbered step-by-step actions |
| **Expected Result** | What should happen at each step and final outcome |
| **Actual Result** | (Filled during execution) |
| **Status** | Pass / Fail / Blocked / Skipped |
| **Screenshots/Evidence** | Attached evidence |
| **Defect ID** | Linked defect if failed |
| **Automation** | Automated (Cypress/Playwright) / Manual |

---

## Example: Fund Transfer UI Test Cases

### TC-UI-TRF-001: Happy Path — Initiate Domestic Transfer

| Field | Value |
|---|---|
| **Scenario** | TS-UI-TRF-001: Complete domestic transfer |
| **Priority** | Critical |
| **Preconditions** | Customer logged in; has Checking ****1234 with $45,000 balance; beneficiary John Smith ****5678 registered |
| **Test Data** | Amount: $5,000; Currency: USD; Reference: "Invoice 2026-001" |

| Step | Action | Expected Result |
|---|---|---|
| 1 | Navigate to Transfers > New Transfer | Transfer form displayed with Step 1 (Details) active |
| 2 | Select source account "Checking ****1234" | Account selected; available balance "$45,000.00" shown |
| 3 | Select beneficiary "John Smith ****5678" | Beneficiary selected |
| 4 | Enter amount "5000" | Amount formatted as "$5,000.00"; within daily limit |
| 5 | Enter reference "Invoice 2026-001" | Reference accepted (alphanumeric, < 140 chars) |
| 6 | Click "Continue" | Step 2 (Confirm) displayed with summary: From ****1234, To ****5678, $5,000.00 USD, Ref: Invoice 2026-001 |
| 7 | Click "Confirm" | MFA step triggered; OTP input displayed with 90-second timer |
| 8 | Enter valid OTP | Transfer initiated; success message: "Transfer of $5,000.00 to John Smith initiated"; transfer reference displayed |
| 9 | Verify transaction in history | Transfer appears in recent transactions with status "Initiated" or "Completed" |

### TC-UI-TRF-002: Validation — Amount Exceeds Daily Limit

| Step | Action | Expected Result |
|---|---|---|
| 1 | Navigate to transfer form | Form displayed |
| 2 | Select account with $45,000 balance | Account selected |
| 3 | Enter amount "$55,000" (exceeds $50,000 daily limit) | Inline error: "Daily transfer limit exceeded. Remaining: $50,000.00" |
| 4 | Verify "Continue" button | Button disabled or click shows error |
| 5 | Verify no API call made | Network tab shows no transfer API call |

### TC-UI-TRF-003: Security — Session Timeout During Transfer

| Step | Action | Expected Result |
|---|---|---|
| 1 | Start transfer flow; reach Step 2 (Confirm) | Confirmation screen displayed |
| 2 | Wait 13 minutes (no activity) | Session timeout warning modal appears: "Session expires in 2:00" |
| 3 | Do not interact for 2 more minutes | Session expires; redirect to login page; message: "Session expired. Please log in again." |
| 4 | Verify sensitive data cleared | No transfer data in localStorage, sessionStorage, or DOM |
| 5 | Log in again | Dashboard displayed; no partial transfer state retained |

### TC-UI-TRF-004: Accessibility — Keyboard-Only Transfer

| Step | Action | Expected Result |
|---|---|---|
| 1 | Tab to source account selector | Focus visible on selector; screen reader announces "From Account, required" |
| 2 | Use arrow keys to select account | Account selected; balance announced |
| 3 | Tab to beneficiary selector | Focus moves to beneficiary; announced correctly |
| 4 | Tab to amount field | Focus on amount; label announced "Transfer Amount, required" |
| 5 | Type amount; Tab to reference | Amount accepted; focus moves to reference |
| 6 | Tab to Continue button; press Enter | Step 2 displayed; focus moves to confirmation summary |
| 7 | Tab to Confirm; press Enter | MFA triggered; focus moves to OTP input |
| 8 | Enter OTP digits (auto-advance) | Each digit accepted; auto-advance to next; auto-submit on 6th digit |
| 9 | Verify success announcement | Screen reader announces "Transfer of $5,000.00 initiated successfully" |

### TC-UI-TRF-005: Responsive — Mobile Transfer Flow

| Step | Action | Expected Result |
|---|---|---|
| 1 | Open transfer form on mobile (375px width) | Single-column layout; all fields stacked; touch targets ≥ 44px |
| 2 | Tap source account | Bottom sheet selector opens (not dropdown) |
| 3 | Tap amount field | Numeric keyboard opens; currency prefix visible |
| 4 | Complete transfer flow | All steps work; confirmation readable; OTP input usable |
| 5 | Verify no horizontal scroll | No content overflows viewport |
