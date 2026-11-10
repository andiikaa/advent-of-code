package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestCheckAttempt(t *testing.T) {
	prev := []attempt{
		{part: 1, verdict: "correct", answer: "42"},
		{part: 2, verdict: "too-high", answer: "1000"},
		{part: 2, verdict: "too-low", answer: "100"},
		{part: 2, verdict: "wrong", answer: "500"},
	}
	cases := []struct {
		part   int
		answer string
		ok     bool
	}{
		{1, "43", false},   // already solved
		{2, "1000", false}, // duplicate
		{2, "1001", false}, // above too-high bound
		{2, "99", false},   // below too-low bound
		{2, "500", false},  // known wrong
		{2, "501", true},
		{2, "abc", true},
	}
	for _, c := range cases {
		err := checkAttempt(prev, c.part, c.answer)
		if (err == nil) != c.ok {
			t.Errorf("part %d answer %s: err=%v, want ok=%v", c.part, c.answer, err, c.ok)
		}
	}
}

func TestToMarkdown(t *testing.T) {
	in := `<article class="day-desc"><h2>--- Day 1: Test ---</h2><p>For example:</p>
<pre><code>1,2
<em>3</em>,4
#&lt;^&lt;#
#.&lt;..&gt;#
</code></pre><p>Answer is <code><em>10</em></code> &gt; 5.</p></article>
<p>Your puzzle answer was <code>123</code>.</p><form method="post"><input name="answer"/></form>`
	got := toMarkdown(in)
	for _, want := range []string{"## --- Day 1: Test ---", "```\n1,2\n3,4\n#<^<#\n#.<..>#\n```", "`**10**` > 5.", "Your puzzle answer was `123`."} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	if regexp.MustCompile(`</?[a-z]`).MatchString(got) {
		t.Errorf("tags left in:\n%s", got)
	}
}

func TestWaitParsing(t *testing.T) {
	m := waitLeftRe.FindStringSubmatch("You gave an answer too recently; you have to wait after submitting an answer before trying again.  You have 4m 3s left to wait.")
	if m == nil || m[1] != "4" || m[2] != "3" {
		t.Errorf("got %v", m)
	}
	m = waitMinRe.FindStringSubmatch("please wait one minute before trying again")
	if m == nil || m[1] != "one" {
		t.Errorf("got %v", m)
	}
}
