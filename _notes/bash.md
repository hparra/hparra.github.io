---
---

# Bash

The Bourne-Again SHell.

## TIPS

`find * | xargs cat`: cat contents of files found

`du -ch *`: show human-readable sizes of files and display total size

`lsof -n -iTCP:$PORT | grep LISTEN`: show who is listening to certain port (Mac OS X)

Use `column` to print things nicely:

## NOTES

[shocco.sh](http://rtomayko.github.io/shocco/) is a documentation for Bash written in a literate programming style.

## REFERENCES

- `man bash` or [bash](http://manpages.debian.org/cgi-bin/man.cgi?query=bash&apropos=0&sektion=0&manpath=Debian+8+jessie&format=html&locale=en)
- `info bash` or [Bash Reference Manual](https://www.gnu.org/software/bash/manual/bashref.html)
- [Bash Hackers Wiki](http://wiki.bash-hackers.org)
- [Advanced Bash-Scripting Guide](http://www.tldp.org/LDP/abs/html/abs-guide.html)
- [GNU Readline](https://en.wikipedia.org/wiki/GNU_Readline)
- [Shell Style Guide](https://google.github.io/styleguide/shell.xml)
