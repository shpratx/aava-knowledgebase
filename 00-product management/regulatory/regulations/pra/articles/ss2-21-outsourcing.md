# PRA SS2/21 — Outsourcing and Third-Party Risk Management

## Regulation Reference
- Regulator: PRA
- Rule: SS2/21
- Effective: 31 March 2022

## Obligation (Plain Language)
Firms must manage risks from outsourcing and third-party dependencies. Material outsourcing requires enhanced governance, contractual protections, and exit planning.

## Technical Controls Required
1. **Third-party register:** Centralized register of all outsourcing arrangements with materiality classification
2. **Due diligence:** Security assessment, financial viability, regulatory compliance check before engagement
3. **Contractual controls:** Data protection clauses; audit rights; incident notification; exit provisions; sub-outsourcing approval
4. **Ongoing monitoring:** SLA tracking; security posture monitoring; incident tracking; periodic reassessment
5. **Exit planning:** Documented exit strategy for every material outsourcing; transition plan; data retrieval
6. **Concentration risk:** Map dependencies; assess impact of single-provider failure
7. **Cloud-specific:** Data residency; encryption; access control; portability; regulatory access

## Evidence Required
- Third-party register with materiality classification
- Due diligence records per provider
- Contracts with required clauses
- SLA monitoring reports
- Exit plans (tested for critical providers)
- Concentration risk assessment
- Board-level reporting on third-party risk

## Acceptance Criteria
Given a material outsourcing arrangement,
When assessed for SS2/21 compliance,
Then the arrangement is in the third-party register with materiality classification
  AND due diligence (security, financial, regulatory) is completed
  AND contract includes: data protection, audit rights, incident notification, exit provisions
  AND SLA monitoring is active with regular reporting
  AND an exit plan exists and is tested for critical providers.
