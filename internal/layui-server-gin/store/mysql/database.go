package mysql

import (
	"fmt"
	"github.com/lihaicheng/layui-server-go/internal/layui-server-gin/pkg/config"
	"github.com/lihaicheng/layui-server-go/internal/layui-server-gin/store/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

// InitDB 初始化mysql数据库
func InitDB(cfg *config.Settings) error {
	user := cfg.MysqlSettings.User
	password := cfg.MysqlSettings.Password
	host := cfg.MysqlSettings.Host
	port := cfg.MysqlSettings.Port
	database := cfg.MysqlSettings.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user, password, host, port, database)

	dialector := mysql.Open(dsn)
	dbConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	var err error
	DB, err = gorm.Open(dialector, dbConfig)
	if err != nil {
		return err
	}
	db, _ := DB.DB()
	err = db.Ping()
	if err != nil {
		return err
	}
	maxIdleConns := cfg.MysqlSettings.MaxIdleConns
	if maxIdleConns > 0 {
		db.SetMaxIdleConns(maxIdleConns)
	} else {
		db.SetMaxIdleConns(10)
	}

	maxOpenConns := cfg.MysqlSettings.MaxOpenConns
	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	} else {
		db.SetMaxOpenConns(100)
	}
	db.SetConnMaxLifetime(time.Hour)
	err = autoMigrate()
	if err != nil {
		zap.L().Error("database: auto migrate database schema failed")
		return err
	}

	err = insertDefaultRecord()
	if err != nil {
		zap.L().Error("database: insert default record failed")
		return err
	}

	return nil
}

func autoMigrate() error {
	tables := []model.Table{
		model.Menu{},
	}
	for _, table := range tables {
		err := DB.AutoMigrate(table)
		if err != nil {
			zap.L().Error(fmt.Sprintf("database: table %s auto migrate failed", table.TableName()))
			return err
		}
	}
	return nil
}

func insertDefaultRecord() error {
	var err error
	err = insertMenuRecord()
	if err != nil {
		return err
	}
	return nil
}

