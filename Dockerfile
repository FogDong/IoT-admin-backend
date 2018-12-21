#源镜像
FROM golang:latest
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR $GOPATH/src/IoT-admin-backend
COPY . $GOPATH/src/IoT-admin-backend
RUN go build .
#暴露端口
EXPOSE 9002
#最终运行docker的命令
ENTRYPOINT  ["./IoT-admin-backend"]