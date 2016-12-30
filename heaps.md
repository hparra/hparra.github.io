Heaps
=====

## Introduction

Usually the word "heap" refers to an unordered pile of things, a mess.
Many times it refers to a _trash heap_, a pile or fill of garbage.

In computer science, a heap usually refers to two different things:
- the _memory heap_
- the _heap_ data structure

The use of the word "heap" for these is not a coincidence!
Memory heaps hold dynamically (randomly) allocated memory with no particular order.
Heap data structures do not maintain sequential order of their content.
In other words, in some way, they are each a respective mess.

Furthering the analogy, some programming languages manage the automatic deallocation of memory within a heap with something called a _garbage collector_.

You may have also heard of _heapsort_.
This is a sorting algorithm that uses a heap to do all the work.

## Heap Data Structures

Heaps
- binary heaps
- fibonacci heaps

_heap-labeled tree_:
- is a binary tree such that each node _dominates_ key labeling of its children
- is a _min-heap_ when node dominates with smaller keys
- is a _max-heap_ when node dominates with larger keys

### Binary Heaps

Usually when we talk about heaps we're referring to _binary heaps_.

- can represent binary trees without using pointers (by using an array)
  - left child => `2k`
  - right child => `2k + 1`
  - parent => `floor(n/2)`
- cannot be efficiently searched because it is not a BST
  - we don't know any facts that will improve a linear search

We can store any binary tree in an array without pointers but:
- array still requires empty spots for missing nodes
- methods to save memory make it less flexible

### Constructing Heaps

```
place new in left most spot (n+1)
check dominance
if new element dominates parent
then switch parent with new element
(note other element should be happy)
but now old parent may not dominate new children
so recurse
```

Swaps happen between levels, and a tree will have lg(n) levels.
Since there are n items to be inserted, insertion will take `O(nlogn)`.

### Extracting the Minimum

Top of the heap sits in first position of array.
If popping heap then replace top element whole with right-most (bottom) leaf.

(Book in unclear here)