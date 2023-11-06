helm install prometheus prometheus-community/kube-prometheus-stack --create-namespace --namespace restoring --set grafana.adminPassword='admin' -f restoring-prometheus.values
