package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// Global Variables
var userid int

// Main
func main() {

  http.HandleFunc("/", index)
  http.HandleFunc("/users",  usersHandler)
  http.HandleFunc("/users/", usersHandler)

  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    fmt.Println("Serve Http: ", err)
  }

}

func usersHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"User Input\n")
  fmt.Fprint(w,"\n")

  fmt.Println("usersHandler: connecting to the database..")

  // Open database connection
  var db_host string
  var db_database string
  var db_user string
  var db_password string
  db_host     = os.Getenv("MYSQL_HOST")
  db_database = os.Getenv("MYSQL_DATABASE")
  db_user     = os.Getenv("MYSQL_USER")
  db_password = os.Getenv("MYSQL_PASSWORD")

  //connectionString := "mysqluser:mysqlpw@tcp(172.17.0.2:3306)/openmrs"
  connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_host, db_database)

  dbc, err := sql.Open("mysql", connectionString)
  if err != nil {
    log.Fatal(err)
  }
  //defer dbc.Close()

  // initialise some db variables
  user_id := 1
  var col string
  var system_id string
  last_id := 0
  new_id := 0


  // single row - read admin user

  fmt.Println("usersHandler: reading admin user..")

  sqlStatement := `SELECT system_id FROM users WHERE user_id=? LIMIT 1`
  row := dbc.QueryRow(sqlStatement, user_id)
  dberr := row.Scan(&col)
  if dberr != nil {
    if dberr == sql.ErrNoRows {
      fmt.Println("usersHandler: no user entry found for user_id")
    } else {
    }
      log.Fatal(dberr)
  }
  fmt.Println("user_id:", user_id, "has system_id:", col)


  // read last user_id value

  fmt.Println("usersHandler: reading last user_id value..")

  sqlStatement2 := `select max(user_id) from users`
  row2 := dbc.QueryRow(sqlStatement2)
  dberr2 := row2.Scan(&last_id)
  if dberr2 != nil {
    if dberr2 == sql.ErrNoRows {
      fmt.Println("usersHandler: no user entres found")
    } else {
    }
      log.Fatal(dberr2)
  }
  new_id = last_id+1
  fmt.Println("next user_id is:", new_id)



  // insert a new dummy user

  new_name := fmt.Sprintf("user%d", new_id)
  fmt.Println("usersHandler: inserting new user entry, ", new_name)

  sqlStatement3 := "insert into users values (?,?,?,'6f0be51d599f59dd1269e12e17949f8ecb9ac963e467ac1400cf0a02eb9f8861ce3cca8f6d34d93c0ca34029497542cbadda20c949affb4cb59269ef4912087b','c788c6ad82a157b712392ca695dfcf2eed193d7f',NULL,NULL,1,'2019-06-01 00:00:00',1,'2019-06-01 00:00:00',1,0,NULL,NULL,NULL,'45ce6c2e-dd5a-11e6-9d9c-0242ac150002')"
  stmtIns, dberr3 := dbc.Prepare(sqlStatement3)
  if dberr3 != nil {
      log.Fatal("usersHandler: new user insert Prepare failed! ", dberr3)
  }
  _, dberr3 = stmtIns.Exec(new_id, new_name, new_name)
  if dberr3 != nil {
      log.Fatal("usersHandler: new user insert Exec failed! ", dberr3)
  }


  // read and display user entries

  fmt.Println("usersHandler: Current users are:")

  sqlStatement4 := `SELECT user_id, system_id FROM users`
  rows, dberr4 := dbc.Query(sqlStatement4)
  if dberr4 != nil {
    if dberr4 == sql.ErrNoRows {
      fmt.Println("usersHandler: no user entres found")
    } else {
      log.Fatal(dberr4)
    }
  }
  //defer rows.Close()
  for rows.Next() {
    rows.Scan(&user_id, &system_id)
    fmt.Println("user row: ", user_id, "\t", system_id)
    fmt.Fprint(w,"user row:\t", user_id, "\t", system_id, "\n")
  }
  fmt.Fprint(w,"New user added.\n")

}

func index(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w,"$$$$$$$\\  $$$$$$\\       $$\\ $$$$$$\\  $$$$$$\\ $$\\ \n")
  fmt.Fprint(w,"$$  __$$\\$$  __$$\\     $$  $$  __$$\\$$  __$$\\$$ |\n")
  fmt.Fprint(w,"$$ |  $$ $$ /  \\__|   $$  /$$ /  $$ $$ /  \\__$$ |\n")
  fmt.Fprint(w,"$$ |  $$ $$ |        $$  / $$ |  $$ \\$$$$$$\\ $$ |\n")
  fmt.Fprint(w,"$$ |  $$ $$ |       $$  /  $$ |  $$ |\\____$$\\__|\n")
  fmt.Fprint(w,"$$ |  $$ $$ |  $$\\ $$  /   $$ |  $$ $$\\   $$ |   \n")
  fmt.Fprint(w,"$$$$$$$  \\$$$$$$  $$  /     $$$$$$  \\$$$$$$  $$\\ \n")
  fmt.Fprint(w,"\\_______/ \\______/\\__/      \\______/ \\______/\\__|\n")
  fmt.Fprint(w,"\n")
}
