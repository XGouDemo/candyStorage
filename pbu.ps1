cd  C:\Users\D058009\go\candyStorage
docker rm -vf $(docker ps -a -q)
docker container prune -f
docker-compose build
docker-compose up