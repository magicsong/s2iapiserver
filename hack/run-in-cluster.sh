#/bin/bash

# git add -A -- .
# git commit --quiet --allow-empty-message --file - --all -m "debug auto commit"

#tag=$(git rev-parse --short HEAD)
tag=93f63fd5
image_name=dockerhub.qingcloud.com/magicsong/s2iapiserver:$tag
namespace=devops

# docker build -f build/Dockerfile bin/ -t $image_name
# docker push $image_name

rm -f config/apiserver.yaml
apiserver-boot build config --name s2iapiserver --namespace $namespace --image $image_name --service-account=s2i-service-account \
--storage-class="gluster-heketi" --image-pull-secrets="s2i-pull-secret" --apiserver-args "--loglevel=3" --controller-args "-v=3" --controller-args "-logtostderr=true"

kubectl apply -f config/apiserver.yaml

# watch
timeout=20
time=0
while true; do
    kubectl get pod -n $namespace -l 'apiserver=true' | grep "2/2"
    if [ $? -eq 0 ];then
        echo "apiserver is ready now"
        break
    fi
    sleep 1
    time=$[$time + 1]
    if [ $time -gt $timeout ]; then
        echo "Timeout waiting for apiserver to be ready"
        exit 1
    fi
done 