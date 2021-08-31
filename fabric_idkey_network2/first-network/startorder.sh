docker rm -f $(docker ps -aq)
docker volume prune
docker network prune
docker ps -a
docker volume ls
docker-compose -f orderer.yaml up -d
docker ps -a
docker volume ls
