# High-Level Design (HLD) Document Template
### Banking Domain — Agentic Knowledge Base

---

## Document Metadata

| Field | Description |
|---|---|
| **Document ID** | HLD-[Domain]-[Feature]-[Version] (e.g., HLD-PAY-DomesticTransfer-v1.0) |
| **Feature / Epic** | Epic ID and title |
| **Author** | Solution/Technical Architect |
| **Reviewers** | Security Architect, Compliance Officer, Domain Expert, Tech Lead |
| **Status** | Draft / In Review / Approved / Superseded |
| **Version** | Semantic version (major.minor) |
| **Created** | Date |
| **Last Updated** | Date |
| **Approved By** | Names, roles, dates |
| **Related Documents** | Requirements (FR, NFR, CR, SR IDs), LLD reference, ADRs |

---

## 1. Executive Summary

| Section | Content |
|---|---|
| **Purpose** | What this design achieves — business capability delivered |
| **Scope** | What is covered and explicitly excluded |
| **Business Context** | Business drivers, strategic alignment, value proposition |
| **Key Decisions** | Top 3-5 architectural decisions with rationale |
| **Constraints** | Regulatory, technical, budgetary, timeline constraints |
| **Assumptions** | Assumptions that influence the design |

---

## 2. Requirements Traceability

| Requirement ID | Type | Description | Design Section |
|---|---|---|---|
| FR-PAY-012 | Functional | Initiate domestic fund transfer | §4 Component Design |
| NFR-PERF-003 | Non-Functional | Transfer API p95 < 500ms | §7 Performance |
| CR-PSD2-001 | Compliance | Strong Customer Authentication | §6 Security |
| SR-AUTH-003 | Security | MFA for fund transfers | §6 Security |
| PM-BP-008 | Data Classification | Restricted (account numbers, amounts) | §5 Data Architecture |
| PM-BP-009 | Audit Trail | Full transaction audit logging | §6 Security |
| PM-BP-010 | DR/BCP | Tier 1 — RTO < 1 hour | §8 Availability & DR |

---

## 3. Architecture Overview

### 3.1 Context Diagram (C4 Level 1)

Shows the system in its environment — external actors, systems, and boundaries.

**Required Elements:**
- Users/personas interacting with the system
- The system being designed (as a black box)
- External systems integrated with (core banking, fraud engine, payment network, notification)
- Data flows between actors and systems (labeled with protocol and data classification)
- Trust boundaries (internal network, DMZ, external)
- Regulatory boundaries (PCI-DSS CDE scope if applicable)

**Template:**
```
[Retail Customer] --HTTPS/TLS 1.3--> [API Gateway] --mTLS--> [Transfer Service]
                                                              |
[Mobile App] --HTTPS/Cert Pin--> [API Gateway]                |--mTLS--> [Account Service]
                                                              |--mTLS--> [Fraud Service]
                                                              |--mTLS--> [Core Banking] (external)
                                                              |--Kafka--> [Notification Service]
                                                              |--Kafka--> [Audit Service]
```

### 3.2 Container Diagram (C4 Level 2)

Shows the high-level technology choices — applications, data stores, message brokers.

**Required Elements:**
- Application containers (services, web apps, mobile apps)
- Data stores (databases, caches, search indexes) with technology choice
- Message brokers (Kafka, RabbitMQ) with topic/queue names
- API gateway and load balancers
- Technology stack per container (e.g., Spring Boot 3.x, React 18, PostgreSQL 16)
- Communication protocols between containers
- Data classification per data store

### 3.3 Deployment Diagram

Shows how containers map to infrastructure.

**Required Elements:**
- Cloud provider and region(s)
- Availability zones
- Kubernetes clusters / compute instances
- Database instances (primary, replicas, DR)
- Network topology (VPC, subnets, security groups)
- Load balancers and CDN
- DR site topology
- Monitoring and logging infrastructure

---

## 4. Component Design

### 4.1 Bounded Context Map

