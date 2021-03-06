/*

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

package operatorversion

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kudoapi "github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
)

// Reconciler reconciles an OperatorVersion object
type Reconciler struct {
	client.Client
}

// SetupWithManager registers this reconciler with the controller manager
func (r *Reconciler) SetupWithManager(
	mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kudoapi.OperatorVersion{}).
		Complete(r)
}

// Reconcile reads that state of the cluster for an OperatorVersion object and makes changes based on the state read
// and what is in the OperatorVersion.Spec.
//
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
func (r *Reconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	// Fetch the operator version
	operatorVersion := &kudoapi.OperatorVersion{}
	err := r.Get(context.TODO(), request.NamespacedName, operatorVersion)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	log.Printf("OperatorVersionController: Received Reconcile request for an operatorVersion named: %v", request.Name)

	// TODO: Validate OperatorVersion is appropriate.
	return reconcile.Result{}, nil
}
