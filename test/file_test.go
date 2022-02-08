package test

import (
	"testing"

	"github.com/raviMukti/training-golaang-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
