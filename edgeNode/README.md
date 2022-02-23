
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

https://cloud.tencent.com/developer/article/1701451
