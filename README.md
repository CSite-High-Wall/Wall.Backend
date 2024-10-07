## Wall.Backend
> 工地高墙后端仓库 

### 部分实现说明

用户密码在数据库中已加密  
用户鉴权方式采用 JWT，同时使用 Cookies 存储 AccessToken，  
AccessToken 具有一定有效期，且可以被立即吊销  
**不允许用户同时多台设备登录**

数据库使用 MySQL，并部署在远程服务器

### 后端接口

详见 [router.go](https://github.com/CSite-High-Wall/Wall.Backend/blob/main/router.go)

### 运行调试

`conf/config.yaml` 该文件为程序配置文件，  
由于其中包含了数据库密钥等信息，本仓库并不提供该文件  
此处仅提供文件模板  

``` yaml
mysql:
  database_name: ""
  host: ""
  port: 3306
  user: ""
  password: ""

server:
  staticFs_schema: "http"
  staticFs_host: "localhost:8000"
# 上面两行将影响返回的静态资源的 url
```

### 部署
将项目编译到 Linux 可执行文件 `详见 build.bat` ，后上传服务器，在服务器使用 systemd 守护进程，运行服务

### 贡献者
<a href="https://github.com/CSite-High-Wall/Wall.Backend/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=CSite-High-Wall/Wall.Backend" />
</a>

  
  
