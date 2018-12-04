#! /bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace


target="${1}"
ns="kasten-io"

backup_name=$(docker run --rm -ti -p 5432 -p 9000-9030 -p 41134 --volumes-from $(hostname) -e HOME=${SHIPPABLE_BUILD_DIR} -e KUBECONFIG=${SHIPPABLE_BUILD_DIR}/kubectl_cfg lachlanevenson/k8s-kubectl:v1.11.5 get restorepoints -n "${ns}" -ojson | jq ' .items | sort_by(.metadata.creationTimestamp) | .[-1] | .metadata.name ')

cat << EOF | kubectl create -n ${ns} -f -
apiVersion: actions.kio.kasten.io/v1alpha1
kind: RestoreAction
metadata:
  generateName: restore-latest-
spec:
  actionMeta:
    subject:
       apiVersion: apps.kio.kasten.io/v1alpha1
       kind:       RestorePoint
       namespace:  ${ns}
       name:       ${backup_name}
  params:
    targetNamespace: ${target}
    pointInTime: 2030-01-01T00:00:00Z
EOF

restore_name=$(docker run --rm -ti -p 5432 -p 9000-9030 -p 41134 --volumes-from $(hostname) -e HOME=${SHIPPABLE_BUILD_DIR} -e KUBECONFIG=${SHIPPABLE_BUILD_DIR}/kubectl_cfg lachlanevenson/k8s-kubectl:v1.11.5 get restoreactions -n "${ns}" -ojson | jq -r ' .items | sort_by(.metadata.creationTimestamp) | .[-1] | .metadata.name ')

state=""
while [[ "${state}" != "Passed" ]] && [[ "${state}" != "Failed" ]]
do
    sleep 3
    state=$(docker run --rm -ti -p 5432 -p 9000-9030 -p 41134 --volumes-from $(hostname) -e HOME=${SHIPPABLE_BUILD_DIR} -e KUBECONFIG=${SHIPPABLE_BUILD_DIR}/kubectl_cfg lachlanevenson/k8s-kubectl:v1.11.5 get restoreactions -n "${ns}" "${restore_name}" -ojson | jq -r '.status.state')
done

if [[ "${state}" == "Failed" ]]
then
    exit 1
fi
