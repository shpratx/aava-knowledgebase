# Best Practices & Standards for RESTful API Design and Architecture — Part 1

---

## 1. RESTful Design Principles (ref: DA-BP-006)

### HTTP Methods

| Method | Purpose | Idempotent | Safe | Banking Example |
|---|---|---|---|---|
| GET | Retrieve resource(s) | Yes | Yes | GET /v1/accounts/{id}/balance |
| POST | Create resource or trigger action | No | No | POST /v1/transfers |
| PUT | Full resource replacement | Yes | No | PUT /v1/beneficiaries/{id} |
| PATCH | Partial resource update | No* | No | PATCH /v1/customers/{id}/address |
| DELETE | Remove resource | Yes | No | DELETE /v1/beneficiaries/{id} |
| HEAD | Retrieve headers only (no body) | Yes | Yes | HEAD /v1/accounts/{id} (existence check) |
| OPTIONS | Retrieve allowed methods (CORS preflight) | Yes | Yes | OPTIONS /v1/transfers |

*PATCH can be made idempotent with conditional headers (If-Match).

### Resource Naming Conventions

| Principle | Standard | Good Example | Bad Example |
|---|---|---|---|
| Use nouns, not verbs | Resources are things, not actions | /v1/transfers | /v1/createTransfer |
| Use plural nouns | Collections are plural | /v1/accounts | /v1/account |
| Use kebab-case | Lowercase with hyphens | /v1/standing-orders | /v1/standingOrders |
| Hierarchical relationships | Nest related resources | /v1/accounts/{id}/transactions | /v1/account-transactions |
| Max nesting depth | 3 levels maximum | /v1/accounts/{id}/transactions/{txId} | /v1/banks/{bId}/branches/{brId}/accounts/{aId}/transactions/{tId} |
| No file extensions | Content negotiation via Accept header | /v1/statements (Accept: application/pdf) | /v1/statements.pdf |
| No trailing slashes | Consistent URL format | /v1/accounts | /v1/accounts/ |
| No CRUD in URL | HTTP method conveys the action | POST /v1/transfers | POST /v1/transfers/create |

### Banking API Resource Model

```
/v1/customers
/v1/customers/{customerId}
/v1/customers/{customerId}/accounts
/v1/accounts/{accountId}
/v1/accounts/{accountId}/balance
/v1/accounts/{accountId}/transactions
/v1/accounts/{accountId}/transactions/{transactionId}
/v1/accounts/{accountId}/statements
/v1/transfers
/v1/transfers/{transferId}
/v1/transfers/{transferId}/status
/v1/beneficiaries
/v1/beneficiaries/{beneficiaryId}
/v1/payments
/v1/payments/{paymentId}
/v1/standing-orders
/v1/standing-orders/{orderId}
/v1/cards
/v1/cards/{cardId}
/v1/cards/{cardId}/limits
```

### HTTP Status Codes

| Code | Meaning | When to Use |
|---|---|---|
| 200 OK | Successful GET, PUT, PATCH, DELETE | Return resource or confirmation |
| 201 Created | Successful POST creating a resource | Return created resource with Location header |
| 202 Accepted | Request accepted for async processing | Transfer queued for processing; return status URL |
| 204 No Content | Successful DELETE or PUT with no response body | Beneficiary deleted |
| 400 Bad Request | Invalid request syntax or parameters | Validation failure (missing field, wrong format) |
| 401 Unauthorized | Authentication required or failed | Missing/invalid/expired token |
| 403 Forbidden | Authenticated but not authorized | Customer accessing another customer's account |
| 404 Not Found | Resource does not exist | Account ID not found |
| 405 Method Not Allowed | HTTP method not supported for this resource | PUT on /v1/transfers (transfers are immutable) |
| 409 Conflict | Request conflicts with current state | Duplicate transfer (idempotency check) |
| 415 Unsupported Media Type | Content-Type not supported | Sending XML when only JSON is accepted |
| 422 Unprocessable Entity | Valid syntax but semantic errors | Transfer amount exceeds daily limit |
| 429 Too Many Requests | Rate limit exceeded | Include Retry-After header |
| 500 Internal Server Error | Unexpected server error | Log details server-side; generic message to client |
| 502 Bad Gateway | Upstream service failure | Core banking system unavailable |
| 503 Service Unavailable | Service temporarily unavailable | Maintenance window; include Retry-After |

