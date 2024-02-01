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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/hypershift/api/hypershift/v1alpha1"
)

// ServicePublishingStrategyMappingApplyConfiguration represents an declarative configuration of the ServicePublishingStrategyMapping type for use
// with apply.
type ServicePublishingStrategyMappingApplyConfiguration struct {
	Service                                      *v1alpha1.ServiceType `json:"service,omitempty"`
	*ServicePublishingStrategyApplyConfiguration `json:"servicePublishingStrategy,omitempty"`
}

// ServicePublishingStrategyMappingApplyConfiguration constructs an declarative configuration of the ServicePublishingStrategyMapping type for use with
// apply.
func ServicePublishingStrategyMapping() *ServicePublishingStrategyMappingApplyConfiguration {
	return &ServicePublishingStrategyMappingApplyConfiguration{}
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *ServicePublishingStrategyMappingApplyConfiguration) WithService(value v1alpha1.ServiceType) *ServicePublishingStrategyMappingApplyConfiguration {
	b.Service = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ServicePublishingStrategyMappingApplyConfiguration) WithType(value v1alpha1.PublishingStrategyType) *ServicePublishingStrategyMappingApplyConfiguration {
	b.ensureServicePublishingStrategyApplyConfigurationExists()
	b.Type = &value
	return b
}

// WithNodePort sets the NodePort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePort field is set to the value of the last call.
func (b *ServicePublishingStrategyMappingApplyConfiguration) WithNodePort(value *NodePortPublishingStrategyApplyConfiguration) *ServicePublishingStrategyMappingApplyConfiguration {
	b.ensureServicePublishingStrategyApplyConfigurationExists()
	b.NodePort = value
	return b
}

// WithLoadBalancer sets the LoadBalancer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LoadBalancer field is set to the value of the last call.
func (b *ServicePublishingStrategyMappingApplyConfiguration) WithLoadBalancer(value *LoadBalancerPublishingStrategyApplyConfiguration) *ServicePublishingStrategyMappingApplyConfiguration {
	b.ensureServicePublishingStrategyApplyConfigurationExists()
	b.LoadBalancer = value
	return b
}

// WithRoute sets the Route field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Route field is set to the value of the last call.
func (b *ServicePublishingStrategyMappingApplyConfiguration) WithRoute(value *RoutePublishingStrategyApplyConfiguration) *ServicePublishingStrategyMappingApplyConfiguration {
	b.ensureServicePublishingStrategyApplyConfigurationExists()
	b.Route = value
	return b
}

func (b *ServicePublishingStrategyMappingApplyConfiguration) ensureServicePublishingStrategyApplyConfigurationExists() {
	if b.ServicePublishingStrategyApplyConfiguration == nil {
		b.ServicePublishingStrategyApplyConfiguration = &ServicePublishingStrategyApplyConfiguration{}
	}
}
