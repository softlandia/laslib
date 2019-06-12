package main

import (
	"fmt"
	"os"

	"github.com/softlandia/laslib"
)

func main() {
	//test file "1.las"
	las := laslib.NewLas()
	n, err := las.Open("1.las")
	if err != nil {
		fmt.Printf("on open file: '1.las' error:%v", err)
		os.Exit(1)
	}
	if n == 7 {
		fmt.Println("read 1.las OK", err)
	} else {
		fmt.Printf("TEST read 1.las ERROR, n = %d, must 7\n", n)
		fmt.Println(err)
	}
	las.SaveWarning("1.warning.md")
	if las.WarningCount() > 0 {
		fmt.Printf("warning count: %d,\nsave warningto file: 1.warning.md\n", las.WarningCount())
	}

	err = las.SetNull(-999.99)
	fmt.Println("set new null value done, error: ", err)

	err = las.Save("-1.las")
	if err != nil {
		fmt.Println("TEST save -1.las ERROR: ", err)
	} else {
		fmt.Println("TEST save -1.las OK")
	}

	las = nil
	las = laslib.NewLas()
	fmt.Println("reopen file: -1.las")
	n, err = las.Open("-1.las")
	if (n == 7) && (las.Null == -999.99) {
		fmt.Println("TEST read -1.las OK")
		fmt.Println(err)
	} else {
		fmt.Println("TEST read -1.las ERROR")
		fmt.Println("NULL not -999.99 or count dept points != 7")
		fmt.Println(err)
	}
}
