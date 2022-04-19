package main

import (
	"regexp"
	"strings"
)

type LbcDetail struct {
	EncodedID    string `json:"encodedId"`
	ContactName  string `json:"contactName"`
	Description  string `json:"description"`
	Goals        string `json:"goals"`
	Skills       string `json:"skills"`
	Deliverables string `json:"deliverables"`
}

func (detail *LbcDetail) missionTextToDetailObjectConverter(missionText string) error {
	r := regexp.MustCompile(`Resp.\sOp.\s:\s(.*?)\s\d+`)
	matchStrings := r.FindStringSubmatch(missionText)
	if len(matchStrings) > 0 {
		detail.ContactName = matchStrings[1]
	}

	r = regexp.MustCompile(`mission\n-{22}\n\n([\S\s]*?)\n\n-{33}\nObjectifs`)
	matchStrings = r.FindStringSubmatch(missionText)
	if len(matchStrings) > 0 {
		detail.Description = strings.ReplaceAll(matchStrings[1], "\n", " ")
	}

	r = regexp.MustCompile(`Compétences\sinformatiques\n-{25}\n\n([\S\s]*?)\n\n-{26}\nCompétences\sfonctionnelles`)
	matchStrings = r.FindStringSubmatch(missionText)
	if len(matchStrings) > 0 {
		detail.Skills = strings.ReplaceAll(matchStrings[1], "\n", " ")
	} else {
		r = regexp.MustCompile(`Compétences\sinformatiques\n-{25}\n\n([\S\s]*?)\n\n-{24}\nAdresse`)
		matchStrings = r.FindStringSubmatch(missionText)
		if len(matchStrings) > 0 {
			detail.Skills = strings.ReplaceAll(matchStrings[1], "\n", " ")
		}
	}

	r = regexp.MustCompile(`tâches\n-{33}\n\n([\S\s]*?)\n\n-{9}\nLivrables`)
	matchStrings = r.FindStringSubmatch(missionText)
	if len(matchStrings) > 0 {
		detail.Goals = strings.ReplaceAll(matchStrings[1], "\n", " ")
	}

	r = regexp.MustCompile(`Livrables\n-{9}\n\n([\S\s]*?)\n\n-{25}\nCompétences`)
	matchStrings = r.FindStringSubmatch(missionText)
	if len(matchStrings) > 0 {
		detail.Deliverables = strings.ReplaceAll(matchStrings[1], "\n", " ")
	}

	return nil
}
