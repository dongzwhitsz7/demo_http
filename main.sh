curl --location --request GET "http://localhost:8888/healthz"
curl --location --request GET "http://10.0.2.5:32523/healthz"
curl --location --request GET "http://localhost:8888/api/v1/hello"

function docker_build() {
    docker build . -t httpserver
    docker tag httpserver dongzw/httpserver:v0.0.0
    docker push dongzw/httpserver:v0.0.0
}

function k8s_common() {
  kubectl logs -f nginx-deployment-76d6c9b8c-jq9pf
  kubectl explain service.spec
  kubectl exec nginx-deployment-76d6c9b8c-j4hmw -c nginx -it -- /bin/bash
}

function cri_find() {
    crictl inspect some_container_id | grep pid
    nsenter -t ${PID} -n ip addr # 查看目标网络namepace和pid对应的ip地址
    nsenter -t ${PID} -n arping some_ip
}