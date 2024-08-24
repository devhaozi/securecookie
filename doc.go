// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package securecookie encodes and decodes authenticated and optionally
encrypted cookie values.

Secure cookies can't be forged, because their values are encrypted and validated
using ChaCha20-Poly1305, the content is inaccessible to malicious eyes.

To use it, first create a new SecureCookie instance:

	var key = []byte("32-byte-long-auth-key")
	var s, err = securecookie.New(key)

The key is required and must be 32 bytes, used to authenticate and encrypt cookie values.

Strong keys can be created using the convenience function GenerateRandomKey().

Once a SecureCookie instance is set, use it to encode a cookie value:

	func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
		value := map[string]string{
			"foo": "bar",
		}
		if encoded, err := s.Encode("cookie-name", value); err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
	}

Later, use the same SecureCookie instance to decode and validate a cookie
value:

	func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("cookie-name"); err == nil {
			value := make(map[string]string)
			if err = s2.Decode("cookie-name", cookie.Value, &value); err == nil {
				fmt.Fprintf(w, "The value of foo is %q", value["foo"])
			}
		}
	}

Secure cookies can hold any value thatcan be encoded using encoding/json.
it works out of the box.
*/
package securecookie
