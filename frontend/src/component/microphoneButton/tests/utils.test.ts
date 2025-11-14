import { describe, it, beforeEach, afterEach, vi, expect } from 'vitest'
import { createAudioRecorder } from '../utils'
import type { AudioRecorderCallbacks } from '../types'

describe('createAudioRecorder', () => {
	let getUserMediaMock: () => Promise<MediaStream>
	let mockStream: MediaStream
	let mockTrack: MediaStreamTrack

	beforeEach(() => {
		mockTrack = {
			stop: vi.fn(),
		} as unknown as MediaStreamTrack

		mockStream = {
			getTracks: vi.fn().mockReturnValue([mockTrack]),
		} as unknown as MediaStream

		getUserMediaMock = vi.fn().mockResolvedValue(mockStream)
		vi.spyOn(navigator.mediaDevices, 'getUserMedia').mockImplementation(
			getUserMediaMock
		)
	})

	afterEach(() => {
		vi.restoreAllMocks()
	})

	it('should request audio permission with correct constraints', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		await createAudioRecorder(callbacks)

		expect(getUserMediaMock).toHaveBeenCalledOnce()
		expect(getUserMediaMock).toHaveBeenCalledWith({ audio: true })
	})

	it('should create MediaRecorder with default mimeType', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		expect(recorder.getRecorder()).toBeInstanceOf(MediaRecorder)
	})

	it('should create MediaRecorder with custom mimeType', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const customMimeType = 'audio/mp4'
		const recorder = await createAudioRecorder(callbacks, customMimeType)

		expect(recorder.getRecorder().mimeType).toBe(customMimeType)
	})

	it('should start recording when start is called', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()

		expect(recorder.getState()).toBe('RECORDING')
	})

	it('should not start recording if already recording', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)
		const mediaRecorder = recorder.getRecorder()
		const startSpy = vi.spyOn(mediaRecorder, 'start')

		recorder.start()

		recorder.start()

		expect(startSpy).toHaveBeenCalledOnce()
	})

	it('should pause recording when pause is called during recording', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()
		recorder.pause()

		expect(recorder.getState()).toBe('PAUSED')
	})

	it('should not pause if not recording', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.pause()

		expect(recorder.getState()).toBe('INACTIVE')
	})

	it('should resume recording when resume is called during pause', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()
		recorder.pause()
		recorder.resume()

		expect(recorder.getState()).toBe('RECORDING')
	})

	it('should not resume if not paused', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()
		const mediaRecorder = recorder.getRecorder()
		const resumeSpy = vi.spyOn(mediaRecorder, 'resume')

		recorder.resume()

		expect(resumeSpy).not.toHaveBeenCalled()
	})

	it('should stop recording and stop all tracks', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()
		recorder.stop()

		expect(mockTrack.stop).toHaveBeenCalledOnce()
	})

	it('should not stop if already inactive', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		const mediaRecorder = recorder.getRecorder()
		const stopSpy = vi.spyOn(mediaRecorder, 'stop')

		recorder.stop()

		expect(stopSpy).not.toHaveBeenCalled()
	})

	it('should call onDataChunk callback when data is available', async () => {
		const mockBlob = new Blob(['test'], { type: 'audio/webm' })
		const onDataChunkMock = vi.fn()
		const callbacks: AudioRecorderCallbacks = {
			onDataChunk: onDataChunkMock,
		}

		const recorder = await createAudioRecorder(callbacks)
		const mediaRecorder = recorder.getRecorder()
		mediaRecorder.ondataavailable?.(
			new BlobEvent('dataavailable', { data: mockBlob })
		)

		expect(onDataChunkMock).toHaveBeenCalledWith(mockBlob)
	})

	it('should not call onDataChunk if data size is zero', async () => {
		const onDataChunkMock = vi.fn()
		const callbacks: AudioRecorderCallbacks = {
			onDataChunk: onDataChunkMock,
		}
		const recorder = await createAudioRecorder(callbacks)

		const emptyBlob = new Blob([], { type: 'audio/webm' })
		const mediaRecorder = recorder.getRecorder()
		mediaRecorder.ondataavailable?.(
			new BlobEvent('dataavailable', { data: emptyBlob })
		)

		expect(onDataChunkMock).not.toHaveBeenCalled()
	})

	it('should call onStop callback with final blob', async () => {
		const mockBlob = new Blob(['test'], { type: 'audio/webm' })
		const onStopMock = vi.fn()
		const callbacks: AudioRecorderCallbacks = {
			onStop: onStopMock,
		}

		const recorder = await createAudioRecorder(callbacks)
		const mediaRecorder = recorder.getRecorder()
		mediaRecorder.ondataavailable?.(
			new BlobEvent('dataavailable', { data: mockBlob })
		)
		mediaRecorder.onstop?.(new Event('stop'))

		expect(onStopMock).toHaveBeenCalledOnce()
		const calledBlob = onStopMock.mock.calls[0][0]
		expect(calledBlob).toBeInstanceOf(Blob)
		expect(calledBlob.type).toBe('audio/webm')
	})

	it('should call onError callback when error occurs', async () => {
		const mockError = new Error('Recording error')
		const onErrorMock = vi.fn()
		const callbacks: AudioRecorderCallbacks = {
			onError: onErrorMock,
		}

		const recorder = await createAudioRecorder(callbacks)
		const mediaRecorder = recorder.getRecorder()
		mediaRecorder.onerror?.(new ErrorEvent('error', { error: mockError }))

		expect(onErrorMock).toHaveBeenCalledWith(mockError)
	})

	it('should call onError with default error if error is undefined', async () => {
		const onErrorMock = vi.fn()
		const callbacks: AudioRecorderCallbacks = {
			onError: onErrorMock,
		}

		const recorder = await createAudioRecorder(callbacks)
		const mediaRecorder = recorder.getRecorder()
		mediaRecorder.onerror?.(new ErrorEvent('error', { error: undefined }))

		expect(onErrorMock).toHaveBeenCalledOnce()
		const calledError = onErrorMock.mock.calls[0][0]
		expect(calledError).toBeInstanceOf(Error)
		expect(calledError.message).toBe('MediaRecorder error')
	})

	it('should return INACTIVE state initially', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		expect(recorder.getState()).toBe('INACTIVE')
	})

	it('should return RECORDING state when recording', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()

		expect(recorder.getState()).toBe('RECORDING')
	})

	it('should return PAUSED state when paused', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		recorder.start()
		recorder.pause()

		expect(recorder.getState()).toBe('PAUSED')
	})

	it('should handle recording lifecycle from start to stop', async () => {
		const callbacks: AudioRecorderCallbacks = {}
		const recorder = await createAudioRecorder(callbacks)

		expect(recorder.getState()).toBe('INACTIVE')

		recorder.start()
		expect(recorder.getState()).toBe('RECORDING')

		recorder.pause()
		expect(recorder.getState()).toBe('PAUSED')

		recorder.resume()
		expect(recorder.getState()).toBe('RECORDING')

		recorder.stop()
		expect(recorder.getState()).toBe('INACTIVE')
	})
})
