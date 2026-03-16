# Best Practices for Composing Functional Requirements

---

## 1. Writing the Requirement ID & Title

**Best Practices:**
- Use consistent naming: `FR-[Module]-[Seq]` (e.g., FR-PAY-001, FR-LND-012, FR-ONB-003)
- Module codes should be standardized across the organization (PAY=Payments, LND=Lending, ONB=Onboarding, TRE=Treasury, CRD=Cards, CUS=Customer, AUTH=Authentication, RPT=Reporting)
- Titles should be action-oriented and concise: "Initiate Domestic Fund Transfer" not "The system should be able to do fund transfers"
- Avoid technical jargon in titles — they should be understandable by business stakeholders

**Common Mistakes:**
- Generic IDs like FR-001 without module context
- Titles that describe implementation ("Call REST API to transfer funds") instead of capability ("Initiate Fund Transfer")

---

## 2. Writing the Description

**Best Practices:**
- Start with "The system shall..." for mandatory requirements or "The system should..." for desirable ones
- One requirement = one capability. If you use "and" to join two distinct behaviors, split into two requirements
- Be specific about channels: "via mobile and web banking" not "via digital channels"
- Specify the user role explicitly: "authenticated retail customers" not "users"
- Include what the system does, not how it does it (avoid implementation details)
- State the business intent — why this capability exists

**Formula:**
> The system shall [allow/enable/prevent/validate/calculate/display] [specific action] for [specific actor] [via specific channel] [under specific conditions].

**Good Example:**
> The system shall allow authenticated retail customers to initiate domestic fund transfers to registered beneficiaries via mobile and web banking channels, subject to daily transfer limits and real-time fraud screening.

**Bad Example:**
> The system should let users transfer money.

---

## 3. Defining Actors

**Best Practices:**
- List all human actors by their banking role (Retail Customer, Relationship Manager, Branch Teller, Compliance Officer, System Administrator)
- List all system actors (Core Banking System, Fraud Detection Engine, Payment Gateway, Notification Service)
- Distinguish between primary actors (who triggers the action) and secondary actors (who participates)
- Use role names consistent with your organization's role definitions and RBAC model

**Banking-Specific Roles to Consider:**
- Customer-facing: Retail Customer, Corporate Customer, Premium Customer
- Operations: Branch Teller, Operations Officer, Trade Finance Officer
- Risk & Compliance: Compliance Officer, Risk Analyst, AML Analyst, Fraud Analyst
- Management: Branch Manager, Relationship Manager, Product Owner
- Technical: System Administrator, DBA, DevOps Engineer
- External: Regulator, Auditor, Third-Party Provider

---

## 4. Writing Preconditions

**Best Practices:**
- State authentication and authorization status explicitly
- Include data prerequisites (e.g., "beneficiary is registered and active")
- Include system state prerequisites (e.g., "core banking system is available")
- Include business prerequisites (e.g., "customer has completed KYC")
- Each precondition should be independently verifiable
- Reference related requirements that satisfy preconditions (e.g., "Customer is authenticated — see FR-AUTH-003")

**Banking-Specific Preconditions to Always Consider:**
- Authentication level (basic auth, MFA, step-up auth)
- KYC/AML status of the customer
- Account status (active, dormant, frozen, closed)
- Regulatory eligibility (sanctions screening passed)
- System availability (core banking, payment network)
- Business hours / cut-off times

---

## 5. Writing Flows (Main, Alternative, Exception)

**Main Flow Best Practices:**
- Number each step sequentially
- Each step should describe one action by one actor
- Use present tense: "System validates balance" not "System will validate balance"
- Include validation steps explicitly — don't assume they happen
- End with the successful outcome and confirmation

**Alternative Flow Best Practices:**
- Start with the condition that triggers the alternative: "If [condition], then..."
- Reference the main flow step where the branch occurs
- Alternative flows are valid paths, not errors — they lead to a successful (but different) outcome
- Common banking alternatives: scheduled vs. immediate, different approval paths based on amount, different channels

**Exception Flow Best Practices:**
- Start with the failure condition: "If [failure], then..."
- Describe the system's response to the failure
- Include user notification (what message, via which channel)
- Include escalation path if applicable
- Include rollback/compensation actions for partial failures
- Always end with a defined state (not left hanging)

**Banking-Specific Exception Scenarios to Always Consider:**
- Insufficient funds
- Daily/transaction limit exceeded
- Fraud detection flag
- Sanctions screening hit
- System timeout (core banking, payment network)
- Duplicate transaction detection
- Account frozen/blocked
- Beneficiary validation failure
- Regulatory cut-off time exceeded

