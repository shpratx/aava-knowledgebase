# Guardrails for API Design and Architecture

---

## 1. API Authentication Guardrails (ref: DA-GR-006)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-AG-001 | All APIs must require authentication — no anonymous access to sensitive data (ref: DA-GR-006) | Reject if any endpoint serving sensitive data has no authentication |
| API-AG-002 | Authentication must use OAuth 2.0 / OpenID Connect — custom auth schemes are prohibited | Reject if custom authentication is proposed |
| API-AG-003 | Access tokens must be JWT signed with RS256 or ES256 — HS256 and "none" algorithm are prohibited | Reject if weak or no signing algorithm is used |
| API-AG-004 | Access token lifetime must not exceed 15 minutes | Reject if token lifetime exceeds 15 minutes |
| API-AG-005 | Refresh tokens must be opaque (not JWT), single-use, and rotated on each use | Reject if refresh tokens are JWT or reusable |
| API-AG-006 | Token validation must verify: signature, issuer (iss), audience (aud), expiry (exp), and not-before (nbf) | Reject if any validation step is missing |
| API-AG-007 | Tokens must not contain sensitive data (passwords, account numbers, balances, PII) in the payload | Reject if sensitive data is found in token payload |
| API-AG-008 | Service-to-service authentication must use mTLS or client credentials grant — not shared API keys | Reject if internal services use shared API keys |
| API-AG-009 | Authentication failures must return 401 with no information leakage — no differentiation between invalid, expired, or missing token | Reject if auth error reveals failure reason |
| API-AG-010 | All authentication events (success, failure, token refresh, revocation) must be logged with correlation ID | Reject if auth events are not logged |
| API-AG-011 | Public endpoints (health check, OpenAPI spec) must be explicitly whitelisted — default is authenticated | Reject if unauthenticated endpoints are not explicitly approved |

---

## 2. API Key & Secret Management Guardrails (ref: DA-GR-007)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-KG-001 | API keys, secrets, tokens, and credentials must not be hardcoded in source code (ref: DA-GR-007) | Reject — CI/CD secret scanning must block |
| API-KG-002 | API keys must not be committed to version control — pre-commit hooks must scan for secrets | Reject if secrets detected in repository |
| API-KG-003 | API keys must not appear in URLs, query parameters, or request logs | Reject if keys are transmitted via URL or logged |
| API-KG-004 | API keys must not be exposed in client-side code (JavaScript bundles, mobile app binaries) | Reject if keys are embedded in client artifacts |
| API-KG-005 | Secrets must be stored in secure vaults (HashiCorp Vault, AWS Secrets Manager) — not in env vars, config files, or images | Reject if secrets are stored outside secure vault |
| API-KG-006 | API keys must be rotated at minimum every 90 days | Flag if key rotation exceeds 90 days |
| API-KG-007 | Compromised keys must be revoked immediately — incident response process must be defined | Reject if no key compromise process exists |
| API-KG-008 | Each environment must use separate keys — no key sharing across environments | Reject if same key is used across environments |
| API-KG-009 | Each consuming application must have its own API key — no shared keys between applications | Reject if multiple applications share a key |
| API-KG-010 | API key usage must be monitored — alert on unusual patterns (volume spike, new IP, off-hours) | Flag if no key usage monitoring exists |

---

## 3. Input Validation Guardrails (ref: DA-GR-008)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-VG-001 | All API input parameters must be validated server-side using whitelist approach (ref: DA-GR-008) | Reject if any parameter lacks server-side validation |
| API-VG-002 | Validation must reject unknown/unexpected fields — no mass assignment | Reject if API accepts undocumented fields |
| API-VG-003 | String inputs must have maximum length enforced | Reject if string fields have no max length |
| API-VG-004 | String inputs must be validated against allowed character patterns (regex whitelist) | Reject if string fields accept arbitrary characters |
| API-VG-005 | Numeric inputs must have min/max range enforced | Reject if numeric fields have no range validation |
| API-VG-006 | Enum inputs must only accept documented values | Reject if enum fields accept undocumented values |
| API-VG-007 | Date inputs must be validated for format (ISO 8601), range, and business rules | Reject if date fields have no validation |
| API-VG-008 | Request body size must be limited — default 1MB; 10MB for file uploads | Reject if no request size limit is configured |
| API-VG-009 | Content-Type header must be validated — reject unexpected content types | Reject if Content-Type validation is missing |
| API-VG-010 | Array inputs must have maximum item count enforced | Reject if array fields have no max count |
| API-VG-011 | Nested object depth must be limited | Flag if no nesting depth limit exists |
| API-VG-012 | Validation errors must return 400 with field-level details | Reject if validation errors are generic |
| API-VG-013 | Validation must occur before any business logic or database access | Reject if validation happens after business processing |

---

