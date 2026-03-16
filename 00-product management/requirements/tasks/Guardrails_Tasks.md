# Guardrails for Tasks

---

## 1. Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-SG-001 | Every task must have a unique ID following TK-[Type]-[Module]-[Seq] convention | Reject if ID is missing, duplicated, or non-conformant |
| TK-SG-002 | Every task must specify a Task Type: DEV / SEC / TST / CMP / OPS / RBK / DOC / REV (ref: PM-BP-020) | Reject if Task Type is empty or non-standard |
| TK-SG-003 | Every task must be linked to a parent User Story and Epic | Reject if User Story or Epic field is empty |
| TK-SG-004 | Every task must have an Assignee — unassigned tasks must not enter a sprint | Reject if task enters sprint without assignee |
| TK-SG-005 | Every task must have an Estimated Effort — unestimated tasks must not enter a sprint | Reject if task enters sprint without estimate |
| TK-SG-006 | Every task must have Acceptance Criteria with pass/fail conditions | Reject if Acceptance Criteria is empty |
| TK-SG-007 | Every task must have a Definition of Done checklist | Flag if Definition of Done is empty |
| TK-SG-008 | Tasks must be completable by one person in 1-2 days maximum — larger tasks must be split | Flag if Estimated Effort exceeds 16 hours |
| TK-SG-009 | Every task must inherit the Data Sensitivity tag from its parent User Story | Reject if Data Sensitivity is empty |

---

## 2. Task Separation Guardrails (ref: PM-BP-020)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-TSG-001 | Every user story must be decomposed into separate tasks by type — development, security, testing, and compliance must not be combined into a single task (ref: PM-BP-020) | Reject if a single task combines development with security scanning, testing, or compliance validation |
| TK-TSG-002 | Every story must have at minimum: one DEV task, one TST task, and one SEC task | Reject if story has no testing or security task |
| TK-TSG-003 | Stories tagged Confidential or Restricted must additionally have: DAST task, security code review task, and compliance validation task | Reject if sensitive story is missing these task types |
| TK-TSG-004 | Every story deploying to production must have at least one RBK (rollback) task (ref: PM-BP-023) | Reject if production-bound story has no rollback task |
| TK-TSG-005 | Security tasks must be assigned to someone other than the developer who wrote the code | Reject if security review/scan task is assigned to the code author |
| TK-TSG-006 | Compliance tasks must be reviewed or approved by the Compliance Officer — developer self-certification is not acceptable | Reject if compliance task has no Compliance Officer involvement |

---

## 3. Security Review Guardrails (ref: PM-GR-019)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-SRG-001 | Security review tasks are mandatory for all features — no feature can be deployed without security review (ref: PM-GR-019) | Reject if any feature has no security review task |
| TK-SRG-002 | SAST scan task is mandatory for every story — must run on completed code with zero critical/high findings before merge | Reject if story has no SAST task; reject merge if critical/high findings exist |
| TK-SRG-003 | SCA dependency scan task is mandatory for every story — must identify and remediate critical CVEs before merge | Reject if story has no SCA task; reject merge if critical CVEs exist |
| TK-SRG-004 | Security code review task is mandatory for all stories tagged Confidential or Restricted — minimum 1 reviewer (2 for Restricted) | Reject if sensitive story has no security code review task |
| TK-SRG-005 | DAST scan task is mandatory for all stories tagged Confidential or Restricted — must run on staging with zero critical/high findings | Reject if sensitive story has no DAST task |
| TK-SRG-006 | Auth implementation review task is mandatory for any story involving authentication or authorization changes | Reject if auth story has no auth review task |
| TK-SRG-007 | Input validation review task is mandatory for any story accepting user input | Flag if user-input story has no input validation review task |
| TK-SRG-008 | Encryption verification task is mandatory for stories handling Restricted data | Reject if Restricted story has no encryption verification task |
| TK-SRG-009 | Security scan findings classified as Critical must be remediated before merge — no deferral permitted | Reject merge with open Critical security findings |
| TK-SRG-010 | Security scan findings classified as High must be remediated before production deployment — deferral requires Security Engineer approval with documented justification | Reject deployment with open High findings unless formally approved |
| TK-SRG-011 | Security review tasks must produce documented evidence (scan reports, review checklists, finding logs) retained for audit | Reject if security task produces no evidence artifact |
| TK-SRG-012 | Security review must cover OWASP Top 10 risks relevant to the story's functionality | Flag if security review scope does not reference applicable OWASP risks |

