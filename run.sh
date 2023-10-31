go mod tidy
go mod vendor
./vendor/k8s.io/code-generator/generate-groups.sh client,informer zookeeper-operator/pkg/client zookeeper-operator/pkg/apis zookeeper:v1beta1 --go-header-file ./hack/boilerplate.go.txt --output-base $(pwd)/pkg/client

