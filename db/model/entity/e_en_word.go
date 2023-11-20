// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EEnWord is the golang structure for table e_en_word.
type EEnWord struct {
	Id           int64       `json:"id"           description:""`
	Pid          int64       `json:"pid"          description:""`
	HasKids      int         `json:"hasKids"      description:""`
	PartOfSpeech string      `json:"partOfSpeech" description:""`
	NounType     string      `json:"nounType"     description:""`
	Gender       string      `json:"gender"       description:""`
	Parents      string      `json:"parents"      description:""`
	Subject      string      `json:"subject"      description:""`
	Word         string      `json:"word"         description:""`
	Zh           string      `json:"zh"           description:""`
	Level        int         `json:"level"        description:""`
	Deep         int         `json:"deep"         description:""`
	Status       string      `json:"status"       description:""`
	En           string      `json:"en"           description:""`
	Root         string      `json:"root"         description:""`
	Prefix       string      `json:"prefix"       description:""`
	Suffix       string      `json:"suffix"       description:""`
	Desc         string      `json:"desc"         description:""`
	DemoSentence string      `json:"demoSentence" description:""`
	Img          string      `json:"img"          description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
}
