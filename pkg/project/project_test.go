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
		Version: "v1.0.0",
	})
	out, err := p.Apply()
	if err != nil {
		t.Error("apply err", err)
	}
	fmt.Println("out: ", out)
}

func TestDelete(t *testing.T) {
	p := New("demo", "hello", Project{
		Version: "v1.0.0",
	})
	out, err := p.Delete()
	if err != nil {
		t.Error("apply err", err)
	}
	fmt.Println("out: ", out)
}

func TestCheckImageExist(t *testing.T) {
	p := New("ops", "netshoot3", Project{
		Version: "latest",
	})
	out, err := p.CheckImageExist()
	if err != nil {
		t.Error("check err", err)
	}
	fmt.Println("out: ", out)
}

// func TestGetOldSpec(t *testing.T) {
// 	last := "{\"apiVersion\":\"project.haodai.com/v1alpha1\",\"kind\":\"Project\",\"metadata\":{\"annotations\":{},\"name\":\"demo\",\"namespace\":\"default\"},\"spec\":{\"branch\":\"v1.0.0\"}}\n"

// 	p, err := getOldSpec(last)
// 	if err != nil {
// 		t.Error("apply err", err)
// 	}
// 	fmt.Println("p: ", p)
// }

func TestGetNameAndEnv(t *testing.T) {
	cases := []struct {
		pname string
		name  string
		env   string
	}{
		{pname: "", name: "", env: ""},
		{pname: "api-", name: "api", env: ""},
		{pname: "-online", name: "", env: "online"},
		{pname: "api-online", name: "api", env: "online"},
		{pname: "xdy-api-online", name: "xdy-api", env: "online"},
	}
	for _, v := range cases {
		n, e := GetNameAndEnv(v.pname)
		if n != v.name || e != v.env {
			t.Errorf("for %v err: want name: %v, got %v, want env: %v, got %v\n",
				v.pname, v.name, n, v.env, e)
		}
	}

}
