fail
====

A package of Go functions that fail the program instead of returning errors.

I admit this might be a silly idea. I'm gonna try it anyways in sprok to see how it looks.

```go

package fail

// A collection of fail-fast wrapper functions that save a few lines of code.
// Since I do a lot of OS interactions with my Go programs, I end up with a lot
// of boilerplate code around common operations such as opening files and other things
// that *should* never fail, so if they do the whole program crashes. Over time
// I've used this technique consistently to get programs to a state where they can
// run unattended a lot faster and with less code.
//
// This doesn't save a whole lot, but for quick programs like https://github.com/tobert/sprok

// Require can wrap a function that returns error. If there is an error, it is output
// with log.Fatalf("fail.Require: %s\n", err).
// Example: fail.Require(foo.Close())
func Require(err error) {
    if err != nil {
        log.Fatalf("fail.Require: %s\n", err)
    }
}

// whatever os.Open is with the error triggering a fail
func Open(name string) (file *File) { }
func OpenWrite(name string) (file *File) { }
// etc.
```
