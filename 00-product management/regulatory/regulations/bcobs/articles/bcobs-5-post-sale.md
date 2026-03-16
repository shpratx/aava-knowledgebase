# BCOBS 5 — Post-Sale Requirements

## Regulation Reference
- Sourcebook: BCOBS (FCA Handbook)
- Chapter: 5
- Enforcing Body: FCA

## Obligation (Plain Language)
Firms must provide ongoing information to banking customers including regular statements, notification of material changes, and annual summaries. Dormant accounts must be managed appropriately.

## Technical Controls Required
1. **Regular statements:** Monthly or quarterly statements (per account type); available online and on request in paper
2. **Annual summary:** Annual summary of fees, interest paid/received, and account usage
3. **Material change notification:** Advance notice of rate changes, fee changes, terms changes
4. **Dormant account management:** Detect dormancy (no customer-initiated activity for 12+ months); attempt contact; manage per Dormant Assets Act
5. **Switching support:** Support Current Account Switch Service (CASS); 7-day switch guarantee
6. **Account closure:** Clear process; no unreasonable barriers; final statement

## Evidence Required
- Statement generation and delivery records
- Annual summary generation records
- Change notification records with advance notice verification
- Dormant account detection and contact records
- CASS switching records
- Account closure process records

## Acceptance Criteria
Given an active banking customer,
When post-sale obligations are assessed,
Then statements are generated and available per schedule
  AND annual summary is produced with fees, interest, and usage
  AND material changes are notified in advance
  AND dormant accounts are detected and managed per policy
  AND account switching is supported via CASS within 7 days.
