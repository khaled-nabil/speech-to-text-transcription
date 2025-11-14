import { ListItem, ListItemText, Box, Typography } from '@mui/material'
import type {
	TranscriptionReady,
	TranscriptionResponse,
} from 'types/transcription'

interface TranscriptionRowProps {
	transcription: TranscriptionResponse
}

const TranscriptionRow = ({ transcription }: TranscriptionRowProps) => {
	const formattedDate = new Date(
		(transcription as TranscriptionReady).uploadDate
	).toLocaleString()

	return (
		<ListItem>
			<Box>
				<Typography variant="caption" color="text.primary">
					{formattedDate}
				</Typography>
				<ListItemText
					primary={transcription.transcriptText}
				/>
			</Box>
		</ListItem>
	)
}

export default TranscriptionRow
