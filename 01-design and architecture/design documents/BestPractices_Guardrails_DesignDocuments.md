# Best Practices & Guardrails for HLD and LLD Documents
### Banking Domain — Agentic Knowledge Base

---

## Part A: Best Practices

### 1. HLD Composition Best Practices

**Structure:**
- Start with business context, not technology — explain what and why before how
- Use C4 model consistently: Context (L1) → Container (L2) → Component (L3) in HLD
- Every diagram must have a legend, title, and version
- Every integration must show protocol, authentication method, and data classification
- Every data store must show technology, encryption, and classification

**Content Quality:**
- Trace every design decision to a requirement (FR, NFR, CR, SR)
- Document what was considered and rejected (ADRs), not just what was chosen
- Include degraded-mode behavior for every external dependency
- Include security controls mapped to STRIDE threats
- Include compliance controls mapped to specific regulation articles
- Include performance targets with measurement methods

**Diagrams Required:**

| Diagram | Purpose | Audience | When Required |
|---|---|---|---|
| C4 Context | System in environment | All stakeholders | Always |
| C4 Container | Technology choices | Architects, leads | Always |
| C4 Component | Internal structure per service | Developers, leads | Always |
| Deployment | Infrastructure mapping | DevOps, architects | Always |
| Data flow | Data movement with classification | Security, compliance | Always |
| Sequence (key flows) | Interaction between components | Developers | Always |
| ER (conceptual) | Data relationships | Developers, DBAs | Always |
| Network topology | Network architecture | DevOps, security | If infra changes |
| Threat model | Security threats and controls | Security team | If Confidential/Restricted |
| State machine | Entity lifecycle | Developers | If complex state transitions |

**Common Mistakes:**
- HLD that is actually an LLD — too much implementation detail
- No requirements traceability — can't verify design satisfies requirements
- Missing security section — security must be designed, not added later
- No DR/degraded mode — only happy path designed
- Diagrams without legends or version numbers
- No milestones or project plan — design without delivery timeline is incomplete

### 1b. Milestones & Project Plan Best Practices

**Structure:**
- Define 7-9 milestones covering: Design → Development → Integration → Security → Performance → Compliance → UAT → Deployment → Post-Deployment
- Every milestone must have explicit exit criteria — not just a date
- Include a phased plan showing activities, deliverables, and dependencies per phase
- Include a dependency timeline showing the critical path and external dependencies
- Include a resource plan mapping roles to phases with allocation percentages

**Content Quality:**
- Milestones must be sequenced with gates — security and compliance milestones must precede deployment
- External dependencies must be identified with owners, expected dates, and mitigation if delayed
- Include a risk-adjusted timeline — identify top 3-5 delivery risks with impact on timeline and contingency
- Include Go/No-Go criteria — explicit checklist that must all pass before production deployment
- Plan must be realistic — account for review cycles, remediation time, environment provisioning

**Common Mistakes:**
- No exit criteria on milestones — "Development Complete" without defining what "complete" means
- Security and compliance as last-minute activities — must be parallel, not sequential at the end
- No external dependency tracking — surprises from other teams cause delays
- No contingency for risks — optimistic plan with no buffer
- Go/No-Go criteria defined after deployment — must be defined upfront in the design

### 2. LLD Composition Best Practices

**Structure:**
- One LLD per microservice per feature — not a monolithic document
- Reference the HLD — don't repeat; extend with implementation detail
- Include actual code structures (package layout, class names, interfaces)
- Include actual SQL (schema, indexes, migrations, queries)
- Include actual API contracts (OpenAPI snippets, request/response examples)
- Include actual configuration (connection pools, timeouts, circuit breaker settings)

**Content Quality:**
- Every class/interface must have a clear responsibility
- Every database table must show columns, types, constraints, indexes, encryption, and audit triggers
- Every API endpoint must show request/response schemas, validation rules, error codes, auth requirements
- Every integration must show timeout, retry, circuit breaker, and fallback configuration
- Every error scenario must show HTTP status, error code, user message, logging level, and alert trigger
- Include state machine for entities with complex lifecycles

**Diagrams Required:**

