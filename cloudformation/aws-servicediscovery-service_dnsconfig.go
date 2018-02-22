package cloudformation

// AWSServiceDiscoveryService_DnsConfig AWS CloudFormation Resource (AWS::ServiceDiscovery::Service.DnsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html
type AWSServiceDiscoveryService_DnsConfig struct {

	// DnsRecords AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-dnsrecords
	DnsRecords []AWSServiceDiscoveryService_DnsRecord `json:"DnsRecords,omitempty"`

	// NamespaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-namespaceid
	NamespaceId string `json:"NamespaceId,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryService_DnsConfig) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::Service.DnsConfig"
}
