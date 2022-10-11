Billing
=======

## Recurring Subscriptions

Recurring Billing

Platforms that have recurring billing:
- Recurly
- Chargify
- Stripe
- Braintree
- PayPal

Fundamental Models:
- _Customer_ is who is buying.
- _Plan_ or _Product_ is what a Customer is buying.
- _Subscription_ is the relationship state of Customer and a Plan.
- _Transaction_ or _Invoice_ is an Subscription related-event where money changed hands.

How do I specify how and when a subscription renews?

Recurly delegates this to a [Plan](https://dev.recurly.com/docs/create-plan):
- `plan_interval_unit` (days, months)
- `plan_interval_length`
- `total_billing_cycles`

Chargify delegate this to a [Product](https://reference.chargify.com/v1/products/create-a-product-1):
- `interval_unit` (days, months)
- `interval`
- `price_in_cents`

Chargify will version Products anytime there is a change.

[Create Subscription](https://dev.recurly.com/docs/create-subscription)

https://articles.braintreepayments.com/guides/recurring-billing/overview
