# `container` 容器设计规范
> `container` 是一个在 `GO` 语言中实现 `IOC` 容器的一个自定义规范

[English](./README.md) | 中文



- `Bean` 的定义
  - `BeanDefinition`
- `Bean` 的排序
  - `Ordered`
    - `HighPriority` 高优先级
    - `LowPriority` 低优先级
    - `DefaultPriority` 默认优先级
  - `PriorityOrdered`
    - 高优先级
- `Bean` 作用域
  - `Singleton` 单例
  - `Prototype` 多例
- `Container` 定义
  - `RegisterBean`
    - 注册 `Bean`
  - `GetBean`
    - 获取 `Bean`
  - `ContainsBean`
    - 容器是否包含 `Bean`
  - `IsSingleton`
    - 是否为 单例 `Bean`
  - `IsPrototype`
    - 是否为 多例 `Bean`
  - `GetAlias`
    - 获取 `Bean`在/ 别名列表