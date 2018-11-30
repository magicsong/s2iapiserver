#/bin/bash

git add -A -- .
git commit --quiet --allow-empty-message --file - --all

tag=$(git rev-parse --short HEAD)
image_name=dockerhub.qingcloud.com/magicsong/s2iapiserver:$tag

docker build -f build/Dockerfile bin/ -t $image_name
docker push $image_name

apiserver-boot build config --name s2iapiserver --namespace default --image $image_name --apiserver-args "-v 3 -logtostderr=true" --controller-args "-v 3 -logtostderr=true"
kubectl apply -f config/apiserver.yaml