| Context | Service(s) | Relationship | Integration Pattern |
|---|---|---|---|
| Payment | transfer-service | Partnership with Account | Sync (REST) for balance check; Async (Kafka) for debit/credit |
| Account | account-service | Open Host Service | REST API for queries; Kafka consumer for balance updates |
| Risk | fraud-service | Customer-Supplier | Sync (gRPC) for real-time scoring |
| Notification | notification-service | Conformist | Async (Kafka) consumer |
| Audit | audit-service | Conformist | Async (Kafka) consumer |

### 4.2 Component Diagram (C4 Level 3)

For each service in scope, show internal components:

**Required Elements per Service:**
- API layer (controllers, DTOs, validation)
- Application layer (use cases, orchestration)
- Domain layer (entities, value objects, domain services, events)
- Infrastructure layer (repositories, external clients, message publishers)
- Key design patterns used (hexagonal, CQRS, event sourcing, saga)

### 4.3 API Design Summary

| Endpoint | Method | Purpose | Auth | Scope | Rate Limit |
|---|---|---|---|---|---|
| /v1/transfers | POST | Initiate transfer | Bearer + MFA | transfers:write | 10/min |
| /v1/transfers/{id} | GET | Get transfer status | Bearer | transfers:read | 100/min |
| /v1/transfers/{id}/cancel | POST | Cancel pending transfer | Bearer + MFA | transfers:write | 10/min |

### 4.4 Domain Event Catalog

| Event | Producer | Consumer(s) | Schema Version | Topic |
|---|---|---|---|---|
| TransferInitiated | transfer-service | fraud-service, audit-service | 1.0 | payment.transfer.initiated |
| TransferCompleted | transfer-service | account-service, notification-service, audit-service | 1.0 | payment.transfer.completed |
| TransferFailed | transfer-service | notification-service, audit-service | 1.0 | payment.transfer.failed |

---

## 5. Data Architecture

### 5.1 Data Classification

| Data Element | Classification | Encryption | Masking (non-prod) | Retention |
|---|---|---|---|---|
| Account number | Restricted | AES-256-GCM + tokenization | Format-preserving random | Duration of account + 7 years |
| Transaction amount | Confidential | TDE | Proportional randomization | 7 years |
| Customer name | Confidential | AES-256-GCM | Fake name generation | Duration of relationship + 30 days |
| Correlation ID | Internal | TDE | No masking needed | 7 years |

### 5.2 Data Model (Conceptual)

Entity-relationship diagram showing key entities, relationships, and cardinality.

### 5.3 Data Flow Diagram

Show how data moves through the system with classification labels at each stage:

**Required Elements:**
- Data origin (user input, external system, internal event)
- Processing steps (validation, enrichment, transformation)
- Storage points (database, cache, queue)
- Data classification at each point
- Encryption state at each point (encrypted in transit, at rest)
- Cross-boundary data flows (PCI-DSS CDE, cross-border)

### 5.4 Data Residency & Cross-Border

| Data Type | Storage Location | Cross-Border Transfer? | Mechanism |
|---|---|---|---|
| Customer PII | EU region | No | N/A |
| Transaction data | EU region | Yes (to clearing network) | SCCs + TIA |

---

## 6. Security Architecture

### 6.1 Authentication & Authorization

| Component | Standard | Implementation |
|---|---|---|
| Customer auth | OAuth 2.0 + PKCE | Auth service → JWT (RS256, 15-min) |
| Step-up MFA | TOTP / Push notification | Required for transfers, beneficiary changes |
| Service-to-service | mTLS | Istio service mesh |
| API authorization | OAuth scopes + resource-level | Spring Security |

### 6.2 Security Controls Matrix

| Threat (STRIDE) | Control | Implementation | Testing |
|---|---|---|---|
| Spoofing | MFA, session management | TOTP, 15-min session, token rotation | DAST: auth bypass |
| Tampering | Input validation, CSRF, integrity checks | Bean Validation, CSRF tokens, HMAC | SAST + DAST |
| Repudiation | Audit logging, tamper-proof logs | Kafka → immutable audit store | Audit log verification |
| Information Disclosure | Encryption, masking, access control | AES-256, RLS, output encoding | Pen test |
| Denial of Service | Rate limiting, circuit breakers | API gateway rate limits, Resilience4j | Load test |
| Elevation of Privilege | RBAC + resource-level auth, least privilege | Spring Security, RLS | IDOR testing |

