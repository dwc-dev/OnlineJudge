Name: microservice.competition
ListenOn: 0.0.0.0:6004
Timeout: 60000

Etcd:
  Hosts:
    - "etcd:2379"
  Key: "microservice.competition"

UserRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: "microservice.user"
  NonBlock: true

QuestionRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: "microservice.question"
  NonBlock: true

JudgeRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: "microservice.judge"
  NonBlock: true
  Timeout: 60000

Mysql:
  DataSource: "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(mysql:3306)/online_judge?charset=utf8mb4&parseTime=True&loc=Local"
