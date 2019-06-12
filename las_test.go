//(c) softland 2019
//softlandia@gmail.com
package laslib

import (
	"testing"
)

//CodePageDetect
func TestReadWellParam(t *testing.T) {
	las := NewLas()
	err := las.ReadWellParam("NULL.\t-999.99 : desc")
	res, err := CodePageDetect("test_files\\test.txt", "~X~")
	if err != nil {
		t.Errorf("<CodePageDetect> on file '%s' return error: %v", "test.txt", err)
	}
	if res != Cp866 {
		t.Errorf("<CodePageDetect> on file '%s' expected 866 got: %s", "test.txt", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\test.txt")
	if res != CpWindows1251 {
		t.Errorf("<CodePageDetect> on file '%s' expected 1251 got: %s", "test.txt", CodePageAsString(res))
	}

	_, err = CodePageDetect("-.-")
	if err == nil {
		t.Errorf("<CodePageDetect> on file '-.-' must return error, but return nil")
	}

	res, _ = CodePageDetect("test_files\\test2.txt")
	if res != CpEmpty {
		t.Errorf("<CodePageDetect> on file 'test2.txt' expect CpEmpty got: %s", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\test3.txt")
	if (res != CpEmpty) || (err != nil) {
		t.Errorf("<CodePageDetect> on file 'test3.txt' expect CpEmpty and no error got: %s and %v", CodePageAsString(res), err)
	}
}
