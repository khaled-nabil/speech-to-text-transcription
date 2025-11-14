import type {
	AudioRecorder,
	AudioRecorderCallbacks,
	RecordingState,
} from './types'

export const createAudioRecorder = async (
	callbacks: AudioRecorderCallbacks,
	mimeType = 'audio/webm'
): Promise<AudioRecorder> => {
	const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
	const chunks: Blob[] = []
	const recorder = new MediaRecorder(stream, { mimeType })

	recorder.ondataavailable = (e) => {
		if (e.data && e.data.size > 0) {
			chunks.push(e.data)
			callbacks.onDataChunk?.(e.data)
		}
	}

	recorder.onstop = () => {
		const finalBlob = new Blob(chunks, { type: mimeType })
		stream.getTracks().forEach((t) => t.stop())
		callbacks.onStop?.(finalBlob)
	}

	recorder.onerror = (e) => {
		callbacks.onError?.(e.error ?? new Error('MediaRecorder error'))
	}

	return {
		start() {
			if (recorder.state === 'inactive') recorder.start()
		},
		pause() {
			if (recorder.state === 'recording') recorder.pause()
		},
		resume() {
			if (recorder.state === 'paused') recorder.resume()
		},
		stop() {
			if (recorder.state !== 'inactive') recorder.stop()
		},
		getState(): RecordingState {
			switch (recorder.state) {
				case 'recording':
					return 'RECORDING'
				case 'paused':
					return 'PAUSED'
				default:
					return 'INACTIVE'
			}
		},
		getRecorder(): MediaRecorder {
			return recorder
		},
	}
}
