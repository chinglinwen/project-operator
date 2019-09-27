package project

import (
	"fmt"
	projectv1alpha1 "wen/project-operator/pkg/apis/project/v1alpha1"
	"wen/project-operator/pkg/project"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

func convertSpec(in *projectv1alpha1.ProjectSpec) project.Project {
	return project.Project{
		// Project:        in.Project,
		Branch:         in.Branch,
		UserName:       in.UserName,
		UserEmail:      in.UserEmail,
		ReleaseMessage: in.ReleaseMessage,
		ReleaseAt:      in.ReleaseAt,
		CommitId:       in.CommitId,
	}
}

func updateProjectForCR(cr *projectv1alpha1.Project) (err error) {
	ns := cr.GetNamespace()
	name := cr.GetName()

	log.Info("creating project:", "ns", ns, "name", name)
	pretty("project cr", cr)

	// last := cr.GetAnnotations()["kubectl.kubernetes.io/last-applied-configuration"]
	// n := cr.GetGeneration()

	spec := convertSpec(&cr.Spec)
	p := project.New(ns, name, project.Project(spec))
	// project.SetLastApplied(last),
	// project.SetGeneration(n))

	err = p.UpdateProject()
	if err != nil {
		err = fmt.Errorf("UpdateProject err:%v", err)
		log.Info("update project", "error", err)
		return
	}
	return
}

func pretty(prefix, a interface{}) {
	out, _ := prettyjson.Marshal(a)
	fmt.Printf("%v: %s\n", prefix, out)
}
