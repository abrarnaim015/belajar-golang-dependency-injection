package test

import (
	"testing"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T)  {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
// 	Close Connection &{Database}
//  Close file Database
}