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

package mass_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/message/mass"
)

func ExampleMediaUploadNews() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.MediaUploadNews(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSendAll() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.SendAll(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMediaUploadVideo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.MediaUploadVideo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePreview() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.Preview(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpeedGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.SpeedGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpeedSet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := mass.SpeedSet(ctx, payload)

	fmt.Println(resp, err)
}
