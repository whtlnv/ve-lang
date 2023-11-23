# The EV programming language

## Guiding principle

> This is the core principle that guide the design of the EV language. It is immutable and is not subject to change. If a software engineer does not agree with this principle, she will have little use for the language.

The foundation of the EV language can be seen from 2 different perspectives:

### Testability

> Software units and their compositions should be trivially tested

This was the original motivation for the language, pushing the testability of the code as far as possible; specifically at the edges of the system.

Good automated software tests are a core component of any non-trivial software project. But writing good automated software tests is hard. The best we can do, at the moment, is to write tests before implementing a solution. This makes the process easier, but it is still hard when testing at the edges of the SUT.

#### Goal

The EV language should allow the engineer to write tests with minimal knowledge of the implementation details of the SUT. This will benefit the engineer when:

- First implementing a solution by requiring minimal effort to write the tests, maximizing the time spent actually solving the problem.
- Maintaining the solution by allowing the engineer to quickly write tests that protect her when refactoring the code.

#### Example (trivial)

**Requirement:** Write a function that prints the string "Hello World!" to the console.

```
// TEST 1.1
when ("printHello() is called")
  it ("should print 'Hello World!' to the console")
    printHello()
    ??? // How do we test this?
```

Most programming languages will allow the engineer to interface with the machine I/O through a function call of an existing standard library. Spying on this function call can be challenging or outright impossible in some of them.

This can be solved with many different approaches. To name a few:

##### Testing the business logic of the SUT.

We can call on a design principle such as [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) to justify this approach.

```
// TEST 1.1.1
when ("buildHello() is called")
  it ("should return 'Hello World!'")
    got = buildHello()
    want = "Hello World!"

    assert(got == want)
```

This test diminish the scope of the SUT, since its responsible of forming the string "Hello World!" but not of printing it to the console. We did cover some ground with this test, but **the responsibility of testing that the string is printed to the console is pushed outwards.**

The engineer might be satisfied with this coverage and it might be enough for the requirements of the project, but the required behaviour of the SUT is not fully tested. For that, we can use any of the following approaches:

##### Making the SUT depend on the implementation of the interface.

This is a common pattern known as [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection).

```
// TEST 1.2.1
when ("printHello() is called")
  it ("should invoke its arg with 'Hello World!'") // !!!
    printToConsole = mock()

    printHello(printToConsole)

    assertCalledWith(printToConsole, "Hello World!")
```

This test diminish the scope of the SUT ever so slightly. The SUT is responsible of forming the string "Hello World!" and invoking the `printToConsole` function, but what that function does is not the responsibility of the SUT. Ideally, the given function will be also tested, but **the responsibility of testing that a function that actually prints to the console is passed to the SUT is pushed outwards.**

Adding an abstraction layer to the problem we can solve it if our language allow us to compare a function reference.

```
// TEST 1.2.2
when ("helloController() is called")
  it ("should invoke printHello() with system.console")
    printHello = mock()

    helloController(printHello).sayHello()

    assertCalledWith(printHello, system.console)
```

##### Spying on the implementation of the interface.

We can achive full coverage, if the language allows us to spy on a function call.

```
// TEST 1.3
when ("printHello() is called")
  it ("should print 'Hello World!' to the console")
    printHello()

    assertCalledWith(system.console, "Hello World!")
```

This test covers the full scope of the SUT. The SUT is responsible of forming the string "Hello World!" and printing it to the console.

But this solution couples the test to the implementation. If the implementation changes, the test will break. Of all the proposed solutions, this is the least desirable even if its the most succinct.

##### Achieving full coverage with the EV language

The previous approaches prove that the problem can be solved with existing languages, but the solution is far from obvious. The given requirements are as simple as they can be, but testing their behaviour requires effort. This particular combination, trivial implementation and complex testing, detracts engineers from adopting vital testing practices.

The EV language should allow the engineer to easily write the test in a way that is not coupled to the implementation, but still achieves full coverage.

```
// TEST 1.4.1
when ("a sayHello event is published")
  publish("sayHello")

  it ("should emit 'Hello World!' to system_console")
      assertEmitted("Hello World!", system_console)


// TEST 1.4.2
when ("the application starts")
  publish("application_start")

  it ("should register Console as listener of system_console")
      assertRegistered(Console, system_console)
```

By decoupling modules through an event broker, the EV language allows the engineer to write tests that are not coupled to the implementation. The SUT is responsible of forming the string "Hello World!" and publishing it. The Console module is responsible of listening to the event broker and printing the string to the console.

Ideally the language should provide a robust set of compilation tools that makes TEST 1.4.2 redundant. That is, if a module is emiting an event, the language should warn the engineer if no other module is listening to it.

### SOLID principles

The language should enforce
