#！/bin/bash

# 获取脚本的执行目录
BasePath=$(cd `dirname $0`; pwd)
ProtoPath=$(cd `dirname $0`; pwd)

## 向外输出
echo "[  工作目录  ]:  ${BasePath}"
echo "[ ProtoPath ]:  $ProtoPath"

# 进入工作目录
cd $BasePath

protoc --go_out=plugins=grpc:. *.proto
protoc --doc_out=. --doc_opt=html,document.html *.proto
protoc --doc_out=. --doc_opt=markdown,document.md *.proto

echo "[  完成  ]"
