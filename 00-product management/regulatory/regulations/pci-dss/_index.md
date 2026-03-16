# PCI-DSS — Payment Card Industry Data Security Standard

## Overview
| Field | Value |
|---|---|
| Full Name | Payment Card Industry Data Security Standard |
| Current Version | 4.0 (March 2022); mandatory compliance by March 2025 |
| Jurisdiction | Global — applies wherever card data is processed, stored, or transmitted |
| Enforcing Bodies | PCI Security Standards Council; enforced by acquiring banks and card brands |
| Penalties | Fines $5,000-$100,000/month; increased transaction fees; loss of card processing ability |
| Scope | Cardholder Data Environment (CDE) — all systems that store, process, or transmit cardholder data |

## Applicability to Banking
- Card issuance and management systems
- Payment processing and authorization
- ATM and POS systems
- Online/mobile card transactions
- Card data storage (PAN, expiry, cardholder name)
- Any system connected to the CDE

## Key Requirements for Banking

| Requirement | Title | Priority |
|---|---|---|
| Req 1 | Install and maintain network security controls | Critical |
| Req 2 | Apply secure configurations to all system components | Critical |
| Req 3 | Protect stored account data | Critical |
| Req 4 | Protect cardholder data with strong cryptography during transmission | Critical |
| Req 5 | Protect all systems and networks from malicious software | High |
| Req 6 | Develop and maintain secure systems and software | Critical |
| Req 7 | Restrict access to system components and cardholder data by business need to know | Critical |
| Req 8 | Identify users and authenticate access to system components | Critical |
| Req 9 | Restrict physical access to cardholder data | High |
| Req 10 | Log and monitor all access to system components and cardholder data | Critical |
| Req 11 | Test security of systems and networks regularly | Critical |
| Req 12 | Support information security with organizational policies and programs | High |

## Cardholder Data Elements

| Element | Storage Permitted | Protection Required | Render Unreadable |
|---|---|---|---|
| PAN (Primary Account Number) | Yes | Yes | Yes (encryption, hashing, truncation, tokenization) |
| Cardholder Name | Yes | Yes | No (but recommended) |
| Expiration Date | Yes | Yes | No (but recommended) |
| Service Code | Yes | Yes | No (but recommended) |
| Full Track Data | **No** (never after authorization) | N/A | N/A |
| CVV/CVC | **No** (never after authorization) | N/A | N/A |
| PIN / PIN Block | **No** (never after authorization) | N/A | N/A |
