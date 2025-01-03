package ent

import (
	"context"
	"fmt"

	"github.com/naoto67/entgql/ent/schema/puuid"
	"github.com/naoto67/entgql/ent/todo"
)

var prefixMap = map[puuid.ID]string{
	"TD": todo.Table,
}

// IDToType maps a pulid.ID to the underlying table.
func IDToType(ctx context.Context, id puuid.ID) (string, error) {
	if len(id) < 2 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:2]
	typ := prefixMap[prefix]
	if typ == "" {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
