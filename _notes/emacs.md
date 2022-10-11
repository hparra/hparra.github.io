---
---

# Emacs

[GNU Emacs Reference Card](https://www.gnu.org/software/emacs/refcards/pdf/refcard.pdf)

## Control-based sorted by keyboard layout

- C-q : quoted-insert
- C-w : kill-region
- C-e : end-of-line
- C-r : isearch-backward
- C-t : transpose-chars
- C-y : yank
- C-u : universal-argument
- C-i : _free_
- C-o : open-line
- C-p : previous-line
- C-[ :
- C-] : abort-recursive-edit
- C-| :

---

- C-a : beginning-of-line
- C-s : isearch-forward
- C-d : kill-character-forward
- C-f : forward-char
- C-g : keyboard-quit
- C-h : help-command / kill-character-backward
- C-j : newline-and-indent
- C-k : kill-whole-line
- C-l : recenter
- C-; : _free_
- C-' :

---

- C-z : suspend-emacs / iconify-or-deiconify-frame
- C-x :
- C-c :
- C-v : scroll-up
- C-n : next-line
- C-m :
- C-, :
- C-. :
- C-/ : undo

## "But I don't use emacs anymore..."

**Evaluating expressions and buffers**
- evaluate expression: `c-x c-e`
- some people will bind _eval-buffer_, e.g. `c-x E`
- (Atom) Install [script](https://atom.io/packages/script)
  - bind `c-x c-e` to _Script: Run_
  - _Script: Run_ will evaluate a selection or the entire buffer
  - `c-g` will still close that window
  - remember that 