### Response Structure Standards

**Success Response:**
```json
{
  "data": {
    "transferId": "TRF-2026-0315-001",
    "status": "COMPLETED",
    "amount": 5000.00,
    "currency": "USD",
    "sourceAccount": "****1234",
    "beneficiary": "****5678",
    "timestamp": "2026-03-15T15:30:00Z"
  },
  "meta": {
    "requestId": "req-abc-123",
    "timestamp": "2026-03-15T15:30:00.123Z"
  }
}
```

**Error Response:**
```json
{
  "error": {
    "code": "TRANSFER_LIMIT_EXCEEDED",
    "message": "Daily transfer limit exceeded. Remaining: $5,000.00",
    "details": [
      {
        "field": "amount",
        "issue": "Exceeds daily limit",
        "limit": 50000.00,
        "used": 45000.00,
        "remaining": 5000.00
      }
    ]
  },
  "meta": {
    "requestId": "req-abc-124",
    "timestamp": "2026-03-15T15:30:01.456Z"
  }
}
```

**Collection Response (with pagination):**
```json
{
  "data": [ ... ],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "totalItems": 156,
    "totalPages": 8
  },
  "links": {
    "self": "/v1/accounts/1234/transactions?page=1&pageSize=20",
    "next": "/v1/accounts/1234/transactions?page=2&pageSize=20",
    "last": "/v1/accounts/1234/transactions?page=8&pageSize=20"
  },
  "meta": {
    "requestId": "req-abc-125",
    "timestamp": "2026-03-15T15:30:02.789Z"
  }
}
```

### Pagination Standards

| Approach | When to Use | Implementation |
|---|---|---|
| Offset-based | Simple lists, moderate data | ?page=2&pageSize=20 |
| Cursor-based | Large datasets, real-time data | ?cursor=eyJpZCI6MTIzfQ&limit=20 |
| Keyset-based | Time-series data (transactions) | ?after=2026-03-01T00:00:00Z&limit=20 |

**Banking Defaults:**
- Default page size: 20
- Maximum page size: 100 (prevent bulk data extraction)
- Transaction history: cursor-based (for consistency with real-time updates)
- Always return total count and navigation links

### Filtering, Sorting, Searching

```
# Filtering
GET /v1/accounts/{id}/transactions?status=COMPLETED&type=TRANSFER&dateFrom=2026-03-01&dateTo=2026-03-15

# Sorting
GET /v1/accounts/{id}/transactions?sort=-timestamp,+amount  (- descending, + ascending)

# Searching
GET /v1/beneficiaries?search=John  (server-side sanitized)

# Field selection (sparse fieldsets)
GET /v1/accounts/{id}?fields=balance,currency,status
```

---

## 2. API Versioning (ref: DA-BP-007)

### Versioning Strategy

| Approach | Format | Pros | Cons | Recommendation |
|---|---|---|---|---|
| URL path | /v1/accounts | Clear, easy to route, cacheable | URL changes on version bump | **Recommended for banking** |
| Header | Api-Version: 1 | Clean URLs | Hidden, harder to test/debug | Acceptable for internal APIs |
| Query param | ?version=1 | Easy to add | Pollutes query string, caching issues | Not recommended |
| Content negotiation | Accept: application/vnd.bank.v1+json | RESTful purist approach | Complex, poor tooling support | Not recommended |

### Versioning Standards

