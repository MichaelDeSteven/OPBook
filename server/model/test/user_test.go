package test

import (
	"os"
	"strings"
	"testing"

	"github.com/MichaelDeSteven/OPBook/server/core"
	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/initialize"
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	name := "server"
	abs, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(abs, "\\")
	var path strings.Builder
	for _, t := range dirs {
		path.WriteString(t)
		path.WriteByte('\\')
		if t == name {
			break
		}
	}
	path.WriteString(utils.ConfigFile)
	abs = path.String()
	global.VP = core.Viper(abs)   // 初始化Viper
	global.DB = initialize.Gorm() // gorm连接数据库
	return global.DB
}

func TestAddUser(t *testing.T) {
	db := initDB()
	user := model.User{
		Nickname: "Steven",
		Mobile:   "13006609678",
		Role:     1,
		Email:    "127072@qq.com",
	}
	// tx := db.Select("Nickname", "Mobile", "Role", "Description").Create(&user)
	tx := db.Create(&user)

	assert.NoError(t, tx.Error)
	assert.Equal(t, int64(1), tx.RowsAffected)
}

func TestDelete(t *testing.T) {
	db := initDB()
	user := model.User{}
	tx := db.Where("nickname = ?", "Steven").Delete(&user)
	assert.NoError(t, tx.Error)
	assert.Equal(t, int64(1), tx.RowsAffected)
}

func TestFindUserByEmail(t *testing.T) {
	db := initDB()
	var user model.User
	db.Where("email = ?", "127072@qq.com").First(&user)
	assert.NotNil(t, user, "user is nil")
	assert.Equal(t, "13006609678", user.Mobile)
}
