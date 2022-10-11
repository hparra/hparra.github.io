# Angular Notes

These are notes to references regarding advanced Angular.js concepts. Working knowledge of Angular is assumed.

## Modules

HTML: `<html ng-app="myApp">`

JS: `var myAppModule = angular.module('myApp', []);`

Modules:
- consist of multiple configuration and run blocks
- can only be loaded once per injector (apps usually only have one)
- can list other modules as their dependencies (array)
- will configure and run each dependency first
- will always run `contsant()` methods first
- will run all other configuration blocks in the order they are registered
- can be unit tested: https://docs.angularjs.org/guide/module#unit-testing

Configuration blocks:
- are executed during the provider registration and configuration phase
- only allow injection of providers and constants (as to prevent accidental instantiation of services before they have been fully configured)
- takes a single function to execute on module load as its only parameter

Run blocks:
- are the closest thing to a main method
- are executed after all services have been configured and after the injector has ben created
- only allow injection instances and constants (as to prevent further system configuration during application run time)
- should be declared in isolated modules because they are difficult to unit test
- takes a single function to execute after injector creation as its only parameter

Convenience methods on module can be used to configure providers.

```
angular.module('myModule', []).
  value('a', 123).
  factory('a', function() { return 123; }).
  directive('directiveName', ...).
  filter('filterName', ...);

// is same as

angular.module('myModule', []).
  config(function($provide, $compileProvider, $filterProvider) {
    $provide.value('a', 123);
    $provide.factory('a', function() { return 123; });
    $compileProvider.directive('directiveName', ...);
    $filterProvider.register('filterName', ...);
  });
```

