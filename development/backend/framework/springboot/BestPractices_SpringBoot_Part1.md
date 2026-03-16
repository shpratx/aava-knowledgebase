# Spring Boot Application Standards & Best Practices
---

## 1. Project Structure (Hexagonal Architecture)

```
transfer-service/
├── pom.xml (or build.gradle.kts)
├── src/
│   ├── main/
│   │   ├── java/com/bank/transfer/
│   │   │   ├── TransferServiceApplication.java          # @SpringBootApplication
│   │   │   │
│   │   │   ├── domain/                                  # Domain layer (ZERO framework deps)
│   │   │   │   ├── model/
│   │   │   │   │   ├── Transfer.java                    # Aggregate root
│   │   │   │   │   ├── TransferStatus.java              # Enum
│   │   │   │   │   ├── Money.java                       # Value object (BigDecimal + Currency)
│   │   │   │   │   └── AccountReference.java            # Value object
│   │   │   │   ├── event/
│   │   │   │   │   ├── TransferInitiated.java           # Domain event
│   │   │   │   │   ├── TransferCompleted.java
│   │   │   │   │   └── TransferFailed.java
│   │   │   │   ├── service/
│   │   │   │   │   └── TransferDomainService.java       # Domain logic
│   │   │   │   ├── repository/
│   │   │   │   │   └── TransferRepository.java          # Port (interface)
│   │   │   │   └── exception/
│   │   │   │       ├── InsufficientFundsException.java
│   │   │   │       ├── DailyLimitExceededException.java
│   │   │   │       └── AccountFrozenException.java
│   │   │   │
│   │   │   ├── application/                             # Application layer (use cases)
│   │   │   │   ├── port/
│   │   │   │   │   ├── in/                              # Inbound ports
│   │   │   │   │   │   ├── InitiateTransferUseCase.java
│   │   │   │   │   │   ├── GetTransferStatusUseCase.java
│   │   │   │   │   │   └── CancelTransferUseCase.java
│   │   │   │   │   └── out/                             # Outbound ports
│   │   │   │   │       ├── AccountPort.java
│   │   │   │   │       ├── FraudPort.java
│   │   │   │   │       ├── EventPublisherPort.java
│   │   │   │   │       └── AuditPort.java
│   │   │   │   └── service/
│   │   │   │       ├── TransferApplicationService.java  # Use case impl
│   │   │   │       └── TransferSagaOrchestrator.java
│   │   │   │
│   │   │   └── infrastructure/                          # Infrastructure layer (adapters)
│   │   │       ├── web/                                 # Inbound: REST API
│   │   │       │   ├── TransferController.java
│   │   │       │   ├── dto/
│   │   │       │   │   ├── TransferRequest.java         # API request DTO
│   │   │       │   │   ├── TransferResponse.java        # API response DTO
│   │   │       │   │   └── ErrorResponse.java
│   │   │       │   ├── mapper/
│   │   │       │   │   └── TransferDtoMapper.java       # DTO <-> Domain
│   │   │       │   └── validation/
│   │   │       │       └── TransferRequestValidator.java
│   │   │       ├── persistence/                         # Outbound: Database
│   │   │       │   ├── TransferJpaRepository.java       # Spring Data JPA
│   │   │       │   ├── TransferJpaAdapter.java          # Implements TransferRepository port
│   │   │       │   ├── entity/
│   │   │       │   │   └── TransferEntity.java          # JPA entity
│   │   │       │   └── mapper/
│   │   │       │       └── TransferEntityMapper.java    # Entity <-> Domain
│   │   │       ├── messaging/                           # Outbound: Events
│   │   │       │   ├── KafkaEventPublisher.java
│   │   │       │   ├── KafkaEventConsumer.java
│   │   │       │   └── schema/                          # Avro schemas
│   │   │       ├── client/                              # Outbound: External services
│   │   │       │   ├── AccountServiceClient.java        # REST client
│   │   │       │   ├── FraudServiceClient.java          # gRPC client
│   │   │       │   └── config/
│   │   │       │       ├── AccountClientConfig.java
│   │   │       │       └── FraudClientConfig.java
│   │   │       ├── security/                            # Security config
│   │   │       │   ├── SecurityConfig.java
│   │   │       │   ├── JwtAuthenticationFilter.java
│   │   │       │   └── ResourceOwnershipChecker.java
│   │   │       ├── observability/                       # Logging, tracing, metrics
│   │   │       │   ├── CorrelationIdFilter.java
│   │   │       │   ├── RequestLoggingFilter.java
│   │   │       │   └── AuditInterceptor.java
│   │   │       └── config/                              # Framework config
│   │   │           ├── DatabaseConfig.java
│   │   │           ├── KafkaConfig.java
│   │   │           ├── ResilienceConfig.java
│   │   │           ├── OpenApiConfig.java
│   │   │           └── ActuatorSecurityConfig.java
│   │   │
│   │   └── resources/
│   │       ├── application.yml                          # Default config
│   │       ├── application-dev.yml                      # Dev profile
│   │       ├── application-staging.yml                  # Staging profile
│   │       ├── application-prod.yml                     # Production profile
│   │       ├── db/migration/                            # Flyway migrations
│   │       │   ├── V1.0.0__create_transfers_table.sql
│   │       │   ├── V1.0.1__add_audit_trigger.sql
│   │       │   └── V1.0.2__add_indexes.sql
│   │       └── openapi/
│   │           └── transfer-api.yaml                    # OpenAPI 3.0 spec
│   │
│   └── test/
│       ├── java/com/bank/transfer/
│       │   ├── domain/                                  # Domain unit tests
│       │   ├── application/                             # Use case tests
│       │   ├── infrastructure/
│       │   │   ├── web/                                 # Controller tests (@WebMvcTest)
│       │   │   ├── persistence/                         # Repository tests (@DataJpaTest)
│       │   │   ├── messaging/                           # Kafka tests (Testcontainers)
│       │   │   └── client/                              # Client tests (WireMock)
│       │   └── integration/                             # Full integration tests
│       └── resources/
│           ├── application-test.yml
│           └── fixtures/                                # Test data
│
├── Dockerfile
├── docker-compose.yml                                   # Local dev dependencies
├── .github/workflows/ci.yml                             # CI/CD pipeline
└── README.md
```

