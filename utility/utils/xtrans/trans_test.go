package xtrans

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestT(t *testing.T) {
	Init("../../../app/api-user_svc/resource/i18n")
	g.Dump(T("en", "-200"))
}