## 4. Output Encoding Guardrails (ref: DA-GR-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-OG-001 | All API responses must use context-aware output encoding (ref: DA-GR-009) | Reject if output encoding is not implemented |
| API-OG-002 | JSON responses must properly escape special characters | Reject if JSON serialization does not escape |
| API-OG-003 | Responses must set Content-Type: application/json with charset=utf-8 | Reject if Content-Type is missing or incorrect |
| API-OG-004 | X-Content-Type-Options: nosniff must be set on all responses | Reject if header is missing |
| API-OG-005 | API responses must not include user-supplied input without encoding | Reject if user input is reflected unencoded |
| API-OG-006 | Error messages must not include raw user input | Reject if errors echo raw input |
| API-OG-007 | SQL errors, stack traces, and internal paths must never appear in responses | Reject if technical details leak |
| API-OG-008 | XML responses must disable external entity processing (XXE prevention) | Reject if XML allows external entities |

---

## 5. Resource-Level Authorization Guardrails (ref: DA-GR-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-RG-001 | APIs must enforce authorization at resource level — role checks alone are insufficient (ref: DA-GR-010) | Reject if authorization is role-only |
| API-RG-002 | Every endpoint accessing customer data must verify the authenticated user owns that resource | Reject if resource ownership check is missing |
| API-RG-003 | IDOR prevention must be implemented and tested | Reject if IDOR testing is not in security tests |
| API-RG-004 | Authorization must be enforced on every request — not cached | Reject if authorization is cached across requests |
| API-RG-005 | Each endpoint must declare required OAuth scopes | Reject if endpoint has no scope requirement |
| API-RG-006 | Maker-checker must be enforced for high-value operations | Reject if high-value ops allow single-user completion |
| API-RG-007 | Horizontal privilege escalation must be prevented | Reject if horizontal escalation is possible |
| API-RG-008 | Vertical privilege escalation must be prevented | Reject if vertical escalation is possible |
| API-RG-009 | Authorization failures must return 403 with no resource details | Reject if 403 reveals resource information |
| API-RG-010 | All authorization decisions must be logged | Reject if authorization events are not logged |

---

## 6. Sensitive Data Masking Guardrails (ref: DA-GR-011)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-DG-001 | Sensitive data in API responses must be masked or tokenized (ref: DA-GR-011) | Reject if sensitive data is returned unmasked |
| API-DG-002 | Account numbers: show last 4 digits only (****1234) | Reject if full account numbers returned by default |
| API-DG-003 | Card numbers (PAN): first 6 + last 4 only (411111****1234) per PCI-DSS | Reject if full PAN is returned |
| API-DG-004 | CVV, PIN, track data: never returned in any response | Reject — no exceptions |
| API-DG-005 | National ID / tax ID: last 4 digits only | Reject if full ID is returned |
| API-DG-006 | Email: partially masked in list responses (j***@example.com) | Flag if fully visible in lists |
| API-DG-007 | Phone: partially masked in list responses (***-***-1234) | Flag if fully visible in lists |
| API-DG-008 | Passwords, hashes, auth secrets: never returned | Reject — no exceptions |
| API-DG-009 | Internal system IDs (database IDs): not exposed — use external identifiers | Reject if internal IDs are exposed |
| API-DG-010 | Responses must not include more data than requested — support field filtering | Flag if responses include unnecessary fields |
| API-DG-011 | Sensitive data must not appear in API logs | Reject if logs contain unmasked sensitive data |
| API-DG-012 | Error responses must not contain sensitive data | Reject if errors contain sensitive data |

---

## 7. Protocol Security Guardrails (ref: DA-GR-012)

| ID | Guardrail | Enforcement |
|---|---|---|
| API-PG-001 | All APIs must use HTTPS — HTTP is prohibited (ref: DA-GR-012) | Reject if any endpoint is accessible via HTTP |
| API-PG-002 | TLS 1.2 minimum; TLS 1.3 preferred — SSL, TLS 1.0, TLS 1.1 prohibited | Reject if protocols below TLS 1.2 are enabled |
| API-PG-003 | Weak cipher suites must be disabled (RC4, DES, 3DES, MD5, export-grade) | Reject if weak ciphers are enabled |
| API-PG-004 | HSTS: max-age ≥ 31536000; includeSubDomains; preload | Reject if HSTS is missing or misconfigured |
| API-PG-005 | Certificates: valid, trusted CA, ≥ 2048-bit RSA or ECC P-256 | Reject if cert is self-signed, expired, or weak |
| API-PG-006 | Certificate expiry must be monitored — alert 30 days before | Flag if no certificate monitoring exists |
| API-PG-007 | Internal service-to-service must use mTLS | Reject if internal services lack mTLS |
| API-PG-008 | WebSocket must use WSS — WS is prohibited | Reject if unsecured WebSocket is used |
| API-PG-009 | Response headers must not reveal server technology | Reject if technology-revealing headers present |
| API-PG-010 | TLS configuration must be verified with SSL Labs / testssl.sh | Reject if TLS is not verified |

---

