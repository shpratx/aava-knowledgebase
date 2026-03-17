# Integration Testing — Best Practices
### Banking Domain — Agentic Knowledge Base

---

## Best Practices

 Best Practices

### Design Principles
| Practice | Standard |
|---|---|
| Test with real dependencies where possible | Use Testcontainers for DB, Kafka, Redis; real services in staging |
| Use WireMock/MockServer for external APIs | Simulate core banking, fraud engine, payment network |
| Test failure modes, not just success | Timeout, error, unavailable are more important than happy path |
| Test circuit breaker lifecycle | Closed → open → half-open → closed |
| Test event contracts | Publish and consume with schema validation |
| Test idempotency for every consumer | Deliver same message twice; verify no duplicate effects |
| Test correlation ID end-to-end | Verify propagation through entire chain |
| Test with production-like data volume | Realistic number of records in database |
| Isolate integration tests | Each test manages its own state; no shared data between tests |
| Test transactional outbox | DB commit + event publish atomicity |

### Environment Strategy
| Environment | Dependencies | Use |
|---|---|---|
| Local (Testcontainers) | DB, Kafka, Redis in containers; WireMock for external | Developer testing; CI pipeline |
| Staging | Real services; real infrastructure; synthetic data | Pre-production validation |
| Contract testing | Pact broker; no live services needed | API contract verification |

### Anti-Patterns
| Anti-Pattern | Fix |
|---|---|
| Mocking everything | Use Testcontainers for real DB/Kafka; mock only external APIs |
| Only testing happy path | Test timeout, error, unavailable for every dependency |
| No idempotency tests | Test duplicate delivery for every event consumer |
| Shared test database | Each test manages its own data; Testcontainers per test class |
| No circuit breaker tests | Test open/half-open/close lifecycle |
| Ignoring event ordering | Test out-of-order delivery; verify handling |

---