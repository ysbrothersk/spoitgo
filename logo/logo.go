package logo

import (
	"fmt"
)

const logo = " ____              _ _    ____\n" +
	"/ ___| _ __   ___ (_) |_ / ___| ___\n" +
	"\\___ \\| '_ \\ / _ \\| | __| |  _ / _ \\\n" +
	" ___) | |_) | (_) | | |_| |_| | (_) |\n" +
	"|____/| .__/ \\___/|_|\\__|\\____|\\___/\n" +
	"      |_|"

// Print spoitgoロゴを出力します
func Print() {
	fmt.Println(logo)
}
