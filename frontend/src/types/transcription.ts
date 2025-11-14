export type TranscriptionStatus = 'PENDING' | 'SUCCESS' | 'ERROR'

export interface TranscriptionResponse {
	id: string
	uploadDate: string
	transcriptText?: string
	status: TranscriptionStatus
}

export interface Transcription extends TranscriptionResponse {
	audioURL: string
}

export interface TranscriptionReady extends Required<Transcription> {
	status: 'SUCCESS'
}
