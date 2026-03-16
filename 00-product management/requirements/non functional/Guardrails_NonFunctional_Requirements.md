# Guardrails for Non-Functional Requirements

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-SG-001 | Every NFR must have a unique ID following NFR-[Category]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| NFR-SG-002 | Every NFR must specify a Category (Performance/Scalability/Availability/Reliability/Usability/Maintainability) | Reject if Category is empty or non-standard |
| NFR-SG-003 | Every NFR must have at least one quantifiable metric with a specific target value | Reject if metric or target is missing, vague, or uses subjective language |
| NFR-SG-004 | Every NFR must specify a measurement method (tool, technique, frequency) | Reject if measurement method is empty |
| NFR-SG-005 | Every NFR must specify environment applicability (Production/Staging/All) | Flag if environment is not specified |
| NFR-SG-006 | Every NFR must link to at least one Functional Requirement or state "System-wide" | Reject if Related Functional Req(s) is empty |

---

## 2. Metric & Target Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-MG-001 | Latency metrics must use percentiles (p95, p99), not averages | Reject if latency target uses "average" or "mean" only |
| NFR-MG-002 | Latency targets must specify the load condition (concurrent users, TPS) | Reject if latency target has no load condition |
| NFR-MG-003 | Availability targets must specify the measurement window and exclusions (e.g., planned maintenance) | Flag if availability target has no measurement window |
| NFR-MG-004 | Scalability targets must specify the scaling dimension (users, data volume, TPS) and growth expectation | Reject if scalability target is vague ("handle growth") |
| NFR-MG-005 | Targets must not use subjective language: "fast", "responsive", "scalable", "reliable", "adequate" | Reject if any subjective language is found in metric or target fields |
| NFR-MG-006 | Targets must be achievable and benchmarked — unrealistic targets (e.g., 100% uptime, 0ms latency) must be flagged | Flag unrealistic targets for architecture review |
| NFR-MG-007 | Performance targets must distinguish between normal load and peak load conditions | Flag if only one load condition is specified for critical systems |
| NFR-MG-008 | RTO and RPO values must align with the business criticality of the linked functional requirements | Flag if RTO/RPO seems misaligned with FR priority (e.g., 24-hour RTO for critical payment system) |

---

