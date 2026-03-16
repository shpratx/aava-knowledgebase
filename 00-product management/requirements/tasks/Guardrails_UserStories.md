# Guardrails for User Stories

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| US-SG-001 | Every story must have a unique ID following US-[Module]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| US-SG-002 | Every story must follow the "As a [role], I want [capability], so that [benefit]" format (ref: PM-BP-015) | Reject if story statement is missing or does not follow the standard format |
| US-SG-003 | The role must be a specific banking persona — "user", "someone", "admin" without context are not acceptable | Reject if role is generic or unspecified |
| US-SG-004 | The "so that" benefit clause is mandatory — it drives prioritization and design | Reject if benefit clause is missing |
| US-SG-005 | One story = one testable capability — stories combining multiple independent actions must be split | Flag if story description contains multiple independent behaviors joined by "and" |
| US-SG-006 | Every story must be linked to a parent Epic | Reject if Epic field is empty |
| US-SG-007 | Every story must be completable within one sprint — stories exceeding sprint capacity must be split | Flag if Story Points exceed team's sprint velocity threshold |
| US-SG-008 | Every story must have a Regulatory Linkage field populated — "None applicable" is valid, empty is not | Reject if Regulatory Linkage is empty |

---

## 2. Data Privacy & Purpose Guardrails (ref: PM-GR-014)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-DPG-001 | Stories accessing customer data must specify the purpose of data access and the legal basis for processing (ref: PM-GR-014) | Reject if story accesses customer data and no purpose/legal basis is documented |
| US-DPG-002 | Legal basis must reference a specific GDPR Art. 6 ground: consent, contract, legal obligation, vital interest, public task, or legitimate interest | Reject if legal basis is vague ("we need it") instead of citing a specific ground |
| US-DPG-003 | Stories collecting new customer data must justify each field against the data minimization principle — only data necessary for the stated purpose | Reject if data fields cannot be justified against the stated purpose |
| US-DPG-004 | Stories must not collect or process customer data beyond what is specified in the purpose statement | Reject if data processing scope exceeds stated purpose |
| US-DPG-005 | Stories involving special category data (health, biometrics, race, political opinions, religious beliefs) must have explicit consent or legal exemption documented | Reject if special category data has no explicit legal basis |
| US-DPG-006 | Stories sharing customer data with third parties must specify: which data, to whom, for what purpose, under what agreement (DPA), and the legal basis | Reject if third-party data sharing lacks any of these elements |
| US-DPG-007 | Stories must specify data retention expectations — how long the data created/accessed will be retained and the basis for that period | Flag if no retention expectation is stated |
| US-DPG-008 | Stories involving customer profiling or automated decision-making must document the right to human review per GDPR Art. 22 | Reject if automated decision-making has no human review provision |

---

## 3. Authentication & Authorization Guardrails (ref: PM-GR-015)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-AAG-001 | Every story must specify the Authentication Requirement — the exact auth level required (None/Basic/MFA/Step-Up MFA) (ref: PM-BP-018) | Reject if Authentication Requirement is empty |
| US-AAG-002 | Every story must specify the Authorization Requirement — role, scope, and resource-level checks (ref: PM-BP-018) | Reject if Authorization Requirement is empty |
| US-AAG-003 | Authentication and authorization must follow the principle of least privilege — users get the minimum access needed for the specific action (ref: PM-GR-015) | Reject if auth grants broader access than the story requires |
| US-AAG-004 | Stories involving financial transactions must require MFA or Step-Up MFA — basic session auth is insufficient | Reject if financial transaction story has auth level below MFA |
| US-AAG-005 | Stories involving profile modifications (address, email, phone, beneficiary changes) must require re-authentication or Step-Up MFA | Reject if profile modification story has no re-authentication |
| US-AAG-006 | Authorization must be enforced at the resource level — role-only checks are insufficient (prevents IDOR) | Reject if authorization is role-only without resource-level validation |
| US-AAG-007 | Stories involving access to another customer's data (e.g., branch teller, RM) must specify the authorization model and audit requirements | Reject if cross-customer access has no authorization model |
| US-AAG-008 | Stories involving administrative or privileged operations must require MFA + just-in-time access with approval workflow | Reject if admin story has no JIT access or approval requirement |
| US-AAG-009 | Stories must specify what happens on authentication failure: error message (no information leakage), lockout threshold, and alert trigger | Flag if auth failure behavior is not specified |
| US-AAG-010 | Stories must specify session timeout behavior for the operation — 15 minutes inactivity for sensitive operations, 8 hours absolute | Flag if session timeout is not specified for sensitive operations |
| US-AAG-011 | Stories must not grant "all accounts" or "all customers" access unless the role explicitly requires it and is documented with justification | Reject if broad access is granted without documented justification |

