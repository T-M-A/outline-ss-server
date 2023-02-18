package dbase

type Repo interface {
	Load() ([]ConfigRow, error)
}
