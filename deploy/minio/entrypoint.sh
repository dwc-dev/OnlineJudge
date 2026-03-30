#!/bin/sh

set -eu

# 启动 MinIO 服务（后台）
minio server /data --console-address ":9001" &
MINIO_PID=$!

# 等待 MinIO 就绪
echo "Waiting for MinIO to start..."
until mc alias set local http://localhost:9000 "$MINIO_ROOT_USER" "$MINIO_ROOT_PASSWORD" > /dev/null 2>&1; do
  sleep 1
done

# 创建 bucket（已存在则跳过）
mc mb local/online-judge --ignore-existing

# 设置匿名可下载
mc anonymous set download local/online-judge

# 上传默认头像（已存在则跳过）
if ! mc stat local/online-judge/avatar/default_avatar.jpg > /dev/null 2>&1; then
  mc cp /tmp/default_avatar.jpg local/online-judge/avatar/default_avatar.jpg
  echo "Default avatar uploaded."
else
  echo "Default avatar already exists, skipping."
fi

echo "MinIO initialized successfully."

# 把 MinIO 进程拉回前台
wait $MINIO_PID