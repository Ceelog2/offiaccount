// Copyright 2020 FastWeGo
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

package membercard_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/card/membercard"
)

func ExampleActivateGetUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.ActivateGetUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleActivateTempInfoGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.ActivateTempInfoGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleActivate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.Activate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleActivateUserFormSet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.ActivateUserFormSet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserinfoGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.UserinfoGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateUser() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.UpdateUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePayGiftcardAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.PayGiftcardAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePayGiftcardDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.PayGiftcardDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePayGiftcardGetById() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.PayGiftcardGetById(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePayGiftcardBatchGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := membercard.PayGiftcardBatchGet(ctx, payload)

	fmt.Println(resp, err)
}
