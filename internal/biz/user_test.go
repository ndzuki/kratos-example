package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("sdf")
	spew.Dump(s)
	panic(1)
}

func TestVerifyPassword(t *testing.T) {

	Equal(t, verifyPassword("$2a$10$VamKs4nfNfrVmLGYy9lzWenzc8N3Br0aiIt49qjPW8fnPTk9G6riK", "sdf"), true)
	NotEqual(t, verifyPassword("$2a$10$dtBh3B5d5PRYQ53VlfyC", "sdf"), true)
}
