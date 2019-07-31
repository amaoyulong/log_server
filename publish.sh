#/bin/sh

cd $(dirname $0)

#打包

docker run -it -d --rm --name bide2-instance -v /Users/yulongli/Code/Go/src/bide2:/go/src/bide2 -w /go/src/bide2 beego-image
docker exec -it bide2-instance /bin/bash
../../bin/bee pack -be GOOS=linux 
scp -r /Users/yulongli/Code/Go/src/bide2/bide2.tar.gz q99:/usr/local/app/bide2
ssh q99 "tar -xzvf /usr/local/app/bide2/bide2.tar.gz -C /usr/local/app/bide2; docker stop bide2-instance; docker run -it --rm -d --name bide2-instance -p 8080:8080 -v /usr/local/app/bide2:/app/bide2  -v /usr/local/app/bide2/static/upload:/app/bide2/static/upload  exec-image:v1 /app/bide2/bide2; docker ps"



scp -r /Users/yulongli/Code/Go/src/static_bide2/*   q99:/usr/local/app/static_bide2/html

scp -r /Users/yulongli/Code/Go/src/bide2/views/admin.tpl   q99:/usr/local/app/bide2/views




----------------
scp  -r /Users/yulongli/Code/Go/src/www_bide2/*   q99:/usr/local/app/www_bide2/