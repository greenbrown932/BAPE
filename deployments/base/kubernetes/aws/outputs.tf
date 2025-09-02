# deployments/kubernetes/aws/outputs.tf

output "cluster_endpoint" {
  value       = aws_eks_cluster.this.endpoint
  description = "The Kubernetes cluster endpoint"
}

output "kubeconfig" {
  value       = aws_eks_cluster.this.kubeconfig
  description = "The kubeconfig for the Kubernetes cluster"
  sensitive   = true
}

output "cluster_name" {
  value       = aws_eks_cluster.this.name
  description = "The Kubernetes cluster name"
}
