# Best Practices & Standards for Microservices and Domain-Driven Design

---

## 1. Domain-Driven Design Foundations

### Strategic Design

#### Bounded Contexts for Banking

A bounded context is a boundary within which a domain model is consistent and a ubiquitous language applies.

```
Banking Domain
├── Customer Context
│   ├── Customer Profile (name, address, contact)
│   ├── KYC/AML (identity verification, risk scoring)
│   └── Consent Management (GDPR consents, preferences)
│
├── Account Context
│   ├── Account Lifecycle (open, close, freeze, dormant)
│   ├── Balance Management (available, ledger, hold)
│   └── Statement Generation
│
├── Payment Context
│   ├── Transfer (domestic, international, scheduled)
│   ├── Bill Payment
│   ├── Standing Orders / Direct Debits
│   └── Payment Processing (clearing, settlement)
│
├── Card Context [PCI-DSS Boundary]
│   ├── Card Lifecycle (issue, activate, block, replace)
│   ├── Card Limits
│   ├── PIN Management
│   └── Card Transactions
│
├── Lending Context
│   ├── Loan Origination (application, decisioning)
│   ├── Loan Servicing (repayment, restructuring)
│   └── Collections
│
├── Risk & Compliance Context
│   ├── Fraud Detection
│   ├── AML Transaction Monitoring
│   ├── Regulatory Reporting
│   └── Sanctions Screening
│
├── Notification Context
│   ├── Email, SMS, Push
│   ├── In-App Notifications
│   └── Regulatory Notifications
│
└── Shared Kernel
    ├── Money (amount + currency value object)
    ├── Date/Time (business day calendar)
    ├── Country/Currency reference data
    └── Audit Event schema
```

#### Context Mapping

| Relationship | Description | Banking Example |
|---|---|---|
| **Partnership** | Two contexts cooperate; changes coordinated | Account ↔ Payment (balance check during transfer) |
| **Customer-Supplier** | Upstream supplies, downstream consumes | Customer (supplier) → Account (customer) |
| **Conformist** | Downstream conforms to upstream model | Internal services conform to core banking model |
| **Anti-Corruption Layer (ACL)** | Translate between models | ACL between modern services and legacy core banking |
| **Open Host Service** | Published API for multiple consumers | Account service exposes standard API for payments, cards, lending |
| **Published Language** | Shared schema/protocol | Domain events published in standard Avro/Protobuf schema |
| **Separate Ways** | No integration; independent | Marketing analytics separate from core banking |

### Tactical Design

#### Building Blocks

| Building Block | Definition | Banking Example |
|---|---|---|
| **Entity** | Object with identity that persists over time | Account (accountId), Customer (customerId), Transfer (transferId) |
| **Value Object** | Immutable object defined by attributes, no identity | Money (amount + currency), Address, IBAN, DateRange |
| **Aggregate** | Cluster of entities/VOs with a root entity; consistency boundary | Transfer Aggregate (Transfer root + TransferLines + FraudCheck) |
| **Aggregate Root** | Entry point to the aggregate; enforces invariants | Account is root of Account Aggregate (Account + Balance + Holds) |
| **Domain Event** | Something that happened in the domain | TransferInitiated, TransferCompleted, AccountFrozen, KYCCompleted |
| **Domain Service** | Stateless operation that doesn't belong to an entity | FraudScreeningService, ExchangeRateService, LimitValidationService |
| **Repository** | Abstraction for aggregate persistence | AccountRepository, TransferRepository |
| **Factory** | Complex object creation | TransferFactory (creates Transfer with validation, fraud check, audit) |

#### Aggregate Design Rules

| Rule | Standard | Banking Rationale |
|---|---|---|
| Small aggregates | Prefer small; one entity + value objects | Performance; concurrent access to accounts |
| Reference by ID | Aggregates reference other aggregates by ID, not object | Account references CustomerId, not Customer object |
| Eventual consistency between aggregates | Use domain events for cross-aggregate consistency | Transfer debit and credit are separate aggregates; eventual consistency via events |
| Transactional consistency within aggregate | Single transaction per aggregate | Balance update + hold creation in single transaction |
| Invariant enforcement | Aggregate root enforces all business rules | Account enforces: balance ≥ 0, daily limit, account status |

---

## 2. Microservice Design Standards

### Service Boundaries

| Principle | Standard | Banking Example |
|---|---|---|
| One bounded context = one service (or small set) | Service boundary aligns with DDD bounded context | Payment Service owns the entire Payment Context |
| Single responsibility | Each service owns one business capability | Transfer Service handles transfers; doesn't manage beneficiaries |
| Data ownership | Each service owns its data store exclusively | Account Service owns account database; no other service reads it directly |
| Autonomous deployment | Each service deployable independently | Payment Service deploys without Account Service redeployment |
| Technology agnostic | Services can use different tech stacks | Java/Spring Boot for core services; Python for ML-based fraud detection |

### Service Decomposition for Banking

