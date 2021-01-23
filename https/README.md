# hello-world-https

## 证书

### 生成证书


```sh
export EMAIL=test@qq.com
export COMMON_NAME=hello-world.com
export COUNTRY_NAME=CN
export STATION=ShangHai
export LOCALITY_NAME=ShangHai
export ORGANIZATION=Global Security
export ORGANIZATIONAL_UNIT_NAME=DEV

openssl req -newkey \
  rsa:4096 \
  -nodes \
  -sha256 \
  -keyout server.key \
  -subj "/C=${COUNTRY_NAME}/ST=${STATION}/L=${LOCALITY_NAME}/O=${ORGANIZATION}/OU=${ORGANIZATIONAL_UNIT_NAME}/CN=${COMMON_NAME}"  \
  -x509 \
  -days 3650 \
  -out server.pem
```

### 使系统信任证书

### 测试

#### 配置`/etc/hosts`

```
127.0.0.1 hello-world.com
```

```sh
curl https://hello-world.com:8443
```

## 打包镜像

```sh
## build
docker build -t  wanyanchengli/hello-world-https:latest .


## push
docker push wanyanchengli/hello-world-https:latest


## run
docker run --rm  -p 8443:8443 -v $PWD/cert:/cert wanyanchengli/hello-world-https:latest
```