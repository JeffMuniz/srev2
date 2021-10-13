sudo wget https://github.com/kubernetes/minikube/releases/download/v0.30.0/minikube-linux-amd64 -O /usr/local/bin/minikube
sudo chmod +x /usr/local/bin/minikube

sudo wget https://storage.googleapis.com/kubernetes-release/release/v1.12.2/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl
sudo chmod +x /usr/local/bin/kubectl

minikube start --memory 4096 --cpus 4

minikube status

minikube addons enable ingress

minikube addons enable dashboard
echo "OK"
exit 0
minikube dashboard


cd /data/work/git/elasticsearch-kubed/clusters/elastic

for proc in `ls -d ?_*`; do kubectl apply -f $proc; sleep 180; done
