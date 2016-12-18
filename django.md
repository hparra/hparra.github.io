django
======

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

[Two Scoops of Django: Best Practices for Django 1.8](http://amzn.to/2gLzQlw).  Daniel Roy Greenfeld & Audrey Roy Greenfeld. 2015-05-15.

[Tango with Django](http://leanpub.com/tangowithdjango19). Leif Azzopardi & David Maxwell. 2016-10-04. Seems best for developers without prior MVC experience.