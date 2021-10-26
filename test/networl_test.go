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

	// 	defer terraform.Destroy(t, terraformOptions)

	// 	terraform.InitAndApply(t, terraformOptions)

	actualNetworkName := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, "vpc-dev", actualNetworkName)

}
