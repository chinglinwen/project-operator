package project

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"

	prettyjson "github.com/hokaccha/go-prettyjson"
	resty "gopkg.in/resty.v1"
)

var BaseURL string

func Init(baseurl string) {
	BaseURL = baseurl
}

const (
	EnvOnline    = "online"
	EnvPreOnline = "pre"
	EnvTest      = "test"
)

func GetNameAndEnv(pname string) (name, env string) {
	s := strings.Split(pname, "-")
	n := len(s)
	if n < 2 {
		return pname, ""
	}
	name = strings.Join(s[:n-1], "-")
	env = s[n-1]
	return
}

// // Project project release info
// type Project struct {
// 	Namespace string `json:"namespace,omitempty"`
// 	Project   string `json:"project,omitempty"` // event.Project.PathWithNamespace
// 	Version    string `json:"branch,omitempty"`  // parseVersion(event.Ref)
// 	// Env       string    `json:"env,omitempty"`  // default detect from branch, can be overwrite here
// 	UserName  string    `json:"user_name,omitempty"`
// 	UserEmail string    `json:"user_email,omitempty"`
// 	Message   string    `json:"message,omitempty"`
// 	Time      time.Time `json:"time,omitempty"`
// }

// Project project release info
type Project struct {
	// Namespace string    `yaml:"namespace,omitempty" json:"namespace,omitempty"`

	// event.Project.PathWithNamespace
	// Project string `yaml:"project,omitempty" json:"project,omitempty"`

	// parseVersion(event.Ref)
	Version string `yaml:"version,omitempty" json:"version,omitempty"`

	UserName       string `yaml:"userName,omitempty" json:"userName,omitempty"`
	UserEmail      string `yaml:"userEmail,omitempty" json:"userEmail,omitempty"`
	ReleaseMessage string `yaml:"releaseMessage,omitempty" json:"releaseMessage,omitempty"`
	ReleaseAt      string `yaml:"releaseAt,omitempty" json:"releaseAt,omitempty"`

	// test env need this to generate image tag
	// CommitId string `yaml:"commitid,omitempty" json:"commitid,omitempty"`

	// helper to avoid duplicate fields
	namespace string
	name      string
	env       string

	// lastApplied string
	// generation int64
}

func pretty(prefix, a interface{}) {
	out, _ := prettyjson.Marshal(a)
	fmt.Printf("%v: %s\n", prefix, out)
}

type ProjectOption func(*Project)

// func SetLastApplied(last string) ProjectOption {
// 	return func(p *Project) {
// 		p.lastApplied = last
// 	}
// }

// if program stop, it will invalid all cache
// we can detect creationTimestamp, or just store the cache?

// func SetGeneration(n int64) ProjectOption {
// 	return func(p *Project) {
// 		p.generation = n
// 	}
// }

func New(ns, pname string, spec Project, options ...ProjectOption) *Project {
	name, env := GetNameAndEnv(pname)
	p := &Project{
		Version:        spec.Version,
		UserName:       spec.UserName,
		UserEmail:      spec.UserEmail,
		ReleaseMessage: spec.ReleaseMessage,
		ReleaseAt:      spec.ReleaseAt,

		namespace: ns,
		name:      name,
		env:       env,
	}
	for _, op := range options {
		op(p)
	}
	return p
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type Alias Project
	return json.Marshal(&struct {
		Projectpath string `yaml:"project,omitempty" json:"Project"`
		*Alias
	}{
		Projectpath: p.getprojectpath(),
		Alias:       (*Alias)(p),
	})
}

type ProjectStatus struct {
	Status string `json:"status,omitempty"`

	// default detect from branch, can be overwrite here
	// Env       string    `yaml:"env,omitempty"`

	// to know if the deploy ready or not
	Image string `yaml:"image,omitempty" json:"image,omitempty"`
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

func (p *Project) getprojectpath() string {
	return p.namespace + "/" + p.name
}

// Apply call release api to apply to create project's yamls
func (p *Project) Apply() (out string, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	url := fmt.Sprintf("/api/apply/%v/%v", p.getprojectpath(), p.env)
	resp, e := resty. //SetDebug(true).
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
	code, msg := parseResponse(out)
	if code != 0 {
		err = fmt.Errorf("%v", msg)
		return
	}
	p.setcache()
	log.Printf("apply ok, output: %v\n", out)
	return
}

func parseResponse(body string) (int, string) {
	code := gjson.Get(body, "code").Int()
	message := gjson.Get(body, "message").String()
	return int(code), message
}

// Delete call release api to delete to create project's yamls
func (p *Project) Delete() (out string, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	url := fmt.Sprintf("/api/delete/%v/%v", p.getprojectpath(), p.env)
	resp, e := resty. //SetDebug(true).
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
	code, msg := parseResponse(out)
	if code != 0 {
		err = fmt.Errorf("%v", msg)
		return
	}

	p.delcache()
	log.Printf("delete ok, output: %v\n", out)
	return
}

// udate status
// deploy name? let's delegate?
// only release status ( apply ok or error )

// only if we can get current state? call api? or check k8s?
// we only need to compare the spec, not the underline deploy
// how to get the current spec? spec is generated by ci?
// so every apply, we just try to apply? let's cache last state
func (p *Project) UpdateProject() (err error) {
	log.Printf("try update project: %v/%v\n", p.namespace, p.name)

	// how to handle program restart? persist cache?
	// though, it can be apply again with the same yaml
	old := getcache(p.getprojectpath())
	if old == nil || changed(old, p) {
		pretty("project", p)

		var out string
		out, err = p.Apply()
		if err != nil {
			log.Printf("apply oerr, output: %v\n", out)
			err = fmt.Errorf("apply err: %v", err)
			return
		}
		log.Printf("apply ok, output: %v\n", out)
		return

	}
	log.Println("spec doesn't changed, so skip")

	// only image need to change? but the fields doesn't have image
	return
}

func (p *Project) CheckImageExist() (exist bool, err error) {
	url := fmt.Sprintf("/api/imagecheck/%v", p.getprojectpath())
	resp, e := resty. //SetDebug(true).
				R().
				SetHeader("Content-Type", "application/json").
				SetQueryParam("tag", p.Version).
				Get(BaseURL + url)
	if e != nil {
		err = e
		log.Printf("check image for %v, err: %v\n", url, err)
		return
	}
	out := string(resp.Body())
	code, msg := parseResponse(out)
	if code != 0 {
		err = fmt.Errorf("%v", msg)
		return
	}

	exist = gjson.Get(out, "data.exist").Bool()
	return
}

// see if updated, if so re-apply
func changed(old, new *Project) bool {
	// if new.generation == 1 {
	// 	return true
	// }
	if old.Version != new.Version {
		return true
	}
	if old.ReleaseAt != new.ReleaseAt {
		return true
	}
	return false
}

// last applied is always the same with current
// func getOldSpec(last string) (p *Project, err error) {
// 	type T struct {
// 		APIVersion string `json:"apiVersion"`
// 		Kind       string `json:"kind"`
// 		Metadata   struct {
// 			Annotations struct {
// 			} `json:"annotations"`
// 			Name      string `json:"name"`
// 			Namespace string `json:"namespace"`
// 		} `json:"metadata"`
// 		Spec Project `json:"spec"`
// 	}
// 	t := &T{}
// 	if err = json.Unmarshal([]byte(last), &t); err != nil {
// 		return
// 	}
// 	p = New(t.Metadata.Namespace, t.Metadata.Name, t.Spec)
// 	return

// }
