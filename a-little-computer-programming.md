A Little Computer Programming
=============================

_This is a serious WIP!_

Dimension 1:
- Basics
  - [ ] repeat 
  - [x] range
  - [x] first
  - [x] last
  - [x] initial
  - [x] rest
  - [x] slice
  - [x] concat
- Using Linear Search
  - [x] index_of
  - [x] includes
  - [x] find
  - [x] some
  - [x] every
  - [ ] equal
  - [x] filter
- Intro to Map/Reduce
  - [x] map
  - [x] reduce
  - [x] reduceRight
  - common reductions
    - sum
    - min/max
    - includes
    - some
    - every
    - factorial
    - gcd/lcm
- Partitioning
  - [x] flatten
  - [x] partition
  - [x] interleave
  - [x] deinterleave
  - [x] zip
  - [x] unzip
  - [x] groups_of
  - [x] groups_prioritized
  - [x] groups
  - [x] paritition_at_pivot
- A Little Set Theory (Part 1?)
  - [x] unique (set)
  - [x] subset
  - [x] union
  - [x] intersection
  - [x] xor (symmetric difference)
  - [x] difference
  - [ ] combinations
  - [ ] permutations
  - [ ] powerset
- Working with a Sorted List
  - [ ] unique No. 2
  - [ ] sorted_merge
  - [ ] binary_search
- Sorting
  - [ ] reverse
  - [ ] bubble_sort
  - [ ] selection_sort
  - [ ] insertion_sort
  - [ ] quick_sort
  - [ ] merge_sort
  - [ ] bucket_sort
- Some Statistics
  - [ ] sample
  - [ ] mean
  - [ ] medium
  - [ ] mode
  - [ ] shuffle
- Data Structures
  - [ ] list
  - [ ] queue
  - [ ] stack
  - [ ] bloom_filter
  - [ ] unordered_set
- Extra
  - [ ] unique_chunk
  - [ ] sorted_unique
  - [ ] radix_sort
  
Dimension 1.5:
- Lists
  - [ ] linked_list
  - [ ] skip_list
- Trees
  - [ ] The Traversals  
  - [ ] binary_heap
  - [ ] binary_search_tree
  - [ ] trie

Dimension 2:
- Graph
- Depth First Search
- Breadth First Search
- Backtracking



## Dimension 1

*Helpers*:
- random

*Arrays*:
- range:        `array = range(start, stop, step)`
- rangeRandom:  `array = range_random(start, stop, count)`

### The Basics

- first
- last
- initial
- rest
- slice:        `array = slice(array, start, stop, step)`
- concat:       `array = concat(array1, array2, ...)`

### Linear Search

- indexOf
- includes
- find
- some
- every
- equal
- filter

### Map/Reduce

- map:          `array = map(function, array)`
- reduce:       `value = reduce(function, array, initial_value)`
- reduceRight:  `value = reduce_right(function, array, initial_value)`
also:
  - sum
  - min/max
  - includes
  - some
  - every
  - factorial
  - gcd/lcm

### Partitioning

This is a challenging group. No joke intended.

`flatten(array)`
  - flatten an array of arrays into a single array
  - e.g. `flatten([[0,1],[2],[3,4,5]]) = [0,1,2,3,4,5]`
  - e.g. "flatten this paged data into a single array"
  - What happens if a subarray has arrays too?
    - e.g. `flatten([[0,[1]],[2],[3,[4],5]]) = [0,1,2,3,4,5]`

`partition(fn, array)`
  - partition an array into array of two arrays using predicate function
  - e.g. `partition(lambda x: x % 2 == 0, [1,2,3,4,5]) = [[2,4],[1,3,5]]`

`partition_at_pivot(fn, array)`
  - partition an array in-place such that when some value is selected from the array:
    - all other values less than it are to the left
    - all other values greater than it are to the right
    - the value itself is in it's proper location should the array be sorted
  - e.g. `partition(lambda a,b: a-b, [2,1,9,8,4,5,7,0,3,6]) = [2,1,3,0,4,5,6,7,8,9]`
  - ```python
    # python example test
    comp = lambda a,b: a-b
    array = [2,1,9,8,4,5,7,0,3,6]
    expected = [2,1,3,0,4,5,6,7,8,9]
    self.assertEqual(6, paritition_at_pivot(comp, array))
    self.assertEqual(expected, array)
    ```

`groups_of(array, n)`
  - partition an array into an array of arrays with n items each
  - e.g. `groups_of([1,2,3,4,5], 3) = [[1,2,3],[4,5]]`
  - e.g. "paginate with each page having n items"
  - also known as _paginate_ or _chunk_
  - What happens when not every array can have n items?
  - You may have programmed this by iterating over either the source or the destination. Now try it the other way. Which looks better?