| Diagram | Purpose | When Required |
|---|---|---|
| Package/class diagram | Code structure | Always |
| Sequence (all flows) | Happy path + every error path | Always |
| State machine | Entity lifecycle transitions | If entity has > 3 states |
| ER (physical) | Actual database schema | Always |
| Integration detail | Timeout, retry, circuit breaker per dependency | Always |

**Common Mistakes:**
- LLD that is actually an HLD — too abstract, no implementation detail
- Missing error/exception flows — only happy path documented
- No database detail — schema, indexes, queries missing
- No configuration — timeouts, pool sizes, thresholds missing
- No testing strategy — what and how to test

### 3. Diagram Standards

**All Diagrams Must Include:**
- Title and version number
- Legend explaining all symbols, colors, and line styles
- Data classification labels on all data flows and stores
- Protocol and authentication on all communication lines
- Trust boundaries clearly marked
- Date of last update

**Color Coding Standard:**

| Color | Meaning |
|---|---|
| Blue | Internal services/components |
| Green | External systems |
| Red | Security boundaries / PCI-DSS CDE |
| Orange | Data stores |
| Yellow | Message brokers / queues |
| Gray | Infrastructure (load balancers, gateways) |
| Purple | Monitoring / observability |

**Data Flow Labels:**
```
[Source] --protocol/auth/classification--> [Destination]
Example: [Transfer Service] --REST/mTLS/Confidential--> [Account Service]
```

---

## Part B: Guardrails

### 4. HLD Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-HG-001 | Every HLD must have a Requirements Traceability section mapping design elements to FR, NFR, CR, SR IDs | Reject if traceability is missing |
| DD-HG-002 | Every HLD must include C4 Context (L1) and Container (L2) diagrams at minimum | Reject if L1 or L2 diagrams are missing |
| DD-HG-003 | Every HLD must include a Security Architecture section with STRIDE threat mapping and controls | Reject if security section is missing |
| DD-HG-004 | Every HLD must include a Compliance Controls section mapping regulations to design controls | Reject if compliance section is missing for regulated features |
| DD-HG-005 | Every HLD must include Data Architecture with classification, encryption, and retention per data element | Reject if data architecture is missing |
| DD-HG-006 | Every HLD must include Availability & DR section with RTO, RPO, degraded-mode behavior | Reject if DR section is missing |
| DD-HG-007 | Every HLD must include Performance targets with measurement methods | Reject if performance section is missing |
| DD-HG-008 | Every HLD must include at least one ADR for each significant architectural decision | Flag if no ADRs are documented |
| DD-HG-009 | Every HLD must include a Risks & Mitigations section | Reject if risk section is missing |
| DD-HG-010 | Every HLD must be reviewed and approved by: Solution Architect, Security Architect, Compliance Officer (if regulated) | Reject if required approvals are missing |

### 5. LLD Structural Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-LG-001 | Every LLD must reference its parent HLD document | Reject if HLD reference is missing |
| DD-LG-002 | Every LLD must include package/class structure with clear layer separation (domain, application, infrastructure) | Reject if code structure is missing |
| DD-LG-003 | Every LLD must include physical database schema with columns, types, constraints, indexes, encryption, and audit triggers | Reject if database detail is missing |
| DD-LG-004 | Every LLD must include API specification detail (request/response schemas, validation rules, error codes, auth) | Reject if API detail is missing |
| DD-LG-005 | Every LLD must include sequence diagrams for happy path AND all error/exception flows | Reject if only happy path is documented |
| DD-LG-006 | Every LLD must include integration detail with timeout, retry, circuit breaker, and fallback per dependency | Reject if integration detail is missing |
| DD-LG-007 | Every LLD must include error handling matrix (scenario, HTTP status, error code, message, logging, alert) | Reject if error handling is missing |
| DD-LG-008 | Every LLD must include testing strategy with coverage targets per test type | Reject if testing strategy is missing |
| DD-LG-009 | Every LLD must include configuration per environment (dev, staging, production) | Flag if configuration is missing |
| DD-LG-010 | Every LLD must be reviewed by: Tech Lead, Architect, Security Engineer, DBA (if DB changes) | Reject if required reviews are missing |

