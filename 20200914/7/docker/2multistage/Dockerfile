FROM golang:onbuild as BuildStage
ADD . /src
RUN cd /src && go build -o app

FROM scratch
WORKDIR /appdir
COPY --from=BuildStage /src/app /appdir/
CMD ["./app"]