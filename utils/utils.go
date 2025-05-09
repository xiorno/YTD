package utils

import (
	"encoding/binary"

	"github.com/pterm/pterm"
)

func IntToByte(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// ignore the name :(
func Ytdpretty(cyantasks [][]string) {
	tableData := pterm.TableData{
		{"ID", "TaskName", "timestamp"},
	}
	for _, taskRow := range cyantasks {
		tableData = append(tableData, taskRow)
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
