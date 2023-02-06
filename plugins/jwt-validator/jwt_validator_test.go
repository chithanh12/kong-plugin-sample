package jwt_validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	tk := ""

	p, e := VerifyToken(tk)
	assert.NoError(t, e)
	fmt.Println(p)
}
