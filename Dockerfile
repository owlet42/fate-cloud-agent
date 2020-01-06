FROM golang as build

WORKDIR /workdir

COPY . ./

RUN go build -a -ldflags '-s' -installsuffix cgo -o kubefate kubefate.go

FROM scratch

COPY --from=0 /workdir/kubefate /workdir/

EXPOSE 8080

CMD ["/workdir/kubefate"]
