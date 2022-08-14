/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package definition

// BeanDefinition an IOC bean definition
//
// Refer to Spring's definition of BeanDefinition
//
// @see org.springframework.beans.factory.config.BeanDefinition
//
// IOC Bean 的定义, 参考了 Spring 对 Bean 的定义 BeanDefinition
type BeanDefinition interface {
	Scope() int                // Return the scope(Singleton or Prototype) of the bean in IOC
	Name() string              // Return the ID of the bean in IOC
	Struct() string            // Return the name of bean(struct)
	PkgPath() string           // Return the package path of bean(struct)
	Parent() string            // Return the full name of parent
	IsAutowireCandidate() bool // Return whether this bean is a candidate for getting autowired into some other bean.
	IsPrimary() bool           // Return whether this bean is a primary autowire candidate.
	FactoryBeanName() string   // Return the factory bean name, if any.
	FactoryMethodName() string // Return a factory method, if any.
	InitMethodName() string    // Return the name of the initializer method.
	DestroyMethodName() string // Return the name of the destroy method.
	Description() string       // Return a human-readable description of this bean definition.
	IsSingleton() bool         // Return whether this a Singleton, with a single, shared instance returned on all calls.
	IsPrototype() bool         // Return whether this a Prototype, with an independent instance returned for each call.
	Aliases() []string         // Return all the aliases of this bean in IOC
}