### Layer Rules

| Layer | Can Depend On | Cannot Depend On | Framework Deps |
|---|---|---|---|
| Domain | Nothing (pure Java) | Application, Infrastructure | None |
| Application | Domain | Infrastructure | None (interfaces only) |
| Infrastructure | Domain, Application | — | Spring, JPA, Kafka, etc. |

---

## 2. Spring Security Configuration (ref: DV-BP-009)

```java
@Configuration
@EnableWebSecurity
@EnableMethodSecurity
public class SecurityConfig {

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        return http
            .csrf(csrf -> csrf.disable()) // Stateless API; CSRF via SameSite cookies
            .sessionManagement(session -> session
                .sessionCreationPolicy(SessionCreationPolicy.STATELESS))
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/actuator/health", "/actuator/info").permitAll()
                .requestMatchers("/v3/api-docs/**", "/swagger-ui/**").permitAll()
                .requestMatchers("/v1/transfers/**").hasAuthority("SCOPE_transfers:write")
                .anyRequest().authenticated())
            .oauth2ResourceServer(oauth2 -> oauth2
                .jwt(jwt -> jwt.jwtAuthenticationConverter(jwtAuthConverter())))
            .headers(headers -> headers
                .contentTypeOptions(Customizer.withDefaults())
                .frameOptions(frame -> frame.deny())
                .cacheControl(Customizer.withDefaults()))
            .build();
    }
}
```

### Security Standards

| Standard | Implementation |
|---|---|
| Authentication | OAuth 2.0 Resource Server with JWT validation |
| Authorization | Scope-based (@PreAuthorize) + resource-level ownership check |
| CORS | Restrictive — specific origins only; no wildcard in production |
| Headers | X-Content-Type-Options, X-Frame-Options: DENY, Cache-Control: no-store |
| Actuator | Health/info public; all others secured behind admin role |
| Secrets | Spring Cloud Vault or AWS Secrets Manager — never in application.yml |

