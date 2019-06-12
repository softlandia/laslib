package laslib

const (
	defWarningCount = 10
	warningUNDEF    = 0
	directOnRead    = 1
	directOnWrite   = 2
)

//TWarning - class to store warning
type TWarning struct {
	direct  int    // 0 - undefine (warningUNDEF), 1 - on read (directOnRead), 2 - on write (directOnWrite)
	section int    // 0 - undefine (warningUNDEF), lasSecVertion, lasSecWellInfo, lasSecCurInfo, lasSecData
	line    int    // number of line in source file
	desc    string // description of warning
}

//TLasWarnings - class to store and manipulate warnings
type TLasWarnings struct {
}
