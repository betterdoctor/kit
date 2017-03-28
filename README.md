kit
===

`kit` is BetterDoctor's Golang standard library.


### Design Philosophy

Packages must be purposeful, usable, and portable.

* **To be purposeful, packages must provide, not contain.**
    * Packages must be named with the intent to describe what it provides.
    * Packages must not become a dumping ground of disparate concerns.
* **To be usable, packages must be designed with the user as their focus.**
    * Packages must be intuitive and simple to use.
    * Packages must respect their impact on resources and performance.
    * Packages must protect the user's application from cascading changes.
    * Packages must prevent the need for type assertions to the concrete.
    * Packages must reduce, minimize and simplify its code base.
* **To be portable, packages must be designed with reusability in mind.**
    * Packages must aspire for the highest level of portability.
    * Packages must reduce taking on opinions when it's reasonable and practical.
    * Packages must not become a single point of dependency.

The above borrowed from [Ardan Labs Go Training](https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/packaging/README.md)

