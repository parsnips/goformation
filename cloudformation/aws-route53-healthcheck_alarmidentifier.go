package cloudformation

// AWSRoute53HealthCheck_AlarmIdentifier AWS CloudFormation Resource (AWS::Route53::HealthCheck.AlarmIdentifier)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html
type AWSRoute53HealthCheck_AlarmIdentifier struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-name
	Name string `json:"Name,omitempty"`

	// Region AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-region
	Region string `json:"Region,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheck_AlarmIdentifier) AWSCloudFormationType() string {
	return "AWS::Route53::HealthCheck.AlarmIdentifier"
}
