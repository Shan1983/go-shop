docker-build:
	docker build -t golangshop .

docker-run:
	docker run -p 8080:8080 -tid golangshop

docker-container-stop:
	docker container stop 2fcfe210c7e5

docker-container-list:
	docker container ls

docker-up:
	docker-compose up --build server


docker: docker-container-stop docker-build docker-run docker-container-list

.PHONEY: docker-build docker-run docker-up docker-compose docker