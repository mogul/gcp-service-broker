// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import "testing"

type osbNameStruct struct {
	Name string `validate:"osbname"`
}

type jsonStruct struct {
	Json []byte `validate:"json"`
}

type jsonStringStruct struct {
	Json string `validate:"json"`
}

func TestValidateStruct(t *testing.T) {
	cases := map[string]struct {
		Validate  interface{}
		ExpectErr bool
	}{
		"osbname missing": {
			Validate:  osbNameStruct{},
			ExpectErr: true,
		},
		"osbname valid": {
			Validate:  osbNameStruct{Name: "google-storage"},
			ExpectErr: false,
		},
		"osbname spaces": {
			Validate:  osbNameStruct{Name: " google-storage  "},
			ExpectErr: true,
		},
		"osbname dots": {
			Validate:  osbNameStruct{Name: "google.storage"},
			ExpectErr: false,
		},
		"osbname alpha": {
			Validate:  osbNameStruct{Name: "googlestorage"},
			ExpectErr: false,
		},
		"osbname upper": {
			Validate:  osbNameStruct{Name: "GOOGLESTORAGE"},
			ExpectErr: false,
		},
		"osbname numeric": {
			Validate:  osbNameStruct{Name: "12345"},
			ExpectErr: false,
		},
		"json bytes blank": {
			Validate:  jsonStruct{},
			ExpectErr: true,
		},
		"json bytes empty object": {
			Validate:  jsonStruct{Json: []byte("{}")},
			ExpectErr: false,
		},
		"json bytes bad object": {
			Validate:  jsonStruct{Json: []byte("")},
			ExpectErr: true,
		},
		"json bytes full object": {
			Validate:  jsonStruct{Json: []byte(`{"a":42, "s":"foo"}`)},
			ExpectErr: false,
		},
		"json string blank": {
			Validate:  jsonStringStruct{},
			ExpectErr: true,
		},
		"json string empty object": {
			Validate:  jsonStringStruct{Json: "{}"},
			ExpectErr: false,
		},
		"json string bad object": {
			Validate:  jsonStringStruct{Json: ""},
			ExpectErr: true,
		},
		"json string full object": {
			Validate:  jsonStringStruct{Json: `{"a":42, "s":"foo"}`},
			ExpectErr: false,
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			err := ValidateStruct(tc.Validate)
			gotErr := err != nil
			if gotErr != tc.ExpectErr {
				t.Errorf("expected error? %v got %v", tc.ExpectErr, err)
			}
		})
	}
}