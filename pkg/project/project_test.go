package project

import (
	"fmt"
	"testing"
)

func init() {
	BaseURL = "http://release-test.newops.haodai.net"
}

func TestApply(t *testing.T) {
	p := &Project{
		Project: "demo/hello",
		Branch:  "v1.0.0",
	}
	out, err := p.Apply()
	if err != nil {
		t.Error("apply err", err)
	}
	fmt.Println("out: ", out)
}
