# fate-cloud-agent

config


init service pod
```bash
$ kubefate init -f ./config.yaml
```
kubefate service pod status 
```bash
$ kubefate status
```

*Service pod must run successfully*

cluster deploy
```bash
$ kubefate install <clusterName> -f ./cluster.yaml
```

cluster upgrade 
```bash
$ kubefate upgrade <clusterName> -f ./cluster.yaml
```

cluster delete 
```bash
$ kubefate upgrade <clusterName>
```

cluster list 
```bash
$ kubefate list
```

cluster info 
```bash
$ kubefate describe <clusterName>
```

job info
```bash
$ kubefate job <jobUUID>
```

job list
```bash
$ kubefate job list
```

job delete
```bash
$ kubefate job delete <jobUUID>
```

user info
```bash
$ kubefate user <userUUID>
```

user list
```bash
$ kubefate user list
```

user delete
```bash
$ kubefate user delete <userUUID>
```