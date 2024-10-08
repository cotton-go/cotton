package cotton

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"sync"
	"time"

	"github.com/cotton-go/cotton/internal/config"
	"github.com/cotton-go/cotton/runtime/codegen"
)

type SingleWeavelet struct {
	ctx context.Context // the propagated context

	// Registrations.
	regs       []*codegen.Registration                // registered components
	regsByName map[string]*codegen.Registration       // registrations by component name
	regsByIntf map[reflect.Type]*codegen.Registration // registrations by component interface type
	regsByImpl map[reflect.Type]*codegen.Registration // registrations by component implementation type

	// Options, config, and metadata.
	opts         Options   // options
	deploymentId string    // globally unique deployment id
	id           string    // globally unique weavelet id
	createdAt    time.Time // time at which the weavelet was created

	// Logging, tracing, and metrics.
	// pp *logging.PrettyPrinter // pretty printer for logger

	// Components and listeners.
	mu         sync.Mutex              // guards the following fields
	components map[string]any          // components, by name
	listeners  map[string]net.Listener // listeners, by name
}

func NewSingleWeavelet(ctx context.Context, regs []*codegen.Registration, opts Options) (*SingleWeavelet, error) {

	return nil, nil
}

func (w *SingleWeavelet) GetImpl(t reflect.Type) (any, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.getImpl(t)
}

func (w *SingleWeavelet) getImpl(t reflect.Type) (any, error) {
	reg, ok := w.regsByImpl[t]
	if !ok {
		return nil, fmt.Errorf("component implementation %v not found; maybe you forgot to run weaver generate", t)
	}
	return w.get(reg)
}

func (w *SingleWeavelet) get(reg *codegen.Registration) (any, error) {
	if c, ok := w.components[reg.Name]; ok {
		// The component has already been created.
		return c, nil
	}

	if fake, ok := w.opts.Fakes[reg.Iface]; ok {
		// We have a fake registered for this component.
		return fake, nil
	}

	// Create the component implementation.
	v := reflect.New(reg.Impl)
	obj := v.Interface()

	// Fill config.
	if cfg := config.Config(v); cfg != nil {
		fmt.Println("todo: fill config")
		// if err := runtime.ParseConfigSection(reg.Name, "", w.config.App.Sections, cfg); err != nil {
		// 	return nil, err
		// }
	}

	// Set logger.
	// if err := SetLogger(obj, w.logger(reg.Name)); err != nil {
	// 	return nil, err
	// }

	// // Set application runtime information.
	// if err := SetWeaverInfo(obj, w.weaverInfo); err != nil {
	// 	return nil, err
	// }

	// // Fill ref fields.
	// if err := FillRefs(obj, func(t reflect.Type) (any, error) {
	// 	return w.getIntf(t, reg.Name)
	// }); err != nil {
	// 	return nil, err
	// }

	// // Fill listener fields.
	// if err := FillListeners(obj, func(name string) (net.Listener, string, error) {
	// 	lis, err := w.listener(name)
	// 	return lis, "", err
	// }); err != nil {
	// 	return nil, err
	// }

	// Call Init if available.
	if i, ok := obj.(interface{ Init(context.Context) error }); ok {
		if err := i.Init(w.ctx); err != nil {
			return nil, fmt.Errorf("component %q initialization failed: %w", reg.Name, err)
		}
	}

	w.components[reg.Name] = obj
	return obj, nil
}
