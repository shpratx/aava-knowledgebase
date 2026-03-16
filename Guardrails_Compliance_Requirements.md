# Guardrails for Compliance Requirements

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-SG-001 | Every CR must have a unique ID following CR-[Regulation]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| CR-SG-002 | Every CR must cite the specific regulation, article/section, and enforcing body | Reject if Regulation or Article/Section is empty |
| CR-SG-003 | Every CR must have a clearly stated Compliance Obligation separate from the Description | Reject if Compliance Obligation is empty |
| CR-SG-004 | Every CR must specify a Compliance Owner (role, not individual) | Reject if Compliance Owner is empty |
| CR-SG-005 | Every CR must specify a Review Frequency | Reject if Review Frequency is empty |
| CR-SG-006 | Every CR must have testable Acceptance Criteria | Reject if Acceptance Criteria is empty or untestable |
| CR-SG-007 | Every CR must specify the Retention Period for data and evidence | Reject if Retention Period is empty |

---

## 2. Regulatory Citation Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-RCG-001 | Regulation version/year must be current — outdated regulation versions must be flagged | Flag if regulation version is not the latest |
| CR-RCG-002 | Article/section citation must be specific enough to locate the exact clause (not just "GDPR" or "PCI-DSS") | Reject if citation is regulation-name-only without article/section |
| CR-RCG-003 | If a local regulation implements an international standard, both must be cited | Flag if only one is cited when both apply |
| CR-RCG-004 | Regulatory amendments must trigger review of all CRs linked to that regulation within 48 hours | Process guardrail — enforce via workflow |
| CR-RCG-005 | CRs must not misquote or misinterpret regulatory text — legal/compliance review required | Review by Compliance Officer before approval |

---

## 3. Data & Privacy Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-DPG-001 | Every CR must specify the Data Involved and its Data Classification | Reject if Data Involved or Data Classification is empty |
| CR-DPG-002 | CRs involving PII must specify the legal basis for processing (consent, contract, legal obligation, legitimate interest) | Reject if PII is involved and no legal basis is stated |
| CR-DPG-003 | CRs involving cross-border data transfer must specify the transfer mechanism (adequacy decision, SCCs, BCRs) | Reject if cross-border transfer has no mechanism specified |
| CR-DPG-004 | CRs must not permit data retention beyond the regulatory/business maximum | Reject if retention period exceeds regulatory maximum without documented exemption |
| CR-DPG-005 | CRs involving data erasure must address conflicts with other retention obligations (e.g., GDPR erasure vs. AML retention) | Reject if erasure CR does not address retention conflicts |
| CR-DPG-006 | CRs must specify whether anonymization or pseudonymization is required and the method | Flag if data protection CR does not specify anonymization approach |

---

## 4. Geographic Scope Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-GSG-001 | Every CR must specify the Geographic Scope — which jurisdictions it applies to | Reject if Geographic Scope is empty |
| CR-GSG-002 | CRs with multi-jurisdiction scope must identify conflicting requirements across jurisdictions | Flag if multi-jurisdiction CR has no conflict analysis |
| CR-GSG-003 | CRs must account for extraterritorial reach of regulations (GDPR applies to EU data subjects globally, FATCA applies to US persons globally) | Flag if extraterritorial scope is not considered |
| CR-GSG-004 | Data localization requirements must be explicitly stated if applicable | Flag if jurisdiction has known data localization rules and CR doesn't address them |

---

## 5. Evidence & Audit Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-EAG-001 | Every CR must specify the Evidence Required to demonstrate compliance to an auditor | Reject if Evidence Required is empty |
| CR-EAG-002 | Every CR must specify Audit Trail Requirements (what is logged, format, retention) | Reject if Audit Trail Requirements is empty |
| CR-EAG-003 | Audit logs for compliance events must be tamper-proof and centrally stored | Reject if audit trail does not specify immutability |
| CR-EAG-004 | Evidence retention must meet or exceed the regulatory retention period | Reject if evidence retention is shorter than the regulatory requirement |
| CR-EAG-005 | Evidence must be independently verifiable — self-reported compliance without supporting artifacts is insufficient | Flag if evidence is self-attestation only |
| CR-EAG-006 | CRs must specify who has access to compliance evidence and under what conditions | Flag if evidence access control is not defined |

---

## 6. Reporting Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-RPG-001 | CRs with regulatory reporting obligations must specify: report name, content, frequency, format, recipient, and deadline | Reject if any reporting field is incomplete |
| CR-RPG-002 | Event-triggered reports (breach notification, SAR, CTR) must specify the triggering condition and maximum response time | Reject if trigger or response time is missing |
| CR-RPG-003 | Reporting processes must include data quality validation before submission | Flag if no data quality check is specified |
| CR-RPG-004 | Late or inaccurate report escalation process must be defined | Flag if no escalation process for reporting failures |
| CR-RPG-005 | Automated reporting is preferred over manual — manual reporting must justify why automation is not feasible | Flag if reporting is manual without justification |

---

