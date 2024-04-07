---
---

# Journal

These are some of my journal entries.
Anything useful from these usually end up in [notes](../notes/).
HGPA

## index

<ul>
{% assign pages = site.pages | where_exp: 'page', 'page.url contains "/journal"' | sort: 'name' %}
{% for page in pages %}
  <li>
    <a href="{{ page.url | remove: '.html' }}">{{ page.name | remove: ".md" }}</a>
  </li>
{% endfor %}
</ul>
