# Design Patterns in Golang

## Solid Principles

- `Single Responsibility` A type should have one primary responsibility, and as a result it should have
  one reason to change that reason being somehow related to it's primary responsibility.
- `Open-Closed` Types should be open for extension (grab the interface and implement this somewhere in your code)
  but closed for modication. We don't want to intefere with what has already been written and tested.
- `Liskov Substitution` If you have some API that takes a base class,
  and works correctly with that base class, it should also work correctly with the derived classes.
- `Interface Segregation` You shouldn't put too much into a interface (break into several interfaces).
- `Dependency Inversion` HLM (High Level Module) should not depend on LLM (Low Level Module). Both should depend on abstractions.