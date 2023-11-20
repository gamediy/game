// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EEnWord is the golang structure of table e_en_word for DAO operations like Where/Data.
type EEnWord struct {
	g.Meta       `orm:"table:e_en_word, do:true"`
	Id           interface{} //
	Pid          interface{} //
	HasKids      interface{} //
	PartOfSpeech interface{} //
	NounType     interface{} //
	Gender       interface{} //
	Parents      interface{} //
	Subject      interface{} //
	Word         interface{} //
	Zh           interface{} //
	Level        interface{} //
	Deep         interface{} //
	Status       interface{} //
	En           interface{} //
	Root         interface{} //
	Prefix       interface{} //
	Suffix       interface{} //
	Desc         interface{} //
	DemoSentence interface{} //
	Img          interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
