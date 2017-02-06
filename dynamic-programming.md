Dynamic Programming (DP)
========================

## tl;dr

Dynamic Programming:
- is a way to design algorithms that search all possibilities
- stores results to avoid recomputing
- usually starts with correct recursive algorithm first
- trades space for time

Technique (Skiena):
1. formulate answer as recurrence relation or recursive function
2. show number of different parameter values taken on by recurrence is bounded by a hopefully small polynomial
3. specify an order of evaluation for the recurrence so that the partial results you need are always available

## Examples

Three traditional examples:
- fibonacci numbers
- binomial coefficients
- coin change problem

## Common Applications

**Longest Common Substring**: Given a set of strings, find the longest substring common to all strings. This could also be solved with a suffix tree.

**Longest Common Subsequence (LCS)**: Given a set of sequences, find the longest subsequence common to all sequences. A subsequence of a string is a set of characters that appear in left-to-right order, but may not be consecutive. This is precisely what `diff` does.

**Knapsack problem**: Given a set of items, each with a weight and a value, determine the number of each item to include in a collection so that the total weight is less than or equal to a given limit and the total value is as large as possible.

**Subset Sum problem**: Given set of integers is there is non-zero subset whose sum is zero? Special-case of knapsack.

**Partition problem**: Given a multiset of positive integers, can it be partitioned into two subsets such that the sum of the numbers in each subset are equal. Special-case of subset sum.

**Cocke–Younger–Kasami (CYK)**: Parses context-free grammars.

TODO:
- Longest Increasing Sequence
- Levenshtein (edit) distance
- Floyd's all-pairs shortest path algorithm
- Bellman–Ford -- finding the shortest distance in a graph

## REFERENCES

[Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming). Wikipedia.

[Longest Common Subsequences](https://www.ics.uci.edu/~eppstein/161/960229.html). David Eppstein. 1996-02-29.