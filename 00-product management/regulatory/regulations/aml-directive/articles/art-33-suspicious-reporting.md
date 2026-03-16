# AML Directive Article 33 — Suspicious Transaction Reporting

## Regulation Reference
- Directive: 4AMLD (2015/849)
- Article: 33
- Enforcing Body: National FIUs

## Obligation (Plain Language)
Banks must report suspicious transactions to the national Financial Intelligence Unit (FIU). The bank must not inform the customer that a report has been made (tipping-off prohibition).

## Technical Controls Required
1. Transaction monitoring engine: real-time rule-based + ML-based anomaly detection
2. Alert generation: suspicious patterns trigger alerts within 5 minutes
3. Alert investigation workflow: analyst queue, investigation tools, decision recording
4. SAR filing: automated SAR generation with required fields; submission to FIU
5. Tipping-off prevention: no customer notification of SAR; restricted access to SAR data
6. CTR filing: Currency Transaction Reports for transactions >= $10,000 (or local equivalent)
7. Record retention: SAR records retained for 5 years

## Common Suspicious Patterns
| Pattern | Description |
|---|---|
| Structuring | Multiple transactions just below reporting threshold |
| Rapid movement | Funds received and immediately transferred out |
| Unusual geography | Transactions to/from high-risk jurisdictions |
| Inconsistent profile | Transaction patterns inconsistent with customer profile |
| Round amounts | Large round-number transactions without business justification |
| Third-party funding | Account funded by unrelated third parties |

## Evidence Required
- Transaction monitoring rules and configuration
- Alert generation and investigation records
- SAR filing records with FIU acknowledgment
- CTR filing records
- Tipping-off prevention controls
- Staff training records on suspicious activity identification

## Acceptance Criteria
Given a completed financial transaction,
When the transaction matches suspicious activity patterns,
Then an alert is generated within 5 minutes
  AND the alert is routed to the AML analyst queue
  AND if confirmed suspicious, a SAR is filed with the FIU within 30 days
  AND the customer is NOT notified of the SAR (tipping-off prevention)
  AND all records are retained for 5 years.
