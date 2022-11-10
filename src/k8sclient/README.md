# K8s Client
A `Go` application to manage reources in the target container platform.

## Running in local environment
Requirements:
* `Go` > 1.19
* `curl` or similar HTTP client

Run the following to update the dependencies and build the executable for your platform:
```bash
go mod tidy
go build main.go
```

Run as:
```bash
./main
```

In alternative, build and run as:
```bash
go run main.go
```

Use the following environment variables to update the default configuration:
| Environment variable | Definition | Default |
|----------------------|-------------|---------|
|LOG_LEVEL| Log level, one of debug, info, warn|`info`|
|SERVER_PORT|Server port|`8181`|

## Running in containerized environment
Requirements:
* `Docker`

Run this command to generate the `k8sclient` image and then run it:
```bash
docker build -t k8sclient:dev .
docker run --rm -p 8080:8080 k8sclient:dev
```

## Manual validation
Validate the execution as:
```bash
> curl localhost:8080/health
{"status":200,"message":"Hello from me!"}
```

## Unit testing
Unit tests are executed using:
```bash
go test -v ./...
```

To generate code coverage and view the results:
```bash
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Deploying on Openshift
First push the image generated using the instructions at [](#running-in-containerized-environment) on your image registry like:
```bash
docker tag k8sclient:dev quay.io/USER/k8sclient:dev
docker push quay.io/USER/k8sclient:dev
```
(replace `USER` with your user ID in `quay.io`)

docker tag k8sclient:dev quay.io/dmartino/k8sclient:dev
docker push quay.io/dmartino/k8sclient:dev


Edit the file [deploy/deployment.yaml](./deploy/deployment.yaml) and replace `<REGISTRY>/<NAME>:<TAG>` with the actual
image registry, name and tag. Following the previous example, it should be something like: `quay.io/USER/k8sclient:dev` 

Deploy the application and the access service:
```bash
oc apply -f deploy/
```
