---
---

# lsof -- list open files

## EXAMPLES

  `lsof -c $PROGRAM` - list files for program like $PROGRAM

  `lsof -i TCP` - list TCP files

## NOTES

Remember to use `-a` to join rules, e.g. `lsof -a -c node -i TCP`

Recall that socket connections are represented as files in *nix systems.

## REFERENCES

[An lsof Primer](https://danielmiessler.com/study/lsof/)