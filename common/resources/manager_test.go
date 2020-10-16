package resources

import (
	"testing"
)

func TestReadResources(t *testing.T) {
	_, err := ReadResource("/resources/jetbrainsmono")
	t.Fatal(err)
}
