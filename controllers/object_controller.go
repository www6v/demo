/*
Copyright 2022 xiong.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	infrav1 "demo/api/v1"
)

// ObjectReconciler reconciles a Object object
type ObjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infra.demo.com,resources=objects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infra.demo.com,resources=objects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infra.demo.com,resources=objects/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Object object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ObjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = log.FromContext(ctx)

    // ctx := context.Background()
    _ = log.Log.WithValues("object", req.NamespacedName)
    // your logic here

    // 1. Print Spec.Detail and Status.Created in log
    obj := &infrav1.Object{}
    if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
        fmt.Errorf("couldn't find object:%s", req.String())
    } else {
    //打印Detail和Created
        log.Log.V(1).Info("Successfully get detail", "Detail", obj.Spec.Detail)
        log.Log.V(1).Info("", "Created", obj.Status.Created)
    }
    // 2. Change Created
    if !obj.Status.Created {
        obj.Status.Created = true
        r.Update(ctx, obj)
    }

    return ctrl.Result{}, nil

	//return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ObjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.Object{}).
		Complete(r)
}
