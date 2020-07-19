package data

/*
 * id
 * name
 * education
 * age
 * marriage
 * height
 * sex
 */

import (
	"log"

	"github.com/crawers/go-crawers/zhenaiwang/save/data"
	"github.com/crawers/go-crawers/zhenaiwang/save/model"
)

func SaveData(user *model.Profile) error {
	db, err := data.InitMysql("root", "Damion123$", "127.0.0.1:3306")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO profile (id, name, education, age, marriage, height, sex) values(?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE education=values(education), age=values(age), marriage=values(marriage), height=values(height), sex=values(sex);")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.MemberId, user.NickName, user.Education, user.Age, user.Marriage, user.Height, user.Sex)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	affeted, _ := res.RowsAffected()
	log.Printf("LastInsertId: %d, RowAffected: %d\n", id, affeted)
	return nil
}
