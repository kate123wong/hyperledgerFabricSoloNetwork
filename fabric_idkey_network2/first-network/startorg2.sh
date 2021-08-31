
docker rm -f $(docker ps -aq)
docker volume prune
docker network prune
export $(cat .env_key)
docker ps -a
docker volume ls
docker-compose -f peer0.org2.yaml -f peer1.org2.yaml up -d
docker ps -a
docker volume ls
