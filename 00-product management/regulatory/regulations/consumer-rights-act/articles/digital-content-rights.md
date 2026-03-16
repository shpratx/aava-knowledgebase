# Consumer Rights Act 2015 — Digital Content Rights (s33-47)

## Regulation Reference
- Act: Consumer Rights Act 2015
- Sections: 33-47 (Chapter 3 — Digital Content)
- Enforcing Body: CMA / Trading Standards / FCA

## Obligation (Plain Language)
Digital content supplied to consumers must be of satisfactory quality, fit for purpose, and as described. If faulty, consumers have the right to repair, replacement, or price reduction.

## Applicability to Banking
- Mobile banking applications (iOS, Android)
- Online banking web application
- Banking APIs consumed by customers (Open Banking)
- Digital tools (calculators, budgeting tools, statements)
- Any digital service provided as part of the banking relationship

## Technical Controls Required
1. **Quality assurance:** Comprehensive testing (functional, performance, security, accessibility) before release
2. **Fitness for purpose:** Features work as documented; edge cases handled; error states managed
3. **As described:** UI matches marketing materials and product descriptions; no misleading feature claims
4. **Bug fix SLA:** Critical bugs fixed within 24 hours; high within 7 days; medium within 30 days
5. **Update mechanism:** Ability to push fixes to mobile apps; web updates deployed without customer action
6. **Fallback access:** If digital channel is unavailable, alternative access provided (phone, branch)
7. **Compatibility:** Supported on stated platforms/browsers; degradation handled gracefully

## Evidence Required
- Test reports (functional, performance, security, accessibility) per release
- Bug tracking records with SLA compliance
- Feature documentation matching UI implementation
- App store ratings and customer feedback monitoring
- Incident records and resolution times
- Platform compatibility test results

## Acceptance Criteria
Given a banking digital product (mobile app, online banking),
When assessed against Consumer Rights Act digital content rights,
Then the product is tested for quality before each release
  AND features work as described in documentation and marketing
  AND critical bugs are fixed within 24 hours
  AND the product is compatible with stated platforms
  AND alternative access is available if digital channel is down.
