c -- C Programming Language
===========================

C has inspired various other languages.
Languages that are extensions of C include C++ and Objective-C.

Major interpreters and virtual machines written in C:
- CPython (Python)
- Ruby MRI (Ruby)
- Zend Engine (PHP)

Major operating systems written in C:
- Unix
- BSD
- Linux
- Windows
- Darwin (Mac OS X, iOS)

Major systems writtin in C:
- Postgres
- Redis
- Memcached
- Nginx
- Apache

You really should know C.
I recommend learning C before C++ or Objective-C.

## Memory

Memory management is what makes C (and C++) a significantly different beast.
Precautions must be taken, even by programmers transitioning from C++.

- `&` is "address of" -- use on variable to get its memory address
- `*` is "dereference" -- use on pointer variable to get/set value
- `->` is "access member" -- use with struct/union to access member variable. Recall that it is sugar for dotting into dereferenced variable

Use `malloc` to dynamically allocate memory _on the heap_:
- e.g. `int* integers = (int*) malloc(sizeof(int) * 10);`
- only argument is the number of bytes to allocate
- calculate bytes by multiplying `sizeof` data type by number of that data
  - e.g. `sizeof(int) * 10` for an integer array of 10
- returns `void*` (void pointer) which you should cast
- remember it DOES NOT initialize the array
- remember to deallocate memory using `free` function
- remember to set pointer to `NULL` after freeing memory

## REFERENCES

[C dynamic memory allocation](https://en.wikipedia.org/wiki/C_dynamic_memory_allocation). Wikipedia.

[Dangling pointer](https://en.wikipedia.org/wiki/Dangling_pointer) Wikipedia.

[malloc](http://man7.org/linux/man-pages/man3/malloc.3.html). Linux man page. Covers `malloc`, `free`, `calloc`, `realloc`.

[Troubleshooting Segmentation Violations/Faults](http://web.mit.edu/10.001/Web/Tips/tips_on_segmentation.html)  