| Standard | Requirement |
|---|---|
| Format | /v{major}/ — e.g., /v1/, /v2/ |
| When to increment | Breaking changes only (removed fields, changed types, removed endpoints, changed behavior) |
| Non-breaking changes | Add new fields, add new endpoints, add optional parameters — no version bump needed |
| Deprecation notice | Minimum 6 months before removal; Deprecation header + Sunset header |
| Concurrent versions | Support maximum 2 concurrent major versions (current + previous) |
| Default version | No default — version is mandatory in URL |
| Documentation | Each version fully documented independently |

### Deprecation Headers

```http
HTTP/1.1 200 OK
Deprecation: Sun, 15 Sep 2026 00:00:00 GMT
Sunset: Sun, 15 Mar 2027 00:00:00 GMT
Link: </v2/accounts>; rel="successor-version"
```

### Breaking vs. Non-Breaking Changes

| Change | Breaking? | Action |
|---|---|---|
| Add optional field to response | No | No version bump |
| Add optional query parameter | No | No version bump |
| Add new endpoint | No | No version bump |
| Remove field from response | Yes | New version |
| Rename field | Yes | New version |
| Change field type (string → number) | Yes | New version |
| Remove endpoint | Yes | New version |
| Change error response structure | Yes | New version |
| Make optional field required | Yes | New version |
| Change authentication mechanism | Yes | New version |

---

## 3. Rate Limiting & Throttling (ref: DA-BP-008)

### Rate Limiting Standards

| Tier | Limit | Scope | Banking Use Case |
|---|---|---|---|
| Per-user | 100 requests/minute | Authenticated user | Normal banking operations |
| Per-user (sensitive) | 10 requests/minute | Authenticated user on sensitive endpoints | Transfer initiation, beneficiary changes |
| Per-IP (unauthenticated) | 20 requests/minute | IP address | Login, password reset, public endpoints |
| Per-API-key (B2B) | 1,000 requests/minute | API key / client ID | Open banking partners, third-party integrations |
| Global | 10,000 requests/minute | Entire API | System protection |

### Rate Limit Response Headers

```http
HTTP/1.1 200 OK
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 87
X-RateLimit-Reset: 1710510600

# When limit exceeded:
HTTP/1.1 429 Too Many Requests
Retry-After: 30
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 0
X-RateLimit-Reset: 1710510600
Content-Type: application/json

{
  "error": {
    "code": "RATE_LIMIT_EXCEEDED",
    "message": "Too many requests. Please try again later.",
    "retryAfter": 30
  }
}
```

### Throttling Strategies

| Strategy | How It Works | When to Use |
|---|---|---|
| Fixed window | Count requests per fixed time window | Simple, general-purpose |
| Sliding window | Rolling window based on request timestamps | Smoother rate enforcement |
| Token bucket | Tokens replenished at fixed rate; request consumes token | Allow bursts within overall limit |
| Leaky bucket | Requests processed at fixed rate; excess queued | Smooth output rate |

### Banking-Specific Rate Limiting:
- Login attempts: 5 per 5 minutes per account → account lockout
- OTP requests: 3 per 10 minutes per user → temporary block
- Transfer initiation: 10 per minute per user → prevent automated abuse
- Statement download: 5 per hour per user → prevent bulk extraction
- Beneficiary operations: 10 per hour per user → prevent rapid changes
- Rate limits must be enforced server-side — never rely on client-side throttling
- Rate limit events must be logged for security monitoring

---

## 4. Authentication & Authorization (ref: DA-BP-009)

### OAuth 2.0 / OpenID Connect Standards

| Flow | When to Use | Banking Use Case |
|---|---|---|
| Authorization Code + PKCE | Customer-facing web/mobile apps | Mobile banking, web banking |
| Client Credentials | Service-to-service (no user context) | Batch processing, internal microservices |
| Device Authorization | Limited-input devices | ATM, kiosk |
| Refresh Token | Extend session without re-authentication | Mobile app background refresh |

