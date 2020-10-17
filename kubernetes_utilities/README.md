# INSTRUCTIONS ON SETTING UP THE BUYTRANCE K8S CLUSTER ONE SERVICE AT A TIME

COCKROACH-DB
** Setting up cockraochdb cluster on Kubernetes **

Following instructions on this page: https://github.com/cockroachdb/cockroach/tree/master/cloud/kubernetes

Run the following:

  1. kubectl create -f cockroachdb.yaml
  2. kubectl create -f cockroachdb_cluster_init.yaml

  3. To view the dashboard run:
     kubectl port-forward cockroachdb-0 8080
  4. Optionaly run a load generator example app for CockRoachDB
     
---

NATS (STAN) STREAMING
** Seting up NATS + STAN (i.e. NATS-Streaming) server HA mode  with prometheus operator**
Following instructions on https://github.com/nats-io/k8s

 RUN: curl -sSL https://nats-io.github.io/k8s/setup.sh | sh

---

TRAEFIK
** Deploy Traefik Ingress controller to the cluster **
RUN: kubectl create -r traefik_ingress.yaml

---

LINKERD
** Deploy Linkerd Service Mesh to the cluster **
RUN: curl -sL https://run.linkerd.io/install | sh

Inject all namespaces into linkerd using the following command
RUN: kubectl get ns --all-namespaces -o yaml | linkerd inject - | kubectl apply -f -

Inject your deployments into linkerd using the following command
RUN: kubectl get deploy -o yaml | linkerd inject - | kubectl apply -f -

Inject your pods into linkerd using the following command
RUN: kubectl get pods -o yaml | linkerd inject - | kubectl apply -f -

TO BE RE-VISITED

---

TO BE CONTINUED
 - Deploy NATS Services                                           *
 - Deploy CockroachDB                                             *
 - Deploy Ingress Controller Traefik                              *
 - Deploy Service Mesh Linkerd                                    *
 - Deploying BuyTrance Services
    -  Landing Page                                               *
    -  Oriserver
    -  Message-Subscribers
        - Domainer [Searchers, Registrar]
        - Hoster [CreateHosting, SuspendHosting, DeleteHosting]
        - FileMover [CopyFiles to Host]
        - 
    -  Message-Producers
    -  REST APIs for front-ends
    -  Front-End