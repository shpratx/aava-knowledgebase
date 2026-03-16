# User Story Template

---

## Template Fields

| Field | Description |
|---|---|
| **Story ID** | US-[Module]-[Seq] (e.g., US-PAY-042, US-ONB-015) |
| **Title** | Short action-oriented title |
| **Epic** | Parent epic ID and title |
| **User Story Statement** | As a [role], I want [capability], so that [benefit] (ref: PM-BP-015) |
| **Priority** | Critical / High / Medium / Low |
| **Story Points** | Estimated effort |
| **Data Sensitivity Tag** | Public / Internal / Confidential / Restricted (ref: PM-BP-017) |
| **Data Fields Involved** | List of data fields this story creates, reads, updates, or deletes — with classification per field |
| **Authentication Requirement** | Required auth level: None / Basic Auth / MFA / Step-Up MFA (ref: PM-BP-018) |
| **Authorization Requirement** | Required permissions: role(s), resource-level checks, segregation of duties (ref: PM-BP-018) |
| **Description** | Detailed context, business logic, and user interaction |
| **Acceptance Criteria** | Given/When/Then — including security and compliance checks (ref: PM-BP-016) |
| **Security Acceptance Criteria** | Dedicated security conditions that must pass (ref: PM-BP-016) |
| **Compliance Acceptance Criteria** | Dedicated regulatory conditions that must pass (ref: PM-BP-016) |
| **Audit Logging Requirements** | Events to log, fields to capture, retention period (ref: PM-BP-019) |
| **Regulatory Linkage** | Applicable regulation(s) with article reference, or "None applicable" |
| **Dependencies** | Other stories, APIs, systems, or teams |
| **Assumptions** | Assumptions made |
| **UI/UX Notes** | Wireframe references, interaction details, accessibility requirements |
| **API Contract** | Endpoint, method, request/response schema (if applicable) |
| **Definition of Done** | Checklist for story completion |

---

## Data Sensitivity Tagging Detail (ref: PM-BP-017)

Tag every story with the highest sensitivity level of data it touches:

| Tag | Data Examples | Implications |
|---|---|---|
| **Restricted** | Card numbers (PAN), CVV, PINs, biometrics, encryption keys, passwords | AES-256 encryption, HSM keys, MFA mandatory, full access logging, tokenization in non-prod, no storage post-auth for CVV/PIN |
| **Confidential** | Customer PII (name, address, DOB, email, phone), account balances, transaction history, financial records | AES-256 encryption, MFA for modifications, modification logging, data masking in non-prod |
| **Internal** | Internal policies, employee IDs, system configurations, non-public business data | Access-controlled, auth event logging, no public exposure |
| **Public** | Published rates, branch locations, product brochures, marketing content | Standard controls, no special handling |

---

## Authentication & Authorization Detail (ref: PM-BP-018)

| Operation Type | Auth Level | Authorization | Examples |
|---|---|---|---|
| View public information | None | None | Interest rates, branch locator |
| View own account data | Basic Auth (session) | Own-account only | Balance inquiry, transaction history |
| Modify own profile | Basic Auth + verification | Own-profile only | Update address, change email |
| Initiate financial transaction | MFA / Step-Up MFA | Own-account + limit check | Fund transfer, bill payment |
| Approve financial transaction | MFA | Maker-checker role | Approve bulk payment, authorize limit change |
| Administrative operation | MFA + privileged access | Admin role + JIT approval | User management, config change, role assignment |
| Access other customer's data | MFA + privileged access | Authorized role + audit | Branch teller serving customer, RM accessing portfolio |

---

## Audit Logging Detail (ref: PM-BP-019)

| Story Involves... | What to Log | Fields | Retention |
|---|---|---|---|
| Financial transaction | Every attempt (success + failure) | User ID, account, amount, beneficiary, timestamp, outcome, correlation ID, fraud score | 7 years |
| Data read (sensitive) | Every access | User ID, data accessed, account viewed, timestamp, source IP | 7 years |
| Data modification | Before/after values | User ID, field changed, old value, new value, timestamp, reason | 7 years |
| Authentication event | Every attempt | User ID, auth method, timestamp, IP, device, outcome | 3 years |
| Approval/rejection | Decision + context | Approver ID, request ID, decision, timestamp, reason | 7 years |
| Configuration change | Every change | Admin ID, setting changed, old/new value, timestamp | 7 years |
| No sensitive data involved | Minimal or none | — | — |

---

## Example — User Story

