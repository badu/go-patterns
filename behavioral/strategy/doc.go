package strategy

// =====================================================================================================================
// Strategy Design Pattern
// =====================================================================================================================
// Strategy design pattern is a behavioral design pattern
// This design pattern allows you to change the behavior of an object at run time
// without any change in the class of that object
// It lets you define a family of algorithms, put each of them into a separate class,
// and make their objects interchangeable
// The Strategy pattern suggests that you take a class that does something specific in a lot of different ways and
// extract all of these algorithms into separate classes called strategies
// The original class, called context, must have a field for storing a reference to one of the strategies
// The context delegates the work to a linked strategy object instead of executing it on its own
// The context isn’t responsible for selecting an appropriate algorithm for the job
// Instead, the client passes the desired strategy to the context
// In fact, the context doesn’t know much about strategies
// It works with all strategies through the same generic interface,
// which only exposes a single method for triggering the algorithm encapsulated within the selected strategy
// This way the context becomes independent of concrete strategies,
// so you can add new algorithms or modify existing ones without changing the code of the context or other strategies
//
// Example 1:
// ===========================================
// One day you decided to create a navigation app for casual travelers
// The app was centered around a beautiful map which helped users quickly orient themselves in any city.
// One of the most requested features for the app was automatic route planning
// A user should be able to enter an address and see the fastest route to that destination displayed on the map
// The first version of the app could only build the routes over roads
// People who traveled by car were bursting with joy
// But apparently, not everybody likes to drive on their vacation
// So with the next update, you added an option to build walking routes
// Right after that, you added another option to let people use public transport in their routes
// However, that was only the beginning. Later you planned to add route building for cyclists
// And even later, another option for building routes through all of a city’s tourist attractions
// While from a business perspective the app was a success,
// the technical part caused you many headaches
// Each time you added a new routing algorithm, the main class of the navigator doubled in size
// At some point, the beast became too hard to maintain
// Any change to one of the algorithms,
// whether it was a simple bug fix or a slight adjustment of the street score,
// affected the whole class, increasing the chance of creating an error in already-working code
// In addition, teamwork became inefficient
// Your teammates, who had been hired right after the successful release,
// complain that they spend too much time resolving merge conflicts
// Implementing a new feature requires you to change the same huge class,
// conflicting with the code produced by other people.
//
// Solution:
// ===========================================
// In our navigation app, each routing algorithm can be extracted to its own class with a single navigate method
// The method accepts an origin and destination and shows a route from start to end point
// Even though given the same arguments, each routing class might build a different route,
// the main navigator class doesn’t really care which algorithm is selected since its primary job is
// to render a set of checkpoints on the map
// The class has a method for switching the active routing strategy, so its clients,
// such as the buttons in the user interface, can replace the currently selected routing behavior with another one
// Reference: https://refactoring.guru/design-patterns/strategy
//
// When to use:
// ===========================================
// 1. Use the Strategy pattern when you want to use different variants of an algorithm within
//    an object and be able to switch from one algorithm to another during runtime
// 2. Use the pattern when your class has a massive conditional operator that
//    switches between different variants of the same algorithm
// 3. Use the Strategy when you have a lot of similar classes that only differ in the way they execute some behavior
// 4. Use the pattern to isolate the business logic of a class from the implementation details
//    of algorithms that may not be as important in the context of that logic
//
// Pros
// ===========================================
// 1. You can swap algorithms used inside an object at runtime.
// 2. You can isolate the implementation details of an algorithm from the code that uses it.
// 3. You can replace inheritance with composition.
// 4. Open/Closed Principle:
//    You can introduce new strategies without having to change the context
//
// Cons
// ===========================================
// 1. If you only have a couple of algorithms and they rarely change, there’s no real reason to
//    over-complicate the program with new classes and interfaces that come along with the pattern
// 2. Clients must be aware of the differences between strategies to be able to select a proper one
// 3. A lot of modern programming languages have functional type support that lets
//    you implement different versions of an algorithm inside a set of anonymous functions.
//    Then you could use these functions exactly as you’d have used the strategy objects,
//    but without bloating your code with extra classes and interfaces
//
// Relations with Other Patterns
// ===========================================
// Bridge, State, Strategy (and to some degree Adapter) have very similar structures
// Indeed, all of these patterns are based on composition, which is delegating work to other objects
// However, they all solve different problems
// A pattern isn’t just a recipe for structuring your code in a specific way
// It can also communicate to other developers the problem the pattern solves.
//
// Command and Strategy may look similar because you can use both to parameterize
// an object with some action. However, they have very different intents
// You can use Command to convert any operation into an object
// The operation’s parameters become fields of that object
// The conversion lets you defer execution of the operation, queue it,
// store the history of commands, send commands to remote services, etc
// On the other hand, Strategy usually describes different ways of doing the same thing,
// letting you swap these algorithms within a single context class
//
// Decorator lets you change the skin of an object, while Strategy lets you change the guts
//
// Template Method is based on inheritance:
// it lets you alter parts of an algorithm by extending those parts in subclasses.
// Strategy is based on composition:
// you can alter parts of the object’s behavior by supplying it with different
// strategies that correspond to that behavior.
// Template Method works at the class level, so it’s static.
// Strategy works on the object level, letting you switch behaviors at runtime
//
// State can be considered as an extension of Strategy.
// Both patterns are based on composition:
// they change the behavior of the context by delegating some work to helper objects
// Strategy makes these objects completely independent and unaware of each other
// However, State doesn’t restrict dependencies between concrete states,
// letting them alter the state of the context at will
//
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// route := Navigator{}
//
// walk := &WalkRoute{}
// route.SetStrategy(walk)
// route.Navigate("A", "B")
//
// bus := &BusRoute{}
// route.SetStrategy(bus)
// route.Navigate("A", "B")
//
// car := &CarRoute{}
// route.SetStrategy(car)
// route.Navigate("A", "B")
// }
//
// Key Points
// ===========================================
// Components: context, strategy, ConcreteStrategies
// Client sets strategy
// Strategies do not communicate with each other
// Multiple ways(algorithms) to do the same task
// Uses composition
// New algorithm = new class
// All algorithms implement the same interface
// Context class is the original class
// Context class doesn't know about strategies
