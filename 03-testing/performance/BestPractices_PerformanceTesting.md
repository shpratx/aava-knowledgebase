# Performance Testing — Best Practices
### Banking Domain — Agentic Knowledge Base

---

## Best Practices

 Best Practices

| Practice | Standard |
|---|---|
| Test in production-equivalent environment | Same infrastructure, same data volume, same network topology |
| Use realistic data volume | Production-equivalent number of records in database |
| Use realistic user behavior | Think time between actions (3-10s); realistic user journeys, not just single endpoints |
| Establish baseline before optimization | Measure current state before making changes |
| Test after every significant change | Code changes, infrastructure changes, dependency upgrades |
| Monitor server-side during tests | CPU, memory, GC, connection pools, thread pools, disk I/O |
| Monitor dependencies during tests | Database query time, Kafka lag, Redis latency, external service response |
| Test with authentication | Include token generation in test flow; don't bypass auth |
| Test with realistic payload sizes | Not empty requests; use production-representative payloads |
| Automate in CI/CD | Run baseline + load test on every release; alert on regression |
| Define pass/fail criteria before testing | Not after seeing results |
| Test database queries independently | Identify slow queries with EXPLAIN ANALYZE before load test |
| Test under failure conditions | Performance with one instance down; with degraded dependency |
| Report trends, not just snapshots | Track p95 over releases; detect gradual degradation |

### Anti-Patterns
| Anti-Pattern | Fix |
|---|---|
| Testing with empty database | Use production-equivalent data volume |
| Testing without auth | Include token generation; auth adds latency |
| Testing single endpoint in isolation | Test realistic user journeys with multiple endpoints |
| No think time | Add realistic pauses (3-10s) between actions |
| Testing only happy path | Include error scenarios (validation failures, auth failures) |
| Running from same network as server | Test from realistic network distance; or account for it |
| No server-side monitoring | Monitor CPU, memory, GC, connections during test |
| Optimizing before measuring | Establish baseline first; optimize based on data |
| One-time test before release | Automate in CI/CD; test continuously |
| Ignoring p99 | p95 hides tail latency; p99 shows worst-case user experience |

---