package job

import (
	"encoding/gob"
	"fmt"
	"os"

	"../utils"
)

// Dat struct store date create from program
type Dat struct {
	Jobs map[string]DatJob
}

// NewDat create a new dat
func NewDat() *Dat {
	return &Dat{
		make(map[string]DatJob),
	}
}

// AddDat add a datjob
func (d *Dat) AddDat(e Exec) {
	d.Jobs[e.GetNameID()] = *SetDatJob(e)
}

// AddDatJob add a datjob
func (d *Dat) AddDatJob(id string, job DatJob) {
	d.Jobs[id] = job
}

// SetDatS add datjobs
func (d *Dat) SetDatS(job map[string]Exec) {
	for id, j := range job {
		d.Jobs[id] = *SetDatJob(j)
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
	fmt.Println(d)
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
	fmt.Println(d)
	return nil
}

// DatJob struct
type DatJob struct {
	Name    string
	Command string
	LogName string
	Time    string
}

// SetDatJob set a job
func SetDatJob(e Exec) *DatJob {
	d := &DatJob{
		e.Name,
		e.Command,
		e.LogName,
		e.Time,
	}
	return d
}
