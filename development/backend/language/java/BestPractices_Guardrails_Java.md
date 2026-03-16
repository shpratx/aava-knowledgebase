# Java Application Development Standards, Best Practices & Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Part A: Standards & Best Practices

### 1. Language Standards

| Standard | Requirement |
|---|---|
| Java version | 21+ (LTS) — use latest LTS in production |
| Source encoding | UTF-8 |
| Compilation | -Werror (warnings as errors); -Xlint:all |
| Preview features | Prohibited in production — only GA features |
| Build tool | Maven or Gradle (Kotlin DSL preferred); reproducible builds |

### 2. Type System & Immutability

| Standard | Requirement | Example |
|---|---|---|
| Records | Use for DTOs, value objects, events, config | `record Money(BigDecimal amount, Currency currency)` |
| Sealed classes | Use for domain event hierarchies, type-safe enums | `sealed interface TransferEvent permits Initiated, Completed, Failed` |
| Final classes | Default to final; extend only with justification | `public final class TransferService` |
| Immutable collections | Return unmodifiable collections from public APIs | `List.copyOf()`, `Collections.unmodifiableList()` |
| Optional | Return type only; never as parameter, field, or collection element | `Optional<Transfer> findById(UUID id)` |
| var | Local variables only when type is obvious from right side | `var transfers = repository.findAll()` |
| Null safety | @Nullable/@NonNull annotations; Optional for return; never return null for collections | Return `List.of()` not `null` |

### 3. Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Packages | lowercase, domain-aligned, no underscores | `com.bank.transfer.domain.model` |
| Classes | PascalCase, noun, descriptive | `TransferApplicationService` |
| Interfaces | PascalCase, no I prefix, capability-named | `TransferRepository` (not `ITransferRepository`) |
| Methods | camelCase, verb-first | `initiateTransfer()`, `getBalance()`, `isActive()` |
| Constants | UPPER_SNAKE_CASE, static final | `MAX_TRANSFER_AMOUNT`, `DEFAULT_CURRENCY` |
| Enums | PascalCase class, UPPER_SNAKE_CASE values | `TransferStatus.COMPLETED` |
| Type parameters | Single uppercase letter or descriptive | `<T>`, `<E extends DomainEvent>` |
| Boolean methods | is/has/can/should prefix | `isActive()`, `hasPermission()`, `canTransfer()` |
| Test methods | descriptive with underscores | `initiateTransfer_insufficientFunds_throwsException()` |

### 4. Monetary & Financial Data

