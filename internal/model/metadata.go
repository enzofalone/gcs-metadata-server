package model

import "time"

type Metadata struct {
	Bucket       string    `json:"bucket" db:"bucket"`
	Name         string    `json:"name" db:"name"`
	Parent       string    `json:"parent" db:"parent"`
	Size         int64     `json:"size" db:"size"`
	Count        int64     `json:"count" db:"count"`
	StorageClass string    `json:"storage_class" db:"storage_class"`
	Created      time.Time `json:"created" db:"created"`
	Updated      time.Time `json:"updated" db:"updated"`
}
