package repository

import "github.com/noah-platform/noah/question-bank/server/core"

func mapSections(docs []SectionDocument) []core.SectionEntry {
	sections := make([]core.SectionEntry, len(docs))
	for i, doc := range docs {
		sections[i] = core.SectionEntry{
			ID:            doc.ID.Hex(),
			Title:         doc.Title,
			Instruction:   doc.Instruction,
			AudioUrl:      doc.AudioUrl,
			QuestionCount: doc.QuestionCount,
			QuestionSet:   mapQuestionSets(doc.QuestionSet),
		}
	}
	return sections
}

func mapQuestionSets(docs []QuestionSetDocument) []core.QuestionSetEntry {
	sets := make([]core.QuestionSetEntry, len(docs))
	for i, doc := range docs {
		sets[i] = core.QuestionSetEntry{
			Instruction:  doc.Instruction,
			ResponseType: doc.ResponseType,
			Questions:    mapQuestions(doc.Questions),
		}
	}
	return sets
}

func mapQuestions(docs []QuestionDocument) []core.QuestionEntry {
	questions := make([]core.QuestionEntry, len(docs))
	for i, doc := range docs {
		questions[i] = core.QuestionEntry{
			Body:      doc.Body,
			ImageUrls: doc.ImageUrls,
		}
	}
	return questions
}