## 7. Implementation Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-IMG-001 | Implementation approach must address both technical controls and operational processes | Flag if implementation is technical-only or process-only |
| CR-IMG-002 | Conflict resolution between regulations must be explicitly documented with legal review | Reject if conflicting regulations are identified but no resolution is documented |
| CR-IMG-003 | Implementation must not create new compliance risks while addressing the stated obligation | Review by Compliance Officer |
| CR-IMG-004 | Automation must be preferred for compliance controls — manual controls must have compensating monitoring | Flag if manual controls have no monitoring |
| CR-IMG-005 | Implementation approach must be reviewed by both legal/compliance and technical teams | Process guardrail — enforce via approval workflow |

---

## 8. Penalty & Risk Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-PRG-001 | Every CR must document the Penalties for Non-Compliance (fines, sanctions, license risk, reputational damage) | Reject if Penalties field is empty |
| CR-PRG-002 | CRs with penalties exceeding €1M or equivalent must be classified as Critical priority | Auto-escalate to Critical if penalty threshold is met |
| CR-PRG-003 | CRs classified as Critical must have Compliance Officer sign-off before development begins | Reject if Critical CR lacks Compliance Officer approval |
| CR-PRG-004 | Non-compliance issues identified during testing must block production deployment | Enforce as release gate |

---

## 9. Review & Maintenance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-RMG-001 | All CRs must be reviewed at minimum quarterly | Process guardrail — enforce via scheduled review |
| CR-RMG-002 | Regulatory changes must trigger impact assessment on all related CRs within 48 hours | Process guardrail — enforce via regulatory change monitoring |
| CR-RMG-003 | Deprecated or superseded regulations must be flagged and CRs updated within 30 days | Flag stale CRs |
| CR-RMG-004 | CR review must involve both Compliance Officer and the technical team responsible for implementation | Process guardrail — enforce via review workflow |
| CR-RMG-005 | Changes to CRs must be version-controlled with change history and rationale | Reject if CR changes have no version history |

---

## 10. DR/BCP Compliance Guardrails (ref: PM-BP-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-DRG-001 | Every CR must specify how compliance is maintained during DR/BCP events in the DR/BCP Compliance field | Flag if DR/BCP Compliance field is empty |
| CR-DRG-002 | Regulatory reporting deadlines remain in effect during outages — contingency reporting procedures must be specified | Reject if event-triggered reporting CR has no DR contingency |
| CR-DRG-003 | Data protection controls (encryption, access control, audit logging) must be maintained in DR environments | Reject if DR environment has weaker compliance controls |
| CR-DRG-004 | CRs governed by DORA must include explicit ICT business continuity requirements | Reject if DORA-scoped CR has no continuity requirements |
| CR-DRG-005 | Compliance evidence must be replicated to DR site and accessible during/after DR events | Flag if evidence availability during DR is not addressed |
| CR-DRG-006 | AML/sanctions screening must not be interrupted during DR — transaction monitoring must continue | Reject if AML CR permits monitoring gaps during DR |
| CR-DRG-007 | DR events that result in data exposure must trigger breach notification assessment per GDPR/local regulations | Flag if no breach assessment process is defined for DR events |

---

## 11. Data Classification Guardrails (ref: PM-BP-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-DCG-001 | Every CR must map the regulation's data categories to the organization's classification scheme (Public/Internal/Confidential/Restricted) | Flag if no classification mapping exists |
| CR-DCG-002 | CRs involving Restricted data (card data, biometrics, encryption keys) must have the strictest controls with no exceptions | Reject if Restricted data CR permits exceptions |
| CR-DCG-003 | CRs must specify different control requirements for different classification levels where the regulation allows tiering | Flag if one-size-fits-all controls are applied across all data levels |
| CR-DCG-004 | Data that cannot be stored post-authorization (PCI-DSS sensitive authentication data) must be explicitly flagged | Reject if sensitive auth data storage is not prohibited |
| CR-DCG-005 | Classification must be reviewed when regulation scope changes — new data types may require reclassification | Process guardrail — enforce via regulatory change review |

---

## 12. Cross-Reference Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| CR-XRG-001 | Every CR must link to the Functional Requirements it governs | Flag if Related Functional Req(s) is empty |
| CR-XRG-002 | CRs with security implications must link to corresponding Security Requirements | Flag if security-related CR has no SR cross-reference |
| CR-XRG-003 | CRs must not conflict with other CRs — conflicts must be identified and resolved | Review during CR creation |
| CR-XRG-004 | A traceability matrix (CR → FR → SR → Test Case) must be maintained | Process guardrail — enforce via tooling |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | CR cannot proceed until fixed | Missing regulation citation, no evidence requirements, no retention period, no penalties documented |
| **Flag** | CR can proceed but must be addressed before release | Missing cross-references, manual controls without monitoring, no conflict analysis |
| **Review** | CR requires additional review by specified role | Compliance Officer for all CRs, Legal for cross-border/conflict CRs |
| **Auto-Escalate** | CR priority automatically elevated | Penalties > €1M → Critical priority |
