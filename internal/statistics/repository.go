package statistics

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
	"time"

	"gorm.io/datatypes"
)

type StatisticsRepository struct {
	Db *db.Db
}

func NewStatisticsRepository(db *db.Db) *StatisticsRepository {
	return &StatisticsRepository{
		Db: db,
	}
}

func (repo *StatisticsRepository) AddClick(link_id uint) {
	var statistic models.Stats
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&statistic, "link_id = ?", link_id)
	if statistic.ID == 0 {
		repo.Db.Create(&models.Stats{
			Link_id: link_id,
			Clicks:  1,
			Date:    currentDate,
		})
	} else {
		statistic.Clicks += 1
		repo.Db.Save(&statistic)
	}
}
