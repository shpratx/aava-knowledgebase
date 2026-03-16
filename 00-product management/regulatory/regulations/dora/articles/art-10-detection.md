# DORA Article 10 — Detection

## Regulation Reference
- Regulation: DORA (2022/2554) — Article 10

## Obligation
Financial entities must have mechanisms to promptly detect anomalous activities, including ICT network performance issues, ICT-related incidents, and potential single points of failure.

## Technical Controls
1. SIEM: centralized security event monitoring and correlation
2. Anomaly detection: baseline behavior profiling; alert on deviations
3. Network monitoring: traffic analysis; DDoS detection; lateral movement detection
4. Application monitoring: APM; error rate monitoring; performance degradation detection
5. Log analysis: automated log analysis for security events
6. Threat intelligence: integration with threat feeds; IOC matching
7. Single point of failure identification: architecture review; dependency mapping
8. Alert management: severity classification; escalation procedures; SLA for response

## Acceptance Criteria
Given ICT systems, Then SIEM monitors all security events AND anomaly detection identifies deviations from baseline AND network and application monitoring is continuous AND alerts are classified by severity with defined response SLAs AND single points of failure are identified and documented.
