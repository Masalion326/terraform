output "name" {
  value = module.vpc.name
}

output "public_subnets_cidr_blocks" {
  value = module.vpc.public_subnets_cidr_blocks
}

output "private_subnets_cidr_blocks" {
  value = module.vpc.private_subnets_cidr_blocks
}

output "security_group_id" {
  value = module.security-group.security_group_id
}

output "security_group_name" {
  value = module.security-group.security_group_name
}

output "security_group_vpc_id" {
  value = module.security-group.security_group_vpc_id
}