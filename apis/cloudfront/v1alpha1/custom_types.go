/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// CustomDistributionParameters includes the custom fields of Distribution.
type CustomDistributionParameters struct {
	DistributionConfig *CustomDistributionConfig `json:"distributionConfig"`
}

// CustomCachePolicyParameters includes the custom fields of CachePolicy.
type CustomCachePolicyParameters struct{}

// CustomCloudFrontOriginAccessIDentityParameters includes the custom fields of CloudFrontOriginAccessIDentityParameters.
type CustomCloudFrontOriginAccessIDentityParameters struct{}

// CustomOrigin includes XP references or OriginAccessIdentities
type CustomOrigin struct {
	ConnectionAttempts *int64 `json:"connectionAttempts,omitempty"`

	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// A complex type that contains the list of Custom Headers for each origin.
	CustomHeaders *CustomHeaders `json:"customHeaders,omitempty"`
	// A custom origin. A custom origin is any origin that is not an Amazon S3 bucket,
	// with one exception. An Amazon S3 bucket that is configured with static website
	// hosting (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html)
	// is a custom origin.
	CustomOriginConfig *CustomOriginConfig `json:"customOriginConfig,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	ID *string `json:"id,omitempty"`

	OriginPath *string `json:"originPath,omitempty"`
	// CloudFront Origin Shield.
	//
	// Using Origin Shield can help reduce the load on your origin. For more information,
	// see Using Origin Shield (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html)
	// in the Amazon CloudFront Developer Guide.
	OriginShield *OriginShield `json:"originShield,omitempty"`
	// A complex type that contains information about the Amazon S3 origin. If the
	// origin is a custom origin or an S3 bucket that is configured as a website
	// endpoint, use the CustomOriginConfig element instead.
	S3OriginConfig *CustomS3OriginConfig `json:"s3OriginConfig,omitempty"`
}

// CustomS3OriginConfig allows a reference toa OriginAccessIdentity
type CustomS3OriginConfig struct {
	// +crossplane:generate:reference:type=OriginAccessIDentity
	OriginAccessIDentity *string `json:"originAccessIDentity,omitempty"`

	// +optional
	OriginAccessIDentityRef *xpv1.Reference `json:"originAccessIDentityRef,omitempty"`

	// +optional
	OriginAccessIDentitySelector *xpv1.Selector `json:"originAccessIDentitySelector,omitempty"`
}

// CustomOrigins includes references
type CustomOrigins struct {
	Items *[]CustomOrigin `json:"items,omitempty"`
}

// CustomDistributionConfig allows for setting a reference
type CustomDistributionConfig struct {
	// A complex type that contains information about CNAMEs (alternate domain names),
	// if any, for this distribution.
	Aliases *Aliases `json:"aliases,omitempty"`
	// A complex type that contains zero or more CacheBehavior elements.
	CacheBehaviors *CacheBehaviors `json:"cacheBehaviors,omitempty"`

	Comment *string `json:"comment,omitempty"`
	// A complex type that controls:
	//
	//    * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range
	//    with custom error messages before returning the response to the viewer.
	//
	//    * How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
	//
	// For more information about custom error pages, see Customizing Error Responses
	// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html)
	// in the Amazon CloudFront Developer Guide.
	CustomErrorResponses *CustomErrorResponses `json:"customErrorResponses,omitempty"`
	// A complex type that describes the default cache behavior if you don’t specify
	// a CacheBehavior element or if request URLs don’t match any of the values
	// of PathPattern in CacheBehavior elements. You must create exactly one default
	// cache behavior.
	DefaultCacheBehavior *DefaultCacheBehavior `json:"defaultCacheBehavior,omitempty"`

	DefaultRootObject *string `json:"defaultRootObject,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	HTTPVersion *string `json:"httpVersion,omitempty"`

	IsIPV6Enabled *bool `json:"isIPV6Enabled,omitempty"`
	// A complex type that controls whether access logs are written for the distribution.
	Logging *LoggingConfig `json:"logging,omitempty"`
	// A complex data type for the origin groups specified for a distribution.
	OriginGroups *OriginGroups `json:"originGroups,omitempty"`
	// Contains information about the origins for this distribution.
	Origins *CustomOrigins `json:"origins,omitempty"`

	PriceClass *string `json:"priceClass,omitempty"`
	// A complex type that identifies ways in which you want to restrict distribution
	// of your content.
	Restrictions *Restrictions `json:"restrictions,omitempty"`
	// A complex type that determines the distribution’s SSL/TLS configuration
	// for communicating with viewers.
	//
	// If the distribution doesn’t use Aliases (also known as alternate domain
	// names or CNAMEs)—that is, if the distribution uses the CloudFront domain
	// name such as d111111abcdef8.cloudfront.net—set CloudFrontDefaultCertificate
	// to true and leave all other fields empty.
	//
	// If the distribution uses Aliases (alternate domain names or CNAMEs), use
	// the fields in this type to specify the following settings:
	//
	//    * Which viewers the distribution accepts HTTPS connections from: only
	//    viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication)
	//    (recommended), or all viewers including those that don’t support SNI.
	//    To accept HTTPS connections from only viewers that support SNI, set SSLSupportMethod
	//    to sni-only. This is recommended. Most browsers and clients support SNI.
	//    To accept HTTPS connections from all viewers, including those that don’t
	//    support SNI, set SSLSupportMethod to vip. This is not recommended, and
	//    results in additional monthly charges from CloudFront.
	//
	//    * The minimum SSL/TLS protocol version that the distribution can use to
	//    communicate with viewers. To specify a minimum version, choose a value
	//    for MinimumProtocolVersion. For more information, see Security Policy
	//    (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy)
	//    in the Amazon CloudFront Developer Guide.
	//
	//    * The location of the SSL/TLS certificate, AWS Certificate Manager (ACM)
	//    (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) (recommended)
	//    or AWS Identity and Access Management (AWS IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html).
	//    You specify the location by setting a value in one of the following fields
	//    (not both): ACMCertificateArn IAMCertificateId
	//
	// All distributions support HTTPS connections from viewers. To require viewers
	// to use HTTPS only, or to redirect them from HTTP to HTTPS, use ViewerProtocolPolicy
	// in the CacheBehavior or DefaultCacheBehavior. To specify how CloudFront should
	// use SSL/TLS to communicate with your custom origin, use CustomOriginConfig.
	//
	// For more information, see Using HTTPS with CloudFront (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https.html)
	// and Using Alternate Domain Names and HTTPS (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-alternate-domain-names.html)
	// in the Amazon CloudFront Developer Guide.
	ViewerCertificate *ViewerCertificate `json:"viewerCertificate,omitempty"`

	WebACLID *string `json:"webACLID,omitempty"`
}
