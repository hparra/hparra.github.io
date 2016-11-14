#!/usr/bin/env bash
# relative to root
# FIXME: spin up another server that uses a key (why is this no the case already???)
sshpass -p $PASSWORD rsync -vrltz -e ssh ./.gliki/build/ hgpa@hectorparra.com:/var/www/hectorparra.com/html/