# Best Practices for Composing Tasks

---

## 1. Separating Tasks by Type (ref: PM-BP-020)

**Best Practices:**
- Decompose every user story into distinct task types: Development, Security, Testing, Compliance, DevOps, Rollback, Documentation
- Never combine security work into development tasks — security tasks need dedicated attention and often different assignees
- Never combine compliance work into testing tasks — compliance validation requires specific expertise and sign-off authority
- Each task type has different skills, tools, and Definition of Done criteria
- Task separation ensures nothing is skipped — a single "implement and test" task often means testing gets cut when time is short

**Why Separation Matters in Banking:**

| Risk of Combining | Consequence |
|---|---|
| Dev + Security in one task | Security scanning skipped under time pressure; vulnerabilities shipped |
| Dev + Testing in one task | Testing reduced to "it works on my machine"; edge cases missed |
| Dev + Compliance in one task | Compliance evidence not collected; audit findings |
| Deployment without rollback task | No rollback plan when production fails; extended outage |

**Minimum Task Set per Story:**

| Story Sensitivity | Minimum Tasks |
|---|---|
| Public | DEV + TST (unit) + SEC (SAST, SCA) + OPS + DOC |
| Internal | DEV + TST (unit, integration) + SEC (SAST, SCA, code review) + OPS + RBK + DOC |
| Confidential | DEV + TST (unit, integration, security) + SEC (SAST, SCA, DAST, code review, auth review) + CMP + OPS + RBK + DOC |
| Restricted | DEV + TST (unit, integration, security, compliance, performance) + SEC (SAST, SCA, DAST, code review, auth review, encryption verification) + CMP (review + sign-off) + OPS + RBK (plan + test) + DOC |

**Task Naming Convention:**
> TK-[Type]-[Module]-[Seq]: [Action verb] [what] [context]

**Good Examples:**
- TK-DEV-PAY-103: Implement transfer API endpoint
- TK-SEC-PAY-108: Run SAST scan and remediate findings
- TK-TST-PAY-113: Develop security tests for MFA bypass and IDOR
- TK-CMP-PAY-120: Collect PSD2 SCA compliance evidence
- TK-RBK-PAY-121: Test rollback procedure in staging

**Bad Examples:**
- TK-001: Do the transfer feature (too vague, combines everything)
- TK-DEV-PAY-103: Implement and test and deploy transfer (combines types)

---

## 2. Including Security Scanning, Code Review & Compliance Validation Tasks (ref: PM-BP-021)

**Best Practices:**
- Security tasks are not optional — they are mandatory deliverables, not nice-to-haves
- Schedule security tasks in the delivery plan with time estimates — they take real effort
- SAST and SCA scans should run on every build (automated), but remediation is a manual task that needs time
- DAST requires a deployed environment — schedule after staging deployment
- Code review for security must be done by someone other than the developer — ideally with security expertise
- Compliance validation tasks must produce evidence artifacts — not just a checkbox

**Security Task Checklist:**

| Task | Tool/Method | When | Effort Estimate | Mandatory |
|---|---|---|---|---|
| SAST scan | SonarQube, Checkmarx, Semgrep | After code complete | 2-4h (scan + remediation) | Always |
| SCA dependency scan | Snyk, OWASP Dependency-Check | After code complete | 1-2h (scan + remediation) | Always |
| Security code review | Manual review by peer with security focus | After code complete | 2-4h (2 reviewers for Restricted) | Always for Confidential/Restricted |
| DAST scan | OWASP ZAP, Burp Suite | After staging deployment | 2-4h (scan + remediation) | Confidential/Restricted |
| Auth implementation review | Manual verification against spec | After auth code complete | 2-4h | If auth involved |
| Input validation review | Manual + automated fuzzing | After code complete | 2-4h | If user input involved |
| Encryption verification | Manual verification + automated check | After staging deployment | 1-2h | Confidential/Restricted |

**Compliance Validation Task Checklist:**

| Task | Who | When | Effort Estimate | Mandatory |
|---|---|---|---|---|
| Compliance review request | Developer → Compliance Officer | Before development | 2-4h (preparation + review) | If regulated |
| Regulatory evidence collection | Developer + QA | After testing complete | 2-4h | If regulated |
| PIA conduct/update | DPO/Privacy Officer | Before development | 4-8h | If new PII processing |
| PCI-DSS scope assessment | QSA/Security team | Before development | 4-8h | If card/payment data |
| Compliance sign-off | Compliance Officer | Before production deployment | 2-4h | If regulated |

