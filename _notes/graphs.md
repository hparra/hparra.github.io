---
---

Graphs
======

Graph algorithms traverse the graph to solve various problems.
There are two different way to traverse a graph:
- Breadth-First Search (BFS)
- Depth-First Search (DFS)

Both BFS and DFS start by visiting a vertex then visiting adjacent vertices, until all vertices have been visited.
They differ in the order in which they accomplish this.

### Breadth-First Search (BFS)

BFS explores oldest unexplored vertices first.
Think LIFO.

```
# BFS in English
Begin with an empty queue of of vertices to be explored,
and add your starting vertex to that list.
For each vertex to explore,
get all edges incident to that vertex,
and for each edge,
get the next vertex,
and if next vertex has not been discovered,
then mark it as discovered,
and add that vertex to queue to be explored.
```

BFS needs:
- a graph -- pass it in
- a starting vertex -- pass it in
- a list to keep track of vertices to be explored -- use a queue
- a way to determine if a vertex has been explored -- use an array or dictionary


### Depth-First Search (DFS)

DFS explores newest unexplored vertices first.
Think FIFO.

```
# Stack-based DFS in English
Begin with an empty stack of of vertices to be explored,
and add your starting vertex to that stack.
While there are vertices in the stack,
pop a vertex off the top,
and if that vertex has not been discovered,
then mark it as discovered,
and get all edges incident to that vertex,
and for each edge get the next vertex,
and push that next vertex onto the stack.
```

Traditionally, DFS is implemented using recursion.
This makes for a much simpler algorithm.
Recursion allows us to "borrow" the function stack.
Also it's easy to erroneously implement a Stack-based DFS,
such that is produces a different DFS tree than recursive DFS.
Just do it recursively unless you have a "Real World" reason not to.

```
# Traditional DFS in English
Begin with a starting vertex,
and for all edges incident to that vertex,
get the next vertex,
and if next vertex has not been discovered,
then mark it as discovered,
and run DFS again on that vertex.
```

### Solving Problems with BFS and DFS

Many graph problems can be solved by modifying BFS or DFS to additional work during traversal:
- when exploring a vertex
- when looking at an edge
- after exploring a vertex

## REFERENCES

[Stack-based graph traversal ≠ depth first search](http://11011110.livejournal.com/279880.html). David Eppstein. 2013-12-17. Clarification on a common misconception held by many, myself included. Fun fact: I never had Eppstein as a professor while at UCI but someone who took his graph theory course told me it actually made them cry.

[Python Patterns - Implementing Graphs](https://www.python.org/doc/essays/graphs/). Building a graph using Python lists and dictionaries.
