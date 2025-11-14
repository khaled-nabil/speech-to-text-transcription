import { useRef } from 'react'
import { Box, CardMedia } from '@mui/material'
import Lottie, { type LottieRefCurrentProps } from 'lottie-react'
import MicrophoneButton from '../../component/microphoneButton'
import audiWaveLottie from './assets/wave.json'

import style from './transcribe.module.scss'

const Transcribe = () => {
	const lottieRef = useRef<LottieRefCurrentProps | null>(null)
	const audioPlayerRef = useRef<HTMLAudioElement | null>(null)

	const handleLottiState = (active: boolean) => {
		if (lottieRef.current) {
			if (active) {
				lottieRef.current.play()
			} else {
				lottieRef.current.pause()
				lottieRef.current.goToAndStop(0, true)
			}
		}
	}

	const handleRecordingCompleted = (blob: Blob) => {
		if (!audioPlayerRef.current) return

		audioPlayerRef.current.src = URL.createObjectURL(blob)
		audioPlayerRef.current.load()
	}

	return (
		<Box className={style.container}>
			<CardMedia
				ref={audioPlayerRef}
				component="audio"
				controls
				className={style.audioPlayer}
			/>
			<Lottie
				lottieRef={lottieRef}
				animationData={audiWaveLottie}
				autoplay={false}
				loop={true}
				className={style.animation}
			/>
			<MicrophoneButton
				onRecordingComplete={handleRecordingCompleted}
				onRecording={handleLottiState}
			/>
		</Box>
	)
}

export default Transcribe
