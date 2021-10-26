# Region
variable "aws_region" {
  default = "ap-northeast-1"
}

# VPC
variable "vpc_name" {
    type = string
    default = "vpc-dev"
}

variable "vpc_cidr" {
    type = string
    default = "10.0.0.0/16"  
}

variable "vpc_azs" {
    type = list(string)
    default = ["ap-northeast-1a", "ap-northeast-1c"]   
}

variable "private_subnets" {
    type = list(string)
    default = ["10.0.1.0/24", "10.0.2.0/24"]
   
}

variable "public_subnets" {
    type = list(string)
    default = ["10.0.101.0/24", "10.0.102.0/24"]
}

variable "enable_nat_gateway" {
    type        = bool
    default = false    
}

variable "enable_vpn_gateway" {
    type        = bool
    default = false
}

# Security Group
variable "security_group_name" {
    type = string
    default = "security-group"
}

variable "ingress_cidr_blocks" {
    type = list(string)
    default = ["10.10.0.0/16"]
}

variable "ingress_rules" {
    type = list(string)
    default = ["http-80-tcp","https-443-tcp"]
}

variable "egress_rules" {
    type = list(string)
    default = ["all-all"]
}



