https://www.youtube.com/watch?v=6zuJXIbOyhs&t=2676s&ab_channel=ProgrammingPercy

## How To Implement Domain-Driven Design (DDD) in Go
**Domain modeling session:** It's a session where the subject matter expert(SME) explains to the engineers about the domain.

**ubiquitous language**: It's about talking the same language between the SMEs and engineers. A unified language which everyone on the
team understand.

### Entity
Entities are uniquely identifiable structures and they're mutable. For example a person is uniquely identifiable by their social security number.
We can also modify them, so they can change.

Entities are structures that must have a unique identifier and their immutable.

### Value objects
Value objects are non-identifiable(don't have unique identifiers) and immutable(we can't change them).

For example a transaction is a value object, once it's done, we can't change it. In real world, a transaction would have an id connected to it
but it's not mutable and that's the kicker here(makes it a value object, although it's identifiable). 

Start with entities and value objects.

Create domain folder which will be storing all the subdomains and also entity folder.

Keep entities separate from domains because that allows us to reuse the entities in different domains.

We made Transaction fields unexported so no other domain can reach and change them.

### Aggregates
Sometimes we can't explain a real life thing with entities and value objects, we need to combine them. That's what aggregate does.

So an aggregate is a combination of entities and value objects.

For example a customer has a root identity and the root identity is used to identify aggregates. This is why an entity has to have a unique identifier.
So whenever we create a customer, we can say a customer is a person and they are identifiable by the personID(id of person). So customer 
is not a root entity. Another reason is when we have products, we don't connect the productID to the customerID because that doesn't make sense, so 
it's(customer) not a root entity, it's a sub entity.

The business logic for for example customers should be inside the aggregate, not inside the entity. Entities are dumb, they're just placeholders
for information. Aggregate can contain business logic.

We made all the fields of customer aggregate to unexported because aggregates should not be accessible directly to grab the data.
Also we're not using any json tags or anything like that in aggregates, because it's not up to the aggregate to decide how the data is
supposed to be formatted.

The data of an aggregate is not accessible from outside of the aggregate. Nothing outside of the aggregate can modify the data of aggregate,
the changes are done by exposing functions like(`SetName`) that allows changes. So if we should be able to modify the name, we expose a function
which allows you to do that(like SetName func). You don't directly modify an aggregate.

Also we set all the entity fields of aggregate to pointer because we can change their state and therefore it will be reflected across the places
it's been used whenever sth in that pointer changes. So if we have a `Person` in multiple places, that change should reflect everywhere.
Note we didn't make transaction field (value object fields) a pointer in the aggregate because it can not be changed, so there's no reason to make
it as a pointer.

### Factories
Factory pattern is a design pattern that is used to encapsulate complex logic inside function for creating the instances, without caller
knowing anything about the actual implementation details.

DDD suggests factories for creating complex aggregates, repositories and services.

Note: Custom errors(sentinel errors) make testing easier.

When we have business logic like NewCustomer() func, we should always test it.

By using `%v` verb, even if the value that will be substituted with %v, is nil, we won't get a nil pointer exception.

### Repository pattern
Aggregates(like customer) don't have any json, bson, csv ... tags and that's because aggregates are stored by repositories.
So an aggregate is a combination of entities and value objects, but when we store and manage them, we use a repository.

The repository pattern relies on hiding the implementation details behind an interface and with this pattern, we can have for example
whenever we do unit tests, we can have in-memory repository for storing customers but we can also have mysql repository. So if the db changes,
we can build a new repository for mongodb and satisfy the same interface that the previous db was satisfying and then we can swap them.
So we just refactor the repository and it will propagate to all other domains.

Create customer folder that holds our repository for customers, in domain folder.

It's not good to have the word `repository` in `CustomerRepository` interface, but it works.

By defining CustomerRepository, it doesn't matter if the customer is fetched from mongodb, mysql, in memory, ... , as long as we call the `Get` method,
it's fine.

Now create `memory` folder in `domain/customer` which will hold our in-memory solution for the `CustomerRepository` which we can use in unit test.

**A repository is used to manage aggregates.** One repository, only handles one aggregate. We want loose coupling. But in real world,
we can't rely on one repository, so we have to start coupling somewhere. For example if we have an order, we need to get the customer, make a billing
and ... (coupling) and that brings us to the next point of DDD and it's services.

Note: If we have an aggregate called product, we need a repo to manage them.

### Services & configuration patterns & more repositories
A service will tie together all the loosely coupled repositories into a business logic that fulfills the needs of the domain.
So in our tavern, we might need an order service. An order service is responsible for making the repositories work together to perform an order.
So getting the customer using the customer repo, getting the product with product repo(for instance) and then making the billing using the billing
service.

So the service takes these loosely coupled pieces and couples them together.

Create services folder and order.go .

Whenever you have services, the factories tend to get larger and more complex. Because you accept multiple repositories as inputs.

One trick for fixing this problem, is sorta a service configuration generator pattern and it allows us to create flexible, modular services
where you can replace the repositories easily.

The `OrderService` needs `CustomerRepository` because whenever somebody makes an order, they are a customer, so we need to handle the
customer aggregate, so we need the `CustomerRepository` in the service.

The reason we would use this pattern is for example if in the future change from in-memory to MySQL and we know the order service itself is used
in a lot of places and we don't want to change everything because of this, you would simply build repository for mysql or mongodb and the
you replace the `WithMemoryCustomerRepository` with `WithMySQLCustomerRepository` or ... and everything would continue working as long as your
repository is working as expected.

Another example is if you have a mail service but you don't want to send mails when in unit tests. For this, you replace the repository for the emails.

Let's add some functions for business logic in order service.

The helper functions for exposing data of aggregate, depend on what you actually need to expose.

Note: A service can hold multiple repositories. But a service is also allowed to hold other services. So the order service could hold
the billing service for example. So a service holds repositories and potentially sub-services, that builds together the business flow.

### Tavern service & sub-services
Create a tavern and the tavern is a service.

Tavern service will hold sub services.

One common approach in repositories(like mongo repo) is to have some functions for converting the related aggregate type into it's repo type(customer
aggregate type into customer mongo repo type). Why? Because we're not working with aggregate types directly in the repo, we need to convert
the aggregate type to it's respected repo type. So we have a function like `NewFromCustomer`. So now we can easily go from an aggregate type
to the related internal repo struct.