---

## 4. Compliance Sign-Off Guardrails (ref: PM-GR-020)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-CSG-001 | Compliance sign-off is required before production deployment for any story linked to regulations (ref: PM-GR-020) | Reject production deployment if compliance sign-off is not obtained |
| TK-CSG-002 | Compliance review task must be submitted before development begins — not after | Reject if compliance review is scheduled after development for regulated stories |
| TK-CSG-003 | Compliance sign-off must be documented: approver name, role, date, scope, and any conditions | Reject if sign-off record is incomplete |
| TK-CSG-004 | Conditional compliance approvals must have all conditions resolved before production deployment | Reject deployment if compliance conditions are open |
| TK-CSG-005 | Compliance evidence collection task must gather all artifacts required to demonstrate regulatory compliance | Reject if evidence collection task produces incomplete artifacts |
| TK-CSG-006 | Stories involving payment/card data must have PCI-DSS compliance verification before deployment | Reject deployment of payment story without PCI-DSS verification |
| TK-CSG-007 | Stories involving customer data processing must have GDPR/data protection compliance verification before deployment | Reject deployment of data processing story without privacy compliance verification |
| TK-CSG-008 | Stories involving AML/KYC functionality must have AML compliance verification before deployment | Reject deployment of AML story without AML compliance verification |
| TK-CSG-009 | Compliance sign-off cannot be retroactive — it must be obtained before the deployment occurs, not after | Reject retroactive compliance sign-off |
| TK-CSG-010 | Changes to existing compliance controls require impact assessment and re-certification before deployment | Reject deployment of compliance control changes without impact assessment |

---

## 5. Technical Debt & Risk Guardrails (ref: PM-GR-021)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-TDG-001 | Tasks must not create technical debt that increases operational, security, or compliance risk (ref: PM-GR-021) | Reject if task introduces known risk-increasing technical debt |
| TK-TDG-002 | Tasks must not introduce TODO/FIXME/HACK comments for security or compliance-related code — these must be resolved before merge | Reject merge if security/compliance TODOs exist |
| TK-TDG-003 | Tasks must not disable or skip existing tests to make new code pass | Reject if existing tests are disabled or removed without explicit approval |
| TK-TDG-004 | Tasks must not reduce code coverage below the minimum threshold (80% overall, 90% critical paths) | Reject if task reduces coverage below threshold |
| TK-TDG-005 | Tasks must not introduce deprecated libraries, APIs, or patterns when current alternatives exist | Flag if deprecated technology is introduced |
| TK-TDG-006 | Tasks must not bypass code quality gates (linting, formatting, complexity thresholds) | Reject if quality gates are bypassed |
| TK-TDG-007 | If a task must introduce technical debt due to time constraints, it must be documented as a follow-up task with priority and deadline | Flag if technical debt is introduced without a documented remediation task |
| TK-TDG-008 | Technical debt tasks must include a risk assessment: what risk does this debt create and what is the impact if not resolved | Flag if debt task has no risk assessment |
| TK-TDG-009 | Tasks must not introduce hardcoded values for configurable parameters (URLs, timeouts, limits, thresholds) | Flag if hardcoded configuration values are introduced |
| TK-TDG-010 | Tasks must not introduce code duplication for security-critical logic (auth, encryption, validation) — use shared libraries | Flag if security-critical code is duplicated |

---

## 6. Peer Review Guardrails (ref: PM-GR-022)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-PRG-001 | Peer review is required for all critical and high-risk tasks before merge (ref: PM-GR-022) | Reject merge without peer review for Critical/High priority tasks |
| TK-PRG-002 | All code changes must have at least one peer review — no self-approved merges | Reject self-approved merges |
| TK-PRG-003 | Tasks involving Restricted data must have two peer reviewers, at least one with security expertise | Reject if Restricted data task has fewer than 2 reviewers |
| TK-PRG-004 | Tasks involving authentication, authorization, or encryption logic must be reviewed by a reviewer with security expertise | Reject if auth/crypto task reviewer lacks security expertise |
| TK-PRG-005 | Tasks involving financial calculations (interest, fees, FX rates, limits) must be reviewed by a reviewer with domain expertise | Flag if financial calculation task reviewer lacks domain knowledge |
| TK-PRG-006 | Tasks involving database schema changes must be reviewed by a DBA or database-experienced reviewer | Flag if schema change task has no DBA review |
| TK-PRG-007 | Peer review must verify: functional correctness, security controls, input validation, error handling, audit logging, and test coverage | Reject if review checklist is not followed |
| TK-PRG-008 | Peer review findings classified as Critical or High must be resolved before merge — no deferral | Reject merge with open Critical/High review findings |
| TK-PRG-009 | Peer review must be documented: reviewer name, date, findings, resolution, and approval | Reject if review record is incomplete |
| TK-PRG-010 | Reviewer must not be the same person as the task assignee — independent review is mandatory | Reject if reviewer = assignee |

