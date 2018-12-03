# Zipcode

Store your orders by Zipcode! A demo app for [CICD, Kubernetes, and Databases:
Better Together
](https://kccna18.sched.com/event/GrSq/cicd-kubernetes-and-databases-better-together-niraj-tolia-tom-manville-kasten)

## Prerequisites

### Go Setup

* go 1.11.x+
* [envdir](http://manpages.ubuntu.com/manpages/trusty/man8/envdir.8.html)

```bash
# Ensure go modules are enabled
$ export GO111MODULE=on

# Download dependencies
$ go mod download
```

Test Install
```bash
$ go build -v ./cmd/zipcode
```

### PostgreSQL

This app requires a access to a Postgres instance.

We recommend installing it using the Kanister example [helm chart](https://docs.kanister.io/helm_instructions/pgsql_instructions.html).

```bash
# Install Kanister-enabled PostgreSQL
$ helm install kanister/kanister-postgresql -n zipcode \
    --namespace postgresql \
    --set postgresDatabase=zipcode \
    --set postgresPassword=admin \
    --set postgresUser=admin
```

Assuming your settings match those in `./env`, you can run:
```bash
$ kubectl port-forward --namespace zipcode svc/postgresql-kanister-postgresql 5432:5432
$ docker run --network host -it --rm jbergknoff/postgresql-client postgresql://admin:admin@127.0.0.1:5432/zipcode
```

## Local Testing

You can run tests locally, but you'll still need to connect to a PostgreSQL
instance.
```bash
# Expost PostgreSQL at `localhost:5432`
$ kubectl port-forward --namespace zipcode svc/postgresql-kanister-postgresql 5432:5432

# Use environment vars in `./env` to connect to PostgreSQL and run unit tests.
$ envdir ./env go test -v ./pkg/zipcode/ -run TestResetInsert -count 1

# Run the full server. The service is available at `localhost:8000`.
$ envdir env go run -v ./cmd/zipcode/
```


## Deploy

Build and push docker image

```bash
# We suggest adding a version tag to the image. You'll need to update
# ./deploy/deployment.yaml
$ docker build . -t kastenhq/zipcode
```

Deploy service to Kubernetes
```bash
$ kubectl apply -f ./deploy

$ kubectl port-forward --namespace zipcode svc/zipcode 8000:8000
```

The zipcode app will now be available at [localhost:8000]
