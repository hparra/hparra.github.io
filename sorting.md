sorting
=======

Sorting is a fundamental problem in computer science.
Each method of sorting has trade-offs in time and space complexity.
Sorting is often the first step in solving more complex problems.

## Sorting Functions

Most of the time we don't care about any of this theoretical stuff.
We just want the damn thing sorted!
Most languages include some sort of sorting,
either through their standard libraries,
or as method on a fundamental data type.

**C** has the `qsort` sorting function in its standard library.
Despite its name it is not necessarily quick sort.
The implementation is vendor specific.
`qsort` requires 4 parameters:
- array
- array length
- size of an individual array element
- comparator function

```c
#include <stdlib.h>

/* basic integer comparator function */
int cmp(const void* elem1, const void* elem2) {
    int *x = (int *) elem1;
    int *y = (int *) elem2;
    return *x - *y;
}

qsort(array, array_length, sizeof(int), cmp);
```
**

## Sorting Implementations

Methods of sorting include:
- bubble sort
- selection sort
- insertion sort
- [merge sort](merge-sort.md)
- quick sort
- bucket sort
- radix sort
- heap sort
- Tim sort
- shell sort

Implementations of sorting algorithms can be tricky. Why?
- is list a continuous array or a linked-list?
- is function pure or does it mutate?
- do we manage memory ourselves?