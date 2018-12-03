# Zipcode

```bash
export GO111MODULE=on
```

# CICD Test
go test

# Local Test

```bash
kubectl port-forward --namespace zipcode svc/postgresql-kanister-postgresql 5432:5432
envdir ./env go test -v ./pkg/zipcode/ -run TestResetInsert -count 1
```


# Connect to DB

```bash
//

docker run --network host -it --rm jbergknoff/postgresql-client \
    postgresql://admin:admin@127.0.0.1:5432/zipcode
```


docker build . -t kastenhq/zipcode:tom-test


