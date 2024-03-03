DOCKER_USERNAME ?= jaysoftware
APPLICATION_NAME ?= klubbot
 
builddocker:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME} .

push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}

start:
	docker run -p 8080:8080 -d --name ${APPLICATION_NAME} ${DOCKER_USERNAME}/${APPLICATION_NAME}