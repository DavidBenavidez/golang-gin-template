GOPRIVATE=dev.azure.com CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -mod=mod -o app ./cmd/
docker build . -t dockerrepository/servicename:1.0.0-local
docker push dockerrepository/servicename:1.0.0-local
kubectl rollout restart deployment servicename -n namespace
kubectl get pods -n namespace | grep servicename