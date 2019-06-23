// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule() *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "target_name",
		max:           200,
		min:           1,
		pattern:       regexp.MustCompile(`^[-\.;a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_target_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_name must be 200 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`target_name does not match valid pattern ^[-\.;a-z0-9]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
