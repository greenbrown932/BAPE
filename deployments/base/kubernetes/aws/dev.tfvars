# deployments/base/kubernetes/aws/dev.tfvars

# General Configuration
cluster_name = "bape-dev-cluster"
region       = "us-east-1" # Replace with your AWS region
vpc_cidr     = "10.0.0.0/16"

# Subnet Configuration
private_subnet_count = 2

# Node Group Configuration
node_group_ami_type       = "AL2_x86_64" # Amazon Linux 2 (x86_64)
node_group_capacity_type  = "SPOT"       # Or "SPOT" for spot instances
node_group_desired_size   = 2
node_group_max_size       = 3
node_group_min_size       = 1
node_group_instance_types = ["t3.medium"] # Or your preferred instance types

# Kubernetes Version
kubernetes_version = "1.28" # Replace with your desired Kubernetes version
