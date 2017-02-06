Gliki
=====

Gliki is a git-based bliki (blog + wiki). 

_WARNING: Gliki was a quick weekend hack!_

Gliki reads Markdown files stored in a git repository and converts each one into an HTML file using a Handlebars template. Metadata from `git`, particularly file creation and modification dates, are used to sort the bliki table of contents.

Each markdown file will be converted into HTML using a Handlebars file of the same name. If this file does not exist then Gliki will use `__default__.hbs`.

Like Github, `README.md` files are analogous to `index.html`. This file will not appear in the Gliki table of contents.

Dependencies:
- git
- nodejs