| Standard | Requirement |
|---|---|
| Money representation | BigDecimal — NEVER float/double for financial calculations |
| Scale | Explicit scale per currency (USD: 2, JPY: 0, BHD: 3) |
| Rounding | RoundingMode.HALF_EVEN (banker's rounding) — always explicit |
| Comparison | Use compareTo(), never equals() for BigDecimal (scale-sensitive) |
| Money value object | Encapsulate amount + currency together; validate in constructor |
| Arithmetic | Use BigDecimal methods (add, subtract, multiply, divide with scale + rounding) |
| Currency | java.util.Currency or ISO 4217 string (CHAR(3)) |

```java
public record Money(BigDecimal amount, Currency currency) {
    public Money {
        Objects.requireNonNull(amount, "Amount must not be null");
        Objects.requireNonNull(currency, "Currency must not be null");
        if (amount.scale() > currency.getDefaultFractionDigits()) {
            throw new IllegalArgumentException(
                "Scale %d exceeds max %d for %s".formatted(
                    amount.scale(), currency.getDefaultFractionDigits(), currency));
        }
    }

    public Money add(Money other) {
        requireSameCurrency(other);
        return new Money(amount.add(other.amount), currency);
    }

    public Money subtract(Money other) {
        requireSameCurrency(other);
        return new Money(amount.subtract(other.amount), currency);
    }

    public boolean isPositive() { return amount.compareTo(BigDecimal.ZERO) > 0; }
    public boolean isNegative() { return amount.compareTo(BigDecimal.ZERO) < 0; }

    private void requireSameCurrency(Money other) {
        if (!currency.equals(other.currency))
            throw new IllegalArgumentException("Currency mismatch: %s vs %s".formatted(currency, other.currency));
    }
}
```

### 5. Exception Handling

| Standard | Requirement |
|---|---|
| Domain exceptions | Extend RuntimeException; specific per business rule |
| Checked exceptions | Avoid in domain; wrap in unchecked at infrastructure boundary |
| Never swallow | Every catch must log or rethrow — no empty catch blocks |
| Never catch Throwable | Catch specific exceptions; Exception as last resort |
| Error messages | Descriptive for developers; no sensitive data in messages |
| Stack traces | Log server-side; never expose to API consumers |
| Try-with-resources | Always for AutoCloseable resources (streams, connections) |
| Multi-catch | Use when handling multiple exceptions the same way |

```java
// Domain exception hierarchy
public sealed class TransferException extends RuntimeException
    permits InsufficientFundsException, DailyLimitExceededException, AccountFrozenException {
    private final String errorCode;
    public TransferException(String errorCode, String message) {
        super(message);
        this.errorCode = errorCode;
    }
}

public final class InsufficientFundsException extends TransferException {
    public InsufficientFundsException(Money available, Money requested) {
        super("INSUFFICIENT_FUNDS",
            "Insufficient funds: available=%s, requested=%s".formatted(available, requested));
    }
}
```

### 6. Collections & Streams

| Standard | Requirement |
|---|---|
| Prefer List.of(), Map.of(), Set.of() | Immutable factory methods for small collections |
| Return unmodifiable | Public methods return unmodifiable collections |
| Streams for transformation | Use streams for map/filter/reduce; not for side effects |
| No side effects in streams | forEach only for terminal output; never modify external state |
| Parallel streams | Avoid unless proven performance benefit with benchmarks |
| Collectors | Use toList() (Java 16+) instead of Collectors.toList() |
| Null in collections | Prohibited — never add null to collections |

### 7. Concurrency

| Standard | Requirement |
|---|---|
| Thread safety | Immutable objects preferred; synchronized only when necessary |
| Virtual threads | Use for I/O-bound operations (Java 21+) |
| CompletableFuture | For async composition; always handle exceptionally |
| Executors | Use Executors factory methods; never raw Thread creation |
| Shared mutable state | Minimize; use ConcurrentHashMap, AtomicReference when needed |
| Locks | Prefer ReentrantLock over synchronized for complex locking |
| Thread-local | Clean up in finally block; avoid in virtual threads |

### 8. Date & Time

| Standard | Requirement |
|---|---|
| API | java.time exclusively — never java.util.Date or Calendar |
| Storage | Instant (UTC) for timestamps; LocalDate for dates without time |
| Timezone | Store in UTC; convert to user timezone for display only |
| Business days | Use custom BusinessDayCalendar for banking day calculations |
| Formatting | DateTimeFormatter (thread-safe); ISO 8601 for API exchange |
| Comparison | Use isBefore(), isAfter(), isEqual() — not compareTo() for readability |

### 9. String Handling

| Standard | Requirement |
|---|---|
| Text blocks | Use for multi-line strings (SQL, JSON templates) |
| String.formatted() | Prefer over String.format() (Java 15+) |
| StringBuilder | Use for loop-based concatenation; + is fine for simple cases |
| Null checks | Use Objects.requireNonNull() for parameters; Optional for returns |
| Comparison | Use "constant".equals(variable) to avoid NPE; or Objects.equals() |
| Sensitive data | Never include in toString(); override toString() to mask |

### 10. Logging

| Standard | Requirement |
|---|---|
| Framework | SLF4J API with Logback implementation |
| Format | Structured JSON (logstash-logback-encoder) |
| Parameterized | Use `log.info("Transfer {} completed", transferId)` — never concatenation |
| Levels | ERROR: unhandled; WARN: business violations; INFO: operations; DEBUG: dev |
| Sensitive data | NEVER log passwords, tokens, PAN, CVV, full account numbers |
| Correlation ID | In every log entry via MDC |
| Performance | Use `log.isDebugEnabled()` guard for expensive debug messages |
| toString() | Override on domain objects to mask sensitive fields |

```java
// Mask sensitive data in toString
public record Transfer(UUID id, String sourceAccount, Money amount, TransferStatus status) {
    @Override
    public String toString() {
        return "Transfer[id=%s, source=****%s, amount=%s, status=%s]".formatted(
            id, sourceAccount.substring(sourceAccount.length() - 4), amount, status);
    }
}
```

### 11. Testing Standards

| Level | Coverage | Tool | What to Test |
|---|---|---|---|
| Unit | >= 90% domain; >= 80% overall | JUnit 5 + Mockito + AssertJ | Domain logic, validators, mappers, utils |
| Integration | All integration points | Testcontainers + WireMock | DB, Kafka, REST clients |
| Contract | All API endpoints + events | Pact / Spring Cloud Contract | API contracts, event schemas |
| Security | OWASP Top 10 applicable | Spring Security Test | Auth bypass, IDOR, injection |
| Performance | p95 < target | JMH (micro); Gatling (load) | Hot paths, critical operations |

| Rule | Standard |
|---|---|
| Test naming | `methodName_condition_expectedResult()` |
| Arrange-Act-Assert | Clear separation in every test |
| One assertion concept per test | Multiple asserts OK if testing one logical concept |
| No test interdependence | Tests must run independently in any order |
| No production data | Synthetic test fixtures only |
| Test behavior, not implementation | Test public API; don't test private methods |
| AssertJ preferred | Fluent assertions: `assertThat(result).isEqualTo(expected)` |

---

## Part B: Guardrails

### 12. Code Quality Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| JV-CQ-001 | No float/double for monetary calculations — BigDecimal only | Reject if float/double used for money |
| JV-CQ-002 | No java.util.Date or Calendar — java.time only | Reject if legacy date API used |
| JV-CQ-003 | No raw types — all generics must be parameterized | CI/CD gate — compiler warning as error |
| JV-CQ-004 | No empty catch blocks — every catch must log or rethrow | CI/CD gate — SonarQube/SpotBugs |
| JV-CQ-005 | No System.out/System.err — use SLF4J logging | CI/CD gate — ESLint equivalent (Checkstyle) |
| JV-CQ-006 | No field injection (@Autowired on fields) — constructor injection only | CI/CD gate — ArchUnit rule |
| JV-CQ-007 | No circular dependencies between packages/classes | CI/CD gate — ArchUnit rule |
| JV-CQ-008 | Maximum method length: 50 lines; maximum class length: 500 lines | CI/CD gate — Checkstyle |
| JV-CQ-009 | Maximum cyclomatic complexity: 10 per method | CI/CD gate — SonarQube |
| JV-CQ-010 | Test coverage >= 80% overall; >= 90% domain layer | CI/CD gate — JaCoCo |

### 13. Security Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| JV-SG-001 | No string concatenation in SQL/JPQL queries — parameterized only | Reject — SAST scan |
| JV-SG-002 | No hardcoded secrets (passwords, keys, tokens) in source code | Reject — secret scanning |
| JV-SG-003 | No sensitive data in log statements (passwords, PAN, CVV, tokens) | Reject — code review + SAST |
| JV-SG-004 | No sensitive data in toString() — override to mask | Reject if domain objects expose sensitive data in toString |
| JV-SG-005 | No sensitive data in exception messages | Reject if exceptions contain PII or credentials |
| JV-SG-006 | All external input must be validated before processing | Reject if input validation missing |
| JV-SG-007 | Serialization/deserialization must use allowlists — no arbitrary class deserialization | Reject if ObjectInputStream used without filtering |
| JV-SG-008 | No Runtime.exec() or ProcessBuilder with user input — command injection risk | Reject if user input in process execution |
| JV-SG-009 | XML parsing must disable external entities (XXE prevention) | Reject if XML parser allows external entities |
| JV-SG-010 | Random number generation for security must use SecureRandom — not Random | Reject if java.util.Random used for tokens/keys |

### 14. Architecture Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| JV-AG-001 | Domain layer must have zero framework dependencies (no Spring, JPA, Kafka annotations) | CI/CD gate — ArchUnit rule |
| JV-AG-002 | Infrastructure layer must not be imported by domain or application layers | CI/CD gate — ArchUnit rule |
| JV-AG-003 | DTOs must be separate from domain entities — never expose JPA entities as API response | Reject if entity used as response |
| JV-AG-004 | Repository interfaces in domain; implementations in infrastructure | CI/CD gate — ArchUnit rule |
| JV-AG-005 | Domain events must be plain Java objects — no framework annotations | Reject if domain events have framework deps |
| JV-AG-006 | Business logic must be in domain layer — not in controllers or infrastructure | Reject if business logic in controller |
| JV-AG-007 | Each bounded context/service must have its own package root | Reject if contexts share packages |

### 15. Dependency Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| JV-DG-001 | All dependencies must be scanned for CVEs on every build (OWASP Dependency-Check) | CI/CD gate — block on critical |
| JV-DG-002 | Critical CVEs must be patched within 24 hours | Reject deployment |
| JV-DG-003 | High CVEs must be patched within 7 days | Reject deployment |
| JV-DG-004 | Only approved licenses (Apache 2.0, MIT, BSD, EPL) — GPL prohibited for proprietary | CI/CD gate — license plugin |
| JV-DG-005 | New dependencies require architecture review | Reject if added without review |
| JV-DG-006 | Dependency versions must be managed centrally (BOM/parent POM) | Reject if versions scattered |
| JV-DG-007 | Snapshot dependencies prohibited in production builds | CI/CD gate — enforcer plugin |

### 16. Performance Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| JV-PG-001 | No N+1 query patterns — use JOIN FETCH, @EntityGraph, or batch fetching | Reject if N+1 detected in tests |
| JV-PG-002 | All collection queries must be paginated — no unbounded result sets | Reject if findAll() without pagination on large tables |
| JV-PG-003 | Database connections must use connection pooling (HikariCP) | Reject if no connection pool |
| JV-PG-004 | No blocking I/O inside @Transactional — keep transactions short | Reject if HTTP calls inside transaction |
| JV-PG-005 | String concatenation in loops must use StringBuilder | Flag if + used in loops |
| JV-PG-006 | Resource cleanup must use try-with-resources for all AutoCloseable | Reject if manual close without try-with-resources |
| JV-PG-007 | Lazy loading must be explicit — open-in-view: false | Reject if open-in-view is true |

---

## ArchUnit Rules Reference

```java
@AnalyzeClasses(packages = "com.bank.transfer")
class ArchitectureTest {

    @ArchTest
    static final ArchRule domainHasNoFrameworkDeps = noClasses()
        .that().resideInAPackage("..domain..")
        .should().dependOnClassesThat()
        .resideInAnyPackage("org.springframework..", "javax.persistence..", "jakarta.persistence..");

    @ArchTest
    static final ArchRule infrastructureNotImportedByDomain = noClasses()
        .that().resideInAPackage("..domain..")
        .should().dependOnClassesThat()
        .resideInAPackage("..infrastructure..");

    @ArchTest
    static final ArchRule noFieldInjection = noFields()
        .should().beAnnotatedWith(Autowired.class)
        .orShould().beAnnotatedWith(Inject.class);

    @ArchTest
    static final ArchRule noCircularDependencies = slices()
        .matching("com.bank.transfer.(*)..")
        .should().beFreeOfCycles();

    @ArchTest
    static final ArchRule controllersOnlyInWeb = classes()
        .that().areAnnotatedWith(RestController.class)
        .should().resideInAPackage("..infrastructure.web..");

    @ArchTest
    static final ArchRule entitiesOnlyInPersistence = classes()
        .that().areAnnotatedWith(Entity.class)
        .should().resideInAPackage("..infrastructure.persistence..");
}
```

---

## Enforcement Summary

| Severity | Action | Examples |
|---|---|---|
| **Reject** | Cannot merge | float for money, empty catch, SQL concatenation, hardcoded secrets, sensitive data in logs/toString, entity as API response, N+1 queries |
| **Flag** | Address before release | String concat in loops, missing toString override, snapshot dependencies |
| **CI/CD Gate** | Automated | SonarQube, SpotBugs, Checkstyle, ArchUnit, JaCoCo, OWASP Dependency-Check, license plugin |

---

## Pre-Merge Checklist

| # | Check |
|---|---|
| 1 | BigDecimal for all monetary calculations; no float/double |
| 2 | java.time for all dates; no java.util.Date |
| 3 | No sensitive data in logs, toString(), or exception messages |
| 4 | No hardcoded secrets in source code |
| 5 | All queries parameterized; no SQL concatenation |
| 6 | Constructor injection only; no field injection |
| 7 | Domain layer has zero framework dependencies (ArchUnit passes) |
| 8 | DTOs separate from entities; sensitive data masked in responses |
| 9 | All external input validated |
| 10 | Test coverage >= 80%; domain >= 90% |
| 11 | No critical/high CVEs in dependencies |
| 12 | No empty catch blocks; all exceptions handled |
| 13 | Cyclomatic complexity <= 10; method <= 50 lines |
| 14 | No N+1 queries; all collections paginated |
