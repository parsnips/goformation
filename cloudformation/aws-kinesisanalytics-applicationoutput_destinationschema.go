package cloudformation

// AWSKinesisAnalyticsApplicationOutput_DestinationSchema AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html
type AWSKinesisAnalyticsApplicationOutput_DestinationSchema struct {

	// RecordFormatType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html#cfn-kinesisanalytics-applicationoutput-destinationschema-recordformattype
	RecordFormatType string `json:"RecordFormatType,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_DestinationSchema) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema"
}
