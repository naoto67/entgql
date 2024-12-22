// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package entgql

import (
	"time"

	"github.com/naoto67/entgql/ent/todo"
)

// TodoWhereInput is used for filtering Todo objects.
// Input was generated by ent.
type TodoWhereInput struct {
	Not *TodoWhereInput   `json:"not,omitempty"`
	And []*TodoWhereInput `json:"and,omitempty"`
	Or  []*TodoWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *int  `json:"id,omitempty"`
	IDNeq   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGt    *int  `json:"idGT,omitempty"`
	IDGte   *int  `json:"idGTE,omitempty"`
	IDLt    *int  `json:"idLT,omitempty"`
	IDLte   *int  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// status field predicates
	Status      *todo.Status  `json:"status,omitempty"`
	StatusNeq   *todo.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []todo.Status `json:"statusIn,omitempty"`
	StatusNotIn []todo.Status `json:"statusNotIn,omitempty"`
	// priority field predicates
	Priority      *int  `json:"priority,omitempty"`
	PriorityNeq   *int  `json:"priorityNEQ,omitempty"`
	PriorityIn    []int `json:"priorityIn,omitempty"`
	PriorityNotIn []int `json:"priorityNotIn,omitempty"`
	PriorityGt    *int  `json:"priorityGT,omitempty"`
	PriorityGte   *int  `json:"priorityGTE,omitempty"`
	PriorityLt    *int  `json:"priorityLT,omitempty"`
	PriorityLte   *int  `json:"priorityLTE,omitempty"`
	// text field predicates
	Text             *string  `json:"text,omitempty"`
	TextNeq          *string  `json:"textNEQ,omitempty"`
	TextIn           []string `json:"textIn,omitempty"`
	TextNotIn        []string `json:"textNotIn,omitempty"`
	TextGt           *string  `json:"textGT,omitempty"`
	TextGte          *string  `json:"textGTE,omitempty"`
	TextLt           *string  `json:"textLT,omitempty"`
	TextLte          *string  `json:"textLTE,omitempty"`
	TextContains     *string  `json:"textContains,omitempty"`
	TextHasPrefix    *string  `json:"textHasPrefix,omitempty"`
	TextHasSuffix    *string  `json:"textHasSuffix,omitempty"`
	TextEqualFold    *string  `json:"textEqualFold,omitempty"`
	TextContainsFold *string  `json:"textContainsFold,omitempty"`
	// value field predicates
	Value      *int  `json:"value,omitempty"`
	ValueNeq   *int  `json:"valueNEQ,omitempty"`
	ValueIn    []int `json:"valueIn,omitempty"`
	ValueNotIn []int `json:"valueNotIn,omitempty"`
	ValueGt    *int  `json:"valueGT,omitempty"`
	ValueGte   *int  `json:"valueGTE,omitempty"`
	ValueLt    *int  `json:"valueLT,omitempty"`
	ValueLte   *int  `json:"valueLTE,omitempty"`
	// parent edge predicates
	HasParent     *bool             `json:"hasParent,omitempty"`
	HasParentWith []*TodoWhereInput `json:"hasParentWith,omitempty"`
	// children edge predicates
	HasChildren     *bool             `json:"hasChildren,omitempty"`
	HasChildrenWith []*TodoWhereInput `json:"hasChildrenWith,omitempty"`
	CreatedToday    *bool             `json:"createdToday,omitempty"`
}