### Token Standards

| Token | Type | Lifetime | Storage | Content |
|---|---|---|---|---|
| Access token | JWT (RS256/ES256) | 15 minutes | In-memory (web), secure enclave (mobile) | sub, scope, iat, exp, jti, iss, aud — no PII |
| Refresh token | Opaque | 24 hours | HTTP-only secure cookie (web), Keychain/Keystore (mobile) | Reference to server-side session |
| ID token | JWT | Single-use (authentication) | In-memory, discard after validation | sub, name, email (if consented) |

### JWT Best Practices

| Practice | Standard |
|---|---|
| Algorithm | RS256 or ES256 — never HS256 for distributed systems, never "none" |
| Validation | Verify signature, issuer (iss), audience (aud), expiry (exp), not-before (nbf) |
| No sensitive data | Never include passwords, account numbers, balances, or PII in JWT payload |
| Token size | Keep small (< 1KB) — use token introspection for additional claims |
| Key rotation | Rotate signing keys every 90 days; publish via JWKS endpoint |
| Revocation | Short lifetime (15 min) + server-side revocation list for immediate invalidation |
| JTI (JWT ID) | Include unique ID for replay detection |

### Authorization Standards

| Level | Implementation | Banking Example |
|---|---|---|
| Scope-based | OAuth 2.0 scopes per API operation | accounts:read, transfers:write, beneficiaries:manage |
| Role-based (RBAC) | User role determines allowed operations | customer, teller, manager, admin |
| Resource-level | Verify user owns/has access to specific resource | Customer can only access own accounts |
| Attribute-based (ABAC) | Dynamic rules based on attributes | Transfer allowed only during business hours for corporate accounts |
| Transaction-level | Additional auth for high-value operations | Step-up MFA for transfers > $10,000 |

### Scope Definitions for Banking APIs

```
# Account scopes
accounts:read          - View account details and balance
accounts:transactions  - View transaction history
accounts:statements    - Download statements

# Transfer scopes
transfers:read         - View transfer status and history
transfers:write        - Initiate transfers
transfers:approve      - Approve pending transfers (maker-checker)

# Beneficiary scopes
beneficiaries:read     - View beneficiary list
beneficiaries:manage   - Add, modify, delete beneficiaries

# Card scopes
cards:read             - View card details
cards:manage           - Activate, block, set limits

# Admin scopes
admin:users            - User management
admin:config           - System configuration
```

### API Security Headers

```http
# Request
Authorization: Bearer eyJhbGciOiJSUzI1NiIs...
X-Request-ID: req-abc-123
X-Correlation-ID: corr-xyz-789
X-MFA-Token: mfa-token-456  (for step-up operations)

# Response
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
Cache-Control: no-store
Pragma: no-cache
```

## 5. API Documentation (ref: DA-BP-010)

### OpenAPI 3.0+ Specification Standards

Every API must be documented using OpenAPI 3.0+ (formerly Swagger). The specification is the single source of truth.

**Required Sections:**

| Section | Content | Standard |
|---|---|---|
| info | API title, description, version, contact, license | Version must match URL path version |
| servers | Base URLs per environment | Production, staging, sandbox |
| paths | All endpoints with operations | Every endpoint documented |
| components/schemas | Request/response models | All models with field descriptions, types, constraints, examples |
| components/securitySchemes | Authentication methods | OAuth 2.0 flows, scopes, token URLs |
| security | Global security requirements | Applied to all endpoints unless overridden |
| tags | Logical grouping of endpoints | Group by domain: Accounts, Transfers, Beneficiaries |

**Per-Endpoint Documentation:**

