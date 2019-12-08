package main

// package test_base

import (
	"fmt"
	"testing"

	ggUuid "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

func TestUuid001(t *testing.T) {
	uid1 := ggUuid.New().String()
	uid2 := ggUuid.New().String()
	fmt.Println("uid1:", uid1)
	fmt.Println("uid2:", uid2)

	// uid1: b1f1928d-5617-40c9-8e53-1a8f3fc52686
	// uid2: 1d807942-cbef-487a-b603-f6f158cd3492
}

func TestUuid002(t *testing.T) {
	u1 := uuid.Must(uuid.FromString("123e4567-e89b-12d3-a456-426655440000"))
	fmt.Printf("u1: %s\n", u1)

	u2 := uuid.NewV4()
	fmt.Printf("u2: %s\n", u2.String())

	// todo: sdfsdf
	// Parsing UUID from string input
	u3, err3 := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err3 != nil {
		fmt.Printf("Something went wrong: %s", err3)
	}
	fmt.Printf("Successfully parsed: %s", u3)

	// u1: 123e4567-e89b-12d3-a456-426655440000
	// u2: fd0d00f0-a7b6-485f-8fae-9aa835f19cbe
}
