package biz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswork(t *testing.T) {
	s := hashPassword("123456")
	t.Log(s)                        // 输出加密后的密码
	assert.NotEqual(t, s, "123456") // 加密后的密码不等于原始密码
}

func TestVerifyPassword(t *testing.T) {
	hashPwd := "$2a$10$C.EFL6UY9NewoPSUd6bCeuJs0/ihHdGfgIb0q5hceB35CUm68Iu3C"
	assert.True(t, verifyPassword(hashPwd, "123456"))
	assert.False(t, verifyPassword(hashPwd, "1234567"))
}