| Service | Bounded Context | Data Store | Key APIs |
|---|---|---|---|
| customer-service | Customer | PostgreSQL | CRUD customers, KYC status |
| account-service | Account | PostgreSQL | Balances, transactions, statements |
| transfer-service | Payment (transfers) | PostgreSQL | Initiate, status, history |
| payment-service | Payment (processing) | PostgreSQL | Clearing, settlement |
| beneficiary-service | Payment (beneficiaries) | PostgreSQL | CRUD beneficiaries |
| card-service | Card [PCI-DSS] | PostgreSQL (encrypted) | Card lifecycle, limits, PIN |
| lending-service | Lending | PostgreSQL | Applications, servicing |
| fraud-service | Risk & Compliance | PostgreSQL + Redis | Real-time fraud scoring |
| aml-service | Risk & Compliance | PostgreSQL + Elasticsearch | Transaction monitoring, SAR |
| notification-service | Notification | PostgreSQL + message queue | Email, SMS, push delivery |
| audit-service | Cross-cutting | Elasticsearch / immutable store | Audit event ingestion, query |
| auth-service | Cross-cutting | PostgreSQL + Redis | OAuth 2.0, token management |

### Database per Service

| Standard | Requirement |
|---|---|
| Exclusive ownership | Each service owns its database schema; no shared databases |
| No cross-service queries | Services query their own data only; use APIs or events for cross-service data |
| Schema migration | Each service manages its own migrations (Flyway/Liquibase) |
| Data duplication | Acceptable — services maintain local copies of data they need (via events) |
| Referential integrity | Enforced within service; eventual consistency across services |

### Service Communication Standards

| Pattern | Protocol | When to Use | Banking Example |
|---|---|---|---|
| Synchronous REST | HTTP/JSON | Query, real-time response needed | GET account balance |
| Synchronous gRPC | HTTP/2 + Protobuf | High-performance internal calls | Fraud scoring during transfer (low latency) |
| Asynchronous events | Kafka/RabbitMQ | State changes, eventual consistency | TransferCompleted → Account updates balance |
| Asynchronous commands | Message queue | Deferred processing | Schedule transfer for future date |
| Request-reply (async) | Kafka with reply topic | Async with response needed | Batch payment processing with status callback |

---

## 3. Service Template (Chassis)

Every microservice should start from a standard template (service chassis) that includes:

| Component | Implementation | Purpose |
|---|---|---|
| Health check | /actuator/health (Spring Boot) | Liveness and readiness probes |
| Metrics | Micrometer → Prometheus | Performance monitoring |
| Distributed tracing | OpenTelemetry → Jaeger | Request tracing across services |
| Structured logging | JSON format with correlation ID | Centralized log aggregation |
| Configuration | Spring Cloud Config / Vault | Externalized, environment-specific config |
| Security | Spring Security + OAuth 2.0 resource server | Token validation, scope enforcement |
| API documentation | SpringDoc OpenAPI | Auto-generated OpenAPI spec |
| Database migration | Flyway | Version-controlled schema changes |
| Circuit breaker | Resilience4j | Fault tolerance for downstream calls |
| Audit logging | Custom audit interceptor | Compliance audit trail |
| Error handling | Global exception handler | Consistent error responses |
| Input validation | Bean Validation (JSR-380) | Request validation |
| Correlation ID | MDC filter + header propagation | Traceability |

### Project Structure (Hexagonal Architecture)

```
transfer-service/
├── src/main/java/com/bank/transfer/
│   ├── domain/                    # Domain layer (no framework dependencies)
│   │   ├── model/                 # Entities, Value Objects, Aggregates
│   │   │   ├── Transfer.java
│   │   │   ├── Money.java
│   │   │   └── TransferStatus.java
│   │   ├── event/                 # Domain Events
│   │   │   ├── TransferInitiated.java
│   │   │   └── TransferCompleted.java
│   │   ├── service/               # Domain Services
│   │   │   └── TransferDomainService.java
│   │   ├── repository/            # Repository interfaces (ports)
│   │   │   └── TransferRepository.java
│   │   └── exception/             # Domain exceptions
│   │       └── InsufficientFundsException.java
│   │
│   ├── application/               # Application layer (use cases)
│   │   ├── port/
│   │   │   ├── in/                # Inbound ports (use case interfaces)
│   │   │   │   └── InitiateTransferUseCase.java
│   │   │   └── out/               # Outbound ports (external dependencies)
│   │   │       ├── AccountPort.java
│   │   │       ├── FraudPort.java
│   │   │       └── EventPublisherPort.java
│   │   └── service/               # Use case implementations
│   │       └── TransferApplicationService.java
│   │
│   └── infrastructure/            # Infrastructure layer (adapters)
│       ├── web/                   # Inbound adapter: REST API
│       │   ├── TransferController.java
│       │   └── dto/
│       ├── persistence/           # Outbound adapter: Database
│       │   ├── TransferJpaRepository.java
│       │   └── entity/
│       ├── messaging/             # Outbound adapter: Events
│       │   └── KafkaEventPublisher.java
│       ├── client/                # Outbound adapter: External services
│       │   ├── AccountServiceClient.java
│       │   └── FraudServiceClient.java
│       └── config/                # Framework configuration
│           ├── SecurityConfig.java
│           └── KafkaConfig.java
```

