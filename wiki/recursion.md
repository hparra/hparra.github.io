---
---

Recursion
=========

```python
def fibonacci(n):
    if n < 2:
        return n
    else:
        return fibonacci(n-1) + fibonacci(n-2)
```

```python
def fibonacci(num):
    hashmap = {}

    def f(n):
        if n in hashmap:
            return hashmap[n]
        if n < 2:
            return n
        hashmap[n] = f(n-1) + f(n-2)
        return hashmap[n]

    return f(num)
```

- always solve base case(s)!
- how do you get next case from these bases?