package cloudformation

// AWSCloudFrontDistribution_LambdaFunctionAssociation AWS CloudFormation Resource (AWS::CloudFront::Distribution.LambdaFunctionAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-lambdafunctionassociation.html
type AWSCloudFrontDistribution_LambdaFunctionAssociation struct {

	// EventType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-lambdafunctionassociation.html#cfn-cloudfront-distribution-lambdafunctionassociation-eventtype
	EventType string `json:"EventType,omitempty"`

	// LambdaFunctionARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-lambdafunctionassociation.html#cfn-cloudfront-distribution-lambdafunctionassociation-lambdafunctionarn
	LambdaFunctionARN string `json:"LambdaFunctionARN,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_LambdaFunctionAssociation) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.LambdaFunctionAssociation"
}
