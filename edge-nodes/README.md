
## 说明

当前虚拟机作为OpenYurt集群的边。

```shell
vagrant up
```

## 环境准备

### 固定网卡（可选）

如果虚拟机已经设置，可以不用搞。

```yaml

cat<<EOF > /etc/netplan/00-installer-config.yaml
network:
  ethernets:
    eth0:     #配置的网卡的名称
      addresses: [192.168.68.101/24]    #配置的静态ip地址和掩码
      dhcp4: no    #关闭DHCP，如果需要打开DHCP则写yes
      optional: true
      gateway4: 192.168.68.1    #网关地址
  version: 2
EOF

netplan apply

```

### 备份CentOS-Base.repo备份

```shell
cd /etc/yum.repos.d
cp CentOS-Base.repo CentOS-Base.repo.bak
```

### 使用阿里云源替换本地源

```shell
curl -o  /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```

### 清除yum缓存并更新

```shell
yum makecache
yum -y update
```

### 安装docker

```shell

yum remove docker  docker-common docker-selinux docker-engine

yum install -y yum-utils device-mapper-persistent-data lvm2 ipvsadm wget curl

yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# yum list docker-ce --showduplicates | sort -r

yum -y install docker-ce-20.10.12-3.el7

yum -y install socat conntrack-tools

```

### 启动相关服务
```shell

systemctl enable kubelet

systemctl enable docker.service

systemctl start  docker

# https://cloud.tencent.com/developer/article/1701451

# 关闭swap
sudo swapoff -a
#https://wsgzao.github.io/post/swap/
```


### 节点添加

```shell
#wget https://github.com/openyurtio/openyurt/releases/download/v0.6.0/yurtctl; chmod u+x yurtctl;

./yurtctl join 114.67.122.239:6443 --token=mqtv7p.dhqv7szz6csif7o1 --discovery-token-unsafe-skip-ca-verification --organizations=openyurt:tenant:19818831111  --ignore-preflight-errors all --v=5

./yurtctl join 114.67.122.239:6443 --token=mqtv7p.dhqv7szz6csif7o1 --discovery-token-unsafe-skip-ca-verification --organizations=openyurt:tenant:19818831111 --v=5
# kubeadm join 114.67.122.239:6443 --token u88fhf.n4hq2bax2586xeh2 --discovery-token-ca-cert-hash sha256:0dbeba021a613e48e8313eebe386b204c6c9302050886fbe91c8d730948fa3f7  --ignore-preflight-errors all 
# --config /tmp/kubeadm.yaml
```


```shell
kubectl label node huannan-01 openyurt.io/node-type=edge
```

## 参考

https://gist.github.com/nilesh93/c743205d34fedb5f48ae4d37d959ba4b
https://www.jianshu.com/p/872e2e2e502d