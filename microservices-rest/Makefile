all: web router repository

web: web_image
	docker push campoy/web-rest

web_image:
	docker build -f web.Dockerfile -t campoy/web-rest .

router: router_image
	docker push campoy/router-rest

router_image:
	docker build -f router.Dockerfile -t campoy/router-rest .

repository: repository_image
	docker push campoy/repository-rest

repository_image:
	docker build -f repository.Dockerfile -t campoy/repository-rest .

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
