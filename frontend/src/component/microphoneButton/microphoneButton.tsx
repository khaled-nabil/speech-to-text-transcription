import { IconButton } from '@mui/material'
import MicIcon from '@mui/icons-material/Mic'

import style from './microphoneButton.module.scss'

interface MicrophoneButtonProps {
	onClick: () => void
}

const MicrophoneButton = ({ onClick }: MicrophoneButtonProps) => {
	return (
		<IconButton
			onClick={onClick}
			className={style.button}
		>
			<MicIcon sx={{ fontSize: 40 }} />
		</IconButton>
	)
}

export default MicrophoneButton
