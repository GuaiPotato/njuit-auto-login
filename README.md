# Univer Login

## 项目简介
Univer Login 是一个校园网认证登录工具，用户只需运行 `univer-login.exe` 文件，系统会根据配置文件自动完成校园网的认证登录。

## 项目结构
项目根目录包含以下文件：
- `config.yaml`：配置文件，存储登录所需的账号、密码和相关网络信息。
- `go.mod` 和 `go.sum`：Go 项目的模块管理文件（无需修改）。
- `main.go`：主程序代码（无需修改）。
- `univer-login.exe`：编译后的可执行文件，用户运行此文件即可完成登录。
## 配置说明

### `config.yaml`
```yaml
# 校园网认证url
url: "http://10.255.200.11/api/portal/v1/login"
# 校园网账号
username: "username"
# 校园网密码
password: "password"
# 校园网： "default", 中国移动： "cmcc", 中国电信： "telecom"
domain: "cmcc"
