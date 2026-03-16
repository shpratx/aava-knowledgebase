# Task Template

---

## Template Fields

| Field | Description |
|---|---|
| **Task ID** | TK-[Type]-[Module]-[Seq] (e.g., TK-DEV-PAY-042, TK-SEC-PAY-043, TK-TST-PAY-044) |
| **Title** | Action-oriented title describing the work |
| **Task Type** | DEV (Development) / SEC (Security) / TST (Testing) / CMP (Compliance) / OPS (DevOps/Deployment) / RBK (Rollback/Remediation) / DOC (Documentation) / REV (Review) (ref: PM-BP-020) |
| **User Story** | Parent user story ID |
| **Epic** | Parent epic ID |
| **Priority** | Critical / High / Medium / Low |
| **Assignee** | Person or role responsible |
| **Estimated Effort** | Hours or story points |
| **Data Sensitivity** | Inherited from parent story: Public / Internal / Confidential / Restricted |
| **Description** | What needs to be done, context, and technical details |
| **Acceptance Criteria** | Specific conditions for task completion |
| **Dependencies** | Predecessor tasks that must complete first (ref: PM-BP-022) |
| **Blocked By** | Tasks or external items currently blocking this task |
| **Blocks** | Downstream tasks that cannot start until this completes |
| **Sequence Order** | Execution order within the story's task set (ref: PM-BP-022) |
| **Rollback Plan** | How to undo this task's changes if needed (ref: PM-BP-023) |
| **Remediation Steps** | What to do if this task fails or produces defects (ref: PM-BP-023) |
| **Definition of Done** | Checklist for task completion |

---

## Task Type Separation (ref: PM-BP-020)

Every user story should be decomposed into separate tasks by type:

### Development Tasks (DEV)
| Task | Description | When Required |
|---|---|---|
| Backend implementation | API/service logic, business rules, data access | Always |
| Frontend implementation | UI components, forms, validation, UX | If UI involved |
| Database changes | Schema changes, migrations, indexes | If data model changes |
| Integration implementation | External system connectors, API clients, message handlers | If integrations involved |
| Audit logging implementation | Implement audit events per story specification | If Confidential/Restricted data |
| Configuration | Environment configs, feature flags, secrets setup | Always |

### Security Tasks (SEC) (ref: PM-BP-021)
| Task | Description | When Required |
|---|---|---|
| Security code review | Dedicated security-focused peer review | Always for Confidential/Restricted |
| SAST scan & remediation | Static analysis scan, fix findings | Always |
| DAST scan & remediation | Dynamic analysis scan, fix findings | Confidential/Restricted stories |
| Dependency scan (SCA) | Scan third-party libraries for CVEs | Always |
| Auth implementation review | Verify auth/authz implementation matches spec | If auth involved |
| Input validation review | Verify server-side validation, injection prevention | If user input involved |
| Encryption verification | Verify data encrypted at rest and in transit per spec | If Confidential/Restricted data |
| Penetration test support | Support pen testers with environment and context | Annual or major changes |

### Testing Tasks (TST)
| Task | Description | When Required |
|---|---|---|
| Unit test development | Write unit tests for business logic | Always (≥ 80% coverage) |
| Integration test development | Write integration tests for API/service interactions | Always |
| Security test development | Write tests for auth bypass, injection, IDOR | Confidential/Restricted stories |
| Compliance test development | Write tests for regulatory acceptance criteria | If regulated |
| Performance test development | Write load/stress tests against NFR targets | Customer-facing stories |
| Accessibility test | WCAG 2.1 AA verification | If UI involved |
| UAT support | Support business users during acceptance testing | Always |

### Compliance Tasks (CMP) (ref: PM-BP-021)
| Task | Description | When Required |
|---|---|---|
| Compliance review request | Submit story for compliance/legal review | If regulated (PM-GR-017) |
| Regulatory evidence collection | Gather audit artifacts demonstrating compliance | If regulated |
| Privacy impact assessment | Conduct or update PIA | If new customer data processing |
| PCI-DSS scope assessment | Verify CDE scope impact | If card/payment data |
| Compliance sign-off | Obtain formal compliance approval | If regulated |

### DevOps/Deployment Tasks (OPS)
| Task | Description | When Required |
|---|---|---|
| CI/CD pipeline update | Update build/deploy pipeline for new components | If new services/components |
| Environment configuration | Configure staging/production environments | Always |
| Monitoring & alerting setup | Configure dashboards, alerts, log queries | Always for production |
| Database migration execution | Run schema migrations in staging/production | If database changes |
| Feature flag configuration | Set up feature flags for gradual rollout | Recommended for high-risk changes |

