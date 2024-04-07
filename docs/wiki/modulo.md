---
---

modulo (%)
==========

The _modulo_ operation:
- `divdend % divisor = remainder`
- `divident mod divisor = remainder`
- `a % 1` is always 0 because `a / 1 = a`
- `a % a` is always 0 because `a / a = 1`
- `a % 0` is not defined because `a / 0` is not
- `a % b = c` where c will always in [0,b) when a > b
- `a % b = a` when b > a
  - "b divides a zero times, so everything remains"

You can check if something is odd by seeing if it is divisible by 2: `n % 2 != 0`

TODO: Examples and clever use:
- clock
- circular arrays (buffer)