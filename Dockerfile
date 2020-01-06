FROM golang as build

WORKDIR /workdir

COPY . ./

RUN go build -o kubefate kubefate.go

FROM centos

COPY --from=0 /workdir/kubefate /workdir/

EXPOSE 8080

CMD ["/workdir/kubefate"]
