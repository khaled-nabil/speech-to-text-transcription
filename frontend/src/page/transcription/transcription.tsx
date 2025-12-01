import { Box, Alert, List, Backdrop, CircularProgress } from '@mui/material'
import TranscriptionRow from 'component/transcriptionRow'
import type { Transcription } from 'types/transcription'
import { useQuery } from '@tanstack/react-query'
import { fetchAPI } from 'utils/api'

import style from 'page/transcribe/transcribe.module.scss'

const Transcription = () => {
	const { data, isLoading, isError, error } = useQuery<Transcription[]>({
		queryKey: ['transcriptionList'],
		queryFn: () => fetchAPI(`/api/v1/transcriptions`),
	})

	return (
		<Box className={style.container}>
			<Backdrop
				sx={(theme) => ({
					color: '#fff',
					zIndex: theme.zIndex.drawer + 1,
				})}
				open={isLoading}
			>
				<CircularProgress color="primary" />
			</Backdrop>
			<List>
				{data
					?.filter(
						(t) =>
							t.status === 'SUCCESS' &&
							t.transcriptText &&
							t.transcriptText.trim().length > 0
					)
					.sort(
						(a, b) =>
							new Date(b.uploadDate).getTime() -
							new Date(a.uploadDate).getTime()
					)
					.map((transcription) => (
						<TranscriptionRow
							key={transcription.id}
							transcription={transcription}
						/>
					))}
			</List>
			{isError && <Alert severity="error">{error.message}</Alert>}
		</Box>
	)
}

export default Transcription
