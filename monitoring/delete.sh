helm uninstall prometheus -n prometheus

kubectl delete -f dcgm-exporter.yaml
kubectl delete -f service-monitor.yaml

