#! /bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace


target="${1}"

backup=$(kubectl get restorepoints -n kasten-io -ojson | jq ' .items | sort_by(.metadata.creationTimestamp) | .[-1] | .metadata.name ')

cat << EOF | kubectl create -f -
apiVersion: actions.kio.kasten.io/v1alpha1
kind: RestoreAction
metadata:
  generateName: restore-latest-
spec:
  actionMeta:
    subject:
       apiVersion: apps.kio.kasten.io/v1alpha1
       kind:       RestorePoint
       name:       ${backup}
       namespace:  kasten-io
  params:
    targetNamespace: ${target}
    pointInTime: 2030-01-01T00:00:00Z
EOF

kubectl get restoreaction.actions.kio.kasten.io -oyaml

