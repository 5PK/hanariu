// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package hanariu

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func IconMoon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M5 5v2H3v2h2v2h2V9h2V7H7V5Zm15.281 2.938L18.625 8C13.281 8.191 9 12.578 9 17.969c0 5.511 4.488 10 10 10 5.39 0 9.777-4.282 9.969-9.625l.062-1.625-1.468.687A5.94 5.94 0 0 1 25 17.97c-3.324 0-6-2.676-6-6 0-.914.223-1.75.594-2.531Zm-2.906 2.375c-.125.554-.375 1.062-.375 1.656 0 4.406 3.594 8 8 8 .605 0 1.121-.246 1.688-.375-.762 3.625-3.829 6.375-7.688 6.375-4.43 0-8-3.57-8-8 0-3.852 2.758-6.887 6.375-7.657Z\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
