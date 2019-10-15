package project

import (
	"fmt"
	projectv1alpha1 "wen/project-operator/pkg/apis/project/v1alpha1"
	"wen/project-operator/pkg/project"

	"errors"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

func convertSpec(in *projectv1alpha1.ProjectSpec) project.Project {
	return project.Project{
		// Project:        in.Project,
		Version:        in.Version,
		UserName:       in.UserName,
		UserEmail:      in.UserEmail,
		ReleaseMessage: in.ReleaseMessage,
		ReleaseAt:      in.ReleaseAt,
	}
}

var ErrImageNotExist = errors.New("image not exist yet, waiting...")

func (r *ReconcileProject) updateProjectForCR(instance *projectv1alpha1.Project) (err error) {
	ns := instance.GetNamespace()
	name := instance.GetName()

	log.Info("creating project:", "ns", ns, "name", name)
	pretty("project instance", instance)

	// last := cr.GetAnnotations()["kubectl.kubernetes.io/last-applied-configuration"]
	// n := cr.GetGeneration()

	spec := convertSpec(&instance.Spec)
	p := project.New(ns, name, project.Project(spec))
	// project.SetLastApplied(last),
	// project.SetGeneration(n))

	exist, err := p.CheckImageExist()
	if err != nil {
		log.Info("image check err: ", err)
		err = nil
		// return
	}
	_ = exist
	// if !exist {
	// 	err = ErrImageNotExist
	// 	return
	// }

	err = p.UpdateProject()
	if err != nil {
		// err = fmt.Errorf("UpdateProject err:%v", err)
		log.Info("update project", "error", err)
		return
	}
	return
}

func pretty(prefix, a interface{}) {
	out, _ := prettyjson.Marshal(a)
	fmt.Printf("%v: %s\n", prefix, out)
}
