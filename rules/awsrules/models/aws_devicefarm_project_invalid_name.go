// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsDevicefarmProjectInvalidNameRule checks the pattern is valid
type AwsDevicefarmProjectInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsDevicefarmProjectInvalidNameRule returns new rule with default attributes
func NewAwsDevicefarmProjectInvalidNameRule() *AwsDevicefarmProjectInvalidNameRule {
	return &AwsDevicefarmProjectInvalidNameRule{
		resourceType:  "aws_devicefarm_project",
		attributeName: "name",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsDevicefarmProjectInvalidNameRule) Name() string {
	return "aws_devicefarm_project_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmProjectInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsDevicefarmProjectInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmProjectInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmProjectInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
