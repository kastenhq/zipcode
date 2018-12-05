#! /bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace


set +o xtrace
echo -e "${ZIPCODEK8SCFG_INTEGRATION_KUBECONFIGCONTENT}" > kubectl_cfg
set -o xtrace

export KUBECONFIG=kubectl_cfg
export GO111MODULE=on

export helm_cmd=( docker run --rm -t --volumes-from $(hostname) -e HOME=${SHIPPABLE_BUILD_DIR} -e KUBECONFIG=${SHIPPABLE_BUILD_DIR}/kubectl_cfg lachlanevenson/k8s-helm:v2.11.0 )
export kubectl_cmd=( docker run --rm -ti -p 5432 -p 9000-9030 -p 41134 --volumes-from $(hostname) -e HOME=${SHIPPABLE_BUILD_DIR} -e KUBECONFIG=${SHIPPABLE_BUILD_DIR}/kubectl_cfg lachlanevenson/k8s-kubectl:v1.11.5 )

go mod download
go build -v ./cmd/zipcode


${helm_cmd[@]} init --upgrade --force-upgrade --wait
${helm_cmd[@]} repo add kanister https://charts.kanister.io/
${helm_cmd[@]} install kanister/kanister-postgresql \
-n "zipcode-${BRANCH}-${BUILD_NUMBER}" \
--namespace "postgresql-${BRANCH}-${BUILD_NUMBER}" \
--set postgresDatabase=zipcode \
--set postgresPassword=admin \
--set postgresUser=admin \
--wait

sleep 20

./restore.sh "postgresql-${BRANCH}-${BUILD_NUMBER}"

${kubectl_cmd[@]} get all -n "postgresql-${BRANCH}-${BUILD_NUMBER}"
${kubectl_cmd[@]} get pods --selector=app="zipcode-${BRANCH}-${BUILD_NUMBER}-kanister-postgresql" -n "postgresql-${BRANCH}-${BUILD_NUMBER}"
${kubectl_cmd[@]} port-forward $(${kubectl_cmd[@]} get pods --selector=app="zipcode-${BRANCH}-${BUILD_NUMBER}-kanister-postgresql" -n "postgresql-${BRANCH}-${BUILD_NUMBER}" --output=jsonpath={.items..metadata.name}) -n "postgresql-${BRANCH}-${BUILD_NUMBER}" 5432:5432 &

export pf_pid="${!}"

for fi in $(ls env/); do export $fi=$(cat env/$fi); done

go test -v ./pkg/zipcode/ -run TestCities -count 1
