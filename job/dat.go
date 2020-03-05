package job

import (
	"encoding/gob"
	"os"

	"../utils"
)

// Dat struct store date create from program
type Dat struct {
	JobID map[string]int  // running jobs id
	Jobs  map[string]Exec // running jobs
}

// NewDat create a new dat
func NewDat() *Dat {
	return &Dat{
		make(map[string]int),
		make(map[string]Exec),
	}
}

// SaveEncode encode Dat and save to file
func (d *Dat) SaveEncode(name string) error {
	// open file
	f, err := os.OpenFile(name, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		return utils.Errs("Open Encode Dat Error: ", err)
	}

	// encode
	encode := gob.NewEncoder(f)
	if err := encode.Encode(d); err != nil {
		panic(utils.Errs("Encode Dat Error: ", err))
	}
	return nil
}

// ReadDecode read Dat and decode
func (d *Dat) ReadDecode(name string) error {
	// check file exist
	// if not exist, return
	isFile, err := utils.IsFile(name)
	if err != nil {
		return utils.Errs("Open Decode Dat Error: ", err)
	}
	if !isFile {
		return nil
	}

	// open && decode
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return utils.Errs("Decode Dat Error: ", err)
	}
	decode := gob.NewDecoder(f)
	if err := decode.Decode(d); err != nil {
		panic(utils.Errs("Decode Dat Error: ", err))
	}
	return nil
}

// SaveEncodeDat save dat by JobID & Jobs
func SaveEncodeDat(name string, jobID map[string]int, jobs map[string]Exec) error {
	d := &Dat{
		jobID,
		jobs,
	}
	return d.SaveEncode(name)
}
