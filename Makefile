docker-version:
	@echo "Current Version of Docker:"
	@docker --version

kube-version:
	@echo "Current Version of K8s:"
	
	@kubectl version

kube-run:
	@echo "Running K8s Deployments..."
	
	@kubectl apply -f .\K8s\platforms-depl.yml

	@echo "Running K8s Services..."
	
	@kubectl apply -f .\K8s\platforms-np-srv.yml

kube-run:
	@echo "Stopping K8s Deployments..."
	
	@kubectl delete deployment platforms-depl

	@echo "Stopping K8s Services..."
	
	@kubectl delete service platformnpservice-srv

kube-get:
	@echo "Getting Deployments..."

	@kubectl get deployments

	@echo "Getting Pods..."
	
	@kubectl get pods

	@echo "Getting Services..."
	
	@kubectl get services

# https-clean-trust
dev-certs:
	@echo "Cleaning..."
	@dotnet dev-certs https --clean
	@dotnet dev-certs https --trust
