/*
Copyright 2024 pf93.

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

package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var customapplicationlog = logf.Log.WithName("customapplication-resource")

func (r *CustomApplication) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-apps-pf93-cn-v1-customapplication,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.pf93.cn,resources=customapplications,verbs=create;update,versions=v1,name=mcustomapplication.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &CustomApplication{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CustomApplication) Default() {
	customapplicationlog.Info("default", "name", r.Name)

	if r.Spec.Deployment.Replicas == nil {
		r.Spec.Deployment.Replicas = new(int32)
		*r.Spec.Deployment.Replicas = 3
	}	
	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-pf93-cn-v1-customapplication,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.pf93.cn,resources=customapplications,verbs=create;update,versions=v1,name=vcustomapplication.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CustomApplication{}

func (r *CustomApplication) validateApplication() error {
	if *r.Spec.Deployment.Replicas > 10 {
		return fmt.Errorf("replicas too many error")
	}
	return nil
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomApplication) ValidateCreate() (admission.Warnings, error) {
	customapplicationlog.Info("validate create", "name", r.Name)

	r.validateApplication()
	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomApplication) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	customapplicationlog.Info("validate update", "name", r.Name)

	r.validateApplication()
	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CustomApplication) ValidateDelete() (admission.Warnings, error) {
	customapplicationlog.Info("validate delete", "name", r.Name)

	r.validateApplication()
	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
