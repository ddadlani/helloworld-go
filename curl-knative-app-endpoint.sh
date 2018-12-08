#!/bin/bash


IP_ADDRESS=$(kubectl get svc knative-ingressgateway --namespace istio-system --output 'jsonpath={.status.loadBalancer.ingress[0].ip}')

PORT=$(kubectl get svc knative-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')

DOMAIN=$(kubectl get ksvc helloworld-go  --output 'jsonpath={.status.domain}')

curl -H "Host: $DOMAIN" http://${IP_ADDRESS}