### Rollback & Remediation Tasks (RBK) (ref: PM-BP-023)
| Task | Description | When Required |
|---|---|---|
| Rollback plan creation | Document step-by-step rollback procedure | Always |
| Rollback testing | Test rollback procedure in staging | High-risk changes |
| Database rollback script | Create reverse migration script | If database changes |
| Data remediation plan | Plan for fixing data if deployment causes corruption | If data modifications |
| Incident response preparation | Prepare runbook for deployment-related incidents | Production deployments |

### Documentation Tasks (DOC)
| Task | Description | When Required |
|---|---|---|
| API documentation | Update OpenAPI spec, endpoint docs | If API changes |
| Runbook update | Update operational runbooks | If operational procedures change |
| Architecture decision record | Document significant design decisions | If architectural decisions made |
| Audit trail documentation | Document audit events and log formats | If audit logging implemented |

---

## Dependency & Sequencing Detail (ref: PM-BP-022)

### Standard Task Sequence

```
Phase 1: Planning & Review
  TK-CMP: Compliance review request
  TK-DOC: Architecture decision record (if needed)
  TK-RBK: Rollback plan creation
    ↓
Phase 2: Development
  TK-DEV: Backend implementation
  TK-DEV: Frontend implementation (parallel with backend if API contract defined)
  TK-DEV: Database changes (before backend if schema dependency)
  TK-DEV: Integration implementation
  TK-DEV: Audit logging implementation
  TK-DEV: Configuration
    ↓
Phase 3: Security & Quality
  TK-SEC: SAST scan & remediation
  TK-SEC: Dependency scan (SCA)
  TK-SEC: Security code review
  TK-TST: Unit test development (parallel with dev or immediately after)
  TK-TST: Integration test development
  TK-TST: Security test development
  TK-TST: Compliance test development
    ↓
Phase 4: Validation
  TK-SEC: DAST scan & remediation (requires deployed to staging)
  TK-SEC: Auth implementation review
  TK-SEC: Input validation review
  TK-SEC: Encryption verification
  TK-TST: Performance test development & execution
  TK-TST: Accessibility test
  TK-CMP: Regulatory evidence collection
  TK-RBK: Rollback testing (in staging)
    ↓
Phase 5: Approval & Deployment
  TK-CMP: Compliance sign-off
  TK-TST: UAT support
  TK-OPS: Environment configuration
  TK-OPS: Monitoring & alerting setup
  TK-OPS: Database migration execution
  TK-OPS: Feature flag configuration
  TK-DOC: API documentation, runbook update
    ↓
Phase 6: Post-Deployment
  TK-RBK: Incident response preparation (active monitoring)
```

### Dependency Types

| Type | Symbol | Meaning | Example |
|---|---|---|---|
| Finish-to-Start (FS) | A → B | B cannot start until A finishes | Backend dev → Integration test |
| Start-to-Start (SS) | A ⇒ B | B can start when A starts | Backend dev ⇒ Unit test dev |
| Finish-to-Finish (FF) | A ⇔ B | B cannot finish until A finishes | All dev tasks ⇔ SAST scan |
| External | A ⟶ B | Dependency on external team/system | Core Banking API ready → Integration dev |

---

## Rollback & Remediation Detail (ref: PM-BP-023)

### Rollback Plan Template

| Field | Description |
|---|---|
| **Rollback Trigger** | What conditions trigger a rollback (error rate > X%, P1 incident, data corruption) |
| **Rollback Decision Maker** | Who authorizes the rollback (on-call lead, release manager) |
| **Rollback Steps** | Numbered step-by-step procedure |
| **Rollback Time Estimate** | Expected time to complete rollback |
| **Data Impact** | What happens to data created/modified during the deployment window |
| **Rollback Verification** | How to verify the rollback was successful |
| **Communication** | Who to notify (stakeholders, customers, support) |

### Rollback by Change Type

| Change Type | Rollback Approach | Considerations |
|---|---|---|
| Code deployment | Redeploy previous version via CI/CD | Feature flags preferred for instant rollback |
| Database schema (additive) | Keep new columns/tables, rollback code only | Backward-compatible migrations preferred |
| Database schema (destructive) | Execute reverse migration script | Must be tested in staging first |
| Data migration | Execute data remediation script | Must preserve audit trail of remediation |
| Configuration change | Revert config via version control | Config changes should be versioned |
| Third-party integration | Disable integration via feature flag/circuit breaker | Must not break dependent flows |

