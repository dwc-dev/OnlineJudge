Name: microservice.question
ListenOn: 0.0.0.0:6002

Etcd:
  Hosts:
    - "etcd:2379"
  Key: "microservice.question"

Mysql:
  DataSource: "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(mysql:3306)/online_judge?charset=utf8mb4&parseTime=True&loc=Local"
