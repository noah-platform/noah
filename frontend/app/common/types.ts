import type { components } from '~/clients/openapi';

export enum Module {
  LISTENING = 'LISTENING',
  READING = 'READING',
  WRITING = 'WRITING',
  SPEAKING = 'SPEAKING',
}

export type Session = components['schemas']['handler.BeginTestSessionResponse'];
export type Exam = components['schemas']['handler.BeginTestSessionResponse']['test'];
export type ExamCover = components['schemas']['handler.GetTestCoverResponse'];
export type QuestionSet = components['schemas']['core.QuestionSetEntry'];
export type Question = components['schemas']['core.QuestionEntry'];
export enum ResponseType {
  FREE_TEXT = 'FREE_TEXT',
  NOTE_COMPLETION = 'NOTE_COMPLETION',
  MULTIPLE_CHOICE = 'MULTIPLE_CHOICE',
  MULTIPLE_SELECT = 'MULTIPLE_SELECT',
}
export enum TestSessionStatus {
  NOT_STARTED = 'NOT_STARTED',
  IN_PROGRESS = 'IN_PROGRESS',
  COMPLETED = 'COMPLETED',
}

export enum LocalStorageKey {
  VOLUME = 'volume',
  LAST_SELECTED_MODULE = 'last_selected_module',
}
