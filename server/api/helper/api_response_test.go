package helper

import (
	"testing"
	"fmt"
)

func TestWrapSuccessResponse(t *testing.T) {
	bytes, _ := WrapResponseSuccess([]byte(`{"foo": "bar"}`))
	fmt.Println(string(bytes))
}

