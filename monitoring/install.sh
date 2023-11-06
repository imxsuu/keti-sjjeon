kubectl create -f dcgm-exporter.yaml
kubectl create -f service-monitor.yaml

helm install prometheus prometheus-community/kube-prometheus-stack --create-namespace --namespace prometheus --set grafana.adminPassword='admin' -f kube-prometheus-stack.values

