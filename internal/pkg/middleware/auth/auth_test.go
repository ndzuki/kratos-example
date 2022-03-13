package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerarteToken("realworld", "nd")
	spew.Dump(tk)
	panic(1)
}
