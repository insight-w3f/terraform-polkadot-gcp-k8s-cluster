package test

import (
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"log"
	"os"
	"testing"
)

func TestTerraformDefaults(t *testing.T) {
	t.Parallel()

	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/defaults")
	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)

	_, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	terraformOptions := &terraform.Options{
		TerraformDir: exampleFolder,
		Vars: map[string]interface{}{
			"gcp_project": projectID,
		},
	}

	defer test_structure.RunTestStage(t, "teardown", func() {
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "setup", func() {
		terraform.InitAndApply(t, terraformOptions)
	})
}
