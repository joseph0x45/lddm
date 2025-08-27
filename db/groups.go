package db

import (
	"fmt"
	"server/types"
)

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
	err := conn.db.Select(&groups, query)
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
