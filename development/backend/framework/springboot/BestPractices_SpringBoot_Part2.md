# Spring Boot Application Standards & Best Practices — Part 2
### Security, Resilience, Integration & Testing
### Banking Domain — Agentic Knowledge Base
### References: DV-BP-014, DV-BP-017, DV-BP-018, DV-BP-019, DV-BP-020, DV-BP-021, DV-BP-022

---

## 11. Circuit Breakers & Resilience (ref: DV-BP-014)

```java
// Resilience4j configuration for external dependencies
@Configuration
public class ResilienceConfig {

    @Bean
    public CircuitBreakerConfig coreBankingCircuitBreaker() {
        return CircuitBreakerConfig.custom()
            .slidingWindowSize(10)
            .failureRateThreshold(50)
            .waitDurationInOpenState(Duration.ofSeconds(30))
            .permittedNumberOfCallsInHalfOpenState(3)
            .slowCallDurationThreshold(Duration.ofSeconds(2))
            .slowCallRateThreshold(80)
            .build();
    }
}

// Usage in service client
@Component
public class AccountServiceClient implements AccountPort {

    @CircuitBreaker(name = "account-service", fallbackMethod = "getBalanceFallback")
    @Retry(name = "account-service")
    @TimeLimiter(name = "account-service")
    public CompletableFuture<Balance> getBalance(String accountId) {
        return CompletableFuture.supplyAsync(() ->
            restClient.get()
                .uri("/v1/accounts/{id}/balance", accountId)
                .retrieve()
                .body(Balance.class));
    }

    private CompletableFuture<Balance> getBalanceFallback(String accountId, Throwable t) {
        log.warn("Account service unavailable, circuit breaker open", t);
        throw new ServiceUnavailableException("Account service temporarily unavailable");
    }
}
```

```yaml
# application.yml — Resilience4j config
resilience4j:
  circuitbreaker:
    instances:
      account-service:
        sliding-window-size: 10
        failure-rate-threshold: 50
        wait-duration-in-open-state: 30s
        permitted-number-of-calls-in-half-open-state: 3
      fraud-service:
        sliding-window-size: 5
        failure-rate-threshold: 50
        wait-duration-in-open-state: 10s
  retry:
    instances:
      account-service:
        max-attempts: 3
        wait-duration: 1s
        exponential-backoff-multiplier: 2
        retry-exceptions:
          - java.net.ConnectException
          - java.net.SocketTimeoutException
  timelimiter:
    instances:
      account-service:
        timeout-duration: 5s
      fraud-service:
        timeout-duration: 2s
```

### Resilience Standards

| Pattern | When | Configuration |
|---|---|---|
| Circuit breaker | Every external service call | Open at 50% failure; 30s wait; 3 half-open calls |
| Retry | Idempotent operations; transient failures | Max 3; exponential backoff with jitter |
| Timeout | Every external call | Connect: 5s; read: 30s (configurable per service) |
| Bulkhead | Isolate thread pools per dependency | Prevent one slow service from exhausting all threads |
| Rate limiter | Protect downstream services | Match downstream service's capacity |
| Fallback | Every circuit breaker | Graceful degradation; never silent failure |

---

## 12. Asynchronous Messaging — Kafka (ref: DV-BP-017)

### Producer

```java
@Component
public class KafkaEventPublisher implements EventPublisherPort {
    private final KafkaTemplate<String, Object> kafkaTemplate;

    public void publish(DomainEvent event) {
        var record = new ProducerRecord<String, Object>(
            event.getTopic(),                    // payment.transfer.initiated
            event.getPartitionKey(),             // transferId (ordering guarantee)
            event.toAvro());                     // Avro-serialized payload

        record.headers().add("correlationId", MDC.get("correlationId").getBytes());
        record.headers().add("eventType", event.getType().getBytes());

        kafkaTemplate.send(record)
            .whenComplete((result, ex) -> {
                if (ex != null) {
                    log.error("Failed to publish event: {}", event.getType(), ex);
                    // Transactional outbox handles retry
                }
            });
    }
}
```

### Consumer

