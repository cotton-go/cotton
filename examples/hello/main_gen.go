package main

import (
	"reflect"

	"github.com/cotton-go/cotton/runtime/codegen"
)

func init() {
	codegen.Register(codegen.Registration{
		Name: "github.com/cotton-go/cotton/Main",
		// Iface: reflect.TypeOf((*cotton.Main)(nil)).Elem(),
		Impl: reflect.TypeOf(app{}),
		// Listeners: []string{"hello"},
		// LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
		// 	return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		// },
		// ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		// ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
		// 	return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		// },
		// ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
		// 	return main_reflect_stub{caller: caller}
		// },
		// RefData: "⟦17f36ff9:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→hello⟧\n",
	})
}
