package version_test

import (
	"fmt"
	"testing"

	"github.com/summingyu/restful-api-demo/version"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
