export type TranscriptionStatus = 'PENDING' | 'SUCCESS' | 'ERROR'

export interface Transcription {
	id: string
	uploadDate: string
	transcriptText: string
	status: TranscriptionStatus
}
