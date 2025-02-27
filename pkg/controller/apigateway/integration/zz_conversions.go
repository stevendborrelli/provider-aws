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

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetIntegrationInput returns input for read
// operation.
func GenerateGetIntegrationInput(cr *svcapitypes.Integration) *svcsdk.GetIntegrationInput {
	res := &svcsdk.GetIntegrationInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}

	return res
}

// GenerateIntegration returns the current state in the form of *svcapitypes.Integration.
func GenerateIntegration(resp *svcsdk.Integration) *svcapitypes.Integration {
	cr := &svcapitypes.Integration{}

	if resp.CacheKeyParameters != nil {
		f0 := []*string{}
		for _, f0iter := range resp.CacheKeyParameters {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		cr.Spec.ForProvider.CacheKeyParameters = f0
	} else {
		cr.Spec.ForProvider.CacheKeyParameters = nil
	}
	if resp.CacheNamespace != nil {
		cr.Spec.ForProvider.CacheNamespace = resp.CacheNamespace
	} else {
		cr.Spec.ForProvider.CacheNamespace = nil
	}
	if resp.ConnectionId != nil {
		cr.Spec.ForProvider.ConnectionID = resp.ConnectionId
	} else {
		cr.Spec.ForProvider.ConnectionID = nil
	}
	if resp.ConnectionType != nil {
		cr.Spec.ForProvider.ConnectionType = resp.ConnectionType
	} else {
		cr.Spec.ForProvider.ConnectionType = nil
	}
	if resp.ContentHandling != nil {
		cr.Spec.ForProvider.ContentHandling = resp.ContentHandling
	} else {
		cr.Spec.ForProvider.ContentHandling = nil
	}
	if resp.Credentials != nil {
		cr.Spec.ForProvider.Credentials = resp.Credentials
	} else {
		cr.Spec.ForProvider.Credentials = nil
	}
	if resp.HttpMethod != nil {
		cr.Spec.ForProvider.HTTPMethod = resp.HttpMethod
	} else {
		cr.Spec.ForProvider.HTTPMethod = nil
	}
	if resp.PassthroughBehavior != nil {
		cr.Spec.ForProvider.PassthroughBehavior = resp.PassthroughBehavior
	} else {
		cr.Spec.ForProvider.PassthroughBehavior = nil
	}
	if resp.RequestParameters != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range resp.RequestParameters {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		cr.Spec.ForProvider.RequestParameters = f8
	} else {
		cr.Spec.ForProvider.RequestParameters = nil
	}
	if resp.RequestTemplates != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range resp.RequestTemplates {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		cr.Spec.ForProvider.RequestTemplates = f9
	} else {
		cr.Spec.ForProvider.RequestTemplates = nil
	}
	if resp.TimeoutInMillis != nil {
		cr.Spec.ForProvider.TimeoutInMillis = resp.TimeoutInMillis
	} else {
		cr.Spec.ForProvider.TimeoutInMillis = nil
	}
	if resp.TlsConfig != nil {
		f11 := &svcapitypes.TLSConfig{}
		if resp.TlsConfig.InsecureSkipVerification != nil {
			f11.InsecureSkipVerification = resp.TlsConfig.InsecureSkipVerification
		}
		cr.Spec.ForProvider.TLSConfig = f11
	} else {
		cr.Spec.ForProvider.TLSConfig = nil
	}
	if resp.Type != nil {
		cr.Spec.ForProvider.Type = resp.Type
	} else {
		cr.Spec.ForProvider.Type = nil
	}
	if resp.Uri != nil {
		cr.Spec.ForProvider.URI = resp.Uri
	} else {
		cr.Spec.ForProvider.URI = nil
	}

	return cr
}

// GeneratePutIntegrationInput returns a create input.
func GeneratePutIntegrationInput(cr *svcapitypes.Integration) *svcsdk.PutIntegrationInput {
	res := &svcsdk.PutIntegrationInput{}

	if cr.Spec.ForProvider.CacheKeyParameters != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.CacheKeyParameters {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetCacheKeyParameters(f0)
	}
	if cr.Spec.ForProvider.CacheNamespace != nil {
		res.SetCacheNamespace(*cr.Spec.ForProvider.CacheNamespace)
	}
	if cr.Spec.ForProvider.ConnectionID != nil {
		res.SetConnectionId(*cr.Spec.ForProvider.ConnectionID)
	}
	if cr.Spec.ForProvider.ConnectionType != nil {
		res.SetConnectionType(*cr.Spec.ForProvider.ConnectionType)
	}
	if cr.Spec.ForProvider.ContentHandling != nil {
		res.SetContentHandling(*cr.Spec.ForProvider.ContentHandling)
	}
	if cr.Spec.ForProvider.Credentials != nil {
		res.SetCredentials(*cr.Spec.ForProvider.Credentials)
	}
	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}
	if cr.Spec.ForProvider.IntegrationHTTPMethod != nil {
		res.SetIntegrationHttpMethod(*cr.Spec.ForProvider.IntegrationHTTPMethod)
	}
	if cr.Spec.ForProvider.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*cr.Spec.ForProvider.PassthroughBehavior)
	}
	if cr.Spec.ForProvider.RequestParameters != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range cr.Spec.ForProvider.RequestParameters {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetRequestParameters(f9)
	}
	if cr.Spec.ForProvider.RequestTemplates != nil {
		f10 := map[string]*string{}
		for f10key, f10valiter := range cr.Spec.ForProvider.RequestTemplates {
			var f10val string
			f10val = *f10valiter
			f10[f10key] = &f10val
		}
		res.SetRequestTemplates(f10)
	}
	if cr.Spec.ForProvider.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*cr.Spec.ForProvider.TimeoutInMillis)
	}
	if cr.Spec.ForProvider.TLSConfig != nil {
		f12 := &svcsdk.TlsConfig{}
		if cr.Spec.ForProvider.TLSConfig.InsecureSkipVerification != nil {
			f12.SetInsecureSkipVerification(*cr.Spec.ForProvider.TLSConfig.InsecureSkipVerification)
		}
		res.SetTlsConfig(f12)
	}
	if cr.Spec.ForProvider.Type != nil {
		res.SetType(*cr.Spec.ForProvider.Type)
	}
	if cr.Spec.ForProvider.URI != nil {
		res.SetUri(*cr.Spec.ForProvider.URI)
	}

	return res
}

// GenerateUpdateIntegrationInput returns an update input.
func GenerateUpdateIntegrationInput(cr *svcapitypes.Integration) *svcsdk.UpdateIntegrationInput {
	res := &svcsdk.UpdateIntegrationInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}

	return res
}

// GenerateDeleteIntegrationInput returns a deletion input.
func GenerateDeleteIntegrationInput(cr *svcapitypes.Integration) *svcsdk.DeleteIntegrationInput {
	res := &svcsdk.DeleteIntegrationInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
