# Best Practices for Composing Compliance Requirements

---

## 1. Writing the Requirement ID & Title

**Best Practices:**
- Use consistent naming: `CR-[Regulation]-[Seq]` (e.g., CR-GDPR-001, CR-PCIDSS-005, CR-PSD2-003)
- Standard regulation codes: GDPR, PCIDSS, PSD2, SOX, BASEL, AML, DORA, LOCAL (for jurisdiction-specific)
- Titles should state the compliance obligation clearly: "Customer Consent for Data Processing" not "GDPR stuff"
- Include the regulatory article in the title when it adds clarity: "Right to Data Erasure (GDPR Art. 17)"

**Common Mistakes:**
- Vague IDs like CR-001 without regulation context
- Titles that name the regulation but not the obligation ("GDPR Requirement" — which one?)

---

## 2. Citing the Regulation Correctly

**Best Practices:**
- Always cite the specific article, section, or control number — not just the regulation name
- Include the version/year of the regulation if it has been updated
- For standards with numbered controls (PCI-DSS, NIST), cite the control ID
- If a local regulation implements an international standard, cite both
- Include the regulatory body that enforces it

**Citation Format:**
> [Regulation Name] [Version/Year] — [Article/Section/Control] — [Enforcing Body]

**Examples:**
- GDPR (2016/679) — Article 17 (Right to Erasure) — Enforced by national DPAs
- PCI-DSS v4.0 — Requirement 3.4.1 (Encryption of PAN) — Enforced by PCI SSC / acquiring banks
- PSD2 (2015/2366) — Article 97 (Strong Customer Authentication) — Enforced by EBA / national regulators
- SOX (2002) — Section 404 (Internal Controls) — Enforced by SEC
- Basel III — BCBS 239 (Risk Data Aggregation) — Enforced by national banking regulators

**Common Mistakes:**
- Citing "GDPR" without the article — GDPR has 99 articles covering very different obligations
- Using outdated regulation versions
- Not tracking regulatory amendments

---

## 3. Writing the Description & Compliance Obligation

**Best Practices:**
- Separate what the regulation says (description) from what the system must do (obligation)
- Description: paraphrase the regulatory text in plain language, explaining how it applies to this system
- Obligation: state the specific, actionable requirement the system must fulfill
- Use "must" for mandatory obligations, "should" for recommended practices
- Include the regulatory intent — why this obligation exists

**Formula for Description:**
> [Regulation] requires that [regulated entity] [obligation] in order to [regulatory intent]. This applies to [this system/feature] because [reason].

**Formula for Obligation:**
> The system must [specific action] [within specific timeframe] [for specific data/users] [with specific evidence].

**Good Example:**
> **Description:** GDPR Article 17 grants data subjects the right to have their personal data erased without undue delay. This applies to our customer data management system as it processes EU customer PII.
> **Obligation:** The system must anonymize all customer PII within 30 days of a verified erasure request, notify all downstream data processors, and retain an audit log of the erasure process. Data required for AML compliance (5-year retention) is exempt but must be documented with legal basis.

**Bad Example:**
> We need to comply with GDPR for data deletion.

---

## 4. Identifying Data Involved & Classification

**Best Practices:**
- List every data type subject to the compliance requirement
- Classify each data type using the organization's data classification scheme
- Map data types to their storage locations (database, file system, logs, backups, third-party systems)
- Identify data flows — where does this data travel? (channels, APIs, integrations, reports)
- Consider derived data and aggregated data — are they also in scope?

**Banking Data Classification:**

| Classification | Definition | Examples | Handling |
|---|---|---|---|
| Restricted | Highest sensitivity, regulatory protected | Card numbers, passwords, biometrics, encryption keys | Encrypted, tokenized, strict access, full audit |
| Confidential | Business-sensitive, customer PII | Account balances, transaction history, customer name/address/email | Encrypted at rest/transit, role-based access |
| Internal | Not for public, low sensitivity | Internal policies, system configurations, employee IDs | Access-controlled, no public exposure |
| Public | No sensitivity | Published interest rates, branch locations, product brochures | No restrictions |

**Data to Always Consider in Banking Compliance:**
- Customer PII (name, address, DOB, national ID, phone, email)
- Financial data (account numbers, balances, transaction history)
- Card data (PAN, CVV, expiry — PCI-DSS scope)
- Authentication data (passwords, tokens, biometrics)
- Communication records (emails, chat logs, call recordings)
- Audit logs and system logs
- Third-party data (credit bureau, sanctions lists)

---

## 5. Defining Geographic Scope

