# Functional Requirements Template

---

## Template Fields

| Field | Description |
|---|---|
| **Requirement ID** | FR-[Module]-[Seq] (e.g., FR-PAY-001) |
| **Title** | Short descriptive title |
| **Epic / Deliverable** | Parent epic or deliverable reference |
| **Business Domain** | e.g., Payments, Lending, Onboarding, Treasury, Cards |
| **Priority** | Critical / High / Medium / Low |
| **Source** | Business stakeholder, regulation, or initiative reference |
| **Description** | Detailed description of the functional capability |
| **Actor(s)** | User roles or systems involved (e.g., Relationship Manager, Core Banking System) |
| **Preconditions** | What must be true before this function executes |
| **Trigger** | Event or action that initiates the function |
| **Main Flow** | Step-by-step expected behavior: 1. … 2. … 3. … |
| **Alternative Flows** | Variations from the main flow |
| **Exception Flows** | Error scenarios and how the system should respond |
| **Business Rules** | Specific rules governing the behavior (e.g., transfer limit = $50,000/day for retail) |
| **Input Data** | Data fields required, with types and constraints |
| **Output Data** | Expected outputs, responses, or state changes |
| **Data Classification** | Public / Internal / Confidential / Restricted (ref: PM-BP-008) |
| **Data Classification Justification** | Why this classification level was assigned; regulatory or policy basis |
| **Audit Trail Requirements** | What must be logged for this function: events, data captured (who/what/when/outcome), retention period (ref: PM-BP-009) |
| **Audit Log Detail** | Specific fields to log: user ID, action, timestamp, source IP, before/after values, outcome, correlation ID |
| **DR/BCP Requirements** | Disaster recovery and business continuity needs for this function (ref: PM-BP-010) |
| **Degraded Mode Behavior** | How the function behaves when a dependency is unavailable (e.g., queue for retry, read-only mode, graceful rejection) |
| **Recovery Priority** | Recovery tier for this function: Tier 1 (< 1 hour) / Tier 2 (< 4 hours) / Tier 3 (< 24 hours) |
| **Integration Points** | Systems or APIs this requirement interacts with (e.g., Core Banking, SWIFT, Card Processor) |
| **Acceptance Criteria** | Measurable conditions for sign-off: Given [context], When [action], Then [outcome] |
| **Regulatory Linkage** | Applicable regulation(s) if any (e.g., PSD2 SCA, AML/KYC) |
| **Dependencies** | Other requirements, systems, or teams this depends on |
| **Assumptions** | Assumptions made during requirement definition |
| **Attachments** | Wireframes, process flows, reference documents |

---

## Example — Functional Requirement

| Field | Value |
|---|---|
| **Requirement ID** | FR-PAY-012 |
| **Title** | Initiate Domestic Fund Transfer |
| **Epic / Deliverable** | EP-PAY-003: Digital Payments Modernization |
| **Business Domain** | Payments |
| **Priority** | Critical |
| **Source** | Product Owner — Retail Banking |
| **Description** | The system shall allow authenticated retail customers to initiate domestic fund transfers to registered beneficiaries via the mobile and web banking channels. |
| **Actor(s)** | Retail Customer, Core Banking System, Fraud Detection Engine |
| **Preconditions** | Customer is authenticated with MFA; beneficiary is registered and active; source account has sufficient balance |
| **Trigger** | Customer selects "Transfer Funds" and submits the transfer form |
| **Main Flow** | 1. Customer selects source account and beneficiary 2. Enters amount and optional reference 3. System validates balance and daily limits 4. System performs real-time fraud check 5. System debits source and credits beneficiary 6. Confirmation displayed with transaction reference |
| **Alternative Flows** | If amount exceeds daily limit, prompt customer with remaining limit and option to schedule |
| **Exception Flows** | If fraud check flags transaction → hold transaction, notify customer, escalate to fraud team |
| **Business Rules** | Daily transfer limit: $50,000 retail / $500,000 premium; Cut-off time: 16:00 local for same-day |
| **Input Data** | Source account (dropdown), beneficiary (dropdown), amount (decimal, 2dp), reference (text, max 140 chars) |
| **Output Data** | Transaction reference number, status, timestamp, updated balance |
| **Data Classification** | Confidential |
| **Data Classification Justification** | Contains customer account balances and transaction details — classified per Data Policy DP-2024-003 |
| **Audit Trail Requirements** | Log every transfer attempt (success and failure): user ID, source account, beneficiary, amount, timestamp, fraud check result, outcome, transaction reference. Retain for 7 years per financial regulation. |
| **Audit Log Detail** | User ID, session ID, source IP, device fingerprint, action=FUND_TRANSFER, source account, beneficiary account, amount, currency, fraud score, outcome (success/held/rejected), transaction reference, timestamp (UTC) |
| **DR/BCP Requirements** | Fund transfer is Tier 1 critical — must be recoverable within 1 hour. In-flight transactions must be either completed or safely rolled back on failover. No duplicate transactions on recovery. |
| **Degraded Mode Behavior** | If core banking is unavailable: reject real-time transfers with user-friendly message and offer scheduled transfer option. If fraud engine is unavailable: hold transaction for manual review (do not auto-approve). |
| **Recovery Priority** | Tier 1 (< 1 hour) |
| **Integration Points** | Core Banking (T24), Fraud Detection (NICE Actimize), Notification Service |
| **Acceptance Criteria** | Given an authenticated customer with sufficient balance, When they submit a valid transfer, Then the transfer completes within 5 seconds and a confirmation with reference number is displayed |
| **Regulatory Linkage** | PSD2 — Strong Customer Authentication |
| **Dependencies** | FR-PAY-008 (Beneficiary Management), FR-AUTH-003 (MFA) |
| **Assumptions** | Real-time gross settlement available for domestic transfers |

---

## Usage Guidelines

1. **Every requirement must have a unique ID** following the FR-[Module]-[Seq] naming convention
2. **Regulatory linkage is mandatory** — even if "None applicable", explicitly state it
3. **Data classification must be specified** — this drives encryption, access control, and audit requirements
4. **Acceptance criteria must be testable** — use Given/When/Then format for clarity
5. **Cross-reference** to related Non-Functional (NFR), Compliance (CR), and Security (SR) requirements
6. **Review cadence**: Per sprint
7. **Approval**: Product Owner sign-off required
