package project

import "time"

// Project project release info
type Project struct {
	Namespace string    `json:"namespace,omitempty"`
	Project   string    `json:"project,omitempty"` // event.Project.PathWithNamespace
	Branch    string    `json:"branch,omitempty"`  // parseBranch(event.Ref)
	Env       string    `json:"env,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	UserEmail string    `json:"user_email,omitempty"`
	Message   string    `json:"message,omitempty"`
	Time      time.Time `json:"time,omitempty"`
}

type ProjectStatus struct {
	Status string `json:"status,omitempty"`
}

// mostly let's just call api

// get yaml ( should we check if yaml updated? )
// we can try example to validate

// validate before apply

// apply

// see if updated, if so re-apply

// udate status
// deploy name? let's delegate?
// only release status ( apply ok or error )

// delete as needed

// so later, let trx generate this project yaml?
