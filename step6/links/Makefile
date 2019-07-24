all: web router repository

web: web_image
	docker push campoy/web-grpc

web_image:
	docker build -f web.Dockerfile -t campoy/web-grpc .

router: router_image
	docker push campoy/router-grpc

router_image:
	docker build -f router.Dockerfile -t campoy/router-grpc .

repository: repository_image
	docker push campoy/repository-grpc

repository_image:
	docker build -f repository.Dockerfile -t campoy/repository-grpc .

kubenuke: nuke_deployments nuke_services nuke_pods

nuke_deployments:
	kubectl delete --all deployments

nuke_services:
	kubectl delete --all services

nuke_pods:
	kubectl delete --all pods

deploy:
	kubectl create -f manifest.yaml

redeploy: kubenuke deploy
