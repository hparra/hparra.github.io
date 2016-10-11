dnsmasq
=======

## Installing

`brew install dnsmasq`

## Modifying dnsmasq

Since there are a few ways to install `dnsmasq` we can't assume certain paths.

Find all related files using `find / -iname "dnsmasq" 2>/dev/null`. This may take a while and won't display additional finds in subdirectories.

If already loaded _plist_ then use `launchctl list | grep dnsmasq` to determine name used registered with `launchd`. First number is currnet PID, seconde is current status.

The CLI for `launchctl` has changed in recent versions. The load/unload/start/stop commands are deprecated. `launchctl stop` will still stop a daemon, but will allow it to immediately restart if configured that way. Use it to restart after config change.

_/Library/LaunchDaemons/dev.dnsmasq.plist_ is the configuration of dnsmasq for `launchd`

For a `boxen`-installed `brew`: 

```sh
# edit config file
sudo emacs /opt/boxen/config/dnsmasq/dnsmasq.conf

# reload
sudo launchctl stop dev.dnsmasq
```