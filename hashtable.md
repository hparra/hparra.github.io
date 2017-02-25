Hash Tables
===========

Can we implement a Dictionary such that all operations: insertion, deletion, and searching, are O(1)? 

Yes. Use a hash table.

Hash tables are fundamental data structures.
They're also analogous to the Dictionary ADT, aka associative arrays.
Most dictionaries are implemented with a hash table.
You really should learn them!

I assume you already know about the Dictionary ADT and  [hashing](hashing.md).

## tl;dr

Complexity:
- all operations are O(1) on average
- all operations are O(n) in worst-case (all collisions, etc.)
- space is O(n)

```python
# Naive (and Wrong) Implementation
# but it communicates the concept
class HashTable:

  def __init__(self):
    self.array = []
        
  def set(self, key, value):
    index = hash(key) % len(self.array)
    self.array[index] = value

  def get(self, key):
    index = hash(key) % len(self.array)
    return self.array[index]

# but...
```

Two Fundamental Problems:
1. collisions: _"Someone's already in here!"_
2. array sizing: _"There's no more space!"_

**Collision Resolution**:
- separate chaining: _"Each spot should have more spots!"_
  - linked lists
  - self-balancing binary search trees 
- open addressing: _"Let's stick it in another spot!"_
  - linear probing: _"Just pick the next open spot!"_
  - quadratic probing: Uh...
  - double hashing: _"Hash it again and see if that's open!"_

**Array Resizing**:
  - on insert check array size
  - if array is too small (or getting there) then resize
  - create new array twice as big
  - copy everything over



## Implementations

You may see some hash tables or maps described as _weak_.
This means that hash table uses a weak reference for its values.
Once there are no strong (regular) references to the value the garbage collector is free to eliminate it.

Implementations differ:
- collision resolution:
- array resizing: some 2x others 1.75x

### Python

Python's `dict`:
- is a hash table
- is a built-in type, e.g. `somevar = {}`
- uses open addressing

```python
# Example Python's dict
secrets = {}
secrets["You"] = "idk"
secrets["HGPA"] = "doesn't like watermelon"
secrets["Tony"] = "has a WWF belt"
secrets["Tyler"] = "listen's to Bjork"
del secrets["You"]
for k in secrets:
  print("{} {}").format(k, secrets[k])
for k,v in d.iteritems(): # Python 2
  print("{} {}").format(k, v)
```

https://docs.python.org/2/library/weakref.html

### C++

C++'s `unordered_map`:
- is a hash table

```cpp
#include <iostream>
#include <unordered_map>
using namespace std;
int main() {
  unordered_map<string, string> secrets;
  secrets["You"] = "idk";
  secrets["HGPA"] = "doesn't like watermelon";
  secrets["Tony"] = "has a WWF belt";
  secrets["Tyler"] = "listen's to Bjork";
  delete secrets["You"];
  unordered_map<string,string>::iterator i;
  // actually iterator across pair<string, string>
  for (i = secrets.begin(); i != secrets.end(); i++) {
    cout << i->first << "  " << i->second << endl;
  }
}
```

### Java

As usual Java has various implementations.

Java's [`HashMap<K,V>`](https://docs.oracle.com/javase/7/docs/api/java/util/HashMap.html):
  - uses separate chaining with a Linked List (Java 7)
  - uses separate chaining with a Red-Black Tree (Java 8)
  - is not thread-safe (but see `Hashtable`)

```java
// Example Java's HashMap
import java.util.HashMap;
import java.util.Map;
import java.util.Iterator;
import java.util.Set;
public class Example {
  public static void main(String args[]) {
    HashMap<String, String> secrets = new HashMap<String, String>();
    // setting
    secrets.put("You", "idk");
    secrets.put("HGPA", "doesn't like watermelon");
    secrets.put("Tony", "has a WWF belt");
    secrets.put("Tyler", "listen's to Bjork");
    // getting
    System.out.println(secrets.get("HGPA"));
    // removing
    secrets.remove("You");
    // traverse -- also example why I don't like java
    Set set = secrets.entrySet();
    Iterator iterator = set.iterator();
    while (iterator.hasNext()) {
       Map.Entry entry = (Map.Entry) iterator.next();
       System.out.print(entry.getKey() + " " + entry.getValue());
    }
  }
}
```

## RESOURCES

[Hash table](https://en.wikipedia.org/wiki/Hash_table). Wikipedia.

[The Mighty Dictionary](https://www.youtube.com/watch?v=C4Kc8xzcA68). Brandon Craig Rhodes. PyCon 2010. An optimized explanation of Python's Dictionary implementation using a hash table.

[How std::unordered_map is implemented?](http://stackoverflow.com/questions/31112852/how-stdunordered-map-is-implemented/31113618#31113618). Stack Overflow.

## TODO

[A Proposal to Add Hash Tables to the Standard Library (revision 4)](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2003/n1456.html). Matthew Austern. 2003-04-03.

[How does a HashMap work in JAVA](http://coding-geek.com/how-does-a-hashmap-work-in-java/).

[C++ Tutorial: Intro to Hash Tables](http://pumpkinprogrammer.com/2014/06/21/c-tutorial-intro-to-hash-tables/). 

[Hash Tables (C)](http://www.eternallyconfuzzled.com/tuts/datastructures/jsw_tut_hashtable.aspx). Eternally Confuzzled.


