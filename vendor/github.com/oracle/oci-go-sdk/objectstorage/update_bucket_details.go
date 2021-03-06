// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateBucketDetails To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type UpdateBucketDetails struct {

	// The namespace in which the bucket lives.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The name of the bucket.
	Name *string `mandatory:"false" json:"name"`

	// Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// The type of public access available on this bucket. Allows authenticated caller to access the bucket or
	// contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess
	// when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access.
	// When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects.
	PublicAccessType UpdateBucketDetailsPublicAccessTypeEnum `mandatory:"false" json:"publicAccessType,omitempty"`
}

func (m UpdateBucketDetails) String() string {
	return common.PointerString(m)
}

// UpdateBucketDetailsPublicAccessTypeEnum Enum with underlying type: string
type UpdateBucketDetailsPublicAccessTypeEnum string

// Set of constants representing the allowable values for UpdateBucketDetailsPublicAccessType
const (
	UpdateBucketDetailsPublicAccessTypeNopublicaccess UpdateBucketDetailsPublicAccessTypeEnum = "NoPublicAccess"
	UpdateBucketDetailsPublicAccessTypeObjectread     UpdateBucketDetailsPublicAccessTypeEnum = "ObjectRead"
)

var mappingUpdateBucketDetailsPublicAccessType = map[string]UpdateBucketDetailsPublicAccessTypeEnum{
	"NoPublicAccess": UpdateBucketDetailsPublicAccessTypeNopublicaccess,
	"ObjectRead":     UpdateBucketDetailsPublicAccessTypeObjectread,
}

// GetUpdateBucketDetailsPublicAccessTypeEnumValues Enumerates the set of values for UpdateBucketDetailsPublicAccessType
func GetUpdateBucketDetailsPublicAccessTypeEnumValues() []UpdateBucketDetailsPublicAccessTypeEnum {
	values := make([]UpdateBucketDetailsPublicAccessTypeEnum, 0)
	for _, v := range mappingUpdateBucketDetailsPublicAccessType {
		values = append(values, v)
	}
	return values
}
