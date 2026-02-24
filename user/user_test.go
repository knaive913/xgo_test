package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/xhd2015/xgo/runtime/mock"
)

func TestCreateRecordWithoutTable(t *testing.T) {
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
}
