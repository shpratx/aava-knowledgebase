# Guardrails for API Test Scenario & Test Case Generation
### Banking Domain — Agentic Knowledge Base

---

## 1. Coverage Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-CG-001 | Every API endpoint must have: happy path, validation, auth, authz, and error test cases | Reject if any category missing |
| TG-API-CG-002 | Every documented error code in the OpenAPI spec must have a test that triggers it | Reject if any error code untested |
| TG-API-CG-003 | Every request field must have validation tests: missing, wrong type, out of range, invalid format | Reject if field validation untested |
| TG-API-CG-004 | Every response must be validated against the OpenAPI schema (contract test) | Reject if no contract validation |
| TG-API-CG-005 | Both positive and negative scenarios must be covered for every endpoint | Reject if only happy path |
| TG-API-CG-006 | Boundary value tests for all numeric parameters (min, max, limit, limit±1, zero, negative) | Reject if no boundary tests |
| TG-API-CG-007 | Pagination tests: first page, last page, empty results, max page size, invalid page | Reject if pagination untested for collection endpoints |
| TG-API-CG-008 | Idempotency tests for all POST endpoints with Idempotency-Key | Reject if idempotency untested for financial POSTs |

---

## 2. Security Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-SG-001 | Auth test: every endpoint tested with no token, expired token, invalid token → verify 401 | Reject if auth not tested |
| TG-API-SG-002 | IDOR test: every resource endpoint tested with another user's resource ID → verify 403 | Reject if IDOR not tested |
| TG-API-SG-003 | Scope test: every endpoint tested with wrong OAuth scope → verify 403 | Reject if scope enforcement not tested |
| TG-API-SG-004 | MFA test: endpoints requiring MFA tested without MFA token → verify 403 | Reject if MFA not tested |
| TG-API-SG-005 | Injection test: SQL injection payload in every string parameter | Reject if injection not tested |
| TG-API-SG-006 | Rate limit test: exceed limit → verify 429 + Retry-After header | Reject if rate limiting not tested |
| TG-API-SG-007 | Response security test: verify no stack traces, no internal paths, no server headers | Reject if response security not tested |
| TG-API-SG-008 | Token manipulation test: modified JWT payload, wrong algorithm, "none" algorithm → verify 401 | Reject if token manipulation not tested |
| TG-API-SG-009 | Sensitive data masking test: verify account numbers masked, no PII in responses | Reject if masking not verified |
| TG-API-SG-010 | CORS test: verify no wildcard origin in production; only allowed origins accepted | Reject if CORS not tested |

---

## 3. Contract Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-CT-001 | Every API response must be validated against the OpenAPI schema — no extra or missing fields | CI/CD gate |
| TG-API-CT-002 | Field types must match spec exactly (string, number, boolean, array) | CI/CD gate |
| TG-API-CT-003 | Enum fields must only return documented values | CI/CD gate |
| TG-API-CT-004 | Required response headers must be verified (Correlation-ID, RateLimit, Cache-Control) | Reject if headers not verified |
| TG-API-CT-005 | Breaking change detection must run in CI (OpenAPI diff) | CI/CD gate |

---

## 4. Integration Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-IG-001 | Downstream dependency timeout must be tested → verify appropriate error and circuit breaker | Reject if timeout not tested |
| TG-API-IG-002 | Downstream dependency error must be tested → verify no downstream details leaked | Reject if error translation not tested |
| TG-API-IG-003 | Downstream dependency unavailable must be tested → verify fallback behavior | Reject if fallback not tested |
| TG-API-IG-004 | Event publishing must be tested → verify correct event on correct topic with correct schema | Reject if event publishing not tested |
| TG-API-IG-005 | Correlation ID propagation must be tested → verify ID in all downstream calls and response | Reject if correlation not tested |

---

## 5. Performance Test Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-PG-001 | Performance baseline must be established for all endpoints before production | Reject if no baseline |
| TG-API-PG-002 | Load test must verify p95 response time within NFR target at expected peak load | Reject if performance target not met |
| TG-API-PG-003 | Performance test must use production-equivalent data volume | Flag if test data volume is trivial |
| TG-API-PG-004 | Performance regression > 20% from baseline must be investigated | CI/CD gate — alert on regression |

---

## 6. Test Quality Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-API-QG-001 | Each test must be independent — no shared state or execution order dependency | Reject if tests depend on each other |
| TG-API-QG-002 | Test data must be synthetic — no production data | Reject if production data used |
| TG-API-QG-003 | Auth tokens must be generated in test setup — not hardcoded | Reject if hardcoded tokens |
| TG-API-QG-004 | Tests must clean up created resources (or use transaction rollback) | Reject if tests leave orphan data |
| TG-API-QG-005 | API test suite must complete within 10 minutes | Flag if suite exceeds 10 minutes |
| TG-API-QG-006 | Flaky tests quarantined within 24 hours; fixed within 1 week | Process guardrail |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot release | Missing auth test, no IDOR test, no injection test, no contract validation, error code untested |
| **Flag** | Address before next release | Trivial test data volume, suite > 10 min, no performance baseline |
| **CI/CD Gate** | Automated | Schema validation, breaking change detection, performance regression, coverage threshold |

---

## Pre-Release API Test Checklist

| # | Check |
|---|---|
| 1 | Every endpoint has happy path + validation + auth + authz + error tests |
| 2 | Every documented error code is triggered by a test |
| 3 | Every request field has validation tests (missing, wrong type, out of range) |
| 4 | Every response validated against OpenAPI schema |
| 5 | Auth tested: no token, expired, invalid, tampered → 401 |
| 6 | IDOR tested: other user's resource → 403 |
| 7 | Scope tested: wrong scope → 403 |
| 8 | MFA tested: without MFA token → 403 |
| 9 | Injection tested: SQL/XSS payloads in all string fields |
| 10 | Rate limiting tested: exceed limit → 429 + Retry-After |
| 11 | Response masking verified: account numbers masked, no PII |
| 12 | Response headers verified: Correlation-ID, RateLimit, Cache-Control |
| 13 | Idempotency tested for financial POST endpoints |
| 14 | Downstream failure modes tested (timeout, error, unavailable) |
| 15 | Correlation ID propagation verified |
| 16 | Performance baseline established; p95 within target |
| 17 | All tests independent; synthetic data; no hardcoded tokens |
