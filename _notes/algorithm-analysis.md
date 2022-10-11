---
---

algorithm analysis
==================

> If you want to be a good programmer you just program every day for two years -- you'll be an excellent programmer. If you want to be a world-class programmer you can program every day for ten years, or you could program every day for two years and take an algorithms class.
> - As told by [Charles E. Leiserson](http://people.csail.mit.edu/cel/)

## Why do we need to analyze algorithms?

- we're concerned with performance or _efficiency_ of algorithms
- algorithms require time and space to execute
- time and space become limiting factors when input size starts to get significantly large
  - there may not be enough time
  - there may not be enough space  
  - time and space cost money in some (in)direct way
- we need a way to characterize an algorithm's efficiency
- we can compare efficiencies between known algorithms
- we can then try to minimize time and/or space to solve a problem

## Asymptotic Analysis

We analyze an algorithm's **complexity**:
- **running time**, or **time complexity**
- space (memory)

Usually we want an estimate of an algorithm's complexity in the _asymptotic_ sense (large inputs).
- Recall [asymptotes](https://en.wikipedia.org/wiki/Asymptote) from algebra

There are three scenarios:

**worst-case complexity**:
- maximum time (upper bound) on any input of size n
- usually what we're concerned about
- described by big O (or big theta) of a function
- usually also the expected-case

**average-case (expected) complexity**:
- expected time over all inputs of size n
    
**best-case complexity**:
- not a useful metric

## Asymptotic Notation

Asymptotic notations:
- allow us to describe how a complexity scales
- allow us to describe the rate of increase
- is a relative notation focused on the most dominant, dependent terms

**big theta**
- f(n) = Θ(f(n))
- the **asymptotically tight bound** on the running time
  - _asymptotically_ => it matters for only large values of n
  - _tight bound_ => running time to within a constant factor above and below
    - so for constants j, k
      - at most k * f(n)
      - at least j * f(n)
  - if O(f(n)) and Ω(f(n)) then Θ(f(n))
    - Think Θ🍔 (Theta Burger)
      - O(f(n)) is top bun
      - Ω(f(n)) is bottom bun
      - if buns are not the same, then it's not a burger

**big O**
- f(n) = O(g(n))
- the **asymptotic upper bounds** on the running time
- "f(n) takes _at most_ a certain amount of time, O(g(n))"
- "f(n) is in the set of functions like g(n)), O(g(n))"
- `f(n) = O(g(n)) <=> 0 <= f(n) <= c * g(n), c > 0, n >= n0 > 0`
- Θ(f(n)) implies O(f(n))

**big omega**
- f(n) = Ω(g(n))
- the **asymptotic lower bounds** on the running time
- "f(n) takes _at least_ a certain amount of time, g(n)"
- "f(n) is in the set of functions like g(n)), O(g(n))"
- Θ(f(n)) implies Ω(f(n))

## Orders of Growth

<img src="https://upload.wikimedia.org/wikipedia/commons/7/7e/Comparison_computational_complexity.svg" alt="Orders of Growth" style="width:100%;max-width:360px"/>

| Running Time | Name         | Fundamental Concept |
| ------------ | ------------ | ------- |
| O(1)         | constant     | instant operation independent of input _n_ |
| O(log n)     | logarithmic  | "binary" (search, tree) |
| O(n)         | linear       | must read all _n_ inputs |
| O(n log n)   | linearithmic | _O(log n)_ operations _n_ times, divide-and-conquer|
| O(n^2)       | quadratic    | loop within a loop |
| O(c^n)       | exponential  | _n_ values that each can be any _c_ values, multi-call recursion, branches^depth |
| O(n!)        | factorial    | permutations of a list |

## Application

(TODO: WIP. HGPA)

We usually only care about Big O of worst-case. So what is the upper bound of the worst-case scenario of an algorithm?

**WARNING: In industry the use of Big O is nearly synonymous with Big Θ (if applicable) in worst-case**

- drop constants
- drop non-dominant terms
- determine how different inputs relate to each other
  - sequential loops => addition
  - nested loops => multiplication

## References

[Asymptotic Notation](https://www.khanacademy.org/computing/computer-science/algorithms/asymptotic-notation/a/asymptotic-notation). Tom Cormen and Devin Balkcom. Algorithms. Kahn Academy.

Big O. Cracking the Code Interview. Sixth Edition. Gayle Laakmann McDowell.

[Big-O Cheat Sheet](http://bigocheatsheet.com/). Eric Rowell. Tables with complexities of common data structures and algorithms.

[Big O Notation](https://brilliant.org/wiki/big-o-notation/). Brilliant.

[Big O Notation](https://www.interviewcake.com/article/javascript/big-o-notation-time-and-space-complexity). Interview Cake.

[Big O notation](https://en.wikipedia.org/wiki/Big_O_notation). Wikipedia.

Getting Started. Introduction to Algorithms. Second Edition. CLRS. 2002.

Growth of Functions. Introduction to Algorithms. Second Edition. CLRS. 2002.

[Lecture 1: Administrivia; Introduction; Analysis of Algorithms, Insertion Sort, Mergesort](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-046j-introduction-to-algorithms-sma-5503-fall-2005/video-lectures/lecture-1-administrivia-introduction-analysis-of-algorithms-insertion-sort-mergesort/). Charles Leiserson. Introduction to Algorithms. MIT OpenCourseWare. 2005 Fall.

[Lecture 1: Algorithmic Thinking, Peak Finding](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/lecture-videos/lecture-1-algorithmic-thinking-peak-finding/). Srini Devadas. Introduction to Algorithms. MIT OpenCourseWare. 2011 Fall.

[Recitation 1: Asymptotic Complexity, Peak Finding](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/recitation-videos/recitation-1-asymptotic-complexity-peak-finding/). Victor Costan. Introduction to Algorithms. MIT OpenCourseWare. 2011 Fall.
