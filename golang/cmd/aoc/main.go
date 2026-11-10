// Command aoc is a small Advent of Code helper: show star status, fetch a
// puzzle (description, example, input) and submit answers with guards against
// the rate limit.
//
// Run from the golang/ directory:
//
//	go run ./cmd/aoc status 2022
//	go run ./cmd/aoc fetch 2022 19
//	go run ./cmd/aoc submit 2022 19 1 1234
//
// The session cookie is read from $AOC_SESSION or ~/.config/aoc/session.
package main

import (
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const baseURL = "https://adventofcode.com"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 1 && args[0] == "login" {
		return login()
	}
	if len(args) < 2 {
		return errors.New(usage)
	}
	year, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid year %q", args[1])
	}
	switch args[0] {
	case "status":
		return status(year)
	case "fetch":
		if len(args) != 3 {
			return errors.New(usage)
		}
		day, err := parseDay(args[2])
		if err != nil {
			return err
		}
		return fetch(year, day)
	case "submit":
		if len(args) != 5 {
			return errors.New(usage)
		}
		day, err := parseDay(args[2])
		if err != nil {
			return err
		}
		part, err := strconv.Atoi(args[3])
		if err != nil || (part != 1 && part != 2) {
			return fmt.Errorf("invalid part %q", args[3])
		}
		return submit(year, day, part, strings.TrimSpace(args[4]))
	}
	return errors.New(usage)
}

const usage = `usage:
  aoc login                               copy the session cookie from Safari to ~/.config/aoc/session
  aoc status <year>                       star status per day
  aoc fetch  <year> <day>                 scaffold day dir, write puzzle.md, example.txt, input.txt
  aoc submit <year> <day> <part> <answer> submit an answer (guarded)`

func parseDay(s string) (int, error) {
	d, err := strconv.Atoi(s)
	if err != nil || d < 1 || d > 25 {
		return 0, fmt.Errorf("invalid day %q", s)
	}
	return d, nil
}

func dayDir(year, day int) string {
	return filepath.Join(strconv.Itoa(year), fmt.Sprintf("day%02d", day))
}

// ---------------------------------------------------------------- http

func session() (string, error) {
	if s := strings.TrimSpace(os.Getenv("AOC_SESSION")); s != "" {
		return s, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(filepath.Join(home, ".config", "aoc", "session"))
	if err != nil || strings.TrimSpace(string(b)) == "" {
		return "", errors.New("no session: set AOC_SESSION or write the cookie value to ~/.config/aoc/session")
	}
	return strings.TrimSpace(string(b)), nil
}

func request(method, path string, form url.Values) (string, error) {
	sess, err := session()
	if err != nil {
		return "", err
	}
	var body io.Reader
	if form != nil {
		body = strings.NewReader(form.Encode())
	}
	req, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sess})
	ua := os.Getenv("AOC_USER_AGENT")
	if ua == "" {
		ua = "aoc-helper (go stdlib)"
	}
	req.Header.Set("User-Agent", ua)
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
		// A redirect means the session is invalid (AoC redirects to the login page).
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s %s: %s (session expired or puzzle not unlocked?)", method, path, resp.Status)
	}
	return string(b), nil
}

// ---------------------------------------------------------------- status

var calendarRe = regexp.MustCompile(`aria-label="Day (\d+)(?:, (one star|two stars))?"`)

func status(year int) error {
	page, err := request("GET", fmt.Sprintf("/%d", year), nil)
	if err != nil {
		return err
	}
	stars := map[int]int{}
	for _, m := range calendarRe.FindAllStringSubmatch(page, -1) {
		d, _ := strconv.Atoi(m[1])
		switch m[2] {
		case "one star":
			stars[d] = 1
		case "two stars":
			stars[d] = 2
		default:
			stars[d] = 0
		}
	}
	if len(stars) == 0 {
		return errors.New("could not parse calendar (session invalid?)")
	}
	total := 0
	var todo []string
	for d := 1; d <= 25; d++ {
		s, ok := stars[d]
		if !ok {
			continue
		}
		total += s
		mark := map[int]string{0: "  -  none", 1: " *   part 2 missing", 2: " **  done"}[s]
		fmt.Printf("day %2d %s\n", d, mark)
		if s < 2 {
			todo = append(todo, strconv.Itoa(d))
		}
	}
	fmt.Printf("\n%d stars; todo: %s\n", total, strings.Join(todo, " "))
	return nil
}

// ---------------------------------------------------------------- fetch

var (
	mainRe    = regexp.MustCompile(`(?s)<main>(.*)</main>`)
	formRe    = regexp.MustCompile(`(?s)<form.*?</form>`)
	scriptRe  = regexp.MustCompile(`(?s)<script.*?</script>`)
	preRe     = regexp.MustCompile(`(?s)<pre><code>(.*?)</code></pre>`)
	h2Re      = regexp.MustCompile(`(?s)<h2[^>]*>(.*?)</h2>`)
	emRe      = regexp.MustCompile(`(?s)<em[^>]*>(.*?)</em>`)
	codeRe    = regexp.MustCompile(`(?s)<code>(.*?)</code>`)
	liRe      = regexp.MustCompile(`<li>`)
	tagRe     = regexp.MustCompile(`(?s)<[^>]+>`)
	blankRe   = regexp.MustCompile(`\n{3,}`)
	articleRe = regexp.MustCompile(`(?s)<article[^>]*>(.*?)</article>`)
)

