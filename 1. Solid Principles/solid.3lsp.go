package solid

/* Liskov Substitution Principle
Isn't that applicable to Golang because it primarily deals with inheritance
Liskov Substitution Principle states: If you have some API that takes a base class,
and works correctly with that base class, it should also work correctly with the derived classes.

But unfortunately, in go we don't have base classes and derive classes.
This concept simply doesn't exist. So we will try to look at one of it's variation for go.

Base class ki assumption break nahi honi chahiye if you extend.
*/

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

//     vvv !! POINTER
func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func main_lsp() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}

/* Transcript
The Liskov of substitution principle, which is named after Barbara Lisk off, isn't really that applicable

to go because it primarily deals with inheritance.

So what the risk of substitution principle basically states that if you have some API that takes a base

class and works correctly with that base class, it should also work correctly with the derived class.

But unfortunately, in go we don't have base classes and derive classes.

This concept simply doesn't exist.

But I am going to try and show you a variation of the lisker substitution principle that does in fact

apply to go.

So let me show you an example.

Let's suppose that you're trying to deal with geometric shapes of a rectangular nature and you decide

to have an interface called size that allows you to specify certain operations on those types of constructs.

Like, for example, you might have the getters and setters for the width in the height.

So you would have an interface where you would have a get width, you would have set with you would

have to get height and maybe set height.

So now let's suppose you have a type called rectangle.

So a rectangle obviously has a width in height.

We can actually store them in just just ordinary fields, use lowercase here.

And then what you can do is you can go ahead and implement the sized interface on a rectangle like so

except what I would typically go for is having the pointer types here is going to help us out later

on.

When you work with the interface, obviously when you work with interfaces, it's a lot more convenient

to have pointers here.

So here you would return with here you would say are with equals width.

Here you would all return height.

And here you would say our height equals like so and so far so good.

I mean, this is something that you can work with and we can actually start experimenting with it.

So for example, I can write a function called Use It, which uses some sized object sized size.

So let's take a look at how this can works.

Let's suppose that I first get the width of this object.

So width is equal to size Donnette Width.

OK, so that's good.

And then let's suppose I set the height to 10 so sized dots at height.

Let's add it to the value of ten.

Now if you were to calculate the area of this sized object, you would expect the area to be ten multiplied

by the width because I've just said the height to ten and we got the width right here.

So the expected value of the expected area would be ten multiplied by the width.

But we can also calculate the actual area.

So the actual area would be size doget width multiplied by size.

Don't get height like so.

So we have an expected area and we have an actual area and we hope that they are the same value because

if they are not, that would be a big problem.

So here I would just do F.A. the print.

So we expected an area of expected area, but instead we got an actual area like so OK, so we can already

try out this entire thing, just just create a rectangle and try plugging it into this.

This is it.

A function that we created.

So let's make a rectangle, let's put an ampersand here, rectangle two by three and then we can use

it with this rectangle that are OK.

So running this, let's take a look at what we get.

So we expected an area of twenty and we go to an area of twenty.

It looks like everything is correct.

Everything is working fine.

And until you try to break the list of substitution principle, everything is OK.

But let's imagine that you decide to be smart and you decide to make a type called Square.

Notwist Square operates just like a rectangle because, well, it has the same members.

So you can say something like type square struct, which simply aggregates a rectangle because that's

square also has the width in the height.

However, let's imagine that you decide in your infinite wisdom that a square is going to enforce this

idea of with being equal to height.

Always.

So you're always going to enforce this so you will have some sort of constructor.

So a new square which takes a size

like.

So so here you would say askew is a square and askew with equal size, Asgeir height, equal size return

askew, something like that.

But in addition, what you'll do and this is really the part that violates the risk of substitution

principle is you'll have methods for set width and set height that set both the width and the height

in both of those situations.

So you'll have a function which takes a square

which has set width.

And here is the insidious part.

So not only do you set the width, but in order to keep this a square, you also said the height and

then you decide to do the same thing for the height center.

So you decide that you're going to have a set height.

Where, you know, but the height here where you set the within the height to the height value.

OK, so what is the problem with this?

Why is this breaking the risk of substitution principle?

So the risk of substitution principle basically states the following.

If you expecting some sort of something up the hierarchy, so to speak, in in this argument, then

it should continue to work, even if you proceed to extend objects which are already size.

So we took a rectangle and we decided to sort of extend the rectangle and make a square.

It should also continue to work.

Everything here should continue to work, but unfortunately it doesn't.

And if we try to plug in the square, you'll see how it badly goes wrong, because if I make a square

with a side of five, for example, and I call use it, we expect this whole thing to actually work.

We expect all of this to operate correctly.

But if I now run this, you'll see that the results are a bit disappointing.

We expected an area of 50 and we go to an area of 100.

So what happened here?

Well, obviously, what happened is this call to set height actually set not just the height, it also

set the width.

So the internal width of the square became inconsistent with the value of this variable right here.

And as a result, we're getting different values for expected area and actual area.

So the risk of substitution principle basically states that if you continue to use generalisations like

interfaces, for example, then you should not have inherited or you should not have implementors of

those generalisations break some of the assumptions which are set up at the higher level.

So at the higher level, we kind of assume that if you have a SIST object and you said it's height,

you are just setting its height, not both the height and the width.

And here what happened is we broke this assumption by setting both the width and the height, which

is a noble goal.

I mean, you can see how somebody would try to enforce the square invariant by setting both the within

the height.

It's a noble goal, but it doesn't work and it actually violates the risk of substitution principle.

So the risk of substitution from principle is actually one of those situations where there is no right

answer.

There is no right solution to this problem.

I mean, you can take different approaches to how you would take this.

For example, we can say that squares don't exist, that since every square is a rectangle, we don't

work with squares at all.

Or, for example, you could do is you could explicitly make make illegal states on representable,

so to speak.

So basically you can say that a square doesn't really have width and height, a square has a size,

and that's pretty much it.

So you could have some type called square two, which would have a size, which is an end, which would

double as both width and height.

So it would represent both of these states.

So instead of aggregating a rectangle, you have its own member.

And then if you want to represent the square as a rectangle, well, you can have your cake and eat

it, too.

You can have a method which just is called rectangle, for example, which returns a rectangle.

And here you would just construct a new rectangle with that size.

Comma is that size.

That's pretty much all that you would have to do.

So this is one approach to the problem.

But just just to recap, basically, the idea of the risk of substitution principle is that the behaviour

of implementors of a particular type, like in this case, the size the interface should not break the

core fundamental behaviors that you rely on.

So you should be able to continue taking sized objects instead of somehow figuring out in here, for

example, by doing type checks whether you have a rectangle or a square, it should still work in the

generalised case.

And so prime example of the violation of this problem is what we've done here.

So we've broken certain assumptions about the type and as a result, we got incorrect behaviour.

So that's what I wanted to show you about the risk of substitution principle.
*/
