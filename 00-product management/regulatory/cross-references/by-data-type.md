# Cross-Reference: Regulations by Data Type

## Customer PII (Name, Address, DOB, Email, Phone)
| Regulation | Article | Obligation |
|---|---|---|
| GDPR | Art. 6 | Legal basis required for processing |
| GDPR | Art. 17 | Right to erasure (with AML exemption) |
| GDPR | Art. 25 | Data protection by design; minimization |
| GDPR | Art. 32 | Encryption at rest and in transit |
| AML Directive | Art. 11-14 | CDD: collect and verify identity |
| AML Directive | Art. 40 | Retain for 5 years post-relationship |
| PCI-DSS | Req 3 | Protect if associated with card data |

## Card Data (PAN, CVV, Expiry, Track Data)
| Regulation | Article | Obligation |
|---|---|---|
| PCI-DSS | Req 3 | Encrypt PAN; never store CVV/PIN post-auth; tokenize in non-CDE |
| PCI-DSS | Req 4 | Encrypt in transit (TLS 1.2+) |
| PCI-DSS | Req 10 | Log all access to cardholder data |
| GDPR | Art. 32 | Security of processing (encryption, access control) |

## Financial Transaction Data
| Regulation | Article | Obligation |
|---|---|---|
| SOX | Sec 404 | Internal controls; audit trail; segregation of duties |
| AML Directive | Art. 33 | Monitor for suspicious activity; file SARs |
| AML Directive | Art. 40 | Retain for 5 years |
| GDPR | Art. 6 | Legal basis (contractual necessity or legal obligation) |
| GDPR | Art. 17 | Erasure exemption for legal obligation retention |
| PSD2 | Art. 97 | SCA for payment initiation |
| Basel III | BCBS 239 | Accurate, complete, timely risk data aggregation |

## Authentication / Session Data
| Regulation | Article | Obligation |
|---|---|---|
| PSD2 | Art. 97 | SCA with two independent factors |
| PCI-DSS | Req 8 | Identify and authenticate access |
| GDPR | Art. 32 | Security of processing |
| DORA | Art. 9 | Protection and prevention measures |

## Audit Logs
| Regulation | Article | Obligation |
|---|---|---|
| PCI-DSS | Req 10 | Log all access; retain 12 months; review daily |
| SOX | Sec 404 | Audit trail for financial data; before/after values |
| SOX | Sec 802 | Criminal penalties for altering audit records |
| GDPR | Art. 30 | Processing records |
| DORA | Art. 10 | Detection capabilities |
| AML Directive | Art. 40 | Retain records for 5 years |

## ICT Systems / Infrastructure
| Regulation | Article | Obligation |
|---|---|---|
| DORA | Art. 5-7 | ICT risk management framework |
| DORA | Art. 11 | Response and recovery; DR/BCP |
| DORA | Art. 23-25 | Resilience testing |
| DORA | Art. 28-30 | Third-party ICT risk management |
| PCI-DSS | Req 1-2 | Network security; secure configuration |
| PCI-DSS | Req 6 | Secure development; patching |


## UK-Specific Cross-References by Data Type

### Customer PII (UK Context)
| Regulation | Obligation |
|---|---|
| DPA 2018 / UK GDPR | Same as EU GDPR + UK-specific exemptions (crime prevention, regulatory) |
| FCA Consumer Duty | Customer data used to deliver good outcomes; no dark patterns |
| Equality Act s20-21 | Reasonable adjustments for disabled customers accessing their data |
| BCOBS 2 | Communications about data must be clear, fair, not misleading |

### Payment Data (UK Context)
| Regulation | Obligation |
|---|---|
| PSR (CoP) | Confirmation of Payee name-checking before payment |
| PSR (APP Fraud) | Fraud detection; reimbursement for APP fraud victims |
| PSD2 / FCA | SCA for payment initiation |
| PCI-DSS | Card data protection (global, applies in UK) |

### Mortgage Data (UK Context)
| Regulation | Obligation |
|---|---|
| MCOB 11 | Affordability assessment data; income verification |
| MCOB 13 | Arrears data; forbearance records |
| DPA 2018 | Personal data protection for all mortgage data |
| FCA Consumer Duty | Fair treatment; good outcomes for mortgage customers |
