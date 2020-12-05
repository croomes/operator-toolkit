package v1

//go:generate mockgen -destination=../../mocks/mock_reconciler.go -package=mocks github.com/darkowlzz/composite-reconciler/controller/v1 Controller

import (
	"context"

	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Controller is the controller interface that must be implemented by a
// composite controller. It provides methods required for reconciling a
// composite controller.
type Controller interface {
	// InitReconciler initializes the reconciler with a given request and
	// context. Logger setup and initialization of an instance of the primary
	// object can also be done here.
	InitReconcile(context.Context, ctrl.Request)

	// FetchInstance queries the latest version of the primary object the
	// controller is responsible for. If the object is not found, a "not found"
	// error is expected in return.
	FetchInstance() error

	// Apply default values to the primary object spec. Use this in case a
	// defaulting webhook has not been deployed.
	Default()

	// Validate validates the primary object spec before it's created. It
	// ensures that all required fields are present and valid. Use this in case
	// a validating webhook has not been deployed.
	Validate() error

	// SaveClone clones and saves the original request instance after
	// defaulting and validating. This is later used in PatchStatus() to check
	// if there's a change in the status and patches the status if required.
	SaveClone()

	// IsUninitialized checks the primary object instance fetched in
	// FetchInstance() and determines if the object has not been initialized.
	// This is usually done by checking the status of the object. This is
	// checked before running Initialize().
	IsUninitialized() bool

	// Initialize sets the provided initialization condition on the object
	// status. This helps some operations in Operation() know about the
	// creation phase and run initialization specific operations.
	Initialize(conditionsv1.Condition) error

	// UpdateStatus queries the status of the child objects and based on them,
	// sets the status of the primary object instance. It doesn't save the
	// updated object in the API. API update is done in PatchStatus() after
	// collecting and comparing all the status updates.
	UpdateStatus() error

	// Operate runs the core operation of the controller that ensures that
	// the child objects or the other objects and configurations in the
	// environment are in the desired state. It should be able to update any
	// existing resources or create one, if there's a configuration drift,
	// based on the type of objects.
	// The returned result is the returned reconcile result.
	Operate() (result ctrl.Result, err error)

	// PatchStatus compares the original primary object instance status with
	// the reconciled primary object status and patches the API object if
	// required.
	PatchStatus() error

	// GetObjectMetadata returns the resource metadata of the primary object.
	// This is usually used to check if a resource is marked for deletion.
	GetObjectMetadata() metav1.ObjectMeta

	// AddFinalizer adds a finalizer to the primary object's metadata. This is
	// used when the controller has a custom cleanup operation, not based on
	// garbage collection using owner reference of the resources.
	AddFinalizer(string) error

	// Cleanup runs the custom cleanup operation to delete or undo the changes
	// made by the controller. This can be empty for controllers that use owner
	// reference based garbage collection for cleanup. For controllers with
	// custom cleanup requirement, the cleanup logic can be defined here.
	Cleanup() (result ctrl.Result, err error)

	// UpdateConditions updates the status condition of local instance of the
	// primary object with the given conditions. It doesn't update the API
	// object in the API server. Use UpdateStatus() to update the actual status
	// of the object.
	UpdateConditions([]conditionsv1.Condition)
}
