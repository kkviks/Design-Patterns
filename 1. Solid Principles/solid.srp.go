package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

/*
A type should have one primary responsibility, and as a result it should have
one reason to change that reason being somehow related to it's primary responsibility.

## Break : Add another function which deal with another concern
### Separation of concern: different concerns or different problems that the
system solves have to reside in different constructs whether attached to different structs or
reside in different packages

-- Ansi patttern :
God object : When you take everything in the kitchen sink that you are doing and
put it in a single package

*/

var entryCount = 0

type Journal struct {
	// Keep entries and manage those entries
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// breaks srp: Adding functions dealing with other concerns
// Responsibility of the journal is not to deal with the persistence

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// We can have a separate function, so other types can also use it for persistence
var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// A better approach
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main_() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))

	// separate function
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
