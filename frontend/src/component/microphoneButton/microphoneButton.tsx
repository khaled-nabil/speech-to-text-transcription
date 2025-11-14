import { IconButton } from '@mui/material'
import MicIcon from '@mui/icons-material/Mic'
import StopIcon from '@mui/icons-material/Stop'
import PauseIcon from '@mui/icons-material/Pause'
import classnames from 'classnames'
import { useAudioRecorder } from './useAudioRecorder'

import style from './microphoneButton.module.scss'

interface MicrophoneButtonProps {
	onRecordingComplete: (blob: Blob) => void
	onRecording: (active: boolean) => void
}

const MicrophoneButton = ({
	onRecordingComplete,
	onRecording,
}: MicrophoneButtonProps) => {
	const { state, start, pause, resume, stop } =
		useAudioRecorder(onRecordingComplete)

	const handleRecordPause = () => {
		if (state === 'INACTIVE') {
			start()
			onRecording(true)
		} else if (state === 'RECORDING') {
			pause()
			onRecording(false)
		} else {
			resume()
			onRecording(true)
		}
	}

	const handleStop = () => {
		stop()
		onRecording(false)
	}

	return (
		<div className={style.container}>
			<IconButton onClick={handleRecordPause} className={style.button}>
				{state === 'RECORDING' ? (
					<PauseIcon sx={{ fontSize: 40 }} />
				) : (
					<MicIcon sx={{ fontSize: 40 }} />
				)}
			</IconButton>

			{state !== 'INACTIVE' && (
				<IconButton
					onClick={handleStop}
					className={classnames(style.button, style.stop)}
				>
					<StopIcon sx={{ fontSize: 40 }} />
				</IconButton>
			)}
		</div>
	)
}

export default MicrophoneButton
