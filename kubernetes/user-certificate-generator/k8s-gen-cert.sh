#!/bin/bash

#向K8s发起证书申请请求名称
csr_name="my-client-usr"

#Kubeconfig 用户标识名称
name="${1:-wangmin}"

#证书申请文件
csr="${2}"

cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: ${csr_name}
spec:
  groups:
  - system:authenticated
  request: $(cat ${csr} | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - client auth
EOF

echo
echo "Approving signing request."
kubectl certificate approve ${csr_name}

echo
echo "Downloading certificate."
kubectl get csr ${csr_name} -o jsonpath='{.status.certificate}' \
  | base64 --decode > $(basename ${csr} .csr).crt

echo
echo "Cleaning up"
kubectl delete csr ${csr_name}

echo
echo "Add the following to the 'users' list in your kubeconfig file:"
echo "- name: ${name}"
echo "  user:"
echo "    client-certficate: ${PWD}/$(basename ${csr} .csr).crt"
echo "    client-key: ${PWD}/$(basename ${csr} .csr)-key.pem"
echo
echo "Next you may want to add role-binding for this user."
