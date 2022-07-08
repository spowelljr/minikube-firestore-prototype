# minikube-firestore-prototype

## How to run on minikube
```
$ minikube start --driver docker
$ minikube ssh
$ curl -LO https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-392.0.0-linux-x86_64.tar.gz
$ tar -xf google-cloud-cli-392.0.0-linux-x86_64.tar.gz
$ ./google-cloud-sdk/install.sh
$ . .bashrc
$ gcloud components install beta
$ sudo mkdir -p /usr/share/man/man1
$ sudo apt-get update
$ sudo apt-get install default-jre
$ gcloud components install cloud-firestore-emulator
$ gcloud beta emulators firestore start --host-port=0.0.0.0:8086 --project=test

# in another terminal run the following
$ git clone https://github.com/spowelljr/minikube-firestore-prototype.git
$ cd minikube-firestore-prototype
$ eval $(minikube docker-env)
$ docker build . -t firestore
$ kubectl apply -f firestore.yml
```
