## k8s master node's IP
## please edit master ip address
masterIP=10.0.2.131

swapoff -a
kubeadm init --control-plane-endpoint "$masterIP:6443" --upload-certs --pod-network-cidr=192.168.0.0/16 --apiserver-advertise-address=$masterIP --ignore-preflight-errors=NumCPU --v=2

mkdir -p $HOME/.kube
cp /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config

sysctl net.bridge.bridge-nf-call-iptables=1
sysctl net.bridge.bridge-nf-call-ip6tables=1

kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/calico.yaml 
