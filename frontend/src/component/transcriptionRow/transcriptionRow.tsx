import { ListItem, ListItemText, Box, Typography } from '@mui/material'
import type { Transcription } from 'types/transcription'

interface TranscriptionRowProps {
	transcription: Transcription
}

const TranscriptionRow = ({ transcription }: TranscriptionRowProps) => {
	const formattedDate = new Date(transcription.uploadDate).toLocaleString()

	return (
		<ListItem>
			<Box>
				<Typography variant="caption" color="text.primary">
					{formattedDate}
				</Typography>
				<ListItemText primary={transcription.transcriptText} />
			</Box>
		</ListItem>
	)
}

export default TranscriptionRow
