JavaScript Style
================

Please see [feross/standard's RULES.md](https://github.com/feross/standard/blob/master/RULES.md) for a general guide of form to follow for javascript. One of the fundamental concepts is the elimination of semicolons but there are some caveats.

Beside that:

- functions and class methods should be in camelCase
- pure functions should be in their own file and have their own unit tests
- named imports should be broken into multiple lines
- files that export classes, static or instantaible, should have filenames UpperCamelCase
- import/includes of classes should be in UpperCamelCase, even if they are not
- do not use `var` anymore (we have `const` and `let`)


## TIPS

- if you're using `let` you may want to reconsider your function design

This is ugly, why the JS folks do not want to improve this potentially awesome expression is beyond me.

```javascript
const select = (function (x) => {
  switch(x) {
    case 'A': return 0
    case 'B': return 1
    default:  return 9
  }
})(y)
```
