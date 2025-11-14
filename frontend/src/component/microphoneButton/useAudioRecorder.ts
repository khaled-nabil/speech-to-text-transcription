// frontend/src/component/microphoneButton/useAudioRecorder.ts
import { useEffect, useRef, useState, useCallback } from 'react'
import { createAudioRecorder } from './utils'
import type { AudioRecorder, RecordingState } from './types'

export const useAudioRecorder = (
	onComplete: (blob: Blob) => void,
	options?: { mimeType?: string; onChunk?(c: Blob): void }
) => {
	const [state, setState] = useState<RecordingState>('INACTIVE')
	const controllerRef = useRef<AudioRecorder | null>(null)

	useEffect(
		() => () => {
			controllerRef.current?.stop()
			controllerRef.current = null
		},
		[]
	)

	const start = useCallback(() => {
		if (state === 'RECORDING') return
		const initAndStart = async () => {
			if (!controllerRef.current) {
				controllerRef.current = await createAudioRecorder(
					{
						onDataChunk: options?.onChunk,
						onStop: (blob) => {
							onComplete(blob)
							setState('INACTIVE')
							controllerRef.current = null
						},
						onError: (err) => console.error(err),
					},
					options?.mimeType
				)
			}
			controllerRef.current.start()
			setState('RECORDING')
		}
		void initAndStart()
	}, [onComplete, options?.mimeType, options?.onChunk, state])

	const pause = useCallback(() => {
		if (state !== 'RECORDING') return
		controllerRef.current?.pause()
		setState('PAUSED')
	}, [state])

	const resume = useCallback(() => {
		if (state !== 'PAUSED') return
		controllerRef.current?.resume()
		setState('RECORDING')
	}, [state])

	const stop = useCallback(() => {
		if (state === 'INACTIVE') return
		controllerRef.current?.stop()
	}, [state])

	return { state, start, pause, resume, stop }
}
