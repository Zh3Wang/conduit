#! /bin/bash
if [ ! "$1" ]; then
  printf "缺少表名参数, 例如 ./model.sh user \n"
  exit 0
fi

if [ ! "$2" ]; then
  printf "缺少struct参数, 例如 ./model.sh User \n"
  exit 0
fi
cd ./model
if [ ! -d $1_model ]; then
    mkdir $1_model
fi

db2struct \
--host 119.91.235.156 \
--database conduit \
--table $1 \
--package $1Model \
--struct $2 \
--password qwe254511 \
--user wangzhe \
--gorm \
--gotype \
--json \
--target $1_model/$1_model.go

goimports -w $1_model/$1_model.go