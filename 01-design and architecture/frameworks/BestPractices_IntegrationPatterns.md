# Best Practices & Standards for Integration Patterns

---

## 1. Integration Pattern Selection

### Decision Framework

| Factor | Synchronous (REST/gRPC) | Asynchronous (Events/Queues) |
|---|---|---|
| Response needed immediately | ✅ Use sync | ❌ Not suitable |
| Fire-and-forget | ❌ Wasteful | ✅ Use async |
| Temporal coupling acceptable | ✅ Both services must be up | ❌ Producer/consumer independent |
| Data consistency | Strong (immediate) | Eventual |
| Throughput priority | Moderate | High |
| Failure isolation | Low (cascade risk) | High (buffered) |
| Latency sensitivity | Low latency required | Latency tolerance acceptable |

### Banking Integration Map

| Integration | Pattern | Rationale |
|---|---|---|
| Balance inquiry during transfer | **Sync (REST)** | Real-time balance needed before debit |
| Fraud scoring during transfer | **Sync (gRPC)** | Real-time decision needed; low latency critical |
| Transfer completed → update balance | **Async (event)** | Eventual consistency acceptable; decoupled |
| Transfer completed → send notification | **Async (event)** | Fire-and-forget; notification delay acceptable |
| Transfer completed → generate CTR | **Async (event)** | Regulatory report can be generated asynchronously |
| Batch payment processing | **Async (queue)** | High volume; processed sequentially |
| Real-time FX rate lookup | **Sync (REST)** | Current rate needed for transaction |
| KYC verification | **Async (queue + callback)** | Third-party verification takes time; callback on completion |
| Account statement generation | **Async (queue)** | Resource-intensive; processed in background |
| Sanctions screening | **Sync (REST)** | Must block transaction until screening completes |
| Audit event publishing | **Async (event)** | Fire-and-forget; must not block business operation |
| Core banking integration | **Sync (REST) + circuit breaker** | Real-time needed; but must handle unavailability |

---

## 2. Synchronous Integration Standards (Request-Response)

### REST Integration Standards

| Standard | Requirement |
|---|---|
| Protocol | HTTPS only; TLS 1.2+ |
| Authentication | OAuth 2.0 Bearer token for external; mTLS for internal service-to-service |
| Timeout | Connect: 5 seconds; Read: 30 seconds (configurable per endpoint) |
| Retry | Exponential backoff with jitter; max 3 retries; only for idempotent operations or with idempotency key |
| Circuit breaker | Open after 5 consecutive failures or 50% error rate in 60-second window |
| Correlation ID | Propagate X-Correlation-ID on every call |
| Error handling | Map downstream errors to appropriate upstream response; never expose downstream details |
| Logging | Log request/response (masked) with correlation ID, latency, status |
| Rate limiting | Respect downstream rate limits; implement client-side throttling |
| Caching | Cache immutable/slow-changing data (exchange rates, product catalog); never cache balances or PII |

### gRPC Integration Standards

| Standard | Requirement |
|---|---|
| Protocol | HTTP/2 with TLS |
| Serialization | Protocol Buffers (proto3) |
| When to use | Internal service-to-service; low-latency requirements (< 50ms) |
| Timeout | Deadline propagation — set deadline at entry point, propagate to all downstream calls |
| Retry | Built-in gRPC retry policy; idempotent methods only |
| Load balancing | Client-side load balancing (gRPC supports natively) |
| Health check | gRPC Health Checking Protocol |
| Streaming | Use server streaming for large result sets (transaction history); bidirectional for real-time updates |

### Circuit Breaker Configuration

```yaml
# Resilience4j configuration
resilience4j:
  circuitbreaker:
    instances:
      core-banking:
        slidingWindowSize: 10
        failureRateThreshold: 50
        waitDurationInOpenState: 30s
        permittedNumberOfCallsInHalfOpenState: 3
        slowCallDurationThreshold: 2s
        slowCallRateThreshold: 80
      fraud-service:
        slidingWindowSize: 5
        failureRateThreshold: 50
        waitDurationInOpenState: 10s
  retry:
    instances:
      core-banking:
        maxAttempts: 3
        waitDuration: 1s
        exponentialBackoffMultiplier: 2
        retryExceptions:
          - java.net.ConnectException
          - java.net.SocketTimeoutException
  timelimiter:
    instances:
      core-banking:
        timeoutDuration: 5s
      fraud-service:
        timeoutDuration: 2s
```

