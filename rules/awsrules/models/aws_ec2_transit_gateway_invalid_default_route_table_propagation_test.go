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

func Test_AwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected issue.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	default_route_table_propagation = "disabled"
}`,
			Expected: []*issue.Issue{
				{
					Detector: "aws_ec2_transit_gateway_invalid_default_route_table_propagation",
					Type:     "ERROR",
					Message:  `default_route_table_propagation is not a valid value`,
					Line:     3,
					File:     "resource.tf",
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	default_route_table_propagation = "disable"
}`,
			Expected: []*issue.Issue{},
		},
	}

	dir, err := ioutil.TempDir("", "AwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule")
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
		rule := NewAwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule()

		if err = rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		if !cmp.Equal(tc.Expected, runner.Issues) {
			t.Fatalf("Expected issues are not matched:\n %s\n", cmp.Diff(tc.Expected, runner.Issues))
		}
	}
}