### 6. Security Design Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-SG-001 | Every design must include authentication and authorization design per component/endpoint | Reject if auth design is missing |
| DD-SG-002 | Every design must show encryption at rest and in transit for all data classified Confidential or Restricted | Reject if encryption design is missing for sensitive data |
| DD-SG-003 | Every design must include audit trail design specifying events, fields, and retention | Reject if audit design is missing |
| DD-SG-004 | Every design must include input validation rules per API field | Reject if validation rules are missing |
| DD-SG-005 | Every design must show sensitive data handling (masking in responses, logs, non-prod) | Reject if sensitive data handling is not documented |
| DD-SG-006 | Designs involving Restricted data must include a threat model diagram | Reject if Restricted-data design has no threat model |
| DD-SG-007 | Every external integration must show security controls (TLS version, auth method, data classification) | Reject if integration security is not documented |
| DD-SG-008 | Designs must not introduce security anti-patterns (hardcoded secrets, SQL concatenation, sensitive data in URLs) | Reject if security anti-patterns are present |

### 7. Data Design Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-DG-001 | Every data element must have a classification (Public/Internal/Confidential/Restricted) in the design | Reject if any data element lacks classification |
| DD-DG-002 | Data flow diagrams must show classification labels on every flow and storage point | Reject if data flows lack classification labels |
| DD-DG-003 | Cross-boundary data flows (PCI-DSS CDE, cross-border) must be explicitly identified and justified | Reject if cross-boundary flows are not identified |
| DD-DG-004 | Database designs must use DECIMAL/NUMERIC for monetary amounts — FLOAT/DOUBLE is prohibited | Reject if floating-point is used for money |
| DD-DG-005 | External-facing IDs must use UUID — sequential integers must not be exposed | Reject if sequential IDs are exposed externally |
| DD-DG-006 | Every new data store must specify retention policy, backup strategy, and DR replication | Reject if data lifecycle is not specified |

### 8. Compliance Design Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-CG-001 | Designs for regulated features must map each applicable regulation to a specific design control | Reject if regulation-to-control mapping is missing |
| DD-CG-002 | Designs must specify how compliance evidence is generated and stored | Reject if evidence generation is not designed |
| DD-CG-003 | Designs involving PII must include GDPR controls (legal basis, minimization, erasure capability, processing records) | Reject if GDPR controls are missing for PII features |
| DD-CG-004 | Designs involving payments must include PSD2 SCA controls | Reject if SCA is missing for payment features |
| DD-CG-005 | Designs involving card data must show PCI-DSS CDE boundary and controls | Reject if PCI-DSS boundary is not defined for card features |

### 9. Performance & Resilience Design Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-PG-001 | Every design must specify performance targets (response time, throughput) with measurement method | Reject if performance targets are missing |
| DD-PG-002 | Every external dependency must have circuit breaker, timeout, and retry configuration documented | Reject if resilience configuration is missing |
| DD-PG-003 | Every external dependency must have a fallback strategy documented | Reject if fallback is missing |
| DD-PG-004 | Designs must specify caching strategy with TTL and invalidation for each cacheable data type | Flag if caching strategy is not documented |
| DD-PG-005 | Database designs must include index strategy with expected query patterns and target latencies | Reject if index strategy is missing |

### 10. Diagram Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-DiG-001 | Every diagram must have a title, version, legend, and last-updated date | Reject if any diagram lacks these elements |
| DD-DiG-002 | Every data flow must show protocol, authentication method, and data classification | Reject if data flows are unlabeled |
| DD-DiG-003 | Trust boundaries must be clearly marked on all architecture diagrams | Reject if trust boundaries are missing |
| DD-DiG-004 | PCI-DSS CDE boundary must be explicitly shown if card data is in scope | Reject if CDE boundary is missing for card features |
| DD-DiG-005 | Diagrams must use consistent color coding per the standard | Flag if non-standard colors are used |

### 11. Review & Approval Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-RG-001 | HLD must be approved before LLD work begins | Reject if LLD is created without approved HLD |
| DD-RG-002 | Design documents must be version-controlled alongside code | Reject if designs are not in version control |
| DD-RG-003 | Design changes after approval require re-review by original approvers | Process guardrail — enforce via change review |
| DD-RG-004 | Design documents must be updated when implementation deviates from design | Flag if implementation differs from approved design |
| DD-RG-005 | Designs for Restricted data features require Security Architect sign-off | Reject if Restricted-data design lacks security approval |
| DD-RG-006 | Designs for regulated features require Compliance Officer sign-off | Reject if regulated design lacks compliance approval |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot proceed to development | Missing security section, no requirements traceability, no error flows, no DB detail, missing approvals |
| **Flag** | Address before development completes | Missing ADRs, no caching strategy, non-standard diagram colors, missing configuration |
| **Process** | Enforced via workflow | HLD before LLD, re-review on changes, version control |

