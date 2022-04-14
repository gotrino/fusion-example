package proc

import (
	"context"
	"github.com/gotrino/fusion-rt-wasmjs/pkg/web/i18n"
	"github.com/gotrino/fusion/spec/app"
	"github.com/gotrino/fusion/spec/form"
	"github.com/gotrino/fusion/spec/rest"
)

type Details struct {
	Route string `route:"/devices/:id/details"`
	ID    string `path:"id"`
}

func (a *Details) Compose(ctx context.Context) app.Activity {
	return app.Activity{
		Title: "Details",
		Fragments: []app.Fragment{
			form.Form{
				Title:       "Stammdaten",
				Description: "Bearbeiten Sie hier die Stammdaten des Scripts",
				ResourceID:  a.ID,
				Repository: rest.Repository[ProcInfo]{
					Path: "/api/v1/procs",
				},
				CanCancel: true,
				CanWrite:  true,
				Fields: []form.Field{
					form.Text[ProcInfo]{
						Label:       i18n.Text(ctx, "Name"),
						Placeholder: i18n.Text(ctx, "Name of the script"),
						Description: i18n.Text(ctx, "A short but memorable name"),
						ToModel: func(src string, state ProcInfo) (ProcInfo, error) {
							if src == "" {
								return state, app.ValidationError{Message: i18n.Text(ctx, "Name must not be empty")}
							}

							state.Name = src
							return state, nil
						},
						FromModel: func(src ProcInfo) string {
							return src.Name
						},
					},

					form.Text[ProcInfo]{
						Label:       i18n.Text(ctx, "Description"),
						Placeholder: i18n.Text(ctx, "..."),
						Description: i18n.Text(ctx, "A more verbose description of the script."),
						ToModel: func(src string, state ProcInfo) (ProcInfo, error) {
							state.Description = src
							return state, nil
						},
						FromModel: func(src ProcInfo) string {
							return src.Description
						},
					},
				},
			},

			form.Form{
				Title:       "Beispiel",
				Description: "Beispiel für Ein- und Ausgabe-Formate für das Script",
				ResourceID:  a.ID,
				Repository: rest.Repository[ProcInfo]{
					Path: "/api/v1/procs",
				},
				Fields: []form.Field{
					form.CodeEditor{},
				},
			},

			form.Form{
				Title:       "Script",
				Description: "Der definierte Rechenkernel",
				ResourceID:  a.ID,
				Repository: rest.Repository[ProcInfo]{
					Path: "/api/v1/procs",
				},
				Fields: []form.Field{
					form.CodeEditor{},
				},
			},
		},
	}
}