**Best Practices:**
- Specify exactly which jurisdictions the requirement applies to
- Consider where the data is collected, processed, stored, and transferred
- Account for cross-border data flows (EU → non-EU, onshore → offshore)
- Identify conflicting regulations across jurisdictions and document resolution
- Consider where the customer is located vs. where the bank is licensed

**Banking-Specific Considerations:**
- GDPR applies to EU/EEA data subjects regardless of where the bank is headquartered
- PCI-DSS applies globally wherever card data is processed
- Local banking regulations vary significantly (e.g., data localization requirements in some countries)
- Sanctions regulations (OFAC, EU sanctions) apply based on both customer and bank jurisdiction
- Some regulations have extraterritorial reach (US FATCA, EU GDPR)

---

## 6. Defining Implementation Approach

**Best Practices:**
- Describe the technical and process approach to meeting the obligation
- Be specific enough to guide development but not so prescriptive as to constrain design
- Address both the system implementation and the operational process
- Consider automation vs. manual processes — prefer automation for consistency
- Document trade-offs and design decisions with rationale
- Address conflicts between regulations (e.g., GDPR erasure vs. AML retention)

**Pattern:**
> **Technical:** [How the system implements the control]
> **Process:** [How the operational process supports the control]
> **Conflict Resolution:** [How conflicting obligations are handled]

**Example:**
> **Technical:** Implement soft-delete with field-level anonymization. Replace PII with hashed tokens. Retain transaction records with anonymized customer references for AML compliance.
> **Process:** Erasure requests received via customer portal or branch are routed to the data privacy team for identity verification. Automated workflow processes approved requests within 30 days.
> **Conflict Resolution:** AML 5-year retention takes precedence over GDPR erasure for transaction records. Customer is informed that certain records are retained under legal obligation, with specific legal basis documented per record.

---

## 7. Specifying Evidence & Audit Trail

**Best Practices:**
- Define exactly what evidence an auditor would need to verify compliance
- Specify the format, location, and retention period of evidence
- Ensure evidence is tamper-proof and independently verifiable
- Include both automated evidence (system logs) and manual evidence (sign-offs, reviews)
- Design evidence collection into the system from the start — retrofitting is expensive

**Evidence Types for Banking Compliance:**

| Evidence Type | Examples | Retention |
|---|---|---|
| System Logs | Access logs, transaction logs, change logs | 7 years (financial), 3 years (access) |
| Audit Reports | Compliance test results, pen test reports, audit findings | 7 years |
| Approvals | Sign-off records, approval workflows, exception approvals | 7 years |
| Configuration | Security configurations, policy documents, RBAC definitions | Current + 3 years history |
| Training | Training completion records, certification records | 3 years |
| Incident Records | Incident reports, RCA documents, remediation evidence | 7 years |

**Audit Trail Must Capture (for every compliance-relevant action):**
- Who: user ID, role, IP address
- What: action performed, data accessed/modified
- When: timestamp (UTC, synchronized)
- Where: system, component, endpoint
- Why: business justification or triggering event
- Outcome: success/failure, before/after values for modifications

---

## 8. Defining Reporting Requirements

**Best Practices:**
- Specify each regulatory report: name, content, frequency, format, recipient
- Include both scheduled reports and event-triggered reports
- Define report generation process (automated vs. manual)
- Include data quality checks before report submission
- Define escalation process for late or inaccurate reports

**Common Banking Regulatory Reports:**

| Report | Regulation | Frequency | Trigger |
|---|---|---|---|
| Currency Transaction Report (CTR) | BSA/AML | Per transaction | Transactions ≥ $10,000 |
| Suspicious Activity Report (SAR) | BSA/AML | Per event | Suspicious activity detected |
| Data Breach Notification | GDPR Art. 33 | Per event | Within 72 hours of awareness |
| PCI-DSS Compliance Report | PCI-DSS | Annual | Scheduled |
| Risk Data Aggregation Report | Basel III / BCBS 239 | Quarterly | Scheduled |
| ICT Incident Report | DORA Art. 19 | Per event | Major ICT incident |
| SOX Internal Controls Report | SOX Section 404 | Annual | Scheduled |

---

## 9. Writing Acceptance Criteria

**Best Practices:**
- Compliance acceptance criteria must prove the obligation is met, not just that the feature works
- Include evidence generation as part of the criteria
- Test both the happy path (compliant behavior) and violation path (system prevents non-compliance)
- Include timing requirements (e.g., "within 72 hours", "within 30 days")
- Include audit trail verification as a criterion

