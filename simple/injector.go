//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDataBasePostgreSQL,
		NewDatabaseMongDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFoorepository, NewFooService)
var BarSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		BarSet,
		NewFooBarService,
	)
	return nil
}

// Salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBar() *FooBar {
	// wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar")) // sama dengan di bawah tp di panggi
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*")) // * langsung memasukan semuanya
	return nil
}

var FooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

// var fooValue = &Foo{}
// var barValue = &Bar{}

// sama dengan yg di atas hanya saja di tampung di varible menjadi value

func InitializedFooBarUseValue() *FooBar {
	wire.Build(FooBarValueSet, wire.Struct(new(FooBar), "*"))
	// wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// func InitializedConfiguration2() *Configuration {
// 	wire.Build(
// 		NewApplication2,
// 		wire.FieldsOf(new(*Application), "Configuration"),
// 	)
// 	return nil
// }

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
