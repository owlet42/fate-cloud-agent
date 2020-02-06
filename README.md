# fate-cloud-agent

start service
```bash
$ go run kubefate.go service
```

deploy
```bash
http://127.0.0.1:8080/deploy?name=fate-10000&namespace=fate-10000&chart=E:\machenlong\AI\github\owlet42\KubeFATE\k8s-deploy\fate-10000
```
list
```bash
http://127.0.0.1:8080/list?namespace=allnamespaces
```
delete
```bash
http://127.0.0.1:8080/delete?name=fate-10000&namespace=fate-10000
```


```bash
git clone https://gitlab.eng.vmware.com/fate/fate-cloud-agent.git
cd fate-cloud-agent
docker build -t kubefate .
docker run --rm -d -p 8080:8080 --name kubefate -v ~/.kube/:/root/.kube/ -v ~/github/KubeFATE/k8s-deploy/:/workdir/k8s-deploy/ kubefate
```
kubernetes 和 chart 文件通过mount进镜像