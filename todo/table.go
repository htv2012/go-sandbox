package main

import (
	"errors"
	"fmt"
	"strings"
)

type Row []string

type Table struct {
	builder     strings.Builder
	header      []string
	rows        []Row
	columnCount int
	widths      []int
}

func NewTable() *Table {
	return &Table{
		header: []string{},
		rows:   []Row{},
	}
}

func (table *Table) AddHeader(header []string) {
	table.header = header
	table.columnCount = len(header)
	for _, value := range header {
		table.widths = append(table.widths, len(value))
	}
}

func (table *Table) AddRow(row Row) error {
	if len(row) != table.columnCount {
		return errors.New("Number of columns don't match")
	}

	table.rows = append(table.rows, row)
	for i, value := range row {
		table.widths[i] = max(table.widths[i], len(value))
	}

	return nil
}

func (table *Table) printTop() {
	fmt.Print("┌")
	last := len(table.widths) - 1
	for i, w := range table.widths {
		fmt.Print(strings.Repeat("─", w+2))
		if i == last {
			fmt.Print("")
		} else {
			fmt.Print("┬")
		}
	}
	fmt.Print("┐")
	fmt.Println()
}

func (table *Table) printBottom() {
	fmt.Print("└")
	last := len(table.widths) - 1
	for i, w := range table.widths {
		fmt.Print(strings.Repeat("─", w+2))
		if i == last {
			fmt.Print("┘")
		} else {
			fmt.Print("┴")
		}
	}
	fmt.Println()
}

func (table *Table) printHeader() {
	fmt.Print("│")
	for i, value := range table.header {
		fmt.Printf(" %-*s ", table.widths[i], value)
		fmt.Print("│")
	}
	fmt.Println()
}

func (table *Table) printDivider() {
	fmt.Print("├")
	last := len(table.widths) - 1
	for i, w := range table.widths {
		fmt.Print(strings.Repeat("─", w+2))
		if i == last {
			fmt.Print("┤")
		} else {
			fmt.Print("┼")
		}
	}
	fmt.Println()
}

func (table *Table) printRow(row []string) {
	fmt.Print("│")
	for i, value := range row {
		fmt.Printf(" %-*s ", table.widths[i], value)
		fmt.Print("│")
	}
	fmt.Println()
}

func (table *Table) Print() {
	table.printTop()
	table.printHeader()
	table.printDivider()
	for _, row := range table.rows {
		table.printRow(row)
	}
	table.printBottom()
}

func tableDemo() {
	table := NewTable()
	table.AddHeader([]string{"#", "Task", "Done"})
	table.AddRow([]string{"1", "Wash Fido", "x"})
	table.AddRow([]string{"2", "Grocery Shopping", "v"})
	table.Print()
}
