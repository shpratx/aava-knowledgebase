# Spring Boot Application Standards & Best Practices — Part 1
### Project Structure, Core Standards & Coding Guidelines
### Banking Domain — Agentic Knowledge Base
### References: DV-BP-009, DV-BP-010, DV-BP-011, DV-BP-012, DV-BP-013, DV-BP-015, DV-BP-016

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
