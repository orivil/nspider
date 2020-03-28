// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Init(gormDB *gorm.DB) error {
	db = gormDB
	err := db.AutoMigrate(
		&Auth{},
		&Book{},
		&BookPattern{},
		&Chapter{},
		&ChapterPattern{},
		&Platform{},
		&PlatformProxy{},
		&Sock5Proxy{},
	).Error
	return err
}
