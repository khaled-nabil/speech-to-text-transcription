import { IconButton } from '@mui/material'
import MicIcon from '@mui/icons-material/Mic'

interface MicrophoneButtonProps {
	onClick: () => void
}

const MicrophoneButton = ({ onClick }: MicrophoneButtonProps) => {
	return (
		<IconButton
			onClick={onClick}
			sx={{
				width: 80,
				height: 80,
				backgroundColor: 'primary.main',
				color: 'white',
				'&:hover': {
					backgroundColor: 'primary.dark',
				},
			}}
		>
			<MicIcon sx={{ fontSize: 40 }} />
		</IconButton>
	)
}

export default MicrophoneButton
