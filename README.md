# go-gin-api
[gin-gonic](https://github.com/gin-gonic/gin)を使ったgoによるAPIのテスト実装

# Dockerfileを使う

```
$ docker pull ytakky2014/go-gin-api:v1
$ docker run -d -p 8080:8080 ytakky2014/go-gin-api:v1
```


リクエスト

```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```

レスポンス

```
{"message":"Hello World!!"}
```

# GKE上で動作させる
- [cloud sdk](https://cloud.google.com/sdk/?hl=ja)をインストールしておく
- gcloud をインストールしておく

```
gcloud components install kubectl
```

- gke clusterへの認証を通す

```
$ gcloud container clusters get-credentials $GKE_CLUSTER_NAME --zone $GKE_CLUSTER_ZONE --project $GKE_PROJECT
```

- kubernetesのマニフェストでデプロイする。

```
$ kubectl create -f go-gin-api-deployment.yaml
$ kubectl create -f go-gin-api-service.yaml
```

IPを確認する

```
$ kubectl get svc
NAME         TYPE           CLUSTER-IP      EXTERNAL-IP       PORT(S)          AGE
go-gin-api   LoadBalancer   10.27.249.245   130.211.138.145   8080:32251/TCP   14m
```

※ここで作成されるIPはエフェメラルIPなので、固定IPにする場合はIngressを使うか、Google Load BalancerのTCPロードバランサの編集からフロントエンドに固定IPを当てれば良い

リクエスト

```
$ curl -XGET -H 'Content-Type:application/json' http://130.211.138.145:8080/
```

レスポンス

```
{"message":"Hello World!!"}
```
