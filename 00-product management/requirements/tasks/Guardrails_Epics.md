# Guardrails for Epics

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-SG-001 | Every epic must have a unique ID following EP-[Domain]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| EP-SG-002 | Every epic must have a Title framed around a business capability, not a technical component (ref: PM-BP-011) | Reject if title describes a technology or technical task (e.g., "Build API", "Database Migration") |
| EP-SG-003 | Every epic must have a Business Value Statement with quantified outcomes | Reject if Business Value Statement is empty or contains no measurable value |
| EP-SG-004 | Every epic must define Scope — In and Scope — Out explicitly | Reject if either scope field is empty |
| EP-SG-005 | Every epic must list Business Capabilities Delivered (ref: PM-BP-011) | Reject if Business Capabilities Delivered is empty |
| EP-SG-006 | Every epic must have a Definition of Done checklist covering development, quality, security, compliance, and deployment | Reject if Definition of Done is empty or incomplete |
| EP-SG-007 | Every epic must specify at least one Success Metric / KPI with a target value and timeframe | Flag if Success Metrics is empty |
| EP-SG-008 | Every epic must identify a Business Owner and Product Owner | Reject if either owner field is empty |
| EP-SG-009 | Every epic must list Dependencies with classification (hard/soft) and mitigation | Flag if Dependencies is empty — epics rarely have zero dependencies |

---

## 2. Duration & Sizing Guardrails (ref: PM-GR-013)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-DG-001 | Maximum epic duration must not exceed 3 months (ref: PM-GR-013) — for regulatory agility and timely value delivery | Reject if Estimated Duration exceeds 3 months |
| EP-DG-002 | Epics exceeding 3 months must be split into smaller epics, each delivering independent business value | Reject — provide decomposition guidance |
| EP-DG-003 | Each split epic must be independently deployable and deliver measurable value on its own | Flag if split epic has no standalone value |
| EP-DG-004 | Epic scope must be validated against the 3-month constraint before development begins — if scope is too large, split before starting | Process guardrail — enforce at epic approval |
| EP-DG-005 | Epics approaching the 3-month limit must have a mid-point review to assess scope and timeline risk | Process guardrail — schedule at epic kickoff |

---

## 3. Privacy & Data Protection Guardrails (ref: PM-GR-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-PG-001 | Epics involving customer data must include a Privacy Impact Assessment (PIA) before development begins (ref: PM-GR-010) | Reject if epic involves customer PII and PIA is not completed or scheduled |
| EP-PG-002 | PIA must be conducted by or reviewed by the Data Protection Officer (DPO) or Privacy Officer | Reject if PIA has no DPO/Privacy Officer involvement |
| EP-PG-003 | PIA must assess: data collected, purpose, legal basis, retention, sharing, cross-border transfer, and data subject rights | Reject if PIA is incomplete — all assessment areas must be covered |
| EP-PG-004 | PIA findings must be addressed before development proceeds — high-risk findings block epic kickoff | Reject if high-risk PIA findings are unresolved |
| EP-PG-005 | Epics involving special category data (health, biometrics, race, political opinions) require enhanced PIA with explicit legal basis | Reject if special category data is involved without enhanced PIA |
| EP-PG-006 | Epics must specify Data Classification at the epic level based on the most sensitive data handled | Reject if Data Classification is empty |
| EP-PG-007 | Epics classified as Restricted must have Security Architecture Review completed before development | Reject if Restricted epic has no security architecture review |
| EP-PG-008 | Epics involving new data collection must document the legal basis for processing (GDPR Art. 6) | Reject if new data collection has no documented legal basis |
| EP-PG-009 | Epics involving data sharing with third parties must specify the data processing agreement and legal basis | Reject if third-party data sharing has no DPA reference |

---

## 4. PCI-DSS Guardrails (ref: PM-GR-011)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-PCIG-001 | Payment processing epics must reference applicable PCI-DSS requirements (ref: PM-GR-011) | Reject if payment epic has no PCI-DSS reference |
| EP-PCIG-002 | Epics involving card data (PAN, CVV, expiry, track data) must be scoped within the PCI-DSS Cardholder Data Environment (CDE) | Reject if card data epic is not mapped to CDE |
| EP-PCIG-003 | Epics must not permit storage of sensitive authentication data (CVV, PIN, track data) post-authorization | Reject — no exceptions |
| EP-PCIG-004 | Epics involving card data must mandate tokenization for all non-production environments and logging | Reject if tokenization is not specified |
| EP-PCIG-005 | Epics involving card data must mandate encryption (AES-256) at rest and in transit (TLS 1.2+) | Reject if encryption requirements are missing |
| EP-PCIG-006 | Epics involving payment processing must include PCI-DSS compliance testing in the Definition of Done | Reject if DoD does not include PCI-DSS testing |
| EP-PCIG-007 | Epics that expand the CDE scope must trigger a PCI-DSS scope assessment with the QSA (Qualified Security Assessor) | Flag if CDE expansion has no QSA assessment scheduled |
| EP-PCIG-008 | Epics involving third-party payment processors must verify the processor's PCI-DSS compliance status (AOC/ROC) | Reject if third-party payment processor compliance is not verified |

