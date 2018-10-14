package main

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	res := generateToken("testUser")
	fmt.Println(string(res))
}