---

## 4. Data Management Patterns

### CQRS (Command Query Responsibility Segregation)

| Aspect | Command Side | Query Side |
|---|---|---|
| Purpose | Write operations (create, update, delete) | Read operations (queries, reports) |
| Data store | Normalized (PostgreSQL) | Denormalized (Elasticsearch, read replicas) |
| Consistency | Strong (transactional) | Eventual (updated via events) |
| Scaling | Scale for write throughput | Scale for read throughput independently |
| Banking example | Initiate transfer, update balance | Transaction history, account dashboard, statements |

### Event Sourcing (When Appropriate)

| Use When | Avoid When |
|---|---|
| Complete audit trail is mandatory (financial transactions) | Simple CRUD with no audit requirements |
| Need to reconstruct state at any point in time | High-volume, low-value data |
| Regulatory requirement to prove what happened and when | Team unfamiliar with event sourcing patterns |
| Complex business processes with multiple state transitions | Simple domain with few state changes |

**Banking Use Cases for Event Sourcing:**
- Transfer lifecycle (initiated → fraud-checked → approved → executed → settled)
- Loan application workflow (submitted → documents-verified → credit-checked → approved/rejected)
- Account lifecycle (opened → active → dormant → closed)

### Saga Pattern (Distributed Transactions)

| Type | How It Works | When to Use |
|---|---|---|
| **Choreography** | Each service listens for events and reacts | Simple flows with few steps; loosely coupled |
| **Orchestration** | Central orchestrator coordinates steps | Complex flows; need visibility and control |

**Banking Transfer Saga (Orchestration):**

```
TransferOrchestrator:
  1. Validate transfer request
  2. Call fraud-service → FraudCheckPassed / FraudCheckFailed
  3. Call account-service → DebitAccount / DebitFailed (insufficient funds)
  4. Call account-service → CreditBeneficiary / CreditFailed
  5. Call notification-service → NotifyCustomer
  6. Publish TransferCompleted event

Compensation (on failure at any step):
  - Step 4 fails → Reverse debit (step 3)
  - Step 3 fails → Cancel transfer, notify customer
  - Step 2 fails → Reject transfer, notify customer
```

**Saga Standards:**
- Every step must have a compensating action defined
- Saga state must be persisted (survive service restart)
- Saga must have a timeout — don't wait indefinitely
- Saga must be idempotent — retries must not cause duplicates
- Saga must produce audit events at every step

---

## 5. Cross-Cutting Concerns

### Service Mesh

| Concern | Implementation | Standard |
|---|---|---|
| mTLS | Istio / Linkerd | All service-to-service communication encrypted |
| Traffic management | Istio VirtualService | Canary deployments, traffic splitting |
| Observability | Istio telemetry | Automatic metrics, traces, logs |
| Circuit breaking | Istio DestinationRule | Outlier detection, connection limits |
| Rate limiting | Istio EnvoyFilter | Per-service rate limits |
| Access control | Istio AuthorizationPolicy | Service-to-service authorization |

### Observability Standards

| Pillar | Tool | Standard |
|---|---|---|
| Metrics | Prometheus + Grafana | RED metrics (Rate, Errors, Duration) per service |
| Logging | ELK / Splunk | Structured JSON; correlation ID in every entry |
| Tracing | Jaeger / Zipkin via OpenTelemetry | Trace every request across all services |
| Alerting | PagerDuty / OpsGenie | Alert on SLO breach, error rate spike, latency spike |

### Configuration Management

| Standard | Implementation |
|---|---|
| Externalized config | Spring Cloud Config Server or Kubernetes ConfigMaps |
| Secrets management | HashiCorp Vault or AWS Secrets Manager — never in code/config files |
| Environment-specific | Config per environment (dev, staging, production) |
| Feature flags | LaunchDarkly or Unleash — gradual rollout, instant kill switch |
| Config versioning | All config in version control; changes audited |

---

## 6. Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Distributed monolith | Services tightly coupled; must deploy together | Proper bounded context boundaries; async communication |
| Shared database | Services coupled at data layer; schema changes break others | Database per service; events for data sharing |
| Synchronous chain | A → B → C → D; latency compounds; one failure breaks all | Async events where possible; circuit breakers for sync |
| Anemic domain model | Business logic in services, not domain objects | Rich domain model; entities enforce their own invariants |
| God service | One service does everything | Decompose by bounded context |
| Chatty communication | Hundreds of calls between services per request | Aggregate data locally; batch calls; CQRS read models |
| No saga compensation | Distributed transaction left in inconsistent state | Every saga step has a compensating action |
| Shared libraries with business logic | Coupling via shared code | Shared libraries for infrastructure only; business logic in services |
