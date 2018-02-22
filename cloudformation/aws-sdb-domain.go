package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSSDBDomain AWS CloudFormation Resource (AWS::SDB::Domain)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type AWSSDBDomain struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html#cfn-sdb-domain-description
	Description string `json:"Description,omitempty"`

	DependsOn *[]string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSDBDomain) AWSCloudFormationType() string {
	return "AWS::SDB::Domain"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSSDBDomain) MarshalJSON() ([]byte, error) {
	type Properties AWSSDBDomain
	return json.Marshal(&struct {
		Type       string
		DependsOn  *[]string `json:",omitempty"`
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		DependsOn:  r.DependsOn,
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSSDBDomain) UnmarshalJSON(b []byte) error {
	type Properties AWSSDBDomain
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSSDBDomain(*res.Properties)
	}

	return nil
}

// GetAllAWSSDBDomainResources retrieves all AWSSDBDomain items from an AWS CloudFormation template
func (t *Template) GetAllAWSSDBDomainResources() map[string]AWSSDBDomain {
	results := map[string]AWSSDBDomain{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSSDBDomain:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SDB::Domain" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSSDBDomain
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSSDBDomainWithName retrieves all AWSSDBDomain items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSDBDomainWithName(name string) (AWSSDBDomain, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSSDBDomain:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SDB::Domain" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSSDBDomain
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSSDBDomain{}, errors.New("resource not found")
}
