package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Seat struct {
	ID      int
	Student string
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

	// Retrieve env values
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetch the seat data
	seats, err := getSeatData(db)
	if err != nil {
		log.Fatal(err)
	}

	// Swap seat ids of consecutive students
	swapSeatIDs(seats)

	// Update the seat data
	err = updateSeatData(db, seats)
	if err != nil {
		log.Fatal(err)
	}

	// fetch updated seat data
	updatedSeats, err := getSeatData(db)
	if err != nil {
		log.Fatal(err)
	}

	//print upeat data
	printSeatData(updatedSeats)
}


//get seat data
func getSeatData(db *sql.DB) ([]Seat, error) {
	rows, err := db.Query("SELECT id, student FROM seat ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []Seat
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.Student)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	return seats, nil
}


// update seat
func updateSeatData(db *sql.DB, seats []Seat) error {
	// Sort seats by id
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID < seats[j].ID
	})

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			tx.Commit()
		}
	}()

	// Update seat data
	for _, seat := range seats {
		_, err := tx.Exec("UPDATE seat SET student = $1 WHERE id = $2", seat.Student, seat.ID)
		if err != nil {
			return err
		}
	}

	return nil
}


//swap seat id function
func swapSeatIDs(seats []Seat) {
	for i := 0; i < len(seats)-1; i += 2 {
		seats[i].ID, seats[i+1].ID = seats[i+1].ID, seats[i].ID
	}
}


//print updated seat data
func printSeatData(seats []Seat) {
	fmt.Println("+--------+------------------+")
	fmt.Println("| id | student |")
	fmt.Println("+--------+------------------+")
	for _, seat := range seats {
		fmt.Printf("| %d | %s |\n", seat.ID, seat.Student)
	}
	fmt.Println("+--------+------------------+")
}
