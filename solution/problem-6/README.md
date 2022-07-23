1. install helm in system
		# add prometheus Helm repo
		helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

		# add grafana Helm repo
		helm repo add grafana https://grafana.github.io/helm-charts
2. install prometheus 
		kubectl create namespace monitoring

		helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

		helm install prometheus prometheus-community/prometheus \
		    --namespace monitoring \
		    --set alertmanager.persistentVolume.storageClass="gp2" \
		    --set server.persistentVolume.storageClass="gp2"

3.check if prometheus component deployed or not
		kubectl get all -n monitoring

4. port forward it to 8080 to get access
		kubectl port-forward -n monitoring deploy/prometheus-server 8080:9090

5. check in browser if its working or not

Deploying grafana

1. make grafana directory
		mkdir ${HOME}/environment/grafana
2. copy granana.yaml using this command
		mkdir ${HOME}/environment/grafana
		cp grafana.yaml ${HOME}/environment/grafana/grafana.yaml
3. run following command
		helm install grafana grafana/grafana \
    		--namespace monitoring \
    		--set persistence.storageClassName="gp2" \
    		--set persistence.enabled=true \
    		--set adminPassword='EKS!sAWSome' \
    		--values ${HOME}/environment/grafana/grafana.yaml \
    		--set service.type=LoadBalancer

4. check if grafana component deployed or not
		kubectl get all -n grafana
5.get grafana url using 
		export ELB=$(kubectl get svc -n grafana grafana -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
		echo "http://$ELB"
6. get password to login using
		kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
7. login to grafana dashboard and add 6417 dashboard to load prometheus metric into visual form.
