package test

import (
	"fmt"
	"testing"

	"github.com/raviMukti/training-golaang-restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}