**Patterns:**
```
# Data Erasure (GDPR)
Given a verified customer erasure request,
When the erasure workflow completes,
Then all PII is anonymized within 30 days
  AND downstream processors are notified within 48 hours
  AND audit log captures: request date, verification date, erasure date, fields anonymized
  AND AML-exempt records are retained with documented legal basis
  AND customer receives confirmation of erasure.

# Breach Notification (GDPR)
Given a confirmed personal data breach,
When the breach is classified as reportable,
Then the supervisory authority is notified within 72 hours
  AND affected data subjects are notified without undue delay
  AND breach report includes: nature, categories of data, approximate number affected, consequences, measures taken
  AND all notifications are logged with timestamps.

# Transaction Monitoring (AML)
Given a completed financial transaction,
When the transaction matches suspicious activity patterns,
Then an alert is generated within 5 minutes
  AND the alert includes: transaction details, customer profile, pattern matched, risk score
  AND the alert is routed to the AML analyst queue
  AND the investigation timeline begins (SAR filing within 30 days if confirmed).
```

---

## 10. Data Classification in Compliance Requirements (ref: PM-BP-008)

**Best Practices:**
- Every CR must specify the data classification of the data it governs — this determines the control intensity
- Classification must align with the organization's data policy and the regulation's own data categorization
- Some regulations define their own categories (GDPR: personal data, special category data; PCI-DSS: cardholder data) — map these to the organization's classification scheme
- Classification drives: encryption requirements, access control stringency, audit depth, retention handling, breach notification urgency

**Regulation-to-Classification Mapping:**

| Regulation | Data Type | Organization Classification |
|---|---|---|
| GDPR | Personal data (name, email, phone) | Confidential |
| GDPR | Special category data (health, biometrics, race) | Restricted |
| PCI-DSS | Cardholder data (PAN, CVV, expiry) | Restricted |
| PCI-DSS | Sensitive authentication data (track data, PIN) | Restricted (must not be stored post-authorization) |
| AML/KYC | Customer due diligence data | Confidential |
| SOX | Financial reporting data | Confidential |
| Basel III | Risk data | Confidential |

---

## 11. DR/BCP in Compliance Requirements (ref: PM-BP-010)

**Best Practices:**
- Compliance obligations don't pause during disasters — CRs must specify how compliance is maintained during DR/BCP events
- Regulatory reporting deadlines remain in effect during outages — specify contingency reporting procedures
- Data protection obligations (encryption, access control, audit logging) must be maintained in DR environments
- Some regulations explicitly mandate DR/BCP (DORA Article 11, PCI-DSS Requirement 12.10, Basel BCBS 239)
- Specify how compliance evidence is preserved and accessible during and after DR events
- Define the compliance impact assessment process for DR events

**DR/BCP Compliance Considerations by Regulation:**

| Regulation | DR/BCP Requirement | What to Specify |
|---|---|---|
| DORA | ICT business continuity management (Art. 11) | Recovery plans, testing frequency, communication procedures |
| PCI-DSS | Incident response plan (Req 12.10) | Cardholder data protection during incidents, forensic readiness |
| GDPR | Data protection during processing (Art. 32) | Encryption and access control maintained in DR; breach notification if DR event causes data exposure |
| SOX | Internal controls continuity | Financial reporting controls maintained during DR; manual controls documented |
| Basel III | Risk data availability (BCBS 239) | Risk data aggregation capability during DR; reporting accuracy maintained |
| AML/KYC | Transaction monitoring continuity | Sanctions screening and transaction monitoring must not be interrupted |

---

## 12. Common Anti-Patterns to Avoid

| Anti-Pattern | Problem | Fix |
|---|---|---|
| "Comply with GDPR" | Too vague, GDPR has 99 articles | Cite specific article and obligation |
| No geographic scope | Unclear where it applies | Specify jurisdictions explicitly |
| No evidence requirements | Can't prove compliance to auditors | Define evidence from day one |
| Compliance as afterthought | Expensive retrofitting, audit failures | Compliance-by-design from inception |
| Ignoring regulation conflicts | Contradictory implementations | Document conflict resolution explicitly |
| No retention period | Data kept forever or deleted too soon | Specify retention per regulation |
| Manual-only compliance | Inconsistent, error-prone, unscalable | Automate wherever possible |
| No review frequency | Requirements become stale | Quarterly review + on regulatory change |
| Missing penalties | Stakeholders underestimate importance | Document fines and business impact |
| No data classification mapping | Controls misaligned with data sensitivity | Map regulation data types to org classification |
| Assuming compliance pauses during DR | Regulatory penalties during outages | Define DR compliance continuity procedures |
| Audit evidence not DR-protected | Evidence lost during disaster | Replicate evidence to DR site |
