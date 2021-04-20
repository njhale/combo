# combo

`combo` is a [Kubernetes controller](https://kubernetes.io/docs/concepts/architecture/controller/) that generates and applies resources for all combinations of a manifest template and its arguments.

## Usage

## On the command line

Evaluate a template directly:

```sh
$ cat <<EOF | combo -a 'NAMESPACE=foo,bar' -a 'NAME=baz' -
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: deployment-reader
rules:
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "watch", "list"]
---
kind: Namespace
metadata:
  name: NAMESPACE
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: NAME
  namespace: NAMESPACE
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: deployment-reader
  namespace: NAMESPACE
subjects:
- kind: ServiceAccount
  name: NAME
  namespace: NAMESPACE
roleRef:
  kind: ClusterRole
  name: deployment-reader
  apiGroup: rbac.authorization.k8s.io
EOF
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: deployment-reader
rules:
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "watch", "list"]
---
kind: Namespace
metadata:
  name: foo
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: baz
  namespace: foo
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: deployment-reader
  namespace: foo
subjects:
- kind: ServiceAccount
  name: baz
  namespace: foo
roleRef:
  kind: ClusterRole
  name: deployment-reader
  apiGroup: rbac.authorization.k8s.io
---
kind: Namespace
metadata:
  name: bar
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: baz
  namespace: bar
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: deployment-reader
  namespace: bar
subjects:
- kind: ServiceAccount
  name: baz
  namespace: bar
roleRef:
  kind: ClusterRole
  name: deployment-reader
  apiGroup: rbac.authorization.k8s.io
```

### As a controller

If `combo` is running as a controller in the current kubectl context's cluster, create a `Pattern`:

```sh
$ cat <<EOF | kubectl create -f -
apiVersion: combo.io/v1alpha1
kind: Pattern
metadata:
  name: feature
spec:
  template: |
    ---
    apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
      name: feature-controller
      namespace: TARGET_NAMESPACE
    subjects:
    - kind: ServiceAccount
      name: controller
      namespace: feature
    roleRef:
      kind: ClusterRole
      name: feature-controller
      apiGroup: rbac.authorization.k8s.io
    ---
    apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
      generateName: feature-user-
      namespace: TARGET_NAMESPACE
    subjects:
    - kind: Group
      name: TARGET_GROUP
      namespace: TARGET_NAMESPACE
      apiGroup: rbac.authorization.k8s.io
    roleRef:
      kind: ClusterRole
      name: feature-user
      apiGroup: rbac.authorization.k8s.io
  parameters:
  - key: TARGET_GROUP
  - key: TARGET_NAMESPACE
EOF
```

Assuming the existance of the `feature-controller` and `feature-user` `ClusterRoles` as well as the `feature`, `staging`, and `prod` `Namespaces`, instantiate all resource/argument combinations with a `CombinationSet`:

```sh
$ cat <<EOF | kubectl create -f -
apiVersion: combo.io/v1alpha1
kind: CombinationSet
metadata:
  name: enable-feature
spec:
  pattern:
    name: feature
  arguments:
  - key: TARGET_GROUP
    values:
    - "sre"
    - "system:serviceaccounts:ci"
  - key: TARGET_NAMESPACE
    values:
    - staging
    - prod
EOF
```

Get all of the `RoleBindings` applied by `combo` for the `CombinationSet` above:

```sh
$ kubectl get rolebindings -A -l 'combo.io/set'=enable-feature
NAMESPACE   NAME                 ROLE                             AGE
prod        feature-controller   clusterrole/feature-controller   1m
prod        feature-user-gpl8b   clusterrole/feature-user         1m
prod        feature-user-19ssf   clusterrole/feature-user         1m
staging     feature-controller   clusterrole/feature-controller   1m
staging     feature-user-r9s81   clusterrole/feature-user         1m
staging     feature-user-6aa3c   clusterrole/feature-user         1m
```

