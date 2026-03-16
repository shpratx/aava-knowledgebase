# BCOBS 2 — Communications with Banking Customers

## Regulation Reference
- Sourcebook: BCOBS (FCA Handbook)
- Chapter: 2
- Enforcing Body: FCA

## Obligation (Plain Language)
Communications with banking customers must be clear, fair, and not misleading. This applies to all channels: website, mobile app, email, SMS, letters, in-branch.

## Technical Controls Required
1. **Plain language:** All customer-facing text reviewed for clarity; no jargon; reading level appropriate
2. **Prominence:** Key information (fees, rates, risks) must be prominent, not buried in small print
3. **Balance of information:** Benefits and risks presented with equal prominence
4. **Timeliness:** Information provided at the right point in the customer journey
5. **Channel consistency:** Same information across all channels (web, mobile, branch)
6. **Rate/fee changes:** Advance notice of changes (minimum per terms; typically 30-60 days)
7. **Marketing vs information:** Clear distinction; marketing must be identifiable as such
8. **Accessibility:** Communications accessible to all customers including those with disabilities

## Evidence Required
- Content review records (plain language assessment)
- UI/UX review showing prominence of key information
- Channel consistency audit
- Rate/fee change notification records with advance notice verification
- Accessibility audit results (WCAG 2.1 AA)

## Acceptance Criteria
Given a customer-facing communication (UI, email, notification),
When assessed against BCOBS 2,
Then the communication uses plain language without jargon
  AND key information (fees, rates, risks) is prominently displayed
  AND benefits and risks are presented with equal prominence
  AND the communication is accessible (WCAG 2.1 AA)
  AND rate/fee changes include advance notice per terms.
