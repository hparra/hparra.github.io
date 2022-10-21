---
---

Bubble Sort
===========

## tl;dr

Bubble Sort:
- while array is not sorted:
  - assume it is sorted
  - iterate through array
    - if pair of elements out of order
      - swap them
      - mark array as unsorted

## Implementation

```python
def bubblesort(a):
  """Sort an array, `a` in ascending order"""
  is_sorted = False
  while not is_sorted:
    is_sorted = True
    for i in xrange(len(a) - 1):
      if a[i] > a[i+1]:
        a[i], a[i+1] = a[i+1], a[i]
        is_sorted = False
  return is_sorted
```