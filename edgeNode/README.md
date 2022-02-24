
## 说明

作为边，部署k8snode




```shell
vagrant up

```

## 环境准备
### 备份CentOS-Base.repo备份

```shell

cd /etc/yum.repos.d
cp CentOS-Base.repo CentOS-Base.repo.bak

```

### 使用阿里云源替换本地源
```shell
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```

### 清除yum缓存并更新
```shell
yum makecache
yum -y update
```


### 安装docker

```shell

yum remove docker  docker-common docker-selinux docker-engine

yum install -y yum-utils device-mapper-persistent-data lvm2

yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# yum list docker-ce --showduplicates | sort -r

yum -y install docker-ce-18.06.3.ce-3.el7

```

systemctl enable docker.service

systemctl start  docker

https://cloud.tencent.com/developer/article/1701451



wget https://github.com/openyurtio/openyurt/releases/download/v0.6.0/yurtctl; chmod u+x yurtctl;

./yurtctl join 114.67.122.239:6443 --token=mqtv7p.dhqv7szz6csif7o1 --discovery-token-unsafe-skip-ca-verification --organizations=openyurt:tenant:19818831111 --v=5



kubeadm join 114.67.122.239:6443 --token u88fhf.n4hq2bax2586xeh2 --discovery-token-ca-cert-hash sha256:0dbeba021a613e48e8313eebe386b204c6c9302050886fbe91c8d730948fa3f7  --ignore-preflight-errors all --config /tmp/kubeadm.yaml


https://gist.github.com/nilesh93/c743205d34fedb5f48ae4d37d959ba4b


https://www.jianshu.com/p/872e2e2e502d