---

## 7. Dependency & Sequencing Guardrails (ref: PM-BP-022)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-DSG-001 | Every task must have Dependencies and Sequence Order fields populated | Flag if dependencies are not mapped |
| TK-DSG-002 | Tasks must not start if their predecessor tasks (Finish-to-Start dependencies) are not complete | Reject if task starts with incomplete predecessors |
| TK-DSG-003 | External dependencies must be identified and have a mitigation plan before the sprint starts | Reject if external dependency has no mitigation plan |
| TK-DSG-004 | Compliance review tasks must be sequenced before development tasks for regulated stories | Reject if compliance review is sequenced after development |
| TK-DSG-005 | DAST scan tasks must be sequenced after staging deployment — they require a running environment | Reject if DAST is scheduled before staging deployment |
| TK-DSG-006 | Compliance sign-off tasks must be sequenced after all security and testing tasks complete | Reject if compliance sign-off is scheduled before security/testing completion |
| TK-DSG-007 | Production deployment tasks must be sequenced after compliance sign-off (ref: PM-GR-020) | Reject if deployment is scheduled before compliance sign-off |
| TK-DSG-008 | Rollback testing tasks must be sequenced before production deployment | Reject if rollback testing is scheduled after production deployment |

---

## 8. Rollback & Remediation Guardrails (ref: PM-BP-023)

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-RRG-001 | Every story deploying to production must have a rollback plan task created before deployment (ref: PM-BP-023) | Reject deployment if no rollback plan exists |
| TK-RRG-002 | Rollback plan must specify: trigger conditions, decision maker, steps, time estimate, data impact, and verification | Reject if rollback plan is incomplete |
| TK-RRG-003 | Rollback procedure must be tested in staging for high-risk deployments (Confidential/Restricted data, database changes, financial logic) | Reject high-risk deployment if rollback is untested |
| TK-RRG-004 | Database changes must have a reverse migration script created and tested before production deployment | Reject DB deployment without tested reverse migration |
| TK-RRG-005 | Data migration tasks must include a data remediation plan for rollback scenarios | Reject data migration without remediation plan |
| TK-RRG-006 | Rollback must not cause data loss for committed transactions — data consistency must be maintained | Reject if rollback plan permits committed transaction loss |
| TK-RRG-007 | Rollback must not weaken security controls — rolled-back state must maintain all security and compliance controls | Reject if rolled-back state has weaker security posture |
| TK-RRG-008 | Post-rollback tasks must include: incident report, root cause analysis, and fix-forward plan | Flag if no post-rollback tasks are defined |

---

## 9. Deployment Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-DPG-001 | Production deployment must only occur through the CI/CD pipeline — no manual deployments | Reject manual production deployments |
| TK-DPG-002 | Production deployment must have all security scan tasks completed with zero critical/high findings | Reject deployment with open critical/high security findings |
| TK-DPG-003 | Production deployment must have compliance sign-off obtained (ref: PM-GR-020) | Reject deployment without compliance sign-off |
| TK-DPG-004 | Production deployment must have rollback plan created and tested (ref: PM-BP-023) | Reject deployment without rollback readiness |
| TK-DPG-005 | Production deployment must have monitoring and alerting configured before go-live | Reject deployment without monitoring setup |
| TK-DPG-006 | Production deployment must use approved deployment strategy (blue-green, canary, or feature flag) for high-risk changes | Flag if high-risk deployment uses direct cutover |
| TK-DPG-007 | Deployment artifacts must be signed and verified — no unsigned artifacts in production | Reject unsigned deployment artifacts |
| TK-DPG-008 | Deployment must be audited: who deployed, what version, when, approval chain | Reject if deployment audit trail is incomplete |

