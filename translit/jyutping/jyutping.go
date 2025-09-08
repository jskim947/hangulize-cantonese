/*
Package jyutping implements the hangulize.Translit interface for Cantonese characters.
Cantonese characters are converted to Jyutping using pycantonese library via Python script.
*/
package jyutping

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/hangulize/hangulize"
)

// T is a hangulize.Translit for Jyutping.
var T hangulize.Translit = &jyutping{}

// ----------------------------------------------------------------------------

type jyutping struct{}

func (jyutping) Scheme() string {
	return "jyutping"
}

// JyutpingResult represents the result from pycantonese
type JyutpingResult struct {
	Character string `json:"character"`
	Jyutping  string `json:"jyutping"`
}

func (j *jyutping) Transliterate(word string) (string, error) {
	// Create Python script inline
	script := `
import pycantonese
import json
import sys
import io

# Set stdout to UTF-8
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

def convert_to_jyutping(text):
    try:
        result = pycantonese.characters_to_jyutping(text)
        return [{"character": char, "jyutping": jyut} for char, jyut in result]
    except Exception as e:
        return [{"character": text, "jyutping": "", "error": str(e)}]

if __name__ == "__main__":
    text = sys.argv[1] if len(sys.argv) > 1 else ""
    result = convert_to_jyutping(text)
    print(json.dumps(result, ensure_ascii=False))
`

	// Execute Python script inline - try python3 first, then python
	var cmd *exec.Cmd
	var stdout, stderr bytes.Buffer

	// Try python3 first
	cmd = exec.Command("python3", "-c", script, word)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	// If python3 fails, try python
	if err != nil {
		stdout.Reset()
		stderr.Reset()
		cmd = exec.Command("python", "-c", script, word)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
	}

	if err != nil {
		// If both Python executions fail, return the original word
		return word, nil
	}

	// Parse JSON result
	var results []JyutpingResult
	if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
		return word, nil
	}

	// Convert results to jyutping string
	var chunks []string
	for _, result := range results {
		if result.Jyutping != "" {
			// Add spaces between syllables for better processing
			jyutping := result.Jyutping
			// Insert spaces between tone numbers and next syllable
			jyutping = strings.ReplaceAll(jyutping, "1", "1 ")
			jyutping = strings.ReplaceAll(jyutping, "2", "2 ")
			jyutping = strings.ReplaceAll(jyutping, "3", "3 ")
			jyutping = strings.ReplaceAll(jyutping, "4", "4 ")
			jyutping = strings.ReplaceAll(jyutping, "5", "5 ")
			jyutping = strings.ReplaceAll(jyutping, "6", "6 ")
			// Remove trailing space
			jyutping = strings.TrimSpace(jyutping)
			chunks = append(chunks, jyutping)
		} else {
			chunks = append(chunks, result.Character)
		}
	}

	// U+200B: Zero Width Space
	return strings.Join(chunks, "\u200b"), nil
}