---

## 3. Global Exception Handling (ref: DV-BP-010)

```java
@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(InsufficientFundsException.class)
    public ResponseEntity<ErrorResponse> handleInsufficientFunds(InsufficientFundsException ex) {
        return ResponseEntity.status(422).body(ErrorResponse.of(
            "INSUFFICIENT_FUNDS", "Insufficient funds for this transfer"));
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleValidation(MethodArgumentNotValidException ex) {
        var details = ex.getBindingResult().getFieldErrors().stream()
            .map(e -> new FieldError(e.getField(), e.getDefaultMessage(), e.getRejectedValue()))
            .toList();
        return ResponseEntity.badRequest().body(ErrorResponse.validation(details));
    }

    @ExceptionHandler(AccessDeniedException.class)
    public ResponseEntity<ErrorResponse> handleAccessDenied(AccessDeniedException ex) {
        return ResponseEntity.status(403).body(ErrorResponse.of(
            "ACCESS_DENIED", "Access denied"));
    }

    // NEVER expose internal details
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGeneric(Exception ex) {
        log.error("Unhandled exception", ex); // Log full stack trace server-side
        return ResponseEntity.status(500).body(ErrorResponse.of(
            "INTERNAL_ERROR", "Something went wrong. Please try again."));
    }
}
```

### Error Response Standard

```java
public record ErrorResponse(
    ErrorBody error,
    Meta meta
) {
    public record ErrorBody(String code, String message, List<FieldError> details) {}
    public record FieldError(String field, String message, Object rejectedValue) {}
    public record Meta(String requestId, Instant timestamp) {}
}
```

---

## 4. DTO Pattern (ref: DV-BP-011)

```java
// API Request DTO — validated, no domain logic
public record TransferRequest(
    @NotNull @ValidUUID String sourceAccountId,
    @NotNull @ValidUUID String beneficiaryId,
    @NotNull @DecimalMin("0.01") @DecimalMax("999999999.99") BigDecimal amount,
    @NotNull @ValidCurrency String currency,
    @Size(max = 140) @Pattern(regexp = "^[a-zA-Z0-9 \\-/.]*$") String reference
) {}

// API Response DTO — masked sensitive data
public record TransferResponse(
    String transferId,
    String status,
    BigDecimal amount,
    String currency,
    String sourceAccount,    // Masked: ****1234
    String beneficiary,      // Masked: ****5678
    Instant initiatedAt
) {}

// Mapper — DTO <-> Domain (use MapStruct for complex mappings)
@Component
public class TransferDtoMapper {
    public TransferResponse toResponse(Transfer transfer) {
        return new TransferResponse(
            transfer.getId().toString(),
            transfer.getStatus().name(),
            transfer.getAmount().amount(),
            transfer.getAmount().currency().getCurrencyCode(),
            maskAccount(transfer.getSourceAccountId()),
            maskAccount(transfer.getBeneficiaryId()),
            transfer.getInitiatedAt());
    }
    private String maskAccount(String account) {
        return "****" + account.substring(account.length() - 4);
    }
}
```

### DTO Rules

| Rule | Standard |
|---|---|
| Separate DTOs from domain | API DTOs in infrastructure/web/dto; never expose domain entities |
| Validation on DTOs | Bean Validation annotations on request DTOs |
| Immutable | Use Java records for DTOs |
| No business logic | DTOs are data carriers only |
| Sensitive data masking | Response DTOs mask account numbers, PII |
| Mapper layer | Explicit mapper between DTO and domain; use MapStruct for complex cases |

---

## 5. Request/Response Logging with Correlation IDs (ref: DV-BP-012)

