package factory

// =====================================================================================================================
// Factory Design Pattern
// =====================================================================================================================
// Factory design pattern is a creational design pattern
// It is also one of the most commonly used patterns
// This pattern provides a way to hide the creation logic of the instances being created
// Factory Method provides a common interface for creating objects,
// but allows subclasses to alter the type of objects that will be created
// The client only interacts with a factory struct(or method) and tells the kind of instances that needs to be created
// The factory method interacts with the corresponding
// concrete structs(or classes) and returns the correct instance back.
// This pattern is also known as the Virtual Constructor
// Factory method pattern enables us to create an object without exposing the creation
// logic to the client and refer to the newly-created object using a common interface.
// The intent of this pattern, according to Design Patterns by Gamma et al, is to:
//  Define an interface for creating an object, but let subclasses decide which class to instantiate.
//  The Factory method allows a class to defer instantiation to subclasses
//
// Example 1:
// ===========================================
// Imagine that you’re creating a logistics management application
// The first version of your app can only handle transportation by trucks,
// so the bulk of your code lives inside the Truck class
// After a while, your app becomes pretty popular
// Each day you receive dozens of requests from sea transportation companies to incorporate sea logistics into the app.
// Great news, right? But how about the code? At present,
// most of your code is coupled to the Truck class
// Adding Ships into the app would require making changes to the entire codebase
// Moreover, if later you decide to add another type of transportation to the app,
// you will probably need to make all of these changes again.
// As a result, you will end up with pretty nasty code, riddled with conditionals that switch the app’s behavior
// depending on the class of transportation objects.
//
// Solution:
// ===========================================
// The Factory Method pattern suggests that you replace direct object construction calls (using the new operator)
// with calls to a special factory method. Don’t worry: the objects are still created via the new operator,
// but it’s being called from within the factory method
// Objects returned by a factory method are often referred to as “products.”
// At first glance, this change may look pointless: we just moved the constructor call
// from one part of the program to another. However, consider this: now you can override
// the factory method in a subclass and change the class of products being created by the method.
// There’s a slight limitation though: subclasses may return different types of products only if these products
// have a common base class or interface. Also, the factory method in the base class
// should have its return type declared as this interface.
// For example, both Truck and Ship classes should implement the iTransport interface,
// which declares SetType(string) and Type() string. Each class implements this method differently:
// trucks deliver cargo by land, ships deliver cargo by sea
// The factory method in the RoadLogistics class returns truck objects,
// whereas the factory method in the SeaLogistics class returns ships
// The code that uses the factory method (often called the client code)
// doesn’t see a difference between the actual products returned by various subclasses
// The client treats all the products as abstract Transport
// The client knows that all transport objects are supposed to have the deliver method,
// but exactly how it works isn’t important to the client.
//
// Example 2:
// ===========================================
// We have iGun interface which defines all methods a gun should have
// There is gun struct that implements the iGun interface.
// Two concrete guns ak47 and scarl.
// Both embed gun struct and hence also indirectly implement all methods of iGun and hence are of iGun type
// We have a gunFactory struct which creates the gun of type ak47 or scarl
// The main.go acts as a client and instead of directly interacting with ak47 or scarl,
// it relies on gunFactory to create instances of ak47 and scarl
//
// When to use
// ===========================================
// 1. Use the Factory Method when you don’t know beforehand the exact types and dependencies
//    of the objects your code should work with. The Factory Method separates product construction code from the code
//    that actually uses the product. Therefore it’s easier to extend the product construction code
//    independently from the rest of the code
//
// Pros
// ===========================================
// 1. You avoid tight coupling between the creator and the concrete products.
// 2. Single Responsibility Principle:
//    You can move the product creation code into one place in the program, making the code easier to support.
// 3. Open/Closed Principle:
//    You can introduce new types of products into the program without breaking existing client code.
//
// Cons
// ===========================================
// 1. The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern.
//    The best case scenario is when you’re introducing the pattern into an existing hierarchy of creator classes
//
// Difference between a builder and factory design pattern
// ===========================================
// Builder pattern is an object creational software design pattern
// This pattern is often compared with “Factory” method pattern because factory method is also an object creational DP
// The key difference is how the object is being created though
// For example with the factory DP, createCar might return Honda or Audi
// To add attributes, you will need to make additional calls such as setColor, setEngine etc on the object returned
// Now with the Builder pattern, you could be specific about the object all in one statement,
// such as 4 all-wheel drive red Honda with v6 engine etc
// Basically, factory method pattern would create an object via createCar VS Builder
// pattern has the ability to create an object by adding granularity to the object by
// applying chains of function calls to the object being created.
// In summary,
// 1. A Factory Design Pattern is used when the entire object can be easily created and object is not very complex
//    Whereas Builder Pattern is used when the construction process of a complete object is very complex.
// 2. Another difference is if we want to add another object into Factory,
//    we need to modify the factory method- violates open close principle (close for modification and open for extension)
//    Whereas, if want to add another car builder we can just create another builder and
//    no change in director class as director constructor accept all builders, extensible.
//
// Relations with Other Patterns
// ===========================================
// Many designs start by using Factory Method (less complicated and more customizable via subclasses) and
// evolve toward Abstract Factory, Prototype, or Builder (more flexible, but more complicated).
//
// Abstract Factory classes are often based on a set of Factory Methods,
// but you can also use Prototype to compose the methods on these classes.
//
// You can use Factory Method along with Iterator to let collection subclasses
// return different types of iterators that are compatible with the collections.
//
// Prototype isn’t based on inheritance, so it doesn’t have its drawbacks
// On the other hand, Prototype requires a complicated initialization of the cloned object
// Factory Method is based on inheritance but doesn’t require an initialization step.
//
// Factory Method is a specialization of Template Method
// At the same time, a Factory Method may serve as a step in a large Template Method
//
// RUN INSTRUCTIONS Example 1:
// ===========================================
// item := "gold"
//
// ship, _ := TransportFactory("ship")
// ship.SetName("Ship 123")
// Deliver(item, ship)
//
// truck, _ := TransportFactory("truck")
// truck.SetName("Truck 123")
// Deliver(item, truck)
//
//}
//
//
// RUN INSTRUCTIONS Example 2:
// ===========================================
// func main() {
// akm , _:= NewGun("akm")
// Shoot(akm)
//
// scarl , _:= NewGun("scarl")
// Shoot(scarl)
//
//}
//
// Key Points
// ===========================================
// Components: Factory method
// Hides creation logic
// Factory method delegates work to concrete classes
// Subclasses override the method as necessary
// Client dont see the difference in the objects returned
// (ship and truck, since both both implement iTransport interface)
