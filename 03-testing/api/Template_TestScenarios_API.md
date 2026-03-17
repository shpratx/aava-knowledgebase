# Test Scenario Template — API Features
---

## Test Scenario Metadata

| Field | Description |
|---|---|
| **Scenario ID** | TS-API-[Module]-[Seq] (e.g., TS-API-TRF-001) |
| **API Endpoint** | Method + Path (e.g., POST /v1/transfers) |
| **Feature** | Feature/User Story ID |
| **Type** | Functional / Security / Performance / Contract / Integration / Error Handling |
| **Priority** | Critical / High / Medium / Low |
| **Preconditions** | Auth tokens, test data, system state |
| **Test Data** | Request payloads, expected responses |
| **Automation** | Automated (Postman/Newman, REST Assured, Karate) / Manual |

---

## API Test Scenario Categories

### 1. Functional Scenarios

| Category | Scenarios to Cover |
|---|---|
| Happy path | Valid request → expected success response (200/201/204) |
| Request validation | Missing required fields; invalid types; out-of-range values; invalid format; extra unknown fields |
| Business rules | Insufficient funds; daily limit exceeded; account frozen; beneficiary inactive; duplicate idempotency key |
| Response validation | Correct status code; correct response schema; correct data values; masked sensitive data |
| Pagination | First page; last page; page size limits; empty results; total count accuracy |
| Filtering/sorting | Valid filters; invalid filter values; combined filters; sort order verification |
| Idempotency | Same key + same body → same response; same key + different body → 422; expired key → new request |
| State transitions | Resource state changes correctly (INITIATED → COMPLETED; ACTIVE → BLOCKED) |
| Concurrency | Simultaneous requests to same resource; optimistic locking; race conditions |

### 2. Security Scenarios

| Category | Scenarios to Cover |
|---|---|
| Authentication | No token → 401; expired token → 401; invalid token → 401; malformed token → 401 |
| Authorization (scope) | Missing required scope → 403; wrong scope → 403; correct scope → success |
| Authorization (resource) | Access own resource → success; access other user's resource → 403 (IDOR test) |
| MFA | Endpoint requiring MFA without MFA token → 403; with valid MFA → success; expired MFA → 403 |
| Input injection | SQL injection in all string params; XSS payloads; command injection; path traversal (../) |
| Rate limiting | Exceed rate limit → 429 with Retry-After; within limit → success; rate limit headers present |
| CORS | Request from allowed origin → success; disallowed origin → blocked; no wildcard in production |
| TLS | HTTP request → rejected or redirected; TLS 1.0/1.1 → rejected; TLS 1.2+ → success |
| Response security | No stack traces in errors; no internal paths; no server version headers; sensitive data masked |
| Token manipulation | Modified JWT payload → 401; changed algorithm → 401; "none" algorithm → 401 |

### 3. Error Handling Scenarios

| Category | Scenarios to Cover |
|---|---|
| 400 Bad Request | Invalid JSON; wrong Content-Type; validation failures with field-level details |
| 401 Unauthorized | All auth failure scenarios (no differentiation in error message) |
| 403 Forbidden | Resource ownership failure; scope failure; MFA required |
| 404 Not Found | Non-existent resource ID; deleted resource |
| 409 Conflict | Duplicate idempotency key; state conflict (cancel completed transfer) |
| 422 Unprocessable | Business rule violations (insufficient funds, limit exceeded) |
| 429 Rate Limited | Rate limit exceeded; Retry-After header present and accurate |
| 500 Internal Error | Simulated server error → generic message, no technical details |
| 502/503 | Downstream dependency failure → appropriate error, no downstream details exposed |
| Error schema | All errors follow consistent ErrorResponse schema |

### 4. Contract Scenarios

| Category | Scenarios to Cover |
|---|---|
| Schema compliance | Response matches OpenAPI schema exactly (no extra/missing fields) |
| Type compliance | Field types match spec (string, number, boolean, array, null) |
| Enum compliance | Enum fields only return documented values |
| Header compliance | Required response headers present (Correlation-ID, Rate-Limit, Cache-Control) |
| Versioning | v1 endpoint returns v1 schema; deprecated endpoint returns Deprecation + Sunset headers |

### 5. Integration Scenarios

| Category | Scenarios to Cover |
|---|---|
| Downstream available | Full flow with all dependencies → success |
| Downstream timeout | Dependency times out → appropriate error; circuit breaker behavior |
| Downstream error | Dependency returns error → translated error; no downstream details leaked |
| Downstream unavailable | Dependency down → fallback behavior; circuit breaker opens |
| Event publishing | Action triggers correct event on correct topic with correct schema |
| Event consumption | Consuming event triggers correct state change; idempotent on replay |
| Correlation ID | ID propagated through all downstream calls; present in response and logs |

---

## API Test Case Template

