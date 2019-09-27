package project

import (
	"fmt"
	projectv1alpha1 "wen/project-operator/pkg/apis/project/v1alpha1"
	"wen/project-operator/pkg/project"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

func updateProjectForCR(cr *projectv1alpha1.Project) (err error) {
	ns := cr.GetNamespace()
	name := cr.GetName()

	log.Info("creating project:", "name", ns+"/"+name)
	pretty("project:", cr)

	err = project.UpdateProject(ns, name, cr.Spec.Project)
	if err != nil {
		err = fmt.Errorf("UpdateProject err:%v", err)
		log.Println(err)
		return
	}
	return
}

func pretty(prefix, a interface{}) {
	out, _ := prettyjson.Marshal(a)
	fmt.Printf("%v: %s\n", prefix, out)
}
