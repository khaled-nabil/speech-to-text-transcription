import { configureStore } from '@reduxjs/toolkit'
import userReducer from 'page/login/slice/userSlice'
import transcriptionReducer from 'page/transcribe/slice/transcriptionSlice'

export const store = configureStore({
	reducer: {
		user: userReducer,
		transcription: transcriptionReducer,
	},
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
