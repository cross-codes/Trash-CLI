package functions

import (
  "os"
)

func FileIsReadable(filename string) bool  {
  info, err := os.Stat(filename)
  if err != nil {
    return false
  }
  return info.Mode().Perm()&0444 == 0444
}
