FROM golang:1.16-alpine as builder

RUN apk add --no-cache make
WORKDIR /upmon
COPY . ./
RUN make install-tools build

FROM golang:1.16-alpine

ENV USER=docker
ENV UID=70908

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/etc/upmon" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

WORKDIR /etc/upmon
COPY --from=builder /upmon/build/upmon /usr/local/bin/upmon

USER docker
ENTRYPOINT ["upmon"]
