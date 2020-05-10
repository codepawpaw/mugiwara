package incident

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitIncidentRepository(Connection *sql.DB) pRepo.IncidentRepository {
	return &IncidentRepository{
		Connection: Connection,
	}
}

type IncidentRepository struct {
	Connection *sql.DB
}

func (o *IncidentRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Incident, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Incident, 0)
	for rows.Next() {
		data := new(models.Incident)

		err := rows.Scan(
			&data.ID,
			&data.CityName,
			&data.Province,
			&data.Nation,
			&data.Description,
			&data.Date,
			&data.Lat,
			&data.Lang,
			&data.UserId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *IncidentRepository) Create(ctx context.Context, incident *models.Incident) (*models.Incident, error) {
	query := "Insert incidents SET cityName=?, province=?, nation=?, description=?, date=?, lat=?, lang=?, user_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.Incident{}, err
	}

	incidentResponse, err := stmt.ExecContext(ctx, incident.CityName, incident.Province, incident.Nation, incident.Description, incident.Date, incident.Lat, incident.Lang, incident.UserId)
	defer stmt.Close()

	if err != nil {
		return &models.Incident{}, err
	}

	incidentId, _ := incidentResponse.LastInsertId()
	incident.ID = incidentId

	return incident, err
}

func (m *IncidentRepository) GetByCityName(ctx context.Context, cityName string) ([]*models.Incident, error) {
	query := "Select * From incidents where cityName=?"

	return m.fetch(ctx, query, cityName)
}

func (m *IncidentRepository) Update(ctx context.Context, p *models.Incident) (*models.Incident, error) {
	query := "Update incidents set description=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Description,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *IncidentRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From incidents Where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
