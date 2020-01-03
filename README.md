# fate-cloud-agent

start service
```bash
$ go run kubefate.go
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
