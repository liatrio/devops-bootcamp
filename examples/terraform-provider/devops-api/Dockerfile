from ubuntu:latest

run apt-get -y update && apt-get -y install golang git

copy devops-resources /devops-resources

run mkdir devops-api
workdir devops-api

copy devops-api .

run go build

expose 8080

entrypoint [ "./devops-api" ]
