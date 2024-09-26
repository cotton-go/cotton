package cotton

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cotton-go/cotton/internal/single"
	"github.com/cotton-go/cotton/runtime"
	"github.com/cotton-go/cotton/runtime/codegen"
	"github.com/cotton-go/cotton/runtime/protos"
)

func parseSingleConfig(regs []*codegen.Registration, filename, contents string) (*single.SingleConfig, error) {
	// Parse the config file, if one is given.
	config := &single.SingleConfig{App: &protos.AppConfig{}}
	if contents != "" {
		app, err := runtime.ParseConfig(filename, contents, codegen.ComponentConfigValidator)
		if err != nil {
			return nil, fmt.Errorf("parse config: %w", err)
		}
		// if err := runtime.ParseConfigSection(single.ConfigKey, single.ShortConfigKey, app.Sections, config); err != nil {
		// 	return nil, fmt.Errorf("parse config: %w", err)
		// }
		config.App = app
	}

	if config.App.Name == "" {
		config.App.Name = filepath.Base(os.Args[0])
	}

	// Validate listeners in the config.
	listeners := map[string]struct{}{}
	for _, reg := range regs {
		for _, listener := range reg.Listeners {
			listeners[listener] = struct{}{}
		}
	}
	for listener := range config.Listeners {
		if _, ok := listeners[listener]; !ok {
			return nil, fmt.Errorf("listener %s (in the config) not found", listener)
		}
	}

	return config, nil
}
