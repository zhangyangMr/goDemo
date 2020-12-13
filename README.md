# go_demo

go 相关功能测试代码

## 目录规范

每一个go功能类型的demo单独创建一个目录，对于复杂的代码，需要再单独添加readme.md

## GitCommit信息规范

提交信息格式：`type(scope): info`

`type`用于说明 commit 的类别，只允许使用下面7个标识

- feat：新功能（feature）
- fix：修补bug
- docs：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试
- chore：构建过程或辅助工具的变动

`scope`用于表示改动会影响的范围，根据具体项目划分的模块而定

`info`用简单的语句说明本次提交的内容，中英文不限

## 示例

```
feat(web): + ZF-2729 完成启动网络功能
feat(*): + 框架搭建
docs(api) : + 添加文档
```