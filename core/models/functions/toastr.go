package functions

import "fmt"

type Toastr struct {
	Icon  string
	Title string
	Text  string
}

func Toast(t Toastr) string {
	return fmt.Sprintf("<script>$(document).ready(function() {toastr.%s('%s', '%s');});</script>", t.Icon, t.Text, t.Title)
}