---

## 5. Cross-Border & Legal Guardrails (ref: PM-GR-012)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-LG-001 | Cross-border data transfer epics require legal review before development begins (ref: PM-GR-012) | Reject if cross-border epic has no legal review completed or scheduled |
| EP-LG-002 | Legal review must assess: applicable jurisdictions, data transfer mechanisms (adequacy decisions, SCCs, BCRs), and local data localization requirements | Reject if legal review is incomplete |
| EP-LG-003 | Epics transferring data from EU/EEA to non-adequate countries must specify the transfer mechanism (SCCs, BCRs, or derogation) | Reject if no transfer mechanism is specified |
| EP-LG-004 | Epics involving data localization jurisdictions must confirm data residency requirements are met | Reject if data residency is not addressed for localization jurisdictions |
| EP-LG-005 | Epics involving cross-border payments must reference applicable payment regulations per jurisdiction (PSD2, local payment laws) | Flag if cross-border payment epic has no jurisdiction-specific regulatory reference |
| EP-LG-006 | Epics involving customers in sanctioned jurisdictions must include sanctions screening requirements | Reject if sanctions screening is not addressed |
| EP-LG-007 | Legal review findings must be addressed before development proceeds — blocking findings halt epic kickoff | Reject if blocking legal findings are unresolved |
| EP-LG-008 | Epics serving multiple jurisdictions must document conflicting regulatory requirements and their resolution | Flag if multi-jurisdiction epic has no conflict analysis |

---

## 6. Compliance Checkpoint Guardrails (ref: PM-BP-012)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-CG-001 | Every epic must have the Compliance Checkpoints field populated — each checkpoint marked Required or Not Required with justification | Reject if Compliance Checkpoints is empty |
| EP-CG-002 | Data Classification Review checkpoint is mandatory for all epics — no exceptions | Reject if Data Classification Review is not marked Required |
| EP-CG-003 | Security Test Results Review checkpoint is mandatory for all epics — no exceptions | Reject if Security Test Results Review is not marked Required |
| EP-CG-004 | Compliance Sign-Off checkpoint is mandatory for all epics linked to regulations | Reject if regulated epic has Compliance Sign-Off marked Not Required |
| EP-CG-005 | Checkpoint failures must block progression to the next phase — they cannot be deferred | Reject if checkpoint failure is deferred without exception approval |
| EP-CG-006 | Checkpoint outcomes must be documented: Approved / Conditionally Approved / Rejected, with rationale and approver | Reject if checkpoint outcome is not recorded |
| EP-CG-007 | Conditional approvals must have conditions tracked to closure with a deadline | Flag if conditional approval has no closure deadline |
| EP-CG-008 | Post-Deployment Compliance Verification must be completed within 2 weeks of production deployment for regulated epics | Flag if post-deployment verification is not scheduled |

---

## 7. Risk Assessment Guardrails (ref: PM-BP-014)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-RG-001 | Every epic must be assessed against all six risk categories: Operational, Credit, Market, Compliance, Reputational, Technology (ref: PM-BP-014) | Reject if Risk Category is empty or only one category is assessed |
| EP-RG-002 | Risk Assessment must use Impact × Likelihood matrix to determine risk rating (Critical/High/Medium/Low) | Reject if risk rating is not calculated |
| EP-RG-003 | Every identified risk must have a specific mitigation action — generic statements like "will be managed" are not acceptable | Reject if risk mitigation contains vague language |
| EP-RG-004 | Critical-risk epics require executive sponsor and weekly risk review | Reject if Critical-risk epic has no executive sponsor |
| EP-RG-005 | High-risk epics require Business Owner + Compliance Officer sign-off and bi-weekly risk review | Reject if High-risk epic lacks required sign-offs |
| EP-RG-006 | Risk assessment must be reviewed when epic scope changes — scope changes can alter risk profile | Process guardrail — enforce via change review |
| EP-RG-007 | Epics involving money movement must always be assessed as minimum High risk for Operational and Compliance categories | Reject if money-movement epic is rated below High for these categories |
| EP-RG-008 | Epics involving customer PII must always be assessed as minimum Medium risk for Compliance and Reputational categories | Flag if PII epic is rated below Medium for these categories |

---

## 8. Acceptance Criteria Guardrails (ref: PM-BP-013)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-AG-001 | Every epic must have both Business Acceptance Criteria and Regulatory Acceptance Criteria (ref: PM-BP-013) | Reject if either acceptance criteria section is empty |
| EP-AG-002 | Every regulation linked in Regulatory Requirements must have at least one corresponding Regulatory Acceptance Criterion | Reject if a linked regulation has no acceptance criterion |
| EP-AG-003 | Acceptance criteria must be specific, measurable, and testable — no vague language ("properly", "adequately", "as needed") | Reject if criteria contain subjective language |
| EP-AG-004 | Regulatory acceptance criteria must cite the specific regulation article they satisfy | Flag if regulatory criterion has no article reference |
| EP-AG-005 | Acceptance criteria must include evidence requirements — what proof demonstrates the criterion is met | Flag if no evidence requirement is specified |
| EP-AG-006 | Definition of Done must include: security scan passed, compliance sign-off obtained, audit trail verified, DR tested (for Tier 1) | Reject if DoD is missing any of these for applicable epics |

