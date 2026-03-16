# Guardrails for Spring Boot Application Development
---

## 1. Endpoint Security Guardrails (ref: DV-GR-009)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-SG-001 | All endpoints must have explicit security annotations — no endpoint without auth config (ref: DV-GR-009) | Reject if endpoint lacks security rule |
| SB-SG-002 | Default security policy must be deny-all — .anyRequest().authenticated() | Reject if default is permitAll |
| SB-SG-003 | Public endpoints must be explicitly whitelisted with documented justification | Reject if public endpoint not documented |
| SB-SG-004 | Every endpoint must declare required OAuth scopes | Reject if scope missing |
| SB-SG-005 | Resource-level authorization must be enforced — role check alone insufficient | Reject if role-only auth |
| SB-SG-006 | Step-up MFA required for financial transactions and profile changes | Reject if financial endpoint has no MFA |
| SB-SG-007 | Security annotations must be tested — auth bypass and privilege escalation tests required | Reject if security tests missing |

---

## 2. Actuator Security Guardrails (ref: DV-GR-010)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-AG-001 | Only health and info actuator endpoints are public (ref: DV-GR-010) | Reject if sensitive actuator endpoints publicly accessible |
| SB-AG-002 | /actuator/env, /actuator/configprops, /actuator/beans must be disabled in production | Reject if enabled in production |
| SB-AG-003 | /actuator/shutdown must be disabled in all environments | Reject if shutdown enabled |
| SB-AG-004 | Actuator should be on separate management port in production | Flag if actuator shares application port |
| SB-AG-005 | Actuator access must be logged and monitored | Flag if not logged |

---

## 3. Secrets Management Guardrails (ref: DV-GR-011)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-SM-001 | Secrets must be externalized via Vault or Secrets Manager (ref: DV-GR-011) | Reject if secrets in application.yml or code |
| SB-SM-002 | No secrets in source code — CI/CD secret scanning must block | Reject immediately |
| SB-SM-003 | No secrets in version control — pre-commit hooks must scan | Reject if secrets in repo |
| SB-SM-004 | application.yml must use vault/SM placeholders for sensitive values | Reject if plain-text secrets in config |
| SB-SM-005 | Each environment must use separate secrets | Reject if shared across environments |
| SB-SM-006 | Secrets must be rotated every 90 days minimum | Flag if rotation exceeds 90 days |
| SB-SM-007 | Database connection strings must not contain plain-text passwords | Reject if datasource password is plain text |

---

## 4. SQL Injection Prevention Guardrails (ref: DV-GR-012)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-QG-001 | All queries must use JPA/Hibernate parameterized queries — no string concatenation (ref: DV-GR-012) | Reject if concatenation found |
| SB-QG-002 | Native SQL queries require security review | Reject if native query has no security review |
| SB-QG-003 | Dynamic queries must use Criteria API or QueryDSL — never string assembly | Reject if string-built queries |
| SB-QG-004 | Stored procedures must use typed parameters | Reject if dynamic SQL in procedures |
| SB-QG-005 | Spring Data repository methods preferred (auto-parameterized) | Flag if bypassed without justification |
| SB-QG-006 | JPQL must use named parameters (:param) not positional (?) | Flag if positional params |
| SB-QG-007 | SQL errors must never propagate to API responses | Reject if SQL errors reach client |

---

## 5. Input Validation Guardrails (ref: DV-GR-013)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-VG-001 | All request DTOs must use Bean Validation annotations (ref: DV-GR-013) | Reject if DTO has no validation |
| SB-VG-002 | @Valid must be on all controller parameters accepting request bodies | Reject if @Valid missing |
| SB-VG-003 | String fields must have @Size with maxLength | Reject if unbounded strings |
| SB-VG-004 | Numeric fields must have @DecimalMin/@DecimalMax or @Min/@Max | Reject if no range validation |
| SB-VG-005 | @Pattern for structured input (IDs, references) | Flag if structured strings lack pattern |
| SB-VG-006 | Custom validators for complex business rules | Flag if complex validation in controller |
| SB-VG-007 | Validation errors must return 400 with field-level details | Reject if validation returns 500 |

---

## 6. CORS Configuration Guardrails (ref: DV-GR-014)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-CG-001 | No wildcard (*) for allowed origins in production (ref: DV-GR-014) | Reject if wildcard in production |
| SB-CG-002 | Allowed origins must be explicitly listed | Reject if not explicitly configured |
| SB-CG-003 | Allowed methods restricted to required HTTP methods only | Flag if all methods allowed |
| SB-CG-004 | Allowed headers must be explicitly listed | Flag if wildcard headers |
| SB-CG-005 | CORS must be environment-specific — production must be strict | Reject if production uses dev config |
| SB-CG-006 | Credentials only with specific origins — never with wildcard | Reject if credentials + wildcard |

