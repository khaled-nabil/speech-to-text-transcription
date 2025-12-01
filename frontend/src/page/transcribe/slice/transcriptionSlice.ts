import { createSlice, type PayloadAction } from '@reduxjs/toolkit'
import type { Transcription } from 'types/transcription'

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
		updateTranscription: (state, action: PayloadAction<Transcription>) => {
			const index = state.items.findIndex(
				(t) => t.id === action.payload.id
			)
			if (index !== -1) {
				state.items[index] = action.payload
			}
		},
	},
})

export const { addTranscription, updateTranscription } =
	transcriptionSlice.actions

export default transcriptionSlice.reducer