## 3. Performance-Specific Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-PG-001 | Customer-facing API response time (p95) must not exceed 500ms under stated load | Flag if target exceeds 500ms for customer-facing APIs |
| NFR-PG-002 | Internal/service-to-service API response time (p95) must not exceed 200ms under stated load | Flag if target exceeds 200ms for internal APIs |
| NFR-PG-003 | Page load time must not exceed 3 seconds for customer-facing web applications | Flag if target exceeds 3 seconds |
| NFR-PG-004 | Database query response time must not exceed 100ms for transactional queries | Flag if target exceeds 100ms for transactional queries |
| NFR-PG-005 | Batch processing must complete within the defined batch window (typically 4 hours for EOD) | Flag if batch processing has no time window constraint |
| NFR-PG-006 | Performance NFRs for critical financial operations must include degraded-mode targets (what's acceptable when a dependency is slow) | Flag if no degraded-mode target exists for critical operations |

---

## 4. Availability & Reliability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-ARG-001 | Critical banking systems (payments, core banking) must have availability target ≥ 99.95% | Reject if critical system availability target is below 99.95% |
| NFR-ARG-002 | RTO for critical systems must not exceed 1 hour | Flag if critical system RTO exceeds 1 hour |
| NFR-ARG-003 | RPO for critical systems must not exceed 15 minutes | Flag if critical system RPO exceeds 15 minutes |
| NFR-ARG-004 | Failover targets must specify zero data loss or acceptable data loss explicitly | Reject if failover NFR does not address data loss |
| NFR-ARG-005 | Availability NFRs must include planned maintenance window definition | Flag if no maintenance window is defined |
| NFR-ARG-006 | DR/failover NFRs must specify testing frequency (minimum quarterly) | Flag if no DR test frequency is specified |

---

## 5. Scalability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-ScG-001 | Scalability NFRs must specify current baseline and projected growth (1-year, 3-year) | Flag if no growth projection is provided |
| NFR-ScG-002 | Auto-scaling NFRs must specify scale-up trigger, scale-up time, and scale-down policy | Reject if auto-scaling NFR is incomplete |
| NFR-ScG-003 | Data volume scalability must account for regulatory retention requirements (typically 5-7 years) | Flag if data volume target doesn't account for retention |
| NFR-ScG-004 | Scalability NFRs must specify performance targets that must be maintained during and after scaling | Flag if no performance constraint during scaling |

---

## 6. Usability & Accessibility Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-UG-001 | Customer-facing applications must target WCAG 2.1 AA compliance at minimum | Reject if accessibility target is below WCAG 2.1 AA |
| NFR-UG-002 | Browser/device support must specify exact versions, not "all browsers" or "modern browsers" | Reject if browser support uses vague language |
| NFR-UG-003 | Error messages must be user-friendly and must not expose technical details or system internals | Flag if error message guidelines are not specified |
| NFR-UG-004 | Usability NFRs must specify the measurement method (automated audit, manual testing, user testing) | Flag if no measurement method for usability |

---

## 7. Maintainability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-MtG-001 | Code coverage target must be ≥ 80% overall and ≥ 90% for critical business logic | Flag if coverage target is below these thresholds |
| NFR-MtG-002 | Technical debt ratio must be tracked and targeted below 10% | Flag if no technical debt target is defined |
| NFR-MtG-003 | Zero critical or blocker static analysis issues must be allowed in production code | Reject if critical/blocker issues are permitted |
| NFR-MtG-004 | Deployment pipeline must support rollback within defined timeframe | Flag if no rollback capability is specified |

---

## 8. Testing & Validation Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-TG-001 | Every NFR must have a defined test scenario with pass/fail criteria before development begins | Reject if no test scenario is defined |
| NFR-TG-002 | Performance NFRs must be validated in a production-equivalent environment before release | Reject if performance testing is planned only in dev/test environments |
| NFR-TG-003 | Availability/DR NFRs must be validated through actual failover tests, not just documentation | Flag if DR validation is documentation-only |
| NFR-TG-004 | NFR test results must be included in release sign-off criteria | Reject if NFR testing is not part of release gate |
| NFR-TG-005 | Performance test data must be production-representative in volume and variety | Flag if test data is trivial or unrealistic |
| NFR-TG-006 | NFR acceptance criteria must specify duration of sustained load (not just peak burst) | Flag if performance criteria only test burst, not sustained load |

---

## 9. Dependency & Assumption Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-DG-001 | All infrastructure assumptions must be documented (CPU, memory, network, storage) | Flag if assumptions are empty for performance/scalability NFRs |
| NFR-DG-002 | External dependency SLAs must be documented and factored into targets | Flag if external dependencies are referenced without their SLA |
| NFR-DG-003 | Assumptions must be validated before NFR sign-off — unvalidated assumptions must be flagged as risks | Flag unvalidated assumptions |
| NFR-DG-004 | NFRs must not assume unlimited resources (infinite scaling, unlimited bandwidth) | Reject if NFR assumes unbounded resources |

---

## 10. Data Classification Guardrails (ref: PM-BP-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-DCG-001 | Every NFR must specify Data Classification Scope — which classification levels it applies to | Flag if Data Classification Scope is empty |
| NFR-DCG-002 | NFR targets must be tiered by data classification — Restricted data requires stricter targets than Public | Flag if same targets apply to all classification levels without justification |
| NFR-DCG-003 | Performance NFRs for Restricted/Confidential data must account for encryption overhead | Flag if encryption overhead is not factored into performance targets |
| NFR-DCG-004 | Backup and replication NFRs must be tiered by data classification (continuous for Restricted, daily for Internal) | Reject if backup frequency is not aligned with data classification |
| NFR-DCG-005 | Data masking/anonymization performance in non-prod must be specified for Confidential/Restricted data | Flag if no masking performance target exists |

---

## 11. Audit Trail & Logging Guardrails (ref: PM-BP-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-ALG-001 | Logging infrastructure must have its own NFRs — log ingestion latency, query performance, availability, and capacity | Reject if no logging infrastructure NFRs exist |
| NFR-ALG-002 | Log ingestion latency must not exceed 5 seconds from event to central store | Flag if log latency target exceeds 5 seconds |
| NFR-ALG-003 | Log availability must be ≥ 99.99% — logs must be queryable when needed for incident response | Flag if log availability target is below 99.99% |
| NFR-ALG-004 | Log completeness must target 100% of defined audit events — no silent log drops | Reject if log completeness target is below 100% |
| NFR-ALG-005 | Log retention NFRs must align with regulatory requirements (7 years financial, 3 years access) | Reject if retention is below regulatory minimum |
| NFR-ALG-006 | Log storage capacity must be planned for the full retention period at projected growth rate | Flag if no capacity projection exists |
| NFR-ALG-007 | Log integrity verification must be automated and run at minimum daily | Flag if no integrity verification frequency is specified |

---

## 12. DR/BCP Guardrails (ref: PM-BP-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| NFR-DRG-001 | RTO and RPO must be defined as explicit NFRs with specific targets, not just assumptions | Reject if RTO/RPO are mentioned only in assumptions |
| NFR-DRG-002 | Tier 1 critical systems must have RTO < 1 hour and RPO < 15 minutes | Reject if Tier 1 targets exceed these thresholds |
| NFR-DRG-003 | Failover NFRs must specify data consistency guarantee (zero loss for committed transactions) | Reject if failover NFR does not address data consistency |
| NFR-DRG-004 | Degraded-mode performance targets must be defined — acceptable performance during partial failure | Flag if no degraded-mode targets exist for Tier 1/2 systems |
| NFR-DRG-005 | DR test frequency must be specified as an NFR: quarterly for Tier 1, semi-annual for Tier 2/3 | Reject if no DR test frequency is defined |
| NFR-DRG-006 | DR test success criteria must be defined (RTO met, RPO met, data integrity verified, no duplicate transactions) | Flag if DR test has no pass/fail criteria |
| NFR-DRG-007 | Backup verification NFRs must include automated restoration testing (not just backup creation) | Flag if backup verification is creation-only |
| NFR-DRG-008 | Geographic redundancy requirements must be specified for Tier 1 systems | Flag if Tier 1 system has no geographic redundancy NFR |
| NFR-DRG-009 | BCP activation time must be defined — how quickly the organization switches to continuity mode | Flag if no BCP activation time is specified |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | NFR cannot proceed until fixed | Missing metric, subjective language, no measurement method, average-only latency |
| **Flag** | NFR can proceed but must be addressed before release | Missing growth projections, no degraded-mode target, unvalidated assumptions |
| **Review** | NFR requires additional review by specified role | Architect for scalability, DevOps for availability, Security for encryption performance |
