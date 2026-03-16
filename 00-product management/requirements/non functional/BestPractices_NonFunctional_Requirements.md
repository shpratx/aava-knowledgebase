# Best Practices for Composing Non-Functional Requirements

---

## 1. Writing the Requirement ID & Title

**Best Practices:**
- Use consistent naming: `NFR-[Category]-[Seq]` where category maps to the quality attribute
- Standard categories: PERF (Performance), SCAL (Scalability), AVAIL (Availability), REL (Reliability), USAB (Usability), MAINT (Maintainability), PORT (Portability)
- Titles should state the quality attribute and scope: "Fund Transfer API Response Time" not "System should be fast"

**Common Mistakes:**
- Mixing NFRs with functional requirements ("System shall display dashboard quickly" — split into FR for display + NFR for speed)
- Using subjective titles ("Good performance for payments")

---

## 2. Choosing the Right Category

**Decision Guide:**

| If the requirement is about... | Category | Example |
|---|---|---|
| Speed, latency, throughput | Performance | API response time < 500ms |
| Handling growth in users/data/transactions | Scalability | Support 10,000 concurrent users |
| Uptime, failover, disaster recovery | Availability | 99.95% uptime SLA |
| Consistency, fault tolerance, data integrity | Reliability | Zero data loss on failover |
| Ease of use, accessibility, UX | Usability | WCAG 2.1 AA compliance |
| Code quality, deployability, technical debt | Maintainability | 80%+ code coverage |
| Cross-platform, migration, interoperability | Portability | Run on AWS and Azure |

**Banking-Specific Consideration:**
- Regulatory requirements often dictate NFRs (e.g., PCI-DSS mandates encryption performance, DORA mandates recovery times)
- Always check if an NFR has a regulatory driver — if so, cross-reference to the Compliance Requirement

---

## 3. Writing the Description

**Best Practices:**
- State the quality attribute, the scope (which system/component), and the condition
- Be explicit about load conditions: "under normal load" vs. "under peak load" vs. "under stress"
- Separate steady-state from burst/peak requirements
- Avoid vague language: never use "fast", "responsive", "scalable", "reliable" without quantification

**Formula:**
> The [system/component/API] shall [achieve/maintain/support] [specific metric] [under specific conditions] [as measured by specific method].

**Good Example:**
> The Fund Transfer API shall maintain p95 response time below 500ms under a sustained load of 500 concurrent users, as measured by Gatling load tests.

**Bad Example:**
> The system should be fast and responsive.

---

## 4. Defining Metrics & Target Values

**Best Practices:**
- Every NFR must have at least one quantifiable metric
- Use industry-standard measurement units (ms, TPS, %, hours)
- Specify percentile for latency metrics (p50, p95, p99) — averages hide outliers
- Define separate targets for normal load, peak load, and degraded mode
- Include both threshold (minimum acceptable) and target (desired) values where appropriate
- Align targets with SLAs committed to customers/regulators

**Banking-Specific Targets:**

| Quality Attribute | Metric | Typical Banking Target |
|---|---|---|
| API Latency | p95 response time | < 500ms (customer-facing), < 200ms (internal) |
| Page Load | Time to interactive | < 3 seconds |
| Throughput | Transactions per second | 500-5,000 TPS depending on channel |
| Availability | Uptime | 99.95% - 99.99% for critical systems |
| RTO | Recovery time | < 1 hour (critical), < 4 hours (standard) |
| RPO | Data loss window | < 15 minutes (critical), < 1 hour (standard) |
| Failover | Switchover time | < 30 seconds (active-active), < 5 minutes (active-passive) |
| Batch Processing | Throughput | Complete EOD batch within 4-hour window |
| Concurrent Users | Capacity | 10,000-100,000 depending on channel |

**Common Mistakes:**
- Using averages instead of percentiles for latency
- Not specifying load conditions ("fast" under what load?)
- Setting unrealistic targets that can't be tested or achieved
- Not distinguishing between customer-facing and internal system targets

---

## 5. Specifying Measurement Methods

**Best Practices:**
- Define how the metric will be measured before development starts
- Specify the tool or technique (JMeter, Gatling, Lighthouse, Datadog, SonarQube)
- Define the test scenario (user journey, data volume, concurrent load)
- Specify measurement frequency (continuous monitoring vs. periodic testing)
- Define the environment where measurement occurs (production, staging, dedicated perf environment)

