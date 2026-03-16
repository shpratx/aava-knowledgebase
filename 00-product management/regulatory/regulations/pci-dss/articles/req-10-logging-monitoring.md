# PCI-DSS Requirement 10 — Log and Monitor All Access

## Regulation Reference
- Standard: PCI-DSS v4.0
- Requirement: 10

## Obligation (Plain Language)
All access to system components and cardholder data must be logged. Audit trails must be secured and reviewed regularly.

## Technical Controls Required
1. **Log all access:** User access to cardholder data; all actions by admin/root; access to audit trails; invalid access attempts; authentication events; creation/deletion of system objects
2. **Log fields:** User ID, event type, date/time, success/failure, origination, identity/name of affected data/resource
3. **Time synchronization:** NTP across all systems; max 1-second skew
4. **Log protection:** Audit trails cannot be altered; write-once or append-only; access restricted
5. **Log retention:** Online: 3 months immediately available; Total: 12 months retained
6. **Log review:** Daily review of security events; automated alerting for anomalies
7. **File integrity monitoring:** Detect unauthorized changes to critical files and logs

## Evidence Required
- Log configuration showing all required events captured
- Log samples demonstrating required fields
- NTP configuration across all systems
- Log access control configuration (write-once, restricted access)
- Log retention policy and verification
- Daily log review procedures and records
- File integrity monitoring configuration and alerts

## Acceptance Criteria
Given a system component in the CDE,
When assessed for Req 10 compliance,
Then all access to cardholder data is logged with required fields
  AND time is synchronized via NTP (max 1-second skew)
  AND audit trails are write-once and access-restricted
  AND logs are retained: 3 months online, 12 months total
  AND security events are reviewed daily
  AND file integrity monitoring detects unauthorized changes.
