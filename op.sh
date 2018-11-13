cat <<EOF | kubectl create -f -
apiVersion: config.istio.io/v1alpha2
kind: handler
metadata:
  name: test
  namespace: istio-system
spec:
  adapter: ditto
  connection:
    address: ":9090"
  params:
    api_requirements:
    - endpoints:
      - endpoint: /test
        method: GET
        name: endpoint
        requirements:
        - actions:
          - get
          resource: endpoint
      version: null
EOF
cat <<EOF | kubectl create -f -
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
  name: test
  namespace: istio-system
spec:
  actions:
  - handler: test.handler.istio-system
    instances:
    - ditto-authorization.instance.istio-system

EOF
sleep 2
kubectl delete rule test -n istio-system
kubectl delete handler test -n istio-system
