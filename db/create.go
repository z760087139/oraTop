package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	// _ "github.com/ibmdb/go_ibm_db"
)

var (
	OraDB OracleDB
)

func init() {
	OraDB = CreateOraLocal()
}

// 数据库连接信息
type ConnInfo struct {
	// 数据库IP地址
	Host string
	// 数据库端口
	Port int
	// dbname(db2) sid(oracle) db(mysql)
	Sid string
	// 连接用户名
	Username string
	// 密码
	Password string
}

// // 根据信息创建DB2连接
// func (self ConnInfo) CreateDB2() DB2DB {
// 	fmtString := "HOSTNAME=%s;DATABASE=%s;PORT=%d;UID=%s;PWD=%s"
// 	connStr := fmt.Sprintf(
// 		fmtString,
// 		self.Host, self.Sid, self.Port,
// 		self.Username, self, self.Password)
// 	driverName := "go_ibm_db"
// 	db := CreateDB(connStr, driverName)
// 	return DB2DB{db}
// }

// 根据信息创建ORACLE连接
func (self ConnInfo) CreateOracle() OracleDB {
	fmtString := "%s/%s@%s:%d/%s"
	connStr := fmt.Sprintf(
		fmtString,
		self.Username, self.Password,
		self.Host, self.Port, self.Sid,
	)
	driverName := "godror"
	db := CreateDB(connStr, driverName)
	return OracleDB{db}
}

// // 根据信息创建MYSQL连接
// func (self ConnInfo) CreateMysql() MysqlDB {
// 	cfg := mysql.NewConfig()
// 	addr := fmt.Sprintf("%s:%d", self.Host, self.Port)
// 	cfg.Net = "tcp"
// 	cfg.User = self.Username
// 	cfg.Passwd = self.Password
// 	cfg.DBName = self.Sid
// 	cfg.Addr = addr
// 	connStr := cfg.FormatDSN()
// 	driverName := "mysql"
// 	db := CreateDB(connStr, driverName)
// 	return MysqlDB{db}
// }

func CreateDB(con, driverName string) *sql.DB {
	db, err := sql.Open(driverName, con)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db
}

func CreateOraLocal() OracleDB {
	dsn := "oracle://?sysdba=1"
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return OracleDB{db}
}
