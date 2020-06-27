#! /bin/bash

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/software/node-v10.15.3-linux-x64/bin:/usr/local/software/go/bin

cd /var/www/html/wxyadpi

echo "更新代码"
git pull -v --all

name=wxyapi
echo "开始停止 $name 容器"
docker stop $name
echo "停止容器 $name 成功"

echo "开始删除 $name 容器"
docker rm $name
echo "删除 $name 容器成功"

docker images|grep none|awk '{print $3 }'|xargs docker rmi

imagesid=`docker images|grep -i $name|awk '{print $3}'`
if [ "$imagesid" == "" ];then
   echo "镜像不存在！"
else
    echo "删除镜像id $imagesid"
    docker rmi $imagesid -f
    echo "删除成功"
fi
docker build . -t wxyapi

docker run -itd --name wxyapi --link=mysql-dev:mysql-dev   --link redis-dev:redis-dev  -p 9111:9111 wxyapi
