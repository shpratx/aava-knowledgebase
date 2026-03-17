# Performance Testing — Template, Best Practices & Guardrails
### Banking Domain — Agentic Knowledge Base

---

## Part A: Test Scenario Template

### Metadata
| Field | Description |
|---|---|
| **Scenario ID** | TS-PERF-[Type]-[Seq] (e.g., TS-PERF-LOAD-001) |
| **Test Type** | Load / Stress / Endurance / Spike / Scalability / Baseline |
| **Target** | API endpoint / UI page / End-to-end flow / Database query |
| **NFR Reference** | Linked NFR ID with target metrics |
| **Tool** | Gatling / k6 / JMeter / Lighthouse / Artillery |
| **Environment** | Staging (production-equivalent infrastructure and data volume) |

### Performance Test Types

#### 1. Baseline Test
| Aspect | Standard |
|---|---|
| Purpose | Establish performance baseline for each endpoint/page |
| Load | Single user; no concurrency |
| Duration | 5 minutes |
| Metrics | Response time (min, avg, p50, p95, p99, max); throughput; error rate |
| When | Before every release; after significant code changes |

#### 2. Load Test
| Aspect | Standard |
|---|---|
| Purpose | Verify system meets NFR targets at expected peak load |
| Load | Expected peak concurrent users/TPS (e.g., 500 concurrent users) |
| Ramp-up | Gradual ramp over 5 minutes to target load |
| Duration | Sustained for 30 minutes at peak |
| Metrics | p95 response time; p99 response time; throughput; error rate; resource utilization |
| Pass criteria | p95 < NFR target (e.g., 500ms); error rate < 0.1%; no resource exhaustion |

#### 3. Stress Test
| Aspect | Standard |
|---|---|
| Purpose | Find breaking point; verify graceful degradation |
| Load | Incrementally increase beyond expected peak (1x → 1.5x → 2x → 3x) |
| Duration | 10 minutes per increment |
| Metrics | Point where p95 exceeds target; point where errors begin; recovery behavior |
| Pass criteria | System degrades gracefully (no crashes, no data loss); recovers when load reduces |

#### 4. Endurance (Soak) Test
| Aspect | Standard |
|---|---|
| Purpose | Detect memory leaks, connection leaks, resource exhaustion over time |
| Load | Normal expected load (not peak) |
| Duration | 4-8 hours minimum |
| Metrics | Memory usage trend; connection pool usage; response time trend; GC behavior |
| Pass criteria | No upward trend in memory/connections; response time stable; no OOM errors |

#### 5. Spike Test
| Aspect | Standard |
|---|---|
| Purpose | Verify system handles sudden traffic spikes (e.g., salary day, market event) |
| Load | Sudden jump from normal to 5x-10x normal |
| Duration | Spike for 5 minutes; return to normal; observe recovery |
| Metrics | Response time during spike; error rate; auto-scaling response time; recovery time |
| Pass criteria | System handles spike (possibly degraded); recovers within 5 minutes of spike end |

#### 6. Scalability Test
| Aspect | Standard |
|---|---|
| Purpose | Verify horizontal scaling works; measure scaling efficiency |
| Load | Incrementally increase load while adding instances |
| Metrics | Throughput per instance; response time as instances scale; scaling trigger time |
| Pass criteria | Near-linear throughput increase with instances; auto-scale triggers within 60 seconds |

---

### Banking Performance Targets

| Endpoint Type | p95 Target | p99 Target | Throughput | Error Rate |
|---|---|---|---|---|
| Account balance (GET) | < 200ms | < 500ms | 1000 TPS | < 0.01% |
| Transaction list (GET) | < 500ms | < 1000ms | 500 TPS | < 0.01% |
| Transfer initiation (POST) | < 500ms | < 1000ms | 500 TPS | < 0.1% |
| Login + MFA | < 2000ms | < 3000ms | 200 TPS | < 0.1% |
| Statement download (GET) | < 3000ms | < 5000ms | 50 TPS | < 0.1% |
| Search (GET) | < 1000ms | < 2000ms | 200 TPS | < 0.1% |
| UI page load (FCP) | < 1500ms | < 2500ms | — | — |
| UI page load (LCP) | < 2500ms | < 4000ms | — | — |
| UI interaction (FID) | < 100ms | < 300ms | — | — |

### Test Case Template

| Field | Description |
|---|---|
| **Test Case ID** | TC-PERF-[Type]-[Seq] |
| **Scenario** | Parent scenario ID |
| **Endpoint/Page** | Target URL |
| **Virtual Users** | Concurrent users or TPS |
| **Ramp-up** | Time to reach target load |
| **Duration** | Sustained load duration |
| **Think Time** | Pause between user actions (realistic: 3-10 seconds) |
| **Test Data** | Data set (number of accounts, transactions, users) |
| **Pass Criteria** | Specific thresholds (p95 < Xms, error rate < Y%) |
| **Infrastructure** | CPU, memory, instances, database size |

### Example: Transfer API Performance Tests

**TC-PERF-LOAD-001: Transfer API — Load Test at Peak**
```
Endpoint:     POST /v1/transfers
Virtual Users: 500 concurrent
Ramp-up:      5 minutes (0 → 500 users)
Duration:     30 minutes sustained
Think Time:   5 seconds between transfers
Test Data:    10,000 accounts; 5,000 beneficiaries; production-equivalent DB volume
Pass Criteria:
  - p95 response time < 500ms
  - p99 response time < 1000ms
  - Error rate < 0.1%
  - No connection pool exhaustion
  - No memory growth trend
  - CPU < 80% average
```

**TC-PERF-STRESS-001: Transfer API — Stress Test**
```
Endpoint:     POST /v1/transfers
Load Profile: 500 → 750 → 1000 → 1500 → 2000 concurrent users (10 min each)
Pass Criteria:
  - Identify breaking point (where p95 > 1000ms or error rate > 1%)
  - System degrades gracefully (no crashes, no data corruption)
  - System recovers within 5 minutes when load reduces to 500
  - Circuit breakers activate appropriately
  - Rate limiting engages correctly
```

**TC-PERF-ENDURANCE-001: Transfer API — Soak Test**
```
Endpoint:     POST /v1/transfers
Virtual Users: 200 concurrent (normal load)
Duration:     8 hours
Pass Criteria:
  - No upward trend in memory usage
  - No upward trend in response time
  - Connection pool stable (no leaks)
  - No GC pauses > 500ms
  - Error rate stable (no increase over time)
```

**TC-PERF-UI-001: Dashboard Page — Core Web Vitals**
```
Page:         /dashboard (authenticated)
Tool:         Lighthouse CI
Conditions:   Mobile (4G throttled); Desktop (no throttle)
Pass Criteria:
  - FCP < 1.5s
  - LCP < 2.5s
  - FID < 100ms
  - CLS < 0.1
  - TTI < 3s
  - Lighthouse Performance Score ≥ 90
```

---