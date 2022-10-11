---
---

Permutations
============

## Introduction

A **permutation** is a particular ordering of a set of objects.
Formally, a permutation of a finite set S is an ordered sequence all the elements with each one appearing once.

How many different ways can you order {1,2,3}?

Answer -- 6:
- {1,2,3}
- {1,3,2}
- {2,1,3}
- {2,3,1}
- {3,1,2}
- {3,2,1}

For a set of length n, there are n! permutations.
So 3! = 6.

Why is this?
- For each number n,
- you can keep n in place,
- and permute the remaining n - 1
- then keep (n-1) in place and permute (n-2)
- keep going...
- n = 2 has 2 permutations
- n = 1 has 1 permutations

`(n)(n-1)(n-2)...(2)(1)` is the definition of factorial of n.

We can also talk about the **k-permutation** of a set: `P(n,r) , or nPr = n!/(n-r)!`

## Generating Permutations

Since O(permutations) is n! we have something that won't run in our lifetime for any decent size n...

Many algorithms seek the best way to order a set of objects:
- traveling salesman problem -- the least-cord order to visit n cities
- bandwidth minimization -- order vertices of a graph on a line so as to minimize the length of the longest edge
- determining graph isomorphism -- oder the vertices of one graph so that it is identical to another

The _lexiographic order_ of the permutations of {1,2,3} are:

- {1,2,3}
- {1,3,2}
- {2,1,3}
- {2,3,1}
- {3,1,2}
- {3,2,1}

# REFERENCES

Permutations. Sal Khan. Kahn Academy.

The Algorithm Design Manual. Second Addition. Steven Skiena. 2008.