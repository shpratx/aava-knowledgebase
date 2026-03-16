# Epic Template

---

## Template Fields

| Field | Description |
|---|---|
| **Epic ID** | EP-[Domain]-[Seq] (e.g., EP-PAY-003, EP-LND-001, EP-ONB-005) |
| **Title** | Business capability-oriented title (ref: PM-BP-011) — describe what business outcome is delivered, not the technical implementation |
| **Business Domain** | Payments / Lending / Onboarding / Treasury / Cards / Customer Management / Reporting / Operations |
| **Initiative / Program** | Parent initiative or strategic program this epic belongs to |
| **Business Owner** | Stakeholder accountable for business outcomes |
| **Product Owner** | Person responsible for backlog prioritization and acceptance |
| **Priority** | Critical / High / Medium / Low |
| **Business Value Statement** | What business outcome this epic delivers — revenue, cost reduction, risk mitigation, regulatory compliance, customer experience |
| **Target Users / Personas** | Who benefits from this epic (e.g., Retail Customers, Relationship Managers, Operations Team) |
| **Description** | Detailed description of the business capability being delivered. Focus on what and why, not how. |
| **Scope — In** | What is explicitly included in this epic |
| **Scope — Out** | What is explicitly excluded (to prevent scope creep) |
| **Business Capabilities Delivered** | List of business capabilities this epic enables (ref: PM-BP-011) |
| **Risk Category** | Operational / Credit / Market / Compliance / Reputational / Technology (ref: PM-BP-014) |
| **Risk Assessment** | Impact (High/Medium/Low) × Likelihood (High/Medium/Low) = Risk Rating |
| **Risk Mitigation** | How identified risks will be mitigated |
| **Compliance Checkpoints** | Mandatory compliance gates within this epic's lifecycle (ref: PM-BP-012) |
| **Regulatory Requirements** | Applicable regulations with specific articles (ref: PM-BP-013) |
| **Regulatory Acceptance Criteria** | Specific regulatory conditions that must be met for epic completion (ref: PM-BP-013) |
| **Data Classification** | Highest data sensitivity level handled: Public / Internal / Confidential / Restricted (ref: PM-BP-008) |
| **Privacy Impact Assessment** | Required (Yes/No) — mandatory if epic involves customer PII |
| **Audit Trail Requirements** | What audit events this epic must generate across its features (ref: PM-BP-009) |
| **DR/BCP Requirements** | Recovery tier and business continuity needs for capabilities delivered (ref: PM-BP-010) |
| **Acceptance Criteria** | Business and regulatory conditions for epic completion (ref: PM-BP-013) |
| **Definition of Done** | Checklist of conditions that must all be true for the epic to be considered complete |
| **Dependencies** | Other epics, systems, teams, or external parties this depends on |
| **Assumptions** | Assumptions made during epic definition |
| **Constraints** | Technical, regulatory, budgetary, or timeline constraints |
| **Estimated Duration** | Target delivery timeline (max 3 months recommended) |
| **Features / User Stories** | List of child features or user stories (linked by ID) |
| **Success Metrics / KPIs** | How success will be measured post-delivery |
| **Stakeholder Sign-Off** | Required approvals: Business Owner, Compliance Officer, Security (for Restricted data), Architecture |

---

## Compliance Checkpoints Detail (ref: PM-BP-012)

Every epic must include the following compliance gates. Mark each as Required/Not Required based on epic scope:

| Checkpoint | When | Who Approves | Required? |
|---|---|---|---|
| Privacy Impact Assessment (PIA) | Before development starts | DPO / Privacy Officer | Required if PII involved |
| Regulatory Impact Assessment | Before development starts | Compliance Officer | Required if regulation-linked |
| Security Architecture Review | Before development starts | Security Architect | Required if Confidential/Restricted data |
| Data Classification Review | Before development starts | Data Owner + Compliance | Always required |
| Compliance Test Plan Review | Before testing phase | Compliance Officer | Required if regulation-linked |
| Security Test Results Review | Before production deployment | Security Engineer | Always required |
| Compliance Sign-Off | Before production deployment | Compliance Officer | Required if regulation-linked |
| Post-Deployment Compliance Verification | Within 2 weeks of deployment | Compliance Officer | Required if regulation-linked |

---

## Risk Category Detail (ref: PM-BP-014)

| Risk Category | Description | Banking Examples | Typical Mitigation |
|---|---|---|---|
| **Operational** | Risk of loss from failed processes, people, or systems | System outage, data entry error, process failure, vendor failure | DR/BCP, automation, monitoring, SLAs |
| **Credit** | Risk of loss from borrower default | Incorrect credit scoring, limit miscalculation | Validation rules, dual approval, testing |
| **Market** | Risk of loss from market movements | Incorrect pricing, FX exposure, rate calculation errors | Real-time validation, reconciliation |
| **Compliance** | Risk of regulatory penalties or sanctions | GDPR violation, AML failure, PCI-DSS breach | Compliance checkpoints, audit trails, testing |
| **Reputational** | Risk of damage to brand or customer trust | Data breach, service outage, poor UX | Security controls, performance testing, UX review |
| **Technology** | Risk from technology failures or obsolescence | Legacy integration failure, scalability limits, security vulnerability | Architecture review, tech debt management, patching |

---

## Example — Epic