### Fallback Strategies

| Downstream Failure | Fallback Strategy | Banking Example |
|---|---|---|
| Core banking unavailable | Queue for retry; return 202 Accepted | Transfer queued; customer notified of delay |
| Fraud service unavailable | Hold transaction for manual review (never auto-approve) | Transfer held; fraud team alerted |
| Notification service unavailable | Queue notification; complete transaction | Transfer succeeds; notification sent when service recovers |
| Exchange rate service unavailable | Use last known rate with staleness indicator | Show rate with "as of [timestamp]" warning |
| Statement service unavailable | Return cached statement or "temporarily unavailable" | Customer sees last cached statement |

### Service Discovery

| Approach | When to Use | Implementation |
|---|---|---|
| DNS-based | Kubernetes environments | Kubernetes Service DNS (service-name.namespace.svc.cluster.local) |
| Client-side discovery | Spring Cloud ecosystem | Eureka / Consul + Ribbon/Spring Cloud LoadBalancer |
| Server-side discovery | API gateway routing | API gateway routes to service instances |
| Service mesh | Complex microservice topology | Istio/Linkerd handles discovery transparently |

---

## 3. Asynchronous Integration Standards (Event-Driven)

### Event-Driven Architecture

#### Domain Events

| Standard | Requirement |
|---|---|
| Event naming | Past tense: TransferInitiated, AccountDebited, FraudCheckPassed |
| Event schema | Versioned; backward compatible; registered in schema registry |
| Event content | Include all data consumer needs — avoid requiring callback to producer |
| Event immutability | Events are facts; never modified after publication |
| Event ordering | Guaranteed within partition/key; not across partitions |
| Event idempotency | Consumers must handle duplicate events (at-least-once delivery) |

#### Event Schema Standard

```json
{
  "eventId": "evt-uuid-001",
  "eventType": "transfer.completed",
  "eventVersion": "1.0",
  "source": "transfer-service",
  "timestamp": "2026-03-15T15:30:00.123Z",
  "correlationId": "corr-uuid-789",
  "causationId": "evt-uuid-000",
  "data": {
    "transferId": "TRF-2026-0315-001",
    "sourceAccountId": "ACC-001234",
    "beneficiaryAccountId": "ACC-005678",
    "amount": 5000.00,
    "currency": "USD",
    "status": "COMPLETED"
  },
  "metadata": {
    "userId": "USR-001",
    "channel": "MOBILE",
    "schemaVersion": "1.0"
  }
}
```

#### Event Envelope Fields

| Field | Purpose | Required |
|---|---|---|
| eventId | Unique event identifier (UUID) — for deduplication | Yes |
| eventType | Domain event type (dot notation) | Yes |
| eventVersion | Schema version for this event type | Yes |
| source | Producing service name | Yes |
| timestamp | When the event occurred (ISO 8601 UTC) | Yes |
| correlationId | Links related events across the flow | Yes |
| causationId | The event that caused this event | Recommended |
| data | Event payload (domain-specific) | Yes |
| metadata | Context (userId, channel, schemaVersion) | Recommended |

### Apache Kafka Standards

| Configuration | Standard | Rationale |
|---|---|---|
| Replication factor | 3 (minimum) | Durability; survive broker failure |
| Min in-sync replicas | 2 | Prevent data loss on broker failure |
| Acks | all (acks=-1) | Guarantee message durability |
| Partitioning key | Business key (accountId, customerId) | Ordering guarantee per entity |
| Partition count | Start with 12; scale based on throughput | Parallelism for consumers |
| Retention | 7 days for operational topics; 30 days for audit topics | Balance storage vs. replay capability |
| Compression | lz4 or snappy | Reduce network and storage |
| Schema registry | Confluent Schema Registry with Avro or Protobuf | Schema evolution; compatibility checks |
| Consumer group | One per consuming service | Independent consumption |
| Offset management | Commit after processing (at-least-once) | Prevent message loss |

#### Topic Naming Convention

```
{domain}.{entity}.{event-type}

Examples:
payment.transfer.initiated
payment.transfer.completed
payment.transfer.failed
account.balance.updated
customer.kyc.completed
risk.fraud.alert-raised
notification.email.sent
audit.event.created
```

#### Consumer Idempotency

