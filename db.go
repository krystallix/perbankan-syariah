package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB

// ConnectDB opens a connection to the database
func DbConnect(QueryOriginal string) {
	failConnect := 1
	for {
		// db, err := sql.Open("mysql", "aang:aang@tcp(192.168.1.254:3306)/arbiterosv2")
		db, err := sql.Open("mysql", ConfDbUser+":"+ConfDbPass+"@tcp("+ConfDbServer+")/"+ConfDbName)

		// db.SetMaxIdleConns(1000)
		// db.SetMaxOpenConns(1000)
		// db.SetConnMaxLifetime(100 * time.Second)
		db.SetMaxIdleConns(200)
		db.SetMaxOpenConns(200)
		db.SetConnMaxLifetime(10 * time.Second)
		// db.SetConnMaxLifetime(800 * time.Millisecond)

		if err != nil {
			// panic(err.Error())
			fmt.Println("fail connect db")
		} else {
			fmt.Println("\n\n\n")
			fmt.Println("Success Reconnect db, trying to ping...")
			err := db.Ping()
			if err != nil {
				fmt.Println("DB ping Error : ", err)
				fmt.Println("Query : ", QueryOriginal)
			} else {
				failConnect = 0
				DbCon = db
				fmt.Println("Db Ping Success, connection stabilize ")
			}
		}

		if failConnect == 0 {
			break
		}
		fmt.Println("FOR DI DB CONNECT")
	}
	// defer
}

func DbClose() {
	DbCon.Close()
}

func DbQuery(query string) *sql.Rows {
	QuerySuccess := 0
	var results *sql.Rows
	for {
		results, err := DbCon.Query(query)
		if err != nil {
			fmt.Println("query fail : ")
			fmt.Println(query)
			fmt.Println("DbQuery db.go (0) : ", err.Error())
			fmt.Print("reconnecting : ")
			DbPing(query)
		} else {
			QuerySuccess = 1
		}

		if QuerySuccess == 1 {
			return results
		}
		fmt.Println("FOR DI DB QUERY")

	}
	return results
}

func DbUpdate(query string) sql.Result {
	QuerySuccess := 0
	var results sql.Result
	for {
		results, err := DbCon.Exec(query)
		// fmt.Println(reflect.TypeOf(results))

		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("\nDbUpdate fail, \n origin query : ", query, " \n with error : ", err.Error(), "reconnecting : ")
			// panic(err.Error()) // proper error handling instead of panic in your app
			DbPing(query)
		} else {
			QuerySuccess = 1
		}

		if QuerySuccess == 1 {
			return results
		}
		fmt.Println("FOR DI DB UPDATE")

	}
	return results
}

func DbFunc(query string) string {
	QuerySuccess := 0
	// var results sql.Result
	for {
		_, err := DbCon.Exec(query)
		// fmt.Println(reflect.TypeOf(results))

		if err != nil {
			fmt.Println()
			fmt.Print("\nDbFunc fail, \n origin query : ", query, " \n with error : ", err.Error(), "\n reconnecting : ")
			DbPing(query)
			time.Sleep(100 * time.Millisecond)
		} else {
			QuerySuccess = 1
		}

		if QuerySuccess == 1 {
			return "ok"
		}
		fmt.Println("FOR DI DB FUNC")

	}
	return "fail"
}

func DbExec(query string) string {
	QuerySuccess := 0
	// var results sql.Result
	for {
		_, err := DbCon.Exec(query)
		// fmt.Println(reflect.TypeOf(results))

		if err != nil {
			fmt.Println()
			fmt.Print("\nDbExec fail, \n origin query : ", query, " \n with error : ", err.Error(), "\n reconnecting : ")
			DbPing(query)
			time.Sleep(10 * time.Millisecond)
		} else {
			QuerySuccess = 1
		}

		if QuerySuccess == 1 {
			return "ok"
		}
		fmt.Println("FOR DI DB FUNC")

	}
	return "fail"
}

func DbPing(QueryOriginal string) {
	var ConEstablished bool
	for {
		if err := DbCon.Ping(); err != nil {
			fmt.Println("ErrorPing : ", err)
			time.Sleep(10 * time.Millisecond)
			// fmt.Println("Reconnecting to db... ")
			// DbConnect(QueryOriginal)
		} else {
			ConEstablished = true
			fmt.Println("success")
		}
		if ConEstablished == true {
			break
		}
	}
}
