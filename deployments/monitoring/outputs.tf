# deployments/monitoring/outputs.tf

output "prometheus_url" {
  description = "URL to access Prometheus"
  value       = "http://localhost:9090"
}

output "grafana_url" {
  description = "URL to access Grafana"
  value       = "http://localhost:3000"
}

output "kibana_url" {
  description = "URL to access Kibana"
  value       = "http://localhost:5601"
}
