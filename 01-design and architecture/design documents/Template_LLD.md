# Low-Level Design (LLD) Document Template
### Banking Domain — Agentic Knowledge Base

---

## Document Metadata

| Field | Description |
|---|---|
| **Document ID** | LLD-[Service]-[Feature]-[Version] (e.g., LLD-TransferService-DomesticTransfer-v1.0) |
| **HLD Reference** | Parent HLD document ID |
| **Service** | Microservice this LLD covers |
| **Author** | Tech Lead / Senior Developer |
| **Reviewers** | Architect, Security Engineer, DBA, Peer Developer |
| **Status** | Draft / In Review / Approved / Superseded |
| **Version** | Semantic version |
| **Related Documents** | HLD, User Stories, API spec (OpenAPI), ADRs |

---

## 1. Service Overview

| Field | Content |
|---|---|
| **Service name** | transfer-service |
| **Bounded context** | Payment |
| **Responsibility** | Domestic fund transfer initiation, status tracking, cancellation |
| **Technology stack** | Java 21, Spring Boot 3.x, PostgreSQL 16, Kafka, Redis |
| **Architecture pattern** | Hexagonal (Ports & Adapters) + CQRS |
| **Data classification** | Restricted (account numbers, amounts) |
| **PCI-DSS scope** | No (no card data; card transfers handled by card-service) |

---

## 2. Class/Code Diagram (C4 Level 4)

### 2.1 Package Structure

```
com.bank.transfer/
├── domain/
│   ├── model/
│   │   ├── Transfer (Aggregate Root)
│   │   ├── TransferStatus (Enum)
│   │   ├── Money (Value Object)
│   │   ├── AccountReference (Value Object)
│   │   └── TransferLine (Entity)
│   ├── event/
│   │   ├── TransferInitiated (Domain Event)
│   │   ├── TransferCompleted (Domain Event)
│   │   └── TransferFailed (Domain Event)
│   ├── service/
│   │   └── TransferDomainService
│   ├── repository/
│   │   └── TransferRepository (Port)
│   └── exception/
│       ├── InsufficientFundsException
│       ├── DailyLimitExceededException
│       └── AccountFrozenException
├── application/
│   ├── port/in/
│   │   ├── InitiateTransferUseCase
│   │   ├── GetTransferStatusUseCase
│   │   └── CancelTransferUseCase
│   ├── port/out/
│   │   ├── AccountPort
│   │   ├── FraudPort
│   │   ├── EventPublisherPort
│   │   └── AuditPort
│   └── service/
│       ├── TransferApplicationService
│       └── TransferSagaOrchestrator
├── infrastructure/
│   ├── web/
│   │   ├── TransferController
│   │   ├── dto/ (TransferRequest, TransferResponse, ErrorResponse)
│   │   └── validation/ (custom validators)
│   ├── persistence/
│   │   ├── TransferJpaRepository
│   │   ├── TransferEntity (JPA entity)
│   │   └── TransferMapper (Entity ↔ Domain)
│   ├── messaging/
│   │   ├── KafkaEventPublisher
│   │   └── KafkaEventConsumer
│   ├── client/
│   │   ├── AccountServiceClient (REST)
│   │   ├── FraudServiceClient (gRPC)
│   │   └── config/ (circuit breaker, retry, timeout)
│   └── config/
│       ├── SecurityConfig
│       ├── KafkaConfig
│       ├── DatabaseConfig
│       └── ObservabilityConfig
```

### 2.2 Domain Model Detail

**Aggregate: Transfer**

| Field | Type | Constraints | Encrypted | Indexed |
|---|---|---|---|---|
| transferId | UUID | PK, generated | No | PK |
| sourceAccountId | UUID | FK, not null | No | Yes (composite) |
| beneficiaryId | UUID | FK, not null | No | Yes |
| amount | Money (BigDecimal + Currency) | > 0, max 999,999,999.99 | No (TDE) | No |
| status | TransferStatus (enum) | NOT NULL | No | Yes (partial) |
| reference | String | max 140, alphanumeric | No | Yes |
| correlationId | UUID | NOT NULL | No | Yes |
| fraudScore | BigDecimal | nullable | No | No |
| initiatedAt | Instant | NOT NULL, UTC | No | Yes (composite) |
| completedAt | Instant | nullable, UTC | No | No |
| failureReason | String | nullable | No | No |
| version | Long | Optimistic locking | No | No |

