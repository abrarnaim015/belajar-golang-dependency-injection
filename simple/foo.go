package simple

type FooRepository struct {

}

func NewFoorepository() *FooRepository  {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(fooRepository *FooRepository) *FooService  {
	return &FooService{FooRepository: fooRepository}
}