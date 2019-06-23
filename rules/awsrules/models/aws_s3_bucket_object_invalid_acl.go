// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsS3BucketObjectInvalidACLRule checks the pattern is valid
type AwsS3BucketObjectInvalidACLRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketObjectInvalidACLRule returns new rule with default attributes
func NewAwsS3BucketObjectInvalidACLRule() *AwsS3BucketObjectInvalidACLRule {
	return &AwsS3BucketObjectInvalidACLRule{
		resourceType:  "aws_s3_bucket_object",
		attributeName: "acl",
		enum: []string{
			"private",
			"public-read",
			"public-read-write",
			"authenticated-read",
			"aws-exec-read",
			"bucket-owner-read",
			"bucket-owner-full-control",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketObjectInvalidACLRule) Name() string {
	return "aws_s3_bucket_object_invalid_acl"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketObjectInvalidACLRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsS3BucketObjectInvalidACLRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketObjectInvalidACLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketObjectInvalidACLRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`acl is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
