## 打包项目

    go build -ldflags="-s -w" -o sshpass.exe

## 生成图标资源文件

    windres -o main.syso main.rc
