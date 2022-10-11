io -- Input/Ouput
=================

| Language   | Input | Output | Notes |
| ---------- | ----- | ------ | ----- |
| C          | `scanf("%s", &x)` | `printf("%s", x)` | Use `getline` for strings  
| C++        | `x >> cin` | `cout << x` | |
| Python     | `x = sys.stdin.readline()` | `print(x)` | See also `raw_input` and `input`
| JavaScript | `process.stdin.on('data', func)` | `console.log(x)` | Reading from stdin is always async and only available in NodeJS.|

## NOTES

### C

printf uses format specifiers that follow `%[flags][width][.precision][length]specifier`

`scanf` with "promote" a float to a double, so you can use `%f` for both, but in C99 you may explicitly use `%lf` for a double.

There are various C functions to read data from a stream:
- `scanf` reads any data, terminating at whitespace, newlines or EOF
- `gets` (deprecated) reads only character data, terminating at newlines or EOF
- `fgets` is the safer version of `gets`, requiring max numbers of characters, and works for any stream
- `getchar` read a single character
- `getline` (POSIX) reads a string and returns a pointer to dynamically allocated memory with that string
- `getdelim` (POSIX) allow you to specify the delimiter -- `getline` is equivalent to `getdelim('\n')`

TIP: Beware of using `scanf` followed by a different function to read a string, as `scanf` will leave the newline in the stream buffer. The following call, e.g. `fgets` will encounter this newline and immediately return. You can use `getchar` to eat this newline first.

### Python

`raw_input` is similar to a `getline` but is only valid in Python 2.

Use `sys.stdout.write` for more control.
Remember to also use `sys.stdout.flush` to ensure buffer is written immediately.

Some quickies:
- read an integer on a line: `int(raw_input().strip())`
- read array of integer from space-separated integers on a line: `map(lambda x: int(x), raw_input().strip().split(' '))`

`print(",".join(map(str, array)))`

### Java

- `System.out` for STDOUT
- `System.in` for STDIN

```java
// Read a string and integer from stdin
Scanner scanner = new Scanner(System.in);
String str = scanner.next();
int num = scanner.nextInt();
scanner.close();
```

```java
// Read until EOF
Scanner scanner = new Scanner(System.in);
while (scanner.hasNext()) {
    int a = scanner.nextInt();
    System.out.println(a);
}
```