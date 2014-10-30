package fail

import "os"

// whatever os.Open is with the error triggering a fail
func Open(name string) (file *File) {
  file, err := os.Open(name)
  Fail(err)
  return file
}

func OpenAppend(name string) (file *File) {

}
