FROM golang:alpine as BuildStage
ADD . /src
RUN cd /src && go build -o app

FROM alpine
WORKDIR /appdir
EXPOSE 8080
COPY --from=BuildStage /src/app /appdir/
CMD ["./app"]