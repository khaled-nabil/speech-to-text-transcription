import { ListItem } from '@mui/material'
import { useAppSelector } from 'store/hooks'
import TranscriptionItem from '../transcriptionItem'

import style from './transcriptionList.module.scss'

const TranscriptionList = () => {
	const transcriptions = useAppSelector(
		({ transcription }) => transcription.items
	)

	return (
		<ListItem className={style.list}>
			{transcriptions.map((transcription) => (
				<TranscriptionItem {...transcription} />
			))}
		</ListItem>
	)
}

export default TranscriptionList
