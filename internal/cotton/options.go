package cotton

import "reflect"

type Options struct {
	ConfigFilename string               // TOML config filename
	Config         string               // TOML config contents
	Fakes          map[reflect.Type]any // component fakes, by component interface type
	Quiet          bool                 // if true, do not print or log anything
}
