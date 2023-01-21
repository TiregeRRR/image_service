package imagefile

import (
	"github.com/TiregeRRR/image_service/internal/model"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/TiregeRRR/image_service/internal/pkg/storage"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type In struct {
	fx.In

	DB     *gorm.DB
	Logger *zap.Logger
	Config config.Config
}

type Manager struct {
	logger  *zap.Logger
	db      *gorm.DB
	storage *storage.DiskStorage
}

func New(in In) *Manager {
	return &Manager{
		logger:  in.Logger,
		db:      in.DB,
		storage: storage.New(in.Config.StorageDir),
	}
}

func (mgr *Manager) Upsert(m *model.Image) (*model.Image, error) {
	path, err := mgr.storage.Save(m.Name, m.Data)
	if err != nil {
		return nil, err
	}
	m.Path = path
	mgr.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).Create(m)
	m = &model.Image{Name: m.Name}
	mgr.db.Find(m)
	return m, nil
}

func (mgr *Manager) GetImages() ([]*model.Image, error) {
	users := []*model.Image{}
	res := mgr.db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (mgr *Manager) GetImageData(name string) (*model.Image, error) {
	data, err := mgr.storage.Load(name)
	if err != nil {
		return nil, err
	}
	return &model.Image{
		Name: name,
		Data: data,
	}, nil
}
