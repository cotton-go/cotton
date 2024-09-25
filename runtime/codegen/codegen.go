package codegen

// type Registration struct {
// 	Name      string       // full package-prefixed component name
// 	Iface     reflect.Type // interface type for the component
// 	Impl      reflect.Type // implementation type (struct)
// 	Routed    bool         // True if calls to this component should be routed
// 	Listeners []string     // the names of any weaver.Listeners
// 	NoRetry   []int        // indices of methods that should not be retried

// 	// Functions that return different types of stubs.
// 	// LocalStubFn   func(impl any, caller string, tracer trace.Tracer) any
// 	// ClientStubFn  func(stub Stub, caller string) any
// 	// ServerStubFn  func(impl any, load func(key uint64, load float64)) Server
// 	ReflectStubFn func(func(method string, ctx context.Context, args []any, returns []any) error) any

// 	// RefData holds a string containing the result of MakeEdgeString(Name, Dst)
// 	// for all components named Dst used by this component.
// 	RefData string
// }
