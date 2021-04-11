package environments

import (
	"os"
	"testing"
)

func Test_uNameOS(t *testing.T) {
	expected := os.Getenv("USER")
	actual := uNameOS()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

