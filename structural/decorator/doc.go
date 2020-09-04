package decorator

// =====================================================================================================================
// Decorator Design Pattern
// =====================================================================================================================
// Decorator is a structural design pattern
// It lets you attach new behaviors to objects by placing these objects
// inside special wrapper objects that contain the behaviors
// The Decorator pattern adds new functionality to an existing object without altering its structure
// It is a structural pattern as this pattern acts as a wrapper to existing class
//
//
// Example 1:
// ===========================================
// Imagine that you’re working on a notification library which
// lets other programs notify their users about important events
// The initial version of the library was based on the Notifier class that had only a few fields,
// a constructor and a single send method. The method could accept a message argument from a client and
// send email notification to the user.
// A third-party app which acted as a client was supposed to create and configure the notifier object once,
// and then use it each time something important happened.
// At some point, you realize that users of the library expect more than just email notifications.
// Many of them would like to receive an SMS about critical issues.
// Others would like to be notified on Facebook and, of course,
// the corporate users would love to get Slack notifications
// How hard can that be?
// You extended the Notifier class and put the additional notification methods into new subclasses.
// Now the client was supposed to instantiate the desired notification class and use it for all further notifications
// But then someone reasonably asked you,
// “Why can’t you use several notification types at once?
// If your house is on fire, you’d probably want to be informed through every channel.”
// You tried to address that problem by creating special subclasses which
// combined several notification methods within one class.
// However, it quickly became apparent that this approach would bloat the code immensely,
// not only the library code but the client code as well.
//
// Solution:
// ===========================================
// Extending a class is the first thing that comes to mind when you need to alter an object’s behavior.
// However, inheritance has several serious caveats that you need to be aware of.
// 1. Inheritance is static. You can’t alter the behavior of an existing object at runtime.
//    You can only replace the whole object with another one that’s created from a different subclass.
// 2. Subclasses can have just one parent class.
//    In most languages, inheritance doesn’t let a class inherit behaviors of multiple classes at the same time.
// One of the ways to overcome these caveats is by using Aggregation or Composition  instead of Inheritance.
// Aggregation: Object A contains object B, B can live without A
// Composition: Object A consists of objects B, A manages lifecycle of B, B cant live without A
// Both of the alternatives work almost the same way:
// one object has a reference to another and delegates it some work,
// whereas with inheritance, the object itself is able to do that work, inheriting the behavior from its superclass.
// With this new approach you can easily substitute the linked “helper” object with another,
// changing the behavior of the container at runtime.
// An object can use the behavior of various classes,
// having references to multiple objects and delegating them all kinds of work.
// Aggregation/composition is the key principle behind many design patterns, including Decorator.
// Wrapper is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern.
// A “wrapper” is an object that can be linked with some “target” object.
// The wrapper contains the same set of methods as the target and delegates to it all requests it receives.
// However, the wrapper may alter the result by doing something either before or after it passes the request to the target
// When does a simple wrapper become the real decorator?
// As I mentioned, the wrapper implements the same interface as the wrapped object.
// That’s why from the client’s perspective these objects are identical.
// Make the wrapper’s reference field accept any object that follows that interface.
// This will let you cover an object in multiple wrappers, adding the combined behavior of all the wrappers to it.
//
//
// When to use
// ===========================================
// 1. Use the Decorator pattern when you need to be able to assign extra behaviors to objects at
//    runtime without breaking the code that uses these objects
// 2. Use the pattern when it’s awkward or not possible to extend an object’s behavior using inheritance
//
// Pros
// ===========================================
// 1. You can extend an object’s behavior without making a new subclass.
// 2. You can add or remove responsibilities from an object at runtime.
// 3. You can combine several behaviors by wrapping an object into multiple decorators.
// 4. Single Responsibility Principle:
//    You can divide a monolithic class that implements many possible variants of behavior into several smaller classes
//
// Cons
// ===========================================
// 1. It’s hard to remove a specific wrapper from the wrappers stack.
// 2. It’s hard to implement a decorator in such a way
//    that its behavior doesn’t depend on the order in the decorators stack.
// 3. The initial configuration code of layers might look pretty ugly
//
// Relations with Other Patterns
// ===========================================
// Adapter changes the interface of an existing object,
// while Decorator enhances an object without changing its interface.
// In addition, Decorator supports recursive composition, which isn’t possible when you use Adapter.
//
// Adapter provides a different interface to the wrapped object,
// Proxy provides it with the same interface,
// and Decorator provides it with an enhanced interface.
//
// Chain of Responsibility and Decorator have very similar class structures.
// Both patterns rely on recursive composition to pass the execution through a series of objects.
// However, there are several crucial differences.
// The CoR handlers can execute arbitrary operations independently of each other.
// They can also stop passing the request further at any point.
// On the other hand, various Decorators can extend the object’s behavior while
// keeping it consistent with the base interface.
// In addition, decorators aren’t allowed to break the flow of the request.
//
// Composite and Decorator have similar structure diagrams since both rely on recursive composition
// to organize an open-ended number of objects.
// A Decorator is like a Composite but only has one child component.
// There’s another significant difference: Decorator adds additional responsibilities to the wrapped object,
// while Composite just “sums up” its children’s results.
// However, the patterns can also cooperate:
// you can use Decorator to extend the behavior of a specific object in the Composite tree.
//
// Designs that make heavy use of Composite and Decorator can often benefit from using Prototype.
// Applying the pattern lets you clone complex structures instead of re-constructing them from scratch.
//
// Decorator lets you change the skin of an object, while Strategy lets you change the guts.
//
// Decorator and Proxy have similar structures, but very different intents.
// Both patterns are built on the composition principle,
// where one object is supposed to delegate some of the work to another.
// The difference is that a Proxy usually manages the life cycle of its service object on its own,
// whereas the composition of Decorators is always controlled by the client.
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// //version 1 of the app
// notify1 := &Notify{}
// notify1.SendNotification("Sample Message Version 1")
//
// // version 2 of the app
// notify2 := &Notify{}
// sms := &SmsDecorator{}
// fb := &FBDecorator{}
// notifiers := make([]Notifier, 0, 1)
// notifiers = append(notifiers, sms, fb)
// nd := NotifyDecorator{Notify: notify2, Notifiers: notifiers}
// nd.SendNotification("Sample Message Version 2")
// }
//
// Key Points
// ===========================================
// Components: Component, Decorator, ConcreteComponent, ConcreteDecorator
// Uses composition
// Adds features to an existing object
// Decorators should be consistent with the existing interfaces
// Decorators can t break the flow of request
// Composition of Decorators is always controlled by the client.
