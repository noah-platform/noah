dev.up:
	$(MAKE) registry.up
	$(MAKE) cluster.up

dev.down:
	$(MAKE) cluster.down
	$(MAKE) registry.down

dev.apply:
	$(MAKE) cluster.apply

registry.up:
	k3d registry create noah.localhost --port 5500

registry.down:
	k3d registry delete noah.localhost

cluster.up:
	k3d cluster create noah -p "80:80@loadbalancer" --k3s-arg "--disable=traefik@server:0" --registry-use "k3d-noah.localhost:5500"

	$(MAKE) cluster.setup
	$(MAKE) cluster.apply

cluster.setup:
	helm repo add datawire https://app.getambassador.io
	helm repo update
	
	kubectl create namespace emissary
	kubectl apply -f https://app.getambassador.io/yaml/emissary/3.9.1/emissary-crds.yaml
	kubectl wait --for=condition=available --timeout=300s -n emissary-system deployment emissary-apiext
	
	helm install emissary-ingress --namespace emissary --set replicaCount=1 datawire/emissary-ingress
	kubectl scale deployment emissary-apiext -n emissary-system --replicas=1
	kubectl wait --for=condition=available --timeout=300s -n emissary deployment -l app.kubernetes.io/instance=emissary-ingress

	kubectl apply -f https://app.getambassador.io/yaml/v2-docs/3.9.1/quickstart/qotm.yam

cluster.apply:
	kubectl apply -f ./manifests/

cluster.down:
	k3d cluster delete noah
