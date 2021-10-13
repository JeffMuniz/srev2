minikube stop; minikube delete ; rm /usr/local/bin/minikube ; rm -rf ~/.minikube minikube start --memory 5120 --cpus 4
minikube dashboard
kubectl config use-context minikube
kubectl create namespace elastic
kubectl create -f /data/work/minikube/elastic.yml -n elastic
kubectl get pods -n elastic
kubectl get service -n elastic
curl $(minikube ip):30531
kubectl create -f /data/work/minikube/kibana.yml -n elastic
kubectl get pods -n elastic
kubectl get service -n elastic


########################################
#Instalação do Minikube
# Referencia: http://alissonmachado.com.br/minikube-kubernetes-em-ambiente-de-desenvolvimento/
sudo wget https://github.com/kubernetes/minikube/releases/download/v0.30.0/minikube-linux-amd64 -O /usr/local/bin/minikube
sudo chmod +x /usr/local/bin/minikube

sudo wget https://storage.googleapis.com/kubernetes-release/release/v1.12.2/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl
sudo chmod +x /usr/local/bin/kubectl

minikube start --memory 5120 --cpus 4

minikube status

minikube addons enable ingress

minikube addons enable dashboard

minikube dashboard

#Instalacao Elasticsearch



