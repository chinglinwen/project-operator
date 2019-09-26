FROM harbor.haodai.net/base/alpine:3.7cgo
WORKDIR /app

MAINTAINER wenzhenglin(http://g.haodai.net/wenzhenglin/project-operator.git)

COPY cmd/manager/project-operator /app

CMD /app/project-operator

ENTRYPOINT ["./project-operator"]

# EXPOSE 8080