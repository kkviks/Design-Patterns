package solid

/* Interface segregation principle
You shouldn't put too much into a interface
Break into several interfaces
*/

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// ok if you need a multifunction device
type MultiFunctionPrinter struct {
	// ...
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
	// ...
}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct {
	// ...
}

func (m MyPrinter) Print(d Document) {
	// ...
}

// combine interfaces
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	//
}

func (p Photocopier) Print(d Document) {
	//
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main_isp() {

}

/*

Transcript
The interface segregation principle is a really simple principle, it's probably the simplest principle

out of the solid design principles.

Basically, what it states is that you shouldn't put too much into an interface.

You shouldn't try to throw everything and the kitchen sink into just one single interface.

And sometimes it makes sense to break up the interface into several smaller interfaces.

So let me show you a very simple, somewhat contrived example.

Let's suppose that you have some sort of document type.

So just some information about the document and you want to make an interface that allows people to

build the different machines, different constructs for operating on a document.

So doing things like printing the document or scanning the document or sending the document as a fax,

that sort of thing.

So one approach you might take is just make just a single interface type machine interface.

And in this interface you would have methods for printing the document

and also may be faxing the document and scanning the document.

So this is generally OK, it's an OK interface if what you are looking for is a kind of multifunction

printer.

So if you have a multi function printer which can both scan and print and fax documents, then everything

is OK.

You simply implement the struct and then you go ahead and you actually implement all the all the interface

members.

So you just grab the machine interface and you generate basically all of this stuff for the printing

and for the faxing and for the scanning as well.

So there's absolutely no problem here.

However, imagine a different situation.

Imagine a situation where somebody is working with an old fashioned printer.

An old fashioned printer doesn't really have any scanning or faxing capabilities.

But because you want to implement this interface, because maybe some other APIs rely on the machine

interface, you have to implement this anyway.

You are being forced into implementing it.

So you go ahead through all the similar motions to implement the machine interface and you end up with

the same stuff as you would for a multifunction device, except there is a bit of a problem.

So certainly when you're working with an old fashioned printer implementing the printing capability,

it makes sense because a printer can print that.

That's what it does.

But it doesn't you don't really know what to do in the case of the faxing and in the case of the scanning

as well.

So one thing you can do is you can certainly leave the panic messages in here, except that you would

probably say something more meaningful, like operation not supported, because that's really what's

happening here.

It's not the case of implement me.

It's just that we don't support scanning from an old fashioned printer.

You can also, as an additional measure, you can add a comment which begins with the word deprecated.

So once again, these methods, they are not really deprecated where lying to the user a little bit.

But the consequence of having deprecated here would be that if you're using it, let's say you made

this old fashioned printer, you made this old fashioned printer like so and then you try to do a dot

dot scan.

Some IDs will actually cross out the scan option and they will tell you that this method is deprecated,

that you shouldn't be calling it.

So that's certainly what happens in my idea.

So that's one way that you could deal with the situation.

But really, we've created the problem by putting too much into an interface.

So we put both print and fax and scan into a single interface and then we expect everyone to implement

this, even if people don't actually have this functionality as part of the classes.

So they want the support of this interface because perhaps the interface is used in some sort of APIs,

but they really don't have anything to put into some of the implementations.

So how can we deal with this?

Well, we adhere to the interface segregation principle.

So the interface aggregation principle basically states that try to break up an interface into separate

parts that people will definitely need.

So there is no guarantee that if somebody needs printing, they also need faxing.

So in this particular example, it might make more sense to split up the printing and the scanning into

separate interfaces.

So you would have maybe the printer interface where you would have the print method.

You would also have the scanner interface, where you would have the scan method and so on and so forth.

And this way, this allows you to compose different types.

Out of the interfaces that you actually need, so, for example, if you just want something which is

only a printer.

So this is only a printer and nothing else, it doesn't scan and doesn't do anything.

In this case, what you do is you simply implement the printer interface.

So you take the printer interface and you implement it like so everything is OK.

You don't have to implement scan, you don't have to implement fax.

Everything is fine.

And then let's suppose that, for example, you have a photocopier.

So a photocopier can both print as well as a scan, so all you do here is you go ahead and you you implement

the printer interface like so and you also at the same time implement the scanner interface.

And that way you you basically get both of the functionality and you now have a photocopier, which

is both a printer as well as a scanner.

So let's not forget the fact that you can actually combine interfaces so you can compose an interface

out of other interfaces.

So if you want an interface that represents a multifunction device, then you can have your cake and

eat it too, because you can make a type called multifunction device, which is an interface.

And into this interface, you can put all the stuff from a printer and all the stuff from a scanner.

And, you know, if you had like a fax, for example, you would add that interface here as well.

So combining interfaces is fairly easy, not really complicated.

And of course, what you end up with if you adopt this approach is if you want to build a multifunction

machine and you've already got, let's say, a printer in a scanner implemented as separate components,

what you can do is you can use the decorator design pattern and then we'll talk about when we get to

the decorator design pattern.

But let me show you how this would work.

So if you want to build some sort of multi function machine, which is both a printer and a scanner,

what you can do is you can simply have the printer part and you can have this kind of part like so and

then you can implement the necessary interfaces.

So you would have a function which takes a multi function machine for printing, for example, where

you would simply reuse the functionality of the printer that you already have.

So you would say am the printer, not print the document and you would do the same thing for for the

scanner.

So you would have a method called scan, which would just say Amdocs scanner, DOT scan and pass that

in.

So you can see that with the interface aggregation approach, what you can do is, first of all, you

have very granular kind of definitions.

So you just grab the interfaces that you need and you don't have any extra members in those interfaces.

So if you're just building an ordinary printer, you just get the print method, then that's pretty

much it.

And of course, you you always have the ability of combining the interfaces.

So here you can combine the printer in the scanner and have a kind of interface aggregate and you can

subsequently have APIs which actually use this interface aggregate in your code.

So that's all there is to be said about the interface aggregation principle.
*/
