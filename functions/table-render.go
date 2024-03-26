package functions

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

var (
	results *tablewriter.Table
)

// Initialise the table containing the process statuses
func InitialiseTable() {
	results = tablewriter.NewWriter(os.Stdout)
	results.SetHeader([]string{"Item", "Status"})
}

// Add rows to said table
func AppendSlice(slice [][]string) {
	for _, v := range slice {
		results.Append(v)
	}
}

// Render the table
func RenderTable() {
	results.Render()
}
