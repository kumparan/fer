package generator

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	//Migration define db migration file generator
	Migration interface {
		Generate(name string) error
	}

	migration struct{}
)

//NewMigrationGenerator :nodoc:
func NewMigrationGenerator() Migration {
	return &migration{}
}

//GenerateMigration to generate sql migration file
func (m migration) Generate(name string) error {
	err := m.checkMigrationFolderExists()
	if err != nil {
		return err
	}
	migrationFile := []byte(`-- +migrate Up notransaction` + "\n\n" + `-- +migrate Down`)
	migrationFileName := "db/migration/" + m.createUniqueTime() + "_" + strings.ToLower(name) + ".sql"
	err = ioutil.WriteFile(migrationFileName, migrationFile, 0666)
	if err != nil {
		return err
	}
	fmt.Println(migrationFileName + " created")
	return nil
}

func (m migration) checkMigrationFolderExists() error {
	_, err := os.Stat("db/migration/")
	if os.IsNotExist(err) { //check if db/migration folder is already exist
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("db/migration/ folder not found, want to  create (Y/N)? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return errors.New("fail when read user input")
		}
		ans := strings.Contains(strings.ToUpper(input), "Y")
		if !ans {
			return errors.New("cancel create migration")
		}
		err = m.createMigrationFolder()
		if err != nil {
			return err
		}
	}
	return nil

}

func (m migration) createMigrationFolder() error {
	_, err := os.Stat("db/")
	if os.IsNotExist(err) { //check if db folder is already exist
		err := os.Mkdir("db/", os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll("db/migration/", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (m migration) createUniqueTime() string {
	now := time.Now()
	splitDate := strings.Split(now.Format("01/02/2006"), "/") // mm/dd/yyyy
	newDate := splitDate[2] + splitDate[0] + splitDate[1]
	hr, min, sc := now.Clock()
	hour := strconv.Itoa(hr)
	minute := strconv.Itoa(min)
	sec := strconv.Itoa(sc)
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	if len(sec) == 1 {
		sec = "0" + sec
	}

	return newDate + hour + minute + sec
}
