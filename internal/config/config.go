package config

import (
	"fmt"
	"reflect"
	"strings"
)

func Config(v reflect.Value) any {
	// TODO(mwhittaker): Delete this function and use weaver.GetConfig instead.
	// Right now, there are some cyclic dependencies preventing us from doing
	// this.
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("invalid non pointer to struct value: %v", v))
	}
	s := v.Elem()
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		// Check that f is an embedded field of type weaver.WithConfig[T].
		f := t.Field(i)
		if !f.Anonymous ||
			f.Type.PkgPath() != "github.com/ServiceWeaver/weaver" ||
			!strings.HasPrefix(f.Type.Name(), "WithConfig[") {
			continue
		}

		// Call the Config method to get a *T.
		config := s.Field(i).Addr().MethodByName("Config")
		return config.Call(nil)[0].Interface()
	}
	return nil
}
