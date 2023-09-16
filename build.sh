# 安装rsrc
go install -v github.com/akavel/rsrc@latest
# 打包manifest文件和应用图标
rsrc -manifest test.manifest -ico favicon.ico -o rsrc.syso
# 编译windowsgui项目，禁用命令行窗口并压制不必要的调试信息
go build -ldflags="-H windowsgui -w -s"