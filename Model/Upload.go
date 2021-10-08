package Model

func UploadPicpath(filePath string,NowStudent string) {
	DB.Model(&User{}).Where("student_id=?",NowStudent).Update("pic_path",filePath)
}