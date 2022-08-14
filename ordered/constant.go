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

package ordered

import (
	"github.com/luwuioc/container/integer"
)

// Bean priority (Priority) definition
// The smaller the value, the higher the priority
//
// Bean 优先级(Priority) 定义
// 数值越小, 优先级越高
const (
	HighPriority    = integer.MinValue  // 高优先级
	LowPriority     = integer.MaxValue  // 低优先级
	DefaultPriority = integer.ZeroValue // 默认优先级
)