| Element | Required | Example |
|---|---|---|
| Summary | Yes | "Initiate domestic fund transfer" |
| Description | Yes | Detailed behavior, business rules, prerequisites |
| OperationId | Yes | initiateTransfer (unique, camelCase) |
| Parameters | Yes (if any) | Path, query, header params with type, required, description, example |
| Request body | Yes (if any) | Schema with field descriptions, constraints, examples |
| Responses | Yes (all codes) | 200, 201, 400, 401, 403, 404, 409, 422, 429, 500 with schema and examples |
| Security | Yes | Scopes required for this operation |
| Tags | Yes | Domain grouping |
| Deprecated | If applicable | Boolean flag + description of replacement |

**Schema Documentation Standards:**

```yaml
TransferRequest:
  type: object
  required:
    - sourceAccountId
    - beneficiaryId
    - amount
    - currency
  properties:
    sourceAccountId:
      type: string
      description: Source account identifier. Must belong to the authenticated customer.
      example: "ACC-001234"
      pattern: "^ACC-[0-9]{6}$"
    beneficiaryId:
      type: string
      description: Registered beneficiary identifier. Must be active and verified.
      example: "BEN-005678"
    amount:
      type: number
      format: double
      description: Transfer amount. Must be positive and within daily limit.
      minimum: 0.01
      maximum: 999999999.99
      example: 5000.00
    currency:
      type: string
      description: ISO 4217 currency code.
      enum: [USD, EUR, GBP, CHF, JPY]
      example: "USD"
    reference:
      type: string
      description: Optional payment reference visible to beneficiary.
      maxLength: 140
      pattern: "^[a-zA-Z0-9 \\-/.]*$"
      example: "Invoice 2026-001"
```

### Documentation Deliverables

| Deliverable | Format | Audience | Content |
|---|---|---|---|
| OpenAPI spec | YAML/JSON | Developers | Complete API contract |
| API reference | Generated from OpenAPI (Redoc/Swagger UI) | Developers | Interactive documentation |
| Getting started guide | Markdown | New consumers | Authentication setup, first API call, common patterns |
| Error handling guide | Markdown | Developers | Error codes, meanings, resolution steps |
| Migration guide | Markdown | Existing consumers | Changes between versions, migration steps |
| Changelog | Markdown | All consumers | Version history, breaking/non-breaking changes |
| Postman collection | JSON | Developers | Pre-built requests for testing |
| SDK documentation | Language-specific | Developers | Generated client library docs |

### Documentation Maintenance:
- OpenAPI spec must be version-controlled alongside code
- Spec must be validated on every build (lint with Spectral or similar)
- Documentation must be updated before or with the code change — never after
- Breaking changes must be documented in migration guide before release
- API changelog must be maintained with every release
- Examples must use realistic banking data (with masked values)

---

## 6. Request/Response Validation (ref: DA-BP-011)

### Input Validation Standards

| Layer | What | How | Purpose |
|---|---|---|---|
| Schema validation | Structure, types, required fields | OpenAPI schema validation middleware | Reject malformed requests early |
| Business validation | Business rules, limits, permissions | Application logic | Enforce business constraints |
| Security validation | Injection, encoding, size limits | WAF + application layer | Prevent attacks |

### Validation Rules by Field Type

| Field Type | Validations | Banking Example |
|---|---|---|
| String | Max length, pattern (regex), allowed characters, encoding | Reference: max 140 chars, alphanumeric + -/. only |
| Number | Min, max, precision, format | Amount: min 0.01, max 999,999,999.99, 2 decimal places |
| Date | Format (ISO 8601), range, business day | Transfer date: ISO 8601, not in past, must be business day |
| Enum | Allowed values only | Currency: [USD, EUR, GBP, CHF, JPY] |
| ID/Reference | Format pattern, existence check | Account ID: ^ACC-[0-9]{6}$, must exist and belong to user |
| Email | RFC 5322 format, domain validation | Customer email for notifications |
| Phone | E.164 format | Customer phone for OTP delivery |
| IBAN | Country-specific format + check digit validation | Beneficiary IBAN |
| Amount | Positive, decimal precision, currency-specific rules | JPY has no decimal places; USD has 2 |

