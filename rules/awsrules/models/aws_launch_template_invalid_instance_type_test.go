// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/configs/configload"
	"github.com/hashicorp/terraform/terraform"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

func Test_AwsLaunchTemplateInvalidInstanceTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected issue.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_launch_template" "foo" {
	instance_type = "t1.2xlarge"
}`,
			Expected: []*issue.Issue{
				{
					Detector: "aws_launch_template_invalid_instance_type",
					Type:     "ERROR",
					Message:  `instance_type is not a valid value`,
					Line:     3,
					File:     "resource.tf",
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_launch_template" "foo" {
	instance_type = "t2.micro"
}`,
			Expected: []*issue.Issue{},
		},
	}

	dir, err := ioutil.TempDir("", "AwsLaunchTemplateInvalidInstanceTypeRule")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(currentDir)

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		loader, err := configload.NewLoader(&configload.Config{})
		if err != nil {
			t.Fatal(err)
		}

		err = ioutil.WriteFile(dir+"/resource.tf", []byte(tc.Content), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		mod, diags := loader.Parser().LoadConfigDir(".")
		if diags.HasErrors() {
			t.Fatal(diags)
		}
		cfg, tfdiags := configs.BuildConfig(mod, configs.DisabledModuleWalker)
		if tfdiags.HasErrors() {
			t.Fatal(tfdiags)
		}

		runner := tflint.NewRunner(tflint.EmptyConfig(), map[string]tflint.Annotations{}, cfg, map[string]*terraform.InputValue{})
		rule := NewAwsLaunchTemplateInvalidInstanceTypeRule()

		if err = rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		if !cmp.Equal(tc.Expected, runner.Issues) {
			t.Fatalf("Expected issues are not matched:\n %s\n", cmp.Diff(tc.Expected, runner.Issues))
		}
	}
}
