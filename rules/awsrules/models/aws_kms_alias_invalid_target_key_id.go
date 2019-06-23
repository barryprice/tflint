// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsKmsAliasInvalidTargetKeyIDRule checks the pattern is valid
type AwsKmsAliasInvalidTargetKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsKmsAliasInvalidTargetKeyIDRule returns new rule with default attributes
func NewAwsKmsAliasInvalidTargetKeyIDRule() *AwsKmsAliasInvalidTargetKeyIDRule {
	return &AwsKmsAliasInvalidTargetKeyIDRule{
		resourceType:  "aws_kms_alias",
		attributeName: "target_key_id",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Name() string {
	return "aws_kms_alias_invalid_target_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_key_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
