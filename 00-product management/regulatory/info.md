## 1. Structure by Regulation → Obligation → Control → Evidence

The most effective pattern is a four-layer hierarchy:

Regulation (GDPR, PCI-DSS, PSD2, etc.)
  └── Obligation (specific article/requirement)
       └── Control (what the system must do)
            └── Evidence (how to prove compliance)


This maps directly to how auditors think — they start with the regulation, find the obligation, check the control, and ask for evidence.

## 2. Chunking Strategy

Rather than ingesting entire regulation PDFs as monolithic documents, break them into obligation-level chunks:

| Chunk Level | Content | Why |
|---|---|---|
| Regulation summary | Name, scope, applicability criteria, penalties, enforcing body | Quick lookup: "does this regulation apply to my feature?" |
| Article/requirement | Specific obligation text, plain-language interpretation, applicability conditions | Agent can match a feature to specific obligations |
| Control mapping | For each obligation: what technical/process controls satisfy it | Agent can generate design controls from requirements |
| Evidence catalog | For each control: what artifacts demonstrate compliance | Agent can generate test cases and audit checklists |
| Cross-reference index | Obligation → related obligations in other regulations | Agent can find all applicable rules for a given data type or operation |

## 3. Metadata Tagging

Every chunk should be tagged with structured metadata so agents can query precisely:

json
{
  "regulationId": "GDPR",
  "articleId": "Art. 17",
  "obligationTitle": "Right to Erasure",
  "keywords": ["erasure", "deletion", "right to be forgotten", "data subject rights"],
  "applicability": {
    "dataTypes": ["PII", "personal data"],
    "geographies": ["EU", "EEA"],
    "operations": ["data processing", "data storage"],
    "roles": ["data controller", "data processor"]
  },
  "relatedRegulations": ["AML Directive Art. 40 (retention conflict)"],
  "severity": "mandatory",
  "penalty": "Up to €20M or 4% global turnover",
  "lastReviewed": "2026-03-01",
  "version": "2016/679"
}


This lets agents do targeted retrieval: "find all obligations that apply when processing PII for EU customers in a payment context."

## 4. Recommended Knowledge Base Architecture

knowledge-base/
├── regulations/
│   ├── gdpr/
│   │   ├── _index.md              # Regulation overview, scope, penalties
│   │   ├── articles/
│   │   │   ├── art-06-lawfulness.md
│   │   │   ├── art-07-consent.md
│   │   │   ├── art-17-erasure.md
│   │   │   ├── art-25-data-protection-by-design.md
│   │   │   ├── art-30-processing-records.md
│   │   │   ├── art-32-security.md
│   │   │   └── art-33-breach-notification.md
│   │   └── controls/
│   │       ├── consent-management-controls.md
│   │       ├── erasure-controls.md
│   │       └── breach-notification-controls.md
│   ├── pci-dss/
│   │   ├── _index.md
│   │   ├── requirements/
│   │   │   ├── req-03-protect-stored-data.md
│   │   │   ├── req-04-encrypt-transmission.md
│   │   │   ├── req-06-secure-systems.md
│   │   │   └── req-10-track-access.md
│   │   └── controls/
│   ├── psd2/
│   ├── sox/
│   ├── aml-directive/
│   ├── dora/
│   └── basel-iii/
├── cross-references/
│   ├── by-data-type.md            # "PII" → GDPR Art.6, PCI-DSS Req.3, AML Art.40
│   ├── by-operation.md            # "payment" → PSD2 Art.97, PCI-DSS Req.3-4, AML Art.11
│   ├── by-role.md                 # "developer" → relevant obligations per role
│   └── conflicts-and-resolutions.md  # GDPR erasure vs AML retention
├── templates/
│   ├── compliance-requirement-template.md
│   ├── control-mapping-template.md
│   └── evidence-checklist-template.md
└── glossary/
    └── regulatory-terms.md        # Standardized definitions across regulations


## 5. Per-Article Document Format

Each article/requirement document should follow a consistent structure that agents can parse:

markdown
# GDPR Article 17 — Right to Erasure

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 17
- Enforcing Body: National DPAs
- Effective: 25 May 2018

## Obligation (Plain Language)
Data subjects have the right to request erasure of their personal data
without undue delay when [specific grounds apply].

## Applicability
- Applies when: processing EU/EEA data subject personal data
- Data types: all personal data categories
- Exceptions: legal obligation retention, public interest, legal claims

## Technical Controls Required
1. Erasure workflow: request → verify identity → assess exemptions → execute → notify downstream → confirm
2. Anonymization capability at field level
3. Downstream processor notification mechanism
4. Exemption documentation with legal basis per field
5. Audit trail of erasure process (without re-creating erased data)

## Timeline Requirements
- Complete within 30 days of verified request
- Notify downstream processors "without undue delay"

## Evidence Required
- Erasure request log with timestamps
- Identity verification record
- Exemption documentation with legal basis
- Downstream notification confirmations
- Completion confirmation to data subject

## Conflicts with Other Regulations
- AML Directive Art. 40: 5-year retention for transaction records
- Resolution: Anonymize customer reference; retain transaction with anonymized ID; document legal basis

## Acceptance Criteria (for agents generating test cases)
Given a verified erasure request,
When the erasure workflow completes,
Then all PII is anonymized within 30 days
  AND downstream processors notified
  AND audit log created
  AND AML-exempt records retained with documented legal basis
  AND data subject receives confirmation

## Penalties
- Up to €20M or 4% of annual global turnover (whichever is higher)


## 6. Practical Tips for Agent Consumption

Do:
- Keep each document focused on one obligation — agents retrieve better with specific chunks
- Use consistent headings across all documents — agents learn the structure
- Include the acceptance criteria in Given/When/Then — agents can directly generate test cases
- Include the control-to-evidence mapping — agents can generate audit checklists
- Maintain the cross-reference index — agents can find all applicable regulations for a feature
- Include plain-language interpretation alongside legal text — agents produce better output from clear language

Don't:
- Don't dump entire regulation PDFs as single documents — too large, too noisy for retrieval
- Don't use legal jargon without plain-language equivalent — agents produce better output from clear language
- Don't skip the metadata tags — without them, agents can't filter by data type, geography, or operation
- Don't mix multiple articles in one document — one obligation per document for precise retrieval
- Don't forget conflict resolution — agents need to know what to do when regulations contradict

## 7. Keeping It Current

| Activity | Frequency | Process |
|---|---|---|
| Regulatory change monitoring | Continuous | Subscribe to regulatory feeds (EBA, ICO, SEC, national DPAs) |
| Impact assessment | Within 48 hours of change | Identify affected articles, update documents, flag affected features |
| Full review | Quarterly | Review all articles for accuracy; update cross-references |
| Version control | Every change | Git-tracked with change history; diff shows what changed |

The key insight is that regulatory documentation works best as a knowledge base when it's structured around what the system must do (controls) and how to prove it (evidence), not just what the regulation says. Agents
generating requirements, designs, or test cases need actionable guidance, not legal text.