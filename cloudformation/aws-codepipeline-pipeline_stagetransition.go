package cloudformation

// AWSCodePipelinePipeline_StageTransition AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.StageTransition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html
type AWSCodePipelinePipeline_StageTransition struct {

	// Reason AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html#cfn-codepipeline-pipeline-stagetransition-reason
	Reason string `json:"Reason,omitempty"`

	// StageName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html#cfn-codepipeline-pipeline-stagetransition-stagename
	StageName string `json:"StageName,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_StageTransition) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.StageTransition"
}
