package core

// 退完了返回 true
func Drop(user *User) bool {
	user.e = nil
	// 无需退课，速速 return
	if len(user.getTarget()) == 0 {
		if user.GetCorn() != nil {
			user.GetCorn().Stop()
		}
		return true
	}

	courses, e := user.FindCourse().PrintFireCourseList().DropCourse()
	if e != nil {
		log.Print(e)
		return true
	}
	if len(courses) == 0 {
		log.Println("暂未退课成功")
		// 主动停止
	} else {
		log.Println("已退如下课程：")
		log.Println(courses)
	}

	log.Println()
	return len(user.getTarget()) == 0
}
