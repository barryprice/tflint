// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule checks the pattern is valid
type AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCognitoUserPoolInvalidAutoVerifiedAttributesRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidAutoVerifiedAttributesRule() *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule {
	return &AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "auto_verified_attributes",
		enum: []string{
			"phone_number",
			"email",
		},
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule) Name() string {
	return "aws_cognito_user_pool_invalid_auto_verified_attributes"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidAutoVerifiedAttributesRule) Check(runner *tflint.Runner) error {
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
					`auto_verified_attributes is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
