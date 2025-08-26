package db

import (
	"fmt"
	"log"
	"server/types"

	"github.com/google/uuid"
)

func (conn *DBConnection) SeedGroups() {
	log.Println("[INFO]: Seeding Groups")
	groups := []types.Group{
		{
			ID:      uuid.NewString(),
			Name:    "La Verte",
			Picture: "https://ik.imagekit.io/lddm/la_verte2.jpg?updatedAt=1756204812699",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Atadi Fionfion",
			Picture: "https://ik.imagekit.io/lddm/atadi_fionfion.jpg?updatedAt=1756204709945",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Koklo Blend",
			Picture: "https://ik.imagekit.io/lddm/koklo_blend.jpg?updatedAt=1756204709840",
		},
		{
			ID:      uuid.NewString(),
			Name:    "La Totale",
			Picture: "https://ik.imagekit.io/lddm/la_totale2.jpg?updatedAt=1756204709814",
		},
		{
			ID:      uuid.NewString(),
			Name:    "L'Exotique",
			Picture: "https://ik.imagekit.io/lddm/exotique.jpg?updatedAt=1756204709634",
		},
		{
			ID:      uuid.NewString(),
			Name:    "L'Exquise",
			Picture: "https://ik.imagekit.io/lddm/exquise2.jpg?updatedAt=1756204709558",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Persilmix",
			Picture: "https://ik.imagekit.io/lddm/persilmix.jpg?updatedAt=1756204709441",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Gbolan Blend",
			Picture: "https://ik.imagekit.io/lddm/gbolan_blend.jpg?updatedAt=1756204709171",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Ail",
			Picture: "https://ik.imagekit.io/lddm/ail.jpg?updatedAt=1756204709015",
		},
		{
			ID:      uuid.NewString(),
			Name:    "Atadi Delices",
			Picture: "https://ik.imagekit.io/lddm/atadi_delices.jpg?updatedAt=1756204707833",
		},
	}
	const query = `
    insert into groups (
      id, name, picture
    )
    values (
      :id, :name, :picture
    )
    on conflict(name) do nothing
  `
	_, err := conn.db.NamedExec(
		query, groups,
	)
	if err != nil {
		log.Panicf("Seeding failed for groups: %s\n", err.Error())
	}
	log.Println("[INFO]: Done!")
}

func (conn *DBConnection) CreateGroup(group *types.Group) error {
	const query = `
    insert into groups(
      id, name, picture
    )
    values (
      :id, :name, :picture
    )
  `
	_, err := conn.db.NamedExec(query, group)
	if err != nil {
		return fmt.Errorf("failed to create group: %w", err)
	}
	return nil
}

func (conn *DBConnection) FetchGroups() ([]types.Group, error) {
	const query = `
    select * from groups
  `
	groups := make([]types.Group, 0)
	err := conn.db.Select(groups, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups: %w", err)
	}
	return groups, nil
}

func (conn *DBConnection) DeleteGroup(groupID string) error {
	const query = `
    delete from groups where id=$1
  `
	_, err := conn.db.Exec(query, groupID)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}
	return nil
}