---

## 10. Documentation & Evidence Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TK-DEG-001 | Security tasks must produce evidence artifacts (scan reports, review checklists) retained for audit | Reject if security task has no evidence |
| TK-DEG-002 | Compliance tasks must produce evidence artifacts (review records, sign-off documents, test results) retained for audit | Reject if compliance task has no evidence |
| TK-DEG-003 | Peer review tasks must produce documented records (reviewer, date, findings, resolution) | Reject if review has no documentation |
| TK-DEG-004 | Rollback test tasks must produce test results documenting success/failure and observations | Flag if rollback test has no results |
| TK-DEG-005 | All evidence must be retained per regulatory requirements (7 years financial, 3 years access) | Reject if evidence retention is below regulatory minimum |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Task/deployment cannot proceed until fixed | No security review for feature (PM-GR-019), no compliance sign-off before deployment (PM-GR-020), risk-increasing technical debt (PM-GR-021), no peer review for critical task (PM-GR-022), self-approved merge, deployment with open critical findings |
| **Flag** | Task can proceed but must be addressed before story is Done | Missing dependency mapping, hardcoded config values, no post-rollback tasks, deprecated technology |
| **CI/CD Gate** | Automated enforcement in pipeline | SAST/SCA block merge on critical findings, unsigned artifacts blocked, manual deployment blocked |
| **Process** | Enforced via workflow | Compliance review before dev, rollback test before deployment, evidence retention |

---

## Quick Reference: Mandatory Tasks by Story Sensitivity

| Task | Public | Internal | Confidential | Restricted |
|---|---|---|---|---|
| DEV tasks | ✅ | ✅ | ✅ | ✅ |
| Unit tests (TST) | ✅ | ✅ | ✅ | ✅ |
| Integration tests (TST) | Recommended | ✅ | ✅ | ✅ |
| SAST scan (SEC) | ✅ | ✅ | ✅ | ✅ |
| SCA scan (SEC) | ✅ | ✅ | ✅ | ✅ |
| Peer code review (REV) | ✅ (1 reviewer) | ✅ (1 reviewer) | ✅ (1 reviewer) | ✅ (2 reviewers) |
| Security code review (SEC) | — | Recommended | ✅ | ✅ (security expert) |
| DAST scan (SEC) | — | — | ✅ | ✅ |
| Auth review (SEC) | — | If auth involved | ✅ | ✅ |
| Encryption verification (SEC) | — | — | Recommended | ✅ |
| Security tests (TST) | — | — | ✅ | ✅ |
| Compliance tests (TST) | — | — | If regulated | ✅ |
| Performance tests (TST) | — | — | If customer-facing | ✅ |
| Compliance review (CMP) | — | — | If regulated | ✅ |
| Compliance sign-off (CMP) | — | — | If regulated | ✅ |
| Rollback plan (RBK) | Recommended | ✅ | ✅ | ✅ |
| Rollback testing (RBK) | — | — | ✅ | ✅ |
| Monitoring setup (OPS) | Recommended | ✅ | ✅ | ✅ |
| Documentation (DOC) | Recommended | ✅ | ✅ | ✅ |

---

## Pre-Deployment Checklist

Before any production deployment, verify all of the following:

| # | Check | Guardrail Ref |
|---|---|---|
| 1 | All security review tasks completed with zero critical/high findings | PM-GR-019, TK-SRG-001 |
| 2 | Compliance sign-off obtained and documented | PM-GR-020, TK-CSG-001 |
| 3 | No risk-increasing technical debt introduced | PM-GR-021, TK-TDG-001 |
| 4 | All critical/high-risk tasks peer-reviewed | PM-GR-022, TK-PRG-001 |
| 5 | Rollback plan created and tested in staging | PM-BP-023, TK-RRG-001 |
| 6 | All tests passing (unit, integration, security, compliance) | TK-SRG-002 |
| 7 | Monitoring and alerting configured | TK-DPG-005 |
| 8 | Deployment via CI/CD pipeline (no manual deployment) | TK-DPG-001 |
| 9 | Deployment artifacts signed and verified | TK-DPG-007 |
| 10 | Evidence artifacts collected and stored | TK-DEG-001-005 |
