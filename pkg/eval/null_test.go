package eval

import (
	"testing"

	"github.com/qiushiyan/qlang/pkg/object"
)

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