```java
@Component
public class TransferEventConsumer {

    @KafkaListener(
        topics = "payment.transfer.initiated",
        groupId = "saga-orchestrator",
        containerFactory = "kafkaListenerContainerFactory")
    public void handleTransferInitiated(
            @Payload TransferInitiatedEvent event,
            @Header("correlationId") String correlationId,
            Acknowledgment ack) {
        MDC.put("correlationId", correlationId);
        try {
            sagaOrchestrator.handleTransferInitiated(event);
            ack.acknowledge(); // Manual ack after successful processing
        } catch (Exception e) {
            log.error("Failed to process TransferInitiated: {}", event.getTransferId(), e);
            // Don't ack — message will be retried
            // After max retries, goes to DLQ
        } finally {
            MDC.clear();
        }
    }
}
```

### Kafka Standards

| Standard | Configuration |
|---|---|
| Acks | all (acks=-1) — guarantee durability |
| Idempotent producer | enable.idempotence=true |
| Manual ack | enable-auto-commit: false; ack after processing |
| DLQ | Every consumer topic has a .dlq topic |
| Schema | Avro with Confluent Schema Registry |
| Partitioning | Business key (accountId, transferId) for ordering |
| Consumer group | One per service |
| Error handling | Retry 3x → DLQ; alert on DLQ messages |

---

## 13. Idempotency (ref: DV-BP-018)

```java
// Idempotency filter for POST endpoints
@Component
public class IdempotencyFilter extends OncePerRequestFilter {
    private final IdempotencyStore store; // Redis-backed

    @Override
    protected void doFilterInternal(HttpServletRequest request,
            HttpServletResponse response, FilterChain chain) throws Exception {
        if (!"POST".equals(request.getMethod())) {
            chain.doFilter(request, response);
            return;
        }

        String idempotencyKey = request.getHeader("Idempotency-Key");
        if (idempotencyKey == null) {
            response.sendError(400, "Idempotency-Key header required for POST");
            return;
        }

        Optional<CachedResponse> cached = store.get(idempotencyKey);
        if (cached.isPresent()) {
            // Return stored response
            writeCachedResponse(response, cached.get());
            return;
        }

        // Wrap response to capture for caching
        var wrapper = new ContentCachingResponseWrapper(response);
        chain.doFilter(request, wrapper);

        // Store response with 24h TTL
        store.put(idempotencyKey, CachedResponse.from(wrapper), Duration.ofHours(24));
        wrapper.copyBodyToResponse();
    }
}
```

---

## 14. Distributed Tracing — OpenTelemetry (ref: DV-BP-021)

```yaml
# application.yml
management:
  tracing:
    sampling:
      probability: 1.0  # 100% in staging; reduce in production
  otlp:
    tracing:
      endpoint: ${OTEL_EXPORTER_OTLP_ENDPOINT:http://localhost:4318/v1/traces}

# pom.xml dependencies
# spring-boot-starter-actuator
# micrometer-tracing-bridge-otel
# opentelemetry-exporter-otlp
```

### Tracing Standards

| Standard | Implementation |
|---|---|
| Auto-instrumentation | Spring Boot + Micrometer Tracing (auto-instruments REST, JPA, Kafka) |
| Correlation ID | Propagated via W3C Trace Context headers + custom X-Correlation-ID |
| Span naming | {service}.{operation} (transfer-service.initiateTransfer) |
| Span attributes | Business context: transfer.amount, transfer.currency, transfer.status |
| Sampling | 100% in staging; 10-50% in production (configurable) |
| Export | OTLP to Jaeger/Tempo/Datadog |

---

## 15. Integration Contract Documentation (ref: DV-BP-022)

### Integration Contract Template

| Field | Content |
|---|---|
| Service | Provider service name |
| Consumer | Consuming service name |
| Protocol | REST / gRPC / Kafka |
| Endpoint/Topic | URL path or topic name |
| Authentication | OAuth scope / mTLS / SASL |
| Request schema | OpenAPI ref or Avro schema |
| Response schema | OpenAPI ref or Avro schema |
| SLA | Response time (p95), availability, throughput |
| Rate limit | Requests per minute |
| Timeout | Connect + read timeout |
| Retry policy | Max retries, backoff strategy |
| Circuit breaker | Threshold, wait duration |
| Fallback | Behavior when unavailable |
| Error codes | Expected error responses |
| Owner | Team responsible |
| Contact | On-call channel |

---

## 16. Testing Standards

### Test Pyramid

