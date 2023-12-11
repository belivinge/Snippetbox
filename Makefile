IMAGE_NAME = docker_image

DOCKERFILE = Dockerfile 

build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

run:
	docker run -d -p 3306:3306 --name mysql-container $(IMAGE_NAME):latest

clean:
	docker system prune -af
