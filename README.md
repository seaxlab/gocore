# GO Core

## 目录说明

- 核心基类

## dev

````
go install
````

## dependency

- github.com/stretchr/testify/assert
- github.com/jinzhu/copier

## Notice

- 禁止使用`json-iterator`，性能和原生差不多
- 尽量使用返回err，`errors.New("xxx")`