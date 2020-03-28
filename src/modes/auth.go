// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package modes

type Auth struct {
	ID int
	Charset string `desc:"编码, 'GBK' 'UTF-8'等"`
	PlatformID int `gorm:"unique_index"`
	UserAgent string
	Cookie string
	WaitTime int `desc:"等待时间, 单位: 毫秒"`
}

func CreateAuth(auth *Auth) error {
	return db.Create(auth).Error
}

func UpdateAuth(auth *Auth) error {
	return db.Model(auth).Updates(auth).Error
}

func DelAuth(id int) error {
	return db.Where("id=?", id).Delete(&Auth{}).Error
}
