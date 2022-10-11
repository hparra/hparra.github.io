selenium
========


```sh
#!/usr/bin/env bash
# selenium server launcher
# place in PATH
java -jar $(cd $(dirname $0); pwd -P)/selenium-server-standalone-2.53.1.jar
```

## Drivers

- [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/downloads).
- [GeckoDriver](https://github.com/mozilla/geckodriver/releases). Firefox driver.
- [SafariDriver](http://www.seleniumhq.org/download/)

## Tools

- [WebdriverIO API](http://webdriver.io/api.html). Binding for NodeJS.

## PaaS

- TestingBot
- Sauce
- Browserstack

### TestingBot

- [Our Selenium configuration](https://testingbot.com/support/other/configuration). Table of server and driver versions used. IPs to whitelist.
- [TestingBot Tunnel](https://testingbot.com/support/other/tunnel). Tunnel download and documentation.

```sh
#!/usr/bin/env bash
# TestingBot tunnel launcher
# place this in your PATH
# See https://testingbot.com/support/other/tunnel
java -jar $(cd $(dirname $0); pwd -P)/testingbot-tunnel/testingbot-tunnel.jar $TB_KEY $TB_SECRET
```

### SauceLabs

If you use `dnsmasq` make sure you DO NOT use _no-resolve_ in configuration. You must specify the localhost as the DNS server, or else `sc` will read _/etc/resolv.conf_ directly: `sc --dns 127.0.0.1`
