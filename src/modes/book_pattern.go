// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

type BookPattern struct {
	ID              int
	PlatformID      int `gorm:"index"`
	Url             string
	Title           string
	Author          string
	DAuthor         string `desc:"默认作者，当没有作者信息时使用"`
	Category string `desc:"分类"`
	DCategory string `desc:"默认分类，当没有分类信息时使用"`
	Description     string
	Finished        string
	Pic             string
	FirstChapterID  string
	DescDelPattern string `desc:"删除简介正则规则，数组字符串类型"`
}
