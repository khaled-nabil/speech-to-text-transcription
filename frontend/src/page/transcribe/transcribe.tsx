import { Box } from '@mui/material'
import Lottie, { type LottieRefCurrentProps } from 'lottie-react'
import MicrophoneButton from '../../component/microphoneButton'
import audiWaveLottie from './assets/wave.json'
import style from './transcribe.module.scss'
import { useEffect, useRef, useState } from 'react'

const Transcribe = () => {
	const [shouldPlay, setShouldPlay] = useState(false)
	const lottieRef = useRef<LottieRefCurrentProps | null>(null)
	useEffect(() => {
		const animation = lottieRef.current

		if (animation) {
			if (shouldPlay) {
				// 2. Play the animation
				animation.play()
			} else {
				// 3. Pause and jump to the first frame
				animation.pause()
				// Sets the animation progress to a specific frame number (e.g., 0 for frame 1)
				animation.goToAndStop(0, true)
			}
		}
	}, [shouldPlay]) // Rerun this effect whenever the 'shouldPlay' prop changes
	const handleMicrophoneClick = () => {
		// TODO: Implement microphone functionality
		setShouldPlay(!shouldPlay)
		console.log('Microphone clicked')
	}

	return (
		<Box className={style.container}>
			<Lottie
				lottieRef={lottieRef}
				animationData={audiWaveLottie}
				autoplay={false}
				loop={true}
				className={style.animation}
			/>
			<MicrophoneButton onClick={handleMicrophoneClick} />
		</Box>
	)
}

export default Transcribe
