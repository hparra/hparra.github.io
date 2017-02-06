cpp -- C++
==========

_These notes are for people who have not used C++ in a long time -- like myself._

Major interpreters and virtual machines written in C++ include:
- Hotspot (Java)
- CLR (.NET)
- HHVM (PHP/Hack)
- V8 (JavaScript)
- SpiderMonkey (JavaScript)
- Flash

[Rule of Three](https://en.wikipedia.org/wiki/Rule_of_three_%28C%2B%2B_programming%29) -- if a class requires one or more of these, then it probably needs all three:
- destructor
- copy constructor
- copy assignment operator

**Initializer lists**
- constructor can call constructors of members outside of constructor body
- Example: `MyClass::MyClass():someMemberOtherType(1,'param') {}`

There is no true concept of null or nil. `NULL` is simply a macro that expands to the number 0.

A **const function** e.g. `someFunc(param) const {}` makes it a compiler error for class function to change a member variable of the class.
