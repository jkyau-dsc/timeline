# K8s services

### Summary

In this chapter, we just go over a very basic example using k8s services.

We have a server written in Go that runs 2 different versions of an API.

Services rely upon the label and selector constructs

We also get familiar with 2 useful tools: `kubectl port-forward` and `kubectl get endpoints`.

By default, kubectl operates within the default namespace. This is fine for testing purposes but feel free to create a new namespace and work within that if you do so choose.

### Workshop

#### Ensure you have setup kuebctl and minikube properly

_I highly recommend disconnecting from the VPN_

Please ensure that kubectl is using the correct context

```
$ kubectl config current-context
minikube
```

If the above command doesn't return minikube, please walkthrough the minikube installation instructions one more time to make sure everything was set up properly or contact infra for assistance.

#### Allow minikube to use your local docker images

```
$ eval $(minikube docker-env)
```

_Note that you will need to run the above command whenever you interact with a new shell for the changes to be applied_

#### Build the docker image

```
code/k8s-for-app-devs/workshops/session1 $ docker build -t coolapp .
```


##### Apply the k8s manifests

```
code/k8s-for-app-devs/workshops/session1 $ kubectl apply -f k8s/coolappv1-deployment.yaml
code/k8s-for-app-devs/workshops/session1 $ kubectl apply -f k8s/coolappv2-deployment.yaml
code/k8s-for-app-devs/workshops/session1 $ kubectl apply -f k8s/service.yaml
```


You can inspect the various resources we have created using kubectl to find out more

```
$ kubectl get deployment coolapp-v1
$ kubectl get deployment coolapp-v2
$ kubectl get service coolapp

# can try using the following tricks to get more information about a resource (works with deployments, services, and more)
$ kubectl get deployment coolapp-v1 -o json # also supports yaml
$ kubectl describe deployment coolapp-v1
```

##### Port-forwarding

```
$ kubectl port-forward svc/coolapp 8080:8080

# can open up a new shell or a browser to test
$ curl localhost:8080
```

#### Kubectl edit
Stop the port-forwarding. It's time to modify the service so that we can start serving the superior version 2 of this application.

At the beginning of this workshop, the selector looks like

```
selector:
  app: coolapp
  version: v1
```

Now we want to change it to
```
selector:
  app: coolapp
  version: v2
```

We have 2 options to choose from to make this change.

##### 1. Kubectl edit

Kubectl edit allows a user to modify a resource directly using a specified text editor. The change will be applied after saving and exiting the editor.

```
$ kubectl edit service coolapp
```

By default, this uses the vi text editor. This can be changed by reading the [docs](https://github.com/fabric8io/kansible/blob/master/vendor/k8s.io/kubernetes/docs/user-guide/kubectl/kubectl_edit.md)

##### 2. Kubectl apply

You also have the option of modifying the `service.yaml` and updating the selector in that file.

```
code/k8s-for-app-devs/workshops/session1 $ kubectl apply -f k8s/service.yaml
```


#### Kubernetes endpoints

Endpoints are one of the lower level abstractions that support Kubernetes services. As you make changes to the service selector, I recommend running

```
kubectl get endpoints
```

in order to see how the changes manifests. Try the following
##### 1. Modify the service to point at v1 coolapp (should see 3 ip addresses)
##### 2. Modify the service to point at v2 coolapp (should see 3 different ip addresses)
##### 3. Modfiy the service to point at neither v1 coolapp nor v2 coolapp (should see 0 ip addresses)

### References
[Slides](https://docs.google.com/presentation/d/1mVbpLJBFMgk3w6gSJ9m2OsGTGC-RQATaKYEf_az_s9s/edit#slide=id.gc6f919934_0_0)
[Official Kubernetes Service Documentation](https://kubernetes.io/docs/concepts/services-networking/service/)
