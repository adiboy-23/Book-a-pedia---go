FROM cgr.dev/chainguard/go AS builder
COPY . /app
RUN cd /app && go build -o go-finder .

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/go-finder /usr/bin
CMD ["/usr/bin/go-finder"]