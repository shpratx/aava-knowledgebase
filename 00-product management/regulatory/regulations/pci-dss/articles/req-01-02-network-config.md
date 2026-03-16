# PCI-DSS Requirements 1-2 — Network Security and Secure Configuration

## Regulation Reference
- Standard: PCI-DSS v4.0 — Requirements 1 and 2

## Obligation
Req 1: Install and maintain network security controls (firewalls, segmentation). Req 2: Apply secure configurations to all system components (no defaults, hardening).

## Technical Controls
1. Network segmentation isolating CDE from non-CDE networks
2. Firewall rules: default-deny; explicit allow per business need; documented and reviewed semi-annually
3. DMZ for public-facing systems; no direct internet access to CDE
4. No vendor-supplied default passwords or settings in production
5. System hardening per CIS Benchmarks or equivalent
6. Unnecessary services/protocols disabled
7. Configuration standards documented per system type

## Acceptance Criteria
Given a system component in or connected to the CDE,
When assessed, Then network segmentation isolates CDE AND firewall rules are default-deny and documented AND no default passwords exist AND system is hardened per documented standards.
