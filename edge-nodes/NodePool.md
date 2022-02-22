

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


cat << EOF | kubectl apply -f -
apiVersion: apps.openyurt.io/v1alpha1
kind: UnitedDeployment
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
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