---

## 6. Defining Business Rules

**Best Practices:**
- State rules as clear, unambiguous conditions with specific values
- Include the source of the rule (regulation, policy, product definition)
- Separate hard rules (regulatory, non-negotiable) from soft rules (configurable, business-driven)
- Include effective dates if rules change over time
- Rules should be testable — avoid subjective language

**Formula:**
> [When condition], [the system shall/shall not] [action], [with specific threshold/value]. Source: [regulation/policy].

**Banking Business Rules Examples:**
- "Daily domestic transfer limit: $50,000 for retail, $500,000 for premium. Source: Product Policy PP-2024-012"
- "Transactions above $10,000 require CTR filing. Source: BSA/AML Regulation"
- "International transfers require beneficiary bank SWIFT/BIC validation. Source: SWIFT Standards"
- "Dormant accounts (no activity >12 months) require reactivation before transactions. Source: Internal Policy"

---

## 7. Specifying Input/Output Data

**Best Practices:**
- List every field with its data type, format, and constraints
- Specify mandatory vs. optional fields
- Include validation rules per field (min/max length, regex pattern, allowed values)
- Specify data classification for each field (especially PII)
- For outputs, specify the format and destination (screen, API response, notification, report)

**Format:**
> Field Name | Type | Format | Constraints | Mandatory | Classification

**Example:**
> Amount | Decimal | 2 decimal places | Min: 0.01, Max: 999,999,999.99 | Yes | Confidential
> Beneficiary Reference | String | Alphanumeric | Max 140 chars, no special chars except -/. | No | Internal

---

## 8. Writing Acceptance Criteria

**Best Practices:**
- Use Given/When/Then format consistently
- Each acceptance criterion should test one specific behavior
- Include positive (happy path) and negative (rejection) criteria
- Include boundary conditions (limits, edge cases)
- Make criteria measurable — include specific values, times, counts
- Include performance expectations where relevant

**Banking-Specific Acceptance Criteria Patterns:**
```
# Happy path
Given an authenticated customer with sufficient balance and active beneficiary,
When they submit a valid domestic transfer of $5,000,
Then the transfer completes within 5 seconds and a confirmation with reference number is displayed.

# Limit enforcement
Given a customer who has already transferred $45,000 today (limit: $50,000),
When they attempt to transfer $10,000,
Then the transfer is rejected with message "Daily limit exceeded. Remaining: $5,000."

# Fraud detection
Given a transfer flagged by the fraud detection engine,
When the fraud score exceeds the threshold,
Then the transfer is held, the customer is notified, and the fraud team receives an alert.

# Regulatory
Given a transfer of $10,000 or more,
When the transfer is completed,
Then a Currency Transaction Report (CTR) is automatically generated.
```

---

## 9. Specifying Data Classification (ref: PM-BP-008)

**Best Practices:**
- Assign a data classification to every FR based on the most sensitive data it handles
- Use the four-level scheme consistently: Public, Internal, Confidential, Restricted
- Classification drives downstream decisions — encryption, access control, audit depth, backup frequency, masking in non-prod
- Document the justification for the classification — reference the data policy or regulation
- If an FR handles data at multiple levels, classify at the highest level
- Review classification when requirements change — data sensitivity can shift

**Classification Decision Guide:**

| Ask This Question | If Yes → Classification |
|---|---|
| Does it involve card data (PAN, CVV), biometrics, encryption keys, or passwords? | Restricted |
| Does it involve customer PII, account balances, transaction details, or financial records? | Confidential |
| Does it involve internal policies, employee data, system configs, or non-public business data? | Internal |
| Is the data already publicly available (published rates, branch locations, product info)? | Public |

**Common Mistakes:**
- Defaulting everything to "Confidential" without analysis — over-classification wastes resources
- Under-classifying PII as "Internal" — regulatory penalties apply
- Not reclassifying when data is aggregated or anonymized

---

## 10. Specifying Audit Trail & Logging Requirements (ref: PM-BP-009)

**Best Practices:**
- Every FR involving financial transactions, data modifications, or access to sensitive data must specify what gets logged
- Define the audit event for each significant step in the Main Flow and Exception Flow
- Specify the exact fields to capture: who (user ID, role), what (action, resource), when (timestamp UTC), where (IP, device, channel), outcome (success/failure), and correlation ID
- Specify before/after values for data modification events
- State the retention period aligned with regulatory requirements (typically 7 years for financial, 3 years for access)
- Logs must not contain sensitive data in clear text — mask PII, never log passwords or full card numbers
- Specify that audit logs must be immutable and stored separately from application data

