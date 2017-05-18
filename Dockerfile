FROM scratch

ENV USERMAN_SERVICE_PORT 8080

ENV K8S_HOST "https://master.k8s.community"
ENV K8S_TOKEN "Token is for access to k8s API"
ENV TLS_SECRET_NAME "tls-secret"
ENV DOCKER_REGISTRY_SECRET_NAME "registry-pull-secret"

EXPOSE $USERMAN_SERVICE_PORT

COPY user-manager /

CMD ["/user-manager"]