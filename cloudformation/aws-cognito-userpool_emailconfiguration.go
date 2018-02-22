package cloudformation

// AWSCognitoUserPool_EmailConfiguration AWS CloudFormation Resource (AWS::Cognito::UserPool.EmailConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html
type AWSCognitoUserPool_EmailConfiguration struct {

	// ReplyToEmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-replytoemailaddress
	ReplyToEmailAddress string `json:"ReplyToEmailAddress,omitempty"`

	// SourceArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-sourcearn
	SourceArn string `json:"SourceArn,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_EmailConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.EmailConfiguration"
}