---

## 7. Error Response Guardrails (ref: DV-GR-015)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-EG-001 | No stack traces in error responses — server.error.include-stacktrace: never (ref: DV-GR-015) | Reject if stack traces in responses |
| SB-EG-002 | No internal class names, package paths, or SQL in responses | Reject if internal details leak |
| SB-EG-003 | No server technology details in responses | Reject if technology details exposed |
| SB-EG-004 | server.error.include-message: never in production | Reject if default messages exposed |
| SB-EG-005 | GlobalExceptionHandler must catch all exceptions | Reject if whitelabel error page reachable |
| SB-EG-006 | Consistent ErrorResponse schema across all endpoints | Reject if inconsistent format |
| SB-EG-007 | Unhandled exceptions logged server-side; generic message to client | Reject if server logging missing for 500s |

---

## 8. API Versioning Guardrails (ref: DV-GR-016)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-AVG-001 | All APIs versioned in URL path (/v1/) (ref: DV-GR-016) | Reject if no version |
| SB-AVG-002 | Breaking changes must create new major version | Reject if existing version modified |
| SB-AVG-003 | Deprecated endpoints must return Deprecation + Sunset headers | Reject if headers missing |
| SB-AVG-004 | Maximum 2 concurrent major versions | Flag if > 2 active |
| SB-AVG-005 | Version in @RequestMapping at controller class level | Flag if per-method |

---

## 9. Integration TLS Guardrails (ref: DV-GR-017)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-TG-001 | All external integrations must use HTTPS with TLS 1.2+ (ref: DV-GR-017) | Reject if HTTP or TLS < 1.2 |
| SB-TG-002 | Internal service-to-service must use mTLS | Reject if no mTLS for internal |
| SB-TG-003 | No trust-all or skip-verification in production | Reject if cert validation disabled |
| SB-TG-004 | HTTP clients must be configured with TLS | Reject if plain HTTP allowed |
| SB-TG-005 | Kafka must use SSL/SASL — no PLAINTEXT in production | Reject if Kafka plaintext |
| SB-TG-006 | Database connections must use SSL (sslmode=verify-full) | Reject if DB unencrypted |

---

## 10. Key/Token Rotation Guardrails (ref: DV-GR-018)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-RG-001 | API keys/tokens rotated max every 90 days (ref: DV-GR-018) | Flag if > 90 days |
| SB-RG-002 | JWT signing keys rotated every 90 days via JWKS | Flag if not configured |
| SB-RG-003 | Database credentials rotated every 90 days | Flag if > 90 days |
| SB-RG-004 | Kafka SASL credentials rotated every 90 days | Flag if > 90 days |
| SB-RG-005 | Compromised keys revoked immediately — process defined | Reject if no process |

---

## 11. Third-Party Integration Guardrails (ref: DV-GR-019)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-TPG-001 | Vendor security assessment required before implementation (ref: DV-GR-019) | Reject if assessment not done |
| SB-TPG-002 | Third-party API contracts must be documented | Reject if undocumented |
| SB-TPG-003 | Third-party dependencies scanned for CVEs | CI/CD gate |
| SB-TPG-004 | Third-party SLAs documented and monitored | Flag if SLA undocumented |
| SB-TPG-005 | Fallback behavior defined for every third-party dependency | Reject if no fallback |

---

## 12. Integration Error Guardrails (ref: DV-GR-020)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-IEG-001 | Integration failures must not expose downstream details (ref: DV-GR-020) | Reject if downstream details leak |
| SB-IEG-002 | Downstream errors caught, logged, translated to upstream error | Reject if raw errors propagate |
| SB-IEG-003 | Downstream service names must not appear in error messages | Reject if service names in errors |
| SB-IEG-004 | Integration error logging must include correlation ID, service, status, duration | Reject if logging incomplete |

---

## 13. Timeout Guardrails (ref: DV-GR-021)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-TOG-001 | Every external call must have explicit timeout — no infinite waits (ref: DV-GR-021) | Reject if no timeout |
| SB-TOG-002 | HTTP clients must have connect and read timeouts | Reject if no timeout |
| SB-TOG-003 | DB connection pool must have connection-timeout (max 5s) | Reject if missing or > 5s |
| SB-TOG-004 | Kafka producer must have delivery.timeout.ms | Reject if no delivery timeout |
| SB-TOG-005 | gRPC clients must have deadline | Reject if no deadline |
| SB-TOG-006 | Timeout values documented in integration contract | Flag if undocumented |

---