| Field | Description |
|---|---|
| **Test Case ID** | TC-API-[Module]-[Seq] |
| **Scenario ID** | Parent test scenario |
| **Title** | Descriptive title |
| **Type** | Functional / Security / Contract / Integration / Error |
| **Priority** | Critical / High / Medium / Low |
| **Preconditions** | Auth state, test data, system state |
| **Request** | Method, URL, headers, body |
| **Expected Response** | Status code, headers, body (schema + values) |
| **Assertions** | Specific checks to perform |
| **Actual Response** | (Filled during execution) |
| **Status** | Pass / Fail / Blocked |
| **Defect ID** | Linked defect if failed |

---

## Example: Fund Transfer API Test Cases

### TC-API-TRF-001: Happy Path — Initiate Transfer

| Field | Value |
|---|---|
| **Priority** | Critical |
| **Preconditions** | Valid Bearer token (scope: transfers:write); valid MFA token; account ACC-001234 with $45,000 balance; beneficiary BEN-005678 active |

**Request:**
```http
POST /v1/transfers
Authorization: Bearer {valid_token}
X-MFA-Token: {valid_mfa_token}
Idempotency-Key: 550e8400-e29b-41d4-a716-446655440000
X-Correlation-ID: corr-test-001
Content-Type: application/json

{
  "sourceAccountId": "ACC-001234",
  "beneficiaryId": "BEN-005678",
  "amount": 5000.00,
  "currency": "USD",
  "reference": "Invoice 2026-001"
}
```

**Expected Response:**
```http
HTTP/1.1 201 Created
X-Correlation-ID: corr-test-001
X-Request-ID: {uuid}
X-RateLimit-Limit: 10
X-RateLimit-Remaining: 9
Cache-Control: no-store
Content-Type: application/json

{
  "data": {
    "transferId": "{uuid}",
    "status": "INITIATED",
    "amount": 5000.00,
    "currency": "USD",
    "sourceAccount": "****1234",
    "beneficiary": "****5678",
    "initiatedAt": "{ISO-8601}"
  },
  "meta": {
    "requestId": "{uuid}",
    "timestamp": "{ISO-8601}"
  }
}
```

**Assertions:**
1. Status code is 201
2. Response body matches TransferSuccessResponse schema
3. transferId is a valid UUID
4. sourceAccount is masked (****1234, not full number)
5. beneficiary is masked (****5678)
6. status is "INITIATED"
7. amount matches request (5000.00)
8. X-Correlation-ID echoed from request
9. X-RateLimit headers present
10. Cache-Control: no-store

### TC-API-TRF-002: Validation — Missing Required Field

**Request:** Same as TC-001 but body missing `amount` field.

**Expected:** 400 with field-level error:
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Request validation failed",
    "details": [{"field": "amount", "message": "Amount is required"}]
  }
}
```

**Assertions:** Status 400; error.code = "VALIDATION_ERROR"; details[0].field = "amount"

### TC-API-TRF-003: Security — Access Other User's Account (IDOR)

**Request:** Valid token for User A; sourceAccountId belongs to User B.

**Expected:** 403 Forbidden
```json
{"error": {"code": "ACCESS_DENIED", "message": "Access denied"}}
```

**Assertions:** Status 403; no data about User B's account in response; access attempt audit-logged.

### TC-API-TRF-004: Security — SQL Injection in Reference Field

**Request:** reference field contains: `'; DROP TABLE transfers; --`

**Expected:** 400 (pattern validation failure) or 422 (sanitized and rejected)

**Assertions:** Status 400 or 422; no SQL execution; no database error in response; attempt logged.

### TC-API-TRF-005: Security — No Auth Token

**Request:** Same as TC-001 but no Authorization header.

**Expected:** 401 Unauthorized
```json
{"error": {"code": "AUTH_REQUIRED", "message": "Authentication required"}}
```

**Assertions:** Status 401; no differentiation between missing/invalid/expired token; no data leaked.

### TC-API-TRF-006: Idempotency — Duplicate Request

**Request:** Same Idempotency-Key and same body as TC-001 (already processed).

**Expected:** 201 with same response as original (cached response returned).

**Assertions:** Status 201; response body identical to original; no duplicate transfer created.

### TC-API-TRF-007: Rate Limiting — Exceed Limit

**Request:** Send 11 requests in 1 minute (limit is 10/min).

**Expected:** 11th request returns 429:
```http
HTTP/1.1 429 Too Many Requests
Retry-After: 30
```

**Assertions:** Status 429; Retry-After header present; X-RateLimit-Remaining = 0.

### TC-API-TRF-008: Contract — Response Schema Validation

**Request:** Valid transfer request.

**Assertions:**
1. Response matches OpenAPI TransferSuccessResponse schema exactly
2. No extra undocumented fields in response
3. All field types match spec (transferId: string, amount: number, status: string enum)
4. Enum values are only documented values (INITIATED, PROCESSING, COMPLETED, FAILED, CANCELLED)
5. Nullable fields are null or correct type (never wrong type)
