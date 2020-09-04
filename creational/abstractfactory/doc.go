package abstractfactory

// =====================================================================================================================
// Abstract Factory Design Pattern
// =====================================================================================================================
// Abstract Factory Design Pattern is a creational design pattern.
// It lets you create a family of related objects
// It lets you produce families of related objects without specifying their concrete classes
// It is an abstraction over the factory pattern
// Abstract Factory patterns work around a super-factory which creates other factories
// An abstract factory pattern is also called the Factory of Factories or Kit
// The abstract factory pattern has an interface that is responsible for creating a
// factory of related objects without explicitly specifying their classes
// Each generated factory can give the objects as per the factory method pattern
// The first thing the Abstract Factory pattern suggests is to explicitly
// declare interfaces for each distinct product of the product family
//
// Example:
// ===========================================
// Let’s say we have two factories nike and adidas
// Imagine you need to buy a sports kit which has a shoe and short
// Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either nike or adidas
// This is where the abstract factory comes into the picture, as concrete products
// that you want is shoe and a short and these products will be created by the abstract factory of nike and adidas
// Both these two factories  nike and adidas implement iSportsFactory interface
// We have two product interfaces:
// 1. iShoe: This interface is implemented by nikeShoe and adidasShoe concrete product
// 2. iShort: This interface is implemented by nikeShort and adidasShort concrete product
//
// How abstract factory is different from builder design pattern?
// ===========================================
// Abstract factory and Builder, both help in creating objects and are as such part of the creational design patterns
// They vary in the context of their usage. Abstract factory is used for creating a family of objects,
// which share a common interface
// Builder, on the other hand, is concerned with building a single type of object
// Builders may also be abstract and its subclasses may have common logic to build a certain subpart
//
// When to use
// ===========================================
// Use the Abstract Factory when your code needs to work with various families of related products,
// but you don’t want it to depend on the concrete classes of those products
// they might be unknown beforehand or you simply want to allow for future extensibility
//
// Pros
// ===========================================
// 1. You can be sure that the products you’re getting from a factory are compatible with each other.
// 2. You avoid tight coupling between concrete products and client code.
// 3. Single Responsibility Principle:
//    You can extract the product creation code into one place, making the code easier to support.
// 4. Open/Closed Principle:
//    You can introduce new variants of products without breaking existing client code
//
// Cons
// ===========================================
// 1. The code may become more complicated than it should be, since
//    a lot of new interfaces and classes are introduced along with the pattern.
//
// Relations with Other Patterns
// ===========================================
// Many designs start by using Factory Method (less complicated and more customizable via subclasses) and
// evolve toward Abstract Factory, Prototype, or Builder (more flexible, but more complicated)
//
// Builder focuses on constructing complex objects step by step
// Abstract Factory specializes in creating families of related objects
// Abstract Factory returns the product immediately, whereas Builder
// lets you run some additional construction steps before fetching the product
//
// Abstract Factory classes are often based on a set of Factory Methods,
// but you can also use Prototype to compose the methods on these classes
//
// Abstract Factory can serve as an alternative to Facade when you only want to
// hide the way the subsystem objects are created from the client code
//
// You can use Abstract Factory along with Bridge
// This pairing is useful when some abstractions defined by Bridge can only work with specific implementations
// In this case, Abstract Factory can encapsulate these relations and hide the complexity from the client code
//
// Abstract Factories, Builders and Prototypes can all be implemented as Singletons
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// adidas, _ := GetSportsFactory("adidas")
// aShoe := adidas.GetShoe()
// aShoe.SetSize(8)
// aShoe.SetColor("White")
// aShort := adidas.GetShort()
// aShort.SetSize(30)
// aShort.SetColor("White")
// fmt.Printf("Adidas Sports Kit:\nShoe: Color:%s Size:%d\n"+
// 	"Short: Color:%s Size:%d\n", aShoe.Color(), aShoe.Size(), aShort.Color(), aShort.Size())
//
// fmt.Println()
//
// nike, _ := GetSportsFactory("nike")
// nShoe := nike.GetShoe()
// nShoe.SetSize(10)
// nShoe.SetColor("Black")
// nShort := nike.GetShort()
// nShort.SetSize(32)
// nShort.SetColor("Black")
// fmt.Printf("Nike Sports Kit:\nShoe: Color:%s Size:%d\n"+
// 	"Short: Color:%s Size:%d\n", nShoe.Color(), nShoe.Size(), nShort.Color(), nShort.Size())
//}
//
// Key Points
// ===========================================
// Family of related objects, which share a common interface
// Create objects without specifying concrete classes of products
// Super factory or Factory of factories or Kit
// Factory method that creates other factories
// Ensures compatibility between created objects(Furniture example)