| Strategy | Implementation | When to Use |
|---|---|---|
| Event ID deduplication | Store processed eventIds; skip duplicates | All consumers |
| Idempotent operations | Design operations to be naturally idempotent | Balance updates (set to value, not increment) |
| Optimistic locking | Version field on entity; reject stale updates | Concurrent updates to same entity |
| Inbox pattern | Store incoming events in inbox table; process from inbox | Complex processing with transactional guarantees |

### Message Queue Standards (RabbitMQ / SQS)

| Use Case | Pattern | Banking Example |
|---|---|---|
| Task queue | Point-to-point; one consumer processes each message | Statement generation, batch payment processing |
| Work distribution | Competing consumers; load balanced | Notification delivery (email, SMS, push) |
| Delayed processing | Message with TTL / delay queue | Scheduled transfers, retry after delay |
| Priority queue | Messages with priority levels | P1 fraud alerts processed before P3 notifications |

| Configuration | Standard |
|---|---|
| Acknowledgment | Manual ack after successful processing |
| Dead letter queue (DLQ) | Every queue must have a DLQ for failed messages |
| DLQ monitoring | Alert on DLQ depth > 0; investigate within 1 hour |
| Message TTL | Set per use case; prevent infinite queue growth |
| Retry policy | 3 retries with exponential backoff before DLQ |
| Message persistence | Durable queues and persistent messages |
| Poison message handling | Detect and route to DLQ after max retries |

### Dead Letter Queue Standards

| Standard | Requirement |
|---|---|
| Every queue has a DLQ | No exceptions — failed messages must not be lost |
| DLQ monitoring | Real-time alerting when messages arrive in DLQ |
| DLQ investigation SLA | Investigate within 1 hour for financial messages; 4 hours for others |
| DLQ replay | Capability to replay DLQ messages after fix |
| DLQ retention | 30 days minimum — enough time to investigate and replay |
| DLQ audit | Log all DLQ arrivals with reason, original message metadata, and correlation ID |

---

## 4. Schema Evolution & Compatibility

### Schema Compatibility Rules

| Change | Backward Compatible | Forward Compatible | Banking Action |
|---|---|---|---|
| Add optional field | ✅ | ✅ | Safe — no version bump |
| Add required field | ❌ | ✅ | New schema version; migrate consumers first |
| Remove field | ✅ | ❌ | Deprecate first; remove after all consumers updated |
| Rename field | ❌ | ❌ | Add new field + deprecate old; never rename |
| Change field type | ❌ | ❌ | New schema version; coordinate migration |

### Schema Registry Standards

| Standard | Requirement |
|---|---|
| Registry | Confluent Schema Registry (or equivalent) |
| Format | Avro (recommended) or Protobuf |
| Compatibility mode | BACKWARD (default) — new schema can read old data |
| Validation | Schema validated on publish; incompatible schemas rejected |
| Versioning | Auto-versioned by registry; manual version in event envelope |
| Documentation | Every schema documented with field descriptions and examples |

---

## 5. Event Choreography vs. Orchestration

### Choreography

```
TransferService publishes: TransferInitiated
  → FraudService listens → publishes: FraudCheckPassed
    → AccountService listens → publishes: AccountDebited
      → AccountService listens → publishes: BeneficiaryCredited
        → NotificationService listens → sends notification
        → AuditService listens → records audit event
        → TransferService listens → updates status to COMPLETED
```

| Pros | Cons |
|---|---|
| Loose coupling | Hard to understand full flow |
| Easy to add new consumers | Difficult to handle failures/compensation |
| No single point of failure | No central visibility of saga state |

### Orchestration

```
TransferOrchestrator:
  1. Publish: ValidateTransfer command → TransferService
  2. Publish: CheckFraud command → FraudService
  3. Publish: DebitAccount command → AccountService
  4. Publish: CreditBeneficiary command → AccountService
  5. Publish: SendNotification command → NotificationService
  6. Publish: TransferCompleted event (for other consumers)

  On failure at any step: execute compensation in reverse order
```

| Pros | Cons |
|---|---|
| Clear flow visibility | Orchestrator is a coordination point |
| Centralized error handling and compensation | Slightly more coupling to orchestrator |
| Easy to monitor saga state | Orchestrator must be highly available |

