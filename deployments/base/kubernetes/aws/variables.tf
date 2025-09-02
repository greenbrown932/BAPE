# deployments/kubernetes/aws/variables.tf

variable "cluster_name" {
  type        = string
  description = "The name of the EKS cluster"
  default     = "bape-dev-cluster"
}

variable "region" {
  type        = string
  description = "The AWS region to deploy the cluster in"
  default     = "us-east-1"
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR block for the VPC"
  default     = "10.0.0.0/16"
}

variable "private_subnet_count" {
  type        = number
  description = "The number of private subnets to create"
  default     = 2
}

variable "node_group_instance_types" {
  description = "List of instance types for the EKS node group"
  type        = list(string)
  default     = ["t3.small"]
}


variable "kubernetes_version" {
  type        = string
  description = "The Kubernetes version for the cluster"
  default     = "1.28" # Or the latest supported version
}

variable "node_group_ami_type" {
  type        = string
  description = "The AMI type for the node group"
  default     = "AL2_x86_64" # Amazon Linux 2
}

variable "node_group_capacity_type" {
  type        = string
  description = "The capacity type for the node group (ON_DEMAND or SPOT)"
  default     = "ON_DEMAND"
}

variable "node_group_desired_size" {
  type        = number
  description = "The desired number of worker nodes"
  default     = 2
}

variable "node_group_max_size" {
  type        = number
  description = "The maximum number of worker nodes"
  default     = 3
}

variable "node_group_min_size" {
  type        = number
  description = "The minimum number of worker nodes"
  default     = 1
}
