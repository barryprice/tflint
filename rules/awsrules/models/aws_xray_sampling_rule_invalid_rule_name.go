// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsXraySamplingRuleInvalidRuleNameRule checks the pattern is valid
type AwsXraySamplingRuleInvalidRuleNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsXraySamplingRuleInvalidRuleNameRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidRuleNameRule() *AwsXraySamplingRuleInvalidRuleNameRule {
	return &AwsXraySamplingRuleInvalidRuleNameRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "rule_name",
		max:           32,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidRuleNameRule) Name() string {
	return "aws_xray_sampling_rule_invalid_rule_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidRuleNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsXraySamplingRuleInvalidRuleNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidRuleNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidRuleNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"rule_name must be 32 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"rule_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
