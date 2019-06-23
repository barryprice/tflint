// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsResourcegroupsGroupInvalidDescriptionRule checks the pattern is valid
type AwsResourcegroupsGroupInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsResourcegroupsGroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsResourcegroupsGroupInvalidDescriptionRule() *AwsResourcegroupsGroupInvalidDescriptionRule {
	return &AwsResourcegroupsGroupInvalidDescriptionRule{
		resourceType:  "aws_resourcegroups_group",
		attributeName: "description",
		max:           512,
		pattern:       regexp.MustCompile(`^[\sa-zA-Z0-9_\.-]*$`),
	}
}

// Name returns the rule name
func (r *AwsResourcegroupsGroupInvalidDescriptionRule) Name() string {
	return "aws_resourcegroups_group_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsResourcegroupsGroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsResourcegroupsGroupInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsResourcegroupsGroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsResourcegroupsGroupInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`description does not match valid pattern ^[\sa-zA-Z0-9_\.-]*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
