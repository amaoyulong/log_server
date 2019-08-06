#/bin/sh

cd $(dirname $0)

#打包

docker run -it -d --rm --name log-instance -v /Users/yulongli/Code/Go/src/log_server:/go/src/log_server -w /go/src/log_server goapp-image 
docker exec -it log-instance /bin/bash
go build
scp  -r /Users/yulongli/Code/Go/src/log_server/log_server q99:/usr/local/app/log_server_1
ssh q99 "docker stop log-instance; mv /usr/local/app/log_server_1 /usr/local/app/log_server/log_server; docker run -it --rm -d --name log-instance -p 18600:18600/udp -v /usr/local/app/log_server:/app/log_server  exec-image:v1 /app/log_server/entrypoint.sh; docker ps"


