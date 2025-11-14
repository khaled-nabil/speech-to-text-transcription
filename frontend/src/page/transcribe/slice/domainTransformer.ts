import type { Transcription, TranscriptionResponse } from 'types/transcription'

export const createNewTranscriptionFromBlob = (
	t: TranscriptionResponse,
	blob: Blob
): Transcription => {
	return {
		...t,
		audioURL: URL.createObjectURL(blob),
	}
}
