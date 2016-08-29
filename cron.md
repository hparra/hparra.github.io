cron
====

View cron for current user: `crontab -l`

## Ways to specify a cronjob

You can place a shell script in one of the following directories:
- _/etc/cron.hourly_
- _/etc/cron.daily_
- _/etc/cron.weekly_
- _/etc/cron.monthly_

You can place a cron script in _/etc/cron.d_

You can edit the user-level cron using `crontab -e`

## REFERENCES

[cron](https://docs.chef.io/resource_cron.html). _Learn Chef_. How to use `chef` to manage cronjobs.