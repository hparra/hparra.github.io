Lists
=====

A **list** a sequence of things. It is the fundamental abstract data type (ADT).

Most programming languages have at least one native list data type, e.g.:
- JavaScript's `Array`
- Python's `List`
- Java's `ArrayList` and `Vector`
- C++'s `vector` 
- Go's `Slice`

The list above (sorry) may highlight that the difference between **arrays** and lists may not always be clear.
For some high-level languages their "arrays" are lists.
We'll define arrays as continuous memory whose size needs to be defined at time of compilation.
We'll define lists as iterable sequence that support random access. As such, we will not consider linked-lists or doubly linked-lists to be ADTs.

These high-level lists usually have sufficient methods to mimic more restrictive ADTs, like stacks and queues.

## Implementation

- dynamic arrays
- single or doubly linked lists

### Java

Java uses interfaces to define a `List` as an iterable collection.
- [ArrayList](https://docs.oracle.com/javase/7/docs/api/java/util/ArrayList.html)
- [Vector](https://docs.oracle.com/javase/7/docs/api/java/util/Vector.html) is an extended ArrayList with added thread safety.
