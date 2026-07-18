package appauth_test

import (
	"fmt"
	"log"

	"github.com/matthiasbruns/ecwid-go/ecwid/appauth"
)

// ExampleDecrypt decodes an Enhanced Security User Auth payload. In enhanced mode
// Ecwid places a URL-safe base64 blob in the ?payload= query parameter, so a Go
// backend can read it directly — unlike default mode, whose hex payload sits in
// the URL fragment and never reaches the server.
func ExampleDecrypt() {
	// The ?payload= blob Ecwid delivers: AES-128-CBC, keyed by the first 16 bytes
	// of your app's client_secret.
	const payload = "au4MxpzDO1oHhk35oESEDbMfLVK7L05p8VGUmmyh7Kue_rmCJEnNeoLsx3M4UhfmbuDjK7CRQ3WR61pDy0RGRQbrjXqxhjEWM82DjClZXJRat5IZg66b0rMrx1u7J8vMKYqSzmBZzjoszsF5BXTqBCmHeweqnWurIkdI-gzMk9qdID-DxSlpO725KS-_VaFzihsnfCcd9Jmgxw49NzDfyA=="
	const clientSecret = "0123456789abcdef0123456789abcdef"

	p, err := appauth.Decrypt(payload, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("store:", p.StoreID)
	fmt.Println("lang:", p.Lang)
	// %v routes through Payload.Format, which redacts access_token and
	// public_token, so a Payload is safe to print and log.
	fmt.Printf("payload: %v\n", p)

	// Output:
	// store: 1003
	// lang: en
	// payload: appauth.Payload{StoreID:1003 Lang:"en" AccessToken:"***************f456" PublicToken:"" ViewMode:"PAGE" AppState:"" Domain:""}
}

// ExampleDecodeHex decodes a Default User Auth payload: the hex-encoded plaintext
// JSON Ecwid puts in the URL fragment. This is what every app gets unless Ecwid
// switches it to enhanced mode. Because a fragment never reaches the server, this
// runs client-side (via EcwidApp.getPayload); DecodeHex exists for tests and for
// tooling that receives the fragment out of band.
func ExampleDecodeHex() {
	// The fragment content, without the leading '#'.
	const fragment = "7b2273746f72655f6964223a313030332c226c616e67223a22656e222c226163636573735f746f6b656e223a227365637265745f616263313233646566343536222c227075626c69635f746f6b656e223a22222c22766965775f6d6f6465223a2250414745222c226170705f7374617465223a22222c22646f6d61696e223a22227d"

	p, err := appauth.DecodeHex(fragment)
	if err != nil {
		log.Fatal(err)
	}

	// A hex payload is unauthenticated plaintext — re-validate access_token against
	// the Ecwid API before trusting anything it carries.
	fmt.Println("store:", p.StoreID)
	fmt.Println("view mode:", p.ViewMode)

	// Output:
	// store: 1003
	// view mode: PAGE
}
