package connection

import (
	"bufio"
	"fmt"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"os"
)

type Word struct {
	ID   int    `reindex:"id,,pk"` // 'id' is primary key
	Word string `reindex:"name"`   // add index by 'name' field
}

func InitConnection() *reindexer.Reindexer {
	database := reindexer.NewReindex("cproto://user:pass@reindexer_db:6534/db", reindexer.WithCreateDBIfMissing())
	if err := database.Ping(); err != nil {
		panic(err)
	}
	database.OpenNamespace("words", reindexer.DefaultNamespaceOptions(), Word{})

	iterator := database.ExecSQL("SELECT ID FROM words")
	if iterator.Count() < 1 {
		FillDB(database)
	}

	return database
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func FillDB(db *reindexer.Reindexer) {

	file, err := os.Open("static/FiveLettersWords.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		_, err := db.Insert("words", &Word{
			ID:   i,
			Word: scanner.Text(),
		})
		if err != nil {
			panic(err)
		}
		i++
	}

}
