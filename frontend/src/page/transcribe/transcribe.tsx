import { useRef } from 'react'
import { Box } from '@mui/material'
import Lottie, { type LottieRefCurrentProps } from 'lottie-react'
import { useMutation } from '@tanstack/react-query'
import { useAppDispatch } from 'store/hooks'
import MicrophoneButton from '../../component/microphoneButton'
import audiWaveLottie from './assets/wave.json'
import { addTranscription } from 'page/transcribe/slice/transcriptionSlice'
import TranscriptionList from './component/transcriptionList'
import { fetchAPI } from 'utils/api'
import type { Transcription } from 'types/transcription'

import style from './transcribe.module.scss'

const Transcribe = () => {
	const dispatch = useAppDispatch()
	const lottieRef = useRef<LottieRefCurrentProps | null>(null)

	const transcribeMutation = useMutation<Transcription, Error, Blob>({
		mutationFn: async (blob) => {
			const formData = new FormData()
			formData.append('file', blob, 'recording.webm')

			return await fetchAPI<Transcription>('/api/v1/transcriber', {
				method: 'POST',
				body: formData,
			})
		},
		onSuccess: (data) => dispatch(addTranscription(data)),
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
