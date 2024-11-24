package data

import (
	"fmt"
	"time"

	up "github.com/upper/db/v4"
)

// AppStat struct
type AppStat struct {
    ID        int       `db:"id,omitempty"`
    Query           string    `db:"query"`               // The query run by the user
    ExecutionTimeMs int       `db:"execution_time_ms"`   // Execution time in milliseconds
    FileName        string    `db:"file_name"`           // The name of the JSON file associated with the query
    Success         bool      `db:"success"`            // Whether the query ran successfully
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *AppStat) Table() string {
    return "public.general_app_stats"
}

// GetAll gets all records from the database, using upper
func (t *AppStat) GetAll(condition up.Cond) ([]*AppStat, error) {
    collection := upper.Collection(t.Table())
    var all []*AppStat

    res := collection.Find(condition)
    err := res.All(&all)
    if err != nil {
        return nil, err
    }

    return all, err
}

// Get gets one record from the database, by id, using upper
func (t *AppStat) Get(id int) (*AppStat, error) {
    var one AppStat
    collection := upper.Collection(t.Table())

    res := collection.Find(up.Cond{"id": id})
    err := res.One(&one)
    if err != nil {
        return nil, err
    }
    return &one, nil
}

// Update updates a record in the database, using upper
func (t *AppStat) Update(m AppStat) error {
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res := collection.Find(m.ID)
    err := res.Update(&m)
    if err != nil {
        return err
    }
    return nil
}

// Delete deletes a record from the database by id, using upper
func (t *AppStat) Delete(id int) error {
    collection := upper.Collection(t.Table())
    res := collection.Find(id)
    err := res.Delete()
    if err != nil {
        return err
    }
    return nil
}

// Insert inserts a model into the database, using upper
func (t *AppStat) Insert(m AppStat) (int, error) {
    if upper == nil {
        return 0, fmt.Errorf("database session is not initialized")
    }
    m.CreatedAt = time.Now()
    m.UpdatedAt = time.Now()
    fmt.Println(m)
    collection := upper.Collection(t.Table())
    res, err := collection.Insert(m)
    if err != nil {
        return 0, err
    }

    id := getInsertID(res.ID())

    return id, nil
}

// Builder is an example of using upper's sql builder
func (t *AppStat) Builder(id int) ([]*AppStat, error) {
    collection := upper.Collection(t.Table())

    var result []*AppStat

    err := collection.Session().
        SQL().
        SelectFrom(t.Table()).
        Where("id > ?", id).
        OrderBy("id").
        All(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

