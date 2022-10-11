Heaps
=====

## Introduction

Usually the word "heap" refers to an unordered pile of things, a mess.
Many times it refers to a _trash heap_, a pile or fill of garbage.

In computer science, a heap usually refers to two different things:
- the _memory heap_, or area for dynamically allocated memory
- the _heap_ data structure, specifically _binary heaps_

The use of the word "heap" for these may not be a coincidence!
Memory heaps hold dynamically (randomly) allocated memory with no particular order.
Heap data structures do not maintain sequential order of their content.
In other words, in some way, they are each a respective mess.

Furthering the analogy, some programming languages manage the automatic deallocation of memory within a heap with something called a _garbage collector_.

The heap data structure is useful for quickly getting the maximum or minimum element from a set. They are generally associated with priority queues, where the element with the highest priority is always at the front of the line.

You may have also heard of _heapsort_.
This is a sorting algorithm that uses a heap to do all the work.

## Heap Data Structures

_heap-labeled tree_:
- is a binary tree such that each node _dominates_ key labeling of its children
- is a _min-heap_ when node dominates with smaller keys
- is a _max-heap_ when node dominates with larger keys

Usually when we talk about heaps (data structure) we're referring to _binary heaps_.
There are other types as well.

> I'm bigger and bolder and rougher and tougher.
> In other words sucker there is no other.
> I'm the one and only Dominator. 

Say you have a list of numbers: `2 7 8 1 5 9 3`.

We want to construct a data structure such that we can always get the dominating element in constant time.
In the case of a min-heap it would be `1`.
This implies that all the real work needs to happen on element insertion and deletion.
