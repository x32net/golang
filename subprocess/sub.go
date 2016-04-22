package main

import (
	"flag"
	"fmt"
	"os"
)

const str = `<?xml version="1.0" encoding="UTF-8" ?>
<countries>
	<country>
		<code>CH</code>
		<name>Schweiz</name>
		<regions>
			<region>
				<code>10</code>
				<name>Mittelland</name>
				<subregions>
					<subregion>
						<code>11</code>
						<name>Blah blah</name>
					</subregion>
				</subregions>
			</region>
		</regions>
	</country>
</countries>`

var setting string

func init() {
	flag.StringVar(&setting, "setting", "value", "usage")
}

func main() {
	fmt.Fprintln(os.Stderr, "fmt.Fprintln(os.Stderr, setting: ", setting, ")") // os.Stderr
	if _, err := os.Stdout.Write([]byte(str)); err != nil { // os.Stdout
		fmt.Fprintln(os.Stderr, err) // os.Stderr
	}
}