**Value Object: Money**
```java
public record Money(BigDecimal amount, Currency currency) {
    public Money {
        if (amount.scale() > currency.getDefaultFractionDigits())
            throw new IllegalArgumentException("Invalid precision for " + currency);
        if (amount.compareTo(BigDecimal.ZERO) < 0)
            throw new IllegalArgumentException("Amount must be non-negative");
    }
}
```

### 2.3 State Machine

```
INITIATED → FRAUD_CHECK_PENDING → FRAUD_CHECK_PASSED → DEBIT_PENDING → DEBITED
                                → FRAUD_CHECK_FAILED → REJECTED       → CREDIT_PENDING → COMPLETED
                                                                      → CREDIT_FAILED → COMPENSATION_PENDING → COMPENSATED
INITIATED → CANCELLED (if still INITIATED)
```

---

## 3. API Specification Detail

### 3.1 Endpoint: POST /v1/transfers

**Request:**
```yaml
TransferRequest:
  type: object
  required: [sourceAccountId, beneficiaryId, amount, currency]
  properties:
    sourceAccountId:
      type: string
      format: uuid
    beneficiaryId:
      type: string
      format: uuid
    amount:
      type: number
      minimum: 0.01
      maximum: 999999999.99
    currency:
      type: string
      enum: [USD, EUR, GBP]
    reference:
      type: string
      maxLength: 140
      pattern: "^[a-zA-Z0-9 \\-/.]*$"
```

**Response (201):**
```json
{
  "data": {
    "transferId": "uuid",
    "status": "INITIATED",
    "amount": 5000.00,
    "currency": "USD",
    "sourceAccount": "****1234",
    "beneficiary": "****5678",
    "initiatedAt": "2026-03-15T15:30:00Z"
  },
  "meta": { "requestId": "uuid", "timestamp": "..." }
}
```

**Error Responses:** 400 (validation), 401 (auth), 403 (not own account), 409 (duplicate idempotency key), 422 (limit exceeded), 429 (rate limit), 500 (internal)

### 3.2 Headers

| Header | Direction | Required | Purpose |
|---|---|---|---|
| Authorization | Request | Yes | Bearer JWT + MFA token |
| Idempotency-Key | Request | Yes | Duplicate prevention |
| X-Correlation-ID | Both | Yes | Traceability |
| X-Request-ID | Both | Yes | Request identification |

---

## 4. Sequence Diagrams

### 4.1 Happy Path: Initiate Transfer

```
Client → API Gateway → TransferController → TransferApplicationService
  │                                              │
  │  1. Validate request (Bean Validation)       │
  │  2. Verify MFA token                         │
  │  3. Check idempotency key                    │
  │  4. Call AccountPort.getBalance()        → AccountServiceClient → account-service
  │  5. Validate balance ≥ amount                │
  │  6. Validate daily limit                     │
  │  7. Call FraudPort.checkFraud()          → FraudServiceClient → fraud-service (gRPC)
  │  8. Create Transfer aggregate                │
  │  9. Save to database                     → TransferJpaRepository → PostgreSQL
  │  10. Publish TransferInitiated event     → KafkaEventPublisher → Kafka
  │  11. Publish audit event                 → AuditPort → Kafka (audit topic)
  │  12. Return 201 Created                      │
  │← ← ← ← ← ← ← ← ← ← ← ← ← ← ← ← ← ← ←│
```

### 4.2 Saga: Transfer Execution (Async)

```
TransferSagaOrchestrator (consumes TransferInitiated):
  Step 1: DebitAccount command → account-service
    Success: AccountDebited event → proceed to Step 2
    Failure: InsufficientFunds → mark REJECTED, publish TransferFailed
  Step 2: CreditBeneficiary command → account-service
    Success: BeneficiaryCredited → mark COMPLETED, publish TransferCompleted
    Failure: CreditFailed → compensate Step 1 (reverse debit), mark COMPENSATED
  Step 3: SendNotification command → notification-service (fire-and-forget)
```

### 4.3 Error/Exception Flows

Document sequence diagrams for:
- Insufficient funds
- Daily limit exceeded
- Fraud check failure
- Core banking timeout (circuit breaker)
- Duplicate idempotency key
- Account frozen/closed

