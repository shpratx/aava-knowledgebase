# Guardrails for Functional Requirements

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-SG-001 | Every FR must have a unique ID following FR-[Module]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| FR-SG-002 | Every FR must have a title, description, and at least one acceptance criterion | Reject if any of these three fields are empty |
| FR-SG-003 | One FR = one capability. Requirements containing multiple independent behaviors must be split | Flag if description contains multiple "and" joining distinct actions |
| FR-SG-004 | Every FR must be linked to a parent epic or deliverable | Reject if Epic/Deliverable field is empty |
| FR-SG-005 | Every FR must specify at least one actor | Reject if Actor(s) field is empty |
| FR-SG-006 | Every FR must include Main Flow with numbered steps | Reject if Main Flow is empty or unnumbered |
| FR-SG-007 | Every FR must include at least one Exception Flow | Flag if Exception Flows is empty — banking functions always have failure scenarios |

---

## 2. Data & Privacy Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-DG-001 | Every FR must specify Data Classification (Public/Internal/Confidential/Restricted) | Reject if Data Classification is empty |
| FR-DG-002 | FRs involving customer PII must state the purpose and legal basis for data processing | Reject if PII is involved and no legal basis is documented |
| FR-DG-003 | FRs must not require storage of sensitive data beyond what is necessary for the function (data minimization) | Flag if input data includes fields not justified by the business rules or flows |
| FR-DG-004 | FRs must not expose PII in URLs, logs, error messages, or client-side storage | Reject if flows describe PII in URLs or unprotected outputs |
| FR-DG-005 | FRs involving card data (PAN, CVV, expiry) must reference PCI-DSS scope | Reject if card data is mentioned without PCI-DSS reference |
| FR-DG-006 | FRs must specify data retention expectations or reference the applicable retention policy | Flag if no retention guidance is provided for stored data |

---

## 3. Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-SecG-001 | Every FR involving financial transactions must require authentication as a precondition | Reject if financial transaction FR has no authentication precondition |
| FR-SecG-002 | Every FR involving fund movement must require MFA/step-up authentication | Reject if fund transfer/payment FR does not reference MFA |
| FR-SecG-003 | FRs must not describe hardcoded credentials, API keys, or secrets in flows | Reject immediately — no exceptions |
| FR-SecG-004 | FRs must not bypass or weaken existing security controls | Reject if flows describe circumventing auth, authorization, or audit |
| FR-SecG-005 | FRs involving privileged operations must specify authorization level required | Reject if privileged operation has no authorization specification |
| FR-SecG-006 | FRs must specify input validation requirements for all user-supplied data | Flag if Input Data fields lack validation constraints |
| FR-SecG-007 | Error messages described in Exception Flows must not expose internal system details | Flag if exception flows describe technical error details shown to users |

---

## 4. Regulatory & Compliance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-RG-001 | Every FR must have the Regulatory Linkage field populated — "None applicable" is valid, empty is not | Reject if Regulatory Linkage is empty |
| FR-RG-002 | FRs involving payments must reference applicable payment regulations (PSD2, local payment laws) | Flag if payment FR has no regulatory reference |
| FR-RG-003 | FRs involving customer onboarding must reference KYC/AML requirements | Flag if onboarding FR has no KYC/AML reference |
| FR-RG-004 | FRs involving cross-border transactions must reference sanctions screening | Reject if cross-border FR has no sanctions reference |
| FR-RG-005 | FRs involving financial reporting must reference SOX/Basel requirements | Flag if reporting FR has no SOX/Basel reference |
| FR-RG-006 | FRs must not describe functionality that contradicts known regulatory requirements | Reject if FR conflicts with documented compliance requirements |
| FR-RG-007 | FRs involving transactions above regulatory thresholds must include reporting triggers (e.g., CTR for $10K+) | Flag if threshold transactions have no reporting reference |

---

## 5. Audit Trail & Logging Guardrails (ref: PM-BP-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-AG-001 | FRs involving financial transactions must include audit logging in the flow (who, what, when, outcome) | Reject if financial transaction flow has no logging step |
| FR-AG-002 | FRs involving data modification must capture before/after values in audit trail | Flag if data modification flow has no audit capture |
| FR-AG-003 | Every FR must be traceable to a business source (stakeholder, regulation, initiative) | Reject if Source field is empty |
| FR-AG-004 | FRs must reference related NFR, CR, and SR requirements where applicable | Flag if cross-references are empty for Critical/High priority FRs |
| FR-AG-005 | FRs involving approval workflows must log the approver, timestamp, and decision | Reject if approval flow has no audit logging |
| FR-AG-006 | Audit Trail Requirements field must specify: events to log, fields to capture, and retention period | Reject if Audit Trail Requirements field is empty for Critical/High FRs |
| FR-AG-007 | Audit logs must not contain sensitive data in clear text — PII must be masked, passwords/card numbers must never be logged | Reject if audit specification includes unmasked sensitive data |
| FR-AG-008 | Audit log retention must align with regulatory requirements (7 years financial, 3 years access) | Flag if retention period is below regulatory minimum |
| FR-AG-009 | FRs must specify that audit logs are immutable and stored separately from application data | Flag if log immutability is not specified |
| FR-AG-010 | Every step in Main Flow and Exception Flow that involves a decision, transaction, or data access must have a corresponding audit event | Flag if significant flow steps lack audit events |

