---
---

Caching
=======

> There are only two hard things in Computer Science: cache invalidation and naming things.

-- [Attributed to Phil Karlton](http://martinfowler.com/bliki/TwoHardThings.html)

General caching:
- [Cache](https://en.wikipedia.org/wiki/Cache_(computing))

Web-specific caching:
- [HTTP Caching](https://developers.google.com/web/fundamentals/performance/optimizing-content-efficiency/http-caching)
- [Caching best practices & max-age gotchas](https://jakearchibald.com/2016/caching-best-practices/)
- [Caching Tutorial for Web Authors and Webmasters](https://www.mnot.net/cache_docs/)
- [A Guide to Caching with NGINX and NGINX Plus](https://www.nginx.com/blog/nginx-caching-guide/)
  - [Cache Placement Strategies for NGINX and NGINX Plus](https://www.nginx.com/blog/cache-placement-strategies-nginx-plus/)
  - [Nginx Caching](https://serversforhackers.com/nginx-caching)
- [Caching Guide (Apache)](http://httpd.apache.org/docs/2.4/caching.html)
- [Web Caching Basics: Terminology, HTTP Headers, and Caching Strategies](https://www.digitalocean.com/community/tutorials/web-caching-basics-terminology-http-headers-and-caching-strategies)

Framework-specific caching:
- [Django’s cache framework](https://docs.djangoproject.com/en/1.10/topics/cache/)
- [Client side caching (Hapi.js)](http://hapijs.com/tutorials/caching)

Platform-specific:
- [Understanding caching in Postgres - An in-depth guide](https://madusudanan.com/blog/understanding-postgres-caching-in-depth/)
- [Caching Dependencies and Directories (Travis CI)](https://docs.travis-ci.com/user/caching/)
- [CDN Caching](https://www.incapsula.com/cdn-guide/cdn-caching.html)

Advanced:
- [Caching for a Global Netflix](http://techblog.netflix.com/2016/03/caching-for-global-netflix.html)

---

Belady’s optimal page replacement policy

- NRU: Not Recently Used

- LRU: Least Recently Used
- LFU: Least Frequently Used
- FIFO: First In, First Out
- MRU: Most Recently Used
- Second-chance
- Clock


- https://www.varnish-cache.org/trac/wiki/ArchitectureLRU