// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package logging

import (
	"reflect"
	"strings"
	"testing"
)

// AssertLoggedMessages asserts the logged messages dispatched by a mocked logger
func AssertLoggedMessages(t *testing.T, mockLogger Dispatcher, wantMessages []string) {
	if mockLogger == nil {
		return
	}

	mockDispatcher := mockLogger.(*MockDispatcher)
	messagesLogged := mockDispatcher.GetDispatchedMessages()
	if len(wantMessages) > 0 && len(messagesLogged) == 0 {
		if !reflect.DeepEqual(wantMessages, messagesLogged) {
			t.Errorf("%s expected logged messages %v but got \n\nnone", t.Name(),
				strings.Join(wantMessages, "\n"))
		}
	}

	if len(wantMessages) == 0 && len(messagesLogged) > 0 {
		if !reflect.DeepEqual(wantMessages, messagesLogged) {
			t.Errorf("%s expected no logged messages but got \n\n%v", t.Name(),
				strings.Join(messagesLogged, "\n"))
		}
	}
	if len(wantMessages) > 0 && len(messagesLogged) > 0 {
		if len(wantMessages) < len(messagesLogged) {
			t.Errorf("%s more messages logged [%d] than expected [%d] \n\n additional messages logged:\n%+v",
				t.Name(),
				len(messagesLogged),
				len(wantMessages),
				strings.Join(messagesLogged[len(wantMessages):], "\n"))
		}
		for index, m := range wantMessages {
			if len(messagesLogged) <= index {
				t.Errorf("%s expected logged message [%d] \n%v \n\nbut got none",
					t.Name(),
					index,
					m)
				return
			}
			if messagesLogged[index] != m {
				t.Errorf("%s expected logged message [%d] \n%v \n\nbut got \n\n%v",
					t.Name(),
					index,
					m,
					messagesLogged[index])
			}
		}
	}
}
