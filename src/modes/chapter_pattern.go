// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

type ChapterPattern struct {
	ID int
	PlatformID int `gorm:"index"`
	Url string
	Title string
	Content string
	NextID string `desc:"下一章ID"`
	TitleDelPattern string `desc:"删除标题正则规则，用于删除特殊字符，数组字符串类型"`
	ContentDel string `desc:"删除内容正则规则，用于删除特殊字符，数组字符串类型"`
}
