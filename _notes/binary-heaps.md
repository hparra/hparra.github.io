---
---

Binary Heaps
============

_Read [heaps](heaps.md) first._

## tl;dr

Implementations:
- python: [`heapq`](https://docs.python.org/2/library/heapq.html)
- java: [`PriorityQueue`](https://docs.oracle.com/javase/7/docs/api/java/util/PriorityQueue.html)
- c++: [`priority_queue`](http://www.cplusplus.com/reference/queue/priority_queue/)
- ruby: NONE

## Introduction

A binary heap uses a complete binary tree.

Recall that _complete binary trees_:
- have each level filled, with possible exception of last level
- when the last level is incomplete all nodes are filled from the left

**Inserting into a binary heap**:
1. place new element in left most spot (n+1)
2. "Bubble Up": if and while (new) element dominates parent
  - swap them

Swaps happen between levels, and a tree will have lg(n) levels.
Since there are n items to be inserted, insertion will take `O(nlogn)`.

**Extracting the dominator from a binary heap**:
1. remove dominating element from top
2. move last added element (bottom-right most leaf) into top
3. "Bubble Down": if and while that element does not dominate its children:
  - swap it with lesser of two children

### Implementing Binary Heaps

Since a binary heap is a complete binary tree we can implement it using an array.

Example: `2 7 8 1 5 9 3` pushed into min-heap as a tree:
```
    1
  2   3
 7 5 9 8
```

But as an array looks like: `1,2,3,7,5,9,8,`.
Notice this is equivalent to a breadth-first traversal.

This image from Wikipedia explains best:
![Implicit Binary Tree](https://upload.wikimedia.org/wikipedia/commons/8/86/Binary_tree_in_array.svg)

So for some zero-based index `i`:
  - left child = `2i + 1`
  - right child = `2i + 2`
  - parent = `floor((i-1) / 2)`

_WARNING: Beware zero and one-based version of the above equations.
Many references use one-based equations because the math/logic is cleaner,
but programmers always use zero-based arrays._

Notice that `floor(n/2)` is the index of one or more the the middle items of the array:
  - `floor(n/2)` if n is odd
  - `floor(n/2)` and `floor(n/2) - 1` if n is even

To insert:
1. append number to list
2. "Bubble Up":
  - `parent = heap[floor((i-1)/2)]`
  - if parent < number then swap them

- Test your inserts
- Test your removals

Note:
  - cannot be efficiently searched because it is not a BST
  - we don't know any facts that will improve a linear search

Heap Interface:
- get the dominating element, e.g. `min` or `max`
- `add` or `insert` an element
- `remove` or `delete` an element

We can store any binary tree in an array without pointers but:
- array still requires empty spots for missing nodes
- methods to save memory make it less flexible

## RESOURCES

https://en.wikipedia.org/wiki/Binary_heap

[8.4. heapq — Heap queue algorithm](https://docs.python.org/2/library/heapq.html).