package Model

//注册时先检验用户是否已经存在

func CheckUserExist(StudentId string)(bool,error){
	var p2 []User
	DB.Find(&p2,"student_id=?",StudentId)
	var count int
	count=len(p2)
	return count>0,nil
}
//注册时将用户信息存储到数据库中

func InsertUser(StudentId string,Name string,Password string,Department string)(err error)  {
	DB.Create(&User{
		StudentId: StudentId,
		Name: Name,
		Password: Password,
		Department:Department,
	})
	if err!=nil{
		return err
	}
	return nil
}
