# BCBS 239 — Principles for Effective Risk Data Aggregation and Risk Reporting

## Regulation Reference
- Framework: Basel III
- Document: BCBS 239
- Enforcing Body: National banking regulators

## Obligation (Plain Language)
Banks must have strong data aggregation capabilities and risk reporting practices. Risk data must be accurate, complete, timely, and adaptable to changing reporting needs, especially during stress/crisis situations.

## Key Principles and Technical Controls

### Principle 2: Data Architecture and IT Infrastructure
| Control | Implementation |
|---|---|
| Single authoritative data source | Master data management; golden source per data domain |
| Data lineage | Track data from source to report; document transformations |
| Data dictionary | Standardized definitions across the organization |
| Automated data pipelines | ETL/ELT with validation; minimize manual intervention |
| Scalable infrastructure | Handle stress scenario calculations within reporting deadlines |

### Principle 3: Accuracy and Integrity
| Control | Implementation |
|---|---|
| Data quality rules | Automated validation at ingestion; completeness, consistency, accuracy checks |
| Reconciliation | Automated reconciliation between source systems and aggregation layer |
| Error handling | Automated detection; escalation; correction workflow |
| Data quality metrics | Accuracy rate, completeness rate, timeliness — tracked and reported |

### Principle 5: Timeliness
| Control | Implementation |
|---|---|
| Normal reporting | Risk reports produced within defined schedule (daily, weekly, monthly) |
| Stress/crisis reporting | Ad-hoc risk reports producible within hours (not days) |
| Automated reporting | Minimize manual steps; automated data aggregation and report generation |

## Evidence Required
- Data architecture documentation (golden sources, data flows, lineage)
- Data dictionary with standardized definitions
- Data quality metrics and trend reports
- Reconciliation reports
- Stress scenario reporting capability demonstration
- Automated pipeline documentation

## Acceptance Criteria
Given a risk reporting requirement,
When risk data is aggregated,
Then data is sourced from the authoritative golden source
  AND data lineage is traceable from source to report
  AND automated quality checks validate accuracy and completeness
  AND reconciliation confirms consistency with source systems
  AND reports are producible within defined timelines (including stress scenarios).
