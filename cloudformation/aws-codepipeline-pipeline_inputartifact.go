package cloudformation

// AWSCodePipelinePipeline_InputArtifact AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.InputArtifact)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-inputartifact.html
type AWSCodePipelinePipeline_InputArtifact struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-inputartifact.html#cfn-codepipeline-pipeline-inputartifact-name
	Name string `json:"Name,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_InputArtifact) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.InputArtifact"
}
