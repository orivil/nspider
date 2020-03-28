// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

type Chapter struct {
	ID int
	PtBookID int `gorm:"index"`
	PtChapterID int `gorm:"index"`
	CIndex int `gorm:"index"`
	Title string
	Content string
}