Name: gateway-api
Host: 0.0.0.0
Port: 8888

CORSOrigins:
  - "http://localhost"
  - "http://127.0.0.1"
  - "http://localhost:5173"
  - "http://127.0.0.1:5173"

UserRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: microservice.user
  NonBlock: true

QuestionRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: microservice.question
  NonBlock: true

JudgeRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: microservice.judge
  NonBlock: true
  Timeout: 60000

CompetitionRpc:
  Etcd:
    Hosts:
      - "etcd:2379"
    Key: microservice.competition
  NonBlock: true

AIServiceURL: http://ai:6005

JWTConf:
  AccessSecret: "${JWT_SECRET}"