| Field | Value |
|---|---|
| **Story ID** | US-PAY-042 |
| **Title** | Initiate Domestic Fund Transfer |
| **Epic** | EP-PAY-003: Digital Payments Modernization — Domestic Fund Transfers |
| **User Story Statement** | As a retail banking customer, I want to transfer funds to a registered beneficiary from my mobile or web banking app, so that I can make payments without visiting a branch. |
| **Priority** | Critical |
| **Story Points** | 13 |
| **Data Sensitivity Tag** | Restricted (account numbers, transaction amounts) |
| **Data Fields Involved** | Source account (Confidential, read), beneficiary account (Confidential, read), amount (Confidential, create), reference (Internal, create), transaction ID (Internal, create), balance (Confidential, read+update) |
| **Authentication Requirement** | Step-Up MFA — TOTP or push notification required before transfer submission |
| **Authorization Requirement** | Customer role, own-account access only, daily transfer limit check ($50K retail / $500K premium) |
| **Description** | Authenticated customer selects source account and registered beneficiary, enters amount and optional reference. System validates balance, daily limits, and performs real-time fraud screening. On success, funds are debited and beneficiary credited. Confirmation with transaction reference is displayed. |
| **Acceptance Criteria** | **AC1:** Given an authenticated customer with sufficient balance, When they submit a valid transfer of $5,000, Then the transfer completes within 5 seconds and confirmation with reference number is displayed. **AC2:** Given a customer who has transferred $45,000 today (limit $50,000), When they attempt $10,000, Then the transfer is rejected with "Daily limit exceeded. Remaining: $5,000." **AC3:** Given a transfer flagged by fraud engine, When fraud score exceeds threshold, Then the transfer is held, customer notified, and fraud team alerted. |
| **Security Acceptance Criteria** | **SAC1:** Transfer requires step-up MFA — expired OTPs (>90s) and replayed OTPs are rejected. **SAC2:** Customer can only transfer from own accounts — attempting another customer's account returns 403. **SAC3:** 5 consecutive MFA failures lock the transfer function for 30 minutes. **SAC4:** All input fields validated server-side — SQL injection and XSS attempts rejected. |
| **Compliance Acceptance Criteria** | **CAC1:** Transfers ≥ $10,000 automatically generate a Currency Transaction Report (CTR) per BSA/AML. **CAC2:** Step-up MFA satisfies PSD2 SCA (Art. 97) with two independent factors. **CAC3:** All customer data processing has documented legal basis per GDPR Art. 6. |
| **Audit Logging Requirements** | Log: transfer initiated, MFA result, fraud check result, balance validated, transfer executed/rejected/held. Fields: user ID, session ID, source IP, device fingerprint, source account, beneficiary, amount, currency, fraud score, outcome, transaction ref, timestamp (UTC). Retention: 7 years. Logs must be immutable. |
| **Regulatory Linkage** | PSD2 Art. 97 (SCA), BSA/AML (CTR for $10K+), GDPR Art. 6 (lawful processing) |
| **Dependencies** | US-PAY-038 (Beneficiary Management), US-AUTH-012 (Step-Up MFA), Core Banking Transfer API, Fraud Detection Engine |
| **Assumptions** | Core banking API supports real-time debit/credit; fraud engine responds within 200ms |
| **UI/UX Notes** | Transfer form: source account dropdown, beneficiary dropdown (with last-used sorting), amount field with currency, optional reference. Confirmation screen: reference number, amount, beneficiary name, timestamp. Error states: insufficient funds, limit exceeded, fraud hold. WCAG 2.1 AA compliant. |
| **API Contract** | POST /v1/transfers — Request: {sourceAccount, beneficiaryId, amount, currency, reference} — Response: {transactionId, status, timestamp, updatedBalance} — Auth: Bearer token + MFA token |
| **Definition of Done** | ✅ Code complete and peer-reviewed ✅ Unit tests ≥ 80% coverage ✅ Integration tests passing ✅ Security scan (SAST/DAST) — zero critical/high ✅ MFA flow tested (positive + bypass attempts) ✅ Fraud screening integration tested ✅ Audit logging verified (all events captured) ✅ Performance: p95 < 500ms at 500 concurrent users ✅ Accessibility: WCAG 2.1 AA verified ✅ Compliance criteria verified ✅ Deployed to staging |

---

## Usage Guidelines

1. **Always use the standard format** (PM-BP-015): "As a [specific role], I want [specific capability], so that [specific benefit]"
2. **Include security and compliance acceptance criteria** (PM-BP-016) as dedicated sections — don't bury them in functional criteria
3. **Tag every story with data sensitivity** (PM-BP-017) — this drives encryption, access control, audit, and testing requirements
4. **Specify auth requirements explicitly** (PM-BP-018) — auth level and authorization scope per story, not assumed
5. **Include audit logging requirements** (PM-BP-019) — specify events, fields, and retention for every story touching sensitive data or financial transactions
6. **One story = one testable capability** — if you need multiple test sessions, split the story
7. **Acceptance criteria must be testable** — use Given/When/Then with specific values
8. **Regulatory linkage is mandatory** — "None applicable" is valid, empty is not
