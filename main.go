package main

import (
  "database/sql" 
  _ "github.com/lib/pq"
  "fmt"
  "log"
)

const (
  host     = "arjuna.db.elephantsql.com"
  port     = 5432
  user     = "lwzbzejw"
  password = "i0xAI776rL4ETUAljCkEhAxI3hH3Oh39"
  dbname   = "lwzbzejw"
)

func main() {
  // Acessing database
   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("database connected")
  }
  
  defer db.Close()

  // retrieving result
  // using.Query
  var (
    id int
    name string
    gender string
  )

  rows, err := db.Query("select * from students")
  
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  
  for rows.Next() {
    err := rows.Scan(&id, &name, &gender)
    if err != nil {
      log.Fatal(err)
    }
  
    log.Println(id, name)
  }
  
  err = rows.Err()
  
  if err != nil {
       log.Fatal(err)
  }

  // using QueryRow()
  err = db.QueryRow("select name from students where id = 1").Scan(&name)
  
  if err != nil {
    log.Fatal(err)
  }
  
  log.Println(name)

  // using db.Prepare()
  stmt, err := db.Prepare("select * from students where id = $1")

  if err != nil {
    fmt.Println("test")
    log.Fatal(err)
  }
  
  defer stmt.Close()
  
  rows, err = stmt.Query(2)
  
  if err != nil {
    log.Fatal(err)
  }
  
  defer rows.Close() 
  
  for rows.Next() {
    err := rows.Scan(&id, &name, &gender)
    if err != nil {
      log.Fatal(err)
    }
  
    log.Println(id, name)
  }
  
  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  stmt, err = db.Prepare("insert into students(id, name, gender) values($1, $2, $3)")

  if err != nil {
    log.Fatal(err)
  }
  
  res, err := stmt.Exec(3,"test", string('F'))

  fmt.Println(res)
  
  if err != nil {
    log.Fatal(err)
  }
}