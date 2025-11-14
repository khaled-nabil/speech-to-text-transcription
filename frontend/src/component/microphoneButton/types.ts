export type RecordingState = 'INACTIVE' | 'RECORDING' | 'PAUSED'

export interface AudioRecorderCallbacks {
	onDataChunk?(chunk: Blob): void
	onStop?(finalBlob: Blob): void
	onError?(err: Error): void
}

export type AudioRecorder = {
	start(): void
	pause(): void
	resume(): void
	stop(): void
	getState(): RecordingState
	getRecorder(): MediaRecorder
}
