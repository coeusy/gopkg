package cfg

import (
	"fmt"
	"testing"
)

func TestNewOption(t *testing.T) {
	opt := NewOption(
		WithPath("xyz"),
		WithConfigList([]string{"1", "2"}),
		WithoutDatasource(),
		WithFileType(FileTypeYaml),
	)
	fmt.Println(opt)
}
