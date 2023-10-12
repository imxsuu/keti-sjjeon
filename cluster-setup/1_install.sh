####################################
#### Change Version as you want ####
####################################
version=1.25.6-00

apt install apt-transport-https curl -y
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
apt-get update
apt install -qy kubelet=$version kubectl=$version kubeadm=$version kubernetes-cni