func insertMenuRecord() error {
	jsonText := `{
		  "homeInfo": {
			"title": "首页",
			"href": "layuimini/page/welcome-1.html?t=1"
		  },
		  "logoInfo": {
			"title": "LAYUI MINI",
			"image": "layuimini/images/logo.png",
			"href": ""
		  },
		  "menuInfo": [
			{
			  "title": "常规管理",
			  "icon": "fa fa-address-book",
			  "href": "",
			  "target": "_self",
			  "child": [
				{
				  "title": "主页模板",
				  "href": "",
				  "icon": "fa fa-home",
				  "target": "_self",
				  "child": [
					{
					  "title": "主页一",
					  "href": "layuimini/page/welcome-1.html",
					  "icon": "fa fa-tachometer",
					  "target": "_self"
					},
					{
					  "title": "主页二",
					  "href": "layuimini/page/welcome-2.html",
					  "icon": "fa fa-tachometer",
					  "target": "_self"
					},
					{
					  "title": "主页三",
					  "href": "layuimini/page/welcome-3.html",
					  "icon": "fa fa-tachometer",
					  "target": "_self"
					}
				  ]
				},
				{
				  "title": "菜单管理",
				  "href": "layuimini/page/menu.html",
				  "icon": "fa fa-window-maximize",
				  "target": "_self"
				},
				{
				  "title": "系统设置",
				  "href": "layuimini/page/setting.html",
				  "icon": "fa fa-gears",
				  "target": "_self"
				},
				{
				  "title": "表格示例",
				  "href": "layuimini/page/table.html",
				  "icon": "fa fa-file-text",
				  "target": "_self"
				},
				{
				  "title": "表单示例",
				  "href": "",
				  "icon": "fa fa-calendar",
				  "target": "_self",
				  "child": [
					{
					  "title": "普通表单",
					  "href": "layuimini/page/form.html",
					  "icon": "fa fa-list-alt",
					  "target": "_self"
					},
					{
					  "title": "分步表单",
					  "href": "layuimini/page/form-step.html",
					  "icon": "fa fa-navicon",
					  "target": "_self"
					}
				  ]
				},
				{
				  "title": "登录模板",
				  "href": "",
				  "icon": "fa fa-flag-o",
				  "target": "_self",
				  "child": [
					{
					  "title": "登录-1",
					  "href": "layuimini/page/login-1.html",
					  "icon": "fa fa-stumbleupon-circle",
					  "target": "_blank"
					},
					{
					  "title": "登录-2",
					  "href": "layuimini/page/login-2.html",
					  "icon": "fa fa-viacoin",
					  "target": "_blank"
					},
					{
					  "title": "登录-3",
					  "href": "layuimini/page/login-3.html",
					  "icon": "fa fa-tags",
					  "target": "_blank"
					}
				  ]
				},
				{
				  "title": "异常页面",
				  "href": "",
				  "icon": "fa fa-home",
				  "target": "_self",
				  "child": [
					{
					  "title": "404页面",
					  "href": "layuimini/page/404.html",
					  "icon": "fa fa-hourglass-end",
					  "target": "_self"
					}
				  ]
				},
				{
				  "title": "其它界面",
				  "href": "",
				  "icon": "fa fa-snowflake-o",
				  "target": "",
				  "child": [
					{
					  "title": "按钮示例",
					  "href": "layuimini/page/button.html",
					  "icon": "fa fa-snowflake-o",
					  "target": "_self"
					},
					{
					  "title": "弹出层",
					  "href": "layuimini/page/layer.html",
					  "icon": "fa fa-shield",
					  "target": "_self"
					}
				  ]
				}
			  ]
			},
			{
			  "title": "组件管理",
			  "icon": "fa fa-lemon-o",
			  "href": "",
			  "target": "_self",
			  "child": [
				{
				  "title": "图标列表",
				  "href": "layuimini/page/icon.html",
				  "icon": "fa fa-dot-circle-o",
				  "target": "_self"
				},
				{
				  "title": "图标选择",
				  "href": "layuimini/page/icon-picker.html",
				  "icon": "fa fa-adn",
				  "target": "_self"
				},
				{
				  "title": "颜色选择",
				  "href": "layuimini/page/color-select.html",
				  "icon": "fa fa-dashboard",
				  "target": "_self"
				},
				{
				  "title": "下拉选择",
				  "href": "layuimini/page/table-select.html",
				  "icon": "fa fa-angle-double-down",
				  "target": "_self"
				},
				{
				  "title": "文件上传",
				  "href": "layuimini/page/upload.html",
				  "icon": "fa fa-arrow-up",
				  "target": "_self"
				},
				{
				  "title": "富文本编辑器",
				  "href": "layuimini/page/editor.html",
				  "icon": "fa fa-edit",
				  "target": "_self"
				},
				{
				  "title": "省市县区选择器",
				  "href": "layuimini/page/area.html",
				  "icon": "fa fa-rocket",
				  "target": "_self"
				}
			  ]
			},
			{
			  "title": "其它管理",
			  "icon": "fa fa-slideshare",
			  "href": "",
			  "target": "_self",
			  "child": [
				{
				  "title": "多级菜单",
				  "href": "",
				  "icon": "fa fa-meetup",
				  "target": "",
				  "child": [
					{
					  "title": "按钮1",
					  "href": "layuimini/page/button.html?v=1",
					  "icon": "fa fa-calendar",
					  "target": "_self",
					  "child": [
						{
						  "title": "按钮2",
						  "href": "layuimini/page/button.html?v=2",
						  "icon": "fa fa-snowflake-o",
						  "target": "_self",
						  "child": [
							{
							  "title": "按钮3",
							  "href": "layuimini/page/button.html?v=3",
							  "icon": "fa fa-snowflake-o",
							  "target": "_self"
							},
							{
							  "title": "表单4",
							  "href": "layuimini/page/form.html?v=1",
							  "icon": "fa fa-calendar",
							  "target": "_self"
							}
						  ]
						}
					  ]
					}
				  ]
				},
				{
				  "title": "失效菜单",
				  "href": "layuimini/page/error.html",
				  "icon": "fa fa-superpowers",
				  "target": "_self"
				}
			  ]
			}
		  ]
		}`
	menu := model.Menu{Flag: "test", Value: jsonText}
	var count int64
	err := DB.Model(&menu).Where(&model.Menu{Flag: "test"}).Count(&count).Error
	if err != nil {
		zap.L().Error("database: check test record failed")
		return err
	}
	if count == 0 {
		DB.Create(&menu)
		if err != nil {
			zap.L().Error("database: insert test record failed")
			return err
		}
	}
	return nil
}
