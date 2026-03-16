# GDPR Article 20 — Right to Data Portability

## Regulation Reference
- Regulation: GDPR (2016/679) — Article 20
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Data subjects have the right to receive their personal data in a structured, commonly used, machine-readable format, and to transmit it to another controller. Applies only to data processed by automated means based on consent or contract.

## Scope
| Applies To | Does Not Apply To |
|---|---|
| Data provided by the data subject | Data inferred or derived by the controller |
| Processing based on consent (Art. 6(1)(a)) or contract (Art. 6(1)(b)) | Processing based on legal obligation, public interest, legitimate interest |
| Automated processing | Manual/paper processing |

## Banking Data in Scope
| Data | In Scope | Format |
|---|---|---|
| Transaction history | Yes (contractual) | CSV, JSON, OFX |
| Account details | Yes (contractual) | JSON |
| Customer profile | Yes (contractual) | JSON |
| Standing orders | Yes (contractual) | JSON |
| Credit score | No (derived/inferred) | N/A |
| Internal risk assessment | No (derived) | N/A |
| AML flags | No (legal obligation basis) | N/A |

## Technical Controls Required
1. Data export functionality: customer self-service download in structured format (JSON, CSV)
2. Machine-readable format: structured data, not PDF scans
3. Direct transfer: capability to transmit data directly to another controller (where technically feasible)
4. Open Banking alignment: PSD2 AISP access provides portability for account/transaction data
5. Timeline: respond within 1 month (same as Art. 15)
6. Free of charge

## Evidence Required
- Data export functionality (UI and API)
- Supported formats documentation
- Direct transfer capability (if implemented)
- Open Banking API compliance (PSD2 alignment)
- Response timeline tracking

## Acceptance Criteria
Given a data portability request,
When the request is processed,
Then the data subject's data is provided in structured, machine-readable format (JSON/CSV)
  AND only data provided by the data subject under consent/contract is included
  AND derived/inferred data is excluded
  AND response is provided within 1 month
  AND Open Banking APIs support ongoing portability for account data.

## Penalties
- Up to €20M or 4% of annual global turnover
