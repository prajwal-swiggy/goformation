// Code generated by "go generate". Please don't change this file directly.

package securityhub

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// AutomationRule_SeverityUpdate AWS CloudFormation Resource (AWS::SecurityHub::AutomationRule.SeverityUpdate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html
type AutomationRule_SeverityUpdate struct {

	// Label AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-label
	Label *string `json:"Label,omitempty"`

	// Normalized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-normalized
	Normalized *int `json:"Normalized,omitempty"`

	// Product AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-severityupdate.html#cfn-securityhub-automationrule-severityupdate-product
	Product *float64 `json:"Product,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AutomationRule_SeverityUpdate) AWSCloudFormationType() string {
	return "AWS::SecurityHub::AutomationRule.SeverityUpdate"
}
