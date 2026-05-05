package conv

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

func ByteToPerformers(row []byte) ([]types.Performer, error) {
	performers := make([]types.Performer, 0)

	err := json.Unmarshal(row, &performers)
	if err != nil {
		return nil, err
	}

	if hasNoPerformerData(performers) {
		return nil, nil
	}

	return performers, nil
}

func ByteToTickets(row []byte) ([]types.Ticket, error) {
	tickets := make([]types.Ticket, 0)

	err := json.Unmarshal(row, &tickets)
	if err != nil {
		return nil, err
	}

	if hasNoTicketData(tickets) {
		return nil, nil
	}

	return tickets, nil
}

// hasNoPerformerData checks if the performers slice is empty or contains
// only zero-value elements (can happen when PostgreSQL returns empty JSON arrays)
func hasNoPerformerData(performers []types.Performer) bool {
	for _, p := range performers {
		if p.ID != uuid.Nil || p.Bio != "" || p.Genre != "" || p.Name != "" {
			return false
		}
	}
	return true
}

// hasNoTicketData checks if the tickets slice is empty or contains
// only zero-value elements (can happen when PostgreSQL returns empty JSON arrays)
func hasNoTicketData(tickets []types.Ticket) bool {
	for _, ticket := range tickets {
		if ticket.ID != uuid.Nil || ticket.TicketType != "" || ticket.Price != 0 || ticket.TotalTickets != 0 || ticket.AvailableTickets != 0 {
			return false
		}
	}
	return true
}

func ConvertEventSearchParams(req types.SearchEvent) entities.SearchEventsParams {
	if req.Page < 1 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.PageSize

	return entities.SearchEventsParams{
		Column1: pgtype.Text{
			String: req.Keyword,
			Valid:  req.Keyword != "",
		},
		EventDate: pgtype.Timestamp{
			Time:  req.Start,
			Valid: !req.Start.IsZero(),
		},
		EventDate_2: pgtype.Timestamp{
			Time:  req.End,
			Valid: !req.End.IsZero(),
		},
		Limit:  req.PageSize,
		Offset: offset,
	}
}