---

## 4. Security Control Guardrails (ref: PM-GR-016)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-SCG-001 | Stories must not bypass, weaken, or disable existing security controls (ref: PM-GR-016) | Reject immediately — no exceptions |
| US-SCG-002 | Stories must not introduce backdoors, hardcoded credentials, debug endpoints, or test accounts in production code | Reject immediately — no exceptions |
| US-SCG-003 | Stories must not reduce the authentication level for an operation that already has a defined auth requirement | Reject if story lowers auth level for an existing operation |
| US-SCG-004 | Stories must not expose sensitive data in URLs, query parameters, error messages, client-side storage, or browser history | Reject if story design exposes sensitive data in any of these locations |
| US-SCG-005 | Stories must not log sensitive data in clear text (passwords, PAN, CVV, OTP, tokens, encryption keys) | Reject if audit logging specification includes unmasked sensitive data |
| US-SCG-006 | Stories must not disable or reduce audit logging for any operation that currently has logging | Reject if story reduces existing audit coverage |
| US-SCG-007 | Stories must not introduce direct database access from client-side code or bypass the API layer | Reject if architecture bypasses established security layers |
| US-SCG-008 | Stories must not use deprecated or insecure protocols, algorithms, or libraries (SSL, TLS <1.2, MD5, SHA-1 for security, known vulnerable dependencies) | Reject if insecure technology is specified |
| US-SCG-009 | Stories must not store sensitive data in client-side storage (localStorage, sessionStorage, cookies without Secure/HttpOnly flags) | Reject if client-side sensitive data storage is proposed |
| US-SCG-010 | Stories involving file uploads must validate file type, size, and content — executable uploads must be prohibited | Reject if file upload has no validation specification |
| US-SCG-011 | Stories must specify server-side input validation for all user-supplied data using whitelist approach | Reject if input validation is client-side only or uses blacklist approach |
| US-SCG-012 | If a story requires a temporary security exception, it must go through the formal exception process (CISO + Compliance approval, compensating controls, max 6-month expiry) | Reject if security exception bypasses formal process |

---

## 5. Compliance Review Guardrails (ref: PM-GR-017)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-CRG-001 | Stories linked to regulations must be reviewed by the Compliance Officer or legal team before development begins (ref: PM-GR-017) | Reject if regulated story has no compliance review completed or scheduled |
| US-CRG-002 | Stories involving payment processing must be reviewed against PCI-DSS requirements | Reject if payment story has no PCI-DSS review |
| US-CRG-003 | Stories involving customer onboarding or identity verification must be reviewed against KYC/AML requirements | Reject if onboarding story has no KYC/AML review |
| US-CRG-004 | Stories involving customer data collection or processing must be reviewed against GDPR/data protection requirements | Reject if data processing story has no GDPR review |
| US-CRG-005 | Stories involving financial reporting or calculations must be reviewed against SOX requirements | Flag if financial reporting story has no SOX review |
| US-CRG-006 | Stories involving cross-border data or transactions must be reviewed by legal for jurisdictional compliance | Reject if cross-border story has no legal review |
| US-CRG-007 | Compliance review outcomes must be documented: Approved / Conditionally Approved / Rejected, with reviewer name, date, and conditions | Reject if compliance review outcome is not recorded |
| US-CRG-008 | Conditional compliance approvals must have conditions tracked to closure before the story is marked Done | Reject if story is marked Done with open compliance conditions |
| US-CRG-009 | Stories that modify existing compliance controls must undergo impact assessment before development | Reject if compliance control modification has no impact assessment |