## 14. Circuit Breaker Guardrails (ref: DV-GR-022)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-CBG-001 | Circuit breakers required for all external dependencies (ref: DV-GR-022) | Reject if no circuit breaker |
| SB-CBG-002 | Config must include: sliding window, failure threshold, wait duration, half-open calls | Reject if config incomplete |
| SB-CBG-003 | Every circuit breaker must have a fallback method | Reject if no fallback |
| SB-CBG-004 | State changes must be logged and alerted | Reject if not monitored |
| SB-CBG-005 | Metrics exposed via Actuator/Prometheus | Flag if not exposed |

---

## 15. Dead Letter Queue Guardrails (ref: DV-GR-023)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-DLG-001 | Every message consumer must have a DLQ (ref: DV-GR-023) | Reject if no DLQ |
| SB-DLG-002 | 3 retries with exponential backoff before DLQ | Reject if no retry before DLQ |
| SB-DLG-003 | DLQ messages must retain original headers | Reject if metadata lost |
| SB-DLG-004 | DLQ depth monitored — alert immediately | Reject if no monitoring |
| SB-DLG-005 | DLQ investigation SLA: 1h financial, 4h others | Process guardrail |
| SB-DLG-006 | DLQ replay capability must exist | Flag if no replay |
| SB-DLG-007 | DLQ retention minimum 30 days | Reject if < 30 days |

---

## 16. Integration Monitoring Guardrails (ref: DV-GR-024)

| ID | Guardrail | Enforcement |
|---|---|---|
| SB-MG-001 | All integrations must have monitoring and alerting (ref: DV-GR-024) | Reject if no monitoring |
| SB-MG-002 | Monitor per integration: p95 response time, error rate, throughput, availability | Reject if key metrics missing |
| SB-MG-003 | Alert thresholds: p95 > SLA, error rate > 1%, circuit breaker open, DLQ > 0 | Reject if thresholds not configured |
| SB-MG-004 | Distributed tracing enabled across all calls (OpenTelemetry) | Reject if tracing not configured |
| SB-MG-005 | Health checks must include all external dependencies | Reject if health check incomplete |
| SB-MG-006 | Dashboards for each service: RED metrics, dependency health, business metrics | Flag if no dashboard |

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot merge/deploy | No security annotation (009), actuator exposed (010), secrets in code (011), SQL concatenation (012), stack trace in response (015), wildcard CORS (014), no timeout (021), no circuit breaker (022), no DLQ (023) |
| **Flag** | Address before release | Key rotation > 90 days (018), no dashboard (024), positional JPQL params, actuator on app port |
| **CI/CD Gate** | Automated | Secret scanning (011), dependency CVE scan (019), SAST for SQL injection (012) |
| **Process** | Workflow | Vendor assessment (019), DLQ investigation SLA (023), key rotation schedule (018) |

---

## Quick Reference: Guardrail Triggers

| Activity | Triggered Guardrails |
|---|---|
| New REST endpoint | SB-SG (security), SB-VG (validation), SB-EG (errors), SB-AVG (versioning), SB-CG (CORS) |
| New external service call | SB-TG (TLS), SB-TOG (timeout), SB-CBG (circuit breaker), SB-MG (monitoring) |
| New Kafka consumer | SB-DLG (DLQ), SB-TG-005 (SSL), SB-MG (monitoring) |
| New third-party integration | SB-TPG (vendor), SB-TG (TLS), SB-TOG (timeout), SB-CBG (circuit breaker) |
| Configuration change | SB-SM (secrets), SB-CG (CORS), SB-AG (actuator) |
| Database query | SB-QG (SQL injection) |

---

## Pre-Deployment Checklist

| # | Check | Ref |
|---|---|---|
| 1 | All endpoints have explicit security annotations | DV-GR-009 |
| 2 | Actuator secured; sensitive endpoints disabled | DV-GR-010 |
| 3 | All secrets externalized via vault | DV-GR-011 |
| 4 | All queries parameterized; no SQL concatenation | DV-GR-012 |
| 5 | All DTOs validated with Bean Validation | DV-GR-013 |
| 6 | CORS restrictive; no wildcard in production | DV-GR-014 |
| 7 | No stack traces in error responses | DV-GR-015 |
| 8 | API versioned in URL path | DV-GR-016 |
| 9 | All integrations use TLS 1.2+; mTLS internal | DV-GR-017 |
| 10 | Key/token rotation configured (max 90 days) | DV-GR-018 |
| 11 | Third-party vendor assessments completed | DV-GR-019 |
| 12 | Integration errors don't expose downstream details | DV-GR-020 |
| 13 | All external calls have explicit timeouts | DV-GR-021 |
| 14 | Circuit breakers on all external dependencies | DV-GR-022 |
| 15 | DLQ configured for all message consumers | DV-GR-023 |
| 16 | Monitoring and alerting for all integrations | DV-GR-024 |
