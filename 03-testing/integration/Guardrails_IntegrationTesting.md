# Integration Testing — Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Guardrails

 Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-INT-001 | Every synchronous dependency must have timeout + error + unavailable tests | Reject if failure modes untested |
| TG-INT-002 | Every circuit breaker must have open/half-open/close lifecycle test | Reject if circuit breaker untested |
| TG-INT-003 | Every event consumer must have idempotency test (duplicate delivery) | Reject if idempotency untested |
| TG-INT-004 | Every event must be tested for schema compliance (Avro/Protobuf validation) | Reject if schema not validated |
| TG-INT-005 | Every message queue must have DLQ test (poison message → DLQ) | Reject if DLQ untested |
| TG-INT-006 | Correlation ID propagation must be tested end-to-end | Reject if correlation not verified |
| TG-INT-007 | Database transaction boundaries must be tested (commit + rollback) | Reject if transaction behavior untested |
| TG-INT-008 | Downstream error responses must not leak to upstream consumers | Reject if error translation untested |
| TG-INT-009 | Fraud service timeout must result in HOLD, never auto-approve | Reject if fraud timeout behavior untested |
| TG-INT-010 | Transactional outbox must be tested (DB commit ↔ event publish atomicity) | Reject if outbox untested |
| TG-INT-011 | Integration tests must use Testcontainers or equivalent — not shared test databases | Reject if shared DB used |
| TG-INT-012 | Integration test suite must complete within 10 minutes | Flag if exceeds 10 minutes |

### Pre-Release Integration Checklist
| # | Check |
|---|---|
| 1 | Every sync dependency: happy path + timeout + error + unavailable tested |
| 2 | Circuit breakers tested (open/half-open/close) |
| 3 | Every event: publish verified (topic, key, schema, headers) |
| 4 | Every consumer: idempotency tested (duplicate delivery) |
| 5 | DLQ tested for every queue (poison message routing) |
| 6 | Correlation ID propagated end-to-end |
| 7 | Transaction boundaries tested (commit + rollback) |
| 8 | Downstream errors translated (no details leaked) |
| 9 | Fraud timeout → hold (never auto-approve) |
| 10 | Transactional outbox atomicity verified |