Further Reading:
- [Requiring vs Browserifying Angular](http://developer.telerik.com/featured/requiring-vs-browerifying-angular/)

## Providers

Providers:
- are registered with the injector
- define how objects are instantiated by injector service
- belong to a module
- are one of two types: services and special purpose

Services:
- are objects whose API is defined by its developer
- are singleton objects (they are instantiated only once per injector)
- are fundamentally based on `$provide`
- include:
  - `provider(providerObject)` (base)
  - `constant(object)`
  - `value(object)`
  - `factory(function)`
  - `service(classObject)` (the poorly named "service" service)
- provide a way to retain data and communicate across controllers
- are used by Angular itself to provide functionality across framework
- are prefixed by a `$` when they are provided by Angular

Special Purpose:
- are objects that conform to a specific Angular framework API
- include:
  - `controller`
  - `directive`
  - `filter`
  - `animation`

_TODO: See "Service v. Factory" notes_

Further Reading:
- [Providers - AngularJS Developer Guide](https://docs.angularjs.org/guide/providers)
- [$provide - AngularJS API](https://docs.angularjs.org/api/auto/service/$provide)
- [Angular.js: service vs provider vs factory?](https://stackoverflow.com/questions/15666048/angular-js-service-vs-provider-vs-factory)


### Provider Services

Provider Services:
- `provide(object)`
- receive a provider object, which:
	- can be an object with a defined `$get` method, which can contain injectables
	- can be a function that returns such an object
	- can be an array of injectable dependecies whose final entry is such a function
- are responsible for registering the service in the $providerCache
- can be externally configured when used directly
- can be injected into `config()` function

```
myModule.provider('someProvider', {
	$get: function () {
		return {
			message: 'hello!'
		}
	}
})

```

### `Factory` Services

Factory Service:
- are registered using `factory(name, function)`
- function is only invoked once
- is really a `provider()` where the `$get` function is passed in

```
myModule.factory('someFactory', function() {
	return {
		message: 'hello!'
	}
})
```

### Serivce Services

Services:
- are registered using `service(name, constructorFunction)`
- are poorly named!
- will invoke the function using the `new` keyword

```
myModule.service('someService', function() {

	var myMessage = 'hello';

	this.getMessage = function() {
		return {
			message: myMessage
		}
	}
})
```

### Constant Services

Constants:
- are registered using `constant(name, value)`
- are injectable
- are not intercepatble by decorators
- can not be objects or functions
- are good for configuration data

### Value Services

Values:
- are registered using `value(name, value)`
- are _not_ injectable
- can register objects or functions

TODO: Decorators

---

## Scopes

Scopes:
- are connections between Angular views and controllers
- are automatically accessble to the view
- are normal javascript objects
- are used as the data model in Angular
- can be nested to provide isolated properties
- can provide execution environment in which view expressions are evaluated
- can provide `observers` to `watch` for model changes
- can propogate model changes by using `apply()`
- have a root ancestor called `$rootScope`
- can be provided to controllers and directives

Notes:
- if a callback executes inside Angular context then `$scope` will be aware of model mutation
- if a callback executes outside Angular context then `$apply()` must be used to notify `$scope` of mutation

Lifecycle:
- on the creation of a controller or directive, Angular creates a new scope with the `$injector`
- at runtime the new scope is passed into the controller or directive
- on linking the scope to the view, all directives that create scopes register their watches with the parent scope
- on the `$rootScope`'s digest all child scopes will perform dirty checking
- when a scope is no longer needed the scrope's creator will need to call `scope.$destroy` to clean up scope
- on scope destruction the `$destroy` event is broadcasted

Further Reading:
- [Make Your Own AngularJS, Part 1: Scopes And Digest](http://teropa.info/blog/2013/11/03/make-your-own-angular-part-1-scopes-and-digest.html). Tero Parviainen.
- ["Controller as" or "$scope"?](http://codetunnel.com/angularjs-controller-as-or-scope/). Alex Ford. July 10, 2014.
- [6 Common Pitfalls Using Scopes](http://thenittygritty.co/angularjs-pitfalls-using-scopes). Jan Philipp. April 24, 2013.

---

## Directives

Directives:
- are custom HTML elements and attributes
- are built into Angular to provide various types of functionality
- are implemented as a special type of factory service
- are `$compile`d when they are parsed at bootstrap
- are fundamental in providing basic Angular features including:
	- `ng-app`
	- `ng-include`
	- `ng-model`
	- `ng-view`
	- `ng-controller`
- use dash-case in HTML
- use camelcase in JavaScript

Further Reading:
- [Creating Custom Directives - AngularJS Developer Guide](https://docs.angularjs.org/guide/directive)
- [$compile - AngularJS API](https://docs.angularjs.org/api/ng/service/$compile)
- [Deep Dive into Custom Directives - NG-Conf 2014 - Video](https://www.youtube.com/watch?v=UMkd0nYmLzY). Dave Smith. January 16, 2014.
- [Writing Direcives](https://www.youtube.com/watch?v=WqmeI5fZcho). Misko Hevery. November 28, 2012.


---

TODO: routing
TODO: testing

---
## Angular Concepts

### Digest Loop

- [Make Your Own AngularJS, Part 1: Scopes And Digest](http://teropa.info/blog/2013/11/03/make-your-own-angular-part-1-scopes-and-digest.html). Tero Parviainen. November 11, 2013.

### Dependency Injection
- [Dependency Injection](https://docs.angularjs.org/guide/di)
- [Vojta Jina - Dependency Injection - NG-Conf](https://www.youtube.com/watch?v=_OGGsf1ZXMs). Vojta Jina. January 16, 2014.

## Angular Conventions

- [An AngularJS Style Guide for Closure Users at Google](https://google-styleguide.googlecode.com/svn/trunk/angularjs-google-style.html)
- [Opinionated AngularJS styleguide for teams](http://toddmotto.com/opinionated-angular-js-styleguide-for-teams/)
- [AngularJS Git Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit)

## Generators & Boilerplate

- [generator-angular](https://github.com/yeoman/generator-angular)
- [ngBoilerplate](https://github.com/ngbp/ngbp)

## Twitter

Who should you follow you ask?

Angular Team:
- [Jeff Cross](https://twitter.com/jeffbcross)
- [Brian Ford](https://twitter.com/briantford)
- [Brad Green](https://twitter.com/bradlygreen) (Product Manager)
- [Misko Hevery](https://twitter.com/mhevery) ("Father of Angular")
- [Vojta Jína](https://twitter.com/vojtajina)
- [Igor Minar](https://twitter.com/IgorMinar)
- [Matias Niemelä](https://twitter.com/yearofmoo) aka "Year of Moo"
- [Julie Ralph](https://twitter.com/SomeJulie) (Protractor)

Angular:
- [AngularJS](https://twitter.com/AngularJS)
- [Protractor](https://twitter.com/ProtractorTest)

Angular Resources:
- [AngularJS](https://twitter.com/AngularJS_News)

## Cheetsheets

- [AngularJS Expressions](http://teropa.info/blog/2014/03/23/angularjs-expressions-cheatsheet.html). Tero Parviainen. March 23, 2014.

## References and Further Reading

- [Tim Kindberg on Angular UI-Router](https://www.youtube.com/watch?v=dqJRoh8MnBo). Tim Kindberg. November 14, 2013.
. [UI-Router - Apps with Depth](http://slides.com/timkindberg/ui-router). Tim Kindberg. November 14, 2013.

- [Writing a Massive Angular Apps at Google - NG-Conf 2014 - Video](https://www.youtube.com/watch?v=62RvRQuMVyg). Ido Sela, Jeremy Elbourn, Rachel Dale. January 16, 2014.
- [Massive AngularJS Apps - NG-Conf 2014 - Presentation](https://docs.google.com/file/d/0B4F6Csor-S1cNThqekp4NUZCSmc/edit). Ido Sela, Jeremy Elbourn, Rachel Dale. January 16, 2014.

Some Browserify + Angular Links:

- http://mindthecode.com/lets-build-an-angularjs-app-with-browserify-and-gulp/
- https://github.com/greypants/gulp-starter/blob/master/package.json#L19
- https://github.com/thlorenz/browserify-shim
- https://medium.com/@dickeyxxx/best-practices-for-building-angular-js-apps-266c1a4a6917