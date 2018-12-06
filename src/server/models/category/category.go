package category

import (
	"github.com/ycyxuehan/bingo"
)
const TABLE = "category"
type Category struct {
	bingo.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
}

//New
func New(name string, description string)*Category{
	c := Category{
		Name: name,
		Description:description,
		Enabled: true,
	}
	c.SetThis(&c)
	return &c
}

//Filter return query filter
func (c *Category)Filter()map[string]interface{}{
	filter := make(map[string]interface{})
	filter["name"] = c.Name
	return filter
}

func (c *Category)Columns()([]string, []interface{}){
	return []string{"name", "description", "enabled"}, []interface{}{&c.Name, &c.Description, &c.Enabled}
}

func (c *Category)Table()string{
	return TABLE
}