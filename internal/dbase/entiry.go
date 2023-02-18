package dbase

type ConfigRow struct {
	ID     string `db:"id"`
	Port   int    `db:"port"`
	Cipher string `db:"cipher"`
	Secret string `db:"secret"`
}
