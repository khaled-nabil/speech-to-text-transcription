import { Box, List } from '@mui/material'
import TranscriptionRow from 'component/transcriptionRow'
import type { Transcription } from 'types/transcription'
import type { FC } from 'react'

interface TranscriptionProps {
	items: Transcription[]
}

const Transcription: FC<TranscriptionProps> = ({ items }) => {
	return (
		<Box sx={{ p: 2 }}>
			<List>
				{items.map((transcription) => (
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
