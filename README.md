# fate-cloud-agent



```bash
git clone https://gitlab.eng.vmware.com/fate/fate-cloud-agent.git

cd ./fate-cloud-agent
```

deploy 
```bash
$ kubectl apply ./rbac-config.yaml
$ kubectl apply ./kubefate.yaml

$ kubectl get all,ingress -n kube-fate
```

*Service pod must run successfully*

cluster deploy
```bash
$ kubefate install -n <namespaces> -f ./cluster.yaml <clusterName>
```

cluster upgrade 
```bash
$ kubefate upgrade -n <namespaces> -f ./cluster.yaml <clusterName>
```

cluster delete 
```bash
$ kubefate delete <clusterId>
```

cluster list 
```bash
$ kubefate cluster list
```

cluster info 
```bash
$ kubefate cluster describe <clusterId>
```

job info
```bash
$ kubefate job describe <jobUUID>
```

job list
```bash
$ kubefate job list
```

job delete
```bash
$ kubefate job delete <jobUUID>
```
