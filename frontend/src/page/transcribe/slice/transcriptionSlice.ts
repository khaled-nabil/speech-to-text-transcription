import { createSlice, type PayloadAction } from '@reduxjs/toolkit'
import type { Transcription, TranscriptionResponse } from 'types/transcription'

interface TranscriptionState {
	items: Transcription[]
}

const initialState: TranscriptionState = {
	items: [],
}

const transcriptionSlice = createSlice({
	name: 'transcription',
	initialState,
	reducers: {
		addTranscription: (state, action: PayloadAction<Transcription>) => {
			state.items.push(action.payload)
		},
		updateTranscription: (
			state,
			action: PayloadAction<TranscriptionResponse>
		) => {
			const index = state.items.findIndex(
				(t) => t.id === action.payload.id
			)
			if (index !== -1) {
				const { audioURL } = state.items[index]

				state.items[index] = { ...action.payload, audioURL }
			}
		},
		clearTranscriptions: (state) => {
			state.items.forEach((t) => URL.revokeObjectURL(t.audioURL))

			state.items = []
		},
	},
})

export const { addTranscription, updateTranscription, clearTranscriptions } =
	transcriptionSlice.actions
export default transcriptionSlice.reducer
