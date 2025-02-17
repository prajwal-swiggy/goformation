// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_AggregationFunction AWS CloudFormation Resource (AWS::QuickSight::Dashboard.AggregationFunction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationfunction.html
type Dashboard_AggregationFunction struct {

	// CategoricalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationfunction.html#cfn-quicksight-dashboard-aggregationfunction-categoricalaggregationfunction
	CategoricalAggregationFunction *string `json:"CategoricalAggregationFunction,omitempty"`

	// DateAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationfunction.html#cfn-quicksight-dashboard-aggregationfunction-dateaggregationfunction
	DateAggregationFunction *string `json:"DateAggregationFunction,omitempty"`

	// NumericalAggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationfunction.html#cfn-quicksight-dashboard-aggregationfunction-numericalaggregationfunction
	NumericalAggregationFunction *Dashboard_NumericalAggregationFunction `json:"NumericalAggregationFunction,omitempty"`

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
func (r *Dashboard_AggregationFunction) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.AggregationFunction"
}
