docker-build-server:
	docker build -t rizalanggoro/socket-server-sister-tugas-5 -f server.Dockerfile .

docker-push-server:
	docker push rizalanggoro/socket-server-sister-tugas-5:latest

docker-build-client:
	docker build -t rizalanggoro/socket-client-sister-tugas-5 -f client.Dockerfile .

docker-push-client:
	docker push rizalanggoro/socket-client-sister-tugas-5:latest

.PHONY: docker-build-server docker-build-client docker-push-server docker-push-client