#!/bin/sh

set -eu

docker build -t oj-c /workspace/backend/code-sandbox/internal/strategies/c
docker build -t oj-cpp /workspace/backend/code-sandbox/internal/strategies/cpp
docker build -t oj-java /workspace/backend/code-sandbox/internal/strategies/java
docker build -t oj-python /workspace/backend/code-sandbox/internal/strategies/python
docker build -t oj-golang /workspace/backend/code-sandbox/internal/strategies/golang
docker build -t oj-rust /workspace/backend/code-sandbox/internal/strategies/rust
