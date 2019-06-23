// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsServiceDiscoveryServiceInvalidDescriptionRule checks the pattern is valid
type AwsServiceDiscoveryServiceInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServiceDiscoveryServiceInvalidDescriptionRule returns new rule with default attributes
func NewAwsServiceDiscoveryServiceInvalidDescriptionRule() *AwsServiceDiscoveryServiceInvalidDescriptionRule {
	return &AwsServiceDiscoveryServiceInvalidDescriptionRule{
		resourceType:  "aws_service_discovery_service",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsServiceDiscoveryServiceInvalidDescriptionRule) Name() string {
	return "aws_service_discovery_service_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServiceDiscoveryServiceInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsServiceDiscoveryServiceInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsServiceDiscoveryServiceInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServiceDiscoveryServiceInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
