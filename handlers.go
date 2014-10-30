package fail

import "log"

// Require can wrap a function that returns error. If there is an error, it is output
// with log.Fatalf("fail.Require: %s\n", err).
// Example: fail.Fail(foo.Close())
func Fail(err error) {
    if err != nil {
        log.Fatalf("fail: %s\n", err)
    }
}
