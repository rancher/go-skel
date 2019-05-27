module %PKG%

go 1.12

replace (
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
)

require (
	k8s.io/api kubernetes-1.14.1
	k8s.io/apiextensions-apiserver kubernetes-1.14.1
	k8s.io/apimachinery kubernetes-1.14.1
	k8s.io/client-go kubernetes-1.14.1
	k8s.io/code-generator kubernetes-1.14.1
)
