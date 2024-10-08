package reflection

import (
	"fmt"
	"reflect"
)

func Type[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// ComponentName returns the name of the component of type T.
// Note that T must be the interface type of the component, not its implementation type.
func ComponentName[T any]() string {
	t := Type[T]()
	return fmt.Sprintf("%s/%s", t.PkgPath(), t.Name())
}
