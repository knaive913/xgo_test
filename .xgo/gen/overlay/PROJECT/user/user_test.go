//line C:\Users\wfg20\OneDrive\Desktop\demo\user\user_test.go:1
package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/xhd2015/xgo/runtime/mock"
)

func TestCreateRecordWithoutTable(t *testing.T) {__xgo_post_11, __xgo_stop_11 := __xgo_trap_1(__xgo_func_info_1_0,nil,[]interface{}{&t},[]interface{}{});if __xgo_post_11!=nil { defer __xgo_post_11(); };if __xgo_stop_11 { return; };
	// 使用xgo mock一个函数
	mock.Patch(GetUserID, func() uint64 {
		return 456
	})

	// 创建sqlmock
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 创建gorm实例
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal(err)
	}

	// 设置sqlmock期望
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO \"users\"").
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	// 这里会触发 "fatal error: defer on system stack"
	err = CreateRecord(gormDB)
	if err != nil {
		t.Fatal(err)
	}

	// 验证mock期望
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
};type __xgo_func_info_1 struct{Kind int;FullName string;Pkg string;IdentityName string;Name string;RecvType string;RecvPtr bool;Interface bool;Generic bool;Closure bool;Stdlib bool;File string;Line int;PC uintptr;Func interface{};Var interface{};RecvName string;ArgNames []string;ResNames []string;FirstArgCtx bool;LastResultErr bool};var __xgo_register_1 = func(v interface{}){};var __xgo_trap_1 = func(info interface{}, recvPtr interface{}, args []interface{}, results []interface{}) (func(), bool){return nil, false;};var __xgo_trap_var_1 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_trap_varptr_1 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_pkg_1="main/user";var __xgo_file_1="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\user\\user_test.go";var __xgo_file_var_1="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\user\\user_test.go";var __xgo_func_info_1_0=&__xgo_func_info_1{Kind:0,Pkg:__xgo_pkg_1,Name:"TestCreateRecordWithoutTable",IdentityName:"TestCreateRecordWithoutTable",File:__xgo_file_1,Line:11,ArgNames:[]string{"t"}};func init(){__xgo_func_info_1_0.Func=TestCreateRecordWithoutTable;};func init(){__xgo_init_1();__xgo_register_1(__xgo_func_info_1_0);}
//go:noinline
func __xgo_init_1(){__xgo_register_1=__xgo_register_1;__xgo_trap_1=__xgo_trap_1;__xgo_trap_var_1=__xgo_trap_var_1;__xgo_trap_varptr_1=__xgo_trap_varptr_1;}
