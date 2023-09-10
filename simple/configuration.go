package simple

type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{
			Name: "AbrarNaim",
		},
	}
}

// func NewApplication2(configuration *Configuration) *Application {
// 	return &Application{
// 		&Configuration{
// 			Name: "AbrarNaim2",
// 		},
// 	}
// }