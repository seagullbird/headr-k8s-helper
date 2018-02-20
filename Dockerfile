FROM alpine

COPY k8s-helper /

ENTRYPOINT /k8s-helper