### 6.3 Compliance Controls

| Regulation | Requirement | Design Control | Evidence |
|---|---|---|---|
| PSD2 Art. 97 | SCA for payments | Step-up MFA before transfer | MFA event logs |
| GDPR Art. 6 | Lawful processing | Legal basis documented per data field | Processing records |
| AML Directive | Transaction monitoring | Fraud service real-time scoring | Alert logs, SAR records |
| PCI-DSS Req 3 | Protect stored card data | Tokenization + AES-256 | Encryption verification |

### 6.4 Audit Trail Design

| Event | Trigger | Fields Logged | Retention |
|---|---|---|---|
| Transfer initiated | POST /v1/transfers | user_id, account (masked), amount, beneficiary (masked), timestamp, correlation_id | 7 years |
| MFA verified | MFA completion | user_id, method, outcome, device, IP, timestamp | 3 years |
| Fraud check | Fraud service response | correlation_id, fraud_score, decision, timestamp | 7 years |

---

## 7. Performance & Scalability

### 7.1 Performance Targets

| Metric | Target | Measurement |
|---|---|---|
| API response (p95) | < 500ms | Gatling load test |
| API response (p99) | < 1000ms | APM monitoring |
| Throughput | 500 TPS | Load test |
| Database query (p95) | < 50ms | Query profiling |

### 7.2 Scalability Design

| Component | Scaling Strategy | Trigger |
|---|---|---|
| Transfer service | Horizontal (Kubernetes HPA) | CPU > 70% or request queue > 100 |
| Database | Read replicas for queries | Read latency > 100ms |
| Kafka | Partition scaling | Consumer lag > 1000 |

### 7.3 Caching Strategy

| Data | Cache | TTL | Invalidation |
|---|---|---|---|
| Account balance | No cache (real-time) | — | — |
| Exchange rates | Redis | 60s | Time-based |
| Customer profile | Redis | 5 min | Event-based (profile.updated) |

---

## 8. Availability & Disaster Recovery

### 8.1 Availability Design

| Component | Availability Target | HA Strategy |
|---|---|---|
| Transfer service | 99.95% | Multi-AZ, 3+ replicas, health checks |
| Database | 99.99% | Synchronous replication, auto-failover |
| Kafka | 99.99% | 3-broker cluster, replication factor 3 |

### 8.2 DR Design

| Metric | Target | Implementation |
|---|---|---|
| RTO | < 1 hour | Multi-region active-passive; automated failover |
| RPO | < 15 minutes | Synchronous replication to DR |
| DR test frequency | Quarterly | Automated failover drill |

### 8.3 Degraded Mode Behavior

| Dependency Failure | Behavior | User Impact |
|---|---|---|
| Core banking unavailable | Queue transfer; return 202 Accepted | "Transfer queued — processing within 30 minutes" |
| Fraud service unavailable | Hold transfer for manual review | "Transfer under review — you'll be notified" |
| Notification service unavailable | Complete transfer; queue notification | No impact on transfer; delayed notification |

---

## 9. Monitoring & Observability

| Aspect | Tool | Key Metrics/Alerts |
|---|---|---|
| APM | Datadog / Dynatrace | Response time, error rate, throughput |
| Logging | ELK / Splunk | Structured JSON with correlation ID |
| Tracing | Jaeger via OpenTelemetry | End-to-end transfer flow trace |
| Alerting | PagerDuty | p95 > 500ms, error rate > 1%, circuit breaker open |
| Business metrics | Custom dashboard | Transfer volume, success rate, average amount |

---

## 10. Risks & Mitigations

| Risk | Impact | Likelihood | Mitigation |
|---|---|---|---|
| Core banking API latency | Transfer timeout | Medium | Circuit breaker + async fallback |
| Fraud service false positives | Legitimate transfers held | Medium | Tunable thresholds; manual review SLA |
| Data breach | Regulatory penalty, reputation | Low | Encryption, RLS, audit, monitoring |

---

## 11. Architecture Decision Records (ADRs)

