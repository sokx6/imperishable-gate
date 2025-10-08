#!/bin/bash
set -e

echo "正在整理依赖..."
go mod tidy

echo "正在编译服务端..."
go build -o bin/gate-server ./cmd/gate-server

echo "正在编译客户端..."
go build -o bin/gate ./cmd/gate

echo "✅ 编译完成！"
echo ""
echo "启动服务器: ./bin/gate-server start"
echo "使用客户端: ./bin/gate --help"
