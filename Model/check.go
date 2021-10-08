package Model

func CheckName(chooseName string)(flag string ) {
	if correctName==chooseName {
		flag="You are right!"
	}else {
		flag="You are wrong!"
	}
	return flag
}