```java
@Component
@Order(Ordered.HIGHEST_PRECEDENCE)
public class CorrelationIdFilter extends OncePerRequestFilter {
    private static final String CORRELATION_HEADER = "X-Correlation-ID";
    private static final String REQUEST_HEADER = "X-Request-ID";

    @Override
    protected void doFilterInternal(HttpServletRequest request,
            HttpServletResponse response, FilterChain chain) throws Exception {
        String correlationId = Optional.ofNullable(request.getHeader(CORRELATION_HEADER))
            .orElse(UUID.randomUUID().toString());
        String requestId = UUID.randomUUID().toString();

        MDC.put("correlationId", correlationId);
        MDC.put("requestId", requestId);
        response.setHeader(CORRELATION_HEADER, correlationId);
        response.setHeader(REQUEST_HEADER, requestId);

        try {
            chain.doFilter(request, response);
        } finally {
            MDC.clear();
        }
    }
}
```

### Logging Standards

| Standard | Implementation |
|---|---|
| Format | Structured JSON (Logback + logstash-logback-encoder) |
| Correlation ID | In every log entry via MDC; propagated to downstream calls |
| Request logging | Method, URI, status, duration — no request body for sensitive endpoints |
| Sensitive data | NEVER log passwords, tokens, PAN, CVV, full account numbers |
| Log levels | ERROR: unhandled exceptions; WARN: business rule violations; INFO: request/response; DEBUG: dev only |
| Retention | Application logs: 30 days hot; audit logs: 7 years |

---

## 6. Actuator & Health Checks (ref: DV-BP-013)

```yaml
# application.yml
management:
  endpoints:
    web:
      exposure:
        include: health, info, metrics, prometheus
  endpoint:
    health:
      show-details: when-authorized
      probes:
        enabled: true
  health:
    db:
      enabled: true
    kafka:
      enabled: true
    diskSpace:
      enabled: true
```

```java
// Custom health indicator for core banking dependency
@Component
public class CoreBankingHealthIndicator implements HealthIndicator {
    private final AccountServiceClient accountClient;

    @Override
    public Health health() {
        try {
            accountClient.healthCheck();
            return Health.up().withDetail("coreBanking", "available").build();
        } catch (Exception e) {
            return Health.down().withDetail("coreBanking", "unavailable").build();
        }
    }
}
```

### Actuator Security

| Endpoint | Access | Purpose |
|---|---|---|
| /actuator/health | Public (liveness/readiness) | Kubernetes probes |
| /actuator/info | Public | Version, build info |
| /actuator/metrics | Authenticated (admin) | Prometheus metrics |
| /actuator/prometheus | Authenticated (monitoring) | Prometheus scrape |
| /actuator/env, /actuator/configprops | Disabled in production | Exposes secrets |
| /actuator/shutdown | Disabled | Dangerous |

---

## 7. Connection Pooling & Timeouts (ref: DV-BP-015)

```yaml
# HikariCP configuration
spring:
  datasource:
    hikari:
      minimum-idle: 5
      maximum-pool-size: 20
      connection-timeout: 5000      # 5s to get connection
      idle-timeout: 600000          # 10 min idle before release
      max-lifetime: 1800000         # 30 min max connection life
      leak-detection-threshold: 60000  # 60s leak detection
      pool-name: transfer-write-pool

  # Read replica (if CQRS)
  datasource-read:
    hikari:
      minimum-idle: 10
      maximum-pool-size: 50
      pool-name: transfer-read-pool
```

### Timeout Standards

| Component | Connect Timeout | Read Timeout | Total Timeout |
|---|---|---|---|
| Database (HikariCP) | 5s | — | 30s (max lifetime) |
| REST client (account-service) | 5s | 30s | 35s |
| gRPC client (fraud-service) | 2s | 2s | 4s |
| Kafka producer | 5s | — | 30s (delivery timeout) |
| Redis | 2s | 2s | 4s |

---

## 8. Transaction Management (ref: DV-BP-016)

```java
// Application service — transaction boundary
@Service
@Transactional(readOnly = true) // Default read-only
public class TransferApplicationService implements InitiateTransferUseCase {

    @Override
    @Transactional // Write transaction for this method
    public Transfer initiateTransfer(InitiateTransferCommand command) {
        // 1. Validate (no DB write yet)
        Transfer transfer = transferDomainService.createTransfer(command);

        // 2. Save (within transaction)
        transferRepository.save(transfer);

        // 3. Publish event (via transactional outbox — same transaction)
        outboxRepository.save(new OutboxEvent(transfer.getInitiatedEvent()));

        return transfer;
    }
}
```

