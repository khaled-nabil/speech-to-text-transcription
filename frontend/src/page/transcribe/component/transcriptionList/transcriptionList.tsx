import { ListItem } from '@mui/material'
import { useAppSelector } from 'store/hooks'
import TranscriptionItem from 'page/transcribe/component/transcriptionItem/transcriptionItem'

const TranscriptionList = () => {
	const transcriptions = useAppSelector(
		({ transcription }) => transcription.items
	)

	return (
		<ListItem>
			{transcriptions.map((transcription) => (
				<TranscriptionItem {...transcription} />
			))}
		</ListItem>
	)
}

export default TranscriptionList
