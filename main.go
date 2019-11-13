package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	stripChars string
)

func init() {
	flag.StringVar(&stripChars, "strip", `{}`, "strip chars")
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	nodeMap := make(map[string]string)
	for _, line := range lines {
		// e.g.: Node0x7fa5a3d1d5c0 [shape=record,label="{external node}"];
		nodeDeclWithLabelPattern := regexp.MustCompile(`(?P<node>[0-9a-zA-Z]+).*\[.*label[ ]*=[ ]*"(?P<label>.*)".*;`)
		matches := nodeDeclWithLabelPattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			node := matches[1]
			label := matches[2]
			if label == "" {
				label = node
			}
			label = strings.Trim(label, stripChars)
			nodeMap[node] = label
		}
	}

	for _, line := range lines {
		{
			// e.g.: Node0x7fa5a3d1d5c0 [shape=record,label="{external node}"];
			nodeDeclWithLabelPattern := regexp.MustCompile(`(?P<node>[0-9a-zA-Z]+).*\[.*label[ ]*=[ ]*"(?P<label>.*)".*;`)
			matches := nodeDeclWithLabelPattern.FindStringSubmatch(line)
			if len(matches) == 3 {
				node := matches[1]
				label := strconv.Quote(nodeMap[node])
				if label == `""` {
					label = node
				}
				line = strings.Replace(line, node, label, 1)
			}
		}
		// e.g.: Node0x7fa5a3d1d5c0 [shape=record,label="{external node}"];
		{
			srcToDstPattern := regexp.MustCompile("(?P<src>[0-9a-zA-Z]+)[ ]*->[ ]*(?P<dst>[0-9a-zA-Z]+);")
			matches := srcToDstPattern.FindStringSubmatch(line)
			if len(matches) == 3 {
				src := matches[1]
				dst := matches[2]
				srcLabel := strconv.Quote(nodeMap[src])
				if srcLabel == `""` {
					srcLabel = src
				}
				dstLabel := strconv.Quote(nodeMap[dst])
				if dstLabel == `""` {
					dstLabel = dst
				}
				line = strings.Replace(line, src, srcLabel, 1)
				line = strings.Replace(line, dst, dstLabel, 1)
			}
		}
		fmt.Println(line)
	}
}
