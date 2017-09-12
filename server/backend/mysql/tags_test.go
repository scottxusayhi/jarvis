package mysql

import (
	"testing"
	"fmt"
)

func TestJarvisMysqlBackend_AttachTag(t *testing.T) {
	backend, _ := GetBackend()
	backend.AttachTag("20", "master")
	backend.AttachTag("20", "slave")
	backend.AttachTag("20", "as")
	backend.AttachTag("20", "stream")
}

func TestJarvisMysqlBackend_RemoveTag(t *testing.T) {
	backend, _ := GetBackend()
	backend.RemoveTag("20", "master")
	backend.RemoveTag("20", "slave")
	backend.RemoveTag("20", "as")
	backend.RemoveTag("20", "stream")
}

func TestJarvisMysqlBackend_ListTags(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.ListTags())
}