---

## 5. Database Design Detail

### 5.1 Physical Schema

```sql
CREATE TABLE transfers (
    transfer_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_account_id UUID NOT NULL,
    beneficiary_id  UUID NOT NULL,
    amount          DECIMAL(19,4) NOT NULL CHECK (amount > 0),
    currency        CHAR(3) NOT NULL,
    status          VARCHAR(30) NOT NULL DEFAULT 'INITIATED'
                    CHECK (status IN ('INITIATED','FRAUD_CHECK_PENDING','FRAUD_CHECK_PASSED',
                    'FRAUD_CHECK_FAILED','DEBIT_PENDING','DEBITED','CREDIT_PENDING',
                    'COMPLETED','REJECTED','CANCELLED','COMPENSATION_PENDING','COMPENSATED')),
    reference       VARCHAR(140),
    correlation_id  UUID NOT NULL,
    fraud_score     DECIMAL(5,2),
    failure_reason  VARCHAR(500),
    initiated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    completed_at    TIMESTAMPTZ,
    version         BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes
CREATE INDEX idx_transfers_source_initiated ON transfers (source_account_id, initiated_at DESC);
CREATE INDEX idx_transfers_correlation ON transfers (correlation_id);
CREATE INDEX idx_transfers_status ON transfers (status) WHERE status IN ('INITIATED','DEBIT_PENDING','CREDIT_PENDING');
CREATE UNIQUE INDEX idx_transfers_idempotency ON transfers (correlation_id);

-- Audit trigger
CREATE TRIGGER trg_transfers_audit
  AFTER INSERT OR UPDATE OR DELETE ON transfers
  FOR EACH ROW EXECUTE FUNCTION audit_trigger_func();

-- Partitioning (monthly)
-- transfers table partitioned by initiated_at (see Part 2 DB best practices)
```

### 5.2 Migration Scripts

| Version | Description | Rollback |
|---|---|---|
| V1.0.0 | Create transfers table + indexes | DROP TABLE transfers |
| V1.0.1 | Add audit trigger | DROP TRIGGER trg_transfers_audit |
| V1.0.2 | Create idempotency index | DROP INDEX idx_transfers_idempotency |

### 5.3 Query Patterns

| Query | Expected Frequency | Index Used | Target Latency |
|---|---|---|---|
| Get transfer by ID | High | PK | < 5ms |
| List transfers by account + date range | High | idx_transfers_source_initiated | < 50ms |
| Get transfer by correlation ID | Medium | idx_transfers_correlation | < 5ms |
| Count pending transfers (for limits) | High | idx_transfers_status | < 10ms |

---

## 6. Integration Detail

### 6.1 Synchronous Integrations

| Dependency | Protocol | Endpoint | Timeout | Retry | Circuit Breaker | Fallback |
|---|---|---|---|---|---|---|
| account-service | REST | GET /v1/accounts/{id}/balance | 5s | 2 retries, exp backoff | Open at 50% failure | Reject transfer |
| fraud-service | gRPC | FraudCheck.evaluate() | 2s | 1 retry | Open at 50% failure | Hold for manual review |

### 6.2 Asynchronous Integrations

| Event/Command | Topic | Partition Key | Schema | Consumer Group |
|---|---|---|---|---|
| TransferInitiated | payment.transfer.initiated | transferId | Avro v1.0 | saga-orchestrator |
| DebitAccount (command) | account.commands.debit | accountId | Avro v1.0 | account-service |
| TransferCompleted | payment.transfer.completed | transferId | Avro v1.0 | notification-svc, audit-svc |

### 6.3 Event Schemas

```json
// TransferInitiated (Avro-compatible JSON representation)
{
  "eventId": "uuid",
  "eventType": "transfer.initiated",
  "eventVersion": "1.0",
  "source": "transfer-service",
  "timestamp": "ISO-8601",
  "correlationId": "uuid",
  "data": {
    "transferId": "uuid",
    "sourceAccountId": "uuid",
    "beneficiaryId": "uuid",
    "amount": 5000.00,
    "currency": "USD",
    "reference": "Invoice 2026-001"
  }
}
```

---

## 7. Security Implementation Detail

### 7.1 Input Validation Rules

