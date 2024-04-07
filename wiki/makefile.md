# Makefiles

> By default, the goal is the first target in the makefile (not counting targets that start with a period). Therefore, makefiles are usually written so that the first target is for compiling the entire program or programs they describe.

```makefile
.PHONY: run
run:
    go run main.go

.PHONY: install
install:
    go mod download
```

## References

[Makefile Tutorial by Example](https://makefiletutorial.com/).
Nice tutorial for make emphasizing original C/C++ use-case.

[Creating a Golang Makefile](https://earthly.dev/blog/golang-makefile/).
Example of a Makefile for go.

[fx/Makefile](https://sourcegraph.com/github.com/uber-go/fx/-/blob/Makefile).
Makefile for go injection framework used at Uber.
Not sure why this was on the first page of my Google search results.

[GNU make](https://www.gnu.org/software/make/manual/html_node/index.html).
The manual!
