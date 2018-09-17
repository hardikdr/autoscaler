/*
Copyright (c) 2017 SAP SE or an SAP affiliate company. All rights reserved.

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

// Package controller is used to provide the core functionalities of machine-controller-manager
package controller

import (
	"github.com/gardener/machine-controller-manager/pkg/apis/cluster/v1alpha1"
	"k8s.io2/apimachinery/pkg/labels"
)

// existsMachineClassForSecret checks for any machineClass
// referring to the passed secret object
func (c *controller) existsMachineClassForSecret(name string) (bool, error) {
	machineClasses, err := c.findMachineClassForSecret(name)
	if err != nil {
		return false, err
	}

	if len(machineClasses) == 0 {
		return false, nil
	}
	return true, nil
}

// findAWSClassForSecret returns the set of
// AWSMachineClasses referring to the passed secret
func (c *controller) findMachineClassForSecret(name string) ([]*v1alpha1.MachineClass, error) {
	machineClasses, err := c.machineClassLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}
	var filtered []*v1alpha1.MachineClass
	for _, machineClass := range machineClasses {
		if machineClass.SecretRef.Name == name {
			filtered = append(filtered, machineClass)
		}
	}
	return filtered, nil
}