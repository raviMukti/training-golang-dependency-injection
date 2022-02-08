package test

import (
	"fmt"
	"testing"

	"github.com/raviMukti/training-golaang-restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializeService()
	fmt.Println(simpleService.SimpleRepository)
}
