package dont

import (
	"fmt"
	"testing"
)

func TestDont(t *testing.T) {
	if err := Do(func() {
		fmt.Println("Something")
	}); err != nil {
		t.Error(err)
	}
}
