package utils

import (
	"bufio"
	"hash/fnv"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// GenerateRand はシード値を基にした *rand.Rand インスタンスを生成します。
func GenerateRand(seed string) *rand.Rand {
	if seed != "" {
		h := fnv.New64a()
		_, err := h.Write([]byte(seed))
		if err != nil {
			panic("failed to hash seed value: " + err.Error())
		}
		seedInt := int64(h.Sum64())
		return rand.New(rand.NewSource(seedInt))
	}
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Shuffle はスライスの要素をランダムに並び替えます。
func Shuffle(sliceLen int, swap func(i, j int), r *rand.Rand) {
	for i := sliceLen - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		swap(i, j)
	}
}

// Contains はスライスに特定の文字列が含まれているかをチェックします。
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// ClearScreen はコンソールの画面をクリアします。
func ClearScreen() {
	cmd := exec.Command("clear") // Unix/Linux 系の場合
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows の場合
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ReadLine は標準入力から一行読み込みます。
func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
