package utils

import (
    "testing"
)

func TestAbsInt64(t *testing.T) {
    var a int64 = -1
    if AbsInt64(a) != 1 {
        t.Errorf("error")
    }
}

