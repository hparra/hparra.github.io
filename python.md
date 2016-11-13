python
======

## python

- `from MODULE import NAME1, NAME2, ...` to import NAME1 and NAME2 from MODULE
- `from MODULE import *` to import everything from MODULE
- `from .MODULE` is relative import (?)

## pip

`pip` is the python package manager

- `sudo easy_install pip` to install `pip`
- `pip` will show you available commands
  - `pip search <keyword>`
  - `pip install <package>`
  - `pip freeze` prints installed packages in "requirements" format
    -  These are usually saved to _requirements.txt_
- `pip help <command>` will show you extended --help for a command

`setuptools` is the fundamental pip dependency

## virtualenv

`virtualenv` creates a virtual environment for each program with:
  - its own `python`
  - its own `pip`

Installation:
  - `pip install virtualenv`
  - `pip install virtualenvwrapper`
  - `mkvirtualenv <name_of_env>` to create a virtual env
  - `workon <name_of_env>` to source (use) a virtual env

Virtual environments are _not_ saved in you application folder.
They are located in *~/.virtualenvs/name_of_env*
When you `pip install` a package that contains global binary then it is isntall in *~/.virtualenvs/name_of_env/bin*

## django

NOTE: A Django project may contain one or more apps

`pip install Django` will install Django

`django-admin.py startproject <name_of_project>` will create a new Django _project_ with the following directory structure:
  - name_of_app
    - __init__.py: empty script to tell Python interpreter that the directory is a Python package
    - settings.py: Django project settings
    - urls.py: contains the _urlpatterns_ list which routes URLs to views
    - wsgi.py: the development server
  - `manage.py`: the Django task runner script

`manage.py` is the same as `django-admin.py` plus:
  - puts your project’s package on sys.path
  - sets the DJANGO_SETTINGS_MODULE environment variable so that it points to your project’s settings.py file

See [django-admin and manage.py](https://docs.djangoproject.com/en/1.9/ref/django-admin/#django-admin-py-and-manage-py)

`python manage.py startapp <name_of_app>` will create a new Django _app_ within this project with the following directory structure:
  -  __init__.py
  - admin.py: where you can register your models so that you can benefit from some Django
  machinery which creates an admin interface for you;
  - apps.py: that provides a place for any application specific configuration;
  - models.py: a place to store your application’s data models - where you specify the entities
  and relationships between data;
  - tests.py 
  - views.py
  - migrations/

## REFERENCES

- [A non-magical introduction to Pip and Virtualenv for Python beginners](https://www.dabapps.com/blog/introduction-to-pip-and-virtualenv-python/)
- http://www.thomas-cokelaer.info/tutorials/python/basics.html