## 8. API Versioning & Lifecycle Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| API-VLG-001 | All APIs must be versioned in URL path (/v1/) — unversioned APIs prohibited | Reject if no version in URL |
| API-VLG-002 | Breaking changes must create new major version | Reject if breaking change modifies existing version |
| API-VLG-003 | Deprecated versions must include Deprecation and Sunset headers | Reject if deprecated version lacks headers |
| API-VLG-004 | Minimum 6 months notice before version retirement | Reject if notice is less than 6 months |
| API-VLG-005 | Maximum 2 concurrent major versions | Flag if more than 2 versions active |

---

## 9. Rate Limiting & Abuse Prevention Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| API-RLG-001 | All endpoints must have rate limiting — no unlimited access | Reject if any endpoint has no rate limit |
| API-RLG-002 | Rate limit responses: 429 with Retry-After header | Reject if 429 or Retry-After is missing |
| API-RLG-003 | Sensitive endpoints must have stricter limits | Reject if sensitive endpoints have general limits |
| API-RLG-004 | Rate limiting must be server-side | Reject if client-side only |
| API-RLG-005 | Rate limit events must be logged | Reject if not logged |
| API-RLG-006 | Rate limit headers must be in responses | Flag if headers missing |

---

## 10. Documentation Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| API-DCG-001 | Every API must have OpenAPI 3.0+ spec — undocumented APIs must not deploy | Reject if no spec |
| API-DCG-002 | Spec must be validated on every build | CI/CD gate |
| API-DCG-003 | All response codes must be documented with schemas and examples | Reject if undocumented |
| API-DCG-004 | All fields must have description, type, constraints, example | Flag if incomplete |
| API-DCG-005 | Security requirements documented per endpoint | Reject if undocumented |
| API-DCG-006 | Breaking changes detected automatically before release | CI/CD gate |

---

## 11. Correlation & Traceability Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| API-TG-001 | Every request must have correlation ID (X-Correlation-ID) | Reject if not propagated |
| API-TG-002 | Correlation ID propagated to all downstream calls | Reject if dropped |
| API-TG-003 | Correlation ID in every log entry | Reject if missing |
| API-TG-004 | Correlation ID returned in response headers | Reject if omitted |
| API-TG-005 | Correlation ID stored with business entities | Flag if not stored |

---

## 12. Idempotency Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| API-IG-001 | Financial POST endpoints must support Idempotency-Key | Reject if no idempotency |
| API-IG-002 | Duplicate key must return original response — no duplicate transactions | Reject if duplicates created |
| API-IG-003 | Same key + different body must return 422 | Reject if collision not detected |
| API-IG-004 | Keys expire after 24 hours | Flag if expiry exceeds 24 hours |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot proceed | No auth (DA-GR-006), hardcoded keys (DA-GR-007), no validation (DA-GR-008), no encoding (DA-GR-009), no resource auth (DA-GR-010), unmasked PAN (DA-GR-011), TLS < 1.2 (DA-GR-012) |
| **Flag** | Address before release | Key rotation > 90 days, missing rate limit headers, email visible in lists |
| **CI/CD Gate** | Automated | Secret scanning, OpenAPI lint, breaking change detection, TLS verification |

---

## Quick Reference: Guardrail Triggers by API Type

| API Type | Key Guardrails |
|---|---|
| Customer-facing | All sections — full enforcement |
| Open Banking | All + stricter rate limits + per-partner keys |
| Internal service-to-service | mTLS (API-PG-007), resource auth (API-RG-001), validation (API-VG-001), correlation (API-TG-001) |
| Admin/operations | Maker-checker (API-RG-006), vertical escalation prevention (API-RG-008), no internal IDs (API-DG-009) |
| Public | Explicit whitelist (API-AG-011), HTTPS (API-PG-001), rate limiting (API-RLG-001), encoding (API-OG-001) |

---

## Pre-Deployment Checklist

| # | Check | Ref |
|---|---|---|
| 1 | All endpoints authenticated (except whitelisted) | DA-GR-006 |
| 2 | No hardcoded keys in code/config/client | DA-GR-007 |
| 3 | All inputs validated server-side (whitelist) | DA-GR-008 |
| 4 | All outputs encoded; no raw input in responses | DA-GR-009 |
| 5 | Resource-level authorization on every endpoint | DA-GR-010 |
| 6 | Sensitive data masked in responses and logs | DA-GR-011 |
| 7 | HTTPS only; TLS 1.2+; HSTS; no weak ciphers | DA-GR-012 |
| 8 | OpenAPI spec complete and validated | API-DCG-001 |
| 9 | Rate limiting on all endpoints | API-RLG-001 |
| 10 | Correlation IDs propagated and logged | API-TG-001 |
| 11 | Idempotency for financial POSTs | API-IG-001 |
| 12 | API versioned in URL | API-VLG-001 |
| 13 | Secret scanning in CI/CD | API-KG-001 |
| 14 | TLS verified (SSL Labs) | API-PG-010 |
