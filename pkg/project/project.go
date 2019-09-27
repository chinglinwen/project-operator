package project

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	resty "gopkg.in/resty.v1"
)

var BaseURL string

func Init(baseurl string) {
	BaseURL = baseurl
}

// // Project project release info
// type Project struct {
// 	Namespace string `json:"namespace,omitempty"`
// 	Project   string `json:"project,omitempty"` // event.Project.PathWithNamespace
// 	Branch    string `json:"branch,omitempty"`  // parseBranch(event.Ref)
// 	// Env       string    `json:"env,omitempty"`  // default detect from branch, can be overwrite here
// 	UserName  string    `json:"user_name,omitempty"`
// 	UserEmail string    `json:"user_email,omitempty"`
// 	Message   string    `json:"message,omitempty"`
// 	Time      time.Time `json:"time,omitempty"`
// }

// Project project release info
type Project struct {
	// Namespace string    `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Project   string    `yaml:"project,omitempty" json:"project,omitempty"` // event.Project.PathWithNamespace
	Branch    string    `yaml:"branch,omitempty" json:"branch,omitempty"`   // parseBranch(event.Ref)
	Env       string    `yaml:"env,omitempty"`                              // default detect from branch, can be overwrite here
	UserName  string    `yaml:"user_name,omitempty" json:"user_name,omitempty"`
	UserEmail string    `yaml:"user_email,omitempty" json:"user_email,omitempty"`
	Message   string    `yaml:"message,omitempty" json:"message,omitempty"`
	Time      time.Time `yaml:"time,omitempty" json:"time,omitempty"`
}

type ProjectStatus struct {
	Status string `json:"status,omitempty"`
}

// mostly let's just call api

// we can actually just apply

// ready check?
// let's do a check before we can apply

// inited
// yaml ready
// image ready(dockerfile,nginx?)

// if image not ready call the build?

// for image based we don't know branch or tag? use latest tag?
// need somewhat auto mode?

// buildmode
// automode

// let's simplify build first
// let all call the api to build?

// three way of build
// 1. trx (gitlabcimode)
// 2. gitlab event(based on version?)
// 3. manual

// for all we just care the image tag
// what if tag updated ( should we try to update? set an flag? )

// latest tag can't know the env?

// func (p *Project) GetYaml() (out string, err error) {
// 	out = "geted yaml"
// 	return
// }

// get yaml ( should we check if yaml updated? )
// we can try example to validate

// // validate before apply
// func (p *Project) GetYaml() (out string, err error) {
// 	out = "geted yaml"
// 	return
// }

// so later, let trx generate this project yaml?

// Apply call release api to apply to create project's yamls
func (p *Project) Apply() (out string, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	url := fmt.Sprintf("/api/apply/%v", p.Project)
	resp, e := resty.SetDebug(true).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(b).
		Post(BaseURL + url)
	if e != nil {
		err = e
		log.Printf("get yaml for %v, err: %v\n", url, err)
		return
	}
	out = string(resp.Body())
	return
}

// Delete call release api to delete to create project's yamls
func (p *Project) Delete() (out string, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	url := fmt.Sprintf("/api/delete/%v", p.Project)
	resp, e := resty.SetDebug(true).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(b).
		Post(BaseURL + url)
	if e != nil {
		err = e
		log.Printf("get yaml for %v, err: %v\n", url, err)
		return
	}
	out = string(resp.Body())
	return
}

// udate status
// deploy name? let's delegate?
// only release status ( apply ok or error )

func UpdateProject(p Project) (err error) {
	// get exist one? then compare?

	// compare image? call api to fetch project info?
	// parse yaml?

	// only image need to change? but the fields doesn't have image

	return nil
}

// see if updated, if so re-apply
func compare(old, new *Project) bool {
	return false
}
