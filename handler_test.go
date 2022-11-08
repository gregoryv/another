package another

import (
	"net/http"

	"github.com/gregoryv/ex"
)

func ExampleGetGopher() {
	GetGopher(ex.JsonOf(
		http.NewRequest("GET", "/another", nil),
	))
	// output:
	// {
	//     "color":"pink"
	// }
}
