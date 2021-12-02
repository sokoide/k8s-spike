# K8s spike

# testapp1

* For per-node daemonset test
* How to build

```sh
make
# sokoide/testapp1:latest image created
```

* How to run

```sh
docker run -it --rm -e PORT=10001 -p 10001:10001 sokoide/testapp1:latest
```

* How to test

```sh
curl localhost:10001
# 'Hello' printed
```

# testapp2

* For per-app sidecar test
* How to build

```sh
make
# sokoide/testapp2:latest image created
```

* How to run

```sh
docker run -it --rm -e PORT=10001 -p 10001:10001 sokoide/testapp2:latest
```

* How to test

```sh
curl localhost:10001
# 'handler called' written in $(CWD)/testapp2.log
```

## fluent-bit per-node (Daemonset)

* The configuraiton runs fluent-bit container per-node using `configs/fluentbit-loki.yaml` and `configs/logging.configmap.yaml`
* `auto_kubernetes_labels  on` adds `app` and `pod_template_hash` labels automatically which is nice!
* The daemonset reads `/var/log/containers/*.log` and sends to Loki
* It can't read app files in a pod. To support app log files, per-pod sidecar is needed to relay it.

## fluent-bit per-pod (Sidecar)

* The configuration runs fluent-bit container in the same pod as an app
* Delete the daemonset by `k delete daemonset.apps/fluent-bit -n logging`
