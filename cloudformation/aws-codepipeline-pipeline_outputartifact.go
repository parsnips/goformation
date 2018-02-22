package cloudformation

// AWSCodePipelinePipeline_OutputArtifact AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.OutputArtifact)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-outputartifact.html
type AWSCodePipelinePipeline_OutputArtifact struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-outputartifact.html#cfn-codepipeline-pipeline-outputartifact-name
	Name string `json:"Name,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_OutputArtifact) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.OutputArtifact"
}
