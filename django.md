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

## Database

DB Tools:
- `manage.py sqlcreate`: create database. Pipe this into your sql client
- `manage.py sqldiff -a`: check difference between models and db schema

Migrations:
- `manage.py makemigrations`: created automatically from your model definitions
- `manage.py showmigrations`
- `manage.py migrate`

Use `manage.py graph_models` to output DOT stream of database modeling.
You can specify default settings in _settings.py_:
```python
GRAPH_MODELS = {
  'all_applications': True,
  'group_models': True,
}
```

Sometimes you'll make changes so drastic in a new project that it's better to redo your migrations:
- `manage.py migrate zero` to 
- delete migration files
- make changes to models
- `manage.py makemigration`

## Django Admin

Django has an admin area (app) by default.
Two Scoops believes its always easier to create a new admin then to override the defaults.
The official Django documentation has information on overriding templates, adding new views, etc.



## Users

- Django comes with its own authentications system.
- Your custom user model can extend `User` via `from django.contrib.auth.models import User`
- Use `create_user(username, email=None, password=None, **extra_fields)`
- Add various default login, et al. view by adding `url('^', include('django.contrib.auth.urls'))` to your `urlpatterns`
- https://docs.djangoproject.com/en/1.10/topics/auth/default/
- `./manage.py createsuperuser`

TODO for implementing your User models:
- Ensure authentication _settings_
  - `INSTALLED_APPS`
    - django.contrib.auth
    - django.contrib.contenttypes
  - `MIDDLEWARE`
    - django.contrib.sessions.middleware.SessionMiddleware
    - django.contrib.auth.middleware.AuthenticationMiddleware
- Change `PASSWORD_HASHERS`? (BCrypt)
- Add `AUTH_PASSWORD_VALIDATORS`?
- Decide on User modeling
  - Option 1: subclass `AbstractUser`
  - Option 2: subclass `AbstractBaseUser`
  - Option 3: relate to appropriate "UserProfile" model
    - Ensure such a "UserProfile" model exists
    - Add to settings `AUTH_USER_MODEL`
      - django.contrib.auth.models.User (?)
    - `user = models.OneToOne(settings.AUTH_USER_MODEL)`
- Add Groups, Permissions?
- Decide on Forms/Views for auth verbs
  - Option 1: use defaults
    - Add `url('^', include('django.contrib.auth.urls'))` to urlpatterns
    - Implement templates
  - Option 2: user custom defaults
    - Import `auth_views` and append each to urlpatterns
    - Implement templates
  - Option 3: Roll your own forms/views
- Add protected views/template
  - Use `@login_required` on view def

## REFERENCES

[Two Scoops of Django: Best Practices for Django 1.8](http://amzn.to/2gLzQlw).  Daniel Roy Greenfeld & Audrey Roy Greenfeld. 2015-05-15.

[Tango with Django](http://leanpub.com/tangowithdjango19). Leif Azzopardi & David Maxwell. 2016-10-04. Seems best for developers without prior MVC experience.

[The Django admin site](https://docs.djangoproject.com/en/1.10/ref/contrib/admin/). Django Documentation.
