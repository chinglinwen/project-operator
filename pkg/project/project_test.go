package project

import (
	"fmt"
	"testing"
)

func init() {
	BaseURL = "http://release-test.newops.haodai.net"
}

func TestApply(t *testing.T) {
	p := New("demo", "hello", Project{
		Branch: "v1.0.0",
	})
	out, err := p.Apply()
	if err != nil {
		t.Error("apply err", err)
	}
	fmt.Println("out: ", out)
}

func TestDelete(t *testing.T) {
	p := New("demo", "hello", Project{
		Branch: "v1.0.0",
	})
	out, err := p.Delete()
	if err != nil {
		t.Error("apply err", err)
	}
	fmt.Println("out: ", out)
}
