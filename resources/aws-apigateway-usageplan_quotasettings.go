package resources


// AWSApiGatewayUsagePlan_QuotaSettings AWS CloudFormation Resource (AWS::ApiGateway::UsagePlan.QuotaSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
type AWSApiGatewayUsagePlan_QuotaSettings struct {
    
    // Limit AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
    Limit int `json:"Limit,omitempty"`
    
    // Offset AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
    Offset int `json:"Offset,omitempty"`
    
    // Period AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
    Period string `json:"Period,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AWSCloudFormationType() string {
    return "AWS::ApiGateway::UsagePlan.QuotaSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}