**Banking Recommendation:**
- **Orchestration** for complex financial flows (transfers, loan origination, onboarding) — visibility and compensation are critical
- **Choreography** for simple, loosely coupled flows (notifications, analytics, audit logging)

---

## 6. Transactional Outbox Pattern

Ensures reliable event publishing alongside database transactions.

```
1. Service writes to database AND outbox table in same transaction:
   BEGIN TRANSACTION
     INSERT INTO transfers (id, amount, status) VALUES (...)
     INSERT INTO outbox (event_id, event_type, payload, created_at) VALUES (...)
   COMMIT

2. Outbox publisher (separate process) reads outbox table and publishes to Kafka:
   SELECT * FROM outbox WHERE published = false ORDER BY created_at
   → Publish to Kafka
   → UPDATE outbox SET published = true WHERE event_id = ...

3. Alternative: Change Data Capture (CDC) via Debezium
   → Debezium reads database transaction log
   → Publishes changes to Kafka automatically
```

**When to Use:**
- Any service that must update its database AND publish an event atomically
- All financial transaction services (transfer, payment, account)
- Prevents: lost events (DB committed but event not published) and ghost events (event published but DB rolled back)

---

## 7. Integration Security Standards

| Concern | Synchronous | Asynchronous |
|---|---|---|
| Authentication | OAuth 2.0 Bearer token (external); mTLS (internal) | mTLS for broker connection; SASL for Kafka |
| Authorization | Scope-based per endpoint | Topic-level ACLs (produce/consume permissions) |
| Encryption in transit | TLS 1.2+ for all HTTP; mTLS for service mesh | TLS for broker connections; encrypted topics |
| Encryption at rest | N/A (stateless) | Encrypted message storage (Kafka at-rest encryption) |
| Sensitive data | Mask in logs; no PII in URLs | No PII in topic names; mask in event metadata |
| Audit | Log all API calls with correlation ID | Log all event publications and consumptions |
| Schema validation | OpenAPI validation middleware | Schema registry validation on publish |
| Poison message | Return 400/422 with details | Route to DLQ; alert; investigate |

---

## 8. Monitoring & Observability

### Synchronous Monitoring

| Metric | Alert Threshold | Tool |
|---|---|---|
| Response time (p95) | > 500ms | APM (Datadog, Dynatrace) |
| Error rate | > 1% | APM + alerting |
| Circuit breaker state | Open | Custom metric + alert |
| Timeout rate | > 0.5% | APM |
| Retry rate | > 5% | Custom metric |

### Asynchronous Monitoring

| Metric | Alert Threshold | Tool |
|---|---|---|
| Consumer lag | > 1000 messages or > 5 minutes | Kafka monitoring (Burrow, Confluent Control Center) |
| DLQ depth | > 0 | Queue monitoring + immediate alert |
| Event processing time | > 5 seconds (p95) | Custom metric |
| Event publish failure | Any occurrence | Producer metric + alert |
| Schema validation failure | Any occurrence | Schema registry metric + alert |
| Partition rebalance | Frequent rebalances | Kafka monitoring |

### End-to-End Flow Monitoring

| Flow | SLA | Measurement |
|---|---|---|
| Transfer initiation → completion | < 5 seconds (real-time) | Correlation ID tracing |
| Transfer → notification delivery | < 30 seconds | Event timestamp comparison |
| Transfer → CTR generation | < 5 minutes | Event timestamp comparison |
| Transfer → balance update | < 2 seconds | Event timestamp comparison |

---

## 9. Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Sync chain (A→B→C→D) | Latency compounds; one failure breaks all | Use async where immediate response not needed |
| No circuit breaker | Cascade failures across services | Circuit breaker on every external call |
| No DLQ | Failed messages lost forever | DLQ on every queue; monitor and replay |
| No idempotency | Duplicate processing on retry | Deduplicate by eventId; idempotent operations |
| Event without schema | Breaking changes break consumers silently | Schema registry with compatibility checks |
| Sync for fire-and-forget | Unnecessary coupling and latency | Use async events for notifications, audit, analytics |
| No timeout | Thread/connection exhaustion | Timeout on every external call |
| Shared database for integration | Tight coupling; schema changes break consumers | Events or APIs for cross-service data |
| No outbox pattern | Lost or ghost events | Transactional outbox or CDC |
| Consuming own events | Circular dependency, infinite loops | Service should not consume events it produces (for same purpose) |
