# PCI-DSS Requirement 11 — Test Security of Systems and Networks Regularly

## Regulation Reference
- Standard: PCI-DSS v4.0 — Requirement 11

## Obligation
Regularly test security controls, systems, and processes to ensure they remain effective.

## Technical Controls
1. Wireless access point detection: quarterly scan for unauthorized wireless
2. Internal vulnerability scan: quarterly (by qualified personnel); rescan after remediation
3. External vulnerability scan: quarterly by ASV (Approved Scanning Vendor); rescan after remediation
4. Internal penetration test: annually and after significant changes
5. External penetration test: annually and after significant changes
6. Segmentation testing: every 6 months (or annually if segmentation not used)
7. File integrity monitoring (FIM): detect unauthorized changes to critical files; alert on change
8. IDS/IPS: monitor all traffic at CDE perimeter and critical points

## Acceptance Criteria
Given CDE systems, Then internal/external vulnerability scans run quarterly with remediation AND penetration tests run annually AND segmentation is tested every 6 months AND FIM detects and alerts on unauthorized changes AND IDS/IPS monitors CDE traffic.
