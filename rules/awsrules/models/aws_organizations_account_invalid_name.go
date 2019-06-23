// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsAccountInvalidNameRule checks the pattern is valid
type AwsOrganizationsAccountInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsAccountInvalidNameRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidNameRule() *AwsOrganizationsAccountInvalidNameRule {
	return &AwsOrganizationsAccountInvalidNameRule{
		resourceType:  "aws_organizations_account",
		attributeName: "name",
		max:           50,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0020}-\x{007E}]+$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidNameRule) Name() string {
	return "aws_organizations_account_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsAccountInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[\x{0020}-\x{007E}]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
