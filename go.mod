module fate-cloud-agent

go 1.13

require (
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gosuri/uitable v0.0.1
	github.com/pkg/errors v0.8.1
	github.com/spf13/pflag v1.0.5
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/urfave/cli/v2 v2.1.1
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	helm.sh/helm/v3 v3.0.2
	k8s.io/klog v1.0.0
)

replace (
	// github.com/Azure/go-autorest/autorest has different versions for the Go
	// modules than it does for releases on the repository. Note the correct
	// version when updating.
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

	// Kubernetes imports github.com/miekg/dns at a newer version but it is used
	// by a package Helm does not need. Go modules resolves all packages rather
	// than just those in use (like Glide and dep do). This sets the version
	// to the one oras needs. If oras is updated the version should be updated
	// as well.
	github.com/miekg/dns => github.com/miekg/dns v0.0.0-20181005163659-0d29b283ac0f
	gopkg.in/inf.v0 v0.9.1 => github.com/go-inf/inf v0.9.1
	gopkg.in/square/go-jose.v2 v2.3.0 => github.com/square/go-jose v2.3.0+incompatible

	rsc.io/letsencrypt => github.com/dmcgowan/letsencrypt v0.0.0-20160928181947-1847a81d2087
)
