JavaScript
==========

The most popular JavaScript interpreter is V8 used by Google Chrome and NodeJS.
Others include SpiderMonkey (Firefox), JavaScriptCore aka Nitro (Safari) and Chakra (Microsoft Edge).

## Array

**Normal**

A sequential array, e.g. `[1,2,3]`, is implemented in a linear storage buffer (memory). Once array is full it increases in size. While convention is for it to double in size implementations differ.

**Sparse**

JS allows you to specify any i-th element in an array.
For example:

```javascript
const array = [];
array[0] = 1;
array[1] = 2;
array[3] = 3;
```

In this case `array[2]` is not defined.
It is an _array hole_ and this type of array is called a _sparse array_.
In this case V8 will use a hash table to store the elements of the array.
This eliminates the space wasted by empty spots.

HYPOTHESIS: Since a sparse array is implemented as an actual hash table, it may be faster to use an array instead of an object when you desire a hash table for your own application with numeric keys, e.g. sequential primary IDs. For example:

```javascript
const map = []
map[532] = 'blah'
map[2727] = 'foo'
```

RESULT: No difference in performance. Why?

- http://andrewdupont.net/2006/05/18/javascript-associative-arrays-considered-harmful/
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map

## Object

- implemented using _shape trees_
- can be used as a hash table
  - but its better to use an array for this

(TODO)

## REFERENCES

[How JavaScript Objects are Implemented](https://www.infoq.com/presentations/javascript-objects-spidermonkey). Eddy Bruel. MLOC.JS. 2014-06-07. Explanation on implementing SpiderMonkey's Objects as shape trees. After 20 minutes there is discussion regarding arrays.

[Breaking the JavaScript Speed Limit with V8](https://www.youtube.com/watch?v=UJPdhx5zTaw). Daniel Clifford. Google I/O 2012. 2012-06-29.

[The pitfalls of using objects as maps in JavaScript](http://www.2ality.com/2012/01/objects-as-maps.html). Axel Rauschmayer. ②ality – JavaScript and more. 2012-01-03.

[V8: an open source JavaScript engine](https://www.youtube.com/watch?v=hWhMKalEicY). Lars Bak. Google. 2008-08-15.

[Writing Fast, Memory-Efficient JavaScript](https://www.smashingmagazine.com/2012/11/writing-fast-memory-efficient-javascript/) Addy Osmani. Smashing Magazine. 2012-10-05.
