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