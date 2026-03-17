# Best Practices for Test Scenario & Test Case Generation — API Features
### Banking Domain — Agentic Knowledge Base

---

## 1. Test Scenario Design Principles

**Coverage Strategy:**
- Every API endpoint must have: happy path, validation, auth, authz, error, and contract tests
- Every error code documented in the OpenAPI spec must have a test that triggers it
- Every security control (auth, IDOR, injection, rate limiting) must be tested
- Test the API contract independently from the UI — API tests should not depend on UI

**Scenario Derivation:**
| Source | Derive |
|---|---|
| OpenAPI spec | Contract tests (schema, types, enums, headers) |
| Acceptance criteria | Functional tests per criterion |
| Security requirements | Auth, authz, injection, rate limit, IDOR tests |
| Error code catalog | One test per documented error code |
| Integration contracts | Downstream dependency tests (available, timeout, error, unavailable) |
| NFR targets | Performance tests (response time, throughput) |

## 2. API-Specific Best Practices

### Functional Testing
| Practice | Standard |
|---|---|
| Test every HTTP method | GET, POST, PUT, PATCH, DELETE — each has different behavior |
| Test all response codes | Every documented status code must be triggered by at least one test |
| Test request validation exhaustively | Missing fields, wrong types, out of range, invalid format, extra fields |
| Test business rules at boundaries | Amount = limit exactly; amount = limit + $0.01; amount = $0.00 |
| Test idempotency | Same key + same body; same key + different body; expired key |
| Test pagination boundaries | page=0, page=1, page=last, pageSize=0, pageSize=max+1, empty results |
| Test state transitions | Valid transitions succeed; invalid transitions return appropriate error |
| Test concurrency | Simultaneous writes to same resource; optimistic locking behavior |
| Verify response masking | Account numbers masked; no PII in responses; no internal IDs |

### Security Testing
| Practice | Standard |
|---|---|
| Test every auth scenario | No token, expired, invalid, wrong algorithm, tampered payload, "none" algorithm |
| Test IDOR on every resource endpoint | Change resource ID to another user's resource; verify 403 |
| Test scope enforcement | Call with wrong scope; call with no scope; call with correct scope |
| Test MFA enforcement | Sensitive endpoints without MFA token; with expired MFA; with valid MFA |
| Test injection on every input | SQL injection, NoSQL injection, command injection, path traversal in all string params |
| Test rate limiting | Exceed limit; verify 429 + Retry-After; verify limit resets |
| Test CORS | Request from allowed origin; disallowed origin; verify no wildcard |
| Test response headers | No Server header; no X-Powered-By; Cache-Control: no-store on sensitive endpoints |
| Test error response safety | No stack traces; no SQL errors; no internal paths; no technology details |
| Test TLS | Reject HTTP; reject TLS < 1.2; accept TLS 1.2+ |

### Contract Testing
| Practice | Standard |
|---|---|
| Validate against OpenAPI spec | Every response validated against documented schema |
| No undocumented fields | Response must not contain fields not in the spec |
| Type checking | Every field type matches spec (string, number, boolean, array) |
| Enum validation | Enum fields only return documented values |
| Nullable handling | Nullable fields are null or correct type; non-nullable fields are never null |
| Header validation | All documented response headers present with correct values |
| Breaking change detection | Run OpenAPI diff (oasdiff) in CI; alert on breaking changes |

### Integration Testing
| Practice | Standard |
|---|---|
| Test with real dependencies (staging) | Full integration test with actual downstream services |
| Test dependency failure modes | Timeout, error response, unavailable — verify fallback behavior |
| Test circuit breaker | Trigger circuit breaker open; verify fallback; verify half-open recovery |
| Test event publishing | Verify correct event published to correct topic with correct schema |
| Test correlation ID propagation | Verify ID flows through all downstream calls and appears in response |
| Test with Testcontainers | Database, Kafka, Redis in containers for reproducible tests |

### Performance Testing
| Practice | Standard |
|---|---|
| Baseline performance | Establish baseline response times for all endpoints |
| Load test at expected volume | Test at expected peak concurrent users/TPS |
| Stress test beyond capacity | Find breaking point; verify graceful degradation |
| Endurance test | Sustained load for extended period; detect memory leaks, connection leaks |
| Test with realistic data volume | Production-equivalent data volume in test database |

## 3. Test Data Management

| Standard | Implementation |
|---|---|
| Synthetic data only | No production data; generated fixtures |
| Parameterized tests | Same test logic with different data sets (valid, invalid, boundary) |
| Data builders | Use builder pattern for test data construction |
| Database state | Each test sets up required state; cleans up after (or uses transactions) |
| Idempotency keys | Generate unique keys per test run to avoid collisions |
| Auth tokens | Generate test tokens with specific scopes/claims; don't use production tokens |

## 4. Test Automation Standards

| Standard | Implementation |
|---|---|
| Framework | REST Assured (Java), Supertest (Node), httpx/pytest (Python), Karate |
| Contract testing | Pact or Spring Cloud Contract |
| Performance | Gatling, k6, or JMeter |
| Security | OWASP ZAP (DAST); custom injection tests |
| CI/CD integration | Run on every PR; block merge on failure |
| Parallel execution | Tests run in parallel; no shared state |
| Environment | Testcontainers for databases/Kafka; WireMock for external services |
| Reporting | Test results in CI dashboard; failure notifications; trend tracking |

## 5. Common Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Testing through UI only | API bugs hidden behind UI validation | Test API independently |
| No contract tests | Breaking changes discovered by consumers in production | OpenAPI validation + Pact |
| Hardcoded auth tokens | Tests break when tokens expire | Generate tokens in test setup |
| No negative tests | Only happy path tested; errors unhandled | Test every error code |
| No IDOR tests | Authorization bypass vulnerabilities | Test resource access with wrong user |
| Shared test database | Tests interfere with each other | Testcontainers or transaction rollback |
| No injection tests | SQL/XSS vulnerabilities shipped | Injection payloads in every string field |
| Testing only 200 responses | Error handling untested | Assert on 400, 401, 403, 404, 409, 422, 429, 500 |
| No performance baseline | Can't detect regressions | Establish and track baseline per endpoint |
| Ignoring response headers | Security headers missing | Assert on all required headers |
