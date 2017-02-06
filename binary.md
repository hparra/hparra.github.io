Binary
======

> There are only 10 types of people in the world: those who understand binary, and those who don't. -- Unknown

## Converting Decimal to Binary

To convert decimal to binary:
1. take your decimal number, n
2. divide your number by 2 and save the dividend and remainder
3. make the dividend your new number
4. push the remainder onto a stack
5. repeat this process until your number is 0
6. pop all the numbers of the stack to create the binary number

Do this on paper. Stop to understand why the remainder will always be 1 or 0.

FYI: The above technique works with decimal to any base! Just change '2' to the base you want to covert to.

```python
# Returns a binary string
# you would never do this in a real program as Python has `bin(n)`
def decimal2binary(n):
  nba = [] # n as binary array
  div = n # dividend
  while div != 0:
      rem = div % 2
      div = div // 2
      nba.append(rem)
  # list comprehension to reverse array and join elements as a string
  return ''.join(str(i) for i in reversed(nba))
```

## REFERENCES

[Day 10: Binary Numbers](https://www.hackerrank.com/challenges/30-binary-numbers/tutorial). 30 Days of Code. HackerRank.

