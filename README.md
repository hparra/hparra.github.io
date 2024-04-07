# HGPA

## wiki

wiki entries. opinionated.

<ul>
{% assign pages = site.pages | where_exp: 'page', 'page.url contains "/wiki"' | sort: 'name' %}
{% for page in pages %}
  <li>
    <a href="{{ page.url }}">{{ page.name | remove: ".md" }}</a>
  </li>
{% endfor %}
</ul>

## notes

notes from books, articles, videos, etc. facts only.

<ul>
{% assign pages = site.pages | where_exp: 'page', 'page.url contains "/notes"' | sort: 'name' %}
{% for page in pages %}
  <li>
    <a href="{{ page.url }}">{{ page.name | remove: ".md" }}</a>
  </li>
{% endfor %}
</ul>

## journal

personal anecdotes.

<ul>
{% assign pages = site.pages | where_exp: 'page', 'page.url contains "/journal"' | sort: 'name' %}
{% for page in pages %}
  <li>
    <a href="{{ page.url }}">{{ page.name | remove: ".md" }}</a>
  </li>
{% endfor %}
</ul>
