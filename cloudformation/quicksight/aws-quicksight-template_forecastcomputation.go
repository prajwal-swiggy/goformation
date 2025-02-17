// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_ForecastComputation AWS CloudFormation Resource (AWS::QuickSight::Template.ForecastComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html
type Template_ForecastComputation struct {

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-computationid
	ComputationId string `json:"ComputationId"`

	// CustomSeasonalityValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-customseasonalityvalue
	CustomSeasonalityValue *float64 `json:"CustomSeasonalityValue,omitempty"`

	// LowerBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-lowerboundary
	LowerBoundary *float64 `json:"LowerBoundary,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-name
	Name *string `json:"Name,omitempty"`

	// PeriodsBackward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-periodsbackward
	PeriodsBackward *float64 `json:"PeriodsBackward,omitempty"`

	// PeriodsForward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-periodsforward
	PeriodsForward *float64 `json:"PeriodsForward,omitempty"`

	// PredictionInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-predictioninterval
	PredictionInterval *float64 `json:"PredictionInterval,omitempty"`

	// Seasonality AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-seasonality
	Seasonality *string `json:"Seasonality,omitempty"`

	// Time AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-time
	Time *Template_DimensionField `json:"Time"`

	// UpperBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-upperboundary
	UpperBoundary *float64 `json:"UpperBoundary,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastcomputation.html#cfn-quicksight-template-forecastcomputation-value
	Value *Template_MeasureField `json:"Value,omitempty"`

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
func (r *Template_ForecastComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ForecastComputation"
}