**Measurement Method by Category:**

| Category | Tools/Methods | When |
|---|---|---|
| Performance | Gatling, JMeter, k6 (load tests); Datadog, Dynatrace (APM) | Pre-release load test + continuous production monitoring |
| Scalability | Load tests with incremental scaling; infrastructure auto-scaling tests | Pre-release + quarterly capacity review |
| Availability | Uptime monitoring (Pingdom, Datadog); failover drills | Continuous + quarterly DR tests |
| Reliability | Chaos engineering (Chaos Monkey); data integrity checks | Monthly resilience tests |
| Usability | Lighthouse, axe (accessibility); user testing sessions | Per release + annual accessibility audit |
| Maintainability | SonarQube (code quality); CI/CD pipeline metrics | Continuous |

---

## 6. Writing Acceptance Criteria

**Best Practices:**
- Acceptance criteria must be pass/fail — no ambiguity
- Include the load condition, duration, and success threshold
- Specify what constitutes a failure (error rate, timeout count)
- Include warm-up period exclusions if applicable
- Define the test data requirements (volume, variety)

**Patterns for Banking NFR Acceptance Criteria:**

```
# Performance
PASS when: p95 response time ≤ 500ms AND error rate < 0.1%
  at 500 concurrent users sustained for 30 minutes
  after 5-minute warm-up period
  using production-equivalent test data (1M accounts, 10M transactions)

# Availability
PASS when: system recovers within 60 seconds of primary node failure
  AND zero transactions are lost during failover
  AND all in-flight transactions are either completed or safely rolled back

# Scalability
PASS when: system auto-scales from 2 to 8 instances within 60 seconds
  when concurrent users increase from 1,000 to 5,000
  AND p95 response time remains below 500ms during scaling

# Maintainability
PASS when: code coverage ≥ 80% across all modules
  AND critical path coverage ≥ 90%
  AND technical debt ratio < 10%
  AND zero critical/blocker SonarQube issues
```

---

## 7. Environment Applicability

**Best Practices:**
- Specify which environments the NFR applies to (not all NFRs apply everywhere)
- Production NFRs are the primary target
- Staging/pre-prod should mirror production NFRs for validation
- Dev/test environments may have relaxed targets but should still be defined
- Specify if dedicated performance testing environment is required

**Typical Banking Environment Matrix:**

| NFR Category | Dev | Test | Staging | Production |
|---|---|---|---|---|
| Performance | Relaxed (2x target) | Relaxed (1.5x target) | Production target | Production target |
| Availability | Best effort | Best effort | 99.9% | 99.95%+ |
| Security | Full enforcement | Full enforcement | Full enforcement | Full enforcement |
| Scalability | Single instance | Reduced scale | Production-like | Full scale |
| Maintainability | Full enforcement | Full enforcement | Full enforcement | Full enforcement |

---

## 8. Handling Dependencies & Assumptions

**Best Practices:**
- Document infrastructure assumptions (CPU, memory, network bandwidth, storage IOPS)
- Document external dependency assumptions (core banking latency, payment network availability)
- Identify shared resource constraints (database connection pools, thread pools, message queue capacity)
- State assumptions about data volume and growth rate
- Document third-party SLA dependencies

**Banking-Specific Assumptions to Document:**
- Core banking system response time (often the bottleneck)
- Payment network availability (SWIFT, local clearing)
- HSM throughput for encryption operations
- Database replication lag for read replicas
- Network latency between data centers (for DR)
- Batch window constraints (EOD processing, regulatory reporting)

---

## 9. Relating NFRs to Functional Requirements

**Best Practices:**
- Every NFR should reference at least one FR it applies to
- Some NFRs are system-wide (e.g., availability) — reference "All" or the system name
- Some NFRs are specific to a function (e.g., transfer API latency) — reference the specific FR
- Create a traceability matrix: FR → NFR mapping
- Ensure critical FRs have corresponding performance, availability, and security NFRs

**Priority Matrix:**

| FR Priority | Required NFRs |
|---|---|
| Critical (e.g., fund transfer) | Performance + Availability + Scalability + Reliability |
| High (e.g., account inquiry) | Performance + Availability |
| Medium (e.g., statement download) | Performance |
| Low (e.g., preference update) | Basic performance |

---

