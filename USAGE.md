# Create pods

kubectl apply -f client.yaml -f server.yaml

# Make a request

kubectl exec client -c fortio -- fortio curl server:9080

# Network layout

Each pod runs a container on :8080, exposed as a service on :9080

Each envoy template is configured for inbound to send to :8080, and for
outbound to resolve `${ISTIO_SERVICE}` via DNS, and send all HTTP traffic there.

Pods for client and server are set-up so that client sends outbound to the
server and vice versa.

There is an access log enabled on each proxy:

kubectl logs client -c istio-proxy