| Level | Scope | Tool | Coverage |
|---|---|---|---|
| Unit | Domain model, services, mappers | JUnit 5 + Mockito | >= 90% domain; >= 80% overall |
| Integration | Repository, Kafka, REST clients | Testcontainers + WireMock | All integration points |
| Contract | API contract, event schemas | Spring Cloud Contract / Pact | All endpoints + events |
| Security | Auth, authz, injection | Spring Security Test + custom | OWASP Top 10 |
| Performance | Load, stress | Gatling | p95 < target at expected load |

### Test Patterns

```java
// Controller test — @WebMvcTest (no full context)
@WebMvcTest(TransferController.class)
class TransferControllerTest {
    @Autowired MockMvc mockMvc;
    @MockBean InitiateTransferUseCase useCase;

    @Test
    @WithMockJwt(scopes = "transfers:write", customerId = "CUST-001")
    void initiateTransfer_validRequest_returns201() throws Exception {
        mockMvc.perform(post("/v1/transfers")
                .header("Idempotency-Key", UUID.randomUUID().toString())
                .contentType(APPLICATION_JSON)
                .content("""
                    {"sourceAccountId":"ACC-001","beneficiaryId":"BEN-001",
                     "amount":5000.00,"currency":"USD"}"""))
            .andExpect(status().isCreated())
            .andExpect(jsonPath("$.data.transferId").exists())
            .andExpect(jsonPath("$.data.sourceAccount").value("****-001"));
    }

    @Test
    void initiateTransfer_noAuth_returns401() throws Exception {
        mockMvc.perform(post("/v1/transfers")
                .contentType(APPLICATION_JSON)
                .content("{}"))
            .andExpect(status().isUnauthorized());
    }
}

// Repository test — @DataJpaTest (DB only)
@DataJpaTest
@AutoConfigureTestDatabase(replace = NONE)
@Testcontainers
class TransferJpaRepositoryTest {
    @Container
    static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:16");

    @Test
    void save_validTransfer_persists() {
        // ...
    }
}
```

---

## 17. CI/CD Pipeline

```yaml
# .github/workflows/ci.yml
stages:
  - compile:        mvn compile
  - unit-test:      mvn test (fail if coverage < 80%)
  - sast:           SonarQube / Checkmarx scan
  - sca:            mvn dependency-check:check (OWASP)
  - build:          mvn package -DskipTests
  - container:      docker build + Trivy scan
  - integration:    mvn verify -Pintegration (Testcontainers)
  - contract:       mvn verify -Pcontract (Pact/SCC)
  - deploy-staging: helm upgrade --install
  - dast:           OWASP ZAP scan on staging
  - performance:    Gatling test on staging
  - deploy-prod:    helm upgrade --install (after approvals)
```

### CI/CD Gates

| Gate | Blocks | Threshold |
|---|---|---|
| Compilation | Merge | Zero errors |
| Unit test coverage | Merge | < 80% overall |
| SAST findings | Merge | Any critical/high |
| SCA (dependency) | Merge | Any critical CVE |
| Container scan (Trivy) | Merge | Any critical/high |
| Integration tests | Deploy | Any failure |
| DAST findings | Deploy | Any critical/high |
| Performance | Deploy | p95 > target |

---

## 18. Common Anti-Patterns

| Anti-Pattern | Problem | Fix |
|---|---|---|
| Field injection (@Autowired on fields) | Untestable; hidden dependencies | Constructor injection |
| open-in-view: true | Lazy loading in controllers; N+1 queries | open-in-view: false; explicit fetching |
| @Transactional on controller | Transaction too broad; includes HTTP processing | @Transactional on service layer |
| Exposing JPA entities as API response | Tight coupling; sensitive data leakage | DTO pattern (DV-BP-011) |
| No circuit breaker on external calls | Cascade failures | Resilience4j on every external call |
| Catching Exception and swallowing | Silent failures; lost errors | Log + rethrow or handle specifically |
| ddl-auto: update in production | Uncontrolled schema changes | Flyway/Liquibase; ddl-auto: validate |
| Secrets in application.yml | Credential exposure | Spring Cloud Vault / Secrets Manager |
| No idempotency on POST | Duplicate transactions on retry | Idempotency-Key for financial POSTs |
| Synchronous event publishing | Lost events if app crashes after DB commit | Transactional outbox pattern |
| No correlation ID propagation | Can't trace requests across services | MDC + header propagation |
| Actuator endpoints unsecured | Information disclosure; shutdown risk | Secure all except health/info |
