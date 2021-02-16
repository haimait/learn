# gf docker -t gf-demos
FROM golang

LABEL maintainer="whm2416@qq.com"


# 设置固定的项目路径
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn

# 设置环境变量
#ENV WORKDIR /go/src/gin04

# 设置工作目录
WORKDIR /go/src/gin04
# 复制当前文件所有内容到容器的工作环境里
COPY . .

#下载vendor依赖包
#RUN export GO111MODULE=on
#RUN export GOPROXY="https://goproxy.cn"

#RUN go mod vendor
RUN go mod tidy
#打包应用可执行文件
#RUN go build -o gin04
RUN go build -o gin04 main.go
#并设置执行权限
#RUN chmod +x $WORKDIR/gin04.go
#暴露端口
EXPOSE 9999

###############################################################################
#                                   START
###############################################################################

# 运行服务
#CMD ["/bin/bash","-c","/go/src/gin04/gin04"]
CMD ["bash","-c","/go/src/gin04/gin04"]
#CMD ./gin04
#最终运行docker的命令
#ENTRYPOINT  ["/go/src/gin04/gin04"]

## 命令行执行docker build -t wanghaima/gin04-web:1.0 . -f Dockerfile

## 运行容器
## docker run --name gin04-web -p 9999:9999 -d wanghaima/gin04-web:1.0