package cloudformation

// AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration AWS CloudFormation Resource (AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html
type AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-applicationname
	ApplicationName string `json:"ApplicationName,omitempty"`

	// TemplateName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-templatename
	TemplateName string `json:"TemplateName,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration"
}
