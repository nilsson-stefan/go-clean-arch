# go-clean-arch
Attempt to create a Go-project template that **that somewhat fulfills** the structure and rules 
recommended by Uncle Bobs clean architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

<img src="https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg" alt="drawing" width="400"/>

\
Short description of the circles and where their code are available in this project:
## 1. Blue circle 
Code that uses frameworks and drivers. In this example:
* **/persistence/...** using dynamodb driver
* **/web/...** gorilla mux request router

Could also be event listeners or other external input.

## 2. Green circle
Interface adapters that defines the "contract" between the use cases (business logic) and the framework specific code in the outer circle. 
In this example:
* **/adapters/controllers/...** specifies how outer circle (rest api/events) communicates with the use cases (business logic)
* **/adapters/repository/...** specifies an interface towards the persistence layer enabling the inner (business logic) to read/write to our database in the outer layer.

## 3. Red circle
This is where the business logic is implemented using the core entities found in the inner/yellow circle. This code should be
totally unaware of external frameworks and web/db specific models. This is accomplished by following the dependency rule (see below).
And by doing this you get *"a system that is intrinsically testable, with all the benefits that implies. And 
when any of the external parts of the system become obsolete, like the database, or the web framework, 
you can replace those obsolete elements with a minimum of fuss."*

## 4. Yellow circle
Core entities used by the business logic.

## The dependency rule
Source code dependencies must only point inwards

## Configuration
The main function in main.go is using the configuration and initialization code in 
* **/configuration/...**

... to instantiate and wire code according to specific config/setup. This can be done manually or by using a framework. But 
since this code needs to be able to create types from all circles the configuration code (at least in this example) is an 
exception from the dependency rule.

## All the data
While the source code dependencies only should point inwards, the data flows in both directions \
In: \
user --> rest request (web) --> adapter (interface) --> core use case --> \
Out: \
--> adapter (interface) --> db query (persistence)  

