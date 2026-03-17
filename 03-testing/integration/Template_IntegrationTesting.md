# Integration Testing — Template, Best Practices & Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Part A: Test Scenario Template

### Metadata
| Field | Description |
|---|---|
| **Scenario ID** | TS-INT-[Module]-[Seq] (e.g., TS-INT-TRF-001) |
| **Integration** | Source service → Target service/system |
| **Protocol** | REST / gRPC / Kafka / Database / External API |
| **Pattern** | Synchronous / Asynchronous / Request-Reply |
| **Priority** | Critical / High / Medium / Low |
| **Preconditions** | Services running; test data seeded; dependencies available |

### Integration Test Categories

#### 1. Service-to-Service (Synchronous)
| Category | Scenarios |
|---|---|
| Happy path | Valid request → correct response; correct status code; correct data |
| Timeout | Dependency exceeds timeout → appropriate error; no hang; circuit breaker triggered |
| Error response | Dependency returns 4xx/5xx → translated error; no downstream details leaked |
| Unavailable | Dependency down → circuit breaker opens; fallback behavior; graceful degradation |
| Slow response | Dependency responds slowly → timeout fires; bulkhead prevents thread exhaustion |
| Retry | Transient failure → retry with backoff; idempotent retry; max retries respected |
| Circuit breaker | Open → fallback; half-open → test request; close → normal flow |
| Correlation ID | ID propagated from source → target → response; present in all logs |
| Auth propagation | Token/mTLS propagated correctly to downstream; unauthorized downstream → appropriate error |
| Data contract | Response matches expected schema; no breaking changes from dependency |

#### 2. Event-Driven (Asynchronous)
| Category | Scenarios |
|---|---|
| Event publish | Action triggers event on correct topic; correct partition key; correct schema; correct headers |
| Event consume | Event received → correct state change; correct side effects; acknowledgment sent |
| Idempotent consume | Same event delivered twice → processed only once; no duplicate side effects |
| Out-of-order events | Events arrive out of order → handled correctly (or rejected with appropriate error) |
| Poison message | Malformed event → routed to DLQ after retries; not blocking other messages |
| DLQ | Failed message lands in DLQ with original headers; DLQ alert triggered |
| Schema evolution | New schema version consumed by old consumer → backward compatible |
| Consumer lag | Consumer falls behind → lag monitored; alert triggered; catches up without data loss |
| Exactly-once semantics | Transactional outbox → event published if and only if DB committed |

#### 3. Database Integration
| Category | Scenarios |
|---|---|
| CRUD operations | Create, read, update, delete → correct data persisted; correct response |
| Transaction boundaries | Multi-step operation → all committed or all rolled back |
| Connection pool | Pool exhaustion → appropriate error; no hang; pool recovers |
| Migration | Schema migration → application works with new schema; rollback works |
| Concurrent access | Simultaneous writes → optimistic locking; no lost updates |
| Read replica | Read from replica → data eventually consistent; lag within tolerance |

#### 4. External System Integration
| Category | Scenarios |
|---|---|
| Core banking | Balance inquiry, debit, credit → correct response; timeout handling; error handling |
| Fraud engine | Fraud check → approve/hold/reject; timeout → hold for manual review (never auto-approve) |
| Payment network | Payment submission → acknowledgment; settlement; rejection handling |
| Notification service | Trigger → email/SMS/push delivered; failure → queued for retry; no blocking |
| Sanctions screening | Screen → clear/match/error; timeout → block transaction pending manual review |

### Test Case Template

| Field | Description |
|---|---|
| **Test Case ID** | TC-INT-[Module]-[Seq] |
| **Scenario** | Parent scenario ID |
| **Setup** | Services, data, mocks/stubs required |
| **Action** | What triggers the integration |
| **Assertions** | What to verify (response, state change, event, log, metric) |
| **Teardown** | Cleanup after test |

### Example: Transfer Service Integration Tests

**TC-INT-TRF-001: Account Service — Balance Check Success**
```
Setup:   Transfer service running; Account service running (or WireMock stub returning balance $45,000)
Action:  POST /v1/transfers with amount $5,000
Assert:  Account service called with correct accountId
         Balance check passed (sufficient funds)
         Transfer created with status INITIATED
         Correlation ID propagated to account service call
```

**TC-INT-TRF-002: Fraud Service — Timeout → Hold Transfer**
```
Setup:   Fraud service stub configured with 10-second delay (exceeds 2s timeout)
Action:  POST /v1/transfers
Assert:  Fraud service call times out after 2 seconds
         Transfer created with status FRAUD_REVIEW_PENDING
         Circuit breaker records failure
         Customer notified "Transfer under review"
         Fraud team alert triggered
         Transfer NOT auto-approved
```

**TC-INT-TRF-003: Kafka — TransferCompleted Event Published**
```
Setup:   Transfer in DEBITED status; Kafka consumer listening on payment.transfer.completed
Action:  Credit beneficiary succeeds → transfer status → COMPLETED
Assert:  Event published to payment.transfer.completed topic
         Partition key = transferId
         Event schema matches Avro v1.0
         Event contains: transferId, amount, currency, sourceAccountId, beneficiaryId
         Correlation ID in event headers matches original request
```

**TC-INT-TRF-004: Kafka — Duplicate Event Consumption (Idempotency)**
```
Setup:   TransferCompleted event already processed (balance updated)
Action:  Same event delivered again (at-least-once delivery)
Assert:  Event processed without error
         Balance NOT updated again (idempotent)
         No duplicate notification sent
         Log indicates duplicate detected
```

---