**What to Log per FR Type:**

| FR Type | Audit Events | Retention |
|---|---|---|
| Financial transaction (transfer, payment) | Initiation, validation, fraud check, execution, confirmation, failure | 7 years |
| Data access (account inquiry, statement) | Who accessed, which account, what data, when | 7 years |
| Data modification (profile update, beneficiary change) | Before/after values, who changed, approval chain | 7 years |
| Authentication event | Login, logout, MFA, lockout, password change | 3 years |
| Administrative action | Config change, role assignment, system override | 7 years |
| Approval workflow | Request, approval/rejection, approver, timestamp | 7 years |

**Common Mistakes:**
- Logging everything without structure — creates noise, makes forensic analysis difficult
- Logging sensitive data in clear text (full PAN, passwords, PII)
- Not specifying retention — logs deleted too early or kept forever
- Audit logging as an afterthought — must be designed into the flow from the start

---

## 11. Specifying DR/BCP Requirements (ref: PM-BP-010)

**Best Practices:**
- Every FR must specify its Recovery Priority tier (Tier 1/2/3) based on business criticality
- Define degraded-mode behavior — what happens when a dependency is unavailable
- Specify data recovery expectations — can any data be lost? What's the acceptable window?
- Define the expected behavior during failover — does the function continue, queue, or gracefully reject?
- Specify rollback/compensation requirements for in-flight transactions during failure
- Consider the user experience during DR — what message does the customer see?

**Recovery Priority Assignment:**

| Tier | RTO | RPO | Criteria | Banking Examples |
|---|---|---|---|---|
| Tier 1 | < 1 hour | < 15 min | Revenue-critical, regulatory-mandated, customer-facing financial ops | Fund transfers, payments, card authorization, core banking |
| Tier 2 | < 4 hours | < 1 hour | Important business functions, customer-facing non-financial | Account inquiry, statement generation, beneficiary management |
| Tier 3 | < 24 hours | < 4 hours | Internal operations, non-time-sensitive | Reporting, batch processing, internal admin tools |

**Degraded Mode Patterns:**

| Dependency Failure | Degraded Mode Options |
|---|---|
| Core banking unavailable | Queue transactions for retry; show cached balance (read-only); reject new transactions with ETA |
| Fraud engine unavailable | Hold transactions for manual review (never auto-approve); allow low-risk transactions below threshold |
| Payment network unavailable | Queue for next available window; offer alternative payment method; notify customer of delay |
| Notification service unavailable | Complete transaction, queue notification for retry; do not block transaction for notification failure |
| Database read-replica unavailable | Route to primary (accept performance degradation); serve cached data with staleness indicator |

**Common Mistakes:**
- Assuming all functions need Tier 1 recovery — over-engineering DR is expensive
- Not defining degraded mode — system fails completely instead of gracefully
- Not testing DR behavior — untested DR plans fail when needed
- Ignoring in-flight transaction handling — leads to duplicates or lost transactions on recovery

---

## 12. Cross-Referencing

**Best Practices:**
- Link to Non-Functional Requirements that apply (e.g., NFR-PERF-003 for response time)
- Link to Compliance Requirements that govern this function (e.g., CR-PSD2-001 for SCA)
- Link to Security Requirements that protect this function (e.g., SR-AUTH-003 for MFA)
- Link to dependent/related Functional Requirements
- Maintain a traceability matrix mapping FR → NFR → CR → SR

---

## 13. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| "The system should handle transfers" | Too vague, untestable | Specify who, what, where, when, how much |
| Combining multiple capabilities in one FR | Untestable, hard to track | One FR = one capability |
| Implementation-specific language | Constrains design, becomes outdated | Describe what, not how |
| Missing exception flows | Gaps in error handling | Cover every failure scenario |
| No data classification | Security/compliance gaps | Classify every data field |
| Acceptance criteria without values | Untestable | Include specific thresholds |
| No regulatory linkage | Compliance gaps in audit | Always state regulation or "None applicable" |
| No audit trail specification | Undetectable breaches, failed audits | Define audit events for every significant flow step |
| No DR/BCP consideration | Total failure instead of graceful degradation | Define recovery tier and degraded-mode behavior |
| Logging sensitive data in clear text | Data breach via logs | Mask PII, never log passwords or full card numbers |
