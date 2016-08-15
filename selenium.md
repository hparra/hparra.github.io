selenium
========


```sh
#!/usr/bin/env bash
# selenium server launcher
# place in PATH
java -jar $(cd $(dirname $0); pwd -P)/selenium-server-standalone-2.53.1.jar
```

## PaaS

- TestingBot
- Sauce
- Browserstack

### TestingBot

[TestingBot Tunnel](https://testingbot.com/support/other/tunnel)

```sh
#!/usr/bin/env bash
# TestingBot tunnel launcher
# place this in your PATH
# See https://testingbot.com/support/other/tunnel
java -jar $(cd $(dirname $0); pwd -P)/testingbot-tunnel/testingbot-tunnel.jar $TB_KEY $TB_SECRET
```