### Validation Response Format

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Request validation failed",
    "details": [
      {
        "field": "amount",
        "code": "EXCEEDS_MAXIMUM",
        "message": "Amount must not exceed 999,999,999.99",
        "rejectedValue": 1000000000.00
      },
      {
        "field": "currency",
        "code": "INVALID_VALUE",
        "message": "Currency must be one of: USD, EUR, GBP, CHF, JPY",
        "rejectedValue": "XYZ"
      }
    ]
  }
}
```

### Response Validation Standards

| Standard | Implementation |
|---|---|
| Schema conformance | Response must match documented OpenAPI schema — validate in tests |
| No extra fields | Response must not include undocumented fields (prevents data leakage) |
| Sensitive data masking | Account numbers masked (****1234), no PII in responses unless necessary |
| Null handling | Use explicit null or omit field — document the convention |
| Date format | ISO 8601 with timezone (2026-03-15T15:30:00Z) — always UTC |
| Amount format | Number with explicit decimal precision — never string for amounts |
| Enum values | Only documented values — new values require documentation update |
| Error responses | Consistent structure across all endpoints — same error schema |

### Security Validation

| Attack | Validation | Implementation |
|---|---|---|
| SQL injection | Parameterized queries; reject SQL patterns in input | Never concatenate user input into queries |
| XSS | Output encoding; reject script tags in input | Context-aware encoding on output |
| XXE | Disable external entity processing | Disable DTD processing in XML parsers |
| Path traversal | Reject ../ patterns; whitelist allowed paths | Validate file paths against allowed directories |
| Mass assignment | Whitelist allowed fields per endpoint | DTO pattern — only bind documented fields |
| Request size | Max body size limit | 1MB default; 10MB for file uploads |
| Header injection | Validate header values; reject newlines | Sanitize all header values |
| JSON injection | Strict JSON parsing; reject duplicate keys | Use strict JSON parser |

---

## 7. Correlation IDs & Traceability (ref: DA-BP-012)

### Correlation ID Standards

| Header | Purpose | Format | Who Generates |
|---|---|---|---|
| X-Request-ID | Unique ID for this specific request | UUID v4 | Client (or API gateway if not provided) |
| X-Correlation-ID | ID linking related requests across services | UUID v4 | First service in chain; propagated downstream |

### Implementation Standards

| Standard | Requirement |
|---|---|
| Generation | UUID v4 — cryptographically random |
| Propagation | Every downstream service call must include the correlation ID |
| Logging | Every log entry must include the correlation ID |
| Response | Return both X-Request-ID and X-Correlation-ID in response headers |
| Storage | Store correlation ID with the business entity (e.g., transfer record) |
| Error responses | Include request ID in error response for support reference |

### Traceability Architecture

```
Client → API Gateway → Service A → Service B → Database
  │          │             │            │           │
  │  X-Request-ID: req-1   │            │           │
  │  X-Correlation-ID: corr-1          │           │
  │          │             │            │           │
  │          │    X-Request-ID: req-2   │           │
  │          │    X-Correlation-ID: corr-1          │
  │          │             │            │           │
  │          │             │   X-Request-ID: req-3  │
  │          │             │   X-Correlation-ID: corr-1
  │          │             │            │           │
  All log entries include: correlation_id=corr-1
