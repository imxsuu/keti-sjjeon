## k8s master ip addresses
controlplane="10.0.2.131
              10.0.2.132
	      10.0.2.133"

## k8s worker node ip addresses
worker="10.0.2.134 
	10.0.2.135
	10.0.2.136"

## Firstly, delete ketiops-master-1's k8s cluster
kubeadm reset -f
systemctl stop kubelet
systemctl stop docker
rm -rf /var/lib/cni/
rm -rf /var/lib/kubelet/*
rm -rf /etc/cni/
rm -rf /etc/kubernetes
rm -rf $HOME/.kube
ifconfig cni0 down
ifconfig flannel.1 down
ifconfig docker0 down
systemctl start kubelet
systemctl start docker
docker login 10.0.0.255:5000 -u admin -p Ketilinux11
ip link delete cni0

for ipaddr in $controlplane
do
	echo -e "\n"
	echo -e "#####################################################################################################################"
	echo "Deleting $ipaddr k8s cluster......."
	echo -e "#####################################################################################################################"
	echo -e "\n"
	ssh root@$ipaddr "kubeadm reset -f; systemctl stop kubelet; systemctl stop docker; rm -rf /var/lib/cni/; rm -rf /var/lib/kubelet/*; rm -rf /etc/cni/; rm -rf /etc/kubernetes; rm -rf $HOME/.kube; ifconfig cni0 down; ifconfig flannel.1 down; ifconfig docker0 down; systemctl start kubelet; systemctl start docker; docker login 10.0.0.255:5000 -u admin -p Ketilinux11; ip link delete cni0; exit"
done

for ipaddr in $worker
do
	echo -e "\n"
        echo -e "#####################################################################################################################"
        echo "Deleting $ipaddr k8s cluster......."
        echo -e "#####################################################################################################################"
        echo -e "\n"
	ssh root@$ipaddr "kubeadm reset -f; systemctl stop kubelet; systemctl stop docker; rm -rf /var/lib/cni/; rm -rf /var/lib/kubelet/*; rm -rf /etc/cni/; rm -rf /etc/kubernetes; rm -rf $HOME/.kube; ifconfig cni0 down; ifconfig flannel.1 down; ifconfig docker0 down; systemctl start kubelet; systemctl start docker; docker login 10.0.0.255:5000 -u admin -p Ketilinux11; ip link delete cni0; exit"
done
