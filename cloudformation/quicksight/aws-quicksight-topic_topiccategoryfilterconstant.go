// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Topic_TopicCategoryFilterConstant AWS CloudFormation Resource (AWS::QuickSight::Topic.TopicCategoryFilterConstant)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html
type Topic_TopicCategoryFilterConstant struct {

	// CollectiveConstant AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-collectiveconstant
	CollectiveConstant *Topic_CollectiveConstant `json:"CollectiveConstant,omitempty"`

	// ConstantType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-constanttype
	ConstantType *string `json:"ConstantType,omitempty"`

	// SingularConstant AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-singularconstant
	SingularConstant *string `json:"SingularConstant,omitempty"`

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
func (r *Topic_TopicCategoryFilterConstant) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.TopicCategoryFilterConstant"
}
