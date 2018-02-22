package cloudformation

// AWSCodePipelinePipeline_ActionTypeId AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.ActionTypeId)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html
type AWSCodePipelinePipeline_ActionTypeId struct {

	// Category AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-category
	Category string `json:"Category,omitempty"`

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-owner
	Owner string `json:"Owner,omitempty"`

	// Provider AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-provider
	Provider string `json:"Provider,omitempty"`

	// Version AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-version
	Version string `json:"Version,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ActionTypeId) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ActionTypeId"
}
