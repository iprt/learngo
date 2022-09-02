# 测试

## 代码覆盖率测试

```shell

go test -coverprofile=c.out

```

```shell
# 可以测试一下，看看输出什么东西
go tool cover 
```

```shell
go tool cover -html=c.out
```

## 性能测试

```shell
go test -bench .
```