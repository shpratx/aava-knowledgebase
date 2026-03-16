# GDPR Article 6 — Lawfulness of Processing

## Regulation Reference
- Regulation: GDPR (2016/679)
- Article: 6
- Title: Lawfulness of processing
- Enforcing Body: National DPAs
- Effective: 25 May 2018

## Obligation (Plain Language)
Every processing of personal data must have a valid legal basis. There are six legal bases; at least one must apply before any data is collected or processed.

## Legal Bases

| Basis | Art. 6(1) | Banking Use Case | When to Use |
|---|---|---|---|
| Consent | (a) | Marketing emails, analytics, third-party sharing | Optional processing only; never for core banking services |
| Contract | (b) | Account management, transaction processing, statements | Core banking services necessary to fulfill the account contract |
| Legal obligation | (c) | KYC/AML data collection, tax reporting, regulatory reporting | When law requires the processing (AML Directive, tax law) |
| Vital interests | (d) | Emergency fraud prevention | Rarely applicable in banking |
| Public interest | (e) | Regulatory reporting to authorities | When processing is necessary for public interest task |
| Legitimate interest | (f) | Fraud detection, security monitoring, service improvement | When balanced against data subject rights; requires LIA |

## Technical Controls Required
1. Legal basis must be determined and documented BEFORE processing begins
2. Processing activity register (Art. 30) must record the legal basis per activity
3. Consent mechanism must meet Art. 7 requirements (freely given, specific, informed, unambiguous)
4. Consent records must capture: timestamp, purpose, mechanism, privacy notice version
5. Consent withdrawal must be as easy as granting and effective immediately
6. Legal basis must be communicated to data subjects via privacy notice (Art. 13/14)

## Applicability
- **Applies when:** Any personal data is collected, stored, processed, or shared
- **Data types:** All personal data (name, email, phone, account data, transaction data, IP addresses)
- **Geographies:** EU/EEA data subjects regardless of where processing occurs

## Evidence Required
- Processing activity register with legal basis per activity
- Consent records (timestamp, purpose, mechanism, version)
- Privacy notices showing legal basis disclosure
- Legitimate Interest Assessments (LIAs) where Art. 6(1)(f) is used
- Data Protection Impact Assessments where required

## Acceptance Criteria
Given a new data processing activity,
When the activity is designed,
Then a legal basis is documented in the processing register
  AND the legal basis is communicated in the privacy notice
  AND if consent-based, the consent mechanism meets Art. 7 requirements
  AND if legitimate interest, an LIA is completed and documented.

## Penalties
- Up to €20M or 4% of annual global turnover
