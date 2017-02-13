Quick Sort
==========

## tl;dr Review

Quicksort:
  1. _Pivot_
    - select a pivot index
  2. _Partition_:
    - swap values such that:
      - values < pivot are on left side of pivot
      - values > pivot are on right side of pivot
      - pivot is in final sorted location
  3. _Recurse_:
    - Quicksort each of the two partitions

Partition:
  1. move left pointer forward until you find a value > pivot
  2. move right pointer backward until you find a value < pivot
  3. now that you've found values on each side that actually belongs in the other:
    - swap them
    - move both pointers inwards
  4. repeat

Performance:
- Worst: O(n^2) e.g. pivot is first or last index
- Best: O(n lg n)
- Average: O(n lg n)

## Overview

Quicksort is a divide & conquer sorting algorithm.
Primary work is in divide phase: _partition_.

## Implementation

Parameters for both `quicksort` and `partition`:
- the `array`
- index at start of `left` side
- index at start of `right` side

Some implementations pass the `pivot` in as parameter.

Some implementations will refer to left and right as "lo and hi".
This works well for 3-way partitioning, e.g. lo, mid, hi.

**McDowell (2008)**

The following implementation is a Python rework of McDowell, 2015.
It is a two-way quicksort that always chooses the middle of the array as the pivot.

```python
def quicksort(array, left, right):
  index = partition(array, left, right)
  if left < index - 1:
    quicksort(array, left, index - 1)
  if index < right:
    quicksort(array, index, right)

def partition(array, left, right):
  pivot = array[(left + right) / 2]
  while left <= right:
    while array[left] < pivot:
      left += 1
    while array[right] > pivot:
      right -= 1
    if left <= right:
      a[left], a[right] = a[right], a[left]
      left += 1
      right -= 1
  return left
```

## REFERENCES

Sorting and Searching. Cracking the Coding Interview. Sixth Edition. Gayle Laakmann McDowell. 2015.

4.5 Mergesort: Sorting by Divide-and-Conquer. The Algorithm Design Manual. Steven S. Skiena. 2008.
