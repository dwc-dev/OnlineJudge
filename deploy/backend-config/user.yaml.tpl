Name: microservice.user
ListenOn: 0.0.0.0:6001

Etcd:
  Hosts:
    - "etcd:2379"
  Key: "microservice.user"

Mysql:
  DataSource: "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(mysql:3306)/online_judge?charset=utf8mb4&parseTime=True&loc=Local"

JWT:
  SecretKey: "${JWT_SECRET}"

Minio:
  Endpoint: "minio:9000"
  Address: "localhost:80"
  UseSSL: ${MINIO_USE_SSL}
  AccessKey: "${MINIO_ROOT_USER}"
  SecretKey: "${MINIO_ROOT_PASSWORD}"
  Bucket: "online-judge"
  Avatar:
    Prefix: "avatar"
    Default: "default_avatar.jpg"

MyRedis:
  Host: "redis:6379"
  Password: "${REDIS_PASSWORD}"
  DB: 0
