# telepolice [![Build Status](https://github.com/takumakume/telepolice/workflows/build/badge.svg)](https://github.com/takumakume/telepolice/actions) [![GitHub release](https://img.shields.io/github/release/takumakume/telepolice.svg)](https://github.com/takumakume/telepolice/releases)

This tool is [telepresenceio/telepresence](https://github.com/telepresenceio/telepresence) cleaner.

Using telepresence `--swap-deployment` , telepresence deployment does not work if process crashes or laptop is closed. (ref: https://github.com/telepresenceio/telepresence/issues/260)

![image](/docs/telepolice_image.jpeg)

Clean according to the following flow.

1. Collect pods with `telepresence` labels in the specified namespace.
2. Check if a valid sshd process exists for those pods. If it does not exist, it is broken. (*1)
3. Perform the same processing as the Telepresence Cleanup on the invalid Pods.


<details>
<summary>(*1) Broken Pod status</summary>

When telepresence is working:

```sh
~ $ ps -elf
PID   USER     TIME   COMMAND
    1 telepres   0:00 {twistd} /usr/bin/python3.6 /usr/bin/twistd --pidfile= -n -y ./forwarder.py
    8 telepres   0:00 [sshd]
    9 telepres   0:00 /usr/sbin/sshd -e
   14 telepres   0:00 [sshd]
   17 telepres   0:00 sshd: telepresence [priv]
   18 telepres   0:00 sshd: telepresence [priv]
   21 telepres   0:00 sshd: telepresence
   22 telepres   0:00 sshd: telepresence
   28 telepres   0:00 sshd: telepresence [priv]
   30 telepres   0:00 sshd: telepresence@notty
   31 telepres   0:00 ash -c /usr/lib/ssh/sftp-server
   32 telepres   0:00 /usr/lib/ssh/sftp-server
   34 telepres   0:00 sh
   39 telepres   0:00 ps -elf
```

When telepresence is not working:

```sh
~ $ ps -elf
PID   USER     TIME   COMMAND
    1 telepres   0:00 {twistd} /usr/bin/python3.6 /usr/bin/twistd --pidfile= -n -y ./forwarder.py
    8 telepres   0:00 [sshd]
    9 telepres   0:00 /usr/sbin/sshd -e
   14 telepres   0:00 [sshd]
   21 telepres   0:00 [sshd]
   22 telepres   0:00 [sshd]
   30 telepres   0:00 [sshd]
   31 telepres   0:00 [ash]
   34 telepres   0:00 sh
   43 telepres   0:00 ps -elf
```

telepolice sees the state of sshd process.

</details>

## Add annotations to deployment

```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: app-deployment
spec:
  annotations:
    telepolice/original-deployment: app-deployment
    telepolice/original-deployment-replicas: 1
...
```

Previously, metadata.selflink of last-applied-configuration of annotations was used, but it became deprecated.

## Usage

### Get telepresence resources

```sh
$ telepolice get
NAMESPACE STATUS POD
default   true   web-9bc4ceefeade40668638a0b782decec9-57f78598krqr
```

Broken case (STATUS = false)

```sh
$ telepolice get
NAMESPACE STATUS POD
default   false  web-9bc4ceefeade40668638a0b782decec9-57f78598krqr
```

### Cleanup telepresence resources

```sh
% telepolice cleanup
Cleanup: default/web-9bc4ceefeade40668638a0b782decec9-57f78598krqr
```

Dry run mode by adding `--dry-run` option.

## Use cases

### Perform cleanup at intervals

```sh
$ telepolice cleanup -i 60
```

Start as a daemon in the foreground.
Clean every 60 seconds.

### Target some namespaces

```sh
$ telepolice get -n ns1,ns2
```

default: `default` namespace

### Target all namespaces

```sh
$ telepolice get -A
```

### Specify kubeconfig

```sh
$ KUBECONFIG=~/.kube/other_config telepolice get
```

default: `~/.kube/config`

### Use in cluser config (kubernetes ServiceAccount)

```sh
$ telepolice --use-in-cluster-config get
```

## Install as a cleaner on kubernetes 

- use master
  ```sh
  kubectl apply -f https://raw.githubusercontent.com/takumakume/telepolice/master/manifests/release.yaml
  ```
- use tag
  ```sh
  kubectl apply -f https://raw.githubusercontent.com/takumakume/telepolice/v0.0.1/manifests/release.yaml
  ```

#### Custom configration

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: telepolice-config
  namespace: telepolice
data:
  arg: "cleanup --use-in-cluster-config -A -i 30 --verbose"
```

edit `arg`

### Install Cli tool

```sh
go get github.com/takumakume/telepolice
```
