import { Box, List } from '@mui/material'
import TranscriptionRow from '../../component/transcriptionRow'
import type { Transcription } from 'types/transcription.ts'

const dummyTranscriptions: Transcription[] = [
	{
		id: '1',
		uploadDate: '2025-01-10T10:30:00',
		transcriptText:
			'This is a successful transcription of a voice recording.',
		status: "SUCCESS",
	},
	{
		id: '2',
		uploadDate: '2025-01-11T14:15:00',
		transcriptText: 'This transcription is still being processed.',
		status: "PENDING",
	},
	{
		id: '3',
		uploadDate: '2025-01-12T09:00:00',
		transcriptText: 'This transcription failed due to an error.',
		status: "ERROR",
	},
	{
		id: '4',
		uploadDate: '2025-01-12T16:45:00',
		transcriptText:
			'Another successful transcription with more text content.',
		status: "SUCCESS",
	},
	{
		id: '5',
		uploadDate: '2025-01-13T11:20:00',
		transcriptText: 'Pending transcription waiting for processing.',
		status: "PENDING",
	},
]

const Transcription = () => {
	return (
		<Box sx={{ p: 2 }}>
			<List>
				{dummyTranscriptions.map((transcription) => (
					<TranscriptionRow
						key={transcription.id}
						transcription={transcription}
					/>
				))}
			</List>
		</Box>
	)
}

export default Transcription
