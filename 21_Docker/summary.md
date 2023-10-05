
# (21) Docker

Container bukan sebuah virtual machine (VM), melainkan sebuah proses pada docker dengan file system yang terisolasi. Perbedaan Container dengan VM adalah :
- Container abstrak pada app layer sedangkan VM pada physical hardware.
- Container tidak menggunakan storage sebanyak VM.
- Container bisa menghandle lebih banyak app, sedangkan VM sangat lambat saat boot.

Basic Docker Syntax :
- FROM => Get Image from docker registry.
- RUN => Execute bash command when building Container.
- ENV => Set variable inside Container.
- ADD => Copy the file with some other process.
- COPY => Copy the file.
- WORKDIR => Set wirking file directory.
- ENTRYPOINT => Execute command when finish building Container.
- CMD => Execute command but can overwrite.

How to build and Store Your Own Container :
1. Preparation

    `FROM golang:1.16-alpine`

    `WORKDIR /app`

    `COPY go.mod ./`

    `COPY go.sum ./`

    `RUN go mod download`

    `COPY *.go ./`

    `RUN go build -o /docker-gs-ping`

    `EXPOSE 8080`

    `CMD [ "/docker-gs-ping" ]`
2. Build Container

    `docker build -t flask-tutorial:latest .`
3. Login to your container registry

    `docker login --username=yourusername`
4. Tag and push your container image

    `docker tag 557bec4698b6`

    `docker push ian18ishar/flask-tutorial`

    `docker pull ian18ishar/flask-tutorial:1.0`
    
    `docker run -d -p 5000:5000 --name flaskdemo ian18ishar/flask-tutorial:1.0`