---

## 6. Data Classification Guardrails (ref: PM-BP-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-DCG-001 | Every FR must specify Data Classification (Public/Internal/Confidential/Restricted) | Reject if Data Classification is empty |
| FR-DCG-002 | Data Classification Justification must be provided — reference data policy or regulation | Flag if justification is empty |
| FR-DCG-003 | Classification must be based on the most sensitive data the FR handles — if mixed, classify at highest level | Flag if classification appears lower than data sensitivity warrants |
| FR-DCG-004 | FRs classified as Restricted must have corresponding Security Requirements (SR) cross-referenced | Reject if Restricted FR has no SR cross-reference |
| FR-DCG-005 | FRs classified as Confidential or Restricted must specify encryption requirements for input/output data | Flag if no encryption reference for sensitive data |
| FR-DCG-006 | Data classification must be reviewed when FR scope changes — reclassify if data sensitivity shifts | Process guardrail — enforce via change review |

---

## 7. DR/BCP Guardrails (ref: PM-BP-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-DRG-001 | Every FR must specify a Recovery Priority tier (Tier 1/2/3) | Flag if Recovery Priority is empty |
| FR-DRG-002 | Tier 1 FRs must specify Degraded Mode Behavior — what happens when a dependency is unavailable | Reject if Tier 1 FR has no degraded mode defined |
| FR-DRG-003 | FRs involving financial transactions must specify in-flight transaction handling during failover (complete, rollback, or queue) | Reject if financial FR has no failover transaction handling |
| FR-DRG-004 | FRs must not auto-approve or bypass security controls during degraded mode | Reject if degraded mode bypasses security (e.g., skipping fraud check) |
| FR-DRG-005 | Exception Flows must include system unavailability scenarios for each integration point | Flag if integration-dependent FR has no unavailability exception flow |
| FR-DRG-006 | Degraded mode must specify the user-facing message/experience during partial failure | Flag if no user communication is defined for degraded mode |
| FR-DRG-007 | FRs must specify data consistency expectations on recovery — no duplicate transactions, no data loss for committed operations | Reject if recovery data consistency is not addressed for financial FRs |

---

## 8. Quality & Testability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-QG-001 | Acceptance criteria must use Given/When/Then format | Flag if acceptance criteria are not in structured format |
| FR-QG-002 | Acceptance criteria must include specific, measurable values (not "quickly", "properly", "correctly") | Reject if criteria contain subjective/vague language |
| FR-QG-003 | Every FR must have at least one positive (happy path) and one negative (rejection) acceptance criterion | Flag if only happy path criteria exist |
| FR-QG-004 | FRs must not use ambiguous language: "may", "might", "could", "optionally" for mandatory behaviors | Flag ambiguous language in mandatory requirements |
| FR-QG-005 | FRs must not describe implementation details (specific technologies, database tables, API endpoints) | Flag if description contains implementation specifics |
| FR-QG-006 | Business rules must include specific values/thresholds, not placeholders like "TBD" or "to be confirmed" | Reject if business rules contain unresolved placeholders |

---

## 9. Integration Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| FR-IG-001 | FRs involving external system integration must specify the integration point and expected behavior on failure | Reject if integration point is listed without failure handling |
| FR-IG-002 | FRs must not assume external system availability — exception flows must cover timeout and unavailability | Flag if no timeout/unavailability exception flow exists for integrations |
| FR-IG-003 | FRs involving third-party services must reference the security review status of that third party | Flag if third-party integration has no security review reference |
| FR-IG-004 | FRs involving asynchronous operations must specify the expected behavior for delayed responses and retries | Flag if async operations have no delay/retry handling |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | FR cannot proceed to development until fixed | Missing ID, no acceptance criteria, hardcoded secrets, no auth for financial ops |
| **Flag** | FR can proceed but must be addressed before sprint completion | Missing cross-references, vague language, no negative test criteria |
| **Review** | FR requires additional review by specified role | Compliance officer for regulatory FRs, security engineer for auth FRs |
