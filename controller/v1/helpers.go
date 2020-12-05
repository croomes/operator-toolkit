package v1

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DefaultIsUninitialized performs uninitialized check on an object based on
// the status conditions.
func DefaultIsUninitialized(conditions []conditionsv1.Condition) bool {
	return conditions == nil
}

// HasFinalizer returns true if obj has the named finalizer.
func HasFinalizer(obj metav1.Object, name string) bool {
	for _, item := range obj.GetFinalizers() {
		if item == name {
			return true
		}
	}
	return false
}

// AddFinalizer adds the named finalizer to obj, if it isn't already present.
func AddFinalizer(obj metav1.Object, name string) {
	if HasFinalizer(obj, name) {
		// It's already present, so there's nothing to do.
		return
	}
	obj.SetFinalizers(append(obj.GetFinalizers(), name))
}

// RemoveFinalizer removes the named finalizer from obj, if it's present.
func RemoveFinalizer(obj metav1.Object, name string) {
	finalizers := obj.GetFinalizers()
	for i, item := range finalizers {
		if item == name {
			obj.SetFinalizers(append(finalizers[:i], finalizers[i+1:]...))
			return
		}
	}
	// We never found it, so it's already gone and there's nothing to do.
}

func contains(slice []string, s string) bool {
	for _, element := range slice {
		if element == s {
			return true
		}
	}
	return false
}
