# 备注

```bash
# 构建镜像
docker build -t idogo:1.0.0.2 -f devops/docker/1002/Dockerfile .

# 查看镜像
docker images | grep idogo

# run运行
docker run --name idogo_1002 idogo:1.0.0.2

# compose部署
docker compose -f devops/docker/1002/docker-compose.yml up
docker compose -f devops/docker/1002/docker-compose.yml up -d
docker compose -f devops/docker/1002/docker-compose.yml down

# 进入容器
docker exec -it idogo_1002 sh
```
