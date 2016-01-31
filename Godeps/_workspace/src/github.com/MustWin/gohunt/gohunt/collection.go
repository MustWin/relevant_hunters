// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Collection struct {
	ID            int               `json:"id,"`
	Name          string            `json:"name"`
	Title string            `json:"title"`
	Created       string            `json:"created_at"`
	Updated       string            `json:"updated_at"`
    Featured      string            `json:"featured_at"`
	color string            `json:"color"`
	SubscriberCount int               `json:"subscriber_count"`
	category_id int               `json:"category_id"`
	posts_count int `json:"posts_count"`
	User          User              `json:"user"`
	Posts []Post `json:"posts"`
}

func (c Collection) Summary() string {
	return fmt.Sprintf("collection[%s: %s]", c.Name, c.Title)
}