`groups_prioritized(array, n)`
  - partition a prioritized array into an array of n prioritized arrays
  - e.g. `groups([1,2,3,4,5], 3) = [[1,4],[2,5],[3]]`
  - e.g. "split m prioritized jobs amongst n workers"

`groups(array, n)`
  - partition an array into an array of n arrays
  - e.g. `groups([1,2,3,4,5], 3) = [[1,2],[3,4],[5]]`
  - e.g. "split items into n columns (left to right)"
  - if n > length of array then include empty arrays
  - aka _groups_sequential_
  - I found this one difficult!

`interleave(*arrays)`
  - interleaves two or more arrays and into single array
  - e.g. `interleave([0,2,4],[1,3,5]) = [0,1,2,3,4,5]`

`deinterleave(array, n)`
  - deinterleave an array into an array of n arrays
  - e.g. `deinterleave([0,1,2,3,4,5], 2) = [[0,2,4],[1,3,5]]`

`zip(*arrays)`
  - interleave m arrays with n items and return array of n arrays with m items
  - e.g. `zip([0,1,2],[3,4,5]]) = [[0,3],[1,4],[2,5]]`
  - think of "zipper"
  - technically also _transpose_
  - similar to _interleave_
  - What happens when subarrays are not equal size?

`unzip(array[])`
  - deinterleave array of n array with m items into array of m arrays with n items
  - e.g. `unzip([[0,3],[1,4],[2,5]]) = [[0,1,2],[3,4,5]]`
  - What happens if there is only one subarray?
  - What if happens when subarray are not equal size?

### Set Theory

A _set_ is an unordered collection of unique items.
In other words, there are no duplicates.
Hint: Dictionaries can help you remember things!

- unique (set)
- subset
- union
- intersection
- xor
- difference

---
- mean
- median
- mode
---
- chunk_unique (airbnb)
- sort
- merge_sorted (n log n)

### Sorted Lists

- unique
- binary_search
- sorted_merge
- sorted_unique

### Sorting

- The Bubble Sort
- The Selection Sort
- The Insertion Sort

- The Quick Sort
- The Merge Sort
- Bucket Sort
- The Radix Sort
-

## Strings

- replace
- split
- trim
- words
- count
- alien: produce a function that creates a paragraph of ~m sentences with approximately n sentences varying v words.

---

for any `[M][N]` array (Rows x Columns)

```python
# assuming array[m][n]
for m in xrange(array): # rows
  for n in xrange(array[m]): # columns
    print array[m][n]  
```

```

1 x N = [1,2,3]
M x 1 = [
  [1],
  [2],
  [3],
]


The problem with this is that it 

How we usually think of arrays in code:
[
  [1,2,3],
  [4,5,6]
  [7,8,9]
]

Translates to these [M][N] pairs:
[
  [(0,0),(0,1),(0,2)],
  [(1,0),(1,1),(1,2)],
  [(2,0),(2,1),(2,2)],
]

If the center is the origin then move like this.
[
  [(-1,-1),(-1,0),(-1,+1)],
  [( 0,-1),( 0,0),( 0,+1)],
  [(+1,-1),(+1,0),(+1,+1)],
]

```


Just picture the matrix!
FORGET the Cartesian plane! this i Column x Row with origin at bottom-left
FORGET the image plane! this is column Column x Row with origin at top-left
FORGET latitude and longitude. I can't even remember which is which or how it works!

https://en.wikipedia.org/wiki/Row-_and_column-major_order

## Tips

- it may be easier to iterate over the destination versus the source
- if the dimension of the destination will change then you may need to calculate this beforehand
- Python: `[[]] * num` does weird stuff! Don't use it.

## History

This is a little tutorial on computer programming. It falls somewhere between beginner and intermediate, but don't be fooled like I have into thinking you do not fall into this category.

After various failed online, phone, and on-site technical interviews I made this tutorial. It is the culmination of learning by doing various things the wrong way. I can't guarantee it will be useful to you, but I've already learned things from creating and following it, so I'm now a little better than I was before. If you're a new programmer or an experienced programmer whose grown stale then this may be useful to you.

This tutorial is meant to be followed sequentially. I group exercises by theme -- which may be a certain application domain, programming technique, or data structure. I've notice there is usually some correlation between one or two of these. It works best to use a modern interpreted language like Javascript or Python but you should be able to do these in any language. In fact, I hypothesize that repeating these exercises with a new language would kill two birds with one stone.

When we focus on the process of computer programming it is enjoyable time. There is an end-product of solving each of these problems that we should focus on, but it is usually not necessarily the answer, but how we got there, and how we can apply that process to new problems. Please focus on this and not what or who you're supposed to be once you've completed the tutorial.

This is a work in progress. HGPA