| ADR ID | Decision | Status |
|---|---|---|
| ADR-001 | Use Kafka for async events (not RabbitMQ) | Accepted |
| ADR-002 | Use orchestration saga (not choreography) for transfer flow | Accepted |
| ADR-003 | Use PostgreSQL with CQRS (not event sourcing) for transfers | Accepted |

---

## 12. Diagrams Checklist

| Diagram | Required | Tool | Audience |
|---|---|---|---|
| C4 Context (Level 1) | Always | Structurizr, draw.io, Mermaid | All stakeholders |
| C4 Container (Level 2) | Always | Structurizr, draw.io, Mermaid | Architects, tech leads |
| C4 Component (Level 3) | Always | Structurizr, draw.io, Mermaid | Developers, tech leads |
| Deployment diagram | Always | draw.io, Mermaid | DevOps, architects |
| Data flow diagram | Always | draw.io | Security, compliance, architects |
| Sequence diagram (key flows) | Always | Mermaid, PlantUML | Developers |
| ER diagram (conceptual) | Always | dbdiagram.io, draw.io | Developers, DBAs |
| Network topology | If infrastructure changes | draw.io | DevOps, security |
| Threat model diagram | If Confidential/Restricted data | draw.io, OWASP Threat Dragon | Security team |

---

## 13. High-Level Milestones & Project Plan

### 13.1 Milestone Summary

| # | Milestone | Description | Target Date | Owner | Exit Criteria |
|---|---|---|---|---|---|
| M1 | Design Approved | HLD + LLD reviewed and signed off | Week X | Solution Architect | All approvals obtained; no open Reject items |
| M2 | Development Complete | All features coded, unit tested, code reviewed | Week X+N | Tech Lead | Code complete; unit tests ≥ 80%; SAST passed |
| M3 | Integration Complete | All service integrations working in staging | Week X+N | Tech Lead | Integration tests passing; all dependencies connected |
| M4 | Security Validated | SAST, DAST, SCA, security code review complete | Week X+N | Security Engineer | Zero critical/high findings |
| M5 | Performance Validated | Load/stress tests passed against NFR targets | Week X+N | QA Lead | p95 < target; throughput met; no degradation |
| M6 | Compliance Signed Off | All compliance checkpoints passed; evidence collected | Week X+N | Compliance Officer | Compliance sign-off obtained; evidence package complete |
| M7 | UAT Complete | Business acceptance criteria verified | Week X+N | Product Owner | All acceptance criteria met; PO sign-off |
| M8 | Production Deployment | Feature deployed to production | Week X+N | Release Manager | Deployment successful; monitoring active; rollback tested |
| M9 | Post-Deployment Verification | Production validation complete | Week X+N+1 | Tech Lead + Compliance | Production health verified; compliance controls confirmed |

### 13.2 Phase Plan

| Phase | Duration | Activities | Deliverables | Dependencies |
|---|---|---|---|---|
| **Design** | 1-2 weeks | HLD creation, HLD review, LLD creation, LLD review, ADRs | Approved HLD, Approved LLD, ADRs | Requirements approved |
| **Sprint 1: Core Development** | 2 weeks | Backend implementation, database schema + migrations, domain model, API endpoints, unit tests | Working API (happy path), DB deployed to dev | Approved LLD |
| **Sprint 2: Integration & Security** | 2 weeks | External integrations (core banking, fraud), MFA flow, audit logging, security controls, integration tests | Integrated service in staging, audit trail working | Sprint 1 complete; external APIs available |
| **Sprint 3: Quality & Compliance** | 2 weeks | SAST/DAST/SCA remediation, security code review, performance testing, compliance testing, UAT support | Security scan reports, performance results, compliance evidence | Sprint 2 complete; staging environment ready |
| **Deployment** | 1 week | Rollback plan + testing, monitoring setup, production deployment, post-deployment verification | Production deployment, monitoring dashboards, rollback verified | All M1-M7 milestones met |

### 13.3 Dependency Timeline

