tmux -- terminal multiplexer
============================

## SYNOPSIS

### Shell Commands

- start a new session: `tmux`
- start a new named session, e.g. `tmux new -s mynewsession`
- list sessions: `tmux ls`
- attach to a session by number, e.g. `tmux at 0`
- attach to a session by name, e.g. `tmux at -t somename`
- kill a session by name, e.g. `tmux kill-session -t somename`

### Key Bindings

The default command key prefix is `C-b`. Note it conflicts with any Emacs-style bindings.

- `s` list/switch sessions
- `$` name session
- `c` new window
- `,` name window
- `f` find window
- `&` kill window
- `%` split window vertically (columns)
- `"` split window horizontally (rows)


## DESCRIPTION

tmux:
- is a terminal multiplexer
- eliminates the need for multiple terminal windows
- lets you switch easily between several programs in one terminal
- lets you detach programs and keep them running in the background
- lets you reattach programs to a different terminal
- works well with emacs and vi

There are three ways to interact with _tmux_:
- from the shell using `tmux`
- through the tmux console
- through key bindings

## Shell Interactions

- start a new session: `tmux`
- start a new named session, e.g. `tmux new -s mynewsession`
- list sessions: `tmux ls`
- attach to a session by number, e.g. `tmux at 0`
- attach to a session by name, e.g. `tmux at -t somename`
- kill a session by name, e.g. `tmux kill-session -t somename`

There are three layers of organization in _tmux_:
- sessions
- windows (tabs)
- panes (splits)

## Session Interactions
- :new<CR>  new session
- s list sessions
- $ name session

## Window Interactions

Keybindings:
- c new window
- , name window
- f find window
- & kill window

## Tips

Mac OS X has some problems with copying and pasting. See [Notes and workarounds for accessing the Mac OS X pasteboard in tmux sessions](https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard).

## References

[tmux cheatsheet](https://gist.github.com/henrik/1967800)

[tmux man page](http://www.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man1/tmux.1?query=tmux&sec=1)

[tmux Tutorial - Split Terminal Windows Easily](http://lukaszwrobel.pl/blog/tmux-tutorial-split-terminal-windows-easily)