package helper

import (
	"fmt"
	"testing"
)

func TestWrapSuccessResponse(t *testing.T) {
	bytes, _ := WrapResponseSuccess([]byte(`{"foo": "bar"}`))
	fmt.Println(string(bytes))
}
