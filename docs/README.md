# HGPA

## wiki

wiki. opinionated.

<ul>
{{#each (section "wiki")}}
  <li><a href="{{href}}">{{name}}</a></li>
{{/each}}
</ul>

## notes

notes from books, articles, videos.

<ul>
{{#each (section "notes")}}
  <li><a href="{{href}}">{{name}}</a></li>
{{/each}}
</ul>

## journal

personal anecdotes.

<ul>
{{#each (section "journal")}}
  <li><a href="{{href}}">{{name}}</a></li>
{{/each}}
</ul>

