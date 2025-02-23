import type { components } from '~/clients/openapi';

export enum Module {
  LISTENING = 'LISTENING',
  READING = 'READING',
  WRITING = 'WRITING',
  SPEAKING = 'SPEAKING',
}

export type Exam = components['schemas']['handler.GetTestByIdResponse'];
export type QuestionSet = components['schemas']['core.QuestionSetEntry'];
export type Question = components['schemas']['core.QuestionEntry'];
export enum ResponseType {
  FREE_TEXT = 'FREE_TEXT',
}

export enum LocalStorageKey {
  VOLUME = 'volume',
}
