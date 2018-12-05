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