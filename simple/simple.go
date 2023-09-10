package simple

import "errors"

type SimpleRepository struct {
	Eror bool
}

func NewSimpleRepository(isError bool) *SimpleRepository  {
	return &SimpleRepository{
		Eror: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error)  {
	if repository.Eror {
		return nil, errors.New("Failed Create Service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}