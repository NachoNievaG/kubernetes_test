to run

```sh
docker run -d -publish 8080:8080 catfact
```

For kubernetes:
```sh
minikube start --driver docker

# clean all
minikube delete 

# use docker env for minikube, to allow local docker images to be used
eval $(minikube docker-env)

# apply files
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml
kubectl apply -f hpa.yaml #currently not working :( 

# start service
minikube service catfact

# monitor using k9s
k9s

# test
curl $URL/api/v1/catfact
{"fact":"Cats have supersonic hearing","length":28}
```
