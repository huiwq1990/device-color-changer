

```shell

export NODEPOOL_NAME="huanan"
export EDGE_NODE="huannan-01"

# 创建hangzhou节点池
cat <<EOF | kubectl apply -f -
apiVersion: apps.openyurt.io/v1alpha1
kind: NodePool
metadata:
  name: $NODEPOOL_NAME
spec:
  type: Edge
EOF

# 将边缘节点加入hangzhou节点池
kubectl label node $EDGE_NODE apps.openyurt.io/desired-nodepool=${NODEPOOL_NAME}

# 检查节点池状态
kubectl get nodepool

spec:
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: edgex-app-rules-engine
  topology:
    pools:
    - name: huadong
      nodeSelectorTerm:
        matchExpressions:
        - key: apps.openyurt.io/nodepool
          operator: In
          values:
          - huadong
      replicas: 1
  workloadTemplate:
    deploymentTemplate:
      metadata:
        creationTimestamp: null
        labels:
          app: edgex-app-rules-engine
      spec:
        selector:
          matchLabels:
            app: edgex-app-rules-engine
        strategy: {}
        template:
          metadata:
            creationTimestamp: null
            labels:
              app: edgex-app-rules-engine
          spec:
            containers:
            - env:
              - name: SERVI



Error from server: error when creating "STDIN":
 admission webhook "vuniteddeployment.kb.io" denied the request:
  [spec.template.deploymentTemplate.metadata.labels: 
  Invalid value: map[string]string(nil): `selector` does not match template `labels`, 
  spec.template.deploymentTemplate.spec.template.metadata.labels: 
  Invalid value: map[string]string(nil): `selector` does not match template `labels`]

cat << EOF | kubectl apply -f -
apiVersion: apps.openyurt.io/v1alpha1
kind: UnitedDeployment
metadata:
  name: test-yurt-deploy
spec:
  selector:
    matchLabels:
      app: test-yurt-deploy
  workloadTemplate:
    deploymentTemplate:
      metadata:
        labels:
          app: test-yurt-deploy
      spec:
        selector:
          matchLabels:
            app: test-yurt-deploy
        template:
          metadata:
            labels:
              app: test-yurt-deploy
          spec:
            tolerations:
            - operator: Equal
              key: node-role.openyurt.io/edge
              effect: NoSchedule
            containers:
              - name: hostname
                image: mirrorgooglecontainers/serve_hostname
                ports:
                - containerPort: 9376
                  protocol: TCP
  topology:
    pools:
    - name: ${NODEPOOL_NAME}
      nodeSelectorTerm:
        matchExpressions:
        - key: apps.openyurt.io/nodepool
          operator: In
          values:
          - ${NODEPOOL_NAME}
      replicas: 1
EOF



cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: test-yurt-svc
  annotations:
    openyurt.io/topologyKeys: openyurt.io/nodepool
spec:
  selector:
    app: test-yurt-deploy
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 9376
EOF

```


在huanan的节点上执行

```shell
curl test-yurt-svc:80
```
