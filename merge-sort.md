merge-sort -- Merge Sort
========================

## tl;dr

```
   [2,4,3,1]     \
  [2,4] [3,1]    | mergesort
[2] [4] [3] [1]  + 
  [2,4] [1,3]    | merge
   [1,2,3,4]     /
```

Merge Sort:
- divide and conquer via recursion, similar to exchange sort
- O(n logn) -- T(n) = 2T(n/2) + O(n)
- mergesort(a)
  - log(n) recursions
  - BASE CASE: list is length 1 or 0
  - RECURSIVE CASE: list is 2 or more
    - find middle of list: `floor(length / 2)`
    - DIVIDE: mergesort left and right halves (you'll hit base case and return)
    - return of merge left and right halves
- merge(a,b):
  - (You must determine where sorted list goes (new array, reference, etc))
  - iterate through both halves simultaneous
  - select minimum of two and increment that pointer only
  - after reaching end of one half iterate through remainder of other

## CONCEPT

- What if we took a list of items and split it in half?
- Then we took each of those halves and split them again?
- Eventually we'll end up with lists composed of only one element.
- We can then merge each pair of lists into a sorted one.
- Then continue to merge each pair of newly merged lists.
- Until we're left with a single sorted list.
- This is what **merge sort** does.

Merge sort is a _divide-and-conquer_ algorithm:
1. Divide: Divide the list in half (unless it's a list of one, then just return it)
2. Conquer: Recursively sort the two halves
3. Combine: Merge the sorted halves into one

Try this with [5,8,7,2,1,4,6,3].
Remember to recursively half each list.
(So you should end up with a [5] first).
Think about the steps you are taking when you merge the lists.
It's easy for us to do something mentally and forget we actually ran an algorithm!
We need to be able to program a computer to do this.

```
       [5,8,7,2,1,4,6,3]

      [5,8,7,2] [1,4,6,3]

    [5,8] [7,2] [1,4] [6,3]

[5] [8] [7] [2] [1] [4] [6] [3]

    [5,8] [2,7] [1,4] [3,6]

      [2,5,7,8] [1,3,4,6]

       [1,2,3,4,5,6,7,8]
```

```python
# pure mergesort in Python (HGPA)
def mergesort(c):
  """Returns a new sorted list"""
  length = len(c)
  if length <= 1: # base case
    return c
  else: # recursive case
    # divide
    middle = length // 2
    a = c[0:middle]
    b = c[middle:length]
    return merge(mergesort(a), mergesort(b)) # conquer and combine
```

Can you already program the merge?
This is where the majority of the work is.

### Merging Two Lists

- We want to merge two sorted lists into a new one
- The resulting list should be sorted too
- We can do this by simultaneously iterating across both lists
- We can compare a pair of items, one from each list
- And insert the smallest item into the new one
- We then move on to the next item of that list we took from and repeat
- Eventually we'll finish iterating through one of these lists
- So we'll insert the remaining items from the other list into our new one

```python
# pure merge in Python returning a new list (HGPA)
def merge(a,b):
  """Merge lists a, b and return a new list c"""
  c = []
  i = 0
  j = 0
  while i < len(a) and j < len(b):
    if a[i] < b[i]:
      c.append(a[i])
      i += 1
    else:
      c.append(b[j])
      j += 1
  while i < len(a):
    c.append(a[i])
  while j < len(b):
    c.append(b[j])
  return c
```

## IMPLEMENTATIONS

Lets look at some other implementations.
Recall that (sorted) `merge` assumes that each list is sorted.

Sorting implementations usually mutate existing arrays to improve space performance.
In this case you will need to pass in the destination array as a parameter.
These mutating functions will not return anything.

When implementing any sorting you should also ask whether the list is indexable.
For example, linked-lists are not.

### `mergesort`

```python
# mutating mergesort in Python (Goodrich)
def mergesort(c):
  """Sort a list, `c`"""
  length = len(c)
  if length <= 1:
    return
  middle = length // 2
  a = c[0:middle]
  b = c[middle:length]
  mergesort(a)
  mergesort(b)
  merge(a,b,c)
```

### `merge`

```c
/**
 * destination-mutating merge in C (Pohl)
 * This is very traditional.
 */
void merge(int a[], int b[], int c[], int a_length, int b_length) {
  int i = 0, j = 0, k = 0;
  while (i < a_length && j < b_length) {
    if (a[i] < b[j]) {
      c[k++] = a[i++];
    } else {
      c[k++] = b[j++];
    }
  }
  while (i < a_length) {
    c[k++] = a[i++];
  }
  while (j < b_length) {
    c[k++] = b[j++];
  }
}
```

```python
# destination-mutating merge in Python (Goodrich)
def merge(a, b, c):
  """Merge two sorted Python lists a and b into properly sized list S."""
  i=j=0
  while i + j < len(c):
    if j == len(b) or (i < len(a) and a[i] < b[j]):
      c[i+j] = a[i]
    else:
      c[i+j] = b[j]
```

```c
/**
 * mutating in-place "merge" using queues (Skiena)
 * This is "safe" method that can also be applied to non-indexable lists.
 */
merge(item_type s[], int low, int middle, int high) {
  int i;
  queue buffer1, buffer2;
  init_queue(&buffer1);
  init_queue(&buffer2);
  for (i=low; i<=middle; i++) {
    enqueue(&buffer1,s[i]);
  }
  for (i=middle+1; i<=high; i++) {
    enqueue(&buffer2,s[i]);
  }
  i = low;
  while (!(empty_queue(&buffer1) || empty_queue(&buffer2))) {
    if (headq(&buffer1) <= headq(&buffer2)) {
      s[i++] = dequeue(&buffer1);
    } else {
      s[i++] = dequeue(&buffer2);
    }
  }
  while (!empty_queue(&buffer1)) {
    s[i++] = dequeue(&buffer1);
  }
  while (!empty_queue(&buffer2)) {
    s[i++] = dequeue(&buffer2);
  }
}
```

## TODO

- time complexity
- non-recursive merge sort (Pohl), (Goodrich)

## REFERENCES

12.2 Merge-Sort. Data Structures and Algorithms in Python. Michael T. Goodrich, Roberto Tamassia & Michael Goldwasser. 2013.
Good visualizations but a little wordy.

4.5 Mergesort: Sorting by Divide-and-Conquer. The Algorithm Design Manual. Steven S. Skiena. 2008. The best short treatise.

4.6 Merge Sort. Data Structures and Algorithms with Python. Kent D. Lee & Steve Hubbard. 2015.

5.2.4. Sorting by Merging. The Art of Computer Programming, Volume 3: Sorting and Searching. Second Edition.

6.9 An Example: Merge and Merge Sort. A Book on C. Al Kelly & Ira Pohl. Fourth Edition. 1998.
