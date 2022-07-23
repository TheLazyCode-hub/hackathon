1. Create the simple golang app sample is in main.go
2. build the docker image using 
	docker build -t go-img
3. after successfull build run container using it to check if result are coming or not.
	docker run -d -p 3333:3000 --name go-test go-img
4. push the image to dockerhub
	docker build -t go-img:latest .
	docker push gauravtest/go-img:latest
5. create deployment using deployment.yaml
	kubectl create -f deployment.yaml
6. expose deployment
	kubectl expose deployment my-go-app --type=NodePort --name=go-app-svc --target-port=3000 
7. get the port exposed using
	kubectl get svc
8. go to the link using the exposed port to see the service running.
