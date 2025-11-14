import { vi } from 'vitest'

const mockGetUserMedia = vi.fn(async () => {
	return new Promise<MediaStream>((resolve) => {
		const mockStream = {
			getTracks: vi.fn().mockReturnValue([]),
		} as unknown as MediaStream
		resolve(mockStream)
	})
})

Object.defineProperty(global.navigator, 'mediaDevices', {
	value: {
		getUserMedia: mockGetUserMedia,
	},
	writable: true,
	configurable: true,
})

class MockMediaRecorder implements MediaRecorder {
	ondataavailable: ((event: BlobEvent) => void) | null = null
	onerror: ((event: Event) => void) | null = null
	onpause: (() => void) | null = null
	onresume: (() => void) | null = null
	onstart: (() => void) | null = null
	onstop: (() => void) | null = null
	state: RecordingState = 'inactive'
	stream: MediaStream
	mimeType: string
	audioBitsPerSecond = 0
	videoBitsPerSecond = 0
	audioBitrateMode: BitrateMode = 'variable'

	constructor(stream: MediaStream, options?: MediaRecorderOptions) {
		this.stream = stream
		this.mimeType = options?.mimeType || 'audio/webm'
	}

	start = vi.fn(() => {
		this.state = 'recording'
		this.onstart?.()
	})

	stop = vi.fn(() => {
		this.state = 'inactive'
		this.onstop?.(new Event('stop'))
	})

	pause = vi.fn(() => {
		this.state = 'paused'
		this.onpause?.()
	})

	resume = vi.fn(() => {
		this.state = 'recording'
		this.onresume?.()
	})

	requestData = vi.fn()

	addEventListener = vi.fn()
	removeEventListener = vi.fn()
	dispatchEvent = vi.fn()

	static isTypeSupported = vi.fn(() => true)
}

global.MediaRecorder = MockMediaRecorder as any
