# Compliance Requirements Template

---

## Template Fields

| Field | Description |
|---|---|
| **Requirement ID** | CR-[Regulation]-[Seq] (e.g., CR-GDPR-001) |
| **Title** | Short descriptive title |
| **Regulation / Standard** | GDPR / PCI-DSS / PSD2 / SOX / Basel III / AML-KYC / DORA / Local Banking Regulations |
| **Regulatory Article/Section** | Specific article or section reference |
| **Priority** | Critical / High / Medium / Low |
| **Related Functional Req(s)** | Linked FR IDs |
| **Description** | What the regulation requires and how it applies to this system |
| **Compliance Obligation** | Specific obligation (e.g., must obtain explicit consent, must report within 72 hours) |
| **Data Involved** | Types of data subject to this requirement |
| **Data Classification** | Public / Internal / Confidential / Restricted |
| **Geographic Scope** | Jurisdictions where this applies |
| **Implementation Approach** | How the system will satisfy this requirement |
| **Evidence Required** | Audit artifacts needed to demonstrate compliance (logs, reports, sign-offs) |
| **Audit Trail Requirements** | What must be logged and retained |
| **Retention Period** | How long data/evidence must be retained |
| **Reporting Requirements** | Regulatory reports, frequency, and format |
| **Penalties for Non-Compliance** | Fines, sanctions, or business impact |
| **Compliance Owner** | Role responsible for ongoing compliance |
| **Review Frequency** | How often this requirement should be reviewed |
| **Acceptance Criteria** | How compliance will be verified |
| **DR/BCP Compliance** | How compliance is maintained during disaster recovery and business continuity events; regulatory reporting obligations during outages (ref: PM-BP-010) |
| **Dependencies** | Other compliance or functional requirements |

---

## Regulation Quick Reference for Banking

| Regulation | Key Focus | Typical Requirements |
|---|---|---|
| **GDPR** | Data privacy | Consent, right to erasure, data portability, breach notification (72h), DPIAs |
| **PCI-DSS** | Payment card security | Encryption, access control, network segmentation, vulnerability management, logging |
| **PSD2** | Payment services | Strong Customer Authentication (SCA), open banking APIs, transaction monitoring |
| **SOX** | Financial reporting | Internal controls, audit trails, segregation of duties, data integrity |
| **Basel III** | Capital adequacy | Risk data aggregation, reporting accuracy, stress testing |
| **AML/KYC** | Anti-money laundering | Customer due diligence, transaction monitoring, suspicious activity reporting |
| **DORA** | Digital operational resilience | ICT risk management, incident reporting, resilience testing, third-party risk |

---

## Example — Compliance Requirement

| Field | Value |
|---|---|
| **Requirement ID** | CR-GDPR-004 |
| **Title** | Customer Right to Data Erasure |
| **Regulation / Standard** | GDPR |
| **Regulatory Article/Section** | Article 17 — Right to Erasure |
| **Priority** | Critical |
| **Related Functional Req(s)** | FR-CUS-021 (Customer Data Management) |
| **Description** | The system must support erasure of customer personal data upon verified request, except where retention is required by other regulations (e.g., AML 5-year retention) |
| **Compliance Obligation** | Erase personal data within 30 days of verified request; notify all downstream processors |
| **Data Involved** | Customer PII: name, address, email, phone, transaction history, communication logs |
| **Data Classification** | Restricted |
| **Geographic Scope** | EU/EEA customers and any data processed within EU |
| **Implementation Approach** | Soft-delete with anonymization; retain transaction records with anonymized identifiers for AML compliance; automated erasure workflow with approval |
| **Evidence Required** | Erasure request log, processing timestamps, confirmation of downstream notification, exception log for retained data with legal basis |
| **Audit Trail Requirements** | Log: request received, identity verified, erasure initiated, erasure completed, exceptions with justification |
| **Retention Period** | Erasure request logs retained for 3 years; AML-related data retained for 5 years post-relationship |
| **Reporting Requirements** | Monthly erasure request summary to DPO; annual report to supervisory authority if required |
| **Penalties for Non-Compliance** | Up to €20M or 4% of global annual turnover |
| **Compliance Owner** | Data Protection Officer (DPO) |
| **Review Frequency** | Quarterly |
| **Acceptance Criteria** | Given a verified erasure request, When processed, Then all PII is anonymized within 30 days AND downstream systems notified AND audit log created AND AML-exempt data retained with documented legal basis |
| **Dependencies** | CR-GDPR-001 (Consent Management), CR-AML-003 (Data Retention) |

---

## Usage Guidelines

1. **Every requirement must have a unique ID** following the CR-[Regulation]-[Seq] naming convention
2. **Regulatory article/section is mandatory** — cite the specific clause
3. **Geographic scope must be specified** — regulations vary by jurisdiction
4. **Evidence and audit trail requirements are mandatory** — these are what auditors will ask for
5. **Cross-reference** to related Functional (FR) and Security (SR) requirements
6. **Review cadence**: Quarterly + on regulatory change
7. **Approval**: Product Owner + Compliance Officer sign-off required
