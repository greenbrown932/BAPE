# deployments/monitoring/logging.tf

resource "helm_release" "elasticsearch" {
  name       = "elasticsearch"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "elasticsearch"
  namespace  = kubernetes_namespace.monitoring.metadata[0].name
  version    = "19.16.2" # Use a specific version for reproducibility

  values = [
    <<-EOT
    master:
      replicas: 1
    data:
      replicas: 1
    coordinating:
      replicas: 1
    EOT
  ]
}

resource "helm_release" "kibana" {
  name       = "kibana"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "kibana"
  namespace  = kubernetes_namespace.monitoring.metadata[0].name
  version    = "12.6.6" # Use a specific version

  values = [
    <<-EOT
    elasticsearch:
      hosts: ["elasticsearch-master.monitoring.svc.cluster.local"]
      port: 9200
    EOT
  ]

  depends_on = [helm_release.elasticsearch]
}

resource "helm_release" "fluentd" {
  name       = "fluentd"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "fluentd"
  namespace  = kubernetes_namespace.monitoring.metadata[0].name
  version    = "5.8.1" # Use a specific version

  values = [
    <<-EOT
    aggregator:
      enabled: false
    forwarder:
      enabled: true
      elasticsearch:
        host: "elasticsearch-master.monitoring.svc.cluster.local"
        port: 9200
    EOT
  ]

  depends_on = [helm_release.elasticsearch]
}
