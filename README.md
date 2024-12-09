# Sealed Slices and Maps for Go

Package `sealed` is a small library implementing immutable slices and maps for Golang.

In other languages this is also known as constant vectors, read-only collections, persistent data-structures,
final instances, or frozen lists and hash maps.

## Usage
```go
import (
	"cmp"
	"fmt"
	"github.com/kamstrup/sealed"
)

// Sealed slices are created via a fluent API based on sealed.Builder
s := sealed.NewBuilder[string](0, 10). // allocate builder with len=0, cap=10
  Append("hello", "world"). // var args
  Collect(seq). // Append a Go 1.23 iterator
  Sort(cmp.Compare[string]).
  Seal() // converts the builder into a sealed.Slice[string]

// Print the elements in the slice
for _, str := range s.All {
	fmt.Println(str)
}
```

## Motivation

Sealed has grown out of the need for having `const` slices and maps. After writing a lot of Go code,
in particular in big projects with varying team members, it has become clear to me that maps and slices
are not well suited for public APIs in Go. Neither as parameters nor return values.

Slices and maps seem so simple and easy to use, but there are so many subtle pitfalls around ownership
and life cycle, making them extremely hard to use correctly everywhere in the long run.

Consider a simple function
```go
func (udb *UserDatabase) AddUsers(users []User) error {
	// do stuff with users
}
```
If the `UserDatabase` wants to hold on to the `users` slice after `AddUsers` returns, it *must* copy it.
On a small team it might, for example, be possible to agree (and document) that AddUsers steals
the users slice, for performance reasons, and callers should take heed and never use the `users`
slice again after the call.

However; with time, complexity, and changing developers, this is guaranteed to lead to tricky
data sharing bugs. No matter the clear warnings in the docs and the thoroughness of the code reviews.

Similarly, if we have a function that returns a slice of values
```go
func (udb *UserDatabase) Users() []Users {
	return udb.users
}
```
The developer team must agree through docs and conventions that _no one ever changes the returned slice
from Users()_! Given time, complexity, and changing team members, this assumption will inevitably break.
The only recourse is to **always copy the returned slice**.

To put it short: If you want to be absolutely safe about the integrity of your slices you **must
copy them every time they cross an API boundary!** The same thing goes for maps.
If you are dealing with performance sensitive software, this can be a severe restriction.

In an ideal world where "constness" could be enforced by the Go compiler, AddUsers() would look something
like the following
```go
// not valid Go
func (udb *UserDatabase) AddUsers(users const [] const User) error {
	// do stuff with users
}
```
The first `const` indicating that the wrapping slice cannot be modified, and the second `const` signifying
that the `User` elements in the slice cannot be changed either.

This library, `sealed`, deals with the first of these `const`s. We can use a `sealed.Slice[User]` to 
efficiently pass a read-only slice of users to the function. The second const on the User struct itself must
still be solved by the developer team themselves. Fx. by ensuring that the `User` struct is effectively immutable,
having no public fields and no mutator methods.

The completely safe variants of the discussed methods looks like
```go
func (udb *UserDatabase) AddUsers(users sealed.Slice[User]) error {
	// we can hold on to the users slice if we want, no need to copy because it is read-only.
}

func (udb *UserDatabase) Users() sealed.Slice[User] {
    // callers may hold on to the users slice we return, no need to copy because it is read-only.
}
```