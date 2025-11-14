// frontend/src/page/transcribe/transcribe.tsx
import { useRef } from 'react'
import { Box } from '@mui/material'
import Lottie, { type LottieRefCurrentProps } from 'lottie-react'
import { useMutation } from '@tanstack/react-query'
import { useAppDispatch, useAppSelector } from 'store/hooks'
import MicrophoneButton from '../../component/microphoneButton'
import audiWaveLottie from './assets/wave.json'
import { addTranscription } from 'page/transcribe/slice/transcriptionSlice'
import TranscriptionList from './component/transcriptionList'
import { fetchAPI } from 'utils/api'
import type { TranscriptionResponse } from 'types/transcription'
import { createNewTranscriptionFromBlob } from './slice/domainTransformer'

import style from './transcribe.module.scss'

const Transcribe = () => {
	const dispatch = useAppDispatch()
	const userID = useAppSelector(({ user }) => user.userId)
	const lottieRef = useRef<LottieRefCurrentProps | null>(null)

	const transcribeMutation = useMutation<TranscriptionResponse, Error, Blob>({
		mutationFn: async (blob) => {
			const formData = new FormData()
			formData.append('file', blob, 'recording.webm')

			return await fetchAPI<TranscriptionResponse>(
				'/api/v1/transcriber',
				{
					method: 'POST',
					headers: {
						'X-User-ID': userID ?? '231',
					},
					body: formData,
				}
			)
		},
		onSuccess: (data, recordedBlob) => {
			const transcription = createNewTranscriptionFromBlob(
				data,
				recordedBlob
			)
			dispatch(addTranscription(transcription))
		},
	})

	const handleLottiState = (active: boolean) => {
		if (!lottieRef.current) return
		if (active) {
			lottieRef.current.play()
		} else {
			lottieRef.current.pause()
			lottieRef.current.goToAndStop(0, true)
		}
	}

	const handleRecordingCompleted = (blob: Blob) => {
		transcribeMutation.mutate(blob)
	}

	return (
		<Box className={style.container}>
			<TranscriptionList />
			<Lottie
				lottieRef={lottieRef}
				animationData={audiWaveLottie}
				autoplay={false}
				loop
				className={style.animation}
			/>
			<Box className={style.controls}>
				<MicrophoneButton
					onRecordingComplete={handleRecordingCompleted}
					onRecording={handleLottiState}
				/>
			</Box>
		</Box>
	)
}

export default Transcribe
