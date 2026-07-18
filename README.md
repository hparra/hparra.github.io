# www.hgpa.tv

repo: https://github.com/hparra/hparra.github.io  
settings: https://github.com/hparra/hparra.github.io/settings/pages  

Built with [gliki](plugins/site/gliki.ts), a small git-based bliki generator,
and published to GitHub Pages via Actions. Content lives in [docs/](docs/).

```sh
npm install     # once
npm run serve   # build + serve at http://127.0.0.1:4000
npm run build   # build to ./public
npm run ci      # typecheck + lint + test
```
