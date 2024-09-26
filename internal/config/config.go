package config

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/cotton-go/cotton/runtime"
	"github.com/cotton-go/cotton/runtime/protos"
)

func Config(conf *protos.AppConfig, v reflect.Value) any {
	// TODO(mwhittaker): Delete this function and use weaver.GetConfig instead.
	// Right now, there are some cyclic dependencies preventing us from doing
	// this.
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("invalid non pointer to struct value: %v", v))
	}
	s := v.Elem()
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.Anonymous || !strings.HasPrefix(f.Type.Name(), "WithConfig[") {
			continue
		}

		key := f.Tag.Get("config")
		config := s.Field(i).Addr().MethodByName("Config")
		cfg := config.Call(nil)[0].Interface()
		if err := runtime.ParseConfigSection(key, "", conf.Sections, cfg); err != nil {
			slog.Warn("failed to parse config section", "err", err)
			return err
		}
		// slog.Info("config", "key", key, "cfg", cfg)
	}
	return nil
}