func stripTags(s string) string { return html.UnescapeString(tagRe.ReplaceAllString(s, "")) }

// toMarkdown converts the AoC page fragment into readable markdown.
func toMarkdown(s string) string {
	s = scriptRe.ReplaceAllString(s, "")
	s = formRe.ReplaceAllString(s, "")
	s = preRe.ReplaceAllStringFunc(s, func(m string) string {
		inner := preRe.FindStringSubmatch(m)[1]
		// keep the text escaped: the final stripTags would otherwise eat "<...>" in the code
		return "\n```\n" + html.EscapeString(strings.TrimRight(stripTags(inner), "\n")) + "\n```\n"
	})
	s = h2Re.ReplaceAllString(s, "\n## $1\n")
	s = emRe.ReplaceAllString(s, "**$1**")
	s = codeRe.ReplaceAllString(s, "`$1`")
	s = liRe.ReplaceAllString(s, "- ")
	s = strings.NewReplacer("</p>", "\n\n", "</li>", "\n", "</ul>", "\n").Replace(s)
	s = stripTags(s)
	s = blankRe.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s) + "\n"
}

func fetch(year, day int) error {
	dir := dayDir(year, day)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	page, err := request("GET", fmt.Sprintf("/%d/day/%d", year, day), nil)
	if err != nil {
		return err
	}
	m := mainRe.FindStringSubmatch(page)
	if m == nil {
		return errors.New("could not find puzzle content")
	}
	// The description is always refreshed: part 2 only appears after part 1 is solved.
	if err := os.WriteFile(filepath.Join(dir, "puzzle.md"), []byte(toMarkdown(m[1])), 0o644); err != nil {
		return err
	}
	fmt.Println("wrote", filepath.Join(dir, "puzzle.md"))

	// Example: first code block of the description. Never overwritten; verify it by reading puzzle.md.
	desc := m[1]
	if i := strings.Index(desc, "For example"); i >= 0 {
		desc = desc[i:] // the real example usually follows "For example"
	}
	if pre := preRe.FindStringSubmatch(desc); pre != nil {
		if wrote, err := writeIfMissing(filepath.Join(dir, "example.txt"), strings.TrimRight(stripTags(pre[1]), "\n")); err != nil {
			return err
		} else if wrote {
			fmt.Println("wrote", filepath.Join(dir, "example.txt"), "(first code block - check it is the right one)")
		}
	}

	// Input never changes, so it is downloaded only once.
	inputPath := filepath.Join(dir, "input.txt")
	if _, err := os.Stat(inputPath); errors.Is(err, os.ErrNotExist) {
		in, err := request("GET", fmt.Sprintf("/%d/day/%d/input", year, day), nil)
		if err != nil {
			return err
		}
		if err := os.WriteFile(inputPath, []byte(strings.TrimRight(in, "\n")), 0o644); err != nil {
			return err
		}
		fmt.Println("wrote", inputPath)
	}

	title := fmt.Sprintf("Day %d", day)
	if t := h2Re.FindStringSubmatch(m[1]); t != nil {
		title = strings.Trim(stripTags(t[1]), "- ")
	}
	for name, tmpl := range map[string]*template.Template{"main.go": mainTmpl, "main_test.go": testTmpl, "NOTES.md": notesTmpl} {
		var sb strings.Builder
		if err := tmpl.Execute(&sb, map[string]any{"Year": year, "Day": day, "Title": title}); err != nil {
			return err
		}
		if wrote, err := writeIfMissing(filepath.Join(dir, name), sb.String()); err != nil {
			return err
		} else if wrote {
			fmt.Println("wrote", filepath.Join(dir, name))
		}
	}
	return nil
}

func writeIfMissing(path, content string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return false, nil
	}
	return true, os.WriteFile(path, []byte(content), 0o644)
}

var mainTmpl = template.Must(template.New("main").Parse(`package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func part1(in string) int {
	lines := strings.Split(in, "\n")
	_ = lines
	return 0
}

func part2(in string) int {
	return 0
}
`))

var notesTmpl = template.Must(template.New("notes").Parse(`# {{.Year}} · {{.Title}}

https://adventofcode.com/{{.Year}}/day/{{.Day}}

## Part 1

**Approach:**

**Struggles:**

**Result:** runtime ?, peak memory ?, CPU ?

## Part 2

**Approach:**

**Struggles:**

**Result:** runtime ?, peak memory ?, CPU ?
`))

