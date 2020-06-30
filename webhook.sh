#! /bin/bash

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/software/node-v10.15.3-linux-x64/bin:/usr/local/software/go/bin

cd /var/www/html/wxyadpi

echo "更新代码"

git pull -v --all

go build -o wxyapi .

./wxyapi