---

## Example — Task Set for Fund Transfer Story (US-PAY-042)

| Task ID | Type | Title | Seq | Dependencies | Effort |
|---|---|---|---|---|---|
| TK-CMP-PAY-101 | CMP | Compliance review — PSD2 SCA & AML requirements | 1 | None | 4h |
| TK-RBK-PAY-102 | RBK | Create rollback plan for transfer feature | 1 | None | 2h |
| TK-DEV-PAY-103 | DEV | Implement transfer API endpoint (POST /v1/transfers) | 2 | TK-CMP-PAY-101 | 16h |
| TK-DEV-PAY-104 | DEV | Implement transfer UI (form, confirmation, error states) | 2 | API contract defined | 12h |
| TK-DEV-PAY-105 | DEV | Implement fraud engine integration | 2 | TK-DEV-PAY-103 | 8h |
| TK-DEV-PAY-106 | DEV | Implement audit logging for transfer events | 2 | TK-DEV-PAY-103 | 4h |
| TK-DEV-PAY-107 | DEV | Implement MFA step-up flow for transfers | 2 | TK-DEV-PAY-103 | 8h |
| TK-SEC-PAY-108 | SEC | SAST scan & remediation | 3 | All DEV tasks | 4h |
| TK-SEC-PAY-109 | SEC | SCA dependency scan | 3 | All DEV tasks | 2h |
| TK-SEC-PAY-110 | SEC | Security code review (2 reviewers — Restricted data) | 3 | All DEV tasks | 4h |
| TK-TST-PAY-111 | TST | Unit tests — transfer logic, validation, limits | 2 | TK-DEV-PAY-103 (SS) | 8h |
| TK-TST-PAY-112 | TST | Integration tests — core banking, fraud engine, notifications | 3 | TK-DEV-PAY-105 | 8h |
| TK-TST-PAY-113 | TST | Security tests — MFA bypass, IDOR, injection, auth enforcement | 3 | TK-DEV-PAY-107 | 6h |
| TK-TST-PAY-114 | TST | Compliance tests — CTR generation, SCA verification | 3 | TK-DEV-PAY-103 | 4h |
| TK-SEC-PAY-115 | SEC | DAST scan & remediation (staging) | 4 | Deployed to staging | 4h |
| TK-SEC-PAY-116 | SEC | Auth implementation review | 4 | TK-DEV-PAY-107 | 2h |
| TK-SEC-PAY-117 | SEC | Encryption verification (TLS, data at rest) | 4 | Deployed to staging | 2h |
| TK-TST-PAY-118 | TST | Performance test — p95 < 500ms at 500 concurrent | 4 | Deployed to staging | 4h |
| TK-TST-PAY-119 | TST | Accessibility test — WCAG 2.1 AA | 4 | TK-DEV-PAY-104 | 2h |
| TK-CMP-PAY-120 | CMP | Collect regulatory evidence & obtain compliance sign-off | 5 | All SEC + TST tasks | 4h |
| TK-RBK-PAY-121 | RBK | Test rollback procedure in staging | 4 | Deployed to staging | 2h |
| TK-OPS-PAY-122 | OPS | Configure monitoring, alerting, and dashboards | 5 | All validation tasks | 4h |
| TK-OPS-PAY-123 | OPS | Production deployment with feature flag | 5 | TK-CMP-PAY-120 | 2h |
| TK-DOC-PAY-124 | DOC | Update API docs, runbook, audit trail documentation | 5 | All DEV tasks | 4h |

**Total: 24 tasks | ~114 hours | 5 phases**

---

## Usage Guidelines

1. **Separate tasks by type** (PM-BP-020) — never combine dev, security, testing, and compliance work into a single task
2. **Always include security scanning and code review tasks** (PM-BP-021) — SAST/SCA always, DAST for sensitive stories
3. **Define dependencies and sequence explicitly** (PM-BP-022) — use the standard phase sequence, document blockers
4. **Always include rollback and remediation tasks** (PM-BP-023) — rollback plan before deployment, tested in staging
5. **Task granularity**: each task should be completable by one person in 1-2 days max
6. **Inherit data sensitivity** from the parent story — it drives which security/compliance tasks are required
