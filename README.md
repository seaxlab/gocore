# GO Core

## 目录说明

- 核心基类

## dev

````
go install
````

## dependency

`go mod graph` //可视化依赖

- github.com/stretchr/testify/assert 测试断言库

## Notice

- 禁止使用`json-iterator`，性能和原生差不多
- 尽量使用返回err，`errors.New("xxx")`
- 对原生语法扩展后缀以`x`结尾，例如`arrayx` `timex`