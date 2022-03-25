FROM scratch
WORKDIR /root/
COPY ./cmd/app ./app
CMD ["./app"]