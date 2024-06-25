package validator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
)

var (
	testStruct = Test{}
)

type Test struct {
	RequiredString   string   `validate:"required"`
	RequiredNumber   int      `validate:"required"`
	RequiredMultiple []string `validate:"required"`
}

func TestValidator(t *testing.T) {
	should := assert.New(t)
	if should.NoError(Init("zh")) {
		err := Validate(testStruct)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			fmt.Println(err)
		}
		fe := errs.Translate(trans)
		fmt.Println(fe)
	}
}
