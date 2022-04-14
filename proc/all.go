package proc

import (
	"context"
	"github.com/gotrino/fusion/spec/app"
	"github.com/gotrino/fusion/spec/rest"
	"github.com/gotrino/fusion/spec/svg"
	"github.com/gotrino/fusion/spec/table"
)

var ProcInfoRepo = rest.Repository[ProcInfo]{
	Path: "/api/v1/procs",
}

type Overview struct {
	Route string `route:"/procs"` // gets injected, due to app.Route, matches the route-tag
	Limit int    `query:"limit"`  // parsed from the routing system, due to query-tag
}

func (a *Overview) Compose(ctx context.Context) app.Activity {
	return app.Activity{
		Title: "Stored Procedures",
		Launcher: app.Icon{
			Icon:  svg.OutlineCalculator,
			Title: "Stored Procedures",
			Hint:  "Zeigt alle procs",
			Link:  "/procs",
		},
		Visible: true,
		Fragments: []app.Fragment{
			table.DataTable[ProcInfo]{
				Repository: ProcInfoRepo,
				Deletable:  true,
				Columns:    []table.Column{{Name: "Name"}, {Name: "Beschreibung"}},
				OnRender: func(ctx context.Context, t ProcInfo, col int) table.Cell {
					switch col {
					case 0:
						return table.NewSVG(svg.OutlineCalculator, t.Name, "Script")
					case 1:
						return table.NewText(t.Description)

					default:
						return table.NewText("")
					}
				},
				OnClick: func(ctx context.Context, t ProcInfo) {
					app.Navigate(ctx, &Details{ID: t.ID})
				},
			},
		},
	}
}
