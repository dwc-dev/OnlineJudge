# OnlineJudge Docker 部署目录

- `docker-compose.yml`: 一键编排入口
- `.env`: 环境变量
- `backend/`: Go 服务通用构建镜像
- `backend-config/`: Go 服务配置模板
- `frontend/`: 前端构建和 Nginx 配置
- `ai/`: AI 服务容器化覆盖代码
- `mysql/init.sql`: 数据库初始化脚本
- `minio/`: MinIO 自定义镜像与启动脚本
- `scripts/build-sandbox-images.sh`: 判题沙箱镜像构建脚本

## 使用

```bash
docker compose -f deploy/docker-compose.yml --env-file deploy/.env up -d --build
```
