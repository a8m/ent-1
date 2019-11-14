// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromRows scans the sql response data into Group.
func (gr *Group) FromRows(rows *sql.Rows) error {
	var scangr struct {
		ID int
	}
	// the order here should be the same as in the `group.Columns`.
	if err := rows.Scan(
		&scangr.ID,
	); err != nil {
		return err
	}
	gr.ID = strconv.Itoa(scangr.ID)
	return nil
}

// QueryUsers queries the users edge of the Group.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{gr.config}).QueryUsers(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (gr *Group) id() int {
	id, _ := strconv.Atoi(gr.ID)
	return id
}

// Groups is a parsable slice of Group.
type Groups []*Group

// FromRows scans the sql response data into Groups.
func (gr *Groups) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scangr := &Group{}
		if err := scangr.FromRows(rows); err != nil {
			return err
		}
		*gr = append(*gr, scangr)
	}
	return nil
}

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
