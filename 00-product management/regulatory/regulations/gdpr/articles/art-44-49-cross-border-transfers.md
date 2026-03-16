# GDPR Articles 44-49 — Cross-Border Data Transfers

## Regulation Reference
- Regulation: GDPR (2016/679)
- Articles: 44-49
- Enforcing Body: National DPAs

## Obligation (Plain Language)
Personal data may only be transferred to countries outside the EU/EEA if adequate protection is ensured. This can be via adequacy decision, appropriate safeguards (SCCs, BCRs), or specific derogations.

## Transfer Mechanisms

| Mechanism | Article | When to Use | Banking Example |
|---|---|---|---|
| Adequacy decision | Art. 45 | Country deemed adequate by EU Commission | Transfers to UK, Japan, South Korea, etc. |
| Standard Contractual Clauses (SCCs) | Art. 46(2)(c) | Most common for non-adequate countries | Cloud provider in US; payment processor in India |
| Binding Corporate Rules (BCRs) | Art. 47 | Intra-group transfers in multinational | Bank's global operations across non-EU subsidiaries |
| Explicit consent | Art. 49(1)(a) | One-off transfers with informed consent | Customer-initiated international transfer |
| Contractual necessity | Art. 49(1)(b) | Transfer necessary to perform contract | Cross-border payment execution |

## Technical Controls Required
1. Data transfer inventory: map all cross-border data flows
2. Transfer Impact Assessment (TIA): assess destination country's data protection laws
3. SCCs execution: signed SCCs with each non-EU processor
4. Supplementary measures: encryption in transit and at rest; access controls preventing foreign government access
5. Data localization: comply with local data residency requirements where applicable
6. Regular review: annual review of transfer mechanisms and adequacy decisions

## Evidence Required
- Data transfer inventory (source, destination, data types, mechanism, legal basis)
- Signed SCCs or BCR approval
- Transfer Impact Assessments
- Supplementary measures documentation
- Annual review records

## Acceptance Criteria
Given a feature that transfers personal data outside the EU/EEA,
When the transfer is designed,
Then a valid transfer mechanism is identified and documented
  AND a Transfer Impact Assessment is completed for non-adequate countries
  AND SCCs are signed with the data importer (if applicable)
  AND supplementary measures are implemented (encryption, access control)
  AND the transfer is recorded in the data transfer inventory.

## Penalties
- Up to €20M or 4% of annual global turnover
