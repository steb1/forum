docker image build -f Dockerfile -t forum .

docker images

docker container run -p 8080 --detach --name container-test forum

docker pa -a

docker run -p 8080:8080 forum