---

## 6. Definition of Done Guardrails (ref: PM-GR-018)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-DDG-001 | Every story must have a Definition of Done checklist that includes security criteria (ref: PM-GR-018) | Reject if DoD is missing or has no security items |
| US-DDG-002 | DoD must include: code peer-reviewed, unit tests ≥ 80% coverage, SAST scan passed (zero critical/high) | Reject if any of these baseline items are missing from DoD |
| US-DDG-003 | Stories tagged Confidential or Restricted must include in DoD: DAST scan passed, auth flow tested (positive + bypass), input validation tested | Reject if sensitive story DoD is missing these items |
| US-DDG-004 | Stories with audit logging requirements must include in DoD: audit events verified, sensitive data not in logs verified | Reject if audit story DoD is missing log verification |
| US-DDG-005 | Stories with compliance acceptance criteria must include in DoD: compliance criteria verified, compliance review conditions closed | Reject if regulated story DoD is missing compliance verification |
| US-DDG-006 | Stories with UI components must include in DoD: WCAG 2.1 AA accessibility verified | Flag if UI story DoD is missing accessibility check |
| US-DDG-007 | Customer-facing stories must include in DoD: performance tested against NFR targets | Flag if customer-facing story DoD is missing performance verification |
| US-DDG-008 | DoD items must be verifiable (yes/no) — vague items like "code is good quality" are not acceptable | Reject if DoD contains non-verifiable items |
| US-DDG-009 | A story cannot be marked Done if any DoD item is not checked off — partial completion is not Done | Reject if story is closed with unchecked DoD items |

---

## 7. Data Sensitivity Guardrails (ref: PM-BP-017)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-DSG-001 | Every story must have a Data Sensitivity Tag (Public/Internal/Confidential/Restricted) (ref: PM-BP-017) | Reject if Data Sensitivity Tag is empty |
| US-DSG-002 | The tag must reflect the highest sensitivity of any data field the story touches | Flag if tag appears lower than the data fields warrant |
| US-DSG-003 | Stories tagged Restricted must have Security Acceptance Criteria — no exceptions | Reject if Restricted story has no security acceptance criteria |
| US-DSG-004 | Stories tagged Confidential must have Security Acceptance Criteria | Reject if Confidential story has no security acceptance criteria |
| US-DSG-005 | Stories tagged Restricted or Confidential must specify data masking requirements for non-production environments | Flag if no masking requirement is specified for sensitive stories |
| US-DSG-006 | Stories must list Data Fields Involved with per-field classification and CRUD operation | Flag if Data Fields Involved is empty for Confidential/Restricted stories |
| US-DSG-007 | Data Sensitivity Tag must be reviewed when story scope changes — adding a data field can change the classification | Process guardrail — enforce via story change review |

---

## 8. Acceptance Criteria Guardrails (ref: PM-BP-016)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-ACG-001 | Every story must have Acceptance Criteria in Given/When/Then format | Reject if acceptance criteria are missing or not in structured format |
| US-ACG-002 | Stories tagged Confidential or Restricted must have dedicated Security Acceptance Criteria section (ref: PM-BP-016) | Reject if sensitive story has no security acceptance criteria |
| US-ACG-003 | Stories linked to regulations must have dedicated Compliance Acceptance Criteria section (ref: PM-BP-016) | Reject if regulated story has no compliance acceptance criteria |
| US-ACG-004 | Every regulation cited in Regulatory Linkage must have at least one corresponding compliance acceptance criterion | Reject if a cited regulation has no acceptance criterion |
| US-ACG-005 | Acceptance criteria must include both positive (happy path) and negative (rejection/error/bypass) scenarios | Flag if only positive criteria exist |
| US-ACG-006 | Acceptance criteria must include specific, measurable values — no subjective language ("fast", "properly", "securely", "adequately") | Reject if criteria contain vague/subjective language |
| US-ACG-007 | Security acceptance criteria must include bypass-attempt tests, not just positive verification | Flag if security criteria only test the happy path |
| US-ACG-008 | Compliance acceptance criteria must cite the specific regulation article they satisfy | Flag if compliance criterion has no article reference |

