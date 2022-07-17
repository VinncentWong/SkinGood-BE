package util

func HandleError(e error) {
	panic(e.Error())
}
