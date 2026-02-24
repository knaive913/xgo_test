//line C:\Users\wfg20\go\pkg\mod\github.com\xhd2015\xgo\runtime@v1.1.10\trace\trace.go:1
package trace

import (
	_ "github.com/xhd2015/xgo/runtime/internal/trap"
	"github.com/xhd2015/xgo/runtime/trace/stack_model"
)

// NOTE: don't add more functions to this file,
// it is specially instrumented by xgo compiler
// to call trap automatically

type Config struct {
	// OnFinish is called when the trace is finished
	OnFinish func(stack stack_model.IStack) `json:"-"`
	// OutputFile specifies the file to save the trace
	// in json format, which can be open by:
	//      xgo tool trace <OutputFile>
	OutputFile string `json:"OutputFile,omitempty"`
};type __xgo_func_info_1 struct{Kind int;FullName string;Pkg string;IdentityName string;Name string;RecvType string;RecvPtr bool;Interface bool;Generic bool;Closure bool;Stdlib bool;File string;Line int;PC uintptr;Func interface{};Var interface{};RecvName string;ArgNames []string;ResNames []string;FirstArgCtx bool;LastResultErr bool};var __xgo_register_1 = func(v interface{}){};var __xgo_trap_1 = func(info interface{}, recvPtr interface{}, args []interface{}, results []interface{}) (func(), bool){return nil, false;};var __xgo_trap_var_1 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_trap_varptr_1 = func(info interface{}, varAddr interface{}, res interface{}){};var __xgo_pkg_1="github.com/xhd2015/xgo/runtime/trace";var __xgo_file_1="C:\\Users\\wfg20\\go\\pkg\\mod\\github.com\\xhd2015\\xgo\\runtime@v1.1.10\\trace\\trace.go";var __xgo_file_var_1="C:\\Users\\wfg20\\go\\pkg\\mod\\github.com\\xhd2015\\xgo\\runtime@v1.1.10\\trace\\trace.go";var __xgo_func_info_1_0=&__xgo_func_info_1{Kind:0,Pkg:__xgo_pkg_1,Name:"Trace",IdentityName:"Trace",File:__xgo_file_1,Line:22,ArgNames:[]string{"config","request","fn"},ResNames:[]string{"response","err"}};func init(){__xgo_func_info_1_0.Func=Trace;};func init(){__xgo_init_1();__xgo_register_1(__xgo_func_info_1_0);}

// the `request` and `response` are only for recording purpose
func Trace(config Config, request interface{}, fn func() (interface{}, error)) (response interface{}, err error) {__xgo_post_22, __xgo_stop_22 := __xgo_trap_1(__xgo_func_info_1_0,nil,[]interface{}{&config,&request,&fn},[]interface{}{&response,&err});if __xgo_post_22!=nil { defer __xgo_post_22(); };if __xgo_stop_22 { return; };
	return fn()
}
//go:noinline
func __xgo_init_1(){__xgo_register_1=__xgo_register_1;__xgo_trap_1=__xgo_trap_1;__xgo_trap_var_1=__xgo_trap_var_1;__xgo_trap_varptr_1=__xgo_trap_varptr_1;}