---

## 9. Audit Logging Guardrails (ref: PM-BP-019)

| ID | Guardrail | Enforcement |
|---|---|---|
| US-ALG-001 | Stories involving financial transactions must specify audit logging requirements (ref: PM-BP-019) | Reject if financial transaction story has no audit logging specification |
| US-ALG-002 | Stories involving access to or modification of sensitive data (Confidential/Restricted) must specify audit logging | Reject if sensitive data story has no audit logging specification |
| US-ALG-003 | Audit logging must capture: who (user ID, role), what (action, resource), when (timestamp UTC), where (IP, device, channel), outcome (success/failure), and correlation ID | Reject if audit specification is missing mandatory fields |
| US-ALG-004 | Audit logging for data modifications must capture before/after values | Flag if data modification story has no before/after logging |
| US-ALG-005 | Audit log retention must be specified and aligned with regulatory requirements (7 years financial, 3 years access) | Reject if retention is not specified or below regulatory minimum |
| US-ALG-006 | Audit logs must not contain sensitive data in clear text — passwords, full PAN, CVV, OTP, tokens, keys must never be logged | Reject if log specification includes unmasked sensitive data |
| US-ALG-007 | Audit logs must be specified as immutable and stored separately from application data | Flag if log immutability is not specified |
| US-ALG-008 | Stories must not reduce or disable existing audit logging for any operation | Reject if story reduces existing audit coverage (same as US-SCG-006) |

---

## 10. Integration & API Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| US-IAG-001 | Stories involving API endpoints must specify the API contract (endpoint, method, request/response schema, auth) | Flag if API story has no contract specification |
| US-IAG-002 | Stories involving external system integration must specify timeout handling and failure behavior | Reject if integration story has no timeout/failure handling |
| US-IAG-003 | Stories must propagate correlation IDs across all integration points for traceability | Flag if correlation ID propagation is not specified |
| US-IAG-004 | Stories involving third-party APIs must specify rate limiting and circuit breaker behavior | Flag if third-party integration has no resilience specification |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Story cannot enter development until fixed | No purpose/legal basis for customer data access (PM-GR-014), auth grants broader access than needed (PM-GR-015), bypasses security controls (PM-GR-016), no compliance review for regulated story (PM-GR-017), DoD missing security criteria (PM-GR-018) |
| **Flag** | Story can enter development but must be addressed before Done | Missing negative test criteria, no data masking spec, no session timeout spec, no API contract |
| **Review** | Story requires review by specified role before development | Compliance Officer for regulated stories (PM-GR-017), Legal for cross-border, Security for Restricted data, DPO for new data collection |
| **Process** | Enforced via workflow, not automated | Data sensitivity tag review on scope change, compliance condition tracking |

---

## Quick Reference: Guardrail Triggers by Story Type

| Story Involves... | Triggered Guardrails |
|---|---|
| Customer data access | US-DPG-001→008 (purpose + legal basis mandatory, PM-GR-014) |
| Financial transaction | US-AAG-004 (MFA mandatory), US-ALG-001 (audit mandatory), US-SCG-001 (no security bypass) |
| Profile modification | US-AAG-005 (re-auth required), US-ALG-002 (audit mandatory) |
| Payment/card data | US-CRG-002 (PCI-DSS review), US-DSG-003 (security criteria mandatory) |
| Cross-border data | US-CRG-006 (legal review mandatory) |
| Regulated activity | US-CRG-001→009 (compliance review mandatory, PM-GR-017), US-ACG-003 (compliance criteria) |
| Restricted data | US-DSG-003 (security criteria), US-DDG-003 (DAST + auth testing in DoD) |
| Confidential data | US-DSG-004 (security criteria), US-DDG-003 (enhanced DoD) |
| Admin/privileged ops | US-AAG-008 (MFA + JIT), US-ALG-002 (audit mandatory) |
| Third-party integration | US-IAG-002→004 (timeout, circuit breaker, rate limiting) |
| UI component | US-DDG-006 (WCAG 2.1 AA), US-SCG-004 (no sensitive data exposure) |
| File upload | US-SCG-010 (validation mandatory) |
