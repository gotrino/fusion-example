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
				Resource: rest.Resource[ProcInfo]{
					Path: "/api/v1/procs/meta/" + a.ID,
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

					form.Integer[ProcInfo]{
						Text:     "Degree",
						Disabled: true,
						FromModel: func(src ProcInfo) int64 {
							return 100
						},
					},
					form.Select[ProcInfo]{
						Text:        "Select something",
						Hint:        "a cool selector",
						MultiSelect: true,
						ToModel: func(src []form.Item, dst *ProcInfo) error {
							return nil
						},
						FromModel: func(src ProcInfo) []form.Item {
							return nil
						},
					},
				},
			},
		},
	}
}
