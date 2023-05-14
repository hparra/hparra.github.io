# bookmarklets

these things have been around for a long time!
javascript has changed too!

## copy a markdown link

```js
javascript:(async function copyMarkdownLink(s) {
  /* copy a markdown link with page title and href to the clipboard */
  const mdLink = `[${document.title}](${window.location.href})`;
  const type = 'text/plain';
  const blob = new Blob([mdLink], { type });
  const data = [new ClipboardItem({ [type]: blob })];
  await navigator.clipboard.write(data);
})();
```

## copy a markdown citation

i would like to create a citation using a format that i like.

i can use page href, title and description, but many pages have additional data.
some pages have JSON-LD tags and many more have open-graph meta tags.

```js
javascript:(async function copyMarkdownCitation(s) {
  /* copy a markdown link and page citation to the clipboard */
  function meta(q) {
    const el = document.querySelector(`meta${q}`);
    return el && el.content || null;
  }
  const c = {
    title: meta('[property="og:title"]') || document.title,
    href: meta('[property="og:url"]') || window.location.href,
    author: meta('[name="author"]'),
    siteName: meta('[property="og:site_name"]') || window.location.hostname,
    description: meta('[name="description"]'),
  };
  switch (window.location.hostname) {
    case 'en.wikipedia.org':
      c.title = document.title.split(' - ')[0];
      c.author = 'Wikipedia';
      break;
    case 'developer.mozilla.org':
      c.title = document.title.split(' | ')[0];
      c.author = 'MDN';
    default:
      break;
  }
  const mdLink = `[${c.title}](${c.href}). ${c.author}. ${c.siteName}.`;
  const type = 'text/plain';
  const blob = new Blob([mdLink], { type });
  const data = [new ClipboardItem({ [type]: blob })];
  await navigator.clipboard.write(data);
})();
```

## references

[Bookmarklet](https://en.wikipedia.org/wiki/Bookmarklet). Wikipedia. en.wikipedia.org.

[Bookmarklets are Dead…](https://medium.com/making-instapaper/bookmarklets-are-dead-d470d4bbb626). Brian Donohue. Medium.

[Clipboard: write() method - Web APIs](https://developer.mozilla.org/en-US/docs/Web/API/Clipboard/write). MDN. developer.mozilla.org.
