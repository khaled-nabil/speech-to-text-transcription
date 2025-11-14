import { ListItem, ListItemText, Box, Typography } from '@mui/material'
import type {
	Transcription,
	TranscriptionReady,
	TranscriptionStatus,
} from 'types/transcription'

interface TranscriptionRowProps {
	transcription: Transcription
}

const getStatusColor = (status: TranscriptionStatus) => {
	switch (status) {
		case 'PENDING':
			return '#FFF59D'
		case 'ERROR':
			return '#EF5350'
		case 'SUCCESS':
			return '#81C784'
		default:
			return 'transparent'
	}
}

const TranscriptionRow = ({ transcription }: TranscriptionRowProps) => {
	if (transcription.status !== 'SUCCESS') {
		return (
			<ListItem
				sx={{
					backgroundColor: getStatusColor(transcription.status),
					mb: 1,
					borderRadius: 1,
				}}
			>
				<Box sx={{ width: '100%' }}>
					<ListItemText
						primary={
							transcription.transcriptText || 'Processing...'
						}
						secondary={`Status: ${transcription.status}`}
					/>
				</Box>
			</ListItem>
		)
	}

	const formattedDate = new Date(
		(transcription as TranscriptionReady).uploadDate
	).toLocaleString()

	return (
		<ListItem
			sx={{
				backgroundColor: getStatusColor(transcription.status),
				mb: 1,
				borderRadius: 1,
			}}
		>
			<Box sx={{ width: '100%' }}>
				<Typography variant="caption" color="text.secondary">
					{formattedDate}
				</Typography>
				<ListItemText
					primary={transcription.transcriptText}
					secondary={`Status: ${transcription.status}`}
				/>
			</Box>
		</ListItem>
	)
}

export default TranscriptionRow
