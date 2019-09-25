package controller

import (
	"wen/project-operator/pkg/controller/project"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, project.Add)
}
