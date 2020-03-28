// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

import (
	"github.com/orivil/morgine/utils/sql"
	"time"
)

type Book struct {
	ID int
	PtID int `gorm:"index"`
	PtBookID int `gorm:"index"`
	Title string `gorm:"index"`
	Desc string
	Pic string
	Author string `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
	Finished sql.Boolean `gorm:"index"`
	Forbidden sql.Boolean `gorm:"index"`
}

func CreateBook(book *Book) error {
	return db.Create(book).Error
}

func UpdateBook(book *Book)  {
	
}
