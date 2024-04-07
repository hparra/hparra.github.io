---
---

Bitwise Math
============


|  x  |  y  | x & y | x or y | x ^ y |
| --- | --- | ----- | ----- | ----- |
|  0  |  0  |   0   |   0   |   0   |
|  0  |  1  |   0   |   1   |   1   |
|  1  |  0  |   0   |   1   |   1   |
|  1  |  1  |   1   |   1   |   0   |


**XOR**

|  x  |  y  | x ^ y |
| --- | --- | ----- |
|  0  |  0  |   0   |
|  0  |  1  |   1   |
|  1  |  0  |   1   |
|  1  |  1  |   0   |

Always forgetting XOR?
- Recall _eXclusive OR_
- We want what is in each but not in both
- In other words:
  - true when their different
  - false when they're the same

Remember these facts regarding XOR:
- `a ^ a = 0`
- `a ^ !a = 1`

Checkout http://bits.stephan-brumme.com/