**Common Mistakes:**
- "We'll do security scanning later" — later never comes; schedule it now
- Estimating zero time for remediation — findings always need fixing
- One reviewer for Restricted data — require two reviewers for highest sensitivity
- Compliance evidence as afterthought — collect evidence during testing, not after deployment

---

## 3. Defining Dependencies and Sequencing (ref: PM-BP-022)

**Best Practices:**
- Map dependencies before work begins — discovering dependencies mid-sprint causes delays
- Use the four dependency types: Finish-to-Start, Start-to-Start, Finish-to-Finish, External
- Identify the critical path — the longest chain of dependent tasks determines the minimum delivery time
- Parallelize where possible — frontend and backend can often run in parallel if the API contract is defined first
- External dependencies are the highest risk — identify them early and have mitigation plans
- Visualize the dependency chain — a simple diagram prevents misunderstandings

**Dependency Identification Checklist:**

| Ask This | If Yes → Dependency |
|---|---|
| Does this task need another task's output? | Finish-to-Start |
| Can this task start when another starts (but not before)? | Start-to-Start |
| Does this task need a deployed environment? | Depends on OPS deployment task |
| Does this task need an external system/API? | External dependency — highest risk |
| Does this task need another team's deliverable? | Cross-team dependency — coordinate early |
| Does this task need compliance/legal approval? | Approval dependency — submit early |

**Parallelization Opportunities:**

| Can Run in Parallel | Condition |
|---|---|
| Backend dev + Frontend dev | API contract (OpenAPI spec) defined first |
| Backend dev + Unit test dev | Start-to-Start — tests written alongside code |
| SAST scan + SCA scan | Both run on completed code |
| Security code review + Unit test execution | Both happen after code complete |
| Documentation + Testing | Documentation can start during testing phase |
| Rollback plan creation + Development | Rollback plan created in parallel with dev |

**Sequencing Anti-Patterns:**

| Anti-Pattern | Problem | Fix |
|---|---|---|
| All tasks sequential | Unnecessarily long delivery time | Identify parallelization opportunities |
| No dependency mapping | Blocked tasks discovered mid-sprint | Map dependencies before sprint starts |
| External dependencies not tracked | Surprise delays from other teams | Identify external deps in sprint planning |
| Security tasks at the end only | No time for remediation if findings | SAST/SCA during dev; DAST after staging |
| Compliance review after development | Rework if compliance rejects | Submit compliance review before dev starts |

---

## 4. Including Rollback and Remediation Tasks (ref: PM-BP-023)

**Best Practices:**
- Every story that deploys to production must have a rollback plan — no exceptions
- The rollback plan must be created before deployment, not during an incident
- Rollback procedures must be tested in staging before production deployment
- Define clear rollback triggers — what conditions warrant a rollback
- Define the rollback decision maker — who has authority to trigger rollback
- Include data remediation — what happens to data created/modified during the failed deployment
- Include communication plan — who gets notified during rollback

**Rollback Plan Components:**

| Component | What to Define |
|---|---|
| Trigger conditions | Error rate > 1%, P1 incident, data corruption, security vulnerability discovered |
| Decision maker | On-call lead, release manager, or designated authority |
| Time window | How long after deployment can rollback be executed (e.g., within 4 hours) |
| Rollback steps | Numbered, step-by-step procedure anyone on the team can follow |
| Data handling | How to handle data created during the deployment window |
| Verification | How to confirm rollback was successful (health checks, smoke tests, monitoring) |
| Communication | Notify: stakeholders, support team, affected customers (if customer-facing) |
| Post-rollback | RCA, fix-forward plan, re-deployment timeline |

**Rollback Strategies by Change Type:**

| Change Type | Strategy | Best Practice |
|---|---|---|
| Code deployment | Redeploy previous version | Use feature flags for instant disable; blue-green deployment for zero-downtime rollback |
| Additive DB change (new column/table) | Keep schema, rollback code only | Design backward-compatible migrations — new code works with old schema |
| Destructive DB change (drop/rename) | Reverse migration script | Never deploy destructive changes without tested reverse script; prefer additive-only |
| Data migration | Data remediation script | Backup before migration; script to restore original state; preserve audit trail |
| Config change | Revert via version control | All configs in version control; automated config deployment |
| Feature flag change | Toggle flag back | Instant rollback — preferred for high-risk features |
| Third-party integration | Circuit breaker / feature flag | Disable integration without affecting core functionality |

**Remediation Task Components:**

