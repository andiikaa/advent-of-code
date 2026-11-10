package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// safariSession reads the adventofcode.com session cookie from Safari's
// Cookies.binarycookies. macOS only lets processes with Full Disk Access read it.
func safariSession() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, "Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies")
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return "", fmt.Errorf("%w\nmacOS blocks this file: give your terminal app Full Disk Access (System Settings > Privacy & Security), or copy the cookie by hand", err)
		}
		return "", err
	}
	return findCookie(data, "adventofcode.com", "session")
}

// findCookie parses the binarycookies format:
// "cook", page count (BE), page sizes (BE), then pages. Each page: 0x00000100,
// cookie count (LE), cookie offsets (LE). Each cookie record holds LE offsets
// (relative to the record) to NUL-terminated domain, name, path and value.
func findCookie(data []byte, domain, name string) (string, error) {
	if len(data) < 8 || string(data[:4]) != "cook" {
		return "", errors.New("not a binarycookies file")
	}
	pages := int(binary.BigEndian.Uint32(data[4:8]))
	if 8+4*pages > len(data) {
		return "", errors.New("truncated binarycookies file")
	}
	pos := 8 + 4*pages
	for i := 0; i < pages; i++ {
		size := int(binary.BigEndian.Uint32(data[8+4*i:]))
		if pos+size > len(data) {
			return "", errors.New("truncated page")
		}
		page := data[pos : pos+size]
		pos += size
		if len(page) < 8 {
			continue
		}
		n := int(binary.LittleEndian.Uint32(page[4:8]))
		for c := 0; c < n && 8+4*c+4 <= len(page); c++ {
			off := int(binary.LittleEndian.Uint32(page[8+4*c:]))
			if off+32 > len(page) {
				continue
			}
			rec := page[off:]
			str := func(at int) string {
				o := int(binary.LittleEndian.Uint32(rec[at:]))
				if o >= len(rec) {
					return ""
				}
				s := rec[o:]
				if end := bytes.IndexByte(s, 0); end >= 0 {
					s = s[:end]
				}
				return string(s)
			}
			d, nm, val := str(16), str(20), str(28)
			if nm == name && strings.TrimPrefix(d, ".") == domain {
				return val, nil
			}
		}
	}
	return "", fmt.Errorf("no %s cookie for %s in Safari (logged in?)", name, domain)
}

// login copies the Safari cookie to ~/.config/aoc/session.
func login() error {
	sess, err := safariSession()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, ".config", "aoc")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	path := filepath.Join(dir, "session")
	if err := os.WriteFile(path, []byte(sess+"\n"), 0o600); err != nil {
		return err
	}
	fmt.Printf("wrote session cookie (%d chars) to %s\n", len(sess), path)
	return nil
}