var testTmpl = template.Must(template.New("test").Parse(`package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		expected := 0 // TODO from puzzle.md
		if got := part1(example); got != expected {
			t.Errorf("Got %d, expected %d", got, expected)
		}
	})

	t.Run("part2", func(t *testing.T) {
		expected := 0 // TODO from puzzle.md
		if got := part2(example); got != expected {
			t.Errorf("Got %d, expected %d", got, expected)
		}
	})
}

// TestInput locks in the accepted answers (fill in after submitting).
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		t.Skip("not solved yet")
		expected := 0
		if got := part1(input); got != expected {
			t.Errorf("Got %d, expected %d", got, expected)
		}
	})

	t.Run("part2", func(t *testing.T) {
		t.Skip("not solved yet")
		expected := 0
		if got := part2(input); got != expected {
			t.Errorf("Got %d, expected %d", got, expected)
		}
	})
}
`))

// ---------------------------------------------------------------- submit

// attempt is one submission, stored in <day>/submissions.log as
// "<RFC3339 time>\t<part>\t<verdict>\t<answer>".
type attempt struct {
	part    int
	verdict string // correct, wrong, too-high, too-low
	answer  string
}

func logPath(year, day int) string { return filepath.Join(dayDir(year, day), "submissions.log") }
func cooldownPath() string         { return filepath.Join(".aoc", "cooldown") }

func loadAttempts(year, day int) ([]attempt, error) {
	b, err := os.ReadFile(logPath(year, day))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var out []attempt
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		f := strings.Split(line, "\t")
		if len(f) != 4 {
			continue
		}
		p, _ := strconv.Atoi(f[1])
		out = append(out, attempt{part: p, verdict: f[2], answer: f[3]})
	}
	return out, nil
}

// checkAttempt refuses answers that cannot be right given earlier feedback.
func checkAttempt(prev []attempt, part int, answer string) error {
	n, numErr := strconv.ParseInt(answer, 10, 64)
	for _, a := range prev {
		if a.part != part {
			continue
		}
		if a.verdict == "correct" {
			return fmt.Errorf("part %d already solved with %s", part, a.answer)
		}
		if a.answer == answer {
			return fmt.Errorf("%s was already submitted for part %d: %s", answer, part, a.verdict)
		}
		if numErr != nil {
			continue
		}
		bound, err := strconv.ParseInt(a.answer, 10, 64)
		if err != nil {
			continue
		}
		if a.verdict == "too-high" && n >= bound {
			return fmt.Errorf("%s is not below %s, which was too high", answer, a.answer)
		}
		if a.verdict == "too-low" && n <= bound {
			return fmt.Errorf("%s is not above %s, which was too low", answer, a.answer)
		}
	}
	return nil
}

var (
	waitLeftRe = regexp.MustCompile(`You have (?:(\d+)m )?(\d+)s left to wait`)
	waitMinRe  = regexp.MustCompile(`(?i)please wait (one|\d+) minutes?`)
)

func submit(year, day, part int, answer string) error {
	// Day 25 part 2 has no puzzle: its "Start the Blender" button posts answer 0.
	if answer == "" || (answer == "0" && !(day == 25 && part == 2)) {
		return fmt.Errorf("refusing to submit %q", answer)
	}
	prev, err := loadAttempts(year, day)
	if err != nil {
		return err
	}
	if err := checkAttempt(prev, part, answer); err != nil {
		return fmt.Errorf("not submitted: %w", err)
	}
	if b, err := os.ReadFile(cooldownPath()); err == nil {
		if until, err := time.Parse(time.RFC3339, strings.TrimSpace(string(b))); err == nil && time.Now().Before(until) {
			return fmt.Errorf("not submitted: cooldown for another %s", time.Until(until).Round(time.Second))
		}
	}

	page, err := request("POST", fmt.Sprintf("/%d/day/%d/answer", year, day),
		url.Values{"level": {strconv.Itoa(part)}, "answer": {answer}})
	if err != nil {
		return err
	}
	text := page
	if m := articleRe.FindStringSubmatch(page); m != nil {
		text = toMarkdown(m[1])
	}
	fmt.Println(strings.TrimSpace(text))

	verdict := ""
	switch {
	case strings.Contains(text, "That's the right answer"):
		verdict = "correct"
	case strings.Contains(text, "That's not the right answer"):
		verdict = "wrong"
		if strings.Contains(text, "too high") {
			verdict = "too-high"
		} else if strings.Contains(text, "too low") {
			verdict = "too-low"
		}
	}

	var wait time.Duration
	if m := waitLeftRe.FindStringSubmatch(text); m != nil {
		min, _ := strconv.Atoi(m[1])
		sec, _ := strconv.Atoi(m[2])
		wait = time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
	} else if m := waitMinRe.FindStringSubmatch(text); m != nil {
		min := 1
		if m[1] != "one" {
			min, _ = strconv.Atoi(m[1])
		}
		wait = time.Duration(min) * time.Minute
	}
	if wait > 0 {
		_ = os.MkdirAll(filepath.Dir(cooldownPath()), 0o755)
		_ = os.WriteFile(cooldownPath(), []byte(time.Now().Add(wait).Format(time.RFC3339)), 0o644)
	}

	if verdict == "" {
		fmt.Println("\n(no verdict recorded - rate limited, wrong level or already solved)")
		return nil
	}
	f, err := os.OpenFile(logPath(year, day), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%s\t%d\t%s\t%s\n", time.Now().Format(time.RFC3339), part, verdict, answer)
	fmt.Println("\nverdict:", verdict)
	return err
}
