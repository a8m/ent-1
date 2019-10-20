// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
)

// Postgres is a postgres migration driver.
type Postgres struct {
	dialect.Driver
	version string
}

// init loads the Postgres version from the database for later use in the migration process.
// It returns an error if the server version is lower than v10.
func (d *Postgres) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW server_version_num", []interface{}{}, rows); err != nil {
		return fmt.Errorf("querying server version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("server_version_num variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("scanning version: %v", err)
	}
	if len(version) < 6 {
		return fmt.Errorf("malformed version: %s", version)
	}
	d.version = fmt.Sprintf("%s.%s.%s", version[:2], version[2:4], version[4:])
	if compareVersions(d.version, "10.0.0") == -1 {
		return fmt.Errorf("unsupported postgres version: %s", d.version)
	}
	return nil
}

// tableExist checks if a table exists in the database and current schema.
func (d *Postgres) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// tableExist checks if a foreign-key exists in the current schema.
func (d *Postgres) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("constraint_type", "FOREIGN KEY").And().EQ("constraint_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// setRange sets restart the identity column to the given offset. Used by the universal-id option.
func (d *Postgres) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN id RESTART WITH %d", name, value), []interface{}{}, new(sql.Result))
}

// table loads the current table description from the database.
func (d *Postgres) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.Postgres).
		Select("column_name", "data_type", "character_maximum_length", "is_nullable", "column_default").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("postgres: reading table description %v", err)
	}
	// call `Close` in cases of failures (`Close` is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, err
		}
		t.AddColumn(c)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("closing rows %v", err)
	}
	idxs, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// Populate the index information to the table and its columns.
	// We do it manually, because PK and uniqueness information does
	// not exist when querying the INFORMATION_SCHEMA.COLUMNS above.
	for _, idx := range idxs {
		switch {
		case idx.primary:
			for _, name := range idx.columns {
				c, ok := t.column(name)
				if !ok {
					return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
				}
				c.Key = PrimaryKey
				t.PrimaryKey = append(t.PrimaryKey, c)
			}
		case idx.Unique && len(idx.columns) == 1:
			name := idx.columns[0]
			c, ok := t.column(name)
			if !ok {
				return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
			}
			c.Key = UniqueKey
			c.Unique = true
		default:
			t.AddIndex(idx.Name, idx.Unique, idx.columns)
		}
	}
	return t, nil
}

// indexesQuery holds a query format for retrieving
// table indexes of the current schema.
const indexesQuery = `
SELECT i.relname AS index_name,
       a.attname AS column_name,
       idx.indisprimary AS primary,
       idx.indisunique AS unique
FROM pg_class t,
     pg_class i,
     pg_index idx,
     pg_attribute a,
     pg_namespace n
WHERE t.oid = idx.indrelid
  AND i.oid = idx.indexrelid
  AND n.oid = t.relnamespace
  AND a.attrelid = t.oid
  AND a.attnum = ANY(idx.indkey)
  AND t.relkind = 'r'
  AND n.nspname = CURRENT_SCHEMA()
  AND t.relname = '%s';
`

func (d *Postgres) indexes(ctx context.Context, tx dialect.Tx, table string) (Indexes, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, fmt.Sprintf(indexesQuery, table), []interface{}{}, rows); err != nil {
		return nil, fmt.Errorf("querying indexes for table %s", table)
	}
	defer rows.Close()
	var (
		idxs  Indexes
		names = make(map[string]*Index)
	)
	for rows.Next() {
		var (
			name, column    string
			unique, primary bool
		)
		if err := rows.Scan(&name, &column, &primary, &unique); err != nil {
			return nil, fmt.Errorf("scanning index description: %v", err)
		}
		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: unique, primary: primary}
			idxs = append(idxs, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	return idxs, nil
}

// maxCharSize defines the maximum size of limited character types in Postgres (10 MB).
const maxCharSize = 10 << 20

// scanColumn scans the information a column from column description.
func (d *Postgres) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		maxlen   sql.NullInt64
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &maxlen, &nullable, &defaults); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch c.typ {
	case "boolean":
		c.Type = field.TypeBool
	case "smallint":
		c.Type = field.TypeInt16
	case "integer":
		c.Type = field.TypeInt32
	case "bigint":
		c.Type = field.TypeInt64
	case "real":
		c.Type = field.TypeFloat32
	case "double precision":
		c.Type = field.TypeFloat64
	case "text":
		c.Type = field.TypeString
		c.Size = maxCharSize + 1
	case "character":
		c.Type = field.TypeString
		c.Size = maxlen.Int64
	case "timestamp with time zone":
		c.Type = field.TypeTime
	case "bytea":
		c.Type = field.TypeBytes
	case "jsonb":
		c.Type = field.TypeJSON
	}
	switch {
	case !defaults.Valid:
		return nil
	case strings.Contains(defaults.String, "::"):
		parts := strings.Split(defaults.String, "::")
		defaults.String = strings.Trim(parts[0], "'")
		fallthrough
	default:
		return c.ScanDefault(defaults.String)
	}
}
