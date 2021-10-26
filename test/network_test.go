package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../src",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVpcName := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, "vpc-dev", actualVpcName)

	expectedPrivateSubnetCIDRBlocks := []string{"10.0.1.0/24", "10.0.2.0/24"}
	actualPrivateSubnetCIDRBlocks := terraform.OutputList(t, terraformOptions, "private_subnets_cidr_blocks")
	assert.Equal(t, expectedPrivateSubnetCIDRBlocks, actualPrivateSubnetCIDRBlocks)

	expectedPublicSubnetCIDRBlocks := []string{"10.0.101.0/24", "10.0.102.0/24"}
	actualPublicSubnetCIDRBlocks := terraform.OutputList(t, terraformOptions, "public_subnets_cidr_blocks")
	assert.Equal(t, expectedPublicSubnetCIDRBlocks, actualPublicSubnetCIDRBlocks)

}
