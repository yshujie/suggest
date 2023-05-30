package model

// Entry 词条，词库中的最小元素，包含索引和对象
type Entry struct {
	Index  EntryIndex
	Object EntryObject
}

// EntryIndex 词条索引
type EntryIndex interface{}

// EntryObject 词条对象
type EntryObject interface {
	GetCode() string
	ToString() string
	ToJson() string
}

type ChildEntryObject struct {
	Code   string
	Id     int
	Name   string
	Mobile string
}

func (o *ChildEntryObject) GetCode() string {
	return o.Code
}

func (o *ChildEntryObject) ToString() string {
	return ""
}

func (o *ChildEntryObject) ToJson() string {
	return ""
}

type PatientEntryObject struct {
	Code   string
	Id     int
	Name   string
	Mobile string
}

func (o *PatientEntryObject) GetCode() string {
	return o.Code
}

func (o *PatientEntryObject) ToString() string {
	return ""
}

func (o *PatientEntryObject) ToJson() string {
	return ""
}

type DoctorEntryObject struct {
	Code     string
	Id       int
	Name     string
	Mobile   string
	Hospital string
}

func (o *DoctorEntryObject) GetCode() string {
	return o.Code
}

func (o *DoctorEntryObject) ToString() string {
	return ""
}

func (o *DoctorEntryObject) ToJson() string {
	return ""
}

type HospitalEntryObject struct {
	Code   string
	Id     int
	Name   string
	Mobile string
}

func (o *HospitalEntryObject) GetCode() string {
	return o.Code
}

func (o *HospitalEntryObject) ToString() string {
	return ""
}

func (o *HospitalEntryObject) ToJson() string {
	return ""
}
