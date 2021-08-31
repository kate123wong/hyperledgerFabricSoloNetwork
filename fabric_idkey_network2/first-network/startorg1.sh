
docker rm -f $(docker ps -aq)
docker volume prune
docker network prune
docker ps -a
docker volume ls
export $(cat .env_key)
docker-compose -f peer0.org1.yaml -f peer1.org1.yaml up -d
docker ps -a
docker volume ls

