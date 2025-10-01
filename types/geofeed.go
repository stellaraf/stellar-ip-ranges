package types

import "fmt"

type GeofeedEntry struct {
	Name        string
	Prefix      string
	CountryCode string
	RegionCode  string
	City        string
	PostalCode  string
}

type Geofeed []GeofeedEntry

func (g *GeofeedEntry) CSV() string {
	data := fmt.Sprintf("%s,%s,%s,%s,%s,",
		g.Prefix,
		g.CountryCode,
		g.RegionCode,
		g.City,
		g.PostalCode,
	)
	//comment := fmt.Sprintf("# %s", g.Name)
	return fmt.Sprintf(`%s`, data)
}

func (g Geofeed) CSV(title string) string {
	data := fmt.Sprintf("# %s\n", title)
	for _, e := range g {
		data += e.CSV()
		data += "\n"
	}
	return data
}
