#! /bin/bash

if [ ! "$1" ]; then
  printf "缺少服务参数, 例如 ./build user \n"
  exit 0
fi

#cd ./app/$1/service
#
#if [ $? -ne 0 ]; then
#  echo "No such file or directory"
#  exit 1
#fi

docker build -t blog-$1 . --build-arg APP_RELATIVE_PATH=$1/service
docker tag blog-$1 registry.cn-shenzhen.aliyuncs.com/wangzhe0568/blog-$1:1.0.1
docker push registry.cn-shenzhen.aliyuncs.com/wangzhe0568/blog-$1:1.0.1
# 删除构建过程中生成的无用镜像
docker system prune -f


:<<!
docker run  --name blog-article --rm -itd -p 8001:8001 -p 9001:9001 -v '/mnt/d/My Project/Go/src/blog/app/article/service/configs':/data/conf blog-article
docker run  --name blog-user --rm -itd -p 8000:8000 -p 9000:9000 -v '/mnt/d/My Project/Go/src/blog/app/user/service/configs':/data/conf blog-user
!