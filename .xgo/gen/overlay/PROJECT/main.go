//line C:\Users\wfg20\OneDrive\Desktop\demo\main.go:1
package main

import (
	"fmt"
)

func main() {__xgo_post_7, __xgo_stop_7 := __xgo_trap_0(__xgo_func_info_0_0,nil,[]interface{}{},[]interface{}{});if __xgo_post_7!=nil { defer __xgo_post_7(); };if __xgo_stop_7 { return; };
	fmt.Println("hello world")
};type __xgo_func_info_0 struct{Kind int;FullName string;Pkg string;IdentityName string;Name string;RecvType string;RecvPtr bool;Interface bool;Generic bool;Closure bool;Stdlib bool;File string;Line int;PC uintptr;Func interface{};Var interface{};RecvName string;ArgNames []string;ResNames []string;FirstArgCtx bool;LastResultErr bool};var __xgo_register_0 = func(v interface{}){};var __xgo_trap_0 = func(info interface{}, recvPtr interface{}, args []interface{}, results []interface{}) (func(), bool){return nil, false;};var __xgo_trap_var_0 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_trap_varptr_0 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_pkg_0="main";var __xgo_file_0="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\main.go";var __xgo_file_var_0="C:\\Users\\wfg20\\OneDrive\\Desktop\\demo\\main.go";var __xgo_func_info_0_0=&__xgo_func_info_0{Kind:0,Pkg:__xgo_pkg_0,Name:"main",IdentityName:"main",File:__xgo_file_0,Line:7};func init(){__xgo_func_info_0_0.Func=main;};func init(){__xgo_init_0();__xgo_register_0(__xgo_func_info_0_0);}
//go:noinline
func __xgo_init_0(){__xgo_register_0=__xgo_register_0;__xgo_trap_0=__xgo_trap_0;__xgo_trap_var_0=__xgo_trap_var_0;__xgo_trap_varptr_0=__xgo_trap_varptr_0;}