```

### Distributed Tracing Integration

| Standard | Tool | Implementation |
|---|---|---|
| OpenTelemetry | Jaeger, Zipkin, Datadog | Instrument all services with OpenTelemetry SDK |
| W3C Trace Context | traceparent, tracestate headers | Propagate trace context across all service boundaries |
| Span naming | {service}.{operation} | transfer-service.initiateTransfer |
| Span attributes | Include business context | transfer.amount, transfer.currency, transfer.status |

---

## 8. Error Handling Standards

### Error Code Taxonomy

| Category | Code Pattern | Examples |
|---|---|---|
| Validation | VALIDATION_* | VALIDATION_ERROR, VALIDATION_FIELD_REQUIRED, VALIDATION_FORMAT_INVALID |
| Authentication | AUTH_* | AUTH_TOKEN_EXPIRED, AUTH_TOKEN_INVALID, AUTH_MFA_REQUIRED |
| Authorization | AUTHZ_* | AUTHZ_INSUFFICIENT_SCOPE, AUTHZ_RESOURCE_FORBIDDEN |
| Business | BIZ_* | BIZ_INSUFFICIENT_FUNDS, BIZ_LIMIT_EXCEEDED, BIZ_ACCOUNT_FROZEN |
| System | SYS_* | SYS_INTERNAL_ERROR, SYS_SERVICE_UNAVAILABLE, SYS_TIMEOUT |
| Rate limiting | RATE_* | RATE_LIMIT_EXCEEDED |

### Error Response Standards

| Principle | Standard |
|---|---|
| Consistent structure | Same error schema across all endpoints and versions |
| Machine-readable code | Error code for programmatic handling (VALIDATION_ERROR) |
| Human-readable message | Clear message for developers/users |
| Field-level details | For validation errors, identify which field(s) failed |
| No sensitive data | Never include stack traces, SQL, internal paths, server details |
| Request ID | Include for support correlation |
| Idempotent errors | Same request produces same error response |
| Localization | Error messages localizable via Accept-Language header |

---

## 9. Idempotency

### Idempotency Standards

| Method | Naturally Idempotent | Idempotency Key Required |
|---|---|---|
| GET | Yes | No |
| PUT | Yes | No |
| DELETE | Yes | No |
| POST | No | Yes — for create/action operations |
| PATCH | Depends | Recommended for financial operations |

### Implementation

```http
# Client sends:
POST /v1/transfers
Idempotency-Key: idem-key-abc-123
Content-Type: application/json

{
  "sourceAccountId": "ACC-001234",
  "beneficiaryId": "BEN-005678",
  "amount": 5000.00,
  "currency": "USD"
}