```
Week 1-2: Design
  ├── HLD creation & review (M1 gate)
  └── LLD creation & review (M1 gate)
         │
Week 3-4: Sprint 1 — Core Development
  ├── DB schema + migrations
  ├── Domain model + API endpoints
  ├── Unit tests
  └── Compliance review submitted (parallel)
         │
Week 5-6: Sprint 2 — Integration & Security
  ├── External integrations ←── [External: Core Banking API ready]
  ├── MFA flow ←── [External: Auth service MFA endpoint ready]
  ├── Audit logging implementation
  ├── Integration tests
  └── SAST/SCA scans (continuous)
         │
Week 7-8: Sprint 3 — Quality & Compliance
  ├── DAST scan + remediation ←── [Depends: Deployed to staging]
  ├── Security code review
  ├── Performance testing ←── [Depends: Staging with prod-like data volume]
  ├── Compliance testing + evidence collection
  └── UAT ←── [External: Business users available]
         │
Week 9: Deployment
  ├── Rollback plan + testing ←── [Depends: Staging deployment]
  ├── Production deployment ←── [Gate: M4, M5, M6, M7 all met]
  └── Post-deployment verification
```

### 13.4 Resource Plan

| Role | Phase | Allocation | Responsibilities |
|---|---|---|---|
| Solution Architect | Design | 100% | HLD creation, review facilitation |
| Tech Lead | Design + All Sprints | 80% | LLD creation, development oversight, code review |
| Backend Developer(s) | Sprint 1-3 | 100% | Implementation, unit tests |
| Frontend Developer(s) | Sprint 1-3 | 100% | UI implementation (if applicable) |
| QA Engineer | Sprint 2-3 | 100% | Integration tests, security tests, compliance tests, performance tests |
| Security Engineer | Sprint 2-3 + Deployment | 50% | Security review, DAST, pen test support |
| DBA | Design + Sprint 1 | 25% | Schema review, migration review, index optimization |
| Compliance Officer | Design + Sprint 3 | 25% | Compliance review, sign-off, evidence review |
| DevOps Engineer | Sprint 2 + Deployment | 50% | CI/CD, monitoring, deployment, rollback |
| Product Owner | Sprint 3 (UAT) | 25% | UAT, acceptance criteria verification |

### 13.5 Risk-Adjusted Timeline

| Risk | Impact on Timeline | Mitigation | Contingency |
|---|---|---|---|
| External API not ready on time | Sprint 2 delayed 1-2 weeks | Early engagement; mock services for parallel development | Use WireMock stubs; defer integration to Sprint 3 |
| Security findings require major rework | Sprint 3 extended 1 week | SAST from Sprint 1; address findings continuously | Prioritize critical/high; defer medium to follow-up |
| Performance targets not met | Sprint 3 extended 1 week | Performance test early in Sprint 2 with subset | Optimize queries/caching; scale infrastructure |
| Compliance review delayed | Deployment delayed | Submit compliance review in Sprint 1 (parallel) | Escalate to Compliance Officer manager |
| Business users unavailable for UAT | Deployment delayed 1 week | Schedule UAT slots in advance | Conduct UAT with proxy PO; formal UAT in post-deployment |

### 13.6 Go/No-Go Criteria

Production deployment proceeds only when ALL of the following are met:

| # | Criterion | Verified By |
|---|---|---|
| 1 | All acceptance criteria (business + regulatory) met | Product Owner + Compliance Officer |
| 2 | Security scans passed — zero critical/high findings | Security Engineer |
| 3 | Performance targets met (p95, throughput, error rate) | QA Lead |
| 4 | Compliance sign-off obtained | Compliance Officer |
| 5 | Rollback plan created and tested in staging | DevOps + Tech Lead |
| 6 | Monitoring and alerting configured and verified | DevOps |
| 7 | All open defects triaged — no P1/P2 open | Tech Lead |
| 8 | Documentation updated (API docs, runbooks, audit trail docs) | Tech Lead |
| 9 | On-call team briefed on new feature | DevOps + Tech Lead |

---

## Approval Sign-Off

| Role | Name | Date | Decision |
|---|---|---|---|
| Solution Architect | | | Approved / Rejected |
| Security Architect | | | Approved / Rejected |
| Compliance Officer | | | Approved / Rejected |
| Domain Expert | | | Approved / Rejected |
| Tech Lead | | | Approved / Rejected |