| Field | Value |
|---|---|
| **Epic ID** | EP-PAY-003 |
| **Title** | Digital Payments Modernization — Domestic Fund Transfers |
| **Business Domain** | Payments |
| **Initiative / Program** | Digital Banking Transformation 2026 |
| **Business Owner** | Head of Retail Payments |
| **Product Owner** | Product Manager — Digital Channels |
| **Priority** | Critical |
| **Business Value Statement** | Enable retail and premium customers to initiate domestic fund transfers via mobile and web banking with real-time processing, reducing branch dependency by 40% and improving customer satisfaction (NPS +15). |
| **Target Users / Personas** | Retail Customers, Premium Customers, Branch Tellers (fallback) |
| **Description** | Deliver end-to-end domestic fund transfer capability across digital channels, including beneficiary management, real-time fraud screening, regulatory compliance (PSD2 SCA, AML transaction monitoring), and integration with core banking for real-time settlement. |
| **Scope — In** | Domestic transfers (same currency), beneficiary CRUD, real-time and scheduled transfers, fraud screening, transaction limits, notifications, audit trail |
| **Scope — Out** | International transfers (separate epic EP-PAY-004), bulk/batch transfers, standing orders |
| **Business Capabilities Delivered** | 1) Self-service domestic fund transfer 2) Beneficiary management 3) Real-time fraud screening 4) Transaction limit management 5) Transfer scheduling |
| **Risk Category** | Operational, Compliance |
| **Risk Assessment** | Impact: High × Likelihood: Medium = High Risk |
| **Risk Mitigation** | Real-time fraud screening, MFA for all transfers, daily limit enforcement, comprehensive audit trail, DR Tier 1 classification |
| **Compliance Checkpoints** | PIA: Required (customer PII), Regulatory Impact: Required (PSD2, AML), Security Architecture Review: Required (Restricted data), Compliance Sign-Off: Required |
| **Regulatory Requirements** | PSD2 Art. 97 (SCA), AML Directive Art. 11 (Transaction Monitoring), GDPR Art. 6 (Lawful Processing), GDPR Art. 30 (Processing Records) |
| **Regulatory Acceptance Criteria** | 1) All transfers require SCA per PSD2 2) Transactions ≥ $10,000 generate CTR 3) Suspicious patterns trigger SAR workflow 4) All customer data processing has documented legal basis 5) Audit trail retained for 7 years |
| **Data Classification** | Restricted (account numbers, transaction amounts, customer PII) |
| **Privacy Impact Assessment** | Yes — involves customer PII and financial transaction data |
| **Audit Trail Requirements** | All transfer attempts (success/failure), beneficiary changes, limit modifications, fraud alerts, approval decisions. Fields: user ID, timestamp, action, account details, amount, outcome, correlation ID. Retention: 7 years. |
| **DR/BCP Requirements** | Tier 1 — RTO < 1 hour, RPO < 15 minutes. In-flight transactions must complete or safely rollback. Degraded mode: queue transfers if core banking unavailable, hold if fraud engine unavailable. |
| **Acceptance Criteria** | 1) Customers can transfer funds domestically via mobile and web 2) All transfers complete within 5 seconds (p95) 3) Fraud screening runs on 100% of transfers 4) Daily limits enforced per customer tier 5) Full audit trail generated for every transfer 6) All regulatory requirements met and evidenced |
| **Definition of Done** | ✅ All features developed and code-reviewed ✅ Unit test coverage ≥ 80% ✅ Integration tests passing ✅ Security scan (SAST/DAST) — zero critical/high findings ✅ Performance test — p95 < 500ms at 500 concurrent users ✅ Compliance sign-off obtained ✅ DR failover tested ✅ Audit trail verified ✅ User acceptance testing passed ✅ Deployed to production ✅ Post-deployment monitoring active |
| **Dependencies** | EP-AUTH-001 (MFA Infrastructure), Core Banking API availability, Fraud Detection Engine integration, Payment Network connectivity |
| **Assumptions** | Core banking supports real-time settlement API; fraud engine can process within 200ms; payment network available 99.95% |
| **Constraints** | Must comply with PSD2 SCA deadline; max budget $500K; delivery within Q2 2026 |
| **Estimated Duration** | 10 weeks |
| **Features / User Stories** | FR-PAY-008 (Beneficiary Mgmt), FR-PAY-012 (Fund Transfer), FR-PAY-015 (Transfer Scheduling), FR-PAY-018 (Transaction Limits), FR-PAY-020 (Transfer Notifications) |
| **Success Metrics / KPIs** | Digital transfer adoption: 60% within 3 months; Branch transfer reduction: 40%; Transfer success rate: > 99.5%; NPS improvement: +15; Zero compliance findings |
| **Stakeholder Sign-Off** | Business Owner: Head of Retail Payments ✅ Compliance: Compliance Officer ✅ Security: Security Architect ✅ Architecture: Enterprise Architect ✅ |

---

## Usage Guidelines

1. **Structure around business capabilities, not technical components** (PM-BP-011) — "Digital Payments Modernization" not "Build REST APIs for Payment Service"
2. **Include compliance checkpoints** (PM-BP-012) — mark each checkpoint as Required/Not Required and schedule them in the delivery plan
3. **Define regulatory acceptance criteria** (PM-BP-013) — specific, testable conditions tied to regulation articles
4. **Map to risk categories** (PM-BP-014) — assess impact × likelihood and define mitigation
5. **Max duration: 3 months** — longer epics should be split for regulatory agility
6. **Every epic needs a Definition of Done** — including security, compliance, performance, and DR verification
7. **Stakeholder sign-off is mandatory** — Business Owner + Compliance Officer at minimum; Security Architect for Restricted data
