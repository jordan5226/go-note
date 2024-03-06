package main

import (
	"jyutil"
)

func swap(strX string, strY string) (string, string) {
	return strY, strX
}

func isValid(strContent string) (bool, string) {
	if strContent == "" {
		return false, "Invalid Content!"
	}

	return true, ""
}

func validateForm(lpForm *map[string]string) (bool, string) {
	for _, strContent := range *lpForm {
		if bRet, strErr := isValid(strContent); !bRet || strErr != "" {
			return bRet, strErr
		}
	}

	return true, ""
}

func submitForm(lpForm *map[string]string) string {
	if lpForm == nil {
		return "Invalid Form Ptr!"
	} else if bRet, strErr := validateForm(lpForm); !bRet || strErr != "" {
		return strErr
	}

	return ""
}

func testSwitch(iType int) {
	switch iType {
	case 1:
		jyutil.Output("1")
	case 2:
		jyutil.Output("2")
	case 3:
		jyutil.Output("3")
		fallthrough
	case 4:
		jyutil.Output("4")
	case 5, 6:
		jyutil.Output("5 or 6")
	default:
		jyutil.Output("default")
	}
}

func useDeclare() {
	jyutil.Output("====== Declare =======")
	// Declare 1
	/*
		var str string
		str = "Hello World 1!"
		jyutil.Output(str)
	*/

	// Declare 2
	var str2 string = "Hello World 2!"
	jyutil.Output(str2)

	// Declare 3
	str3 := "Hello World 3!"
	jyutil.Output(str3)
}

func useFunction() {
	jyutil.Output("======= Function ======")

	a, b := swap("Hello", "World")

	jyutil.OutputDebugString("%s %s\n", a, b)
}

func useIf() {
	jyutil.Output("========= if ==========")

	if name := "john"; name != "" {
		jyutil.Output(name)
	}
}

func useMapFor() {
	jyutil.Output("===== map, for ====")

	mapForm := map[string]string{
		"Name":   "Jordan",
		"Age":    "18",
		"Gender": "Male",
		"Note":   "",
	}

	jyutil.Output("Submit form result:")
	if strErr := submitForm(&mapForm); strErr != "" {
		jyutil.Output(strErr)
	} else {
		jyutil.Output(mapForm)
	}

	jyutil.Output("Form data:")
	jyutil.Output(mapForm)
}

func useSlice() {
	jyutil.Output("======= Slice =======")

	var arrName []string
	var arrGender []string = []string{"Male", "Female", "Male"}
	arrAge := []int{18, 26, 39}

	jyutil.Output("Name:", arrName)
	jyutil.Output("Gender:", arrGender)
	jyutil.Output("Age:", arrAge)

	jyutil.Output("Push data to arrName...")
	arrName = append(arrName, "Jordan")
	jyutil.Output("Name:", arrName)
	jyutil.Output("Push more data to arrName...")
	arrName = append(arrName, "John", "Alice", "Alpha", "Beta", "Gamma")
	jyutil.Output("Name:", arrName)
	jyutil.Output("Length of Name Array: ", len(arrName))
	jyutil.Output("arrName[1]: ", arrName[1])
	for iIdx, strName := range arrName {
		jyutil.OutputDebugString("arrName[%d]=\"%s\"\n", iIdx, strName)
	}
}

func useSwitch() {
	jyutil.Output("======== switch ========")

	jyutil.Output("testSwitch 1")
	testSwitch(1)
	jyutil.Output("testSwitch 3")
	testSwitch(3)
	jyutil.Output("testSwitch 5 or 6")
	testSwitch(5)
	testSwitch(6)
}
