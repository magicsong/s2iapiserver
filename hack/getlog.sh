#/bin/bash
namespace=devops
POD=$(kubectl get pod -n $namespace -l 'apiserver=true' -o jsonpath="{.items[0].metadata.name}")
kubectl logs -f $POD -c $1 -n $namespace