### Transaction Rules

| Rule | Standard |
|---|---|
| Boundary | @Transactional on application service methods, not on domain or infrastructure |
| Default | readOnly = true at class level; write on specific methods |
| Propagation | REQUIRED (default) for most; REQUIRES_NEW for audit logging |
| Isolation | READ_COMMITTED (default); SERIALIZABLE for balance updates if needed |
| Rollback | Default rollback on RuntimeException; explicit rollbackFor if checked exceptions |
| Scope | Keep transactions short — no external HTTP calls inside transaction |
| Outbox | Use transactional outbox for event publishing (same transaction as DB write) |
| Testing | @Transactional on tests for auto-rollback; @Commit for integration tests that verify DB |

---

## 9. Application Configuration

```yaml
# application.yml — base configuration
spring:
  application:
    name: transfer-service
  profiles:
    active: ${SPRING_PROFILES_ACTIVE:dev}

  jpa:
    open-in-view: false                    # Prevent lazy loading in controllers
    hibernate:
      ddl-auto: validate                   # Flyway manages schema; Hibernate validates
    properties:
      hibernate:
        default_schema: transfer
        jdbc:
          batch_size: 50
        order_inserts: true
        order_updates: true

  flyway:
    enabled: true
    locations: classpath:db/migration
    schemas: transfer
    validate-on-migrate: true

  kafka:
    bootstrap-servers: ${KAFKA_BOOTSTRAP_SERVERS}
    producer:
      acks: all
      retries: 3
      properties:
        enable.idempotence: true
    consumer:
      group-id: transfer-service
      auto-offset-reset: earliest
      enable-auto-commit: false

server:
  port: 8080
  shutdown: graceful
  servlet:
    context-path: /
  error:
    include-message: never               # Never expose error messages
    include-stacktrace: never            # Never expose stack traces
    include-binding-errors: never

logging:
  level:
    root: INFO
    com.bank.transfer: INFO
    org.springframework.security: WARN
  pattern:
    console: '{"timestamp":"%d","level":"%p","service":"${spring.application.name}","correlationId":"%X{correlationId}","requestId":"%X{requestId}","logger":"%logger","message":"%m"}%n'
```

---

## 10. Coding Guidelines

### Java Standards

| Standard | Requirement |
|---|---|
| Java version | 21+ (LTS) |
| Records | Use for DTOs, value objects, events — immutable by default |
| Optional | Return type only; never as parameter or field |
| Null safety | @Nullable/@NonNull annotations; Optional for return types |
| Streams | Prefer streams for collection operations; avoid side effects in streams |
| var | Allowed for local variables when type is obvious from right side |
| Sealed classes | Use for domain event hierarchies, status enums |
| Pattern matching | Use instanceof pattern matching and switch expressions |

### Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Packages | lowercase, domain-aligned | com.bank.transfer.domain.model |
| Classes | PascalCase, noun | TransferApplicationService |
| Interfaces | PascalCase, no I prefix | TransferRepository (not ITransferRepository) |
| Methods | camelCase, verb | initiateTransfer(), getBalance() |
| Constants | UPPER_SNAKE_CASE | MAX_TRANSFER_AMOUNT |
| DTOs | PascalCase + Request/Response | TransferRequest, TransferResponse |
| Entities | PascalCase + Entity suffix | TransferEntity |
| Mappers | PascalCase + Mapper suffix | TransferDtoMapper |
| Config | PascalCase + Config suffix | SecurityConfig |
| Tests | Same name + Test suffix | TransferApplicationServiceTest |

### Dependency Injection

| Rule | Standard |
|---|---|
| Constructor injection | Always — never field injection (@Autowired on fields) |
| Final fields | All injected dependencies are final |
| Single constructor | Implicit @Autowired (no annotation needed) |
| Circular dependencies | Prohibited — redesign if detected |

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
