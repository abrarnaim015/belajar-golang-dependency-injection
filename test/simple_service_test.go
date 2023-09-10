package test

import (
	"testing"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T)  {
	// simpleService, err := simple.InitializedService(true);
	// assert.NotNil(t, err)
	// assert.Nil(t, simpleService)
	if simpleService, err := simple.InitializedService(true); err != nil {
		assert.NotNil(t, err)
		assert.Nil(t, simpleService)
	}
}

func TestSimpleServiceSuccess(t *testing.T)  {
	if simpleService, err := simple.InitializedService(false); err != nil {
		assert.Nil(t, err)
		assert.NotNil(t, simpleService)
	}
}
