package data

import (
	"campyuk-api/features/camp"
	"errors"
	"log"

	"gorm.io/gorm"
)

type campData struct {
	db *gorm.DB
}

func New(db *gorm.DB) camp.CampData {
	return &campData{db: db}
}

func (cd *campData) Add(userID uint, newCamp camp.Core) error {
	// Create camp
	cm := ToData(userID, newCamp)
	tx := cd.db.Create(&cm)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	// Inserting image to camp
	cim := ToImageData(cm.ID, newCamp.Images)
	for _, v := range cim {
		// Kenapa pakai exec ketimbang batch create dari gorm? dikarenakan kena panic error.
		tx = tx.Exec("INSERT INTO images(camp_id, image) VALUES(?, ?)", v.CampID, v.Image)
		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
	}

	tx.Commit()

	return nil
}

func (cd *campData) List(userID uint, role string, limit int, offset int) (int, []camp.Core, error) {
	var cm []CampModel

	switch role {
	case "host":
		res, err := cd.listCampHost(userID, limit, offset)
		if err != nil {
			return 0, nil, err
		}
		cm = res
	case "admin":
		res, err := cd.listCampAdmin(limit, offset)
		if err != nil {
			return 0, nil, err
		}
		cm = res
	default:
		res, err := cd.listCampUser(limit, offset)
		if err != nil {
			return 0, nil, err
		}
		cm = res
	}

	return len(cm), ToListCampCore(cm), nil
}

func (cd *campData) GetByID(userID uint, campID uint) (camp.Core, error) {
	c := CampModel{}
	qc := "SELECT camps.id, camps.verification_status, users.fullname, camps.title, camps.price, camps.description, camps.latitude, camps.longitude, camps.distance, camps.address, camps.city, camps.document FROM camps JOIN users ON users.id = camps.host_id WHERE camps.id = ? AND camps.deleted_at IS NULL"
	tx := cd.db.Raw(qc, campID).First(&c)
	if tx.Error != nil {
		return camp.Core{}, tx.Error
	}

	images := []Image{}
	tx = tx.Raw("SELECT * FROM images WHERE camp_id = ? AND deleted_at IS NULL", campID).Find(&images)
	if tx.Error != nil {
		log.Println(tx.Error)
		log.Println("no image found in camp")
	}

	items := []CampItemModel{}
	tx = tx.Raw("SELECT * FROM items WHERE camp_id = ? AND deleted_at IS NULL", campID).Find(&items)
	if tx.Error != nil {
		log.Println(tx.Error)
		log.Println("no item found in camp")
	}

	c.Images = images
	c.Items = items

	return ToCampCore(c), nil
}

func (cd *campData) Update(userID uint, campID uint, updateCamp camp.Core) error {
	campData := ToData(userID, updateCamp)
	tx := cd.db.Where("id = ? AND host_id = ?", campID, userID).Updates(&campData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
func (cd *campData) Delete(userID uint, campID uint) error {
	data := Camp{}
	qry := cd.db.Where("host_id = ?", userID).Delete(&data, campID)
	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no camp deleted")
	}

	return nil
}
func (cd *campData) RequestAdmin(campID uint, status string) error {
	tx := cd.db.Exec("UPDATE camps SET verification_status=? WHERE id = ?", status, campID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		log.Println("no rows affected")
		return errors.New("status update failed")
	}

	return nil
}

// ------------------------
// Functions that are not include in the contract
// ------------------------

func (cd *campData) listCampUser(limit int, offset int) ([]CampModel, error) {
	cm := []CampModel{}
	// Select camp
	qc := "SELECT camps.id, camps.verification_status, users.fullname, camps.title, camps.price, camps.distance, camps.city FROM camps JOIN users ON users.id = camps.host_id WHERE camps.verification_status = 'ACCEPTED' AND camps.deleted_at IS NULL ORDER BY camps.id DESC LIMIT ? OFFSET ?"
	tx := cd.db.Raw(qc, limit, offset).Find(&cm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Find camp image
	for i := range cm {
		ci := Image{}
		tx = tx.Raw("SELECT id, image FROM images WHERE camp_id = ? AND deleted_at IS NULL ORDER BY id ASC", cm[i].ID).First(&ci)
		if tx.Error != nil {
			return nil, tx.Error
		}
		cm[i].Images = append(cm[i].Images, ci)
	}

	return cm, nil
}

func (cd *campData) listCampHost(userID uint, limit int, offset int) ([]CampModel, error) {
	cm := []CampModel{}
	// Select camp
	qc := "SELECT camps.id, camps.verification_status, users.fullname, camps.title, camps.price, camps.distance,camps.city FROM camps JOIN users ON users.id = camps.host_id WHERE users.id = ? AND camps.deleted_at IS NULL ORDER BY camps.id DESC LIMIT ? OFFSET ?"
	tx := cd.db.Raw(qc, userID, limit, offset).Find(&cm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Find camp image
	for i := range cm {
		ci := Image{}
		tx = tx.Raw("SELECT id, image FROM images WHERE camp_id = ? AND deleted_at IS NULL ORDER BY id ASC", cm[i].ID).First(&ci)
		if tx.Error != nil {
			return nil, tx.Error
		}
		cm[i].Images = append(cm[i].Images, ci)
	}

	return cm, nil
}

func (cd *campData) listCampAdmin(limit int, offset int) ([]CampModel, error) {
	cm := []CampModel{}
	// Select camp
	qc := "SELECT camps.id, camps.verification_status, users.fullname, camps.title, camps.price, camps.distance,camps.city FROM camps JOIN users ON users.id = camps.host_id WHERE camps.verification_status = 'PENDING' AND camps.deleted_at IS NULL ORDER BY camps.id DESC LIMIT ? OFFSET ?"
	tx := cd.db.Raw(qc, limit, offset).Find(&cm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Find camp image
	for i := range cm {
		ci := Image{}
		tx = tx.Raw("SELECT id, image FROM images WHERE camp_id = ? AND deleted_at IS NULL ORDER BY id ASC", cm[i].ID).First(&ci)
		if tx.Error != nil {
			return nil, tx.Error
		}
		cm[i].Images = append(cm[i].Images, ci)
	}

	return cm, nil
}
