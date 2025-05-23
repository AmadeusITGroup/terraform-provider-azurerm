package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoDbLinkedServiceTypeProperties struct {
	AllowSelfSignedServerCert *bool                      `json:"allowSelfSignedServerCert,omitempty"`
	AuthSource                *interface{}               `json:"authSource,omitempty"`
	AuthenticationType        *MongoDbAuthenticationType `json:"authenticationType,omitempty"`
	DatabaseName              interface{}                `json:"databaseName"`
	EnableSsl                 *bool                      `json:"enableSsl,omitempty"`
	EncryptedCredential       *string                    `json:"encryptedCredential,omitempty"`
	Password                  SecretBase                 `json:"password"`
	Port                      *int64                     `json:"port,omitempty"`
	Server                    interface{}                `json:"server"`
	Username                  *interface{}               `json:"username,omitempty"`
}

var _ json.Unmarshaler = &MongoDbLinkedServiceTypeProperties{}

func (s *MongoDbLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowSelfSignedServerCert *bool                      `json:"allowSelfSignedServerCert,omitempty"`
		AuthSource                *interface{}               `json:"authSource,omitempty"`
		AuthenticationType        *MongoDbAuthenticationType `json:"authenticationType,omitempty"`
		DatabaseName              interface{}                `json:"databaseName"`
		EnableSsl                 *bool                      `json:"enableSsl,omitempty"`
		EncryptedCredential       *string                    `json:"encryptedCredential,omitempty"`
		Port                      *int64                     `json:"port,omitempty"`
		Server                    interface{}                `json:"server"`
		Username                  *interface{}               `json:"username,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowSelfSignedServerCert = decoded.AllowSelfSignedServerCert
	s.AuthSource = decoded.AuthSource
	s.AuthenticationType = decoded.AuthenticationType
	s.DatabaseName = decoded.DatabaseName
	s.EnableSsl = decoded.EnableSsl
	s.EncryptedCredential = decoded.EncryptedCredential
	s.Port = decoded.Port
	s.Server = decoded.Server
	s.Username = decoded.Username

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MongoDbLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'MongoDbLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}

	return nil
}
