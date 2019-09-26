package project

import (
	"context"

	projectv1alpha1 "wen/project-operator/pkg/apis/project/v1alpha1"
	"wen/project-operator/pkg/project"

	"github.com/go-logr/logr"
)

const projectFinalizer = "finalizer.project.haodai.com"

func (r *ReconcileProject) finalizeProject(reqLogger logr.Logger, m *projectv1alpha1.Project) error {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted. Examples
	// of finalizers include performing backups and deleting
	// resources that are not owned by this CR, like a PVC.

	// Delete deploy too? let's call api to delete
	err := project.DeleteProject(m.Spec.TaskName)
	if err != nil {
		reqLogger.Error(err, "Failed to delete project with finalizer")
		return err
	}
	reqLogger.Info("Successfully finalized project")
	return nil
}

func (r *ReconcileProject) addFinalizer(reqLogger logr.Logger, m *projectv1alpha1.Project) error {
	reqLogger.Info("Adding Finalizer for the Memcached")
	m.SetFinalizers(append(m.GetFinalizers(), projectFinalizer))

	// Update CR
	err := r.client.Update(context.TODO(), m)
	if err != nil {
		reqLogger.Error(err, "Failed to update project with finalizer")
		return err
	}
	return nil
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func remove(list []string, s string) []string {
	for i, v := range list {
		if v == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}
