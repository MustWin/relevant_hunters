// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"net/http"
	"testing"

	"github.com/MustWin/relevant_hunters/Godeps/_workspace/src/golang.org/x/net/context"
)

func TestContextClient(t *testing.T) {
	rc := &http.Client{}
	RegisterContextClientFunc(func(context.Context) (*http.Client, error) {
		return rc, nil
	})

	c := &http.Client{}
	ctx := context.WithValue(nil, HTTPClient, c)

	hc, err := ContextClient(ctx)
	if err != nil {
		t.Fatalf("want valid client; got err = %v", err)
	}
	if hc != c {
		t.Fatalf("want context client = %p; got = %p", c, hc)
	}

	hc, err = ContextClient(context.TODO())
	if err != nil {
		t.Fatalf("want valid client; got err = %v", err)
	}
	if hc != rc {
		t.Fatalf("want registered client = %p; got = %p", c, hc)
	}
}
