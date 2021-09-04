#!/bin/bash -e
set -o pipefail

NAMESPACE=poc

kubectl apply -f - << EOF
apiVersion: v1
kind: Namespace
metadata:
  name: $NAMESPACE
  labels:
    istio-injection: enabled
EOF

kubectl apply -n $NAMESPACE -f client.yaml
kubectl apply -n $NAMESPACE -f client-l7.yaml
kubectl apply -n $NAMESPACE -f server.yaml
kubectl apply -n $NAMESPACE -f server-l7.yaml

kubectl create cm -n $NAMESPACE client-uproxy --from-file config.yaml=client-uproxy.yaml \
  --dry-run=client -o yaml | kubectl apply -f -
kubectl create cm -n $NAMESPACE server-uproxy --from-file config.yaml=server-uproxy.yaml \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl exec -itn $NAMESPACE client -c sleep -- curl server:8080
