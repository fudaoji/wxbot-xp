## Go-Wxbot

基于`Golang`语言和`Gin`框架的个人微信机器人。

项目基于<a href="https://github.com/cixingguangming55555/wechat-bot" target="_blank">WeChat-bot</a>，用golang实现客户端的同时对外提供RESTFul Api。

## 安装

```shell
# 下载代码
git clone https://github.com/fudaoji/wxbot-xp.git
# 更新依赖
go mod download
# 复制一份生产环境配置文件修改配置文件
cp setting_dev.yaml  setting_prod.yaml
# 编译
go build main.go
# 清理无用mod引用
go mod tidy
```
无需手动安装数据表，随着程序的启动，系统会自动迁移数据表。
## 运行

```shell
# 运行测试环境
./main
# 运行生成环境
./main  -mode="prod"
```

## Thanks

<a href="https://github.com/cixingguangming55555/wechat-bot" target="_blank">WeChat-Bot</a>
