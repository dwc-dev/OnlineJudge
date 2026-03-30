#!/bin/sh

set -eu

: "${CONFIG_TEMPLATE_PATH:=/app/templates/config.yaml.tpl}"
: "${CONFIG_PATH:=/app/etc/config.yaml}"

if [ ! -f "${CONFIG_TEMPLATE_PATH}" ]; then
  echo "config template not found: ${CONFIG_TEMPLATE_PATH}" >&2
  exit 1
fi

# 容器启动时用环境变量渲染配置模板。
envsubst < "${CONFIG_TEMPLATE_PATH}" > "${CONFIG_PATH}"

exec "$@"
