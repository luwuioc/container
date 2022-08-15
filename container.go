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

package container

// Container An IOC container root interface
type Container interface {
	// RegisterBean register bean by name (The name(ID) must be unique)
	// 通过 名称(ID)注册 Bean 实例,且名称(ID)全局唯一
	RegisterBean(name string, bean any) error
	// GetBean get bean by name(ID)
	// 通过 Bean 的名称获取实例, 通过 ok 值表征是否存在与 IOC 中
	GetBean(name string) (any, bool)
	// ContainsBean does this bean factory contain a bean definition or externally registered singleton instance with the given name?
	// 判定 IOC 容器中是否有该 Bean 的实例
	ContainsBean(name string) bool
	// IsSingleton is this bean a shared singleton? That is, will Container#GetBean always return the same instance?
	// 判断 该 Bean 在 IOC 中是否为单例 Bean, 每次调用 Container#GetBean 会返回同一个 Bean 实例
	IsSingleton(name string) bool
	// IsPrototype is this bean a prototype? That is, will Container#GetBean always return independent instances?
	// 判断 该 Bean 在 IOC 中是否为多例 Bean, 每次调用 Container#GetBean 会返回一个相互独立的 Bean 实例
	IsPrototype(name string) bool
	// GetAlias Return the aliases for the given bean name, if any.
	// 获取 该 Bean 在 IOC 中的所有别名
	GetAlias(name string) []string
}
