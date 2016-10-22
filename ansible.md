ansible
=======

ansible will:
- generate a python script that does a task
- copy the script to hosts
- execute the script on hosts
- wait for the script to complete execution on all hosts

ansible:
- runs each task in parallel across all hosts
- waits until all hosts have completed a task before moving to the next task
- runs tasks in the order you specify them

ansible modules:
- are declarative
- idempotent

ansible:
- uses YAML for configuation
- uses Jinja2 templates

## REFERENCES

Ansible: Up & Running. Lorin Hochstein. O'Reilly. May 2015.
[Ansible Playbooks vs Roles](http://stackoverflow.com/questions/32101001/ansible-playbooks-vs-roles)