---

### 12. Milestones & Project Plan Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| DD-MG-001 | Every HLD must include a Milestones & Project Plan section with phased delivery plan | Reject if milestones/plan is missing |
| DD-MG-002 | Milestones must include explicit exit criteria — not just dates | Reject if milestones have no exit criteria |
| DD-MG-003 | Plan must include security validation milestone (SAST/DAST/SCA complete) before deployment | Reject if security milestone is missing |
| DD-MG-004 | Plan must include compliance sign-off milestone before deployment for regulated features | Reject if compliance milestone is missing for regulated features |
| DD-MG-005 | Plan must include performance validation milestone before deployment for customer-facing features | Reject if performance milestone is missing for customer-facing features |
| DD-MG-006 | Plan must identify external dependencies with owners and mitigation for delays | Reject if external dependencies are not identified |
| DD-MG-007 | Plan must include Go/No-Go criteria for production deployment | Reject if Go/No-Go criteria are missing |
| DD-MG-008 | Plan must include rollback testing milestone before production deployment | Reject if rollback testing is not in the plan |
| DD-MG-009 | Plan must include post-deployment verification milestone | Flag if post-deployment verification is missing |
| DD-MG-010 | Risk-adjusted timeline must account for top 3 delivery risks with contingency | Flag if no risk-adjusted timeline exists |

---

## Design Document Checklist

### HLD Checklist

| # | Check | Section |
|---|---|---|
| 1 | Requirements traceability (FR, NFR, CR, SR mapped) | §2 |
| 2 | C4 Context diagram (L1) | §3.1 |
| 3 | C4 Container diagram (L2) | §3.2 |
| 4 | Deployment diagram | §3.3 |
| 5 | Component design with bounded contexts | §4 |
| 6 | API design summary | §4.3 |
| 7 | Domain event catalog | §4.4 |
| 8 | Data classification per element | §5.1 |
| 9 | Data flow diagram with classification labels | §5.3 |
| 10 | Security controls (STRIDE mapping) | §6.2 |
| 11 | Compliance controls (regulation mapping) | §6.3 |
| 12 | Audit trail design | §6.4 |
| 13 | Performance targets | §7.1 |
| 14 | Scalability design | §7.2 |
| 15 | Availability & DR (RTO, RPO, degraded mode) | §8 |
| 16 | Monitoring & observability | §9 |
| 17 | Risks & mitigations | §10 |
| 18 | ADRs for key decisions | §11 |
| 19 | Milestones with exit criteria | §13.1 |
| 20 | Phased project plan with dependencies | §13.2-13.3 |
| 21 | Resource plan | §13.4 |
| 22 | Risk-adjusted timeline | §13.5 |
| 23 | Go/No-Go criteria for deployment | §13.6 |
| 24 | All required approvals obtained | Sign-off |

### LLD Checklist

| # | Check | Section |
|---|---|---|
| 1 | HLD reference | Metadata |
| 2 | Package/class structure | §2.1 |
| 3 | Domain model detail (fields, types, constraints) | §2.2 |
| 4 | State machine (if applicable) | §2.3 |
| 5 | API specification (request/response/errors/auth) | §3 |
| 6 | Sequence diagrams (happy + error paths) | §4 |
| 7 | Physical database schema (DDL, indexes, triggers) | §5.1 |
| 8 | Migration scripts with rollback | §5.2 |
| 9 | Query patterns with target latencies | §5.3 |
| 10 | Integration detail (timeout, retry, circuit breaker, fallback) | §6 |
| 11 | Input validation rules per field | §7.1 |
| 12 | Authorization checks | §7.2 |
| 13 | Sensitive data handling matrix | §7.3 |
| 14 | Error handling matrix | §8 |
| 15 | Testing strategy with coverage targets | §9 |
| 16 | Configuration per environment | §10 |
| 17 | All required reviews completed | Sign-off |
