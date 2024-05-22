package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("hello", "oswin")
	spew.Dump(token) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE0NDQ0Nzg0MDAsInVzZXJuYW1lIjoib3N3aW4ifQ.dx9gRl708hXitRu0ife4ltvzIdVclsDEnuUF--Gzva8
}
