//line C:\Users\wfg20\OneDrive\Desktop\demo\user\user.go:1
package user

import "github.com/jinzhu/gorm"

// Player 定义数据库模型
type Player struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Name       string `gorm:"size:255;not null"`
	Hero       string `gorm:"size:255"`
	MMR        int
};type __xgo_func_info_0 struct{Kind int;FullName string;Pkg string;IdentityName string;Name string;RecvType string;RecvPtr bool;Interface bool;Generic bool;Closure bool;Stdlib bool;File string;Line int;PC uintptr;Func interface{};Var interface{};RecvName string;ArgNames []string;ResNames []string;FirstArgCtx bool;LastResultErr bool};var __xgo_register_0 = func(v interface{}){};var __xgo_trap_0 = func(info interface{}, recvPtr interface{}, args []interface{}, results []interface{}) (func(), bool){return nil, false;};var __xgo_trap_var_0 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_trap_varptr_0 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_pkg_0="main/user";var __xgo_file_0="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\user\\user.go";var __xgo_file_var_0="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\user\\user.go";var __xgo_func_info_0_0=&__xgo_func_info_0{Kind:0,Pkg:__xgo_pkg_0,Name:"CreateRecord",IdentityName:"CreateRecord",File:__xgo_file_0,Line:14,ArgNames:[]string{"db"},ResNames:[]string{"__xgo_auto_res_0"}};var __xgo_func_info_0_1=&__xgo_func_info_0{Kind:0,Pkg:__xgo_pkg_0,Name:"GetUserID",IdentityName:"GetUserID",File:__xgo_file_0,Line:23,ResNames:[]string{"__xgo_auto_res_0"}};func init(){__xgo_func_info_0_0.Func=CreateRecord;__xgo_func_info_0_1.Func=GetUserID;};func init(){__xgo_init_0();__xgo_register_0(__xgo_func_info_0_0);__xgo_register_0(__xgo_func_info_0_1);}

// 简单的数据库操作函数
func CreateRecord(db *gorm.DB) ( __xgo_auto_res_0 error ){__xgo_post_14, __xgo_stop_14 := __xgo_trap_0(__xgo_func_info_0_0,nil,[]interface{}{&db},[]interface{}{&__xgo_auto_res_0});if __xgo_post_14!=nil { defer __xgo_post_14(); };if __xgo_stop_14 { return; };
	record := map[string]interface{}{
		"name": "test",
		"age":  25,
	}
	return db.Table("users").Create(record).Error
}

// 被mock的函数
func GetUserID() ( __xgo_auto_res_0 uint64 ){__xgo_post_23, __xgo_stop_23 := __xgo_trap_0(__xgo_func_info_0_1,nil,[]interface{}{},[]interface{}{&__xgo_auto_res_0});if __xgo_post_23!=nil { defer __xgo_post_23(); };if __xgo_stop_23 { return; };
	return 123
}
//go:noinline
func __xgo_init_0(){__xgo_register_0=__xgo_register_0;__xgo_trap_0=__xgo_trap_0;__xgo_trap_var_0=__xgo_trap_var_0;__xgo_trap_varptr_0=__xgo_trap_varptr_0;}
