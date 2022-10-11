---
---

Java
====

I've forgotten all my Java!

```java
// The bare minimum!
public class HelloWorld {
  public static class main(String[] args) {
    System.out.println("Hello World!"); // stdout
  }
}
```

Primitives:
- `String`
  - Java strings must use double-qoutes!
  - Java uses single-quotes for a single character ala C/C++.
- `boolean`
  - `true` or `false`
- `int`
- `float`
- `double`

Arrays:

```java
// Arrays
int[] array1 = {3,2,1};
Collections.sort(array1);

Comparator<int> comparator = new Comparator<int>() {
    @Override
    public int compare(int left, int right) {
        return left - right;
    }
};
Collections.sort(array1);

int len = array.length;
for (int i = 0; i < len; i++) {
  System.out.println(array[i]);
}
//
int len2 = 10;
int[] array2 = new int[10];
for (int i = 0; i < array2.length; i++) {
  array2[i] = i;
}
```

```java
PriorityQueue<String> queue = 
            new PriorityQueue<String>(10, comparator);

queue.offer(123); // aka push -- add element
queue.size();
queue.peek(); // aka top -- return dominator
queue.poll(); // aka pop -- remove dominator and return it

```

```java
Hashtable<String, Integer> numbers
  = new Hashtable<String, Integer>();
numbers.put("one", 1);
numbers.put("two", 2);
numbers.put("three", 3);
numbers.get("two");
```


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