## 10. Data Classification in NFRs (ref: PM-BP-008)

**Best Practices:**
- NFR targets should vary by data classification level — Restricted data needs stricter performance, availability, and security NFRs
- Specify which data classification levels the NFR applies to in the Data Classification Scope field
- Encryption performance overhead must be factored into performance targets for Confidential/Restricted data
- Backup and replication NFRs must be tiered by data classification
- Data masking/anonymization performance in non-prod environments should be specified

**Classification-Driven NFR Differentiation:**

| NFR Aspect | Restricted | Confidential | Internal | Public |
|---|---|---|---|---|
| Availability | 99.99% | 99.95% | 99.9% | 99% |
| RTO | < 1 hour | < 4 hours | < 24 hours | Best effort |
| RPO | < 15 minutes | < 1 hour | < 4 hours | < 24 hours |
| Backup frequency | Continuous | Hourly | Daily | Weekly |
| Audit log depth | Full access logging | Modification logging | Auth event logging | Minimal |
| Performance testing | Must include encryption overhead | Must include encryption overhead | Standard | Standard |

---

## 11. Audit Trail & Logging NFRs (ref: PM-BP-009)

**Best Practices:**
- Define NFRs for the logging infrastructure itself — it's a critical system
- Specify log ingestion latency (time from event to searchable in central store)
- Specify log query performance (response time for forensic queries)
- Specify log availability SLA (logs must be available when needed for incident response)
- Specify log storage capacity planning (retention × volume × growth rate)
- Specify log integrity verification method and frequency

**Key Audit/Logging NFR Metrics:**

| Metric | Target | Why It Matters |
|---|---|---|
| Log ingestion latency | < 5 seconds | Real-time alerting depends on timely log availability |
| Log query response (24h window) | < 10 seconds | Incident response speed depends on query performance |
| Log availability | 99.99% | Logs unavailable during incident = blind spot |
| Log completeness | 100% of defined events | Missing logs = audit failure |
| Log integrity verification | Daily automated check | Tampered logs invalidate audit evidence |
| Log storage growth | Plan for 3-year projection | Running out of log storage = data loss |

---

## 12. DR/BCP NFRs (ref: PM-BP-010)

**Best Practices:**
- Define RTO and RPO as explicit NFRs, not just assumptions
- Specify failover behavior NFRs: automatic vs. manual, switchover time, data consistency
- Define degraded-mode performance targets — what's acceptable during partial failure
- Specify DR test frequency and success criteria as NFRs
- Define geographic redundancy requirements based on business criticality
- Specify backup verification NFRs — backups must be tested, not just created
- Define BCP activation time — how quickly can the organization switch to continuity mode

**DR/BCP NFR Patterns:**
```
# Failover
NFR: System shall failover to DR site within 30 seconds of primary failure
  with zero loss of committed transactions
  and p95 response time not exceeding 2x normal during failover window.

# Backup Verification
NFR: Automated backup restoration test shall run weekly
  restoring to a test environment within 2 hours
  with 100% data integrity verification.

# Degraded Mode
NFR: During single-AZ failure, system shall maintain 80% of normal throughput
  with p95 response time not exceeding 1.5x normal target
  for all Tier 1 functions.
```

---

## 13. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| "System should be fast" | Unmeasurable, untestable | Specify metric, target, and conditions |
| Average response time only | Hides tail latency issues | Use percentiles (p95, p99) |
| No load condition specified | Target is meaningless without context | Always state concurrent users/TPS |
| Same target for all APIs | Over-engineering or under-engineering | Tier targets by criticality |
| NFR without measurement method | Can't verify compliance | Define tool and test scenario |
| Ignoring degraded mode | System has no graceful degradation | Define acceptable degraded performance |
| Not testing NFRs | Discovered in production | Include NFR tests in CI/CD and release gates |
| Unrealistic targets | Wasted effort, never achieved | Benchmark current state, set achievable targets |
| Same NFR targets for all data classifications | Over/under-engineering | Tier targets by data sensitivity |
| No logging infrastructure NFRs | Logs unavailable during incidents | Define log ingestion, query, and availability SLAs |
| DR/BCP as documentation only | Untested plans fail when needed | Define testable DR NFRs with pass/fail criteria |
| No degraded-mode performance targets | Unknown behavior during partial failure | Define acceptable performance during failover |
