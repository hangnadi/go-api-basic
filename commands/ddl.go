package commands

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ddlFile represents a Data Definition Language (DDL) file
// Given the file naming convention 001-user.sql, the numbers up to
// the first dash are extracted, converted to an int and added to the
// fileNumber field to make the struct sortable using the sort package.
type ddlFile struct {
	filename   string
	fileNumber int
}

// newDDLFile initializes a DDLFile struct. File naming convention
// should be 001-user.sql where 001 represents the file number order
// to be processed
func newDDLFile(f string) (ddlFile, error) {
	i := strings.Index(f, "-")
	fileNumber := f[:i]
	fn, err := strconv.Atoi(fileNumber)
	if err != nil {
		return ddlFile{}, err
	}

	return ddlFile{filename: f, fileNumber: fn}, nil
}

func (df ddlFile) String() string {
	return fmt.Sprintf("%s: %d", df.filename, df.fileNumber)
}

// readDDLFiles reads and returns sorted DDL files from the
// ./scripts/ddl/db-deploy/up or ./scripts/ddl/db-deploy/down directory
func readDDLFiles(dir string) ([]ddlFile, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ddlFiles []ddlFile
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		df, err := newDDLFile(file.Name())
		if err != nil {
			return nil, err
		}
		ddlFiles = append(ddlFiles, df)
	}

	sort.Sort(byFileNumber(ddlFiles))

	return ddlFiles, nil
}

// byFileNumber implements sort.Interface for []ddlFile based on
// the fileNumber field.
type byFileNumber []ddlFile

// Len returns the length of elements in the ByFileNumber slice for sorting
func (bfn byFileNumber) Len() int { return len(bfn) }

// Swap sets up the elements to be swapped for the ByFileNumber slice for sorting
func (bfn byFileNumber) Swap(i, j int) { bfn[i], bfn[j] = bfn[j], bfn[i] }

// Less is the sorting logic for the ByFileNumber slice
func (bfn byFileNumber) Less(i, j int) bool { return bfn[i].fileNumber < bfn[j].fileNumber }

// PSQLArgs takes a slice of DDL files to be executed and builds a
// sequence of command line arguments using the appropriate flags
// psql needs to execute files. The flags returned for psql are as follows:
//
// -w flag is used to never prompt for a password as we are running this as a script
//    see
// -d flag sets the database connection using a Connection URI string.
// -f flag is sent before each file to tell it to process the file
func PSQLArgs(up bool) ([]string, error) {
	dir := "./scripts/ddl/db-deploy"
	if up {
		dir += "/up"
	} else {
		dir += "/down"
	}

	ddlFiles, err := readDDLFiles(dir)
	if err != nil {
		return nil, err
	}

	if len(ddlFiles) == 0 {
		return nil, fmt.Errorf("there are no DDL files to process in %s", dir)
	}

	flgs, err := newFlags([]string{"server"})
	if err != nil {
		return nil, err
	}

	args := []string{"-w", "-d", newPostgreSQLDSN(flgs).ConnectionURI(), "-c", "select current_database(), current_user, version()"}

	for _, file := range ddlFiles {
		args = append(args, "-f")
		args = append(args, dir+"/"+file.filename)
	}

	return args, nil
}
