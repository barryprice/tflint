// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsRoute53QueryLogInvalidZoneIDRule checks the pattern is valid
type AwsRoute53QueryLogInvalidZoneIDRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53QueryLogInvalidZoneIDRule returns new rule with default attributes
func NewAwsRoute53QueryLogInvalidZoneIDRule() *AwsRoute53QueryLogInvalidZoneIDRule {
	return &AwsRoute53QueryLogInvalidZoneIDRule{
		resourceType:  "aws_route53_query_log",
		attributeName: "zone_id",
		max:           32,
	}
}

// Name returns the rule name
func (r *AwsRoute53QueryLogInvalidZoneIDRule) Name() string {
	return "aws_route53_query_log_invalid_zone_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53QueryLogInvalidZoneIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsRoute53QueryLogInvalidZoneIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53QueryLogInvalidZoneIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53QueryLogInvalidZoneIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"zone_id must be 32 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
