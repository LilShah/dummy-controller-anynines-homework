/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	interviewcomv1alpha1 "github.com/LilShah/dummy-controller/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"
)

// DummyReconciler reconciles a Dummy object
type DummyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=interview.com,resources=dummies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=interview.com,resources=dummies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=interview.com,resources=dummies/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=create;update;patch;get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *DummyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	dummyInstance := &interviewcomv1alpha1.Dummy{}
	err := r.Get(ctx, req.NamespacedName, dummyInstance)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info(fmt.Sprintf("Dummy resource %s deleted", req.NamespacedName))
		}
		return ctrl.Result{}, nil
	}

	// Step 2, logging details
	log.Info(fmt.Sprintf("Dummy resource %s in namespace %s, with the message: %s", dummyInstance.Name, dummyInstance.Namespace, dummyInstance.Spec.Message))

	// Step 3, echo in status
	dummyPatchBase := client.MergeFrom(dummyInstance.DeepCopy())
	dummyInstance.Status.SpecEcho = dummyInstance.Spec.Message
	err = r.Status().Patch(ctx, dummyInstance, dummyPatchBase)
	if err != nil {
		log.Error(err, "unable to patch Dummy object's status")
		return ctrl.Result{}, err
	}

	// Step 4, associate Pod to each Dummy
	pod := &corev1.Pod{}
	if err := r.Get(ctx, types.NamespacedName{Name: dummyInstance.Name, Namespace: dummyInstance.Namespace}, pod); err != nil {
		if apierrors.IsNotFound(err) {
			// Create only if pod doesn't exist
			pod = getPodManifest(dummyInstance, r.Scheme)
			if err := r.Create(ctx, pod); err != nil {
				if !apierrors.IsAlreadyExists(err) {
					log.Error(err, "unable to create pod for dummy")
					return ctrl.Result{}, err
				}
			}
			log.Info(fmt.Sprintf("Pod %s created for dummy instance %s", dummyInstance.Name, pod.Name))
			// reconcile after creating, get updated pod manifest
			return ctrl.Result{Requeue: true}, nil
		}
		log.Error(err, "unable to get pod for dummy")
		return ctrl.Result{}, err
	}
	// Pod status in dummy status
	dummyPatchBase = client.MergeFrom(dummyInstance.DeepCopy())
	dummyInstance.Status.PodStatus = string(pod.Status.Phase)
	err = r.Status().Patch(ctx, dummyInstance, dummyPatchBase)
	if err != nil {
		log.Error(err, "unable to patch Dummy object's status")
		return ctrl.Result{}, err
	}

	log.Info("Current dummy status: " + dummyInstance.Status.PodStatus)
	log.Info(fmt.Sprintf("Reconciled %s", dummyInstance.Name))
	return ctrl.Result{}, nil
}

func getPodManifest(owner *interviewcomv1alpha1.Dummy, scheme *runtime.Scheme) *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      owner.Name,
			Namespace: owner.Namespace,
		},
		Spec: corev1.PodSpec{
			Containers: getPodContainers(),
		},
	}
	// Set owner reference, control the pod's lifecycle
	ctrl.SetControllerReference(owner, pod, scheme)

	return pod
}

func getPodContainers() []corev1.Container {
	return []corev1.Container{
		{
			Name:  "nginx",
			Image: "nginx",
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *DummyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&interviewcomv1alpha1.Dummy{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
