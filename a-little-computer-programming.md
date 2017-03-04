A Little Computer Programming
=============================

(This is WIP)

This is a little workbook on computer programming.
It falls somewhere between beginner and intermediate.
I am not an expert.

This workbook is meant to be followed sequentially.
I recommend Python or Javascript.
If you use a Lisp you will identify things you can skip.

Computer programming is a skill.
Deliberate practice is necessary to improve this skill.
Only with skill and effort can we achieve.

Computer programming is fun.
Enjoy these challenges for what they are.
Forget about everything else.

## Checklist

Dimension 1:
- Essential Functions
  - [x] isempty (null)
  - [x] first (car)
  - [x] rest (cdr)
  - [x] push (cons)
- More Basic Stuff
  - [x] repeat
  - [x] last
  - [x] initial
  - [x] range
  - [x] remove
  - [x] slice
- Reversing
  - [x] reverse (pure)
  - [x] reverse (pure and recursive)
  - [x] reverse (in-place mutation)
- Using Linear Search
  - [x] index_of
  - [x] includes
  - [x] find
  - [x] some
  - [x] every
  - [x] equals
  - [ ] issorted
- Intro to Map/Reduce
  - [x] filter
  - [x] map
  - [x] reduce
  - [x] reduceRight
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
  - [x] partition_at_pivot
  - [x] concat
- Working with a Sorted List
  - [ ] unique No. 2
  - [ ] sorted_merge
  - [ ] binary_search
- Sorting
  - [x] reverse
  - [x] bubble_sort
  - [x] selection_sort
  - [x] insertion_sort
  - [ ] quick_sort
  - [ ] merge_sort
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
- (IDK where to put these yet)
  - [ ] bucket_sort
  - [ ] radix_sort
  
Dimension 1.5:
- Lists
  - [ ] linked_list
  - [ ] skip_list
- Trees
  - height
  - is_complete
  - is_perfect
  - is_bst
  - dfs_inorder
  - dfs_postorder
  - dfs_preorder
  - bfs
- [ ] binary_heap
- [ ] binary_search_tree
- [ ] trie

Dimension 2:
- Graph
- Depth First Search
- Breadth First Search
- Backtracking

## Dimension 1

### Essential Functions

Write these basic functions but don't forget them!

`isarray(obj)`
  - return whether the obj is an array
  - e.g. `isarray(1) = False`

`isempty(array)`
  - return whether the array is empty
  - e.g. `isempty([]) = True`
  - aka _null?_

`first(array)`
  - return first element of the array
  - e.g. `first([1,2,3]) = 1`
  - What should we return if the array is empty?
  - aka _car_

`rest(array)`
  - return the array with the first element removed
  - e.g. `rest([1,2,3]) = [2,3]` 
  - What if there is only one item?
  - What if there are no items?
  - Built into Python as `array[1:]`
  - aka _cdr_

`push(val, array)`
  - return copy of array with val prepended to it
  - e.g. `push([1,2], 0) = [0,1,2]`
  - aka _cons_
  - in fact, name this function `cons` instead

### Useful Functions

`repeat(val, count)`
  - return an array with count elements of val 
  - e.g. `repeat(1, 5) = [1,1,1,1,1]`

`last(array)`
  - return last element of the array
  - e.g. `last([1,2,3]) = 3`
  - What should we return if the array is empty?

`initial(array)`
  - return the array with the last element removed
  - e.g. `initial([1,2,3]) = [1,2]`
  - What if there is only one item?
  - What if there are no items?
  - Built into Python as `array[:-1]`

`range(start, stop, step=1)`
  - return an array consisting of sequential numbers starting at start and ending at (but not including) stop, incrementing by step each time
  - e.g. `range(1,5,1) = [1,2,3,4]`

`remove(array, val)`
  - return a new array with the first instance of val removed, if it exists
  - e.g. `remove([1,2,3,2,1], 2) = [1,3,2,1]`

`slice(array, start, stop, step=1)`
  - return a new subarray of array consisting of elements starting at start and ending at (but not including) stop, incrementing by step each time
  - e.g. `slice([1,2,3,4],1,3) = [2,3]`

Now implement all of these recursively without using loops.
Use the _Essential Functions_.
You do not need to know the length of the array!

### Reversing

`reverse(array)`
  - reverse the array!

Please do this three ways:
  - iteratively, returning a new array
  - recursively, returning a new array
  - iteratively, mutating the original array
  
**From here on out please implement each function twice: iteratively and recursively!**

### Linear Search

- indexOf
- includes
- find
- some
- every
- equal

### Map/Reduce

- filter
- map:          `array = map(function, array)`
- reduce:       `value = reduce(function, array, initial_value)`
- reduceRight:  `value = reduce_right(function, array, initial_value)`

Use `reduce` to calculate some common reductions:
  - sum
  - min/max
  - includes
  - some
  - every
  - factorial
  - gcd

Noticing a recurring theme here? 😂

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

(I'm just throwing these last two down here for now)


`concat(*array)`
  - concatenate all the arrays and return a single array

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

### Sorted Lists

- unique
- binary_search
- sorted_merge
- sorted_unique

### Sorting

All these should mutate the array, i.e. do the work in-place.

`reverse_mutate(array)`
  - reverse items of an array in-place
  - e.g. `reverse_mutate([1,2,3,4,5])`
  - ```python
    # python example test
    array = [1,2,3,4,5]
    expected = [5,4,3,2,1]
    reverse_mutate(array)
    self.assertEqual(expected, array)
    ```

Use the following arrays to test:
  ```
  actual = [2,1,9,8,4,5,7,0,3,6]
  expected = [0,1,2,3,4,5,6,7,8,9]
  ```

`bubble_sort(array)`
  - sort the array by comparing each pair of items and swapping them if necessary
  - At most how many times do you need to do this? (worst-case scenario)
  - How do you know when array is sorted?
  
`selection_sort(array)`
  - sort the array by:
    - selecting the left-most value
    - searching the remaining items for the smallest value
    - swapping the left-most value and smallest value
    - repeating by selecting the second left-most value, etc.

`insertion_sort(array)`
  - "card sort"
  - sort the array by:
    - selecting the first two adjacent values
    - swapping them if necessary
    
TODO:
- quick_sort
- merge_sort

## Strings

- split
- word_count

`trim(string, characters=" \t\n\r")`
  - create a new string from string such that the specified characters are removed from the beginning and ends of string
  - e.g. `trim("  hello ") = "hello"`

`parse_int`
  - create a number from a string representing that number
  - e.g. `parse_int("1234") = 1234`
  - What about negative numbers?

- repeat
- index_of
- includes
- replace (non-regex)
- dc (reverse-polish)

## Bits

- parse_int (bitstring)

## Trees



TODO

## Ideas

- alien: produce a function that creates a paragraph of ~m sentences with approximately n sentences varying v words.
- choose your own adventure

## Note regarding Space Matrices

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

After various failed online, phone, and on-site technical interviews I made this workbook. It is the culmination of lessons learned the hard way.

This is a work in progress. HGPA