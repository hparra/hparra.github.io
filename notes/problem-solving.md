---
---

Problem Solving
===============

Problem solving is a problem!

## tl;dr

1. understand solution medium
  - whiteboard, editor, meeting?
- confirm the problem
  - what is input?
  - what is desired output?
- analyze the problem
  - identify variables
  - identify computed facts
    - lengths
    - locations
  - ask question
- work through medium-sized example
  - one may be given to you
  - draw 
- work through tiny examples (edge cases)
  - "0 or 1" examples
  - "perfect" examples
  - "no work" examples, e.g. list already sorted
- reanalyze problem given new context
- state solution idea(s)
  - _always_ state a brute force solution
- iterate "high-level" solution
  - define facts
  - define the primary function that solves problem
  - define additional function that do actual work
- analyze hotspots
  - "off by 1" errors
- tests

## HGPA's Interview Question Heuristic

**Survey the Workspace**

Is it a whiteboard, a piece of paper, or is it a discussion with a group? Perhaps it may be on a computer in a text editor or IDE. Understand how to use each of these mediums effectively. This is the essence of the _programming environment_.

**Understand the Domain**

This is the essence of _programming arts_.

**Understand the Problem**

If a problem is spoken to your then write it down.
Repeat or present your transcription and confirm that it is correct.
This is the essence of _requirements_ and _specification_.

Remember that every problem is essentially function: it has one or more inputs, and an output. Understand the types and domains of your inputs and outputs -- this include upper limits. Sometimes the inputs themselves are complex.

Remember that some problems are still problems because of how they are stated. Information may be implied or hidden. Sometimes necessary information is not supplied at all.

**Develop Test Cases**

Stating various simple and extreme combinations of inputs and their known outputs will help determine error cases, base cases, edge cases, and cases that should run through conditional statements. This is the essence of _unit testing_ and _code coverage_.

Remember that a complex function can usually be composed of various smaller functions. Use this to separate concerns and isolate bugs.

Remember that the development of a solution may imply additional test cases.

**Work the Problem**

Work the problem with the your primary computer: your mind.

Sometimes it is easy to break down the problem into smaller subproblems.

- initializing facts
  - sizes/lengths
- deconstruction

**Develop a Solution**

Understand this may not be the best solution, but it is a solution.

**Test the Solution**

**Analyze the Solution**

Is there another way? Is it better or is it a trade-off?
Do you have the time to consider another solution or is this good enough considering the range of the inputs?
This is the essence of _scaling_ and _project management_.

If the solution needs to improve, then return to _Work the Problem_ with a new attacks.
