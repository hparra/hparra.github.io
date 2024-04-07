---
---

# notes

Contained herein are notes on various subjects.
Some are complete or have had many iterations.
Others are simply scratch.
Most are artifacts of my own learning or documentation of some kind.
Perhaps you may find some of them useful.

## selected directory tree

- [editors](editors.md)
  - [emacs](emacs.md)
  - [vscode](vscode.md)
- langauges

## index

<ul>
{% assign pages = site.pages | where_exp: 'page', 'page.url contains "/wiki"' | sort: 'name' %}
{% for page in pages %}
  <li>
    <a href="{{ page.url  | remove: '.html' }}">{{ page.name | remove: ".md" }}</a>
  </li>
{% endfor %}
</ul>

## references

[hgpa](/README.md)
