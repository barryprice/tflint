// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule checks the pattern is valid
type AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAPIGatewayIntegrationResponseInvalidStatusCodeRule returns new rule with default attributes
func NewAwsAPIGatewayIntegrationResponseInvalidStatusCodeRule() *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule {
	return &AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule{
		resourceType:  "aws_api_gateway_integration_response",
		attributeName: "status_code",
		pattern:       regexp.MustCompile(`^[1-5]\d\d$`),
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule) Name() string {
	return "aws_api_gateway_integration_response_invalid_status_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayIntegrationResponseInvalidStatusCodeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`status_code does not match valid pattern ^[1-5]\d\d$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
