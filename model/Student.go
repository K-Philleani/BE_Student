package model

import (
	"log"
	db "student/connection"
)

type Student struct {
	StuNo      int    `json:"stu_no"`
	StuName    string `json:"stu_name"`
	StuClass   string `json:"stu_class"`
	StuSex     string `json:"stu_sex"`
	StuAge     int    `json:"stu_age"`
	StuMajor   string `json:"stu_major"`
	StuCollege string `json:"stu_college"`
	StuPhone   string `json:"stu_phone"`
	StuCity    string `json:"stu_city"`
}

// 新建学生信息
func (s *Student) AddStudent() (id int64, err error) {
	sqlStr := `insert into stu_info (stu_name, stu_class, stu_sex, stu_age, stu_major, stu_college, stu_phone, stu_city)
			values(?, ?, ?, ?, ?, ?, ?, ?)	`
	stmt, err := db.SqlDB.Prepare(sqlStr)
	if err != nil {
		return
	}
	rs, err := stmt.Exec(s.StuName, s.StuClass, s.StuSex, s.StuAge, s.StuMajor, s.StuCollege, s.StuPhone, s.StuCity)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		return
	}
	return
}

// 通过id查找数据
func (s *Student) GetStudentById() (student Student, err error) {
	sqlStr := `select stu_no, stu_name, stu_class, stu_sex, stu_age, stu_major, stu_college, stu_phone, stu_city
	from stu_info where stu_no = ?`
	stmt, err := db.SqlDB.Prepare(sqlStr)
	if err != nil {
		return
	}
	err = stmt.QueryRow(s.StuNo).Scan(
		&student.StuNo,
		&student.StuName,
		&student.StuClass,
		&student.StuSex,
		&student.StuAge,
		&student.StuMajor,
		&student.StuCollege,
		&student.StuPhone,
		&student.StuCity,
		)
	if err != nil {
		return
	}
	return

}


// 获取全部数据
func (s *Student) GetStudentAll() (students []Student, err error) {
	sqlStr := `select stu_no, stu_name, stu_class, stu_sex, stu_age, stu_major, stu_college, stu_phone, stu_city
	from stu_info`
	stmt, err := db.SqlDB.Prepare(sqlStr)
	if err != nil {
		log.Println("数据预处理失败, err: ", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println("数据查询失败, err: ", err)
		return
	}
	defer rows.Close()
	students = make([]Student, 0)
	for rows.Next() {
		var student Student
		rows.Scan(
			&student.StuNo,
			&student.StuName,
			&student.StuClass,
			&student.StuSex,
			&student.StuAge,
			&student.StuMajor,
			&student.StuCollege,
			&student.StuPhone,
			&student.StuCity,
			)
		students = append(students, student)
		if  err = rows.Err(); err != nil {
			return
		}
	}
	return
}

func (s *Student) DeleteStudent() (id int64, err error) {
	sqlStr := `delete from stu_info where stu_no = ?`
	stmt, err := db.SqlDB.Prepare(sqlStr)
	if err != nil {
		return
	}
	rs, err := stmt.Exec(s.StuNo)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		return
	}
	return
}