package resources

// AWS::AutoScaling::AutoScalingGroup.MetricsCollection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AWSAutoScalingAutoScalingGroupMetricsCollection struct {

	// Granularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-granularity
	Granularity string `json:"Granularity"`

	// Metrics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-metrics
	Metrics []AWSAutoScalingAutoScalingGroupMetricsCollectionstring `json:"Metrics"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroupMetricsCollection) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.MetricsCollection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroupMetricsCollection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}