---

## 9. Audit Trail Guardrails (ref: PM-BP-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-ATG-001 | Every epic must specify Audit Trail Requirements — what events are logged across the epic's features | Reject if Audit Trail Requirements is empty for Critical/High priority epics |
| EP-ATG-002 | Audit trail must cover: financial transactions, data modifications, access to sensitive data, approval decisions, and security events | Flag if audit scope is incomplete |
| EP-ATG-003 | Audit trail retention must be specified and aligned with regulatory requirements (7 years financial, 3 years access) | Reject if retention is not specified or below regulatory minimum |
| EP-ATG-004 | Audit trail verification must be included in the Definition of Done | Reject if DoD does not include audit trail verification |

---

## 10. DR/BCP Guardrails (ref: PM-BP-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-DRG-001 | Every epic must specify DR/BCP Requirements including Recovery Tier (1/2/3) | Flag if DR/BCP Requirements is empty |
| EP-DRG-002 | Tier 1 epics must include DR failover testing in the Definition of Done | Reject if Tier 1 epic DoD does not include DR testing |
| EP-DRG-003 | Epics must specify degraded-mode behavior for each external dependency | Flag if no degraded-mode behavior is defined |
| EP-DRG-004 | Epics involving financial transactions must specify in-flight transaction handling during failover | Reject if financial epic has no failover transaction handling |
| EP-DRG-005 | Security controls must not be weakened during DR — MFA, encryption, and audit logging must be maintained | Reject if DR mode permits security control bypass |

---

## 11. Stakeholder & Approval Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-SAG-001 | Every epic must have Business Owner sign-off before development begins | Reject if Business Owner has not approved |
| EP-SAG-002 | Epics linked to regulations must have Compliance Officer sign-off before development begins | Reject if regulated epic lacks Compliance Officer approval |
| EP-SAG-003 | Epics involving Restricted data must have Security Architect sign-off before development begins | Reject if Restricted-data epic lacks Security Architect approval |
| EP-SAG-004 | Epics involving cross-border data must have Legal review sign-off before development begins (ref: PM-GR-012) | Reject if cross-border epic lacks Legal sign-off |
| EP-SAG-005 | Critical-risk epics must have Executive Sponsor sign-off | Reject if Critical-risk epic has no executive approval |
| EP-SAG-006 | All sign-offs must be documented with approver name, role, date, and any conditions | Reject if sign-off records are incomplete |

---

## 12. Integration & Dependency Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| EP-IDG-001 | Epics involving third-party integrations must verify the third party's security and compliance posture | Reject if third-party security assessment is not completed or scheduled |
| EP-IDG-002 | Epics involving third-party payment processors must verify PCI-DSS compliance (AOC/ROC) (ref: PM-GR-011) | Reject if payment processor compliance is unverified |
| EP-IDG-003 | Hard dependencies must have a mitigation plan — what happens if the dependency is delayed or unavailable | Reject if hard dependency has no mitigation |
| EP-IDG-004 | Epics with external dependencies must define SLA expectations and monitoring approach | Flag if no SLA is defined for external dependencies |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Epic cannot proceed to development until fixed | No PIA for customer data epic (PM-GR-010), no PCI-DSS reference for payment epic (PM-GR-011), no legal review for cross-border epic (PM-GR-012), duration > 3 months (PM-GR-013), no risk assessment, no acceptance criteria |
| **Flag** | Epic can proceed but must be addressed before first feature enters testing | Missing success metrics, incomplete dependency mitigation, no degraded-mode behavior |
| **Review** | Epic requires additional review by specified role | Compliance Officer for regulated epics, DPO for PII epics, Legal for cross-border, Security Architect for Restricted data, QSA for CDE expansion |
| **Process** | Enforced via workflow/calendar, not automated | Mid-point review for long epics, risk reassessment on scope change, post-deployment verification |

---

## Quick Reference: Guardrail Triggers by Epic Type

| Epic Involves... | Triggered Guardrails |
|---|---|
| Customer PII | EP-PG-001→009 (PIA mandatory), EP-RG-008 (min Medium risk) |
| Card/Payment data | EP-PCIG-001→008 (PCI-DSS mandatory), EP-RG-007 (min High risk) |
| Cross-border data transfer | EP-LG-001→008 (Legal review mandatory) |
| Money movement | EP-RG-007 (min High risk), EP-DRG-004 (failover handling), EP-ATG-001 (audit mandatory) |
| Regulated activity | EP-CG-004 (Compliance Sign-Off), EP-AG-001→002 (regulatory acceptance criteria), EP-SAG-002 |
| Restricted data | EP-PG-007 (Security Architecture Review), EP-SAG-003 (Security Architect sign-off) |
| Third-party integration | EP-IDG-001→004 (vendor assessment, SLA) |
| > 3 months estimated | EP-DG-001→005 (must split) |
| Critical risk rating | EP-RG-004 (executive sponsor), EP-SAG-005 (executive sign-off) |
