package models

import "context"

func (s *Conn) CreateJob(ctx context.Context, nj NewJob, CompanyID uint) (Job, error) {
	j := Job{
		JobName: nj.JobName,
	}
	tx := s.db.WithContext(ctx).Create(&j)
	if tx.Error != nil {
		return Job{}, tx.Error
	}
	return j, nil
}
func (s *Conn) ViewJob(ctx context.Context, userId string) ([]Job, error) {
	var inv = make([]Job, 0, 10)
	tx := s.db.WithContext(ctx).Where("user_id = ?", userId)
	err := tx.Find(&inv).Error
	if err != nil {
		return nil, err
	}
	return inv, nil

}