| Field | Validations | Implementation |
|---|---|---|
| sourceAccountId | UUID format; must belong to authenticated user | @NotNull @ValidUUID; ResourceOwnershipCheck |
| beneficiaryId | UUID format; must be active and registered by user | @NotNull @ValidUUID; BeneficiaryOwnershipCheck |
| amount | > 0; ≤ daily limit; ≤ 999,999,999.99; currency-appropriate precision | @DecimalMin @DecimalMax; DailyLimitValidator |
| currency | Must be in allowed list | @ValidEnum |
| reference | Max 140 chars; alphanumeric + -/. only | @Size @Pattern |

### 7.2 Authorization Checks

```java
// Resource-level authorization
@PreAuthorize("hasScope('transfers:write')")
public TransferResponse initiateTransfer(TransferRequest request, Authentication auth) {
    // Verify source account belongs to authenticated user
    Account account = accountPort.getAccount(request.sourceAccountId());
    if (!account.customerId().equals(auth.getCustomerId())) {
        auditPort.logUnauthorizedAccess(auth, request.sourceAccountId());
        throw new ForbiddenException("Access denied");
    }
    // ... proceed with transfer
}
```

### 7.3 Sensitive Data Handling

| Data | In Request | In Response | In Logs | In Database | In Events |
|---|---|---|---|---|---|
| Account ID | UUID (no masking) | Masked (****1234) | Masked | UUID (not masked) | UUID |
| Amount | Plain | Plain | Plain | Plain (TDE) | Plain |
| Customer name | Not in request | Not in response | Never | Encrypted (AES-256) | Never |
| Correlation ID | Header | Header + body | Full | Full | Full |

---

## 8. Error Handling Detail

| Error Scenario | HTTP Status | Error Code | User Message | Logging | Alert |
|---|---|---|---|---|---|
| Validation failure | 400 | VALIDATION_ERROR | Field-level details | INFO | No |
| Unauthorized | 401 | AUTH_REQUIRED | "Authentication required" | WARN | If repeated |
| Not own account | 403 | ACCESS_DENIED | "Access denied" | WARN + audit | Yes |
| Insufficient funds | 422 | INSUFFICIENT_FUNDS | "Insufficient funds" | INFO | No |
| Daily limit exceeded | 422 | LIMIT_EXCEEDED | "Daily limit exceeded. Remaining: $X" | INFO | No |
| Fraud check failed | 422 | TRANSFER_HELD | "Transfer under review" | WARN + audit | Yes |
| Duplicate idempotency | 409 | DUPLICATE_REQUEST | Return original response | INFO | No |
| Core banking timeout | 502 | UPSTREAM_TIMEOUT | "Service temporarily unavailable" | ERROR | Yes |
| Internal error | 500 | INTERNAL_ERROR | "Something went wrong" | ERROR + stack trace (server only) | Yes |

---

## 9. Testing Strategy

| Test Type | Scope | Coverage Target | Tool |
|---|---|---|---|
| Unit | Domain model, validators, mappers | ≥ 90% (domain), ≥ 80% (overall) | JUnit 5, Mockito |
| Integration | Repository, Kafka, REST clients | All integration points | Testcontainers, WireMock |
| Contract | API contract, event schemas | All endpoints and events | Pact, Spring Cloud Contract |
| Security | Auth bypass, IDOR, injection | All OWASP Top 10 applicable | OWASP ZAP, custom tests |
| Performance | Load, stress | p95 < 500ms at 500 concurrent | Gatling |

---

## 10. Configuration

| Config | Dev | Staging | Production |
|---|---|---|---|
| DB connection pool (write) | 5 | 10 | 20 |
| DB connection pool (read) | 5 | 20 | 50 |
| Kafka bootstrap | localhost:9092 | kafka-staging:9092 | kafka-prod:9092 (via vault) |
| Fraud service timeout | 5s | 2s | 2s |
| Circuit breaker threshold | 80% | 50% | 50% |
| Rate limit (transfers) | 100/min | 10/min | 10/min |
| Feature flags | All enabled | Canary | Gradual rollout |

---

## Approval Sign-Off

| Role | Name | Date | Decision |
|---|---|---|---|
| Tech Lead | | | Approved / Rejected |
| Architect | | | Approved / Rejected |
| Security Engineer | | | Approved / Rejected |
| DBA | | | Approved / Rejected |
