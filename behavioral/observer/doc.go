package observer

// =====================================================================================================================
// Observer Design Pattern
// =====================================================================================================================
// Observer Design Pattern is a behavioral design pattern
// This pattern allows an instance (called subject or observable) to
// publish events to other multiple instances (called observers)
// These observers subscribe to the subject and
// hence get notified by events in case of any change happening in the subject
// It lets you define a subscription mechanism to
// notify multiple objects about any events that happen to the object they’re observing.
// If the subject wants to unsubscribe, he can deregister from notifications
// The major components of the observer pattern are:
// a. Subject - It is the instance which publishes an event when anything changes
// b. Observer - It subscribes to the subject and gets notified by the events
// Generally, Subject and Observer are implemented as an interface.
// Concrete implementation of both are used
//
// Example 1
// ===========================================
// In the E-Commerce website, many items go out of stock.
// There can be customers, who are interested in a particular item that went out of stock.
// There are three solutions to this problem:
// 1. The customer keeps checking the availability of the item at some frequency.
// 2. E-Commerce bombard customers with all new items available which are in stock
// 3. The customer subscribes only to the particular item he is interested in and
//    gets notified in the case that item is available. Also, multiple customers can subscribe to the same product
// Option 3 is most viable, and this is what Observer Patter is all about.
// The object that has some interesting state is often called subject,
// but since it’s also going to notify other objects about the changes to its state, we can call it publisher.
// All other objects that want to track changes to the publisher’s state are called subscribers.
//
// When to use
// ===========================================
// 1. Use the Observer pattern when changes to the state of one object may require changing other objects,
//    and the actual set of objects is unknown beforehand or changes dynamically.
// 2. Use the pattern when some objects in your app must observe others,
//    but only for a limited time or in specific cases.
//
// Pros
// ===========================================
// 1. Open/Closed Principle:
//    You can introduce new subscriber classes without having to change the publisher’s code
//    (and vice versa if there’s a publisher interface).
// 2. You can establish relations between objects at runtime.
//
// Cons
// ===========================================
// 1. Subscribers are notified in random order.
//
// Relations with Other Patterns
// ===========================================
// Chain of Responsibility, Command, Mediator and Observer address various ways of connecting senders and
// receivers of requests:
// a. Chain of Responsibility passes a request sequentially along a dynamic chain of potential
//    receivers until one of them handles it.
// b. Command establishes unidirectional connections between senders and receivers.
// c. Mediator eliminates direct connections between senders and receivers, forcing them to
//    communicate indirectly via a mediator object.
// d. Observer lets receivers dynamically subscribe to and unsubscribe from receiving requests.
//
// The difference between Mediator and Observer is often elusive.
// In most cases, you can implement either of these patterns; but sometimes you can apply both simultaneously.
// Let’s see how we can do that:
//
// The primary goal of Mediator is to eliminate mutual dependencies among a set of system components.
// Instead, these components become dependent on a single mediator object.
// The goal of Observer is to establish dynamic one-way connections between objects,
// where some objects act as subordinates of others.
//
// There’s a popular implementation of the Mediator pattern that relies on Observer.
// The mediator object plays the role of publisher,
// and the components act as subscribers which subscribe to and unsubscribe from the mediator’s events.
// When Mediator is implemented this way, it may look very similar to Observer.
//
// When you’re confused, remember that you can implement the Mediator pattern in other ways.
// For example, you can permanently link all the components to the same mediator object.
// This implementation won’t resemble Observer but will still be an instance of the Mediator pattern.
//
// Now imagine a program where all components have become publishers,
// allowing dynamic connections between each other.
// There won’t be a centralized mediator object, only a distributed set of observers.
//
//
// RUN INSTRUCTIONS:
// ===========================================
// customer1 := &e1.Customer{}
// customer1.SetName("Ash")
// customer2 := &e1.Customer{}
// customer2.SetName("John")
// customer3 := &e1.Customer{}
// customer3.SetName("Poor Guy")
//
// item := e1.Item{}
// item.SetName("Galaxy S9")
// item.Subscribe(customer1)
// item.Subscribe(customer2)
// item.Subscribe(customer3)
//
// item.UnSubscribe(customer3)
//
// item.UpdateStock(10)
//}
//
// Key Points
// ===========================================
// Components: Observer(Subscriber), Subject(Observable/Publisher), Concrete Implementations of Observer and Subject
// Subscribe/UnSubscribe from events
// Very similar to mediator pattern