# Server behavior:
# 1. Check if Idempotency-Key exists in store
# 2. If exists: return stored response (same status code and body)
# 3. If not: process request, store response with key, return response
# 4. Key expires after 24 hours
```

**Banking Idempotency Rules:**
- All POST endpoints that create financial transactions must support Idempotency-Key
- Idempotency key must be client-generated UUID
- Server must store: key, response status, response body, timestamp
- Duplicate detection window: 24 hours
- If processing is in-flight for the same key: return 409 Conflict with status URL
- Idempotency key collisions with different request bodies: return 422 Unprocessable Entity

---

## 10. API Performance & Resilience

### Performance Standards

| Metric | Target | Measurement |
|---|---|---|
| Response time (p95) | < 500ms (customer-facing), < 200ms (internal) | APM / load testing |
| Response time (p99) | < 1000ms (customer-facing), < 500ms (internal) | APM |
| Throughput | Per SLA (typically 500-5000 TPS) | Load testing |
| Error rate | < 0.1% | Monitoring |
| Availability | 99.95%+ | Uptime monitoring |

### Resilience Patterns

| Pattern | Implementation | Banking Use Case |
|---|---|---|
| Circuit breaker | Open after N failures; half-open to test recovery | Core banking integration — prevent cascade failure |
| Retry with backoff | Exponential backoff with jitter; max 3 retries | Payment network timeout — retry with increasing delay |
| Timeout | Connect timeout: 5s; read timeout: 30s | Prevent indefinite waits on downstream services |
| Bulkhead | Isolate thread pools per downstream service | Fraud engine failure doesn't block account inquiry |
| Fallback | Graceful degradation when dependency unavailable | Show cached balance when core banking is slow |
| Rate limiting | Protect downstream services from overload | Limit requests to core banking API |

### Caching Standards

| Data Type | Cacheable | TTL | Cache-Control |
|---|---|---|---|
| Account balance | No (real-time) | — | no-store |
| Transaction history | Short-term | 30 seconds | private, max-age=30 |
| Exchange rates | Yes | 60 seconds | public, max-age=60 |
| Product catalog | Yes | 1 hour | public, max-age=3600 |
| Customer profile | No (PII) | — | no-store, private |
| Static content | Yes | 1 year | public, max-age=31536000, immutable |

---

## 11. API Lifecycle Management

### API Lifecycle Stages

| Stage | Activities | Duration |
|---|---|---|
| Design | OpenAPI spec, review, approval | 1-2 weeks |
| Development | Implementation, unit tests, integration tests | Per sprint |
| Testing | Security testing, performance testing, contract testing | 1-2 weeks |
| Beta/Sandbox | Partner testing, feedback collection | 2-4 weeks |
| Production | GA release, monitoring, support | Ongoing |
| Deprecation | Sunset notice, migration support | 6 months minimum |
| Retirement | Endpoint removed, documentation archived | After sunset date |

### Contract Testing

| Test Type | Tool | Purpose |
|---|---|---|
| Provider contract test | Pact (provider side) | Verify API matches documented contract |
| Consumer contract test | Pact (consumer side) | Verify consumer expectations match API |
| Schema validation test | OpenAPI validator | Verify responses match OpenAPI schema |
| Breaking change detection | OpenAPI diff (oasdiff) | Detect breaking changes before release |

---

## 12. API Gateway Standards

### Gateway Responsibilities

| Function | Implementation |
|---|---|
| Authentication | Validate JWT, check token expiry, verify signature |
| Rate limiting | Enforce per-user, per-IP, per-API-key limits |
| Request routing | Route to correct service version |
| TLS termination | Terminate TLS at gateway; mTLS to backend services |
| Request/response transformation | Add correlation IDs, strip internal headers |
| Logging | Log all requests with correlation ID, user, endpoint, status, latency |
| CORS | Enforce CORS policy (restrictive origins, no wildcard in production) |
| Request size limiting | Reject oversized requests before reaching services |
| IP allowlisting | Restrict access for admin/internal APIs |

### CORS Configuration

```
Access-Control-Allow-Origin: https://banking.example.com  (specific origin, never *)
Access-Control-Allow-Methods: GET, POST, PUT, PATCH, DELETE, OPTIONS
Access-Control-Allow-Headers: Authorization, Content-Type, X-Request-ID, X-Correlation-ID, X-MFA-Token, Idempotency-Key
Access-Control-Expose-Headers: X-Request-ID, X-Correlation-ID, X-RateLimit-Limit, X-RateLimit-Remaining
Access-Control-Max-Age: 3600
Access-Control-Allow-Credentials: true
```

---

## 13. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Verbs in URLs (/createTransfer) | Not RESTful; inconsistent | Use nouns + HTTP methods |
| Single version forever | Can't evolve without breaking consumers | URL-based versioning from day one |
| No rate limiting | Abuse, DoS, data scraping | Rate limit all endpoints |
| API keys as sole auth | Insufficient for banking | OAuth 2.0 + PKCE for user context |
| Undocumented APIs | Consumer confusion, integration errors | OpenAPI 3.0+ spec for every API |
| No input validation | Injection attacks, data corruption | Validate at schema + business + security layers |
| No correlation IDs | Can't trace issues across services | Generate and propagate on every request |
| Sensitive data in GET params | Logged in server logs, browser history | POST for sensitive operations |
| No idempotency for POST | Duplicate transactions on retry | Idempotency-Key for all financial POSTs |
| Inconsistent error format | Consumer can't handle errors reliably | Single error schema across all APIs |
| No pagination | Memory issues, data extraction risk | Paginate all collection endpoints |
| Wildcard CORS | Security vulnerability | Specific origins only |
