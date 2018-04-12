APP?=./build/k8s-helper
GOARCH?=amd64
GOOS?=linux
PROJECT?=github.com/seagullbird/headr-k8s-helper
COMMIT?=$(shell git rev-parse --short HEAD)
IMAGE_NAME?=k8s-helper
DEV?=true

clean:
	rm -f ${APP}

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH} go build \
	-ldflags "-s -w -X ${PROJECT}/config.Dev=${DEV}" \
	-o ${APP}

container: build
	docker build -t ${IMAGE_NAME}:${COMMIT} ./build/

minikube: container
	cat k8s/deployment.yaml | gsed -E "s/\{\{(\s*)\.Commit(\s*)\}\}/$(COMMIT)/g" > tmp.yaml
	kubectl apply -f tmp.yaml
	rm -f tmp.yaml