| Component | What to Define |
|---|---|
| Defect triage | How to assess severity and impact of discovered issues |
| Hotfix process | Fast-track process for critical fixes (still requires security scan) |
| Data fix | How to correct corrupted/incorrect data with audit trail |
| Customer communication | Template for notifying affected customers |
| Incident report | Timeline, impact, root cause, corrective actions |
| Preventive measures | What changes prevent recurrence (test gap, monitoring gap, process gap) |

**Common Mistakes:**
- No rollback plan — "we'll figure it out if something goes wrong" leads to extended outages
- Untested rollback — a rollback plan that hasn't been tested in staging will likely fail in production
- No data remediation plan — rolling back code doesn't fix corrupted data
- Rollback plan without decision maker — nobody knows who can authorize rollback during an incident
- No time estimate for rollback — stakeholders need to know how long recovery takes
- Destructive DB changes without reverse script — irreversible changes are the highest risk

---

## 5. Task Estimation

**Best Practices:**
- Each task should be completable by one person in 1-2 days maximum — larger tasks should be split
- Include remediation time in security scan estimates — findings always need fixing
- Include review/approval wait time in compliance task estimates
- Estimate rollback testing separately — it's real work, not a checkbox
- Account for environment setup time in testing and deployment tasks

**Estimation Guide:**

| Task Type | Typical Range | Factors That Increase |
|---|---|---|
| Backend implementation | 4-16h | Complex business rules, multiple integrations, high data sensitivity |
| Frontend implementation | 4-12h | Complex forms, real-time validation, accessibility requirements |
| Unit test development | 2-8h | Complex logic, many edge cases, high coverage target |
| Integration test development | 4-8h | Multiple integration points, async flows, error scenarios |
| SAST scan + remediation | 2-4h | Large codebase, many findings, complex fixes |
| DAST scan + remediation | 2-4h | Many endpoints, auth flows, complex input validation |
| Security code review | 2-4h | Restricted data, auth logic, crypto implementation |
| Compliance review | 2-8h | Multiple regulations, new data processing, cross-border |
| Rollback plan + testing | 2-4h | Database changes, data migrations, multi-service deployment |
| Monitoring setup | 2-4h | Custom dashboards, complex alerting rules, SLA monitoring |

---

## 6. Task Assignment

**Best Practices:**
- Security tasks should be assigned to someone other than the developer who wrote the code
- Compliance tasks should be assigned to or reviewed by the Compliance Officer
- Code review for Restricted data requires two reviewers with security awareness
- Rollback plan should be reviewed by the on-call/operations team who would execute it
- Testing tasks can be assigned to the developer (unit tests) or QA (integration, security, compliance tests)

**Assignment Matrix:**

| Task Type | Primary Assignee | Reviewer/Approver |
|---|---|---|
| DEV | Developer | Peer reviewer (2 for Restricted) |
| SEC (scanning) | Developer or DevSecOps | Security Engineer |
| SEC (code review) | Peer developer (not the author) | Security Engineer for Restricted |
| TST (unit) | Developer | QA Engineer |
| TST (integration, security) | QA Engineer | Security Engineer |
| TST (compliance) | QA Engineer | Compliance Officer |
| CMP | Developer (preparation) | Compliance Officer (approval) |
| OPS | DevOps Engineer | Release Manager |
| RBK | Developer + DevOps | On-call Lead |
| DOC | Developer | Tech Lead |

---

## 7. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Single "implement feature" task | Security, testing, compliance skipped | Separate by type (PM-BP-020) |
| No security tasks | Vulnerabilities shipped to production | Mandatory SAST/SCA/review tasks (PM-BP-021) |
| Dependencies discovered mid-sprint | Blocked work, wasted sprint capacity | Map dependencies before sprint (PM-BP-022) |
| No rollback plan | Extended outages on failed deployments | Rollback task for every production change (PM-BP-023) |
| Security scan with zero remediation time | Findings ignored or deferred indefinitely | Estimate scan + remediation together |
| All tasks sequential | Unnecessarily long delivery | Parallelize where possible |
| Tasks too large (> 2 days) | Hard to track, hard to estimate, hard to complete | Split into 1-2 day tasks |
| Same person develops and reviews | Bias, missed issues | Different assignee for review tasks |
| Compliance review after deployment | Rework or rollback if rejected | Compliance review before development |
| Rollback plan untested | Fails when needed most | Test rollback in staging |
| No monitoring setup task | Blind to production issues | Always include monitoring task |
| Documentation deferred indefinitely | Knowledge loss, onboarding difficulty | Documentation task in every story |
