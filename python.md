python
======

## Implementations of Fundamental Types

### List

- is a sequence
- is implemented as a variable-length array
  - NOT a linked list
- is a different data structure from `array`
  - See https://docs.python.org/2/library/array.html

### Tuple

- is a sequence
- are immutable

## Set

- is a sequence

### Dictionary

- implemented with resizable hash tables
- any immutable object can be used as a dictionary key
- python exposes its hash function as the global `hash()` (32-bit)

## Concurrency

- Python uses reference counting
GIL: Global Interpreter Lock.
- There is only one lock!

## TIPS

`sys.getrefcount(obj)`

## python

- `from MODULE import NAME1, NAME2, ...` to import NAME1 and NAME2 from MODULE
- `from MODULE import *` to import everything from MODULE
- `from .MODULE` is relative import (?)

## pip

`pip` is the python package manager

- `sudo easy_install pip` to install `pip`
- `pip` will show you available commands
  - `pip search <keyword>`
  - `pip install <package>`
  - `pip freeze` prints installed packages in "requirements" format
    -  These are usually saved to _requirements.txt_
- `pip help <command>` will show you extended --help for a command

`setuptools` is the fundamental pip dependency

## virtualenv

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

- [A non-magical introduction to Pip and Virtualenv for Python beginners](https://www.dabapps.com/blog/introduction-to-pip-and-virtualenv-python/)
- http://www.thomas-cokelaer.info/tutorials/python/basics.html

[Design and History FAQ](https://docs.python.org/2/faq/design.html). Python Documentation. 

[The Mighty Dictionary](https://www.youtube.com/watch?v=C4Kc8xzcA68). Brandon Craig Rhodes. PyCon 2010. An optimized explanation of Python's Dictionary implementation using a hash table.
