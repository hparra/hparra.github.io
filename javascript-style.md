JavaScript Style
================

Please see [feross/standard's RULES.md](https://github.com/feross/standard/blob/master/RULES.md) for a general guide of form to follow for javascript. One of the fundamental concepts is the elimination of semicolons but there are some caveats.

Beside that:

- functions and class methods should be in camelCase
- pure functions should be in their own file and have their own unit tests
- files that export classes, static or instantaible, should have filenames UpperCamelCase

- do not use `var` anymore (we have `const` and `let`)
- always leave newline between function or method declarations

- prefer `null` over an empty string

`import`/`require` of classes should be in UpperCamelCase, even if they are not:
```javascript
// OK
const Path = require('path')

// Avoid
const path = require('path')
```

Avoid multiple named `import`s on one line:
```javascript
// OK
import {
  thing1,
  thing2
} from 'whatever';

// Avoid
import { thing1, thing2 } from 'whatever'
```

Avoid defining anonymous static object as function parameter. This improves debugging and avoid Promise-related clutter but I understand the shortcut: 
```javascript
// OK
const myObject = {
  hello: 'Hi!'
}
someFunc(myObject);

// Avoid
someFunc({
  hello: 'Hi!'
})
```

Avoid checking boolean values with expressions that check for value:
```javascript
const myBoolean = true;

// OK
if (myBoolean) {}

// Avoid
if (myBoolean === true) {}
```

## Nice Things

```javascript
  // Best
  return {
    user:         state.user,
    channel:      state.channel,
    newuser:      state.newuser,
    creditcard:   state.creditcard,
    coupon:       state.coupon,
    queryparams:  state.queryparams,
  }

  // OK
  return {
    user: state.user,
    channel: state.channel,
    newuser: state.newuser,
    creditcard: state.creditcard,
    coupon: state.coupon,
    queryparams: state.queryparams,
  }
```

## TIPS

- if you're using `let` you may want to reconsider your function design

This is ugly, why the JS folks do not want to improve this potentially awesome expression is beyond me.

```javascript
var select = (x => {
  switch(x) {
    case 'A': return 0
    case 'B': return 1
    default:  return 9
  }
})(y)
```
