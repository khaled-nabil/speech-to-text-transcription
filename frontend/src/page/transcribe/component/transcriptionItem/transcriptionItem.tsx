import {
	Box,
	CardMedia,
	ListItemText,
	CircularProgress,
	Alert,
} from '@mui/material'
import { useQuery } from '@tanstack/react-query'
import { type FC, useEffect, useRef } from 'react'
import type { Transcription, TranscriptionResponse } from 'types/transcription'
import { useAppDispatch } from 'store/hooks'
import { updateTranscription } from 'page/transcribe/slice/transcriptionSlice'
import { fetchAPI } from 'utils/api'

import style from 'page/transcribe/transcribe.module.scss'

type TranscriptionItemProps = Transcription

const TranscriptionItem: FC<TranscriptionItemProps> = ({
	id,
	audioURL,
	transcriptText,
	status,
}) => {
	const dispatch = useAppDispatch()

	const audioPlayerRef = useRef<HTMLAudioElement | null>(null)

	useEffect(() => {
		if (!audioPlayerRef.current) return

		audioPlayerRef.current.src = audioURL
		audioPlayerRef.current.load()
	}, [audioURL])

	const { data, isLoading, isError, error } = useQuery<TranscriptionResponse>(
		{
			queryKey: ['transcriptionID', id],
			queryFn: () =>
				fetchAPI(`/api/v1/transcriptions/${id}`),
			enabled: status === 'PENDING',
			refetchInterval: 2500,
			refetchIntervalInBackground: true,
			retry: 15,
		}
	)

	useEffect(() => {
		if (data) {
			dispatch(updateTranscription(data))
		}
	}, [data, dispatch, id])

	return (
		<Box sx={{ width: '100%' }}>
			<CardMedia
				ref={audioPlayerRef}
				component="audio"
				controls
				className={style.audioPlayer}
			/>
			{status === 'SUCCESS' && (
				<ListItemText
					primary={transcriptText}
					sx={{ marginTop: '8px' }}
				/>
			)}
			{status === 'PENDING' ||
				(isLoading && (
					<Box
						sx={{
							display: 'flex',
							justifyContent: 'center',
							marginTop: '16px',
						}}
					>
						<CircularProgress />
					</Box>
				))}
			{status === 'ERROR' ||
				(isError && (
					<Alert severity="error" sx={{ marginTop: '8px' }}>
						Failed to transcribe audio. Please try again.
						{error ? ` Error: ${error.message}` : ''}
					</Alert>
				))}
		</Box>
	)
}

export default TranscriptionItem
