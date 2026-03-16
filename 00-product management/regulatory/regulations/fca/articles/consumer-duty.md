# FCA Consumer Duty (PRIN 2A / PS22/9)

## Regulation Reference
- Regulator: FCA
- Rule: PRIN 2A (Consumer Duty)
- Policy Statement: PS22/9
- Effective: 31 July 2023 (existing products); 31 July 2024 (closed products)

## Obligation (Plain Language)
Firms must act to deliver good outcomes for retail customers. This is an overarching obligation that applies across all customer interactions, products, and services.

## Four Outcomes

### Outcome 1: Products and Services
| Control | Implementation |
|---|---|
| Target market definition | Product configured with eligibility criteria; system enforces |
| Product governance | Regular review of product outcomes; data-driven assessment |
| Distribution strategy | Appropriate channels for target market |
| Vulnerability consideration | Identify and support vulnerable customers |

### Outcome 2: Price and Value
| Control | Implementation |
|---|---|
| Fee transparency | All fees displayed clearly before commitment; no hidden charges |
| Value assessment | Regular assessment that price provides fair value |
| Differential pricing | Monitor for unfair pricing differences between customer groups |
| Charges disclosure | Total cost of product clearly shown in UI |

### Outcome 3: Consumer Understanding
| Control | Implementation |
|---|---|
| Clear communications | Plain language; no jargon; reading level appropriate for audience |
| Key information prominence | Important terms/risks prominently displayed, not buried |
| Timely information | Information provided at the right time in the customer journey |
| Testing comprehension | Test that customers understand key information (A/B testing, user research) |

### Outcome 4: Consumer Support
| Control | Implementation |
|---|---|
| Accessible channels | Multiple support channels; accessible to all including vulnerable customers |
| No unreasonable barriers | Cancellation as easy as sign-up; no friction to switch/leave |
| Timely resolution | Complaints resolved promptly; SLA tracked |
| Post-sale support | Ongoing support throughout product lifecycle |

## Technical Controls Required
1. Product eligibility engine enforcing target market criteria
2. Fee calculator showing total cost before commitment
3. Plain language content management; readability scoring
4. Accessibility compliance (WCAG 2.1 AA) across all customer touchpoints
5. Complaint management system with SLA tracking
6. Customer outcome monitoring dashboards
7. Vulnerability flags in customer profiles (with appropriate handling)
8. Cancellation/switching journeys with no more friction than sign-up

## Evidence Required
- Product governance records (target market, value assessment)
- Customer outcome data (complaints, satisfaction, product usage)
- Communications testing results (comprehension, clarity)
- Support channel accessibility audit
- Cancellation journey analysis (friction comparison with sign-up)
- Vulnerability identification and support records

## Acceptance Criteria
Given a retail banking product,
When assessed against Consumer Duty,
Then the product has a defined target market with system-enforced eligibility
  AND all fees are transparently displayed before customer commitment
  AND communications use plain language tested for comprehension
  AND support channels are accessible (WCAG 2.1 AA) with no unreasonable barriers
  AND cancellation is as easy as sign-up
  AND customer outcomes are monitored and reported.

## Penalties
- Unlimited fines; public censure; permission variation/cancellation
