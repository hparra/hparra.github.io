stacks
======

A **stack** mimics a stack of paper.
You add and remove things from the top.
This is also known as "Last In, First Out" (LIFO).

## Requirements

- view top of the stack: `top()` or `peek()`
- remove top of the stack: `pop()`
  - this may or may not return the top item
- add to the top of the stack: `push(item)`

Note that you shouldn't need to (or be able to) iterate through a stack.
Knowing the size of the stack may also be optional.

## Conceptualization

- Keep track of the head, or top of the stack

## Implementation

A list can sometimes be used to mimic a stack:
- for a Python list:
  - `list.append()`
  - `list.pop()`
  - `list[-1]` to access top
- for a JavaScript array:
  - `array.push()`
  - `array.pop()`
  - `array[array.length - 1]` to access top

Other language libraries use another ADT to implement a stack:
- C++ `stack` (uses a `deque`)
- Java's `Stack` (uses a `vector`)

Since lists can be used to implement a stack, it follows that they can also be implemented directly using list implementation techniques.

## Applications

Reversing a string: 
- iterate through a string pushing each letter onto stack
- pop off all elements into a new string

Stack machines, like an interpreter for Reverse Polish Notation.

The _shunting-yard algorithm_ parses infix expression (1 + 2) to produce postfix ones (RPN) or an AST. Uses both stack and queue.

TODO: Backtracking

Compilers & Interpreters:
- space for parameters, local variables, and return address of each function call (_stack frame_)
- compiler's syntax check for matching braces

Virtual (stack) machines:
- JVM (Java, Clojure, Scala, Groovy)
- CLR & Mono (C#)
- CPython (Python)
- Ruby MRI (>=1.9)

## REFERENCES

[`Class Stack<E>`](https://docs.oracle.com/javase/7/docs/api/java/util/Stack.html). Java Collections stack.

[`stack<T, Sequence>`](http://www.sgi.com/tech/stl/stack.html). C++ STL stack.

[Using Lists as Stacks](https://docs.python.org/2/tutorial/datastructures.html#using-lists-as-stacks). Data Structures. Python.

[Array](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array). Javascript reference. MDN.
