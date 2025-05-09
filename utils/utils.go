package utils

import (
	"encoding/binary"

	"github.com/bndr/gotabulate"
	"github.com/pterm/pterm"
)

func IntToByte(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func TabularizeOutput(allTaskRows [][]string) string {
	// strings := []string{allTaskRows}
	headers := []string{"ID", "Task", "TimeStamp"}
	tabulate := gotabulate.Create(allTaskRows)
	tabulate.SetHeaders(headers)
	tabulate.SetMaxCellSize(55)
	tabulate.SetWrapStrings(true)
	tabulate.SetAlign("left")
	return tabulate.Render("simple")
}

func Ytdpretty(cyantasks [][]string) {
	tableData := pterm.TableData{
		{"ID", "TaskName", "timestamp"},
	}
	for _, taskRow := range cyantasks {
		tableData = append(tableData, taskRow)
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
