package controller

import (
	"context"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/njhale/combo/api/v1alpha1"
)

type combinationController struct {
	client.Client

	log logr.Logger
}

func (c *combinationController) manageWith(mgr ctrl.Manager) error {
	// Create multiple controllers for resource types that require automatic adoption
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Combination{}).
		Complete(c)
}

func (c *combinationController) Reconcile(ctx context.Context, req ctrl.Request) (reconcile.Result, error) {
	// Set up a convenient log object so we don't have to type request over and over again
	log := c.log.WithValues("request", req)
	log.V(1).Info("reconciling bundle")

	in := &v1alpha1.Combination{}
	if err := c.Get(ctx, req.NamespacedName, in); err != nil {
		if apierrors.IsNotFound(err) {
			// If the instance is not found, we're likely reconciling because of a DELETE event.
			return reconcile.Result{}, nil
		}

		log.Error(err, "Error requesting Combination")
		return reconcile.Result{Requeue: true}, nil
	}

	return reconcile.Result{}, nil
}
