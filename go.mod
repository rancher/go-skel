module %PKG%

go 1.13

replace k8s.io/client-go => k8s.io/client-go v0.17.0

require (
	github.com/rancher/wrangler v0.4.1
	github.com/rancher/wrangler-api v0.4.1
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.22.2
	k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/code-generator v0.17.0
)
