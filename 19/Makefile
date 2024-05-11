all: create-rootfs docker-run

create-rootfs:
	docker pull alpine
	docker run --name alpine alpine
	docker export alpine > alpine.tar
	docker rm alpine

docker-run:
	docker run -it --rm -v `pwd`:/go-system-programming/19 --name test-container golang:1.21-bullseye
