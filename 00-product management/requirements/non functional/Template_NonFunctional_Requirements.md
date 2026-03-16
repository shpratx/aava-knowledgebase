# Non-Functional Requirements Template

---

## Template Fields

| Field | Description |
|---|---|
| **Requirement ID** | NFR-[Category]-[Seq] (e.g., NFR-PERF-001) |
| **Title** | Short descriptive title |
| **Category** | Performance / Scalability / Availability / Reliability / Usability / Maintainability / Portability / DR-BCP / Audit-Logging |
| **Priority** | Critical / High / Medium / Low |
| **Related Functional Req(s)** | Linked FR IDs |
| **Data Classification Scope** | Which data classification levels this NFR applies to (ref: PM-BP-008). E.g., "Confidential and Restricted data" or "All" |
| **Description** | Detailed description of the non-functional expectation |
| **Metric** | Quantifiable measure (e.g., response time in ms, uptime %) |
| **Target Value** | Specific threshold or SLA |
| **Measurement Method** | How this will be measured (tool, test type, monitoring) |
| **Acceptance Criteria** | Pass/fail conditions |
| **Environment Applicability** | Production / Staging / All |
| **Assumptions** | Infrastructure, load, or configuration assumptions |

---

## Sub-Categories & Standard Targets for Banking

### Performance
| Metric | Target | Measurement |
|---|---|---|
| API response time (p95) | < 500ms | Load testing (JMeter/Gatling) |
| API response time (p99) | < 1000ms | APM monitoring |
| Page load time | < 3 seconds | Lighthouse / RUM |
| Database query time | < 100ms for transactional queries | Query profiling |
| Batch processing throughput | Per SLA (e.g., 1M records/hour) | Batch monitoring |

### Scalability
| Metric | Target | Measurement |
|---|---|---|
| Concurrent users | Support 10,000+ concurrent sessions | Load testing |
| Transaction throughput | 500 TPS minimum | Stress testing |
| Horizontal scaling | Auto-scale within 60 seconds | Infrastructure testing |
| Data volume | Support 5 years of transaction history | Capacity planning |

### Availability & Reliability
| Metric | Target | Measurement |
|---|---|---|
| System uptime | 99.95% (excludes planned maintenance) | Monitoring (Datadog/Dynatrace) |
| RTO (Recovery Time Objective) | < 1 hour for critical systems | DR testing |
| RPO (Recovery Point Objective) | < 15 minutes | Backup verification |
| Failover time | < 30 seconds | Failover testing |
| Mean Time Between Failures | > 720 hours | Incident tracking |

### Usability
| Metric | Target | Measurement |
|---|---|---|
| Accessibility | WCAG 2.1 AA compliant | Automated + manual audit |
| Browser support | Latest 2 versions of Chrome, Firefox, Safari, Edge | Cross-browser testing |
| Mobile responsiveness | iOS 15+, Android 12+ | Device testing |
| Error message clarity | User-friendly, no technical jargon | UX review |

### Maintainability
| Metric | Target | Measurement |
|---|---|---|
| Code coverage | > 80% (90%+ critical paths) | SonarQube |
| Technical debt ratio | < 10% | SonarQube |
| Deployment frequency | On-demand capability | CI/CD metrics |
| Mean time to deploy | < 30 minutes | Pipeline metrics |

### Disaster Recovery & Business Continuity (ref: PM-BP-010)
| Metric | Target | Measurement |
|---|---|---|
| RTO (Recovery Time Objective) | Tier 1: < 1 hour, Tier 2: < 4 hours, Tier 3: < 24 hours | DR failover testing |
| RPO (Recovery Point Objective) | Tier 1: < 15 minutes, Tier 2: < 1 hour, Tier 3: < 4 hours | Backup verification |
| Failover time | Active-active: < 30 seconds, Active-passive: < 5 minutes | Failover drill |
| Backup frequency | Tier 1: continuous replication, Tier 2: hourly, Tier 3: daily | Backup monitoring |
| DR test frequency | Quarterly for Tier 1, Semi-annual for Tier 2/3 | DR test calendar |
| Geographic redundancy | Tier 1: multi-region active-active, Tier 2: multi-region active-passive | Architecture review |
| Data integrity on recovery | Zero data loss for committed transactions | Recovery validation test |
| Degraded mode availability | Core functions available within 5 minutes of primary failure | Failover testing |
| BCP activation time | < 30 minutes from incident declaration | BCP drill |

### Audit Trail & Logging (ref: PM-BP-009)
| Metric | Target | Measurement |
|---|---|---|
| Log ingestion latency | < 5 seconds from event to central log store | Log pipeline monitoring |
| Log availability | 99.99% — logs must be available for query | Log infrastructure monitoring |
| Log query response time | < 10 seconds for 24-hour window queries | Log platform SLA |
| Log retention — financial transactions | 7 years | Retention policy audit |
| Log retention — access/security events | 3 years minimum | Retention policy audit |
| Log completeness | 100% of defined audit events captured | Log completeness audit |
| Log integrity | Tamper-proof with cryptographic verification | Integrity verification test |
| Log storage capacity | Support retention period at projected growth rate | Capacity planning |

### Data Classification Driven Targets (ref: PM-BP-008)
| NFR Aspect | Restricted | Confidential | Internal | Public |
|---|---|---|---|---|
| Encryption at rest | AES-256 mandatory | AES-256 mandatory | Recommended | Not required |
| Encryption in transit | TLS 1.3 mandatory | TLS 1.2+ mandatory | TLS 1.2+ mandatory | TLS 1.2+ recommended |
| Audit logging | Full (all access) | Full (all modifications) | Standard (auth events) | Minimal |
| Backup frequency | Continuous replication | Hourly | Daily | Weekly |
| RTO | < 1 hour | < 4 hours | < 24 hours | Best effort |
| RPO | < 15 minutes | < 1 hour | < 4 hours | < 24 hours |
| Access control | MFA + role + resource-level | Role + resource-level | Role-based | Open |
| Data masking in non-prod | Mandatory (full anonymization) | Mandatory (tokenization) | Recommended | Not required |

---

## Example — Non-Functional Requirement

| Field | Value |
|---|---|
| **Requirement ID** | NFR-PERF-003 |
| **Title** | Fund Transfer API Response Time |
| **Category** | Performance |
| **Priority** | Critical |
| **Related Functional Req(s)** | FR-PAY-012 |
| **Description** | The fund transfer API must respond within acceptable latency under normal and peak load conditions |
| **Metric** | API response time (p95) |
| **Target Value** | < 500ms under 500 concurrent users |
| **Measurement Method** | Gatling load test simulating peak hour traffic |
| **Acceptance Criteria** | p95 response time ≤ 500ms and zero errors at 500 concurrent users sustained for 30 minutes |
| **Environment Applicability** | Production, Staging |
| **Assumptions** | Network latency to core banking < 50ms; database connection pool sized at 100 |

---

## Usage Guidelines

1. **Every requirement must have a unique ID** following the NFR-[Category]-[Seq] naming convention
2. **Metrics must be quantifiable** — avoid vague terms like "fast" or "responsive"
3. **Target values must be specific** — include load conditions and percentile thresholds
4. **Cross-reference** to related Functional (FR) requirements
5. **Review cadence**: Per release
6. **Approval**: Product Owner + Architect sign-off required