We also have the other way around conversions, like: `ToAggregate`.

These are a bit overhead.

So we can switch repositories in a service, like replacing the postgres repo with mongo repo. So the service won't have to care about the repositories.
We would get the same result no matter which repo we use.

To structure a domain driven project, watch the next video.

## How To Structure Domain Driven Design (DDD) In Go
https://www.youtube.com/watch?v=jJHhXaWwM7Y&t=154s&ab_channel=ProgrammingPercy

### intro
We will be moving the aggregates into their own domain packages.

So we have learned what an aggregate is and what rules apply to them, so we don't necessarily have to name the package `aggregate`.
So place aggregates into their respective `domain` package. So place customer files in aggregate package in domain>customer folder.
Do the same thing for `product` aggregate.

Now remove the aggregate folder.

### Start extracting aggregate
We still have the entity and valueobject packages. It's not wrong having them in a separate packages, so that we don't cyclic imports.
But the entities and value objects are related to the tavern domain and the tavern domain is actually the package that we're building(the root package).

Note: We could have a tavern domain inside the domain package(folder).

Now delete the entity and valueobject folders. Rename the moved files package to `tavern` which is the core domain(we have a `tavern` package at root
level). Our root package is named tavern. So we need to change the module name in go.mod to `tavern`.

### splitting services
Rename `CustomerRepository` and `ProductRepository` to `Repository`. Because it's name starts with the package name, so it's unnecessary.
This is only makes sense if we put the repository in the domain.

### making services expose needed functionality

### removing duplicate naming convention

### making the tavern runnable

### testing with a mongodb

### ending