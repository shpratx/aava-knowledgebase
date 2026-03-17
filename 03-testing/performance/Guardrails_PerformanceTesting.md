# Performance Testing — Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Guardrails

 Guardrails

| ID | Guardrail | Enforcement |
|---|---|---|
| TG-PERF-001 | Performance baseline must be established for all customer-facing endpoints before production | Reject if no baseline |
| TG-PERF-002 | Load test must verify p95 within NFR target at expected peak load | Reject if p95 exceeds target |
| TG-PERF-003 | Load test must sustain for minimum 30 minutes at peak | Reject if test duration < 30 min |
| TG-PERF-004 | Error rate must be < 0.1% during load test | Reject if error rate exceeds 0.1% |
| TG-PERF-005 | Test environment must be production-equivalent (infrastructure + data volume) | Reject if test env is not representative |
| TG-PERF-006 | Endurance test (4+ hours) must be run before major releases | Reject major release without endurance test |
| TG-PERF-007 | No memory leaks: memory usage must not trend upward during endurance test | Reject if memory leak detected |
| TG-PERF-008 | No connection leaks: connection pool must be stable during endurance test | Reject if connection leak detected |
| TG-PERF-009 | Stress test must verify graceful degradation (no crashes, no data loss) | Reject if system crashes under stress |
| TG-PERF-010 | System must recover within 5 minutes after stress load is removed | Reject if recovery exceeds 5 minutes |
| TG-PERF-011 | Performance regression > 20% from baseline must be investigated and justified | CI/CD gate — alert on regression |
| TG-PERF-012 | UI Core Web Vitals must meet targets: LCP < 2.5s, FID < 100ms, CLS < 0.1 | CI/CD gate (Lighthouse CI) |
| TG-PERF-013 | Database slow queries (> 100ms) must be identified and optimized before load test | Reject if known slow queries exist |
| TG-PERF-014 | Performance test results must be retained and trended across releases | Process guardrail |
| TG-PERF-015 | Auto-scaling must be tested: verify scale-up within 60 seconds of trigger | Reject if auto-scaling untested |

### Pre-Release Performance Checklist
| # | Check |
|---|---|
| 1 | Baseline established for all customer-facing endpoints |
| 2 | Load test passed: p95 within target at peak load for 30+ minutes |
| 3 | Error rate < 0.1% during load test |
| 4 | Test environment is production-equivalent |
| 5 | Endurance test passed (4+ hours): no memory/connection leaks |
| 6 | Stress test: graceful degradation; recovery within 5 minutes |
| 7 | No performance regression > 20% from baseline |
| 8 | UI Core Web Vitals: LCP < 2.5s, FID < 100ms, CLS < 0.1 |
| 9 | Database slow queries optimized (none > 100ms) |
| 10 | Auto-scaling tested and verified |
| 11 | Server-side metrics monitored during tests (CPU, memory, connections) |
| 12 | Results retained and trended |