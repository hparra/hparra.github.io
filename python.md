python -- Notes on the Python programming Language
==================================================

- Python has an excellent [language reference](https://docs.python.org/2.7/reference/index.html).
- Python has a [style guide](https://www.python.org/dev/peps/pep-0008/).

## Implementations of Fundamental Types

List:
- e.g. `[0,1,2]`
- is a sequence
- is implemented as a variable-length array
  - NOT a linked list
- is a different data structure from `array`
  - See https://docs.python.org/2/library/array.html

Tuple:
- e.g. `(0,1,2)`
- is a sequence
- are immutable

[Set](https://docs.python.org/2.7/library/stdtypes.html#set):
- e.g. `set([0,1,2])`
- e.g. `{0,1,2}`
- is a sequence

Dictionary:
- e.g. `{ 'a':1 , 'b':2 }`
- implemented with resizable hash tables
- any immutable object can be used as a dictionary key
- python exposes its hash function as the global `hash()` (32-bit)

## Built-in functions

Use `range(a)` for loops.
Note that `range(a,b)` returns an empty array if a and b are equal.

`slice(start, stop)` but use `a[start:stop]` instead

Numbers:
- `int(x[, base])`: constructor that can also be used to change base
- `chr`: character from integer as ASCII code
- `divmod(a,b)`: return `(a // b, a % b)` for integers
- `bin`, `oct` and `hex`
- `abs`
- `pow`: but use `x**y` instead

String:
- `ord`: return n

Iterables:
- `len` for length
- `next(iterator[, default])`, e.g. `next(i, None)`
- `all` and `any`: returns true if all/any element are true
- `sorted`
- `min` and `max`
- `map(function, iterable)` and `reduce(function, iterable[, initializer])`
- `filter(function, iterable)` but use [list comprehensions](https://docs.python.org/2.7/tutorial/datastructures.html#list-comprehensions) instead
- `tuple` to create a tuple from iterable
- `zip([iterable, ...])`: return list of tuples by "zippering" multiple lists
- `enumerate`: return list of numbered tuples
- Unfortunately, there is no `some`

## imports

- `from MODULE import NAME1, NAME2, ...` to import NAME1 and NAME2 from MODULE
- `from MODULE import *` to import everything from MODULE
- `from .MODULE` is relative import (?)

## Concurrency

- Python uses reference counting
GIL: Global Interpreter Lock.
- There is only one lock!

## TIPS

Functions can return multiple values:
```python
def return_both(a,b):
  return a,b
x,y = return_both(1,2)
```

You can put multiple statements on one line by using a semi-colon, e.g. `i += 1; j += 1` but this is discouraged.
Perhaps only do this on a whiteboard or doc to save space.

Get an object's reference count: `sys.getrefcount(obj)`

## RELATED TOOLS

### pip

`pip` is the python package manager

- `sudo easy_install pip` to install `pip`
- `pip` will show you available commands
  - `pip search <keyword>`
  - `pip install <package>`
  - `pip freeze` prints installed packages in "requirements" format
    -  These are usually saved to _requirements.txt_
- `pip help <command>` will show you extended --help for a command

`setuptools` is the fundamental pip dependency

### virtualenv

`virtualenv` creates a virtual environment for each program with:
  - its own `python`
  - its own `pip`

Installation:
  - `pip install virtualenv`
  - `pip install virtualenvwrapper`
  - `mkvirtualenv <name_of_env>` to create a virtual env
  - `workon <name_of_env>` to source (use) a virtual env

Virtual environments are _not_ saved in you application folder.
They are located in *~/.virtualenvs/name_of_env*
When you `pip install` a package that contains global binary then it is isntall in *~/.virtualenvs/name_of_env/bin*

## REFERENCES

[A Byte of Python](https://python.swaroopch.com/). Swaroop Chitlur.

[A non-magical introduction to Pip and Virtualenv for Python beginners](https://www.dabapps.com/blog/introduction-to-pip-and-virtualenv-python/). Jamie Matthews. 2013-04-18.

[Design and History FAQ](https://docs.python.org/2/faq/design.html). Python Documentation. 

[The Mighty Dictionary](https://www.youtube.com/watch?v=C4Kc8xzcA68). Brandon Craig Rhodes. PyCon 2010. An optimized explanation of Python's Dictionary implementation using a hash table.
