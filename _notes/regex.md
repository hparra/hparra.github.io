---
---

regex -- Regular Expressions
============================

Every programmer should learn to use these,
but approach them with the respect of learning another programming language,
because that's really what they are.
Get a good book.

There are two popular regex styles:
- Perl Compatible Regular Expressions (PCRE)
- POSIX

PCREs in some form are used natively in some languages, e.g.:
- javascript
- python
- java

Beware that even PCREs amongst languages have slight feature differences and quirks:
- what `.` actually matches
- multiple-line support
- back references
- advanced features

## Review

PCRE basics:
- Characters or Tokens
  - any literal (with some exceptions)
  - escape special characters or rules when you want to use literally
  - Specials:
    - `.` any character
    - `^` beginning of line
    - `$` end of line
- Sets are groups (or ranges) of characters
  - `[ ]` positive set
  - `[^ ]` negative set
  - Use `-` in a set to denote range between characters
  - character classes are sugar for sets:
    - `\d` equivalent to `[0-9]`
    - `\w` equivalent to `[a-zA-Z0-9_]`
    - `\s` equivalent to `[ \t\r\n\f]`
- Groups
  - Use these to group or capture
  - `()` capturing group
  - `(?:)` non-capturing group
  - `(?P<var>)` named capturing group (named _var_)
- Alternation (Boolean or) `|`
  - Separate items, sets or unions
- Quanifiers
  - `{M}` exactly M times
  - `{M,}` mininum M times
  - `{M,N}` minimum M maximum N times
  - "greedy quantifiers" are just sugar:
    - `*` equivalent to `{0,}`
    - `+` equivalent to `{1,}`
    - `?` equivalent to `{0,1}`
- Assertions
  - Lookahead
    - `x(?=y)` match x only if it is followed by y
    - `x(?!y)` match x only if it is not followed by y
  - Lookbehind
    - `(?<=y)x` match x only if it follows y
    - `(?<!y)x` match x only if it does not follow y

## Implementations

- JavaScript is missing quite a few PCRE features:
  - lookbehind assertions (maybe in future)
  - named groups
  - (and more)

## Notes on Bash

Bash globbing (filename expansions) look similar to regex but are not.
They are still very useful:
- `?` single character wildcard
- `*` multi character wildcard
- `[ ]` bracket expression
- `[^ ]` negated bracket expression
- `[-]` options range
  - `^` negate range (place inside bracket)

## TODO:

- back references
- boundaries
- Can you parse HTML with regex? (Regular Grammars & DFA/NFA)
  - http://stackoverflow.com/questions/1732348/regex-match-open-tags-except-xhtml-self-contained-tags/1732454#1732454
  - http://stackoverflow.com/questions/590747/using-regular-expressions-to-parse-html-why-not
  - http://web.mit.edu/6.005/www/fa15/classes/17-regex-grammars/
  - https://perl6.org/archive/doc/design/apo/A05.html
  - https://swtch.com/~rsc/regexp/regexp1.html
  - https://github.com/google/re2/wiki/Syntax
- recursion

REFERENCES
==========

[Globbing](http://www.tldp.org/LDP/abs/html/globbingref.html). Regular Expressions. Advanced Bash-Scripting Guide.

[Filenames and Pathnames in Shell: How to do it Correctly](http://www.dwheeler.com/essays/filenames-in-shell.html). David A. Wheeler. 2016-05-04.

[Perl Style Regular Expressions in Prolog](http://www.cs.sfu.ca/~cameron/Teaching/384/99-3/regexp-plg.html). Robert D. Cameron. CMPT 384 Lecture Notes. 1999.

[Regular Expression](https://en.wikipedia.org/wiki/Regular_expression). Wikipedia.

[RegExp](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp). MDN.

[RegExp lookbehind assertions](http://v8project.blogspot.com/2016/02/regexp-lookbehind-assertions.html). Yang Guo. V8 JavaScript Engine. 2016-02-26.
