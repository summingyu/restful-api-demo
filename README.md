#### 使用说明

```
cmd: CLI
apps: 业务
    host: 主机的增删改查的业务模型
      host.go: 定义业务模型, struct 和 领域接口(关于这个业务对象处理的接口)
```

本地启mysql

`docker run -itd --name mysqlWSL -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0.27`
