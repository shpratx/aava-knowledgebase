# Cross-Reference: Regulations by Banking Operation

## Customer Onboarding
| Regulation | Obligation |
|---|---|
| AML Directive Art. 11-14 | Customer Due Diligence (identity verification, sanctions screening, PEP check) |
| GDPR Art. 6 | Legal basis for data collection (legal obligation for KYC; contract for account) |
| GDPR Art. 13 | Privacy notice at point of collection |
| GDPR Art. 25 | Data minimization — collect only what's needed |
| GDPR Art. 7 | Consent for optional processing (marketing) |

## Payment / Fund Transfer
| Regulation | Obligation |
|---|---|
| PSD2 Art. 97 | Strong Customer Authentication (two factors) |
| AML Directive Art. 33 | Transaction monitoring; SAR filing if suspicious |
| AML Directive (BSA) | CTR for transactions >= $10,000 |
| GDPR Art. 6 | Legal basis (contractual necessity) |
| SOX Sec 404 | Internal controls; segregation of duties |
| PCI-DSS Req 3-4 | If card payment: encrypt PAN, protect in transit |

## Account Management
| Regulation | Obligation |
|---|---|
| GDPR Art. 15 | Right of access (customer can request their data) |
| GDPR Art. 17 | Right to erasure (with retention exemptions) |
| GDPR Art. 20 | Right to data portability |
| AML Directive Art. 40 | Retain account records 5 years post-closure |
| SOX Sec 404 | Audit trail for account modifications |

## Card Operations
| Regulation | Obligation |
|---|---|
| PCI-DSS Req 3 | Protect stored card data; never store CVV post-auth |
| PCI-DSS Req 4 | Encrypt card data in transit |
| PCI-DSS Req 7-8 | Access control; authentication for CDE |
| PCI-DSS Req 10 | Log all access to cardholder data |
| PSD2 Art. 97 | SCA for online card payments |

## Regulatory Reporting
| Regulation | Obligation |
|---|---|
| SOX Sec 302/404 | Financial reporting accuracy; internal controls |
| Basel III BCBS 239 | Risk data aggregation accuracy, completeness, timeliness |
| AML Directive Art. 33 | SAR filing to FIU |
| DORA Art. 17-19 | ICT incident reporting |
| GDPR Art. 33 | Breach notification to DPA within 72 hours |

## System Development / Change
| Regulation | Obligation |
|---|---|
| PCI-DSS Req 6 | Secure development; patching; code review |
| SOX Sec 404 | Change management for financial systems |
| DORA Art. 9 | Protection and prevention in ICT systems |
| GDPR Art. 25 | Data protection by design |
| DORA Art. 23-25 | Resilience testing |

## Disaster Recovery
| Regulation | Obligation |
|---|---|
| DORA Art. 11-12 | ICT business continuity; backup and restoration |
| PCI-DSS Req 12 | Incident response plan |
| SOX Sec 404 | Backup and recovery for financial data |
| Basel III BCBS 239 | Risk data available during stress scenarios |


## UK-Specific Cross-References by Operation

### Retail Banking (UK)
| Regulation | Obligation |
|---|---|
| FCA Consumer Duty | Good outcomes: products, price/value, understanding, support |
| BCOBS 2 | Clear, fair, not misleading communications |
| BCOBS 5 | Statements, annual summaries, dormant account management |
| Consumer Rights Act s33-47 | Digital content quality (banking app) |
| Consumer Rights Act s62-69 | Fair terms and conditions |
| Equality Act s20-21 | Accessible services; reasonable adjustments |
| DPA 2018 | Data protection (UK GDPR + UK exemptions) |

### Payments (UK)
| Regulation | Obligation |
|---|---|
| PSR CoP | Confirmation of Payee before payment execution |
| PSR APP Fraud | Fraud detection; victim reimbursement |
| FCA SYSC 15A | Operational resilience for payment services |
| PRA SS1/21 | Operational resilience (joint with FCA) |

### Mortgage Lending (UK)
| Regulation | Obligation |
|---|---|
| MCOB 11 | Responsible lending; affordability assessment |
| MCOB 13 | Arrears management; forbearance; repossession as last resort |
| FCA Consumer Duty | Good outcomes for mortgage customers |
| Equality Act | No discrimination in lending decisions |

### IT Systems / Operational Resilience (UK)
| Regulation | Obligation |
|---|---|
| FCA SYSC 15A | Important business services; impact tolerances; scenario testing |
| PRA SS1/21 | Same framework (joint policy) |
| PRA SS2/21 | Third-party/outsourcing risk management |
| PRA SS1/23 | Model risk management (AI/ML) |
