FROM ghcr.io/kwitsch/docker-buildimage:main AS build-env

ADD src .
RUN gobuild.sh -o dnscacherefresh

FROM scratch
COPY --from=build-env /builddir/dnscacherefresh /dnscacherefresh

ENTRYPOINT ["/dnscacherefresh"]